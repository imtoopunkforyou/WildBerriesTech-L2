package main

import (
	"reflect"
	"testing"
)

func TestSortWord(t *testing.T) {
	input := []string{"пятка", "листок", "кот"}
	expect := []string{"акптя", "иклост", "кот"}

	for i, v := range input {
		input[i] = sortWord(v)
	}
	if !reflect.DeepEqual(input, expect) {
		t.Errorf("expect %v, got %v", expect, input)
	}
}

func TestDeleteDuplicates(t *testing.T) {
	input := []string{"пятка", "пятка", "листок", "листок", "листок", "кот", "кот"}
	expect := []string{"пятка", "листок", "кот"}
	newStr := deleteDuplicate(input)

	if !reflect.DeepEqual(newStr, expect) {
		t.Errorf("expect %v, got %v", expect, newStr)
	}
}

func TestFindAnagram(t *testing.T) {
	input := []string{"пЯтак", "пяТка", "тяпкА", "листоК", "сЛиток", "столИк", "Кот", "тОк", "оКт"}
	expect := map[string][]string{
		"кот":    {"кот", "окт", "ток"},
		"листок": {"листок", "слиток", "столик"},
		"пятак":  {"пятак", "пятка", "тяпка"}}
	anagramMap := findAnagrams(input)

	if !reflect.DeepEqual(anagramMap, expect) {
		t.Errorf("expect %v, got %v", expect, anagramMap)
	}
}

func TestFindAnagramWithDuplicatesAndOneWord(t *testing.T) {
	input := []string{"пЯтак", "пяТка", "тяпкА", "листоК", "сЛиток", "столИк", "Кот", "пятка", "тОк", "оКт", "АбобА"}
	expect := map[string][]string{
		"кот":    {"кот", "окт", "ток"},
		"листок": {"листок", "слиток", "столик"},
		"пятак":  {"пятак", "пятка", "тяпка"}}
	anagramMap := findAnagrams(input)

	if !reflect.DeepEqual(anagramMap, expect) {
		t.Errorf("expect %v, got %v", expect, anagramMap)
	}
}
