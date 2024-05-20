package main

import (
	"io"
	"log"
)

type Hasher interface {
	io.Writer
	Hash() byte
}

type hash struct {
	result byte
}

func New(init byte) Hasher { //nolint:ireturn // example
	return &hash{result: init}
}

func (h *hash) Write(bytes []byte) (int, error) {
	for _, b := range bytes {
		h.result = (h.result^b)<<1 + b%2 //nolint:mnd // example
	}

	return len(bytes), nil
}

func (h *hash) Hash() byte {
	return h.result
}

func main() {
	buf := make([]byte, 16) //nolint:mnd // example

	hasher := New(0)
	_, err := hasher.Write(buf)
	if err != nil {
		return
	}
	log.Printf("Hash: %v \n", hasher.Hash())
}
