// SPDX-License-Identifier: MIT

package assert

import (
	"errors"
	"fmt"
	"os"
	"path"
	"reflect"
	"runtime"
	"strconv"
	"strings"
	"testing"
)

// 定位错误信息的触发函数。输出格式为：TestXxx(xxx_test.go:17)。
func getCallerInfo() string {
	var info string

	for i := 0; ; i++ {
		pc, file, line, ok := runtime.Caller(i)
		if !ok {
			break
		}

		// 定位以 _test.go 结尾的文件。
		basename := path.Base(file)
		if !strings.HasSuffix(basename, "_test.go") {
			continue
		}

		// 定位函数名为 Test 开头的行。
		// 为什么要定位到 TestXxx 函数，是因为考虑以下情况：
		//  func isOK(val interface{}, t *testing.T) {
		//      // do somthing
		//      assert.True(t, val)  // (1
		//  }
		//
		//  func TestOK(t *testing.T) {
		//      isOK("123", t)       // (2
		//      isOK(123, t)         // (3
		//  }
		// 以上这段代码，定位到 (2、(3 的位置比总是定位到 (1 的位置更直观！
		funcName := runtime.FuncForPC(pc).Name()
		index := strings.LastIndex(funcName, ".Test")
		if -1 == index {
			index = strings.LastIndex(funcName, ".Benchmark")
			if index == -1 {
				continue
			}
		}
		funcName = funcName[index+1:]

		// Go1.5 之后的匿名函数为 TestA.func1
		// 包含以下几种情况：
		// 调用函数内的匿名函数；
		// 采用 go func(){} 的形式调用函数内的匿名函数；
		// 采用 go func(){} 的形式调用外部函数；
		//
		// 但是无法处理 go xx() 的情况，该情况直接开启一个新的堆栈信息，无法定位当前函数中的调用位置。
		if index := strings.IndexByte(funcName, '.'); index > -1 {
			funcName = funcName[:index]
			info = funcName + "(" + basename + ":" + strconv.Itoa(line) + ")"
			continue
		}

		info = funcName + "(" + basename + ":" + strconv.Itoa(line) + ")"
		break
	}

	if info == "" {
		info = "<无法获取调用者信息>"
	}
	return info
}

// 格式化错误提示信息
//
// msg1 中的所有参数将依次被传递给 fmt.Sprintf() 函数，
// 所以 msg1[0] 必须可以转换成 string(如:string, []byte, []rune, fmt.Stringer)
//
// msg2 参数格式与 msg1 完全相同，在 msg1 为空的情况下，会使用 msg2 的内容，
// 否则 msg2 不会启作用。
func formatMessage(msg1 []interface{}, msg2 []interface{}) string {
	msg := msg1
	if len(msg) == 0 {
		msg = msg2
	}

	if len(msg) == 0 {
		return "<未提供任何错误信息>"
	}

	if len(msg) == 1 {
		return fmt.Sprint(msg[0])
	}

	format := ""
	switch v := msg[0].(type) {
	case []byte:
		format = string(v)
	case []rune:
		format = string(v)
	case string:
		format = v
	case fmt.Stringer:
		format = v.String()
	default:
		return fmt.Sprintln(msg...)
	}

	return fmt.Sprintf(format, msg[1:]...)
}

// 当 expr 条件不成立时
//
// expr 返回结果值为bool类型的表达式；
// msg1,msg2 输出的错误信息，之所以提供两组信息，是方便在用户没有提供的情况下，
// 可以使用系统内部提供的信息，优先使用 msg1 中的信息，若不存在，则使用 msg2 的内容。
func assert(t testing.TB, expr bool, msg1 []interface{}, msg2 []interface{}) {
	if !expr {
		t.Error(formatMessage(msg1, msg2) + "@" + getCallerInfo())
	}
}

