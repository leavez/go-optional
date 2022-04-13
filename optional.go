package optional

import (
	"encoding/json"
	"errors"
)

type Type[T any] struct {
	/** Why we use an interface{} type here? Why not a *T?

	Both *T and interface{} type cloud represent a nil value. But if the real
	content of wrapped is a nil value of an interface type, we want optional.Type to
	be Nil. `*T` is hard to implement this behavior.
	*/
	wrapped interface{}
}

func (s Type[T]) IsNil() bool {
	return s.wrapped == nil
}

// ForceValue returns the wrapped content if not nil, or it will panic
func (s Type[T]) ForceValue() T {
	return s.wrapped.(T)
}

// Value is a golang style unwrapping method
func (s Type[T]) Value() (v T, ok bool) {
	if s.IsNil() {
		ok = false
	} else {
		ok = true
		v = s.ForceValue()
	}
	return
}

// ValueOrDefault return d if IsNil
func (s Type[T]) ValueOrDefault(d T) T {
	if w, ok := s.Value(); ok {
		return w
	} else {
		return d
	}
}

// ValueOrLazyDefault return f() if IsNil
func (s Type[T]) ValueOrLazyDefault(f func() T) T {
	if w, ok := s.Value(); ok {
		return w
	} else {
		return f()
	}
}

// --- initializer ---

// New returns an optional.Type. It doesn't always return a non-nil value. If
// wrapped is a nil value of interface type, it will return Nil
func New[T any](wrapped T) Type[T] {
	return Type[T]{wrapped: wrapped}
}

// Nil return a Nil optional.Type, whose IsNil() is true
func Nil[T any]() Type[T] {
	return Type[T]{}
}

func Compact[T any](wrapped Type[Type[T]]) Type[T] {
	if w, ok := wrapped.Value(); ok {
		return w
	} else {
		return Nil[T]()
	}
}

func Map[T, U any](v Type[T], f func(T) U) Type[U] {
	if w, ok := v.Value(); ok {
		return New(f(w))
	} else {
		return Nil[U]()
	}
}

// ---

func FromPtr[T any](wrapped *T) Type[T] {
	if wrapped == nil {
		return Nil[T]()
	} else {
		return New[T](*wrapped)
	}
}

// -- Marshaler and Unmarshaler ---

func (s Type[T]) MarshalJSON() ([]byte, error) {
	if w, ok := s.Value(); ok {
		return json.Marshal(w)
	} else {
		return json.Marshal(nil)
	}
}

func (s *Type[T]) UnmarshalJSON(data []byte) error {
	if s == nil {
		return errors.New("optional.Type: UnmarshalJSON on nil pointer")
	}
	var w *T
	err := json.Unmarshal(data, &w)
	if err != nil {
		return err
	}
	*s = FromPtr(w)
	return nil
}
