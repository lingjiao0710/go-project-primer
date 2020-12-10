package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

type problem struct {
	question string //问题
	answer   string //答案
}

func main() {
	csvFilename := flag.String("csv", "problems.csv", "包含问题和结果的csv文件")
	timeLimit := flag.Int("limit", 5, "答题超时时间，默认30秒")

	flag.Parse()

	file, err := os.Open(*csvFilename)
	if err != nil {
		exit(fmt.Sprintf("打开csv文件%s失败\n", *csvFilename))
	}

	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit("解析csv文件失败\n")
	}

	fmt.Println(lines)

	//解析文件内容到切片中
	problems := parseLines(lines)
	fmt.Println(problems)

	//创建计时器
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	//打分
	correct := 0

loop:
	for i, p := range problems {
		fmt.Printf("Problem #%d,:%s = \n", i+1, p.question)

		//创建一个通道用于接收答案
		answerCh := make(chan string)

		//用协程跑答题输入
		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerCh <- answer
		}()

		select {
		//时间到
		case <-timer.C:
			fmt.Println("答题超时")
			break loop
		case answer := <-answerCh:
			if answer == p.answer {
				correct++
			}
		}
	}

	fmt.Printf("你做对了 %d 题， 总共 %d 题\n", correct, len(problems))

}

func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i] = problem{
			question: line[0],
			answer:   strings.TrimSpace(line[1]),
		}
	}
	return ret
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
