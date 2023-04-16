package pipelines

// Terrible and the most offensive words to detect :)
var junkDataset []string = []string{"fudge", "freak", "fart"}

func isJunk(word string) bool {
	for _, junk := range junkDataset {
		if word == junk {
			return true
		}
	}
	return false
}

func JunkChannelize(wordChannel <-chan string) <-chan bool {
	channel := make(chan bool)
	go func() {
		for word := range wordChannel {
			channel <- isJunk(word)
		}
		close(channel)
	}()
	return channel
}
