package main

import (
	"fmt"
	"io"
	"os"
)

func copyFile(srcFileName, dstFileName string) error {
	srcFile, err := os.Open(srcFileName)
	if err != nil {
		return fmt.Errorf("open source file: %w", err)
	}
	dstFile, err := os.Open(dstFileName)
	if err != nil {
		return fmt.Errorf("open destination file: %w", err)
	}
	n, err := io.Copy(dstFile, srcFile)
	if err != nil {
		return fmt.Errorf("copy: %w", err)
	}
	fmt.Printf("Copied %d bytes from %s to %s", n, srcFileName, dstFileName) //nolint:forbidigo // example

	return nil
}

func main() {
	err := copyFile("file1", "file2")
	if err != nil {
		return
	}
}
