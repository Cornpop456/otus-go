package main

import (
	"errors"
	"fmt"
	"log"
	"regexp"
	"sort"
	"strconv"
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
	text = strings.TrimSpace(text)
	text = strings.ReplaceAll(text, "\n", " ")
	text = strings.ReplaceAll(text, "\t", " ")

	reg1, _ := regexp.Compile(`\s\s+`)
	reg2, _ := regexp.Compile(`[^\wА-Яа-я\s\\/]+`)

	text = reg2.ReplaceAllString(reg1.ReplaceAllString(text, " "), "")

	return text
}

// GetWordsFrequency returns slice of a reverse sorted slice of WordFrequency from a given text with top arg number of elements
func GetWordsFrequency(text string, top uint32) ([]WordFrequency, error) {
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
		return []WordFrequency{}, errors.New("Given [top] argument is bigger than words number: " + strconv.Itoa(len(frequencySlice)))
	}

	res := make([]WordFrequency, top)

	copy(res, frequencySlice[:top])

	frequencySlice = nil

	return res, nil
}

func main() {
	res, err := GetWordsFrequency(Text, 10)

	if err != nil {
		log.Fatal(err)
	}

	for _, v := range res {
		fmt.Println(v)

	}
}
