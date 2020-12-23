package gb2utf8

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"golang.org/x/text/encoding/simplifiedchinese"
)

func GB2UTF8() {
	if len(os.Args) < 4 {
		fmt.Println("expected file path and target path")
		os.Exit(1)
	}
	files, err := filepath.Glob(os.Args[2] + "/*")
	if err != nil {
		panic(err)
	}
	fmt.Println(files)

	decoder := simplifiedchinese.GB18030.NewDecoder()
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

		targetFile, err := os.Create(os.Args[3] + "/" + stat.Name() + ".utf8")
		if err != nil {
			panic(err)
		}
		defer targetFile.Close()

		decodedFile := decoder.Reader(file)

		i, err := io.Copy(targetFile, decodedFile)
		if err != nil {
			panic(err)
		}
		fmt.Printf("从【%s】往【%s】复制了%d个字节\n", fileName, targetFile.Name(), i)
	}
}
