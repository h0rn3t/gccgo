package main

import "p"

func main() {
	var b p.Box[int]
	b.V = p.Id(1)
	n := p.Id[int](2)
	println(b.V, n)
}
