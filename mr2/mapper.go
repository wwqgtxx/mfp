package mr2

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/wwqgtxx/mfp/common"
	"io"
	"os"
	"strings"
)

func getGList() map[string]string {
	gList := make(map[string]string)
	file, err := os.Open("gList.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	err = json.NewDecoder(file).Decode(&gList)
	if err != nil {
		panic(err)
	}
	return gList
}

func Mapper() {
	gList := getGList()
	usedGid := make([]string, 0)
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
		readLine = strings.TrimSpace(readLine)
		items := strings.Fields(readLine) //按空格将句子分割成单个单词
		for i := len(items) - 1; i >= 0; i-- {
			gid := gList[items[i]]
			if common.StringSliceContains(usedGid, gid) {
				continue
			} else {
				usedGid = append(usedGid, gid)
				fmt.Printf("%s\t%s\n", gid, strings.Join(items[:i+1], " "))
			}
		}
	}
}
