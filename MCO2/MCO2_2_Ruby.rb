require 'csv'
require 'date'
# to use the QuickChart API
require 'json'
require 'http'

# Input file
print "Enter the filename: "
filename = gets.chomp

# to do
# while passing, count total words
  # total num of unique words
  # word frequency
    # will later find 20 most freq words
  # frequency for all chars then sort
    # set to lowercase to count accurately
  # identify 10 stop words

  # make a separate hash for non-alphanumeric symbols
  # make a hash for each month

# visualize data 

# Initialize hashes and counters
word_count = Hash.new(0)
char_count = Hash.new(0)
special = Hash.new(0)
monthly_data = Hash.new(0)
total_words = 0
unique_words = 0

# Read CSV file if it exists
  if File.exist?(filename)
    CSV.foreach(filename, headers: true) do |row|
      # Tweet text
      if row['text']
        text = row['text']
        words = text.downcase.split
        words.each do |word|
          word_count[word] += 1
          total_words += 1
    
          word.each_char do |char|
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
          date = DateTime.strptime(row['date_created'], '%Y-%m-%d %H:%M:%S')
          # for getting month in word form
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
  word_count.sort_by { |word, count| -count }.first(20).each do |word, count|
    puts "#{word}: #{count}"
  end

  puts "\nTop 10 most frequent special characters:"
  special.sort_by { |symbol, count| -count }.first(10).each do |char, count|
    puts "#{char}: #{count}"
  end

  puts "\nMonth Tweets:"
  monthly_data.sort_by { |month, count| -count }.each do |month, count|
    puts "#{month}: #{count}"
  end
  
  puts "\nAll Chars (sorted by count):"
  char_count.sort_by { |char, count| -count }.each do |char, count|
    puts "#{char}: #{count}"
  end

  puts "\nALL SPECIAL CHARS"
  special.sort_by { |symbol, count| -count }.each do |symbol, count|
    puts "#{symbol}: #{count}" 
  end

  puts "\nALL WORDS"
  word_count.sort_by { |word, count| -count }.each do |word, count|
    puts "#{word}: #{count}" 
  end

wordcloud_setup = {
  format: "png",
  width: 1000,
  height: 1000,
  fontFamily: "sans-serif",
  scale: "linear",
  text: word_count.map { |word, count| ([word] * count).join(' ') }.join(' ')
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
  format('#%06x', rand(0..0xFFFFFF))
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
