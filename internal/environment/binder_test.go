package environment

import (
	"reflect"
	"testing"
)

type Sub struct {
	X string `binder:"x"`
	Y int    `binder:"y"`
}

type Main struct {
	A   string  `binder:"d"`
	B   int     `binder:"e"`
	C   bool    `binder:"f"`
	Z   float64 `binder:"g.h"`
	Sub Sub     `binder:"sub"`
}

func TestBinder_DefaultBind(t *testing.T) {
	properties := AnyMap{
		"a": AnyMap{
			"b": AnyMap{
				"c": AnyMap{
					"d": "Hello",
					"e": 42,
					"f": true,
					"g": AnyMap{
						"h": 3.14,
					},
					"sub": AnyMap{
						"x": "Nested",
						"y": 123,
					},
				},
			},
		},
	}

	tests := []struct {
		name     string
		prefix   string
		expected Main
	}{
		{
			name:     "Test Case 1",
			prefix:   "a.b.c",
			expected: Main{A: "Hello", B: 42, C: true, Z: 3.14, Sub: Sub{X: "Nested", Y: 123}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			config := Main{}
			binder := &Binder{
				Prefix: tt.prefix,
			}
			binder.DefaultBind(&config, properties)

			if !reflect.DeepEqual(config, tt.expected) {
				t.Errorf("Expected %+v, but got %+v", tt.expected, config)
			}
		})
	}
}

func TestBinder_Bind(t *testing.T) {
	properties := AnyMap{
		"a": AnyMap{
			"b": AnyMap{
				"c": AnyMap{
					"d": "Hello",
					"e": 42,
					"f": true,
					"g": AnyMap{
						"h": 3.14,
					},
					"sub": AnyMap{
						"x": "Nested",
						"y": 123,
					},
				},
			},
		},
	}

	tests := []struct {
		name     string
		prefix   string
		expected Main
	}{
		{
			name:     "Test Case 1",
			prefix:   "a.b.c",
			expected: Main{A: "Hello", B: 42, C: true, Z: 3.14, Sub: Sub{X: "Nested", Y: 123}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			config := Main{}
			binder := &Binder{}
			binder.Bind(tt.prefix, &config, properties)

			if !reflect.DeepEqual(config, tt.expected) {
				t.Errorf("Expected %+v, but got %+v", tt.expected, config)
			}
		})
	}
}
