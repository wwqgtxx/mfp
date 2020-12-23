package mr1

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
		readLine = strings.TrimSpace(readLine)
		items := strings.Fields(readLine) //按空格将句子分割成单个单词
		for _, line := range items {
			fmt.Printf("%s\t%d\n", line, 1)
		}
	}
}
