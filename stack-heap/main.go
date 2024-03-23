package main

func main() {
	_ = stackIt()
	_ = stackIt2()

	y := 2
	_ = stackIt3(&y)
}

//go:noinline
func stackIt() int {
	y := 2
	return y * 2
}

//go:noinline
func stackIt2() *int {
	y := 2
	res := y * 2
	return &res
}

//go:noinline
func stackIt3(y *int) int {
	res := *y * 2
	return res
}
