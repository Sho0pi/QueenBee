package main

import (
	"bytes"
	"fmt"
)

var (
	VowelsToNumbers = map[byte]byte{
		'a': '4',
		'A': '4',
		'e': '3',
		'E': '3',
		'i': '1',
		'I': '1',
		'o': '0',
		'O': '0',
	}
)

type Rule struct {
	Rules []func(word []byte) [][]byte
}

func (r *Rule) AddRule(f func(word []byte) [][]byte) {
	r.Rules = append(r.Rules, f)
}

func (r *Rule) Crack(word []byte) {
	results := [][]byte{}
	for _, f := range r.Rules {
		results = append(results, f(word)...)
	}
	for _, pass := range results {
		fmt.Println(string(pass))
	}
}

func emptyRule(word []byte) [][]byte {
	return [][]byte{word}
}

func vowelsToNumbersRule(word []byte) [][]byte {
	storage := [][]byte{}
	vowelsToNumbersCrack(&storage, &word, 0)
	return storage

}

func camelMixRule(word []byte) [][]byte {
	storage := [][]byte{}
	newWord := bytes.ToLower(word)
	camelMixCrack(&storage, &newWord, 0)
	return storage

}

func vowelsToNumbersCrack(storage *[][]byte, word *[]byte, index int) {
	for index < len(*word) {
		if val, ok := VowelsToNumbers[(*word)[index]]; ok {
			newWord := append((*word)[:0:0], (*word)...)
			newWord[index] = val

			vowelsToNumbersCrack(storage, &newWord, index+1)
		}
		index++
	}
	(*storage) = append((*storage), (*word))
}

func camelMixCrack(storage *[][]byte, word *[]byte, index int) {
	for index < len(*word) {
		if !(((*word)[index] < 'a' || (*word)[index] > 'z') && ((*word)[index] < 'A' || (*word)[index] > 'Z')) {
			tmp := []byte{(*word)[index]}
			tmp = bytes.ToUpper(tmp)

			tmpWord := append((*word)[:0:0], (*word)...)

			newWord := append(tmpWord[:index], tmp[0])
			newWord = append(newWord, (*word)[index+1:]...)
			camelMixCrack(storage, &newWord, index+1)
		}
		index++
	}
	(*storage) = append((*storage), *word)
}

func momboComboRule(word []byte) [][]byte {

	storage := [][]byte{}
	for _, newWord := range vowelsToNumbersRule(word) {
		storage = append(storage, camelMixRule(newWord)...)
	}
	return storage
}
