// Copyright (c) 2021 KongchengPro
// Exception is licensed under Mulan PSL v2.
// You can use this software according to the terms and conditions of the Mulan PSL v2.
// You may obtain a copy of Mulan PSL v2 at:
//          http://license.coscl.org.cn/MulanPSL2
// THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND,
// EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT,
// MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
// See the Mulan PSL v2 for more details.

package tests

import (
	"fmt"
	"testing"
	"time"

	. "github.com/kongchengpro/gexception"
)

func TestA(t *testing.T) {
	var content string
	Try(func(t Thrower) {
		content = MockReadFile(t)
	}).Catch("unknow error", func(e Exception) {
		fmt.Println("handle unknow error")
		content = "[error]"
	})
	fmt.Println(content)
}

func MockReadFile(t Thrower) string {
	Try(func(t Thrower) {
		err := MockIO()
		if err != nil {
			fmt.Println("fail!")
			t.ThrowError(err)
		}
		fmt.Println("success!")
	}).Catch("unknow error", func(e Exception) {
		fmt.Println("unknow error, throw up!")
		t.ThrowException(e)
	}).Catch("IO error", func(e Exception) {
		fmt.Println("catch the error: ", e)
	})
	return "some data"
}

func MockIO() error {
	t := time.Now().UnixNano()
	n := t / 100 % 10 % 3
	fmt.Println(n)
	if n == 0 {
		return nil
	} else if n == 1 {
		return fmt.Errorf("IO error")
	} else {
		return fmt.Errorf("unknow error")
	}
}
