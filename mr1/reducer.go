package mr1

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func Reducer() {
	var currentWord string // 为当前单词
	currentCount := 0      // 当前单词频数
	var word string
	reader := bufio.NewReader(os.Stdin)
	for {
		readLine, err := reader.ReadString('\n')
		if err != nil {
			if err != io.EOF {
				panic(err)
			}
			break
		}
		readLine = strings.TrimRight(readLine, "\r\n")
		readLine = strings.TrimSpace(readLine) //去除字符串首尾的空白字符
		split := strings.Split(readLine, "\t") //按照制表符分隔单词和数量
		if len(split) == 2 {
			word = split[0]
			count, err := strconv.Atoi(split[1]) //将字符串类型的‘1’转换为整型1
			if err != nil {
				continue
			}
			if currentWord == word { //如果当前的单词等于读入的单词
				currentCount += count //单词频数加1
			} else {
				if len(currentWord) > 0 { //如果当前的单词不为空则打印其单词和频数
					fmt.Printf("%s\t%d\n", currentWord, currentCount)
				}
				currentCount = count //否则将读入的单词赋值给当前单词，且更新频数
				currentWord = word
			}
		}
	}
	if currentWord == word {
		fmt.Printf("%s\t%d\n", currentWord, currentCount)
	}
}
