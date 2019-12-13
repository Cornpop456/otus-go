package main

import (
	"errors"
	"fmt"
	"log"
	"regexp"
	"sort"
	"strings"
)

type wordFrequency struct {
	word      string
	frequency uint32
}

func (wf wordFrequency) String() string {
	return fmt.Sprintf("{%s} was found {%d} times", wf.word, wf.frequency)
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

func getWordsFrequency(text string, top uint32) ([]wordFrequency, error) {
	var (
		dict           map[string]int
		frequencySlice []wordFrequency
		counter        = 0
	)

	text = clearText(text)

	s := strings.Split(text, " ")

	dict = make(map[string]int, len(s))

	for _, v := range s {
		dict[v]++
	}

	frequencySlice = make([]wordFrequency, len(dict))

	for key, value := range dict {
		frequencySlice[counter] = wordFrequency{key, uint32(value)}
		counter++
	}

	sort.Slice(frequencySlice, func(i, j int) bool {
		return frequencySlice[i].frequency > frequencySlice[j].frequency
	})

	if top > uint32(len(frequencySlice)) {
		return []wordFrequency{}, errors.New("Given [top] argument is bigger than number of words")
	}

	return frequencySlice[:top], nil
}

func main() {
	res, err := getWordsFrequency(Text, 10)

	if err != nil {
		log.Fatal(err)
	}

	for _, v := range res {
		fmt.Println(v)

	}
}
