package lazywire

import "sync"

/* L is shorthand for Lazy */
type L[T any] struct {
	once sync.Once
	i    func() T
	v    T
}

func (l *L[T]) Get() T {
	l.once.Do(func() {
		l.v = l.i()
	})

	return l.v
}

func Lazy[T any](i func() T) *LR[T] {
	return &LR[T]{
		i: func() (T, error) { return i(), nil },
	}
}

/* LR is shorthand for LazyResult */
type LR[T any] struct {
	once sync.Once
	i    func() (T, error)
	v    T
	e    error
}

func LazyResult[T any](i func() (T, error)) *LR[T] {
	return &LR[T]{
		i: i,
	}
}

func (l *LR[T]) Get() (T, error) {
	l.once.Do(func() {
		l.v, l.e = l.i()
	})

	return l.v, l.e
}
