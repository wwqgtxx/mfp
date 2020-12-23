package to_db

import (
	"bufio"
	"fmt"
	"github.com/go-ego/gse"
	"github.com/wwqgtxx/mfp/common"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func ToDB() {
	var seg gse.Segmenter
	err := seg.LoadDict("dictionary.txt")
	if err != nil {
		panic(err)
	}

	if len(os.Args) < 4 {
		fmt.Println("expected file path and target path")
		os.Exit(1)
	}
	files, err := filepath.Glob(os.Args[2] + "/*.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(files)

	re, err := regexp.Compile("[+\\s]")
	if err != nil {
		panic(err)
	}

	countTotal := 0
	for _, fileName := range files {
		file, err := os.Open(fileName)
		if err != nil {
			panic(err)
		}
		defer file.Close()

		stat, err := file.Stat()
		if err != nil {
			panic(err)
		}
		fmt.Println(stat.Name())

		targetFile, err := os.Create(os.Args[3] + "/" + strings.TrimSuffix(stat.Name(), ".txt") + "_db.txt")
		if err != nil {
			panic(err)
		}
		defer targetFile.Close()

		reader := bufio.NewReader(file)

		uid2keywords := make(map[string][]string)

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
			if len(split) != 6 {
				continue
			}
			uid := split[1]
			keywords := split[2]
			keywords = strings.TrimPrefix(keywords, "[")
			keywords = strings.TrimSuffix(keywords, "]") //去掉中括号

			//通过加号和空白符分开后，再用 jieba 搜索引擎模式分词
			keywordsTempArr := make([]string, 0, len(keywords))
			for _, s := range re.Split(keywords, -1) {
				s = strings.TrimSpace(s)
				if len(s) > 0 {
					keywordsTempArr = append(keywordsTempArr, s)
				}
			}
			keywordsTemp := strings.Join(keywordsTempArr, " ")
			keywordsList := seg.Slice(keywordsTemp, false)

			if uid2keyword, ok := uid2keywords[uid]; !ok {
				uid2keywords[uid] = keywordsList
			} else {
				for _, kw := range keywordsList {
					if !common.StringSliceContains(uid2keyword, kw) {
						uid2keywords[uid] = append(uid2keyword, kw)
					}
				}
			}

		}

		count := 0
		for _, v := range uid2keywords {
			if len(v) > 1 { //只有一个关键词的就不用留下了
				_, err := targetFile.WriteString("\t" + strings.Join(v, " ") + "\n") //事务集 加 \t 是因为第一个 Map 输入键为空
				if err != nil {
					panic(err)
				}
				count++
			}
		}
		fmt.Printf("事务条数： %d\n", count)
		countTotal += count
	}
	fmt.Printf("总和： 事务条数： %d\n", countTotal)
}
