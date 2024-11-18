package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"

	"encoding/csv"
	"os"
	"sort"
	"time"
)

var file_name string
var total_words, unique_words int
var word_counts = make(map[string]int, 0)
var char_counts = make(map[string]int, 0)
var spec_counts = make(map[string]int, 0)
var monthly_data = make(map[string]int, 0)

// each column in the csv file
var user_id, date_created, text []string

// time format
const time_format = "2006-01-02 15:06:05"

type string_value struct {
	String string
	Value  int
}

func get_csv(file_name string) [][]string {
	file, err := os.Open(file_name)
	if err != nil {
		fmt.Println("File is not found.")
	}

	defer file.Close()

	reader := csv.NewReader(file)

	data, err := reader.ReadAll()
	if err != nil {
		return nil
	}

	no_head := data[1:]

	return no_head
}

func arrange_data(data [][]string) {
	for _, row := range data {
		user_id = append(user_id, row[1])
		date_created = append(date_created, row[2])
		text = append(text, row[3])
	}
}

func is_alphanum(r rune) bool {
	return unicode.IsLetter(r) || unicode.IsDigit(r)
}

func count_words() {
	for _, entry := range text {
		word_list := strings.Split(strings.ToLower(entry), " ") // split the words by " "
		for _, word := range word_list {
			word_counts[word] += 1
			total_words += 1

			for _, char := range word {
				if is_alphanum(char) {
					char_counts[string(char)]++
				} else {
					spec_counts[string(char)]++
				}
			}
		}
	}
	unique_words = len(word_counts)
}

func get_date(date_str string) time.Time {
	date, err := time.Parse(time_format, date_str)
	if err != nil {
		fmt.Println("Error parsing date")
		return time.Time{}
	}

	return date
}

func get_month(date time.Time) time.Month {
	return date.Month()
}

func count_tweets_per_month() {
	for _, entry := range date_created {
		date := get_date(entry)
		month := get_month(date)

		monthly_data[month.String()]++
	}
}

func sort_list(list map[string]int) []string_value {
	var sorted_data []string_value

	for s, v := range list {
		sorted_data = append(sorted_data, string_value{s, v})
	}

	sort.Slice(sorted_data, func(i, j int) bool {
		return sorted_data[i].Value > sorted_data[j].Value
	})

	return sorted_data
}

func get_top_n(list []string_value, n int) []string_value {
	top_n := list[:n]
	return top_n
}

func to_str(n int) string {
	return strconv.Itoa(n)
}

func view_results() {
	fmt.Println("--- Results ---")
	fmt.Println("Total Words: " + to_str(total_words))
	fmt.Println("Unique Words: " + to_str(unique_words))

	top_20_words := get_top_n(sort_list(word_counts), 20)
	top_10_specs := get_top_n(sort_list(spec_counts), 10)
	top_months := get_top_n(sort_list(monthly_data), 12)
	sorted_chars := sort_list(char_counts)
	sorted_specs := sort_list(spec_counts)
	sorted_words := sort_list(word_counts)

	fmt.Println()

	fmt.Println("Top 20 most frequent words")
	for _, entry := range top_20_words {
		fmt.Printf("%s: %d\n", entry.String, entry.Value)
	}

	fmt.Println()

	fmt.Println("Top 10 most frequent special characters")
	for _, entry := range top_10_specs {
		fmt.Printf("%s: %d\n", entry.String, entry.Value)
	}

	fmt.Println()

	fmt.Println("Tweets per Month")
	for _, entry := range top_months {
		fmt.Printf("%s: %d\n", entry.String, entry.Value)
	}

	fmt.Println()

	fmt.Println("All Characters (sorted by count)")
	for _, entry := range sorted_chars {
		fmt.Printf("%s: %d\n", entry.String, entry.Value)
	}

	fmt.Println()

	fmt.Println("All Special Characters (sorted by count)")
	for _, entry := range sorted_specs {
		fmt.Printf("%s: %d\n", entry.String, entry.Value)
	}

	fmt.Println()

	fmt.Println("All Words")
	for _, entry := range sorted_words {
		fmt.Printf("%s: %d\n", entry.String, entry.Value)
	}

}

func main() {
	fmt.Print("Enter filename: ")
	fmt.Scan(&file_name)

	data := get_csv(file_name)
	arrange_data(data)

	count_words()
	count_tweets_per_month()
	view_results()
}
