// Copyright (c) 2021 KongchengPro
// Errox is licensed under Mulan PSL v2.
// You can use this software according to the terms and conditions of the Mulan PSL v2.
// You may obtain a copy of Mulan PSL v2 at:
//          http://license.coscl.org.cn/MulanPSL2
// THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND,
// EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT,
// MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
// See the Mulan PSL v2 for more details.

package errox_test

import (
	"fmt"
	"testing"

	. "github.com/kongchengpro/errox"
)

func TestA(t *testing.T) {
	Try(func(t *Thrower) {
		err := ReturnError()
		if err != nil {
			fmt.Println("fail!")
			t.ThrowError(err)
		}
		fmt.Println("success!")
	}).Catch("unknow error", func(erx Errox) {
		fmt.Println("catch A")
	}).Catch("some error", func(erx Errox) {
		fmt.Println("catch the error: ", erx)
	})
}

func ReturnError() error {
	return fmt.Errorf("some error")
}
