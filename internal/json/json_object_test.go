package json

import (
	`reflect`
	`testing`
)

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
				mvp: tt.fields.context,
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
				mvp: tt.fields.context,
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
				mvp: tt.fields.context,
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
				mvp: tt.fields.context,
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
				mvp: tt.fields.context,
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

func TestNewJsoonObjects(t *testing.T) {
	type args struct {
		mv map[string]any
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test new JSOONObject-with-map",
			args: args{
				mv: map[string]any{
					"hello": "world",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			jsoon := NewJsoonObjects(tt.args.mv)
			value, ok := jsoon.GetSafe("hello", "")
			if !ok || value != "world" {
				t.Errorf("GetInt64() got1 = %v, want %v", value, "world")
			}
		})
	}
}

var body = `{
  "id": "9787111558422",
  "name": "The Go Programming Language",
  "authors": [
    "Alan A.A.Donovan",
    "Brian W. Kergnighan"
  ],
  "press": "Pearson Education"
}`

func TestParseJSONObject(t *testing.T) {
	type args struct {
		body string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test new JSOONObject",
			args: args{
				body: body,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseJSONObject(tt.args.body)
			if err != nil {
				t.Errorf("ParseJSONObject error:%v", err)
			}
			safe, ok := got.GetSafe("id", "")
			if !ok || "9787111558422" != safe {
				t.Errorf("ParseJSONObject and GetSafe() error:%v", err)
			}
		})
	}
}
