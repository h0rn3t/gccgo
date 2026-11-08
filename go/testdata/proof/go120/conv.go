package main

func main() {
	s := []byte{1, 2, 3, 4}
	a := [2]byte(s[1:])
	if a[0] != 2 || a[1] != 3 {
		panic("array")
	}

	p := (*[2]byte)(s[1:])
	if p[0] != 2 || &p[0] != &s[1] {
		panic("ptr")
	}

	var n []byte
	if (*[0]byte)(n) != nil {
		panic("nil ptr")
	}
	if [0]byte(n) != [0]byte{} {
		panic("empty")
	}

	nval := 7
	boxes := []struct{ N *int }{{&nval}}
	ab := [1]struct{ N *int }(boxes)
	if ab[0].N != &nval || *ab[0].N != 7 {
		panic("struct")
	}

	func() {
		defer func() {
			if recover() == nil {
				panic("expected panic")
			}
		}()
		_ = [4]byte(s[:1])
	}()
	println("ok")
}
