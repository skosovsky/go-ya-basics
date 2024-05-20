package main

import (
	"encoding/binary"
	"fmt"
	"io"
	"math/rand"
	"os"
	"time"
)

type generator struct {
	rnd rand.Source
}

func New(seed int64) io.Reader {
	return &generator{
		rnd: rand.NewSource(seed),
	}
}

func (g *generator) ReadOld(bytes []byte) (int, error) { //nolint:unparam // it's Read
	for i := range bytes {
		randInt := g.rnd.Int63()
		randByte := byte(randInt)
		bytes[i] = randByte
	}

	return len(bytes), nil
}

func (g *generator) Read(bytes []byte) (int, error) {
	for i := 0; i+8 < len(bytes); i += 8 {
		binary.LittleEndian.PutUint64(bytes[i:i+8], uint64(g.rnd.Int63()))
	}

	return len(bytes), nil
}

func (g *generator) Write(bytes []byte) (int, error) {
	return len(bytes), nil
}

func (g *generator) Dump(n int64, dst *os.File) error {
	_, err := io.CopyN(g, dst, n)
	if err != nil {
		return fmt.Errorf("copy failed: %w", err)
	}

	return nil
}

func main() {
	generatorer := New(time.Now().UnixNano())

	buf := make([]byte, 16) //nolint:mnd // example

	for range 5 {
		n, _ := generatorer.Read(buf)
		fmt.Printf("Genetare bytes: %v size(%d)\n", buf, n) //nolint:forbidigo // example
	}
}
