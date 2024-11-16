require 'csv'
require 'date'

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

else
  puts "File not found."
end
