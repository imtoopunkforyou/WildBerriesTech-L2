package main

import (
	"fmt"
	"sort"
	"strings"
)

/*
=== Поиск анаграмм по словарю ===

Напишите функцию поиска всех множеств анаграмм по словарю.
Например:
'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
'листок', 'слиток' и 'столик' - другому.

Входные данные для функции: ссылка на массив - каждый элемент которого - слово на русском языке в кодировке utf8.
Выходные данные: Ссылка на мапу множеств анаграмм.
Ключ - первое встретившееся в словаре слово из множества
Значение - ссылка на массив, каждый элемент которого, слово из множества. Массив должен быть отсортирован по возрастанию.
Множества из одного элемента не должны попасть в результат.
Все слова должны быть приведены к нижнему регистру.
В результате каждое слово должно встречаться только один раз.

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	str := []string{"пЯтак", "пяТка", "тяпкА", "листоК", "сЛиток", "столИк", "Кот", "пятка", "тОк", "оКт", "АбобА"}
	anagramMap := findAnagrams(str)
	fmt.Println(anagramMap)
}

func findAnagrams(words []string) map[string][]string {
	anagramMap := make(map[string][]string) // Мапа для записи анаграмм.

	for _, word := range words {
		lowerWord := strings.ToLower(word) // Приведение слов к нижнему регистру.
		sortedWord := sortWord(lowerWord)  // Соритруем буквы в словах.
		// Ключ - отсортированное слово, значение - не остортированное, таким образом мы группируем по анаграммы.
		anagramMap[sortedWord] = append(anagramMap[sortedWord], lowerWord)
	}

	resultMap := make(map[string][]string) // Нужна для изменения ключа в читаемый вид.

	for _, value := range anagramMap {
		if len(value) > 1 {
			sort.Strings(value)                // Сортируем слова в строке.
			newValue := deleteDuplicate(value) // Удаляем дубликаты из значений старой мапы.
			resultMap[value[0]] = newValue     // Ключ - 0 элемент из анаграммы, как первое встретившееся
		}
	}

	return resultMap
}

// Для сортировки букв в слове
func sortWord(word string) string {
	runes := []rune(word) // Преобразуем строку в слайс рун.

	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	}) // Сортировка, получаем одинаковую последовательность букв из слов.

	return string(runes)
}

// Для удаления повторяющихся слов.
func deleteDuplicate(words []string) []string {
	var result []string

	for i := 0; i < len(words); i++ {
		if i > 0 { // Условие нужно, что бы 0 элемент был в слайсе.
			if words[i] != words[i-1] { // Проверяем соседние элементы, одинаковые или нет, так как они отсортированы.
				result = append(result, words[i])
			}
		} else {
			result = append(result, words[i]) // Добавляем 0 элемент.
		}
	}

	return result
}
