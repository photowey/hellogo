package sort

import "testing"

func TestSort(t *testing.T) {
	type Student struct {
		Id    int64
		Name  string
		Score float64
	}

	type args[T any] struct {
		sorter  []T
		swapper func(T, T) bool
	}

	type testT[A any] struct {
		name string
		args args[A]
	}

	students := []Student{
		{
			Id:    9527,
			Name:  "Lilei",
			Score: 98.76,
		},
		{
			Id:    8848,
			Name:  "Hanmeimei",
			Score: 88.76,
		},
		{
			Id:    7923,
			Name:  "Tom",
			Score: 78.76,
		},
		{
			Id:    6379,
			Name:  "Jerry",
			Score: 68.76,
		},
	}

	tests := []testT[Student]{
		{
			name: "Test Student sort-true",
			args: args[Student]{
				sorter: students,
				swapper: func(s1, s2 Student) bool {
					return s1.Score < s2.Score
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Sort(tt.args.sorter, tt.args.swapper)

			/*
				=== RUN   TestSort
				=== RUN   TestSort/Test_Student_sort-true
				    sorts_test.go:62: the student:[Jerry] score is:[68.760000]
				    sorts_test.go:62: the student:[Tom] score is:[78.760000]
				    sorts_test.go:62: the student:[Hanmeimei] score is:[88.760000]
				    sorts_test.go:62: the student:[Lilei] score is:[98.760000]
				--- PASS: TestSort (0.00s)
				    --- PASS: TestSort/Test_Student_sort-true (0.00s)
				PASS
			*/
			for _, v := range tt.args.sorter {
				t.Logf("the student:[%s] score is:[%f]", v.Name, v.Score)
			}
		})
	}
}
