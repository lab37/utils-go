package main

import "bufio"
import "flag"
import "fmt"
import "io"
import "os"
import "strconv"
import "time"

/* 使用2种不同的算法排序文件中的数字
   2种算法分别为:冒泡排序和快速排序
   用法:
   sorter -i input文件  -o  output文件 -a 算法[bubblesort|qsort]
*/
var infile *string = flag.String("i", "unsorted.dat", "File contains values for sorting")
var outfile *string = flag.String("o", "sorted.dat", "File to receive sorted values")
var algorithm *string = flag.String("a", "qsort", "Sort algorithm")

//从文件中读取值
func readValues(infile string) (values []int, err error) {
	file, err := os.Open(infile)
	if err != nil {
		fmt.Println("Failed to open the input file ", infile)
		return
	}

	defer file.Close()
	
	br := bufio.NewReader(file)
	
	values = make([]int, 0)
	
	for {
		line, isPrefix, err1 := br.ReadLine()
	
		if err1 != nil {
			if err1 != io.EOF {
				err = err1
			}
			break
		}
	
		if isPrefix {
			fmt.Println("A too long line, seems unexpected.")
			return
		}
	
		str := string(line) // Convert []byte to string
	
		value, err1 := strconv.Atoi(str)
	
		if err1 != nil {
			err = err1
			return
		}
	
		values = append(values, value)
	}
	return
}

//写入输出文件
func writeValues(values []int, outfile string) error {
	file, err := os.Create(outfile)
	if err != nil {
		fmt.Println("Failed to create the output file ", outfile)
		return err
	}
	
	defer file.Close()
	
	for _, value := range values {
		str := strconv.Itoa(value)
		file.WriteString(str + "\n")
	}
	return nil
}

//胃泡排序算法
func bubbleSort(values []int) {
	flag := true

	for i := 0; i < len(values)-1; i++ {
		flag = true
	
		for j := 0; j < len(values)-i-1; j++ {
			if values[j] > values[j+1] {
				values[j], values[j+1] = values[j+1],
					values[j]
				flag = false
			} // end if
		} // end for j = ...
	
		if flag == true {
			break
		}
	
	} // end for i = ...
}
// 快速排序算法
func quickSort(values []int, left, right int) {
	temp := values[left]
	p := left
	i, j := left, right

	for i <= j {
		for j >= p && values[j] >= temp {
			j--
		}
		if j >= p {
			values[p] = values[j]
			p = j
		}
		if values[i] <= temp && i <= p {
			i++
		}
		if i <= p {
			values[p] = values[i]
			p = i
		}
	}
	values[p] = temp
	if p-left > 1 {
		quickSort(values, left, p-1)
	}
	if right-p > 1 {
		quickSort(values, p+1, right)
	}
}


func main() {
	flag.Parse()
	
	if infile != nil {
		fmt.Println("infile =", *infile, "outfile =", *outfile, "algorithm =", *algorithm)
	}
	
	values, err := readValues(*infile)
	if err == nil {
		t1 := time.Now()
		switch *algorithm {
		case "qsort":
			quickSort(values,0, len(values)-1)
		case "bubblesort":
			bubbleSort(values)
		default:
			fmt.Println("Sorting algorithm", *algorithm, "is either unknown or unsupported.")
		}
		t2 := time.Now()
	
		fmt.Println("The sorting process costs", t2.Sub(t1), "to complete.")
	
		writeValues(values, *outfile)
	} else {
		fmt.Println(err)
	}
}
