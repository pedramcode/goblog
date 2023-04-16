package pipelines

import "strings"

func WordChannelize(data string) <-chan string {
	wordsRep := strings.Replace(data, "!", " ", -1)
	wordsRep = strings.Replace(data, "?", " ", -1)
	wordsRep = strings.Replace(data, ".", " ", -1)
	wordsRep = strings.Replace(data, "(", " ", -1)
	wordsRep = strings.Replace(data, ")", " ", -1)
	wordsRep = strings.Replace(data, ":", " ", -1)
	wordsRep = strings.Replace(data, ";", " ", -1)
	wordsRep = strings.Replace(data, "&", " ", -1)
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
