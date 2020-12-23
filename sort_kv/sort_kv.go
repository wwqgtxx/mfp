package sort_kv

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

const (
	I = 100
)

func SortKV() {
	if len(os.Args) < 4 {
		fmt.Println("expected file path and target name")
		os.Exit(1)
	}
	file, err := os.Open(os.Args[2])
	if err != nil {
		panic(err)
	}
	defer file.Close()
	reader := bufio.NewReader(file)

	type kv struct {
		Key   string
		Value int
	}
	lineList := make([]kv, 0)

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
			pattern := split[0]
			supp, err := strconv.Atoi(split[1])
			if err != nil {
				continue
			}
			lineList = append(lineList, kv{pattern, supp})
		}
	}
	itemsCount := len(lineList)
	fmt.Printf("项目总数：%d\n", itemsCount)

	sort.Slice(lineList, func(i, j int) bool {
		return lineList[i].Value > lineList[j].Value
	})
	gListFile, err := os.Create(os.Args[3])
	if err != nil {
		panic(err)
	}
	defer gListFile.Close()

	writer := bufio.NewWriter(gListFile)
	_, err = writer.WriteString("{")
	if err != nil {
		panic(err)
	}
	for i, kv := range lineList {
		if i != 0 {
			_, err = writer.WriteString(",\n  ")
			if err != nil {
				panic(err)
			}
		}
		b, err := json.Marshal(kv.Key)
		if err != nil {
			panic(err)
		}
		_, err = writer.Write(b)
		if err != nil {
			panic(err)
		}
		_, err = writer.WriteString(":")
		if err != nil {
			panic(err)
		}
		b, err = json.Marshal(strconv.Itoa(i/I + 1))
		if err != nil {
			panic(err)
		}
		_, err = writer.Write(b)
		if err != nil {
			panic(err)
		}
	}
	_, err = writer.WriteString("\n}")
	if err != nil {
		panic(err)
	}
	err = writer.Flush()
	if err != nil {
		panic(err)
	}

	//resultMap := make(map[string]string)
	//for i, kv := range lineList {
	//	resultMap[kv.Key] = strconv.Itoa(i/I + 1)
	//}
	//err = json.NewEncoder(gListFile).Encode(resultMap)
	//if err != nil {
	//	panic(err)
	//}
}
