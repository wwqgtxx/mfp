package mr3

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func Mapper() {
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
			pattern := split[0]
			//supp := split[1]
			items := strings.Fields(pattern) //按空格将句子分割成单个单词
			for _, item := range items {
				fmt.Printf("%s\t%s\n", item, readLine) //输出格式： item \t pattern \t supp
			}
		}
	}
}
