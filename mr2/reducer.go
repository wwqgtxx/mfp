package mr2

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
	K = 50 //输出前 K 个
)

func Reducer() {
	root := common.NewNode()
	currentGid := "1" //当前正在处理的 gid
	var gid string
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
		if len(split) == 2 {
			gid = split[0]
			if strings.Contains(gid, " ") {
				continue // 跳过reducer的输出
			}

			_, err := strconv.Atoi(gid)
			if err != nil {
				continue
			}

			items := strings.Fields(split[1])

			if len(items) == 0 {
				continue
			}

			if gid != currentGid {
				//此处通过当前的树弄出最大堆
				root.Scan(hp, make([]string, 0))

				root = common.NewNode()
				currentGid = gid

				//此处输出当前最大堆 top K
				for _, k := range common.StringSliceTopN(common.RankMapStringInt(hp), K) {
					v := hp[k]
					fmt.Printf("%s\t%d\n", k, v)
				}

				hp = make(map[string]int)
			}

			//此处建树
			p := root.GetSelf()
			for _, item := range items {
				child := p.FindChild(item)
				if child != nil {
					child.AddCount()
					p = child.GetSelf()
				} else {
					p = p.AddChild(common.NewNodeWithItem(item))
				}
			}
		}
	}
	if gid == currentGid {
		//此处输出当前最大堆 top K
		for _, k := range common.StringSliceTopN(common.RankMapStringInt(hp), K) {
			v := hp[k]
			fmt.Printf("%s\t%d\n", k, v)
		}
	}
}
