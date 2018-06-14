package sego

import (
	"fmt"
	"testing"

	"github.com/issue9/assert"
)

/*
 * 作者:张晓明 时间:18/6/14
 */

var (
	strs = []Text{
		Text("one"),
		Text("two"),
		Text("three"),
		Text("four"),
		Text("five"),
		Text("six"),
		Text("seven"),
		Text("eight"),
		Text("nine"),
		Text("ten"),
	}
)

func Test_textSliceToString(t *testing.T) {
	a := textSliceToString(strs)
	b := Join(strs)
	assert.Equal(t, a, b)
}

func StringsJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Join(strs)
	}
}

func TextSliceToString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		textSliceToString(strs)
	}
}

func Test_Benchmark(t *testing.T) {
	fmt.Println("strings.Join:")
	fmt.Println(testing.Benchmark(StringsJoin))
	fmt.Println("textSliceToString")
	fmt.Println(testing.Benchmark(TextSliceToString))
}
