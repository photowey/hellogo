package json

import (
	"io"
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
