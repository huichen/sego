sego
====

Go语言实现的中文分词程序

<a href="https://github.com/huichen/sego/blob/master/dictionary.go">词典</a>用前缀树实现，
<a href="https://github.com/huichen/sego/blob/master/segmenter.go">分词器</a>算法为最短路径加动态规划。

支持用户词典，支持词性标注，可以运行JSON格式的<a href="https://github.com/huichen/sego/blob/master/server/server.go">RPC服务器</a>。

分词速度<a href="https://github.com/huichen/sego/blob/master/tools/benchmark.go">单线程</a>2.7MB/s，<a href="https://github.com/huichen/sego/blob/master/tools/goroutines.go">goroutines并发</a>13MB/s, 处理器Core i7-3615QM 2.30GHz 8核。

在线演示地址 http://sego.weiboglass.com

# 安装/更新

```
go get -u github.com/huichen/sego
```

# 使用


```go
package main

import (
	"fmt"
	"github.com/huichen/sego"
)

func main() {
	// 载入词典
	var segmenter sego.Segmenter
	segmenter.LoadDictionary("github.com/huichen/sego/data/dictionary.txt")

	// 分词
	text := []byte("请在一米线以外等候")
	segments := segmenter.Segment(text)
  
	// 处理分词结果
	fmt.Println(sego.SegmentsToString(segments)) 
}
```
