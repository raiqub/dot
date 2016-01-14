/*
 * Copyright 2016 Fabrício Godoy
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package dot_test

import (
	"fmt"

	"github.com/raiqub/dot"
)

type Foo struct {
	Number int
	Text   string
}

func (f *Foo) Dispose() {
	f.Number = 0
	f.Text = ""

	fmt.Println("Foo disposed")
}

func NewFoo(n int, txt string) *Foo {
	return &Foo{
		n,
		txt,
	}
}

func Example_using() {
	dot.Using(NewFoo(50, "lorem"), func(d dot.Disposable) {
		foo := d.(*Foo)
		fmt.Printf("The number is '%d' and text is '%s'\n",
			foo.Number, foo.Text)
	})

	fmt.Println("Outside of execution block")

	// Output:
	// The number is '50' and text is 'lorem'
	// Foo disposed
	// Outside of execution block
}
