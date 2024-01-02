package mapz

import (
	"fmt"
	"maps"
)

type Programmer struct {
	Name string
	Age  int
}

func HelloMap() {

	m1 := map[string]Programmer{
		"zhangsan": {
			Name: "zhangsan",
			Age:  18,
		},
		"lisi": {
			Name: "lisi",
			Age:  19,
		},
		"wangwu": {
			Name: "wangwu",
			Age:  20,
		},
	}

	m11 := map[string]Programmer{}

	m2 := maps.Clone(m1)

	maps.Copy(m1, m11)

	fmt.Printf("m1: %v\n", m1)
	m2["zhangsan"] = Programmer{
		Name: "zhaoliu",
		Age:  21,
	}
	fmt.Printf("m2: %v\n", m2)
	fmt.Printf("m1: %v\n", m1)
	fmt.Printf("m11: %v\n", m11)
}
