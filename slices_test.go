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
	"strings"
	"testing"
)

const (
	SAMPLE_TEXT_MISSING = "Maecenas"
)

var (
	SAMPLE_TEXT_ARRAY = []string{
		"Lorem", "ipsum", "dolor", "sit", "amet,", "consectetur", "adipiscing",
		"elit.", "Sed", "tortor", "justo", "dui", "iaculis", "molestie.",
		"Integer.",
	}
)

func TestStringSliceIndexOf(t *testing.T) {
	sample := StringSlice(SAMPLE_TEXT_ARRAY)

	for i, item := range SAMPLE_TEXT_ARRAY {
		if retIdx := sample.IndexOf(item, false); retIdx != i {
			t.Errorf("Expected index '%d' but got '%d'", i, retIdx)
		}
		if retIdx := sample.IndexOf(strings.ToUpper(item), true); retIdx != i {
			t.Errorf("Expected index '%d' but got '%d'", i, retIdx)
		}
	}

	if sample.IndexOf(SAMPLE_TEXT_MISSING, false) != -1 {
		t.Errorf("The index of '%s' should be -1", SAMPLE_TEXT_MISSING)
	}

	missingUpper := strings.ToUpper(SAMPLE_TEXT_MISSING)
	if sample.IndexOf(missingUpper, true) != -1 {
		t.Errorf("The index of '%s' should be -1", missingUpper)
	}
}

func TestStringSliceExists(t *testing.T) {
	sample := StringSlice(SAMPLE_TEXT_ARRAY)

	for _, item := range SAMPLE_TEXT_ARRAY {
		if !sample.Exists(item, false) {
			t.Errorf("The text '%s' should be found", item)
		}
		itemUpper := strings.ToUpper(item)
		if !sample.Exists(itemUpper, true) {
			t.Errorf("The text '%s' should be found", itemUpper)
		}
	}

	if sample.Exists(SAMPLE_TEXT_MISSING, false) {
		t.Errorf("The text '%s' should not exists", SAMPLE_TEXT_MISSING)
	}

	missingUpper := strings.ToUpper(SAMPLE_TEXT_MISSING)
	if sample.Exists(missingUpper, true) {
		t.Errorf("The text '%s' should not exists", missingUpper)
	}
}

func TestStringSliceExistsAll(t *testing.T) {
	sample := StringSlice(SAMPLE_TEXT_ARRAY)
	testSample := make([]string, 6)
	copy(testSample, sample[2:8])

	if !sample.ExistsAll(testSample, false) {
		t.Error("All elements of specified sample should exists")
	}

	testSample = append(testSample, SAMPLE_TEXT_MISSING)
	if sample.ExistsAll(testSample, false) {
		t.Errorf("The element '%s' should not exists", SAMPLE_TEXT_MISSING)
	}
}

func TestStringSliceTrueForAll(t *testing.T) {
	sample := StringSlice(SAMPLE_TEXT_ARRAY)

	hasVowel := func(s string) bool {
		return strings.IndexAny(s, "aeiou") >= 0
	}

	if !sample.TrueForAll(hasVowel) {
		t.Error(
			"Every element of specified sample should have at least one vowel")
	}

	isBig := func(s string) bool {
		return len(s) > 50
	}
	if sample.TrueForAll(isBig) {
		t.Error(
			"None of elements of specified sample should have more than " +
				"50 characters")
	}
}
