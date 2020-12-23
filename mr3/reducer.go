package mr3

import (
	"bufio"
	"fmt"
	"github.com/wwqgtxx/mfp/common"
	"io"
	"os"
	"strconv"
	"strings"
)

const (
	K = 10 //输出前 K 个
)

func Reducer() {
	currentItem := "1" //当前正在处理的 gid
	var item string
	hp := make(map[string]int) //伪最大堆

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
		if len(split) == 3 {
			item = split[0]
			if len(item) == 0 {
				continue // 跳过reducer的输出
			}

			pattern := split[1]
			supp, err := strconv.Atoi(split[2])
			if err != nil {
				continue
			}

			if item != currentItem {
				if len(currentItem) > 0 {
					//此处输出当前最大堆 top K
					fmt.Printf("\t%s", currentItem)
					for _, k := range common.StringSliceTopN(common.RankMapStringInt(hp), K) {
						//v := hp[k]
						fmt.Printf(" [%s]", k)
					}
					fmt.Print("\n")
				}
				currentItem = item
				hp = make(map[string]int)
			}

			//更新最大堆
			if _, ok := hp[pattern]; ok {
				hp[pattern] += supp
			} else {
				hp[pattern] = supp
			}
		}
	}
	if item == currentItem {
		//此处输出当前最大堆 top K
		fmt.Printf("\t%s", currentItem)
		for _, k := range common.StringSliceTopN(common.RankMapStringInt(hp), K) {
			//v := hp[k]
			fmt.Printf(" [%s]", k)
		}
		fmt.Print("\n")
	}
}
