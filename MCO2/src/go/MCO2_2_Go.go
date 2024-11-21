package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
	"unicode"

	vicanso "github.com/vicanso/go-charts/v2"

	echarts_charts "github.com/go-echarts/go-echarts/v2/charts"
	echarts_compos "github.com/go-echarts/go-echarts/v2/components"
	echarts_option "github.com/go-echarts/go-echarts/v2/opts"
)

var file_name string
var total_words, unique_words int
var word_counts = make(map[string]int, 0)
var char_counts = make(map[string]int, 0)
var spec_counts = make(map[string]int, 0)
var stop_word_counts = make(map[string]int, 0)
var monthly_data = make(map[string]int, 0)

// each column in the csv file
var user_id, date_created, text []string

// time format
const time_format = "2006-01-02 15:06:05"

// initialize stop words
var stop_words = []string{
	"the", "be", "to", "of", "and", "a", "in", "that", "have", "I",
	"it", "for", "not", "on", "with",
}

type string_value struct {
	String string
	Value  int
}

func init() {
	for _, word := range stop_words {
		stop_word_counts[word] = 0
	}
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
			word_counts[word]++
			total_words++

			// Check if the word is a stop word and increment its count
			if _, exists := stop_word_counts[word]; exists {
				stop_word_counts[word]++
			}

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
	sorted_stop_words := sort_list(stop_word_counts)

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

	fmt.Println()

	fmt.Println("Stop Words (sorted by count)")
	for _, entry := range sorted_stop_words {
		fmt.Printf("%s: %d\n", entry.String, entry.Value)
	}
}

func write_file(buff []byte, filename string) error {
	err := os.WriteFile(filename, buff, 0600)
	if err != nil {
		fmt.Println(">>> PLOT_PIE : Cannot save chart")
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

	p, err := vicanso.BarRender(
		[][]float64{values_f},
		vicanso.XAxisDataOptionFunc(months_numeric),
		vicanso.LegendLabelsOptionFunc([]string{
			"Occurrences",
		}, vicanso.PositionRight),
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

	p, err := vicanso.PieRender(
		values_f,
		vicanso.TitleOptionFunc(vicanso.TitleOption{
			Text: "Occurence of Each Special Character",
			Left: vicanso.PositionCenter,
		}),
		vicanso.LegendOptionFunc(vicanso.LegendOption{
			Orient: vicanso.OrientVertical,
			Data:   specs,
			Left:   vicanso.PositionLeft,
		}),
		vicanso.PieSeriesShowLabel(),
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

func generate_plot_wc_data(data map[string]int) (items []echarts_option.WordCloudData) {
	items = make([]echarts_option.WordCloudData, 0)

	for String, Count := range data {
		items = append(items, echarts_option.WordCloudData{Name: String, Value: Count})
	}
	return
}

func plot_wc() *echarts_charts.WordCloud {
	wc := echarts_charts.NewWordCloud()
	wc.SetGlobalOptions(
		echarts_charts.WithTitleOpts(echarts_option.Title{
			Title: "Wordcloud of top 20 words",
		}))

	// FIXME: change this to top 20 words
	wc.AddSeries("wordcloud", generate_plot_wc_data(word_counts)).SetSeriesOptions(
		echarts_charts.WithWorldCloudChartOpts(
			echarts_option.WordCloudChart{SizeRange: []float32{14, 140}}))

	return wc
}

func make_plot_wc() {
	page := echarts_compos.NewPage()
	page.AddCharts(
		plot_wc(),
	)

	file, err := os.Create("go-wordcloud.html")
	if err != nil {
		fmt.Println(">>> WORD_CLOUD : Cannot create HTML file of wordcloud")
	}

	page.Render(io.MultiWriter(file))
}

func main() {
	fmt.Print("Enter filename: ")
	fmt.Scan(&file_name)
	// file_name = "fake_tweets.csv"

	data := get_csv(file_name)
	arrange_data(data)

	count_all()
	count_tweets_per_month()
	view_results()
	plot_bar(to_arrays(monthly_data))
	plot_pie(to_arrays(spec_counts))
	make_plot_wc()
}
