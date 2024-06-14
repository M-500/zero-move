package main

func demo() {
	s := make([]int, 1000000, 1000000)

	for index, _ := range s {
		s[index] = index
	}
}

func main() {
	demo()
}
