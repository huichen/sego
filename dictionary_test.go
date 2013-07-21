package sego

import (
	"fmt"
	"testing"
)

func TestBinarySort(t *testing.T) {
	e := []*node{
		&node{word: []byte("2")},
		&node{word: []byte("3")},
		&node{word: []byte("4")},
		&node{word: []byte("7")},
		&node{word: []byte("8")}}
	expect(t, "0 false", fmt.Sprint(binarySearch(e, []byte("1"))))
	expect(t, "0 true", fmt.Sprint(binarySearch(e, []byte("2"))))
	expect(t, "1 true", fmt.Sprint(binarySearch(e, []byte("3"))))
	expect(t, "3 false", fmt.Sprint(binarySearch(e, []byte("5"))))
	expect(t, "3 false", fmt.Sprint(binarySearch(e, []byte("6"))))
	expect(t, "3 true", fmt.Sprint(binarySearch(e, []byte("7"))))
	expect(t, "4 true", fmt.Sprint(binarySearch(e, []byte("8"))))
	expect(t, "5 false", fmt.Sprint(binarySearch(e, []byte("9"))))
}

func TestUnsert(t *testing.T) {
	e := []*node{
		&node{word: []byte("2")},
		&node{word: []byte("3")},
		&node{word: []byte("4")},
		&node{word: []byte("7")},
		&node{word: []byte("8")}}
	upsert(&e, []byte("1"))
	expect(t, "1 2 3 4 7 8 ", printNodes(e))
	upsert(&e, []byte("2"))
	expect(t, "1 2 3 4 7 8 ", printNodes(e))
	upsert(&e, []byte("3"))
	expect(t, "1 2 3 4 7 8 ", printNodes(e))
	upsert(&e, []byte("5"))
	expect(t, "1 2 3 4 5 7 8 ", printNodes(e))
	upsert(&e, []byte("9"))
	expect(t, "1 2 3 4 5 7 8 9 ", printNodes(e))
}

func TestToken(t *testing.T) {
	var dict Dictionary
	dict.addToken(&Token{text: toWords("1", "2", "3")})
	dict.addToken(&Token{text: toWords("1", "2", "3", "4")})
	dict.addToken(&Token{text: toWords("1", "2", "4")})
	expect(t, "(1(2(3.(4. ) 4. ) ) )", printDictionary(&dict))

	tokens := make([]*Token, 5)
	var numTokens int
	numTokens = dict.lookupTokens(toWords("1", "2", "4"), tokens)
	expect(t, "124 ", printTokens(tokens, numTokens))
	numTokens = dict.lookupTokens(toWords("1", "2", "3", "4"), tokens)
	expect(t, "123 1234 ", printTokens(tokens, numTokens))
	numTokens = dict.lookupTokens(toWords("1", "2", "3"), tokens)
	expect(t, "123 ", printTokens(tokens, numTokens))
	numTokens = dict.lookupTokens(toWords("1", "2", "7", "9"), tokens)
	expect(t, "", printTokens(tokens, numTokens))
}
