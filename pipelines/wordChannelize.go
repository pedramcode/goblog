package pipelines

import "strings"

func WordChannelize(data string) <-chan string {
	words := strings.Split(data, " ")
	channel := make(chan string)
	go func() {
		for _, word := range words {
			channel <- word
		}
		close(channel)
	}()
	return channel
}
