package iterator

import (
	`reflect`
	`testing`
)

func Test_iterator_Next(t *testing.T) {
	type user struct {
		name string
		age  int
	}

	type fields struct {
		index int
		array []any
	}
	tests := []struct {
		name   string
		fields fields
		want   any
	}{
		{
			name: "Test iterator Next-index-0",
			fields: fields{
				index: 0,
				array: []any{
					user{
						name: "tom",
						age:  30,
					}, user{
						name: "jerry",
						age:  24,
					},
				},
			},
			want: user{
				name: "tom",
				age:  30,
			},
		},
		{
			name: "Test iterator Next-index-1",
			fields: fields{
				index: 1,
				array: []any{
					user{
						name: "tom",
						age:  30,
					}, user{
						name: "jerry",
						age:  24,
					},
				},
			},
			want: user{
				name: "jerry",
				age:  24,
			},
		},
		{
			name: "Test iterator Next-index-1",
			fields: fields{
				index: 2,
				array: []any{
					user{
						name: "tom",
						age:  30,
					}, user{
						name: "jerry",
						age:  24,
					},
				},
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &iterator{
				index: tt.fields.index,
				array: tt.fields.array,
			}
			if got := u.Next(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Next() = %v, want %v", got, tt.want)
			}
		})
	}
}
