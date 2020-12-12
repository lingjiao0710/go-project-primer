package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/lingjiao0710/test/goprogramming/sorter/algorithm/bubblesort"
	"github.com/lingjiao0710/test/goprogramming/sorter/algorithm/qsort"
	"io"
	"os"
	"strconv"
	"time"
)

var infile *string = flag.String("i", "unsorted.dat", "需要排序的文件")
var outfile *string = flag.String("o", "sorted.dat", "保存排序结果文件")
var algorithm *string = flag.String("a", "qsort", "排序算法")

func readValues(infile string) (values []int, err error) {
	file, err := os.Open(infile)
	if err != nil {
		fmt.Printf("打开文件 %s 失败\n", infile)
		return
	}

	defer file.Close()

	br := bufio.NewReader(file)
	values = make([]int, 0)

	for {
		line, isPrefix, err1 := br.ReadLine()
		if err1 != nil {
			//文件读取发生错误
			if err1 != io.EOF {
				err = err1
				break
			}
		}

		if isPrefix || err1 == io.EOF {
			//fmt.Println("文件内容过长或读取到文件尾")
			return
		}
		//转换字符数组到字符串
		str := string(line)
		//fmt.Println("读取： ", str, "isprefix: ", isPrefix, err1)

		value, err1 := strconv.Atoi(str)
		if err1 != nil {
			err = err1
			return
		}
		fmt.Println(value)
		values = append(values, value)
	}

	return
}

func writeValues(values []int, outfile string) (err error) {
	file, err := os.Create(outfile)
	if err != nil {
		fmt.Printf("打开文件 %s 失败\n", outfile)
		return
	}

	defer file.Close()

	for _, value := range values {
		str := strconv.Itoa(value)
		file.WriteString(str + "\n")
	}

	return nil
}

func main() {
	flag.Parse()

	if infile != nil {
		fmt.Println("infile =", *infile, "outfile=", *outfile, "algorithm=", *algorithm)
	}

	values, err := readValues(*infile)
	if err == nil {
		tStart := time.Now()
		switch *algorithm {
		case "qsort":
			qsort.QuickSort(values)
		case "bubblesort":
			bubblesort.BubbleSort(values)
		default:
			fmt.Println("Sorting algorithm", *algorithm, "is either unknown or unsupported.")
		}
		tEnd := time.Now()
		fmt.Println("The soring process costs", tEnd.Sub(tStart), "to complete.")

		writeValues(values, *outfile)
	} else {
		fmt.Println(err)
	}

}
