package common

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func Match(expect bool, standard, defaultCandidate string) string {
	if expect {
		return standard
	}
	return defaultCandidate
}

func Exists(path string) bool {
	_, err := os.Stat(path)

	return err != nil || os.IsExist(err)
}

// ListFiles 列出目录树
func ListFiles(dir string, level int) {
	tree := "|--"
	for i := 0; i < level; i++ {
		tree = "|   " + tree
	}

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		fileName := dir + "/" + file.Name()
		fmt.Printf("%s%s\n", tree, fileName)
		if file.IsDir() {
			ListFiles(fileName, level+1)
		}
	}
}
