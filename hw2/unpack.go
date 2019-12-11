package main

import (
	"errors"
	"strconv"
	"strings"
)

var (
	errBadInput        error = errors.New("Incorrect input")
	errIncorrectEscape error = errors.New("Incorrect usage of escape")
)

func runeEscaped(r rune, escaped *bool, b *strings.Builder) bool {
	if *escaped {
		b.WriteRune(r)
		*escaped = false
		return true
	}
	return false
}

// Unpack a given string
func Unpack(s string) (string, error) {
	var res strings.Builder

	escActive := false
	unpacked := false

	runes := []rune(s)

	if _, err := strconv.Atoi(string(runes[0])); err == nil {
		return "", errBadInput
	}

	for i, v := range runes {
		switch v {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			if !runeEscaped(v, &escActive, &res) {
				digit, _ := strconv.Atoi(string(v))
				if unpacked || digit == 0 { // check for two succesive unpackers or unpacker with value zero
					return "", errBadInput
				}
				res.WriteString(strings.Repeat(string(runes[i-1]), digit-1))
				unpacked = true
			}

		case '\\':
			if !runeEscaped(v, &escActive, &res) {
				if i == (len(runes) - 1) { // check for escape is last element
					return "", errIncorrectEscape
				}
				escActive = true
				unpacked = false
			}

		default:
			if escActive { // check for non escapable character
				return "", errIncorrectEscape
			}
			res.WriteRune(v)
			unpacked = false
		}
	}

	return res.String(), nil
}
