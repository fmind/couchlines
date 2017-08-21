// Print one json document per line from the output of CouchDB "/db/_all_docs":

package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

// BEG marks the beginning of the document
const BEG = "\"doc\":"

// END marks the end of the document
const END = "}"

// Trim creates a substring with the content of the document.
// It removes all characters before the first BEG (included)
// It removes all characters after the last END (included)
func Trim(line string) string {
	beg := strings.Index(line, BEG)
	end := strings.LastIndex(line, END)

	if beg == -1 || end == -1 {
		return ""
	}

	return line[beg+len(BEG) : end]
}

func main() {
	var in string
	var err error
	reader := bufio.NewReader(os.Stdin)

	for {
		in, err = reader.ReadString('\r')
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		if out := Trim(in); out != "" {
			fmt.Println(out)
		}
	}
}
