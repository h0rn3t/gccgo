package main

type I interface {
	M[T any]() // ERROR "interface method must have no type parameters"
}
