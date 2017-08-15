package set

import (
	"testing"
)

func TestHashSet(t *testing.T){
	s := NewSet()
	s.Insert([3]string{"1","2","3"})
	s.Insert(true)
	for s.Len()>0{
		item,notempty := s.PopAny()
		if !notempty{
			t.Logf("empty")
		}
		t.Log(item)
	}

	s.Insert([3]string{"1","2","3"},2,2,2,2,false,true)
	//使用Elements生成切片，迭代更方便
	for _,element:=range s.Elements(){
		t.Log(element)
	}

	s.Delete(false,3)
	for _,element:=range s.Elements(){
		t.Log(element)
	}

	s2 := NewSet()
	s2.Insert(3,false,true,2)
	diff := s.Difference(s2)
	t.Log(diff)

	union := s.Union(s2)
	t.Log(union)

	inter := s.Intersection(s2)
	t.Log(inter)

	t.Log(s.Contains(s2))

	s.Clear()

	t.Log(s)
}