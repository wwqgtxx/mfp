package find_pair

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/go-ego/gse"
	"io"
	"os"
	"strings"
)

var (
	seg      gse.Segmenter
	mapper   map[string]string
	FilePath = "./res.txt"
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

func Init() {
	seg.SkipLog = true
	err := seg.LoadDict("dictionary.txt")
	if err != nil {
		panic(err)
	}
	mapper = buildMapper(FilePath)
}

func FindPairFromStringToWriter(word string, out io.Writer) (err error) {
	if len(word) < 1 {
		return nil
	}

	if result, ok := mapper[word]; ok {
		_, err = fmt.Fprintln(out, result)
		if err != nil {
			return err
		}
	} else {
		_, err = fmt.Fprintln(out, "没有找到该关键词，提供一些近似的关键词：")
		if err != nil {
			return err
		}
		keywordsList := seg.Cut(word)
		_, err = fmt.Fprintln(out, keywordsList)
		if err != nil {
			return err
		}
		flags := false
		for _, kw := range keywordsList {
			if result, ok := mapper[kw]; ok {
				_, err = fmt.Fprintf(out, "%s:%s\n", kw, result)
				if err != nil {
					return err
				}
				flags = true
			}
		}
		if !flags {
			keywordsList := seg.CutSearch(word)
			_, err = fmt.Fprintln(out, keywordsList)
			for _, kw := range keywordsList {
				if result, ok := mapper[kw]; ok {
					_, err = fmt.Fprintf(out, "%s:%s\n", kw, result)
					if err != nil {
						return err
					}
					flags = true
				}
			}
		}
		if !flags {
			_, err = fmt.Fprintln(out, "无结果")
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func FindPairFromString(word string) (str string, err error) {
	bufferString := bytes.NewBufferString("")
	err = FindPairFromStringToWriter(word, bufferString)
	str = bufferString.String()
	return
}

func FindPair() {
	if len(os.Args) > 2 {
		FilePath = os.Args[2]
	}
	Init()
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
		err = FindPairFromStringToWriter(word, os.Stdout)
		if err != nil {
			panic(err)
		}
	}
}
