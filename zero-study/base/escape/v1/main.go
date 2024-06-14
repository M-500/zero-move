package main

// @Description 分析内存逃逸的命令： go build -gcflags '-m -m -l'
// @Author 代码小学生王木木

func foo() *int {
	x := 1
	return &x
}

func main() {
	foo()
}
