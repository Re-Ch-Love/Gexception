// Copyright (c) 2021 KongchengPro
// Errox is licensed under Mulan PSL v2.
// You can use this software according to the terms and conditions of the Mulan PSL v2.
// You may obtain a copy of Mulan PSL v2 at:
//          http://license.coscl.org.cn/MulanPSL2
// THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND,
// EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT,
// MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
// See the Mulan PSL v2 for more details.

package errox

import "runtime"

type Thrower struct {
	waitCh chan Errox
}

func (t Thrower) ThrowError(err error) {
	t.waitCh <- Wrap(err)
	runtime.Goexit()
}

func (t Thrower) ThrowErrox(erx Errox) {
	t.waitCh <- erx
	runtime.Goexit()
}

type Tryer struct {
	tryBlock  func(*Thrower)
	errox     Errox
	isCatched bool
}

func Try(tryBlock func(*Thrower)) *Tryer {
	thrower := &Thrower{waitCh: make(chan Errox)}
	go tryBlock(thrower)
	erx := <-thrower.waitCh
	return &Tryer{tryBlock: tryBlock, errox: erx}
}

func (t *Tryer) Catch(erroxType string, catchBlock func(Errox)) *Tryer {
	if !t.isCatched && erroxType == t.errox.Type() {
		catchBlock(t.errox)
		t.isCatched = true
	}
	return t
}

func New(erroxType string) Errox {
	return ErroxString{TypeString: erroxType}
}

type Errox interface {
	Type() string
}

type ErroxString struct {
	TypeString string
}

func (e ErroxString) Type() string {
	return e.TypeString
}

func Wrap(err error) Errox {
	return New(err.Error())
}
