package main

func Id[T any](v T) T { return v }

func Eq[T comparable](a, b T) bool { return a == b }

type Box[T any] struct{ V T }

func NewBox[T any](v T) Box[T] {
	return Box[T]{V: v}
}

func (b Box[T]) Get() T { return b.V }

func (b *Box[T]) Set(v T) { b.V = v }

func main() {
	b := NewBox(int(1))
	n := Id[int](2)
	if !Eq(b.V, 1) || b.Get() != 1 {
		panic("eq")
	}
	b.Set(n)
	if b.Get() != 2 {
		panic("set")
	}
	println(b.V, n)
}