// True 断言表达式 expr 为 true
//
// args 对应 fmt.Printf() 函数中的参数，其中 args[0] 对应第一个参数 format，依次类推，
// 具体可参数 formatMessage() 函数的介绍。其它断言函数的 args 参数，功能与此相同。
func True(t testing.TB, expr bool, args ...interface{}) {
	assert(t, expr, args, []interface{}{"True 失败，实际值为 %#v", expr})
}

// False 断言表达式 expr 为 false
func False(t testing.TB, expr bool, args ...interface{}) {
	assert(t, !expr, args, []interface{}{"False 失败，实际值为 %#v", expr})
}

// Nil 断言表达式 expr 为 nil
func Nil(t testing.TB, expr interface{}, args ...interface{}) {
	assert(t, IsNil(expr), args, []interface{}{"Nil 失败，实际值为 %#v", expr})
}

// NotNil 断言表达式 expr 为非 nil 值
func NotNil(t testing.TB, expr interface{}, args ...interface{}) {
	assert(t, !IsNil(expr), args, []interface{}{"NotNil 失败，实际值为 %#v", expr})
}

// Equal 断言 v1 与 v2 两个值相等
func Equal(t testing.TB, v1, v2 interface{}, args ...interface{}) {
	assert(t, IsEqual(v1, v2), args, []interface{}{"Equal 失败，实际值为\nv1=%#v\nv2=%#v", v1, v2})
}

// NotEqual 断言 v1 与 v2 两个值不相等
func NotEqual(t testing.TB, v1, v2 interface{}, args ...interface{}) {
	assert(t, !IsEqual(v1, v2), args, []interface{}{"NotEqual 失败，实际值为\nv1=%#v\nv2=%#v", v1, v2})
}

// Empty 断言 expr 的值为空(nil,"",0,false)，否则输出错误信息
func Empty(t testing.TB, expr interface{}, args ...interface{}) {
	assert(t, IsEmpty(expr), args, []interface{}{"Empty 失败，实际值为 %#v", expr})
}

// NotEmpty 断言 expr 的值为非空(除 nil,"",0,false之外)，否则输出错误信息
func NotEmpty(t testing.TB, expr interface{}, args ...interface{}) {
	assert(t, !IsEmpty(expr), args, []interface{}{"NotEmpty 失败，实际值为 %#v", expr})
}

// Error 断言有错误发生
//
// 传递未初始化的 error 值(var err error = nil)，将断言失败
func Error(t testing.TB, expr interface{}, args ...interface{}) {
	if IsNil(expr) { // 空值，必定没有错误
		assert(t, false, args, []interface{}{"Error 失败，实际值为 Nil：[%T]", expr})
		return
	}

	_, ok := expr.(error)
	assert(t, ok, args, []interface{}{"Error 失败，实际类型为[%T]", expr})
}

// ErrorString 断言有错误发生且错误信息中包含指定的字符串 str
//
// 传递未初始化的 error 值(var err error = nil)，将断言失败
func ErrorString(t testing.TB, expr interface{}, str string, args ...interface{}) {
	if IsNil(expr) { // 空值，必定没有错误
		assert(t, false, args, []interface{}{"ErrorString 失败，实际值为 Nil：[%T]", expr})
		return
	}

	if err, ok := expr.(error); ok {
		index := strings.Index(err.Error(), str)
		assert(t, index >= 0, args, []interface{}{"Error 失败，实际类型为[%T]", expr})
	}
}

// ErrorType 断言有错误发生且错误的类型与 typ 的类型相同
//
// 传递未初始化的 error 值(var err error = nil)，将断言失败。
//
// 仅对 expr 是否与 typ 为同一类型作简单判断，如果要检测是否是包含关系，可以使用 errors.Is 检测。
func ErrorType(t testing.TB, expr interface{}, typ error, args ...interface{}) {
	if IsNil(expr) { // 空值，必定没有错误
		assert(t, false, args, []interface{}{"ErrorType 失败，实际值为 Nil：[%T]", expr})
		return
	}

	if _, ok := expr.(error); !ok {
		assert(t, false, args, []interface{}{"ErrorType 失败，实际类型为[%T]，且无法转换成 error 接口", expr})
		return
	}

	t1 := reflect.TypeOf(expr)
	t2 := reflect.TypeOf(typ)
	assert(t, t1 == t2, args, []interface{}{"ErrorType 失败，v1[%v]为一个错误类型，但与v2[%v]的类型不相同", t1, t2})
}

