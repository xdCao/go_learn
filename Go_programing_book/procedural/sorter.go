package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	sort "xdCao/golearn/goprogramming/sort"
)

var inputFile *string = flag.String("i", "", "需要排序的文件")
var outputFile *string = flag.String("o", "", "结果文件")
var algo *string = flag.String("a", "", "排序算法")

func main() {
	flag.Parse()
	if inputFile != nil {
		fmt.Println("inputFile =", *inputFile, "outputFile =", *outputFile, "algorithm =", *algo)
	}

	values, err := readFile(*inputFile)
	if err == nil {
		fmt.Println("Read values:", values)
	} else {
		fmt.Println(err)
	}

	switch *algo {
	case "qsort":
		sort.Qsort(values)
	case "bsort":
		sort.BubbleSort(values)
	}

	err1 := writeFile(values, *outputFile)
	if err != nil {
		fmt.Println(err1)
		return
	}

}

func writeFile(values []int, outFile string) error {
	file, err := os.Create(outFile)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer file.Close()
	for _, v := range values {
		file.WriteString(strconv.Itoa(v) + "\n")
	}
	return nil
}

// 从文件读取数据到切片
func readFile(input string) (values []int, err error) {
	file, err := os.Open(input)
	if err != nil {
		fmt.Println(err)
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
			fmt.Println("line toot long")
			return
		}

		str := string(line) // 将字符数组转换成字符串
		value, err1 := strconv.Atoi(str)
		if err1 != nil {
			err = err1
			return
		}
		values = append(values, value)
	}
	return
}
