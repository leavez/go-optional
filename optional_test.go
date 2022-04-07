package optional

import (
	"log"
	"testing"
)

func assert(t *testing.T, condition bool, message string, args ...any) {
	if !condition {
		t.Helper()
		t.Errorf(message, args...)
	}
}

func TestPrimitive(t *testing.T) {
	v := New[int](1)
	assert(t, !v.IsNil(), "IsNil failed")
	assert(t, v.ForceValue() == 1, "ForceValue failed")

	w, ok := v.Value()
	assert(t, ok && w == 1, "Value failed")

	v2 := Nil[int]()
	assert(t, v2.IsNil(), "IsNil failed")
}

func TestStruct(t *testing.T) {
	type S struct {
		A int
	}
	var v Optional[S] = New[S](S{A: 123})
	assert(t, !v.IsNil(), "IsNil failed")
	assert(t, v.ForceValue().A == 123, "ForceValue failed")

	v2 := Nil[S]()
	assert(t, v2.IsNil(), "IsNil failed")
}

type AAA interface {
	print() int
}
type aaa struct{ I int }

func (s aaa) print() int { return s.I }

type bbb struct{ I int }

func (s *bbb) print() int { return s.I }

func TestInterface(t *testing.T) {

	v := New[AAA](aaa{I: 123})
	assert(t, !v.IsNil(), "IsNil failed")
	assert(t, v.ForceValue().print() == 123, "ForceValue failed")

	v2 := Nil[AAA]()
	assert(t, v2.IsNil(), "IsNil failed")

	// ----
	v = New[AAA](&bbb{I: 123})
	assert(t, !v.IsNil(), "IsNil failed")
	assert(t, v.ForceValue().print() == 123, "ForceValue failed")

	// ----
	var inter AAA = nil
	v2 = New(inter)
	assert(t, v2.IsNil(), "should be nil")

	// ----
	func() {
		var inter AAA = nil
		v := New[AAA](inter)
		assert(t, v.IsNil(), "should be nil")
	}()

	func() {
		var ptr *bbb
		var inter AAA = ptr
		log.Println(inter, inter == nil)
		v := New[AAA](inter)
		// NOTE: this case is un-natual, but it works like assigning to an interface
		assert(t, !v.IsNil(), "should NOT be nil")
	}()
}

func TestPointer(t *testing.T) {
	var o = 123
	var ptr = &o

	var v2 = New(ptr)
	assert(t, !v2.IsNil(), "IsNil failed")
	assert(t, v2.ForceValue() == ptr, "ForceValue failed")

	// ----
	var v = FromPtr(ptr)
	assert(t, !v.IsNil(), "IsNil failed")
	assert(t, v.ForceValue() == 123, "ForceValue failed")
}

func TestNestedOptional(t *testing.T) {
	var v Optional[Optional[int]] = New(New(123))
	assert(t, !v.IsNil(), "IsNil failed")
	if w, ok := v.Value(); ok {
		assert(t, !w.IsNil(), "IsNil failed")
		assert(t, w.ForceValue() == 123, "ForceValue failed")
	}

	var v2 Optional[Optional[int]] = Nil[Optional[int]]()
	assert(t, v2.IsNil(), "IsNil failed")

	var v3 Optional[Optional[int]] = New(Nil[int]())
	assert(t, !v3.IsNil(), "IsNil failed")
}

func TestCompact(t *testing.T) {
	var v Optional[Optional[int]] = New(New(123))
	v2 := Compact(v)
	assert(t, !v2.IsNil(), "IsNil failed")
	assert(t, v2.ForceValue() == 123, "ForceValue failed")

	v = Nil[Optional[int]]()
	v2 = Compact(v)
	assert(t, v2.IsNil(), "IsNil failed")

	v = New(Nil[int]())
	v2 = Compact(v)
	assert(t, v2.IsNil(), "IsNil failed")
}

func TestMap(t *testing.T) {
	v := New[int](1)
	v2 := Map(v, func(t int) string {
		return "123"
	})
	assert(t, !v2.IsNil(), "IsNil failed")
	assert(t, v2.ForceValue() == "123", "ForceValue failed")
}

func TestValueOrDefault(t *testing.T) {
	v := New[int](1)
	assert(t, v.ValueOrDefault(10) == 1, "ValueOrDefault failed")
	assert(t, v.ValueOrLazyDefault(func() int { panic("don't call me") }) == 1, "ValueOrLazyDefault failed")
	v2 := Nil[int]()
	assert(t, v2.ValueOrDefault(10) == 10, "ValueOrDefault failed")
	assert(t, v2.ValueOrLazyDefault(func() int { return 100 }) == 100, "ValueOrLazyDefault failed")
}