// NotError 断言没有错误发生
func NotError(t testing.TB, expr interface{}, args ...interface{}) {
	if IsNil(expr) { // 空值必定没有错误
		assert(t, true, args, []interface{}{"NotError 失败，实际类型为[%T]", expr})
		return
	}
	err, ok := expr.(error)
	assert(t, !ok, args, []interface{}{"NotError 失败，错误信息为[%v]", err})
}

// ErrorIs 断言 expr 为 target 类型
//
// 相当于 True(t, errors.Is(expr, target))
func ErrorIs(t testing.TB, expr interface{}, target error, args ...interface{}) {
	err, ok := expr.(error)
	assert(t, ok, args, []interface{}{"ErrorIs 失败，expr 无法转换成 error。"})

	assert(t, errors.Is(err, target), args, []interface{}{"ErrorIs 失败，expr 不是且不包含 target。"})
}

// FileExists 断言文件存在
func FileExists(t testing.TB, path string, args ...interface{}) {
	_, err := os.Stat(path)

	if err != nil && !os.IsExist(err) {
		assert(t, false, args, []interface{}{"FileExists 失败，且附带以下错误：%v", err})
	}
}

// FileNotExists 断言文件不存在
func FileNotExists(t testing.TB, path string, args ...interface{}) {
	_, err := os.Stat(path)

	if err == nil {
		assert(t, false, args, []interface{}{"FileNotExists 失败"})
	}
	if os.IsExist(err) {
		assert(t, false, args, []interface{}{"FileNotExists 失败，且返回以下错误信息：%v", err})
	}
}

// Panic 断言函数会发生 panic
func Panic(t testing.TB, fn func(), args ...interface{}) {
	has, _ := HasPanic(fn)
	assert(t, has, args, []interface{}{"并未发生 panic"})
}

// PanicString 断言函数会发生 panic 且 panic 信息中包含指定的字符串内容
func PanicString(t testing.TB, fn func(), str string, args ...interface{}) {
	if has, msg := HasPanic(fn); has {
		index := strings.Index(fmt.Sprint(msg), str)
		assert(t, index >= 0, args, []interface{}{"panic 中并未包含 %s", str})
		return
	}

	assert(t, false, args, []interface{}{"并未发生 panic"})
}

// PanicType 断言函数会发生 panic 且抛出指定的类型
func PanicType(t testing.TB, fn func(), typ interface{}, args ...interface{}) {
	has, msg := HasPanic(fn)
	if !has {
		return
	}

	t1 := reflect.TypeOf(msg)
	t2 := reflect.TypeOf(typ)
	assert(t, t1 == t2, args, []interface{}{"PanicType 失败，v1[%v]的类型与v2[%v]的类型不相同", t1, t2})

}

// NotPanic 断言函数不会发生 panic
func NotPanic(t testing.TB, fn func(), args ...interface{}) {
	has, msg := HasPanic(fn)
	assert(t, !has, args, []interface{}{"发生了 panic，其信息为[%v]", msg})
}

// Contains 断言 container 包含 item 的或是包含 item 中的所有项
//
// 具体函数说明可参考 IsContains()
func Contains(t testing.TB, container, item interface{}, args ...interface{}) {
	assert(t, IsContains(container, item), args,
		[]interface{}{"container:[%v]并未包含item[%v]", container, item})
}

// NotContains 断言 container 不包含 item 的或是不包含 item 中的所有项
func NotContains(t testing.TB, container, item interface{}, args ...interface{}) {
	assert(t, !IsContains(container, item), args,
		[]interface{}{"container:[%v]包含item[%v]", container, item})
}
