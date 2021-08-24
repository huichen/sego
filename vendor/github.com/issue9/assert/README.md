assert
[![Go](https://github.com/issue9/assert/workflows/Go/badge.svg)](https://github.com/issue9/assert/actions?query=workflow%3AGo)
[![codecov](https://codecov.io/gh/issue9/assert/branch/master/graph/badge.svg)](https://codecov.io/gh/issue9/assert)
[![license](https://img.shields.io/badge/license-MIT-brightgreen.svg?style=flat)](https://opensource.org/licenses/MIT)
======

assert 包是对 testing 的一个简单扩展，提供的一系列的断言函数，
方便在测试函数中使用：

```go
func TestA(t *testing.T) {
    v := true
    assert.True(v)

    a := assert.New(t)
    a.True(v)
}

// 也可以对 testing.B 使用
func Benchmark1(b *testing.B) {
    a := assert.New(b)
    v := false
    a.True(v)
    for(i:=0; i<b.N; i++) {
        // do something
    }
}

// 对 API 请求做测试，可以引用 assert/rest
func TestHTTP( t *testing.T) {
    a := assert.New(t)

    srv := rest.NewServer(a, h, nil)
    a.NotNil(srv)
    defer srv.Close()

    srv.NewRequest(http.MethodGet, "/body").
        Header("content-type", "application/json").
        Query("page", "5").
        JSONBody(&bodyTest{ID: 5}).
        Do().
        Status(http.StatusCreated).
        Header("content-type", "application/json;charset=utf-8").
        JSONBody(&bodyTest{ID: 6})
}
```

文档
----

[![Go Walker](https://gowalker.org/api/v1/badge)](https://gowalker.org/github.com/issue9/assert)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/issue9/assert)](https://pkg.go.dev/github.com/issue9/assert)

版权
----

本项目采用 [MIT](https://opensource.org/licenses/MIT) 开源授权许可证，完整的授权说明可在 [LICENSE](LICENSE) 文件中找到。
