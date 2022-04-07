package optional

type Optional[T any] struct {
	wrapped interface{}
}

func (s Optional[T]) IsNil() bool {
	return s.wrapped == nil
}

// ForceValue returns the wrapped content if not nil, or it will panic
func (s Optional[T]) ForceValue() T {
	return s.wrapped.(T)
}

// Value is a golang style unwrapping method
func (s Optional[T]) Value() (v T, ok bool) {
	if s.IsNil() {
		ok = false
	} else {
		ok = true
		v = s.ForceValue()
	}
	return
}

// ValueOrDefault return d if IsNil() is true
func (s Optional[T]) ValueOrDefault(d T) T {
	if w, ok := s.Value(); ok {
		return w
	} else {
		return d
	}
}

func (s Optional[T]) ValueOrLazyDefault(f func() T) T {
	if w, ok := s.Value(); ok {
		return w
	} else {
		return f()
	}
}

// --- initializer ---

func New[T any](wrapped T) Optional[T] {
	return Optional[T]{wrapped: wrapped}
}

func Nil[T any]() Optional[T] {
	return Optional[T]{}
}

func Compact[T any](wrapped Optional[Optional[T]]) Optional[T] {
	if w, ok := wrapped.Value(); ok {
		return w
	} else {
		return Nil[T]()
	}
}

func Map[T, U any](v Optional[T], f func(T) U) Optional[U] {
	if w, ok := v.Value(); ok {
		return New(f(w))
	} else {
		return Nil[U]()
	}
}

// ---

func FromPtr[T any](wrapped *T) Optional[T] {
	if wrapped == nil {
		return Nil[T]()
	} else {
		return New[T](*wrapped)
	}
}
