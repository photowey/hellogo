package environment

import (
	"reflect"
	"strings"
)

const (
	dot       = "."
	emptyStr  = ""
	binderTag = "binder"
)

type Binder struct {
	Prefix string
}

func New() *Binder {
	return &Binder{}
}

func NewBinder(prefix string) *Binder {
	return &Binder{Prefix: prefix}
}

func (b *Binder) DefaultBind(target any, ctx AnyMap) {
	b.Bind(b.Prefix, target, ctx)
}

func (b *Binder) Bind(prefix string, target any, ctx AnyMap) {
	if prefix != emptyStr && !strings.HasSuffix(prefix, dot) {
		prefix += dot
	}

	tt := reflect.TypeOf(target).Elem()
	tv := reflect.ValueOf(target).Elem()

	for i := 0; i < tt.NumField(); i++ {
		fieldType := tt.Field(i)
		fieldValue := tv.Field(i)

		tag := fieldType.Tag.Get(binderTag)
		key := prefix + strings.ToLower(tag)

		if fieldType.Type.Kind() == reflect.Struct {
			sub := reflect.New(fieldType.Type).Interface()
			b.Bind(key, sub, ctx)
			fieldValue.Set(reflect.ValueOf(sub).Elem())
		} else {
			value := getProperty(ctx, key)

			if value != nil {
				fieldValue.Set(reflect.ValueOf(value).Convert(fieldType.Type))
			}
		}
	}
}

func getProperty(ctx AnyMap, key string) any {
	keys := strings.Split(key, dot)

	for _, k := range keys {
		if value, ok := ctx[k]; ok {
			switch v := value.(type) {
			case AnyMap:
				ctx = v
			default:
				return value
			}
		} else {
			return nil
		}
	}

	return nil
}
