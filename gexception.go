// Copyright (c) 2021 KongchengPro
// Exception is licensed under Mulan PSL v2.
// You can use this software according to the terms and conditions of the Mulan PSL v2.
// You may obtain a copy of Mulan PSL v2 at:
//          http://license.coscl.org.cn/MulanPSL2
// THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND,
// EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT,
// MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
// See the Mulan PSL v2 for more details.

package exception

import "runtime"

type Thrower struct {
	exceptionCh chan Exception
}

func (t Thrower) ThrowError(err error) {
	t.exceptionCh <- WrapError(err)
	runtime.Goexit()
}

func (t Thrower) ThrowException(e Exception) {
	t.exceptionCh <- e
	runtime.Goexit()
}

type Tryer struct {
	tryBlock  func(Thrower)
	exception Exception
	isCatched bool
}

func Try(tryBlock func(Thrower)) *Tryer {
	thrower := Thrower{exceptionCh: make(chan Exception)}
	go tryBlock(thrower)
	e := <-thrower.exceptionCh
	return &Tryer{tryBlock: tryBlock, exception: e}
}

func (t *Tryer) Catch(exceptionType string, catchBlock func(Exception)) *Tryer {
	if !t.isCatched && exceptionType == t.exception.Type() {
		catchBlock(t.exception)
		t.isCatched = true
	}
	return t
}

var (
	BaseException_ = BaseException{
		FatherException_: nil,
		Type_:            "exception",
		Error_:           "nil",
	}
)

type Exception interface {
	FatherType() Exception
	Type() string
	Error() string
}

type BaseException struct {
	FatherException_ Exception
	Type_            string
	Error_           string
}

func (e *BaseException) FatherType() Exception {
	return e.FatherException_
}

func (e *BaseException) Type() string {
	return e.Type_
}

func (e *BaseException) Error() string {
	return e.Error_
}

func WrapError(err error) Exception {
	if err == nil {
		return nil
	}
	return &BaseException{
		FatherException_: &BaseException_,
		Type_:            err.Error(),
		Error_:           err.Error(),
	}
}

// 判断e2是否属于e1，如果e1是e2的父级（父级的父级等也算）
func Is(e1 Exception, e2 Exception) {
	var e Exception
	for {
		e = e.FatherType()
		if e.Type() == e1.Type() {
			break
		}
	}
}
