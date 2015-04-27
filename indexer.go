package main

import (
	"index"
)

func main() {
	// 载入词典
	var segmenter sego.Segmenter
	segmenter.LoadDictionary("./sego/data/dictionary.txt")

	// 分词
	text := []byte("北京双安商场北土城奥森")
	segments := segmenter.Segment(text)

	// 处理分词结果
	// 支持普通模式和搜索模式两种分词，见代码中SegmentsToString函数的注释。
	println(sego.SegmentsToString(segments, false))
}
