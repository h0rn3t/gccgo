package main

func Key[T comparable](v T) bool { return v == v }

type I interface{ M() }

func main() {
	var i I
	if !Key[I](i) {
		panic("iface")
	}
	var a any
	if !Key(a) {
		panic("any")
	}
	println("ok")
}
