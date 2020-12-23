package sort_kv

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/wwqgtxx/mfp/common"
	"io"
	"os"
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
	lineMap := make(map[string]int)

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
			lineMap[pattern] = supp
		}
	}
	itemsCount := len(lineMap)
	fmt.Printf("项目总数：%d\n", itemsCount)

	resultMap := make(map[string]string)
	for i, k := range common.RankMapStringInt(lineMap) {
		//v := lineMap[k]
		resultMap[k] = strconv.Itoa(i/I + 1)
	}
	gListFile, err := os.Create(os.Args[3])
	if err != nil {
		panic(err)
	}
	defer gListFile.Close()
	err = json.NewEncoder(gListFile).Encode(resultMap)
	if err != nil {
		panic(err)
	}
}
