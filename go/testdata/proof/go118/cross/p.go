package p

func Id[T any](v T) T { return v }

type Box[T any] struct{ V T }
