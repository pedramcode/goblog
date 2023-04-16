package pipelines

import "strings"

func WordChannelize(data string) <-chan string {
	replacer := strings.NewReplacer("!", " ", "?", " ", ".", " ", "(",
		" ", ")", " ", ":", " ", ";", " ", "&", " ")
	wordsRep := replacer.Replace(data)
	words := strings.Split(wordsRep, " ")
	channel := make(chan string)
	go func() {
		for _, word := range words {
			channel <- word
		}
		close(channel)
	}()
	return channel
}
