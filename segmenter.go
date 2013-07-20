package sego

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
	"unicode/utf8"
)

const (
	minTokenFrequency = 2 // 仅从字典文件中读取大于等于此频率的分词
)

// 分词器结构体
type Segmenter struct {
	dict Dictionary
}

// 文本中的一个分词
type Segment struct {
	// 分词在文本中的起始字节位置
	Position int

	// 分词信息
	Token *Token
}

// 该结构体用于记录Viterbi算法中某字元处的向前分词跳转信息
type Jumper struct {
	minDistance float32
	token       *Token
}

// 从文件中载入词典
//
// 可以载入多个词典文件，文件名用","分隔，排在前面的词典优先载入分词，比如
// 	"用户词典.txt,通用词典.txt"
// 当一个分词既出现在用户词典也出现在通用词典中，则优先使用用户词典。
//
// 词典的格式为（每个分词一行）
//	分词文本 频率 词性
func (seg *Segmenter) LoadDictionary(files string) {
	for _, file := range strings.Split(files, ",") {
		log.Printf("载入sego词典 %s", file)
		dictFile, err := os.Open(file)
		defer dictFile.Close()
		if err != nil {
			log.Fatalf("无法载入字典文件 \"%s\" \n", file)
		}

		reader := bufio.NewReader(dictFile)
		var text string
		var frequency int
		var pos string

		// 逐行读入分词
		for {
			size, _ := fmt.Fscanf(reader, "%s %d %s\n", &text, &frequency, &pos)

			// 文件结束
			if size == 0 {
				break
			}

			// 过滤频率太小的词
			if frequency < minTokenFrequency {
				continue
			}

			// 将分词添加到字典中
			words := splitTextToWords([]byte(text))
			token := Token{text: words, frequency: frequency, pos: pos}
			seg.dict.addToken(&token)
		}
	}

	// 计算每个分词的路径值，路径值含义见Token结构体的注释
	logTotalFrequency := float32(math.Log2(float64(seg.dict.totalFrequency)))
	for _, token := range seg.dict.tokens {
		token.distance = logTotalFrequency - float32(math.Log2(float64(token.frequency)))
	}
	log.Println("sego词典载入完毕")
}

// 对文本分词
//
// 输入参数：
//	bytes	UTF8文本的字节数组
//
// 输出：
//	[]Segment	划分的分词
func (seg *Segmenter) Segment(bytes []byte) []Segment {
	// 划分字元
	text := splitTextToWords(bytes)
	outputSegments := make([]Segment, len(text))

	// 处理特殊情况
	if len(text) == 0 {
		return outputSegments
	}

	// jumpers定义了每个字元处的向前跳转信息，包括这个跳转对应的分词，以及从文本段开始到该字元的最短路径值
	jumpers := make([]Jumper, len(text))
	jumpers[0] = Jumper{0, nil}

	start := 0
	current := start
	currentSeg := 0
	tokens := make([]*Token, seg.dict.maxTokenLength)
	for current < len(text) {
		max := start
		for current = start; current < len(text); current++ {
			// 在前一个字元处所有路径汇聚了，这表明前段文本已经完成动态搜索得到最短路径，此时跳出循环
			if current > max {
				break
			}

			// 找到前一个字元处的最短路径，以便计算后续路径值
			var baseDistance float32
			if current == start {
				// 当本字元在文本首部时，基础距离应该是零
				baseDistance = 0
			} else {
				baseDistance = jumpers[current-1].token.distance
			}

			// 寻找所有以当前字元开头的分词
			numTokens := seg.dict.lookupTokens(
				text[current:minInt(current+seg.dict.maxTokenLength, len(text))], tokens)

			// 对所有可能的分词，更新分词结束字元处的跳转信息
			for iToken := 0; iToken < numTokens; iToken++ {
				location := current + len(tokens[iToken].text) - 1
				updateJumper(&jumpers[location], baseDistance, tokens[iToken])

				// 更新这段文本所能覆盖的最远位置
				max = maxInt(max, current+len(tokens[iToken].text)-1)
			}

			// 当前字元没有对应分词时补加一个伪分词
			if numTokens == 0 || len(tokens[0].text) > 1 {
				updateJumper(&jumpers[current], baseDistance,
					&Token{text: []Text{text[current]}, frequency: 1, distance: 32, pos: "x"})
				max = maxInt(max, current)
			}
		}

		// 从后向前扫描第一遍得到需要添加的分词数目
		numSeg := 0
		for index := max; index >= start; {
			location := index - len(jumpers[index].token.text) + 1
			numSeg++
			index = location - 1
		}
		oldNumSeg := numSeg

		// 从后向前扫描第二遍添加分词到最终结果
		for index := max; index >= start; {
			location := index - len(jumpers[index].token.text) + 1
			numSeg--
			outputSegments[currentSeg+numSeg] = Segment{Position: 0, Token: jumpers[index].token}
			index = location - 1
		}
		currentSeg += oldNumSeg

		// 开始下一段文本
		start = max + 1
	}

	// 计算各个分词的字节位置
	bytePosition := 0
	for iSeg := 0; iSeg < currentSeg; iSeg++ {
		outputSegments[iSeg].Position = bytePosition
		bytePosition += TextSliceByteLength(outputSegments[iSeg].Token.text)
	}
	return outputSegments[:currentSeg]
}

// 更新跳转信息:
// 	1. 当该位置从未被访问过时(jumper.minDistance为零的情况)，或者
//	2. 当该位置的当前最短路径大于新的最短路径时
// 将当前位置的最短路径值更新为baseDistance加上新分词的概率
func updateJumper(jumper *Jumper, baseDistance float32, token *Token) {
	newDistance := baseDistance + token.distance
	if jumper.minDistance == 0 || jumper.minDistance > newDistance {
		jumper.minDistance = newDistance
		jumper.token = token
	}
}

// 取两整数较小值
func minInt(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// 取两整数较大值
func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 将文本划分成字元
func splitTextToWords(text Text) []Text {
	output := make([]Text, len(text))
	current := 0
	currentWord := 0
	inAlphanumeric := true
	alphanumericStart := 0
	for current < len(text) {
		_, size := utf8.DecodeRune(text[current:])
		if size == 1 &&
			(text[current] >= 'a' && text[current] <= 'z') ||
			(text[current] >= 'A' && text[current] <= 'Z') ||
			(text[current] >= '0' && text[current] <= '9') {
			// 当前是英文字母或者数字
			if !inAlphanumeric {
				alphanumericStart = current
				inAlphanumeric = true
			}
		} else {
			if inAlphanumeric {
				inAlphanumeric = false
				if current != 0 {
					output[currentWord] = toLower(text[alphanumericStart:current])
					currentWord++
				}
			}
			output[currentWord] = text[current : current+size]
			currentWord++
		}
		current += size
	}

	// 处理最后一个字元是英文的情况
	if inAlphanumeric {
		if current != 0 {
			output[currentWord] = toLower(text[alphanumericStart:current])
			currentWord++
		}
	}

	return output[:currentWord]
}

// 将英文词转化为小写
func toLower(text []byte) []byte {
	output := make([]byte, len(text))
	for i, t := range text {
		if t >= 'A' && t <= 'Z' {
			output[i] = t - 'A' + 'a'
		} else {
			output[i] = t
		}
	}
	return output
}
