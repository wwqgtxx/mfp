package find_pair

import (
	"bufio"
	"fmt"
	"github.com/go-ego/gse"
	"io"
	"os"
	"strings"
)

func buildMapper(filePath string) map[string]string {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	reader := bufio.NewReader(file)

	mapper := make(map[string]string)
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
		items := strings.Fields(readLine)
		if len(items) < 2 {
			continue
		}
		mapper[items[0]] = strings.Join(items[1:], " ")
	}
	return mapper
}

func FindPair() {
	var seg gse.Segmenter
	seg.SkipLog = true
	err := seg.LoadDict("dictionary.txt")
	if err != nil {
		panic(err)
	}
	filePath := "./res.txt"
	if len(os.Args) > 2 {
		filePath = os.Args[2]
	}
	mapper := buildMapper(filePath)
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("请输入想要查找关联模式的关键词：(输入“q”退出")
		readLine, err := reader.ReadString('\n')
		if err != nil {
			if err != io.EOF {
				panic(err)
			}
			break
		}
		readLine = strings.TrimRight(readLine, "\r\n")
		readLine = strings.TrimSpace(readLine)
		word := readLine
		if strings.ToLower(word) == "q" {
			return
		}
		if len(word) < 1 {
			continue
		}

		if result, ok := mapper[word]; ok {
			fmt.Println(result)
		} else {
			fmt.Println("没有找到该关键词，提供一些近似的关键词：")
			keywordsList := seg.Slice(word, false)
			fmt.Println(keywordsList)
			flags := false
			for _, kw := range keywordsList {
				if result, ok := mapper[kw]; ok {
					fmt.Printf("%s:%s\n", kw, result)
					flags = true
				}
			}
			if !flags {
				keywordsList := seg.CutSearch(word)
				fmt.Println(keywordsList)
				for _, kw := range keywordsList {
					if result, ok := mapper[kw]; ok {
						fmt.Printf("%s:%s\n", kw, result)
						flags = true
					}
				}
			}
			if !flags {
				fmt.Println("无结果")
			}
		}
	}
}
