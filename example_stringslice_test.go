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
	"regexp"

	"github.com/raiqub/dot"
)

func isValidUrl(url string) bool {
	const pattern = `^(https?:\/\/)?([\da-z\.-]+)\.([a-z\.]{2,6})([\/\w \.-]*)*\/?$`
	r := regexp.MustCompile(pattern)
	return r.MatchString(url)
}

func Example_stringSlice() {
	list := []string{
		"https://www.google.com",
		"https://www.facebook.com",
		"http://www.youtube.com",
		"http://www.yahoo.com",
		"http://www.amazon.com",
		"https://en.wikipedia.org",
	}

	if dot.StringSlice(list).TrueForAll(isValidUrl) {
		fmt.Println("All URLs from list is valid")
	} else {
		fmt.Println("Not all URLs from list is valid")
	}

	if dot.StringSlice(list).IndexOf("HTTP://WWW.YAHOO.COM", true) == 3 {
		fmt.Println("Yahoo URL was found on 4th element of the list")
	} else {
		fmt.Println("Yahoo URL was not found")
	}

	if dot.StringSlice(list).Exists("https://www.facebook.com", false) {
		fmt.Println("Facebook URL was found")
	} else {
		fmt.Println("Facebook URL was not found")
	}

	// Output:
	// All URLs from list is valid
	// Yahoo URL was found on 4th element of the list
	// Facebook URL was found
}
