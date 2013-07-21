package sego

import (
	"testing"
)

var (
	prodSeg = Segmenter{}
)

func TestSplit(t *testing.T) {
	expect(t, "中/国/有/十/三/亿/人/口/",
		bytesToString(splitTextToWords([]byte(
			"中国有十三亿人口"))))

	expect(t, "github/ /is/ /a/ /web/-/based/ /hosting/ /service/ /for/ /software/ /development/ /projects/",
		bytesToString(splitTextToWords([]byte(
			"GitHub is a web-based hosting service for software development projects"))))

	expect(t, "中/国/雅/虎/yahoo/!/ /china/致/力/于/领/先/的/公/益/民/生/门/户/网/站/",
		bytesToString(splitTextToWords([]byte(
			"中国雅虎Yahoo! China致力于领先的公益民生门户网站"))))
}

func TestSegment(t *testing.T) {
	var seg Segmenter
	seg.LoadDictionary("testdata/test_dict1.txt,testdata/test_dict2.txt")
	expect(t, "12", seg.dict.numTokens)
	segments := seg.Segment([]byte("中国有十三亿人口"))
	expect(t, "中国/p8 有/p3 十三亿/p11 人口/p12 ", SegmentsToString(segments, false))
	expect(t, "4", len(segments))
	expect(t, "0", segments[0].Start)
	expect(t, "6", segments[0].End)
	expect(t, "6", segments[1].Start)
	expect(t, "9", segments[1].End)
	expect(t, "9", segments[2].Start)
	expect(t, "18", segments[2].End)
	expect(t, "18", segments[3].Start)
	expect(t, "24", segments[3].End)
}

func TestLargeDictionary(t *testing.T) {
	prodSeg.LoadDictionary("data/dictionary.txt")
	expect(t, "中国/ns 人口/n ", SegmentsToString(prodSeg.Segment(
		[]byte("中国人口")), false))

	expect(t, "中国/ns 人口/n ", SegmentsToString(prodSeg.internalSegment(
		[]byte("中国人口"), false), false))

	expect(t, "中国/ns 人口/n ", SegmentsToString(prodSeg.internalSegment(
		[]byte("中国人口"), true), false))

	expect(t, "中华人民共和国/ns 中央人民政府/nt ", SegmentsToString(prodSeg.internalSegment(
		[]byte("中华人民共和国中央人民政府"), true), false))

	expect(t, "中华人民共和国中央人民政府/nt ", SegmentsToString(prodSeg.internalSegment(
		[]byte("中华人民共和国中央人民政府"), false), false))

	expect(t, "中华/nz 人民/n 共和/nz 共和国/ns 人民共和国/nt 中华人民共和国/ns 中央/n 人民/n 政府/n 人民政府/nt 中央人民政府/nt 中华人民共和国中央人民政府/nt ", SegmentsToString(prodSeg.Segment(
		[]byte("中华人民共和国中央人民政府")), true))
}
