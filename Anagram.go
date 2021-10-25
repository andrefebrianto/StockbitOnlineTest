package main

import "fmt"

type AnagramGroup struct {
	dictionary map[rune]int
	words      []string
}

func main() {
	texts := []string{"kita", "atik", "tika", "aku", "kia", "makan", "kua"}

	printOutAnagramGroups(getAnagramGroups(texts))
}

func printOutAnagramGroups(groups [][]string) {
	for row, words := range groups {
		if row == 0 {
			fmt.Println("[")
		}

		fmt.Printf(" [")
		for col, word := range words {
			if col == 0 {
				fmt.Printf(`"%s"`, word)
			} else {
				fmt.Printf(`, "%s"`, word)
			}
		}

		if row == len(groups)-1 {
			fmt.Printf("]\n]")
		} else {
			fmt.Printf("],\n")
		}
	}
}

func getAnagramGroups(words []string) [][]string {
	anagramGroups := make([]AnagramGroup, 0)

	for _, word := range words {
		dictionary := createDictionary(word)
		found := false

		for index, anagramGroup := range anagramGroups {
			if isAnagram(anagramGroup.dictionary, dictionary) {
				anagramGroups[index].words = append(anagramGroup.words, word)
				found = true
				break
			}
		}

		if !found {
			anagramGroups = append(anagramGroups, AnagramGroup{dictionary: dictionary, words: []string{word}})
		}

	}

	result := make([][]string, 0)

	for _, anagramGroup := range anagramGroups {
		result = append(result, anagramGroup.words)
	}

	return result
}

func createDictionary(word string) map[rune]int {
	dictionary := make(map[rune]int)

	for _, char := range word {
		dictionary[char]++
	}

	return dictionary
}

func isAnagram(sourceDicitonary, targetDictionary map[rune]int) bool {
	if len(sourceDicitonary) != len(targetDictionary) {
		return false
	}

	for key, value := range sourceDicitonary {
		if value != targetDictionary[key] {
			return false
		}
	}

	return true
}
