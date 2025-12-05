package main

import (
	"sort"
	"strconv"
	"strings"
)

func sortSlice(lines []string, column int, number, reverse, unique bool) []string {
	if len(lines) == 0 {
		return lines
	}

	sort.Slice(lines, func(i, j int) bool {
		a := lines[i]
		b := lines[j]

		keyA := getKey(a, column)
		keyB := getKey(b, column)

		var less bool

		if number {
			less = numLess(keyA, keyB)
		} else {
			less = keyA < keyB
		}

		if reverse {
			return !less
		}
		return less
	})

	if unique {
		lines = makeUnique(lines)
	}

	return lines
}

// ключ для сортировки
func getKey(line string, column int) string {
	if column <= 0 {
		return line
	}
	cols := strings.Split(line, "\t")
	idx := column - 1
	if idx >= 0 && idx < len(cols) {
		return cols[idx]
	}
	return ""
}

// сравнение строк как числа
func numLess(a, b string) bool {
	af, errA := strconv.ParseFloat(strings.TrimSpace(a), 64)
	bf, errB := strconv.ParseFloat(strings.TrimSpace(b), 64)

	if errA == nil && errB == nil {
		return af < bf
	}

	return a < b
}

// убираем дубли
func makeUnique(lines []string) []string {
	if len(lines) == 0 {
		return lines
	}

	res := []string{lines[0]}
	for i := 1; i < len(lines); i++ {
		if lines[i] != lines[i-1] {
			res = append(res, lines[i])
		}
	}
	return res
}
