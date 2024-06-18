package main

import (
	"fmt"
	"sync"
)

// @Description go语言实现交叉打印（1. 两个goroutine，一个负责打印数字，另一个负责打印字符，最后的效果：12AB34CD56EF78GH910IJ）
// @Author 代码小学生王木木

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	numCh := make(chan struct{})
	charCh := make(chan struct{})
	go displayNumber(numCh, charCh, &wg)
	go displayWord(charCh, numCh, &wg)
	numCh <- struct{}{}
	wg.Wait()
}

func displayNumber(numCh <-chan struct{}, charCh chan<- struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 6; i++ {
		_ = <-numCh
		fmt.Println(i)
		charCh <- struct{}{}
	}
}

func displayWord(charCh <-chan struct{}, numCh chan<- struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 'A'; i <= 'F'; i++ {
		_ = <-charCh
		fmt.Println(string(i))
		numCh <- struct{}{}
	}
}
