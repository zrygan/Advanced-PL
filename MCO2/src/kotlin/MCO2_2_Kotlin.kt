/*
******************
Last names: Ching
Language: Kotlin
Paradigm(s): Functional, Imperative, Object-oriented, Procedural
******************
*/

import com.kennycason.kumo.CollisionMode
import java.awt.Color
import java.awt.Dimension
import java.io.File
import java.util.*
import com.kennycason.kumo.WordCloud
import com.kennycason.kumo.WordFrequency
import com.kennycason.kumo.bg.CircleBackground
import com.kennycason.kumo.font.scale.LinearFontScalar
import org.knowm.xchart.*
import java.io.FileNotFoundException

/*

Sources:
1. https://kotlinandroid.org/kotlin/kotlin-read-csv-file/

*/

/*
Description: Used to get user input for file name.
Parameters: none
Return type: File
*/

fun fileScan(): File {
    val scanner = Scanner(System.`in`)

    while (true) { // while  loop to check if file exists
        try { // try-catch to not stop program when file doesn't exist
            print("Input file name (including extension): ")
            val fileName = scanner.nextLine()

            val file = File(fileName)

            if (!file.exists()) {
                throw FileNotFoundException("$fileName was not found. Please try again.")
            }

            return file
        }
        catch (e: FileNotFoundException) {
            println("${e.message}")
        }
    }
}

/*
Description: Used to count total amount of words.
Parameters: allLines is a List of List of Strings
Return type: none
*/
fun wordCount(allLines: MutableList<List<String>>) {
    var wordCount = 0

    for (i in 1 until allLines.size) {
        wordCount += allLines[i][3].split(" ").count() // Counts word separated by a space. The 4th index is the 'text' column.
    }

    println("Word count: $wordCount")
}

/*
Description: Used to count unique words.
Parameters: allLines is a List of List of Strings
Return type: none
*/
fun vocabSize(allLines: MutableList<List<String>>) {
    val uniqueWordsSet = mutableSetOf<String>()

    for (i in 1 until allLines.size) {
        val words = allLines[i][3].split(" ") // Counts word separated by a space. The 4th index is the 'text' column.
        uniqueWordsSet.addAll(words)
    }

    println("Vocabulary size: ${uniqueWordsSet.size}")
}

/*
Description: Used to count the frequency of words.
Parameters: allLines is a List of List of Strings
Return type: none
*/
fun wordFrequency(allLines: MutableList<List<String>>, top20: MutableMap<String, Int>) {
    val wordsList = mutableListOf<String>()

    for (i in 1 until allLines.size) {
        val words = allLines[i][3].split(" ") // Counts word separated by a space. The 4th index is the 'text' column.
        wordsList.addAll(words)
    }

    val heap = PriorityQueue<Map.Entry<String, Int>>(compareByDescending {it.value}) // A Priority Queue that takes in a Map
    val counter = mutableMapOf<String, Int>()

    for (i in wordsList) { // This works like the solution for Top K Frequent Elements in LeetCode
        counter[i] = counter.getOrDefault(i, 0) + 1
    }

    heap.addAll(counter.entries)

    println("\nWord frequency:")
    var num = 1
    while (heap.isNotEmpty()) {
        val entry = heap.poll()
        println("$num. ${entry.key}: ${entry.value} times")

        if (num < 21) {
            top20[entry.key] = entry.value
        }

        num++
    }
}

/*
Description: Used to count the frequency of characters.
Parameters: allLines is a List of List of Strings
Return type: none
*/
fun characterFrequency(allLines: MutableList<List<String>>) {
    val characterList = mutableListOf<Char>()

    for (i in 1 until allLines.size) {
        val characters = allLines[i][3].toCharArray().toList() // Counts word separated by a space. The 4th index is the 'text' column.
        characterList.addAll(characters)
    }

    val heap = PriorityQueue<Map.Entry<Char, Int>>(compareByDescending {it.value}) // A Priority Queue that takes in a Map
    val counter = mutableMapOf<Char, Int>()

    for (i in characterList) { // This works like the solution for Top K Frequent Elements in LeetCode
        counter[i] = counter.getOrDefault(i, 0) + 1
    }

    heap.addAll(counter.entries)

    println("\nCharacter frequency:")
    var num = 1
    heap.poll() // Skips the space character because it's the first one to be popped
    while (heap.isNotEmpty()) {
        val entry = heap.poll()
        println("$num. ${entry.key}: ${entry.value} times")

        num++
    }
}

