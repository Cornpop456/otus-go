package main

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
)

// WordFrequency represents a single word frequency
type WordFrequency struct {
	Word      string
	Frequency uint32
}

func (wf WordFrequency) String() string {
	return fmt.Sprintf("{%s} was found {%d} times", wf.Word, wf.Frequency)
}

func clearText(text string) string {
	text = strings.ToLower(text)
	text = strings.ReplaceAll(text, "\n", " ")
	text = strings.ReplaceAll(text, "\t", " ")

	reg1, _ := regexp.Compile(`\s\s+`)
	reg2, _ := regexp.Compile(`[^\wА-Яа-я\s\\/]+`)

	text = reg1.ReplaceAllString(reg2.ReplaceAllString(text, ""), " ")
	text = strings.TrimSpace(text)

	return text
}

// GetWordsFrequency returns slice of a reverse sorted slice of WordFrequency from a given text with top arg number of elements
func GetWordsFrequency(text string, top uint32) []WordFrequency {
	var (
		dict           map[string]int
		frequencySlice []WordFrequency
		counter        = 0
	)

	text = clearText(text)

	splitedText := strings.Split(text, " ")

	dict = make(map[string]int, len(splitedText))

	for _, v := range splitedText {
		dict[v]++
	}

	frequencySlice = make([]WordFrequency, len(dict))

	for key, value := range dict {
		frequencySlice[counter] = WordFrequency{key, uint32(value)}
		counter++
	}

	sort.Slice(frequencySlice, func(i, j int) bool {
		return frequencySlice[i].Frequency > frequencySlice[j].Frequency
	})

	if top > uint32(len(frequencySlice)) {
		return frequencySlice
	}

	res := make([]WordFrequency, top)

	copy(res, frequencySlice[:top])

	return res
}

// Top10 return top 10 words from given text
func Top10(text string) []string {
	val := GetWordsFrequency(text, 10)

	res := make([]string, len(val))

	for i, v := range val {
		res[i] = v.Word
	}

	return res
}

func main() {
	for _, v := range Top10(Text) {
		fmt.Println(v)

	}
}
