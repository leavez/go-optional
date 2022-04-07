package optional

type Type[T any] struct {
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

func New[T any](wrapped T) Type[T] {
	return Type[T]{wrapped: wrapped}
}

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
