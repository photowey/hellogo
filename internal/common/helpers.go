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

func UnSafeZeroValue(target any) any {
	switch t := target.(type) {
	case int8, int16, int32, int64, uint, byte, uint16, uint32, uint64:
		return 0
	case uintptr:
		return 0
	case float32, float64:
		return 0.0
	case []byte:
		return []byte{}
	case string:
		return EmptyString
	case *string:
		return *t
	default:
		return nil
	}
}
