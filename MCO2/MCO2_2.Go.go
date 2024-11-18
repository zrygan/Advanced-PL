/*
 * by: Zhean Robby L. Ganituen
 * Group 2, CSADPRG Major Course Output 2
 *
 * groupmates: 		Justin Ching and Juan Miranda
 * submitted to: 	Mr. Bautista Romualdo
 * date created:	November 17, 2024
 */

package main

import (
	"fmt"
	"time"

	"encoding/csv" // for working with CSV files
	"os"

	"regexp" // for regular expressions and string handling
	"strings"

	"sort" // for sorting algorithm

	"strconv" // for int to string conversion
)

type string_value struct {
	String string
	Value  int
}

type char_value struct {
	Char  rune
	Value int
}

func getCSV(fileName string) [][]string {
	file, err := os.Open(fileName)
	if err != nil {
		return nil
	}
	defer file.Close()

	fileReader := csv.NewReader(file)

	records, err := fileReader.ReadAll()
	if err != nil {
		return nil
	}

	return records
}

func corpus(records [][]string, wordCount *int, vocabSize *int, wordFrequency map[string]int, charFrequency map[rune]int, specialFrequency map[rune]int, monthlyData map[string]int) {
	alphanumeric := regexp.MustCompile(`^[a-zA-Z0-9]+$`)

	for _, row := range records {
		if len(row) > 1 && row[1] != "" {
			text := row[1]
			words := strings.Fields(strings.ToLower(text))

			for _, word := range words {
				wordFrequency[word]++
				*wordCount++

				// Count characters in word
				for _, char := range word {
					if alphanumeric.MatchString(string(char)) {
						charFrequency[char]++
					} else {
						specialFrequency[char]++
					}
				}
			}
		}

		if len(row) > 0 && row[0] != "" {
			dateString := row[0]
			date, err := time.Parse("2006-01-02 15:04:05", dateString)
			if err == nil {
				month := date.Format("January")
				monthlyData[month]++
			}
		}
	}

	*vocabSize = len(wordFrequency)
}

func getTopString(n int, list map[string]int) []string {
	var sorted []string_value
	var topN []string

	for key, value := range list {
		sorted = append(sorted, string_value{String: key, Value: value})
	}

	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i].Value > sorted[j].Value
	})

	for i := 0; i < n; i++ {
		topN = append(topN, sorted[i].String)
	}

	return topN
}

func getTopChar(n int, list map[rune]int) []rune {
	var sorted []char_value
	var topN []rune

	for key, value := range list {
		sorted = append(sorted, char_value{Char: key, Value: value})
	}

	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i].Value > sorted[j].Value
	})

	limit := n
	if len(sorted) < n {
		limit = len(sorted)
	}

	for i := 0; i < limit; i++ {
		topN = append(topN, sorted[i].Char)
	}

	return topN
}

func results(wordCount int, vocabSize int, wordFrequency map[string]int, charFrequency map[rune]int, specialFrequency map[rune]int, monthlyData map[string]int) {
	fmt.Println("########## Results ##########")
	fmt.Println("Total Words: " + strconv.Itoa(wordCount))
	fmt.Println("Unique Words: " + strconv.Itoa(vocabSize))

	fmt.Println("\nTop 20 most frequent words: ")
	topWords := getTopString(20, wordFrequency)
	for _, word := range topWords {
		fmt.Println(word)
	}

	fmt.Println("\nTop 10 most frequent special characters:")
	topSpecialChars := getTopChar(10, specialFrequency)
	for _, char := range topSpecialChars {
		fmt.Printf("%c\n", char)
	}

	fmt.Println("\nTweets per month:")
	for month, count := range monthlyData {
		fmt.Printf("%s: %d tweets\n", month, count)
	}

	fmt.Println("\nAll characters sorted by frequency:")
	var allChars []char_value
	for char, count := range charFrequency {
		allChars = append(allChars, char_value{Char: char, Value: count})
	}
	sort.Slice(allChars, func(i, j int) bool {
		return allChars[i].Value > allChars[j].Value
	})
	for _, char := range allChars {
		fmt.Printf("%c: %d\n", char.Char, char.Value)
	}

	fmt.Println("\nAll special characters sorted by frequency:")
	var allSpecialChars []char_value
	for char, count := range specialFrequency {
		allSpecialChars = append(allSpecialChars, char_value{Char: char, Value: count})
	}
	sort.Slice(allSpecialChars, func(i, j int) bool {
		return allSpecialChars[i].Value > allSpecialChars[j].Value
	})
	for _, char := range allSpecialChars {
		fmt.Printf("%c: %d\n", char.Char, char.Value)
	}

	fmt.Println("\nAll words sorted by frequency:")
	var allWords []string_value
	for word, count := range wordFrequency {
		allWords = append(allWords, string_value{String: word, Value: count})
	}
	sort.Slice(allWords, func(i, j int) bool {
		return allWords[i].Value > allWords[j].Value
	})
	for _, word := range allWords {
		fmt.Println(word.String, word.Value)
	}
}

func main() {
	var fileName string
	var wordCount, vocabSize int
	wordFrequency := make(map[string]int)
	charFrequency := make(map[rune]int)
	specialFrequency := make(map[rune]int)
	monthlyData := make(map[string]int)

	fmt.Print("Enter filename: ")
	fmt.Scan(&fileName)

	records := getCSV(fileName)
	if records == nil {
		fmt.Printf("File name %s is not found", fileName)
		return
	}

	corpus(records, &wordCount, &vocabSize, wordFrequency, charFrequency, specialFrequency, monthlyData)

	results(wordCount, vocabSize, wordFrequency, charFrequency, specialFrequency, monthlyData)
}
