// SPDX-License-Identifier: MIT

package assert

import "testing"

// Assertion 是对 testing.TB 进行了简单的封装。
// 可以以对象的方式调用包中的各个断言函数。
type Assertion struct {
	t testing.TB
}

// New 返回 Assertion 对象。
func New(t testing.TB) *Assertion {
	return &Assertion{t: t}
}

// TB 返回 testing.TB 接口
func (a *Assertion) TB() testing.TB {
	return a.t
}

// True 参照 assert.True() 函数
func (a *Assertion) True(expr bool, msg ...interface{}) *Assertion {
	True(a.t, expr, msg...)
	return a
}

// False 参照 assert.False() 函数
func (a *Assertion) False(expr bool, msg ...interface{}) *Assertion {
	False(a.t, expr, msg...)
	return a
}

// Nil 参照 assert.Nil() 函数
func (a *Assertion) Nil(expr interface{}, msg ...interface{}) *Assertion {
	Nil(a.t, expr, msg...)
	return a
}

// NotNil 参照 assert.NotNil() 函数
func (a *Assertion) NotNil(expr interface{}, msg ...interface{}) *Assertion {
	NotNil(a.t, expr, msg...)
	return a
}

// Equal 参照 assert.Equal() 函数
func (a *Assertion) Equal(v1, v2 interface{}, msg ...interface{}) *Assertion {
	Equal(a.t, v1, v2, msg...)
	return a
}

// NotEqual 参照 assert.NotEqual() 函数
func (a *Assertion) NotEqual(v1, v2 interface{}, msg ...interface{}) *Assertion {
	NotEqual(a.t, v1, v2, msg...)
	return a
}

// Empty 参照 assert.Empty() 函数
func (a *Assertion) Empty(expr interface{}, msg ...interface{}) *Assertion {
	Empty(a.t, expr, msg...)
	return a
}

// NotEmpty 参照 assert.NotEmpty() 函数
func (a *Assertion) NotEmpty(expr interface{}, msg ...interface{}) *Assertion {
	NotEmpty(a.t, expr, msg...)
	return a
}

// Error 参照 assert.Error() 函数
func (a *Assertion) Error(expr interface{}, msg ...interface{}) *Assertion {
	Error(a.t, expr, msg...)
	return a
}

// ErrorString 参照 assert.ErrorString() 函数
func (a *Assertion) ErrorString(expr interface{}, str string, msg ...interface{}) *Assertion {
	ErrorString(a.t, expr, str, msg...)
	return a
}

// ErrorType 参照 assert.ErrorType() 函数
func (a *Assertion) ErrorType(expr interface{}, typ error, msg ...interface{}) *Assertion {
	ErrorType(a.t, expr, typ, msg...)
	return a
}

// NotError 参照 assert.NotError() 函数
func (a *Assertion) NotError(expr interface{}, msg ...interface{}) *Assertion {
	NotError(a.t, expr, msg...)
	return a
}

// ErrorIs 断言 expr 为 target 类型
//
// 相当于 a.True(errors.Is(expr, target))
func (a *Assertion) ErrorIs(expr interface{}, target error, msg ...interface{}) *Assertion {
	ErrorIs(a.t, expr, target, msg...)
	return a
}

// FileExists 参照 assert.FileExists() 函数
func (a *Assertion) FileExists(path string, msg ...interface{}) *Assertion {
	FileExists(a.t, path, msg...)
	return a
}

// FileNotExists 参照 assert.FileNotExists() 函数
func (a *Assertion) FileNotExists(path string, msg ...interface{}) *Assertion {
	FileNotExists(a.t, path, msg...)
	return a
}

// Panic 参照 assert.Panic() 函数
func (a *Assertion) Panic(fn func(), msg ...interface{}) *Assertion {
	Panic(a.t, fn, msg...)
	return a
}

// PanicString 参照 assert.PanicString() 函数
func (a *Assertion) PanicString(fn func(), str string, msg ...interface{}) *Assertion {
	PanicString(a.t, fn, str, msg...)
	return a
}

// PanicType 参照 assert.PanicType() 函数
func (a *Assertion) PanicType(fn func(), typ interface{}, msg ...interface{}) *Assertion {
	PanicType(a.t, fn, typ, msg...)
	return a
}

// NotPanic 参照 assert.NotPanic() 函数
func (a *Assertion) NotPanic(fn func(), msg ...interface{}) *Assertion {
	NotPanic(a.t, fn, msg...)
	return a
}

// Contains 参照 assert.Contains() 函数
func (a *Assertion) Contains(container, item interface{}, msg ...interface{}) *Assertion {
	Contains(a.t, container, item, msg...)
	return a
}

// NotContains 参照 assert.NotContains() 函数
func (a *Assertion) NotContains(container, item interface{}, msg ...interface{}) *Assertion {
	NotContains(a.t, container, item, msg...)
	return a
}
