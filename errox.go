// Copyright (c) 2021 KongchengPro
// Errox is licensed under Mulan PSL v2.
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
	expCh chan Exception
}

func (t Thrower) ThrowError(err error) {
	t.expCh <- WrapError(err)
	runtime.Goexit()
}

func (t Thrower) ThrowErrox(exp Exception) {
	t.expCh <- exp
	runtime.Goexit()
}

type Tryer struct {
	tryBlock  func(Thrower)
	exception     Exception
	isCatched bool
}

func Try(tryBlock func(Thrower)) *Tryer {
	thrower := Thrower{expCh: make(chan Exception)}
	go tryBlock(thrower)
	exp := <-thrower.expCh
	return &Tryer{tryBlock: tryBlock, exception: exp}
}

func (t *Tryer) Catch(exceptionType string, catchBlock func(Exception)) *Tryer {
	if !t.isCatched && exceptionType == t.exception.Type() {
		catchBlock(t.exception)
		t.isCatched = true
	}
	return t
}

func NewException(exceptionType string) Exception {
	return ExceptionString{TypeString: exceptionType}
}

type Exception interface {
	Type() string
}

type ExceptionString struct {
	TypeString string
}

func (e ExceptionString) Type() string {
	return e.TypeString
}

func WrapError(err error) Exception {
	if err == nil {
		return nil
	}
	return NewException(err.Error())
}
