package main

import (
	"errors"
	"strconv"
)

var errorString = errors.New("invalid string")

func unpack(str string) (string, error) {
	//пустая строка -> пустая строка
	if len(str) == 0 {
		return "", nil
	}

	runeStr := []rune(str)
	resRune := []rune{}
	counter := []rune{}
	var letter rune

	//первая цифра подряд -> ошибка
	if runeStr[0] >= rune('0') && runeStr[0] <= '9' {
		return "", errorString
	}

	for i := 0; i < len(runeStr); {
		if runeStr[i] >= '0' && runeStr[i] <= '9' {
			return "", errorString
		}
		letter = runeStr[i]
		i++

		counter = counter[:0] //обнуляем
		for i < len(runeStr) && runeStr[i] >= '0' && runeStr[i] <= '9' {
			counter = append(counter, runeStr[i])
			i++
		}

		repeat := 1

		if len(counter) > 0 {
			numStr := string(counter)
			n, err := strconv.Atoi(numStr)
			if err != nil {
				return "", err
			}
			repeat = n
		}

		for k := 0; k < repeat; k++ {
			resRune = append(resRune, letter)
		}
	}
	return string(resRune), nil

}
