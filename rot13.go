package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rot rot13Reader) Read(b []byte) (int, error) {
	cypher_text := make([]byte, 8)
	cypher_read, cypher_err := rot.r.Read(cypher_text)
	if cypher_err != nil {
		return 0, cypher_err
	}
	written := 0

	for ; written < cypher_read; written++ {
		text := cypher_text[written]

		var start byte = 0

		if text >= 'A' && text <= 'Z' {
			start = 'A'
		} else if text >= 'a' && text <= 'z' {
			start = 'a'
		}

		if start > 0 {
			b[written] = ((text - start + 13) % 26) + start
		} else {
			b[written] = text
		}
	}
	return written, nil
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
