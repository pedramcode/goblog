package test

import (
	"github.com/pedramcode/goblog/pipelines"
	"testing"
)

func TestWordChannelize(t *testing.T) {
	data := "Hello World,This is a test. Is it working?"
	channel := pipelines.WordChannelize(data)
	count := 0
	for _ = range channel {
		count++
	}
	if count != 9 {
		t.Fatalf("Word channelize failed to tokenize words (Counted words: %d)", count)
	}
}

func TestJunkDetection(t *testing.T) {
	data := "What the freak!"
	wChannel := pipelines.WordChannelize(data)
	jChannel := pipelines.JunkChannelize(wChannel)
	count := 0
	for res := range jChannel {
		if res {
			count++
		}
	}
	if count != 1 {
		t.Fatal("Junk detection failed")
	}
}
