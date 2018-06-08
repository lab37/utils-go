package main
// 词频统计
// 用法 word_frequency  input.txt
import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"sort"
	"strings"
	"unicode"
	"unicode/utf8"
)

func main() {
	if len(os.Args) == 1 || os.Args[1] == "-h" || os.Args[1] == "--help" {
		fmt.Printf("usage: %s <file1> [<file2> [... <fileN]]\n", filepath.Base(os.Args[0]))
		os.Exit(1)
	}

	frequencyForWord := map[string]int{} // same as: make(map[string]int)
	for _, filename := range commandLineFiles(os.Args[1:]) {
		updateFrequencies(filename, frequencyForWord)
	}
	reportByWords(frequencyForWord)
	wordsForFrequency := invertStringIntMap(frequencyForWord)
	reportByFrequency(wordsForFrequency)
}

func commandLineFiles(files []string) []string {
	if runtime.GOOS == "windows" {
		args := make([]string, 0, len(files))
		for _, name := range files {
			if matches, err := filepath.Glob(name); err != nil {
				args = append(args, name) // Invalid pattern
			} else if matches != nil { // At least one match
				args = append(args, matches...)
			}
		}
		return args
	}
	return files
}

func updateFrequencies(filename string, frequencyForWord map[string]int) {
	file, err := os.Open(filename)
	if err != nil {
		log.Println("failed to open the file: ", err)
		return
	}
	defer func() {
		if err = file.Close(); err != nil {
			log.Println("failed to close the file: ", err)
		}
	}()

	readAndUpdateFrequencies(bufio.NewReader(file), frequencyForWord)
}

func readAndUpdateFrequencies(reader *bufio.Reader, frequencyForWord map[string]int) {
	for { //用一个无限循环来一行一行的读取文件
		line, err := reader.ReadString('\n')

		for _, word := range splitOnNonLetters(strings.TrimSpace(line)) {
			if len(word) > utf8.UTFMax || utf8.RuneCountInString(word) > 1 {
				frequencyForWord[strings.ToLower(word)]++
			}
		}

		if err != nil {
			if err != io.EOF {
				log.Println("failed to finish reading the file: ", err)
			}
			break
		}
	}
}

// We only want to count words of 2 or more letters. The cheapest way is to
// see if the word has enough bytes to represent at least one UTF-8
// character that uses the most possible bytes (i.e., 4); but if it is less
// than 4 bytes it could still be a 2, 3, or 4 letter word if they're 7-bit
// ASCII so for this case we actually count the runes (which will be cheap
// because there are at most 4 to count)
func splitOnNonLetters(s string) []string {
	notALetter := func(char rune) bool { return !unicode.IsLetter(char) }
	return strings.FieldsFunc(s, notALetter)
}

func invertStringIntMap(intForString map[string]int) map[int][]string {
	stringsForInt := make(map[int][]string, len(intForString))

	for key, value := range intForString {
		stringsForInt[value] = append(stringsForInt[value], key)
	}

	return stringsForInt
}

func reportByWords(frequencyForWord map[string]int) {
	words := make([]string, 0, len(frequencyForWord))
	wordWidth, frequencyWidth := 0, 0

	for word, frequency := range frequencyForWord {
		words = append(words, word)
		if width := utf8.RuneCountInString(word); width > wordWidth {
			wordWidth = width
		}
		if width := len(fmt.Sprint(frequency)); width > frequencyWidth {
			frequencyWidth = width
		}
	}

	sort.Strings(words)
	gap := wordWidth + frequencyWidth - len("Word") - len("Frequency")
	fmt.Printf("Word %*s%s\n", gap, " ", "Frequency")
	for _, word := range words {
		fmt.Printf("%-*s %*d\n", wordWidth, word, frequencyWidth, frequencyForWord[word])
	}
}

func reportByFrequency(wordsForFrequency map[int][]string) {
	frequencies := make([]int, 0, len(wordsForFrequency))

	for frequency := range wordsForFrequency {
		frequencies = append(frequencies, frequency)
	}

	sort.Ints(frequencies)
	width := len(fmt.Sprint(frequencies[len(frequencies)-1]))
	fmt.Println("Frequency → Words")
	for _, frequency := range frequencies {
		words := wordsForFrequency[frequency]
		sort.Strings(words)
		fmt.Printf("%*d %s\n", width, frequency, strings.Join(words, ", "))
	}
}
