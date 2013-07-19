package sego

import (
	"testing"
)

func TestSegment(t *testing.T) {
	var seg Segmenter
	seg.LoadDictionary("testdata/test_dict1.txt,testdata/test_dict2.txt")
	expect(t, "12", seg.dict.numTokens)
	segments := seg.Segment([]byte("中国有十三亿人口"))
	expect(t, "中国/p8 有/p3 十三亿/p11 人口/p12", SegmentsToString(segments))
	expect(t, "4", len(segments))
	expect(t, "0", segments[0].Position)
	expect(t, "6", segments[1].Position)
	expect(t, "9", segments[2].Position)
	expect(t, "18", segments[3].Position)
}

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