/*
Description: Used to count if Stop words exist in the text.
Parameters: allLines is a List of List of Strings
Return type: none
*/
fun stopWords(allLines: MutableList<List<String>>) {
    val stopWords = setOf("a", "yes", "no", "hello", "there", "and", "or", "the", "yesterday", "tomorrow")
    val wordsList = mutableListOf<String>()

    for (i in 1 until allLines.size) {
        val words = allLines[i][3].split(" ") // Counts word separated by a space. The 4th index is the 'text' column.
        wordsList.addAll(words)
    }

    val stopWordsMap = mutableMapOf<String, Boolean>()

    for (i in stopWords) {
        stopWordsMap[i] = wordsList.contains(i)
    }

    println("\nStop words:")
    var num = 1
    for ((word, exists) in stopWordsMap) {
        println("$num. $word: ${if (exists) "Yes" else "No"}")
        num++
    }
}

/*
Description: Used to generate a word cloud of the top 20 words.
Parameters: wordFreq is a List of WordFrequency
Return type: none
*/
fun top20wordCloud(wordFreq: List<WordFrequency>) {
    val dimension = Dimension(250, 250)
    val wordCloud = WordCloud(dimension, CollisionMode.PIXEL_PERFECT)
    wordCloud.setPadding(0)
    wordCloud.setBackground(CircleBackground(100))
    wordCloud.setColorPalette(com.kennycason.kumo.palette.ColorPalette(Color(52, 180, 235), Color(247, 45, 122)))
    wordCloud.setFontScalar(LinearFontScalar(10, 40))
    wordCloud.build(wordFreq)
    wordCloud.writeToFile("C:/Users/My PC/Downloads/top20wordCloud.png")
}

/*
Description: Used to generate a bar graph of the frequency of posts per month.
Parameters: allLines is a List of List of Strings
Return type: none
*/
fun monthlyPostsGraph(allLines: MutableList<List<String>>) {
    val datesList = mutableListOf<String>()

    for (i in 1 until allLines.size) {
        val words = allLines[i][2].split("-") // Counts word separated by a space. The 3rd index is the 'date_created' column'.
        datesList.addAll(words)
    }

    val counter = mutableMapOf<String, Int>()

    for (i in 1 until datesList.size step 3) { // This works like the solution for Top K Frequent Elements in LeetCode
        counter[datesList[i]] = counter.getOrDefault(datesList[i], 0) + 1
    }

    val sorted = counter.toSortedMap() // Sorts the 'counter' Map

    val months = sorted.keys.toList() // Gets keys from 'sorted' Map
    val frequency = sorted.values.toList() // Gets values from 'sorted' Map

    val barChart = CategoryChartBuilder().width(800).height(600).title("Monthly Post Frequency").xAxisTitle("Months").yAxisTitle("Total posts").build() // Basically builds the chart

    barChart.addSeries("Post Frequency", months, frequency)

    SwingWrapper(barChart).displayChart() // Displays the chart
}

/*
Description: Used to generate a pie graph of the frequency of symbols found in the 'text' column.
Parameters: allLines is a List of List of Strings
Return type: none
*/
fun symbolsGraph(allLines: MutableList<List<String>>) {
    val symbolsList = mutableListOf<Char>()

    for (i in 1 until allLines.size) {
        val characters = allLines[i][3].toCharArray().toList() // Counts word separated by a space. The 4th index is the 'text' column.
        symbolsList.addAll(characters.filter { !it.isLetter() && !it.isDigit() && !it.isWhitespace()})
    }

    val symbolsMap = mutableMapOf<Char, Int>()

    for (i in symbolsList) {
        symbolsMap[i] = symbolsMap.getOrDefault(i, 0) + 1
    }

    val sorted = symbolsMap.toSortedMap() // Sorts the 'symbolsMap' Map

    val pieChart = PieChartBuilder().width(800).height(600).title("Symbols Frequency").build() // Basically builds the chart

    for (i in sorted) {
        pieChart.addSeries(i.key.toString(), i.value)
    }

    SwingWrapper(pieChart).displayChart() // Displays the chart
}

fun main() {
    val file = fileScan()
    val allLines = mutableListOf<List<String>>()
    val top20 = mutableMapOf<String, Int>()

    file.forEachLine { line ->
        val wholeLine = line.split(",").map { it.lowercase() }
        allLines.add(wholeLine)
    }

    wordCount(allLines)
    vocabSize(allLines)
    wordFrequency(allLines, top20)
    characterFrequency(allLines)

    println("\nTop 20 most frequent words:")
    var counter = 1;

    for ((key, value) in top20) {
        println("$counter. $key: $value")

        counter++
    }

    stopWords(allLines)

    val wordFreq = top20.map { (word, freq) ->
        WordFrequency(word, freq)
    }

    top20wordCloud(wordFreq)
    monthlyPostsGraph(allLines)
    symbolsGraph(allLines)
}