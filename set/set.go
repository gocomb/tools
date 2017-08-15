package set

import (
	"sync"
	"fmt"
	"bytes"
)

type HashSet struct {
	items map[interface{}]bool  //函数、map、切片皆为引用类型，不可hash，如果传入会报 unhashable type 错误
	mu sync.RWMutex //加锁，防止在多线程使用环境下，数据的不一致
}

type SetImpl interface {
	NewSet() *HashSet
	Insert(items ...interface{})
	Delete(items ...interface{})
	Has (item interface{}) bool
	Clear ()
	PopAny ()  (interface{},bool)
	Elements() []interface{}
	Len() int
	Difference(other *HashSet) *HashSet
	Union (other *HashSet) *HashSet
	Contains(other *HashSet) bool
	Intersection(other *HashSet) *HashSet
}

func NewSet() *HashSet{
	return  &HashSet{
		items:make(map[interface{}]bool),
	}
}

func (s *HashSet) Insert(items ...interface{}){
	s.mu.Lock()
	defer s.mu.Unlock()
	for _,item:=range items{
		s.items[item]= true
	}
}

func (s *HashSet) Delete(items ...interface{}){
	s.mu.Lock()
	defer s.mu.Unlock()
	for _,item:=range items {
		delete(s.items, item)
	}
}

func (s *HashSet) Clear(){
	s.mu.Lock()
	defer s.mu.Unlock()
	s.items = make(map[interface{}]bool)
}

func (s *HashSet) PopAny()(interface{},bool){
	s.mu.Lock()
	defer s.mu.Unlock()
	for item:=range s.items{
		delete(s.items,item)
		return item,true
	}
	return struct{}{},false
}

func (s *HashSet) Has(item interface{}) bool{
	return s.items[item]
}

//Elements生成一个可迭代的切片，便于迭代
func (s *HashSet) Elements() []interface{}{
	s.mu.RLock()
	defer s.mu.RUnlock()
	elements := make([]interface{},s.Len())
	index := 0
	for item,_:=range s.items{
		elements[index] = item
		index++
	}
	return elements
}

func (s *HashSet) Len() int{
	return len(s.items)
}


// Difference returns a set of objects that are not in other
func (s *HashSet) Difference(other *HashSet) *HashSet{
	s.mu.RLock()
	defer s.mu.RUnlock()
	if other.Len() == 0{
		return s
	}
	diff := NewSet()
	for item,_:=range s.items{
		if !other.Has(item){
			diff.Insert(item)
		}
	}
	return diff
}

//items in both a and b
func (s *HashSet)Intersection(other *HashSet) *HashSet{
	s.mu.RLock()
	defer s.mu.RUnlock()
	var a,b *HashSet
	result := NewSet()
	if s.Len()<other.Len(){
		a = s
		b = other
	}else {
		a = other
		b = s
	}
	for item,_:=range a.items{
		if b.Has(item){
			result.Insert(item)
		}
	}
	return result
}

func (s *HashSet) Union (other *HashSet) *HashSet{
	s.mu.RLock()
	defer s.mu.RUnlock()
	if other.Len() == 0{
		return s
	}
	if s.Len() == 0{
		return other
	}
	u := NewSet()
	for item,_:=range s.items{
		u.Insert(item)
	}
	for item,_:=range other.items{
		u.Insert(item)
	}
	return u
}


//whether super set
func (s *HashSet) Contains(other *HashSet) bool{
	s.mu.RLock()
	defer s.mu.RUnlock()
	if other.Len() == 0{
		return true
	} else if s.Len() == 0 || s.Len() < other.Len(){
		return false
	}
	for item,_:=range other.items{
		if !s.Has(item){
			return false
		}
	}
	return true
}

func (s *HashSet) String()string{
	s.mu.RLock()
	defer s.mu.RUnlock()
	var buf bytes.Buffer
	buf.WriteString("{")
	isFirst := true
	for item,_ := range s.items{
		if isFirst{
			isFirst = false
		}else {
			buf.WriteString(" ")
		}
		buf.WriteString(fmt.Sprintf("%v",item))
	}
	buf.WriteString("}")
	return buf.String()
}