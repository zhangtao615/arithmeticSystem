package FIFO

import (
	"container/list"
	"fmt"
	"runtime"
)

// Cache interface

type Cache interface {
	// Set 通过key设置/更新一个值
	Set(key string, value interface{})
	// Get 通过key获取一个值
	Get(key string) interface{}
	// Del 通过key删除一个值
	Del(key string)
	// DelOldest 删除最就久没被访问的值
	DelOldest()
	// UseBytes 缓存中元素所占内存大小
	UseBytes() int
	// Size 缓存中元素数量
}

type Value interface {
	Len() int
}

// 定义fifo结构体
type fifo struct {
	maxBytes  int
	usedBytes int
	ll        *list.List
	cache     map[string]*list.Element
}

func (f *fifo) Set(key string, value interface{}) {
	if element, ok := f.cache[key]; ok {
		f.ll.MoveToFront(element)
		val := element.Value.(*entry)
		f.usedBytes = f.usedBytes - CalcLen(val.Len()) + CalcLen(value) // 更新占用内存大小
		element.Value = value
	} else {
		element := &entry{key, value}
		e := f.ll.PushFront(element)
		f.cache[key] = e
		f.usedBytes += element.Len()
	}

	for f.maxBytes > 0 && f.maxBytes < f.usedBytes {
		f.DelOldest()
	}
}

func (f *fifo) Get(key string) interface{} {
	if e, ok := f.cache[key]; ok {
		return e.Value.(*entry).value
	}

	return nil
}

func (f *fifo) Del(key string) {
	if e, ok := f.cache[key]; ok {
		f.removeElement(e)
	}
}

func (f *fifo) Len() int {
	return f.ll.Len()
}

func (f *fifo) Size() int {
	return f.ll.Len()
}

func (f *fifo) DelOldest() {
	f.removeElement(f.ll.Back())
}

func (f *fifo) removeElement(e *list.Element) {
	if e == nil {
		return
	}
	f.ll.Remove(e)
	en := e.Value.(*entry)
	f.usedBytes -= en.Len()
	delete(f.cache, en.key)
}

func (f *fifo) UseBytes() int {
	return f.usedBytes
}

// 定义key,value 结构
type entry struct {
	key   string
	value interface{}
}

// Len 计算出元素占用内存字节数
func (e *entry) Len() int {
	return CalcLen(e.value)
}

func CalcLen(value interface{}) int {
	var n int
	switch v := value.(type) {
	case Value:
		n = v.Len()
	case string:
		if runtime.GOARCH == "amd64" {
			n = 16 + len(v)
		} else {
			n = 8 + len(v)
		}
	case bool, int8, uint8:
		n = 1
	case int16, uint16:
		n = 2
	case int32, uint32, float32:
		n = 4
	case int64, uint64, float64:
		n = 8
	case int, uint:
		if runtime.GOARCH == "amd64" {
			n = 8
		} else {
			n = 4
		}
	case complex64:
		n = 8
	case complex128:
		n = 16
	default:
		panic(fmt.Sprintf("%T is not implement cache.value", value))
	}
	return n
}

func NewFIFOCache(maxBytes int) Cache {
	return &fifo{
		maxBytes: maxBytes,
		ll:       list.New(),
		cache:    make(map[string]*list.Element),
	}
}
