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

package dot

import (
	"testing"
)

type Foo struct {
	Number int
}

func (f *Foo) Dispose() {
	f.Number = 0
}

func TestExecutionBlock(t *testing.T) {
	val := &Foo{50}
	Using(val, func(v Disposable) {
		vFoo := v.(*Foo)
		if vFoo.Number != 50 {
			t.Errorf("The number should be '%d' but got '%d'",
				50, vFoo.Number)
		}
	})

	if val.Number != 0 {
		t.Errorf("The type was not disposed. Expected '%d' got '%d'",
			0, val.Number)
	}
}

func TestNestedExecutionBlock(t *testing.T) {
	val1 := &Foo{50}
	val2 := &Foo{40}
	val3 := &Foo{30}

	Using(val1, func(v1 Disposable) {
		v1Foo := v1.(*Foo)

		Using(val2, func(v2 Disposable) {
			v2Foo := v2.(*Foo)

			Using(val3, func(v3 Disposable) {
				v3Foo := v3.(*Foo)

				if v1Foo.Number != 50 {
					t.Errorf("The number should be '%d' but got '%d'",
						50, v1Foo.Number)
				}
				if v2Foo.Number != 40 {
					t.Errorf("The number should be '%d' but got '%d'",
						30, v2Foo.Number)
				}
				if v3Foo.Number != 30 {
					t.Errorf("The number should be '%d' but got '%d'",
						30, v3Foo.Number)
				}
			})
		})
	})

	if val1.Number != 0 {
		t.Errorf("The type was not disposed. Expected '%d' got '%d'",
			0, val1.Number)
	}
	if val2.Number != 0 {
		t.Errorf("The type was not disposed. Expected '%d' got '%d'",
			0, val2.Number)
	}
	if val3.Number != 0 {
		t.Errorf("The type was not disposed. Expected '%d' got '%d'",
			0, val3.Number)
	}
}
