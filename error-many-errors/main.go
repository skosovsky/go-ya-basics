package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
	"unicode/utf8"
)

var (
	ErrFoundNumbers = errors.New("found numbers")
	ErrLineTooLong  = errors.New("line is too long")
	ErrNoTwoSpace   = errors.New("no two space")
)

type CheckErrors []error

func (c *CheckErrors) Error() string {
	errs := make([]string, 0, len(*c))

	for _, err := range *c {
		errs = append(errs, err.Error())
	}

	return strings.Join(errs, "\n")
}

func MyCheck(input string) error {
	const maxLength = 20

	var errs CheckErrors

	if utf8.RuneCountInString(input) >= maxLength {
		errs = append(errs, fmt.Errorf("%w: %s", ErrLineTooLong, input))
	}

	if !strings.Contains(input, "  ") {
		errs = append(errs, fmt.Errorf("%w: %s", ErrNoTwoSpace, input))
	}

	for _, symbol := range input {
		if unicode.IsNumber(symbol) {
			errs = append(errs, fmt.Errorf("%w: %s", ErrFoundNumbers, string(symbol)))
		}
	}

	if len(errs) != 0 {
		return &errs
	}

	return nil
}

func Check(input string) error {
	var err error
	const maxLength = 20

	if utf8.RuneCountInString(input) >= maxLength {
		err = errors.Join(err, fmt.Errorf("%w: %s", ErrLineTooLong, input))
	}

	if !strings.Contains(input, "  ") {
		err = errors.Join(err, fmt.Errorf("%w: %s", ErrNoTwoSpace, input))
	}

	for _, symbol := range input {
		if unicode.IsNumber(symbol) {
			err = errors.Join(err, fmt.Errorf("%w: %s", ErrFoundNumbers, string(symbol)))

			break
		}
	}

	if err != nil {
		return err
	}

	return nil
}

func main() {
	for {
		fmt.Printf("Укажите строку (q для выхода): ") //nolint:forbidigo // example
		reader := bufio.NewReader(os.Stdin)
		ret, err := reader.ReadString('\n')

		if err != nil {
			log.Println(err)

			continue
		}

		ret = strings.TrimRight(ret, "\n")
		if ret == `q` {
			break
		}

		if err = Check(ret); err != nil {
			log.Println(err)

			continue
		}

		log.Println(`Строка прошла проверку`)
	}
}
