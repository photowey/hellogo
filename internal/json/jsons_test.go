package json

import (
	"io"
	`reflect`
	"strings"
	"testing"
)

// Book
type Book struct {
	Id      string   `json:"id"`      // 图书 ISBN ID
	Name    string   `json:"name"`    // 图书名称
	Authors []string `json:"authors"` // 图书作者
	Press   string   `json:"press"`   // 出版社
}

var jsonData = `{
  "id": "9787111558422",
  "name": "The Go Programming Language",
  "authors": [
    "Alan A.A.Donovan",
    "Brian W. Kergnighan"
  ],
  "press": "Pearson Education"
}`

func TestToStruct(t *testing.T) {
	type args struct {
		data   []byte
		target any
	}
	var book *Book = &Book{}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Test json string to struct(Unmarshal)",
			args: args{
				data:   []byte(jsonData),
				target: book,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ToStruct(tt.args.data, tt.args.target); (err != nil) != tt.wantErr {
				t.Errorf("ToStruct() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestToStructd(t *testing.T) {
	type args struct {
		reader io.Reader
		target any
	}
	var book *Book = &Book{}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Test json string to struct(Decode)",
			args: args{
				reader: strings.NewReader(jsonData),
				target: book,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ToStructd(tt.args.reader, tt.args.target); (err != nil) != tt.wantErr {
				t.Errorf("ToStructd() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestJSOONObject_Put(t *testing.T) {
	type fields struct {
		context map[string]any
	}
	type args struct {
		key   string
		value any
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "Test JSOONObject Put()",
			fields: fields{
				context: make(map[string]any),
			},
			args: args{
				key:   "hello",
				value: "world",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			jsoon := &JSOONObject{
				context: tt.fields.context,
			}
			jsoon.Put(tt.args.key, tt.args.value)
		})
	}
}

func TestJSOONObject_Get(t *testing.T) {
	type fields struct {
		context map[string]any
	}
	type args struct {
		key string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   any
	}{
		{
			name: "Test JSOONObject Get()",
			fields: fields{
				context: map[string]any{
					"hello": "world",
				},
			},
			args: args{
				key: "hello",
			},
			want: "world",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			jsoon := &JSOONObject{
				context: tt.fields.context,
			}
			if got := jsoon.Get(tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJSOONObject_GetSafe(t *testing.T) {
	type fields struct {
		context map[string]any
	}
	type args struct {
		key     string
		standBy any
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   any
		ok     bool
	}{
		{
			name: "Test JSOONObject GetSafe()",
			fields: fields{
				context: map[string]any{
					"hello": "world",
				},
			},
			args: args{
				key:     "hello",
				standBy: "standBy",
			},
			want: "world",
			ok:   true,
		},

		{
			name: "Test JSOONObject GetSafe()",
			fields: fields{
				context: map[string]any{
					"hello": "world",
				},
			},
			args: args{
				key:     "helloStandBy",
				standBy: "standBy",
			},
			want: "standBy",
			ok:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			jsoon := &JSOONObject{
				context: tt.fields.context,
			}
			got, got1 := jsoon.GetSafe(tt.args.key, tt.args.standBy)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetSafe() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.ok {
				t.Errorf("GetSafe() got1 = %v, want %v", got1, tt.ok)
			}
		})
	}
}

func TestJSOONObject_GetString(t *testing.T) {
	type fields struct {
		context map[string]any
	}
	type args struct {
		key string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
		ok     bool
	}{
		{
			name: "Test JSOONObject GetString()-true",
			fields: fields{
				context: map[string]any{
					"hello": "world",
				},
			},
			args: args{
				key: "hello",
			},
			want: "world",
			ok:   true,
		},
		{
			name: "Test JSOONObject GetString()-false",
			fields: fields{
				context: map[string]any{
					"hello": 8848,
				},
			},
			args: args{
				key: "hello",
			},
			want: "",
			ok:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			jsoon := &JSOONObject{
				context: tt.fields.context,
			}
			got, got1 := jsoon.GetString(tt.args.key)
			if got != tt.want {
				t.Errorf("GetString() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.ok {
				t.Errorf("GetString() got1 = %v, want %v", got1, tt.ok)
			}
		})
	}
}

func TestJSOONObject_GetInt64(t *testing.T) {
	type fields struct {
		context map[string]any
	}
	type args struct {
		key string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int64
		ok     bool
	}{
		{
			name: "Test JSOONObject GetInt64()-int64-true",
			fields: fields{
				context: map[string]any{
					"hello": int64(1234567890912345678),
				},
			},
			args: args{
				key: "hello",
			},
			want: 1234567890912345678,
			ok:   true,
		},
		{
			name: "Test JSOONObject GetInt64()-int-true",
			fields: fields{
				context: map[string]any{
					"hello": 1234567890987654321,
				},
			},
			args: args{
				key: "hello",
			},
			want: 1234567890987654321,
			ok:   true,
		},
		{
			name: "Test JSOONObject GetInt64()-int8-true",
			fields: fields{
				context: map[string]any{
					"hello": int8(127),
				},
			},
			args: args{
				key: "hello",
			},
			want: 127,
			ok:   true,
		},
		{
			name: "Test JSOONObject GetInt64()-int32-true",
			fields: fields{
				context: map[string]any{
					"hello": int32(8848),
				},
			},
			args: args{
				key: "hello",
			},
			want: 8848,
			ok:   true,
		},
		{
			name: "Test JSOONObject GetInt64()-string-false",
			fields: fields{
				context: map[string]any{
					"hello": "world",
				},
			},
			args: args{
				key: "hello",
			},
			want: 0,
			ok:   false,
		},
		{
			name: "Test JSOONObject GetInt64()-float-false",
			fields: fields{
				context: map[string]any{
					"hello": 1.01,
				},
			},
			args: args{
				key: "hello",
			},
			want: 1,
			ok:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			jsoon := &JSOONObject{
				context: tt.fields.context,
			}
			got, got1 := jsoon.GetInt64(tt.args.key)
			if got != tt.want {
				t.Errorf("GetInt64() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.ok {
				t.Errorf("GetInt64() got1 = %v, want %v", got1, tt.ok)
			}
		})
	}
}
