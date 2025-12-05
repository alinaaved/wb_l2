package main

import (
	"fmt"

	"github.com/spf13/pflag"
)

func main() {
	column := pflag.IntP("column", "k", 0, "sort by column n")
	number := pflag.BoolP("number", "n", false, "sort by number")
	reverse := pflag.BoolP("reverse", "r", false, "reversed sort")
	unique := pflag.BoolP("unique", "u", false, "unique strings")

	pflag.Parse()

	args := pflag.Args()
	var text []string
	if len(args) == 0 {
		text = getLinesStdin()
	} else {
		fileName := args[0]
		text = getLinesFile(fileName)
	}

	sorted := sortSlice(text, *column, *number, *reverse, *unique)

	for _, line := range sorted {
		fmt.Println(line)
	}
}
