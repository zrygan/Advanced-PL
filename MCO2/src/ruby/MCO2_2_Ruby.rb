=begin
  Last names: Miranda
  Language: Ruby
  Paradigm(s):Multi-paradigm: Object-Oriented Programming, Functional, and Generic
=end

require 'csv'
require 'date'

# to use the QuickChart API
require 'json'
require 'http'

# Input file
print "Enter the filename: "
filename = gets.chomp

# Initialize hashes and counters
word_count = Hash.new(0)
char_count = Hash.new(0)
special = Hash.new(0)
monthly_data = Hash.new(0)
total_words = 0
unique_words = 0

# For the stop words
# note how $ makes a variable a global variable
# within ruby, only local scopes are accessible within functions
$stop_words = {"a"=>0, "an"=>0, "the"=>0, "and"=>0, "but"=>0, "or"=>0, "in"=>0, "on"=>0, "at"=>0, "with"=>0}
$stops_total = 0


# To update stop word counter
def find_stops(word)
  if $stop_words.key?(word)
    $stop_words[word] += 1
    $stops_total += 1
  end
end

# Read CSV file if it exists
  if File.exist?(filename)
    # Uses the csv folder to read each row
    # headers: true assigns it to a hash
    CSV.foreach(filename, headers: true) do |row|
      # Note how ifs are used, this is to only read the relevant tables

      if row['text']
        text = row['text']
        # Text is set to downcase to read all letters the same regardless of letter case
        # .split will split it to each word
        words = text.downcase.split
        words.each do |word|
          word_count[word] += 1
          total_words += 1

          find_stops(word);
    
          word.each_char do |char|
            # That regex is used to check if it is alphanumeric
            if char =~ /\A\p{Alnum}+\z/
              char_count[char] += 1
            else
              special[char] += 1
            end
          end
        end
      end
  
      # For month counting
      if row['date_created']
        begin
          # uses datetime library
          # makes the date a DateTime object
          date = DateTime.strptime(row['date_created'], '%Y-%m-%d %H:%M:%S')
          # to get the month in the DateTime
          month = date.strftime('%B')
          monthly_data[month] += 1
        end
      end
    end

  # Calculate the number of unique words using keys
  unique_words = word_count.keys.size

  # Display analysis results
  puts "\n--- Results ---"
  puts "Total words: #{total_words}"
  puts "Unique words: #{unique_words}"

  puts "\nTop 20 most frequent words:"

  #.first will be used to easily limit
  word_count.sort_by { |word, count| -count }.first(20).each do |word, count|
    puts "#{word}: #{count}"
  end

  puts "\nAll special characters:"
  special.sort_by { |symbol, count| -count }.each do |char, count|
    puts "#{char}: #{count}"
  end

  puts "\nTotal Stop words: #{$stops_total}"
  $stop_words.each do |word, count|
    puts "#{word}: #{count}"
  end

  puts "\nMonth Tweets:"
  monthly_data.sort_by { |month, count| -count }.each do |month, count|
    puts "#{month}: #{count}"
  end
  
  puts "\nAll Chars (sorted by count):"
  char_count.sort_by { |char, count| -count }.each do |char, count|
    puts "#{char}: #{count}"
  end

  # puts "\nALL SPECIAL CHARS"
  # special.sort_by { |symbol, count| -count }.each do |symbol, count|
  #   puts "#{symbol}: #{count}" 
  # end

  # puts "\nALL WORDS"
  # word_count.sort_by { |word, count| -count }.each do |word, count|
  #   puts "#{word}: #{count}" 
  # end


# GRAPH MAKINGS
# Uses the quickchart.io to create the images
# It allows for easy creation of images without needing to do any downloads or installs

# Ruby's chart creation is limited especially with word clouds
# Other sources such as magic_cloud need the rmagick and imagemagick 
# which will also need installs within the computer and have limited documentation
# to limit potential machine-specific errors, online apis were used instead

# by using Ruby's json library, requests can be easily sent

# The word cloud api groups up character combinations so when given
# user####, it will group up all "users"
# it doesn't show special symbols such as hashtags

wordcloud_setup = {
  format: "png",
  width: 1000,
  height: 1000,
  fontFamily: "sans-serif",
  scale: "linear",
  text: word_count.sort_by { |word, count| -count }.first(20).map { |word, count| ([word] * count).join(' ') }.join(' ')
}

# Convert the setup to json
cloud_json = wordcloud_setup.to_json

# Make the POST request to QuickChart
word_cloud_url = 'https://quickchart.io/wordcloud'
response1 = HTTP.post(word_cloud_url, headers: { 'Content-Type' => 'application/json' }, body: cloud_json)

# Check if the request was successful and save the image
if response1.status.success?
  File.open('ruby-wordcloud.png', 'wb') do |file|
    file.write(response1.body)
    puts "Word Cloud Success"
  end
  else
    puts "Error: #{response1.body.to_s}"
  end
end

# Bar Chart Setups
month_labels = monthly_data.keys
month_data = monthly_data.values

bar_design = {
  type: 'bar',
  data: {
    labels: month_labels,
    datasets: [{
      label: 'Monthly Tweets',
      data: month_data,
      backgroundColor: 'blue',
      borderColor: 'black',
      borderWidth: 1
    }]
  },
  options: {
    plugins: {
      legend: {
        display: true,
        position: 'top'
      }
    },
    scales: {
      y: {
        beginAtZero: true
      }
    }
  }
}

bar_setup = {
  chart: bar_design.to_json,
  width: 800,
  height: 600,
  backgroundColor: 'white',
  format: 'png',
  version: '4'
}

bar_json = bar_setup.to_json

chart_url = 'https://quickchart.io/chart'
response2 = HTTP.post(chart_url, headers: { 'Content-Type' => 'application/json' }, body: bar_json)

if response2.status.success?
  File.open('ruby-bar.png', 'wb') do |file|
    file.write(response2.body)
    puts "Bar Chart Success"
  end
else
  puts "Error: #{response2.body.to_s}"
end

def random_color
  return format('#%06x', rand(0..0xFFFFFF))
end

# Pie Chart Setups
pie_labels = special.keys
pie_data = special.values

background_colors = pie_labels.map { random_color }

pie_design = {
  type: 'pie',
  data: {
    labels: pie_labels,
    datasets: [{
      data: pie_data,
      backgroundColor: background_colors,
      borderColor: 'white',
      borderWidth: 2
    }]
  },
  options: {
    plugins: {
      legend: {
        display: true,
        position: 'right'
      }
    }
  }
}

pie_setup = {
  chart: pie_design.to_json,
  width: 800,
  height: 600,
  devicePixelRatio: 2,
  backgroundColor: 'white',
  format: 'png',
  version: '4'
}

pie_json = pie_setup.to_json

chart_url = 'https://quickchart.io/chart'
response3 = HTTP.post(chart_url, headers: { 'Content-Type' => 'application/json' }, body: pie_json)

if response3.status.success?
  File.open('ruby-pie.png', 'wb') do |file|
    file.write(response3.body)
    puts "Pie Chart Success"
  end
else
  puts "Error: #{response3.body.to_s}"
end