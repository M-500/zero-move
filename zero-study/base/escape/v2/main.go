package main

// @Description
// @Author 代码小学生王木木
func demo() func() int {
	x := 1
	return func() int {
		return x
	}
}

func main() {
	demo()
}
