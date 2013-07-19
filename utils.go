package sego

import (
	"fmt"
)

// 输出分词结果为下面的字符串格式
// 中国/ 有/ 十三亿/ 人口
func SegmentsToString(segs []Segment) (output string) {
	for i, seg := range segs {
		output += fmt.Sprintf("%s/%s", TextSliceToString(seg.Token.text), seg.Token.pos)
		if i != len(segs)-1 {
			output += " "
		}
	}
	return
}

// 将多个字元拼接一个字符串输出
func TextSliceToString(text []Text) string {
	var output string
	for _, word := range text {
		output += string(word)
	}
	return output
}

// 返回多个字元的字节总长度
func TextSliceByteLength(text []Text) (length int) {
	for _, word := range text {
		length += len(word)
	}
	return
}
