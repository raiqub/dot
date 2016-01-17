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

var (
	counter = make([]int, 0)
)

type Bar int

func (f Bar) Dispose() {
	counter = append(counter, int(f))
}

func TestMulticastDispose(t *testing.T) {
	bar1 := Bar(1)
	bar2 := Bar(2)

	md := NewMulticastDispose()
	md.AddDisposable(bar1, bar2)
	md.Add(func() {
		counter = append(counter, 3)
	})
	md.Dispose()

	if len(counter) != 3 {
		t.Error("MulticastDispose was not disposed all registered functions")
	}

	for i, j := 0, 3; i < 3; i++ {
		if counter[i] != j {
			t.Errorf(
				"MulticastDispose was disposed out of order. Expected '%d' got '%d'",
				j, counter[i])
		}
		j--
	}
}
