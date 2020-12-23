package format_file

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/wwqgtxx/mfp/common"
)

func NewQ2B() *strings.Replacer { //全角转半角，制表符转空格
	Q := []rune("ａｂｃｄｅｆｇｈｉｊｋｌｍｎｏｐｑｒｓｔｕｖｗｘｙｚＡＢＣＤＥＦＧＨＩＪＫＬＭＮＯＰＱＲＳＴＵＶＷＸＹＺ０１２３４５６７８９\t")
	B := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789 ")
	list := make([]string, 0, 62*2)
	for i, q := range Q {
		b := B[i]
		list = append(list, string(q), string(b))
	}
	return strings.NewReplacer(list...)
}

func FormatFile() {
	if len(os.Args) < 4 {
		fmt.Println("expected file path and target path")
		os.Exit(1)
	}
	files, err := filepath.Glob(os.Args[2] + "/*.decode.filter.utf8")
	if err != nil {
		panic(err)
	}
	fmt.Println(files)

	//按照搜狗格式处理的正则表达式
	re, err := regexp.Compile("^([A-Za-z0-9_:]+)\\s(\\d+)\\s(\\[.+\\])\\s(\\d+)\\s(\\d+)\\s([,:+~^|;!$*{}\"><\\'\\`@_A-Za-z0-9\\-\\[\\]\\(\\)\\\\./%=?&]+)$")
	if err != nil {
		panic(err)
	}
	replacer := NewQ2B()
	successCountTotal, failCountTotal := 0, 0

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

		targetFile, err := os.Create(os.Args[3] + "/" + strings.TrimRight(stat.Name(), ".decode.filter.utf8") + ".txt")
		if err != nil {
			panic(err)
		}
		defer targetFile.Close()

		reader := bufio.NewReader(file)

		successCount := 0
		failCount := 0

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

			groups := re.FindStringSubmatch(readLine)
			if len(groups) > 1 {
				groups = groups[1:]
			}
			//检查字段数是否正确、各项格式是否正确
			if len(groups) != 6 ||
				!strings.HasPrefix(groups[2], "[") ||
				!strings.HasSuffix(groups[2], "]") ||
				!common.IsNumeric(groups[3]) ||
				!common.IsNumeric(groups[4]) ||
				!strings.Contains(groups[5], ".") {
				failCount++
				continue
			}

			//忽略陋俗
			if strings.Contains(groups[2], "陋俗") {
				failCount++
				continue
			}

			groups[2] = replacer.Replace(groups[2]) //去掉关键词里的 tab，把字母和数字全角转半角
			successCount++

			_, err = targetFile.WriteString(strings.Join(groups[:6], "\t") + "\n")
			if err != nil {
				panic(err)
			}
		}
		fmt.Printf("有效数据条数： %d， 丢弃条数： %d\n", successCount, failCount)
		successCountTotal += successCount
		failCountTotal += failCount
	}
	fmt.Printf("总和： 有效数据条数： %d， 丢弃条数： %d\n", successCountTotal, failCountTotal)
}
