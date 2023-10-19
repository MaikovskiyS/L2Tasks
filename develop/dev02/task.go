package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"unicode"
)

/*
=== Задача на распаковку ===

Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны, например:
  - "a4bc2d5e" => "aaaabccddddde"
  - abcd => "abcd"
  - "45" => "" (некорректная строка)
  - "" => ""

Дополнительное задание: поддержка escape - последовательностей
  - qwe\4\5 => qwe45 (*)
  - qwe\45 => qwe44444 (*)
  - qwe\\5 => qwe\\\\\ (*)

В случае если была передана некорректная строка функция должна возвращать ошибку. Написать unit-тесты.

Функция должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/
var ErrInvalidString = errors.New("invalid string")

func main() {
	ex := "a4bc2d5e"
	str, err := unpackingString(ex)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(str)
	os.Exit(0)
}
func unpackingString(s string) (string, error) {
	var bilder strings.Builder
	sr := []rune(s)
	var s2 string
	var n int
	var backslash bool

	for i, item := range sr {
		if unicode.IsDigit(item) && i == 0 {
			return "", ErrInvalidString
		}

		if unicode.IsDigit(item) && unicode.IsDigit(sr[i-1]) && sr[i-2] != '\\' {
			return "", ErrInvalidString
		}
		if item == '\\' && !backslash {
			backslash = true
			continue
		}
		if backslash && unicode.IsLetter(item) {
			return "", ErrInvalidString
		}
		if backslash {
			bilder.WriteRune(item)
			s2 += string(item)
			backslash = false
			continue
		}
		if unicode.IsDigit(item) {
			fmt.Println("item:", item)

			n = int(item - '0')
			fmt.Println("n:", n)
			if n == 0 {
				//bilder.WriteRune(sr[i-1])
				s2 = s2[:len(s2)-1]
				continue
			}

			for j := 0; j < n-1; j++ {
				bilder.WriteRune(sr[i-1])
				s2 += string(sr[i-1])
			}
			continue
		}

		s2 += string(item)
	}
	fmt.Println("bilder:", bilder.String())
	return s2, nil
}
