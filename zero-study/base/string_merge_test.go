package base

import (
	"bytes"
	"fmt"
	"math/rand"
	"strings"
	"testing"
)

// @Description
// @Author 代码小学生王木木
const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// randomString
//
//	@Description: 随机生成一个长度为n的字符串
//	@param n
//	@return string
func randomString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func normalMerge(n int, str string) string {
	s := ""
	for i := 0; i < n; i++ {
		s += str
	}
	return s
}

func useFmtPrintf(n int, str string) string {
	s := ""
	for i := 0; i < n; i++ {
		s = fmt.Sprintf("%s%s", s, str)
	}
	return s
}

func useStringsbuilder(n int, str string) string {
	var builder strings.Builder
	for i := 0; i < n; i++ {
		builder.WriteString(str)
	}
	return builder.String()
}

func useBytesBuilder(n int, str string) string {
	buf := new(bytes.Buffer)
	for i := 0; i < n; i++ {
		buf.WriteString(str)
	}
	return buf.String()
}

func benchmark(b *testing.B, fn func(int, string) string) {
	b.ResetTimer()
	s := randomString(15)
	for i := 0; i < b.N; i++ {
		fn(10000, s)
	}
}
func BenchmarkPlusConcat(b *testing.B)    { benchmark(b, normalMerge) }
func BenchmarkSprintfConcat(b *testing.B) { benchmark(b, useFmtPrintf) }
func BenchmarkBuilderConcat(b *testing.B) { benchmark(b, useStringsbuilder) }
func BenchmarkBufferConcat(b *testing.B)  { benchmark(b, useBytesBuilder) }

//func BenchmarkByteConcat(b *testing.B)    { benchmark(b, byteConcat) }
//func BenchmarkPreByteConcat(b *testing.B) { benchmark(b, preByteConcat) }
