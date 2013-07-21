package sego

import (
	"bytes"
)

// Dictionary结构体实现了一个字串前缀树，一个分词可能出现在叶子节点也有可能出现在非叶节点
type Dictionary struct {
	root           node     // 根节点
	maxTokenLength int      // 词典中最长的分词
	numTokens      int      // 词典中分词数目
	tokens         []*Token // 词典中所有的分词，方便遍历
	totalFrequency int64    // 词典中所有分词的频率之和
}

// 前缀树节点
type node struct {
	word     Text    // 该节点对应的字元
	token    *Token  // 当此节点没有对应的分词时值为nil
	children []*node // 该字元后继的所有可能字元，当为叶子节点时为空
}

// 词典中最长的分词
func (dict *Dictionary) MaxTokenLength() int {
	return dict.maxTokenLength
}

// 词典中分词数目
func (dict *Dictionary) NumTokens() int {
	return dict.numTokens
}

// 词典中所有分词的频率之和
func (dict *Dictionary) TotalFrequency() int64 {
	return dict.totalFrequency
}

// 向词典中加入一个分词
func (dict *Dictionary) addToken(token *Token) {
	current := &dict.root
	for _, word := range token.text {
		// 一边向深处移动一边添加节点（如果需要的话）
		current = upsert(&current.children, word)
	}

	// 当这个分词不存在词典中时添加此分词，否则忽略
	if current.token == nil {
		current.token = token
		if len(token.text) > dict.maxTokenLength {
			dict.maxTokenLength = len(token.text)
		}
		dict.numTokens++
		dict.tokens = append(dict.tokens, token)
		dict.totalFrequency += int64(token.frequency)
	}
}

// 在词典中查找和字元组words可以前缀匹配的所有分词
// 返回值为找到的分词数
func (dict *Dictionary) lookupTokens(words []Text, tokens []*Token) int {
	// 特殊情况
	if len(words) == 0 {
		return 0
	}

	current := &dict.root
	numTokens := 0
	for _, word := range words {
		// 如果已经抵达叶子节点则不再继续寻找
		if len(current.children) == 0 {
			break
		}

		// 否则在该节点子节点中进行下个字元的匹配
		index, found := binarySearch(current.children, word)
		if !found {
			break
		}

		// 匹配成功，则跳入匹配的子节点中
		current = current.children[index]
		if current.token != nil {
			tokens[numTokens] = current.token
			numTokens++
		}
	}
	return numTokens
}

// 二分法查找字元在子节点中的位置
// 如果查找成功，第一个返回参数为找到的位置，第二个返回参数为true
// 如果查找失败，第一个返回参数为应当插入的位置，第二个返回参数false
func binarySearch(nodes []*node, word Text) (int, bool) {
	start := 0
	end := len(nodes) - 1

	// 特例：
	if len(nodes) == 0 {
		// 当slice为空时，插入第一位置
		return 0, false
	}
	compareWithFirstWord := bytes.Compare(word, nodes[0].word)
	if compareWithFirstWord < 0 {
		// 当要查找的元素小于首元素时，插入第一位置
		return 0, false
	} else if compareWithFirstWord == 0 {
		// 当首元素等于node时
		return 0, true
	}
	compareWithLastWord := bytes.Compare(word, nodes[end].word)
	if compareWithLastWord == 0 {
		// 当尾元素等于node时
		return end, true
	} else if compareWithLastWord > 0 {
		// 当尾元素小于node时
		return end + 1, false
	}

	// 二分
	current := end / 2
	for end-start > 1 {
		compareWithCurrentWord := bytes.Compare(word, nodes[current].word)
		if compareWithCurrentWord == 0 {
			return current, true
		} else if compareWithCurrentWord < 0 {
			end = current
			current = (start + current) / 2
		} else {
			start = current
			current = (current + end) / 2
		}
	}
	return end, false
}

// 将字元加入节点数组中，并返回插入的节点指针
// 如果字元已经存在则返回存在的节点指针
func upsert(nodes *[]*node, word Text) *node {
	index, found := binarySearch(*nodes, word)
	if found {
		return (*nodes)[index]
	}
	*nodes = append(*nodes, nil)
	copy((*nodes)[index+1:], (*nodes)[index:])
	(*nodes)[index] = &node{word: word}
	return (*nodes)[index]
}
