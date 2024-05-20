package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

type LimitedReader struct {
	io.Reader
	left int
}

func LimitReader(r io.Reader, n int) io.Reader {
	return &LimitedReader{
		Reader: r,
		left:   n,
	}
}

func (r *LimitedReader) Read(data []byte) (int, error) {
	if r.left == 0 {
		return 0, io.EOF
	}
	if r.left < len(data) {
		data = data[:r.left]
	}
	count, err := r.Reader.Read(data)
	if err != nil && errors.Is(err, io.EOF) {
		err = fmt.Errorf("read error: %w", err)
	}
	r.left -= count

	return count, err
}

func main() {
	r := strings.NewReader("some io.Reader stream to be read\n")
	lr := LimitReader(r, 4) //nolint:mnd // example

	_, err := io.Copy(os.Stdout, lr)
	if err != nil {
		log.Fatal(err)
	}
}
