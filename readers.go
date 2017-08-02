package main

import "golang.org/x/tour/reader"

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.
func (r MyReader) Read(b []byte) (int, error) {
	to_write := len(b)
	written := 0
	for ; written < to_write; written++ {
		b[written] = 'A'
	}
	return written, nil
}

func main() {
	reader.Validate(MyReader{})
}
