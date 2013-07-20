package sego

// 字串类型，可以用来表达下面三种文本
//	1. 一个字元，比如"中"又如"国", 英文的一个字元是一个词
//	2. 一个分词，比如"中国"又如"人口"
//	3. 一段文字，比如"中国有十三亿人口"
type Text []byte

// 一个分词
type Token struct {
	text      []Text  // 分词的字串，这实际上是个字元的数组
	frequency int     // 分词在语料库中的词频
	distance  float32 // log2(总词频/该分词词频)，这相当于log2(1/p(分词))，用作动态规划中该分词的路径长度。求解prod(p(分词))的最大值相当于求解sum(distance(分词))的最小值，这就是“最短路径”的来历
	pos       string  // 词性标注
}

func (token *Token) Text() string {
	return TextSliceToString(token.text)
}

func (token *Token) Frequency() int {
	return token.frequency
}

func (token *Token) Pos() string {
	return token.pos
}
