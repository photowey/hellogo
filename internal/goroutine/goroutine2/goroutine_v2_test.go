package goroutine2

import (
	"testing"
	"time"
)

func TestGoroutine_Start(t *testing.T) {
	type fields struct {
		options []any
		fx      func(parameters ...any)
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "Test goroutine run()",
			fields: fields{
				options: []any{"hello", "world", 3, struct{ name string }{name: "sharkchili"}},
				fx: func(parameters ...any) {
					t.Logf("executing the goroutine callback fx,the parametes is:%v", parameters)
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actor := Goroutine{
				options: tt.fields.options,
				fx:      tt.fields.fx,
			}
			actor.Start()
		})
	}

	time.Sleep(time.Duration(1) * time.Second)
}

func TestGoroutine_Startpre(t *testing.T) {
	type fields struct {
		options []any
		fx      func(parameters ...any)
	}
	type args struct {
		pre func()
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Goroutine
	}{
		{
			name: "Test goroutine pre() run()",
			fields: fields{
				options: []any{"hello", "world", 3, struct{ name string }{name: "sharkchili"}},
				fx: func(parameters ...any) {
					t.Logf("executing the goroutine callback fx,the parametes is:%v", parameters)
				},
			},
			args: args{
				pre: func() {
					t.Logf("exeucte bedore start and exeucting the goroutine")
				},
			},
			want: Goroutine{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actor := Goroutine{
				options: tt.fields.options,
				fx:      tt.fields.fx,
			}
			__ := actor.Startpre(tt.args.pre)
			_ = __
		})
	}

	time.Sleep(time.Duration(1) * time.Second)
}

func TestGoroutine_Startpost(t *testing.T) {
	type fields struct {
		options []any
		fx      func(parameters ...any)
	}
	type args struct {
		post func()
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Goroutine
	}{
		{
			name: "Test goroutine post() run()",
			fields: fields{
				options: []any{"hello", "world", 3, struct{ name string }{name: "sharkchili"}},
				fx: func(parameters ...any) {
					t.Logf("executing the goroutine callback fx,the parametes is:%v", parameters)
				},
			},
			args: args{
				post: func() {
					t.Logf("exeucte after start and exeucting the goroutine")
				},
			},
			want: Goroutine{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actor := Goroutine{
				options: tt.fields.options,
				fx:      tt.fields.fx,
			}
			__ := actor.Startpost(tt.args.post)
			_ = __
		})
	}

	time.Sleep(time.Duration(1) * time.Second)
}

func TestGoroutine_Startaround(t *testing.T) {
	type fields struct {
		options []any
		fx      func(parameters ...any)
	}
	type args struct {
		pre  func()
		post func()
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Goroutine
	}{
		{
			name: "Test goroutine around() run()",
			fields: fields{
				options: []any{"hello", "world", 3, struct{ name string }{name: "sharkchili"}},
				fx: func(parameters ...any) {
					t.Logf("executing the goroutine callback fx,the parametes is:%v", parameters)
				},
			},
			args: args{
				pre: func() {
					t.Logf("exeucte bedore start and exeucting the goroutine")
				},
				post: func() {
					t.Logf("exeucte after start and exeucting the goroutine")
				},
			},
			want: Goroutine{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actor := Goroutine{
				options: tt.fields.options,
				fx:      tt.fields.fx,
			}
			__ := actor.Startaround(tt.args.pre, tt.args.post)
			_ = __
		})
	}
}
