package list

import "testing"

func TestNewList(t *testing.T) {
	lis := New()

	key := []byte("my_list")
	//p1 := lis.LPush(key, []byte("11"))
	//t.Log(p1)

	lis.RPush(key, []byte("a"))
	//lis.LPush(key, []byte("b"))
	//lis.LPush(key, []byte("c"))
	//lis.LPush(key, []byte("d"))

	v1 := lis.LIndex(key, 0)
	t.Log(string(v1))

	v2 := lis.RPop(key)
	t.Log(string(v2))

	v3 := lis.RPop(key)
	t.Log(string(v3))
}
