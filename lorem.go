// Package loremIpsum retrieves the classic Latin placeholder text from http://www.lipsum.com/.
package lorem

import (
	// "fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

// Function Lorem returns the classic Latin placeholder text from http://www.lipsum.com.
// Parameter amount is how many of 'what'.
// Parameter what is 'bytes', 'words', or 'paras'.
// Parameter startsWithLipson is a boolean to indicate to start at the beginning with Lipson or later in the text.
func Lorem(amount int, what string, startWithLipson bool) string {
	result := ""

	urlStr := "http://www.lipsum.com/feed/xml?amount=" + strconv.Itoa(amount) + "&what=" + what + "&start=" + strconv.FormatBool(startWithLipson)
	response, err := http.Get(urlStr)
	if err != nil {
		return ""
	} else {
		defer response.Body.Close()
		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return ""
		}
		result = string(contents)
	}

	// Result is XML.  No need to build a struct just look for the lipsum element.
	lines := strings.Split(string(result), "\n")
	for _, value := range lines {
		if strings.HasPrefix(value, "<lipsum>") {
			result = value[8 : len(value)-9]
			break
		}
	}

	return string(result)
}
