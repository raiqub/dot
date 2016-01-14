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

type Bar int

func (f Bar) Dispose() {
	fmt.Printf("Bar(%d) disposed\n", int(f))
}

func Example_multicastDispose() {
	f1 := Bar(0)
	f2 := Bar(1)

	md := dot.NewMulticastDispose()
	md.AddDisposable(f1, f2)
	md.Add(func() {
		fmt.Println("Anonymous function called")
	})
	defer md.Dispose()

	// Output:
	// Anonymous function called
	// Bar(1) disposed
	// Bar(0) disposed
}
