package main
//猜测分隔符是什么，读出文件的前5行来统计
//用法: guess_separate  input.txt
import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	if len(os.Args) == 1 || os.Args[1] == "-h" || os.Args[1] == "--help" {
		fmt.Printf("usage: %s file\n", filepath.Base(os.Args[0]))
		os.Exit(1)
	}

	// hold the separators we are interested in; for whitespace-separated files
	// we will adopt the convention that the separator is "" (the empty string)
	separators := []string{"\t", "*", "|", "•"}

	linesRead, lines := readUpToNLines(os.Args[1], 10)
	counts := createCounts(lines, separators, linesRead)
	separator := guessSep(counts, separators, linesRead)
	report(separator)
}

// Reads only the number of lines specified - or fewer if the file has fewer lines - and
// returns the number of lines it actually read as well as the lines themselves.
func readUpToNLines(filename string, maxLines int) (int, []string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal("failed to open the file: ", err)
	}
	defer func() {
		if err = file.Close(); err != nil {
			log.Println("error closing file: ", err)
		}
	}()

	lines := make([]string, maxLines)
	reader := bufio.NewReader(file)
	i := 0

	for ; i < maxLines; i++ {
		line, err := reader.ReadString('\n')
		if line != "" {
			lines[i] = line
		}
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal("failed to finish reading the file: ", err)
		}
	}

	// return the subslice actually used; could be < maxLines
	return i, lines[:i]
}

// Populate a matrix tha tholds the counts of each separator for each line that was read.
func createCounts(lines, separators []string, linesRead int) [][]int {
	counts := make([][]int, len(separators))

	for sepIndex := range separators {
		counts[sepIndex] = make([]int, linesRead)
		for lineIndex, line := range lines {
			counts[sepIndex][lineIndex] = strings.Count(line, separators[sepIndex])
		}
	}

	return counts
}

// Finds the first []int the the counts slices whose counts are all the same - and nonzero
func guessSep(counts [][]int, separators []string, linesRead int) string {
	for sepIndex := range separators {
		same := true
		count := counts[sepIndex][0]

		for lineIndex := 1; lineIndex < linesRead; lineIndex++ {
			if counts[sepIndex][lineIndex] != count {
				same = false
				break
			}
		}
		if count > 0 && same {
			return separators[sepIndex]
		}
	}

	return ""
}

func report(separator string) {
	switch separator {
	case "":
		fmt.Println("whitespace-separated or not separated at all")
	case "\t":
		fmt.Println("tab-separated")
	default:
		fmt.Printf("%s-separated\n", separator)
	}
}
