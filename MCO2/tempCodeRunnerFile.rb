
stop_words = {"a"=>0, "an"=>0, "the"=>0, "and"=>0, "but"=>0, "or"=>0, "in"=>0, "on"=>0, "at"=>0, "with"=>0}

def find_stops(word)
  if stop_words.key?(word)
    stop_words[word] += 1
  end
end