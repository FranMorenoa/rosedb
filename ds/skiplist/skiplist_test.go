package skiplist

import (
	"fmt"
	"testing"
)

type Employee struct {
	id   uint32
	name string
	age  uint8
}

func TestSkipList_Put(t *testing.T) {
	list := New()
	val := []byte("test_val")

	list.Put([]byte("ec"), val)
	list.Put([]byte("dc"), val)
	list.Put([]byte("ac"), val)
	list.Put([]byte("ae"), val)
	list.Put([]byte("bc"), val)
	list.Put([]byte("22"), val)
	list.Put([]byte("2"), val)
	list.Put([]byte("bc"), val)
	list.Put([]byte("xc"), val)
	list.Put([]byte("34"), val)
	list.Put([]byte("13"), val)

	e := list.Front()
	for p := e; p != nil; p = p.Next() {
		t.Logf("key = %+v, val = %+v", p.Key(), p.Value())
	}
}

func TestSkipList_Get(t *testing.T) {
	list := New()
	val := []byte("test_val")

	list.Put([]byte("ec"), val)
	list.Put([]byte("dc"), 123)
	list.Put([]byte("ac"), val)

	list.Put([]byte("111"), Employee{3330912, "mary", 24})

	t.Logf("%v \n", list.Get([]byte("ec")))
	t.Logf("%v \n", list.Get([]byte("ac")))
	t.Logf("%v \n", list.Get([]byte("111")))
}

func TestSkipList_Remove(t *testing.T) {
	list := New()
	val := []byte("test_val")

	list.Put([]byte("ec"), val)
	list.Put([]byte("dc"), 123)
	list.Put([]byte("ac"), val)

	t.Log(list.Size)
	list.Remove([]byte("dc"))
	list.Remove([]byte("ec"))
	list.Remove([]byte("ac"))
	t.Log(list.Size)
}

func TestSkipList_Foreach(t *testing.T) {
	list := New()
	val1 := []byte("test_val1")
	val2 := []byte("test_val2")
	val3 := []byte("test_val3")
	val4 := []byte("test_val4")

	list.Put([]byte("ec"), val1)
	list.Put([]byte("dc"), val2)
	list.Put([]byte("ac"), val3)
	list.Put([]byte("ae"), val4)

	keys := func(e *Element) bool {
		t.Logf("%s ", e.key)
		return false
	}

	list.Foreach(keys)

	vals := func(e *Element) bool {
		t.Logf("%s ", e.value)
		return true
	}

	list.Foreach(vals)
}

func TestSkipList_Foreach2(t *testing.T) {
	list := New()
	val := []byte("test_val")

	list.Put([]byte("ec"), val)
	list.Put([]byte("dc"), val)
	list.Put([]byte("ac"), val)
	list.Put([]byte("ae"), val)

	list.Foreach(func(e *Element) bool {
		e.value = []byte("test_val_002")
		return true
	})

	for p := list.Front(); p != nil; p = p.Next() {
		fmt.Printf("%s %s \n", string(p.Key()), string(p.Value().([]byte)))
	}
}
