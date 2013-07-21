package sego

import (
	"fmt"
	"testing"
)

func expect(t *testing.T, expect string, actual interface{}) {
	actualString := fmt.Sprint(actual)
	if expect != actualString {
		t.Errorf("期待值=\"%s\", 实际=\"%s\"", expect, actualString)
	}
}

func printNodes(nodes []*node) (output string) {
	for _, node := range nodes {
		output += fmt.Sprintf("%s ", node.word)
	}
	return
}

func getNodeString(node *node) string {
	output := string(node.word)
	if node.token != nil {
		output += "."
	}
	if node.children != nil {
		output += "("
		for _, c := range node.children {
			output += getNodeString(c) + " "
		}
		output += ")"
	}
	return output
}

func printDictionary(dict *Dictionary) string {
	return fmt.Sprint(getNodeString(&dict.root))
}

func printTokens(tokens []*Token, numTokens int) (output string) {
	for iToken := 0; iToken < numTokens; iToken++ {
		for _, word := range tokens[iToken].text {
			output += fmt.Sprint(string(word))
		}
		output += " "
	}
	return
}

func toWords(strings ...string) []Text {
	words := []Text{}
	for _, s := range strings {
		words = append(words, []byte(s))
	}
	return words
}

func bytesToString(bytes []Text) (output string) {
	for _, b := range bytes {
		output += (string(b) + "/")
	}
	return
}
