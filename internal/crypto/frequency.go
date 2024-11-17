package crypto

import (
	"fmt"
	"sort"
)

const (
	frequentCharacterEnglish = 'e'
	maxTries                 = 3
)

type Frequency struct {
	charactersCount map[rune]int
}

func NewFrequency() Frequency {
	return Frequency{
		charactersCount: make(map[rune]int),
	}
}

func (f Frequency) Decrypt(s *Stream) error {
	err := f.countCharacters(s)
	if err != nil {
		return err
	}
	f.printCharactersCount(s)
	tries := 3
	mostFrequent := f.getMostFrequentCharacters(tries)
	fmt.Fprintln(s, mostFrequent)
	for char := range mostFrequent {
		err := decryptCesar(s, char)
		if err != nil {
			return err
		}
	}
	return nil
}

func findKey(encryptedChar rune) int {
	return int(encryptedChar) - frequentCharacterEnglish
}

func (f Frequency) getMostFrequentCharacters(quantity int) map[rune]int {
	characters := make([]rune, 0, len(f.charactersCount))
	for char := range f.charactersCount {
		characters = append(characters, char)
	}
	sort.Slice(characters, func(i int, j int) bool {
		return f.charactersCount[characters[i]] < f.charactersCount[characters[j]]
	})
	mostFrequent := make(map[rune]int, quantity)
	for i := 0; i < quantity; i++ {
		char := characters[len(characters)-1-i]
		mostFrequent[char] = f.charactersCount[char]
	}
	return mostFrequent
}

func (f Frequency) countCharacters(s *Stream) error {
	count := func(char byte) error {
		f.charactersCount[rune(char)]++
		return nil
	}
	return s.Scan(count)
}

func (f Frequency) printCharactersCount(s *Stream) {
	fmt.Fprintln(s, "Encrypted characters frequency:")
	for char, number := range f.charactersCount {
		fmt.Fprintf(s, "%v %v\n", string(char), number)
	}
}

func decryptCesar(s *Stream, char rune) error {
	key := findKey(char)
	fmt.Fprintf(s, "Cesar decrypted with key %v assuming character %v represents %v\n", key, string(char), string(frequentCharacterEnglish))
	cesar := NewCesar(key)
	return cesar.Decrypt(s)
}
