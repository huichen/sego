package main

import (
	"flag"
	"fmt"
	"github.com/liu-xuewen/sego"
)

var (
	text = flag.String("text", "中国互联网历史上最大的一笔并购案", "要分词的文本")
)

func main() {
	flag.Parse()

	var seg sego.Segmenter
	seg.LoadDictionary(dataFS, "../data/dictionary.txt")

	segments := seg.Segment([]byte(*text))
	fmt.Println(sego.SegmentsToString(segments, true))
}
