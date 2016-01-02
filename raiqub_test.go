/*
 * Copyright 2015 Fabrício Godoy
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

func TestTrueForAll(t *testing.T) {
	samples := []string{
		"YQcWKe9b",
		"D08tbfRG",
		"7iaTwmh7",
		"k+oDBoWAiFOxD7pknX2kxDvQ+OX6HqwW0uqqvSVWRtU=",
	}

	var predicates []PredicateStringFunc

	predicates = append(predicates, TrueForAll)
	predicates = append(predicates, TrueForAll)
	predicates = append(predicates, TrueForAll)

	for _, s := range samples {
		for _, p := range predicates {
			if !p(s) {
				t.Fatalf("Error testing '%s' input", s)
			}
		}
	}
}
