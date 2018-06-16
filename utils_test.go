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


func Test_Token_TextEquals(t *testing.T){
	token := Token{
		text:[]Text{
			[]byte("one"),
			[]byte("two"),
		},
	}
	assert.True(t,token.TextEquals("onetwo"))
}

func Test_Token_TextEquals_CN(t *testing.T){
	token := Token{
		text:[]Text{
			[]byte("中国"),
			[]byte("文字"),
		},
	}
	assert.True(t,token.TextEquals("中国文字"))
}

func Test_Token_TextNotEquals(t *testing.T){
	token := Token{
		text:[]Text{
			[]byte("one"),
			[]byte("two"),
		},
	}
	assert.False(t,token.TextEquals("one-two"))
}

func Test_Token_TextNotEquals_CN(t *testing.T){
	token := Token{
		text:[]Text{
			[]byte("中国"),
			[]byte("文字"),
		},
	}
	assert.False(t,token.TextEquals("中国文字1"))
}

func Test_Token_TextNotEquals_CN_B(t *testing.T){
	token := Token{
		text:[]Text{
			[]byte("中国"),
			[]byte("文字"),
		},
	}
	assert.False(t,token.TextEquals("中国文"))
}