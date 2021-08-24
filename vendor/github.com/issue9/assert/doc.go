// SPDX-License-Identifier: MIT

// Package assert 是对 testing 包的一些简单包装
//
// 提供了两种操作方式：直接调用包函数；或是使用 Assertion 对象。
// 两种方式完全等价，可以根据自己需要，选择一种。
//  func TestAssert(t *testing.T) {
//      var v interface{} = 5
//
//      // 直接调用包函数
//      assert.True(t, v == 5, "v的值[%v]不等于5", v)
//      assert.Equal(t, 5, v, "v的值[%v]不等于5", v)
//      assert.Nil(t, v)
//
//      // 以 Assertion 对象方式使用
//      a := assert.New(t)
//      a.True(v==5, "v的值[%v]不等于5", v)
//      a.Equal(5, v, "v的值[%v]不等于5", v)
//      a.Nil(v)
//      a.TB().Log("success")
//
//      // 以函数链的形式调用 Assertion 对象的方法
//      a.True(false).Equal(5,6)
//  }
//
//  // 也可以对 testing.B 使用
//  func Benchmark1(b *testing.B) {
//      a := assert.New(b)
//      a.True(false)
//      for(i:=0; i<b.N; i++) {
//          // do something
//      }
//  }
package assert
