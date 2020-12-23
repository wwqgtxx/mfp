package combine_part

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func CombinePart() {
	if len(os.Args) < 4 {
		fmt.Println("expected file path and target name")
		os.Exit(1)
	}
	files, err := filepath.Glob(os.Args[2] + "/part-*")
	if err != nil {
		panic(err)
	}
	fmt.Println(files)
	targetFile, err := os.Create(os.Args[3])
	if err != nil {
		panic(err)
	}
	defer targetFile.Close()
	for _, fileName := range files {
		func() {
			file, err := os.Open(fileName)
			if err != nil {
				panic(err)
			}
			defer file.Close()
			i, err := io.Copy(targetFile, file)
			if err != nil {
				panic(err)
			}
			fmt.Printf("从【%s】往【%s】复制了%d个字节", fileName, targetFile.Name(), i)
		}()
	}
}
