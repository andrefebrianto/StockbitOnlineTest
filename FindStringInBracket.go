package main

import (
	"fmt"
	"strings"
)

func main() {
	text := "Anggota DPR (Dewan Perwakilan Rakyat)"

	fmt.Println(findFirstStringInBracket(text))
}

func findFirstStringInBracket(text string) string {
	indexOfFirstOpenBracket := strings.Index(text, "(")
	indexOfFirstCloseBracket := strings.Index(text, ")")

	if indexOfFirstOpenBracket < 0 || indexOfFirstCloseBracket < 0 || indexOfFirstCloseBracket <= indexOfFirstOpenBracket {
		return ""
	}

	return text[indexOfFirstOpenBracket+1 : indexOfFirstCloseBracket]
}
