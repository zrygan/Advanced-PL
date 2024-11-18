package main

import (
	"fmt"
	"image/color"
	"image/png"
	"strconv"
	"strings"
	"unicode"

	"encoding/csv"
	"os"
	"sort"
	"time"

	"github.com/vicanso/go-charts/v2"

	"github.com/psykhi/wordclouds"
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

func count_all() {
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

func to_arrays(list map[string]int) ([]string, []int) {
	var string_array []string
	var int_array []int

	for i, j := range list {
		string_array = append(string_array, i)
		int_array = append(int_array, j)
	}

	return string_array, int_array
}

func months_to_numeric(months []string) []string {
	month_map := map[string]string{
		"January":   "1",
		"February":  "2",
		"March":     "3",
		"April":     "4",
		"May":       "5",
		"June":      "6",
		"July":      "7",
		"August":    "8",
		"September": "9",
		"October":   "10",
		"November":  "11",
		"December":  "12",
	}

	var months_converted []string

	for _, month := range months {
		if num, ok := month_map[month]; ok {
			months_converted = append(months_converted, num)
		} else {
			fmt.Println("Error <months_to_numeric>, month not found")
		}
	}

	return months_converted
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

func write_file(buff []byte, filename string) error {
	err := os.WriteFile(filename, buff, 0600)
	if err != nil {
		fmt.Println(">>> PLOT_PIE : Cannot save bar chart")
		return err
	}

	return nil
}

func plot_bar(months []string, values []int) {
	var values_f []float64
	for _, val := range values {
		values_f = append(values_f, float64(val))
	}

	months_numeric := months_to_numeric(months)

	p, err := charts.BarRender(
		[][]float64{values_f},
		charts.XAxisDataOptionFunc(months_numeric),
		charts.LegendLabelsOptionFunc([]string{
			"Occurrences",
		}, charts.PositionRight),
	)

	if err != nil {
		fmt.Println(">>> PLOT_BAR : Cannot create bar chart")
	}

	buff, err := p.Bytes()
	if err != nil {
		fmt.Println(">>> PLOT_BAR : Cannot create bar chart")
	}

	err = write_file(buff, "go-bar.png")
	if err != nil {
		fmt.Println(">>> PLOT_BAR : Cannot create bar chart")
	}
}

func plot_pie(specs []string, values []int) {
	var values_f []float64
	for val := range values {
		values_f = append(values_f, float64(val))
	}

	p, err := charts.PieRender(
		values_f,
		charts.TitleOptionFunc(charts.TitleOption{
			Text: "Occurence of Each Special Character",
			Left: charts.PositionCenter,
		}),
		charts.LegendOptionFunc(charts.LegendOption{
			Orient: charts.OrientVertical,
			Data:   specs,
			Left:   charts.PositionLeft,
		}),
		charts.PieSeriesShowLabel(),
	)

	if err != nil {
		fmt.Println(">>> PLOT_PIE : Cannot create pie chart")
	}

	buff, err := p.Bytes()
	if err != nil {
		fmt.Println(">>> PLOT_PIE : Cannot create pie chart")
	}

	err = write_file(buff, "go-pie.png")
	if err != nil {
		fmt.Println(">>> PLOT_PIE : Cannot create pie chart")
	}
}

func plot_cloud(wordCounts map[string]int) {
	wc := wordclouds.NewWordcloud(wordCounts,
		wordclouds.FontMaxSize(100),
		wordclouds.FontMinSize(10),
		wordclouds.RandomPlacement(true),
		wordclouds.Width(800),
		wordclouds.Height(800),
		wordclouds.BackgroundColor(color.RGBA{255, 255, 255, 255}),
	)

	img := wc.Draw()

	outputFile, err := os.Create("go-wordcloud.png")
	if err != nil {
		fmt.Println(">>> PLOT_CLOUD : Cannot create word cloud")
	}
	defer outputFile.Close()

	err = png.Encode(outputFile, img)
	if err != nil {
		fmt.Println(">>> PLOT_CLOUD : Cannot create word cloud")
	}
}

func main() {
	// fmt.Print("Enter filename: ")
	// fmt.Scan(&file_name)
	file_name = "fake_tweets.csv"

	data := get_csv(file_name)
	arrange_data(data)

	count_all()
	count_tweets_per_month()
	view_results()
	plot_bar(to_arrays(monthly_data))
	plot_pie(to_arrays(spec_counts))
	plot_cloud(word_counts)
}
