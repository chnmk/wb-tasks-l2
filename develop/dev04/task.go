package main

import (
	"slices"
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
*/

func Search(input *[]string) *map[string][]string {
	result := make(map[string][]string)

	dublicates := make(map[string]struct{})

	for _, word := range *input {
		// В результате каждое слово должно встречаться только один раз.
		_, ok := dublicates[word]
		if ok {
			continue
		}

		dublicates[word] = struct{}{}

		createNew := true
		wordLower := strings.ToLower(word)

		for key := range result {
			if sortLetters(key) == sortLetters(wordLower) {
				result[key] = append(result[key], wordLower)
				createNew = false
			}
		}

		if createNew {
			result[wordLower] = []string{wordLower}
		}
	}

	for k, v := range result {
		// Множества из одного элемента не должны попасть в результат
		if len(v) < 2 {
			delete(result, k)
		}

		// Массив должен быть отсортирован по возрастанию
		slices.Sort(v)
	}

	return &result
}

func sortLetters(word string) string {
	l := strings.Split(word, "")

	sort.Strings(l)

	return strings.Join(l, "")
}
