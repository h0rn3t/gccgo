package main

func Key[T comparable](v T) bool { return v == v }

func main() {
	var s []byte
	_ = Key[[]byte](s) // ERROR "type argument is not comparable"
}
