package environment

import (
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"strings"

	"github.com/BurntSushi/toml"
	"gopkg.in/yaml.v2"
)

type PropertySource struct {
	FilePath string
	Name     string
	Type     reflect.Type
	Map      AnyMap
}

type Environment struct {
	configMap       AnyMap
	propertySources []PropertySource
}

func NewEnvironment(propertySources ...PropertySource) *Environment {
	return &Environment{
		configMap:       make(AnyMap),
		propertySources: propertySources,
	}
}

func (e *Environment) LoadSystemEnvVars() {
	envVars := make(map[string]string)
	for _, env := range os.Environ() {
		pair := strings.SplitN(env, "=", 2)
		if len(pair) == 2 {
			envVars[pair[0]] = pair[1]
		}
	}

	e.LoadSystemEnv(envVars)
}

func (e *Environment) LoadSystemEnv(envVars map[string]string) {
	for key, value := range envVars {
		e.setProperty(key, value)
	}
}

func (e *Environment) LoadConfig(path, name string, _ reflect.Type) error {
	filePath := filepath.Join(path, name)
	ext := filepath.Ext(name)

	data, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	var configMap AnyMap
	switch ext {
	case ".yaml", ".yml":
		err = yaml.Unmarshal(data, &configMap)
	case ".toml":
		err = toml.Unmarshal(data, &configMap)
	default:
		return fmt.Errorf("unsupported config file format: %s", ext)
	}

	if err != nil {
		return err
	}

	cm := convertAnyMap(configMap)
	e.mergeMap(cm.(AnyMap))

	return nil
}

func (e *Environment) LoadMap(configMap AnyMap) {
	e.mergeMap(configMap)
}

func (e *Environment) LoadConfigSources() error {
	for _, source := range e.propertySources {
		if source.FilePath != "" {
			if err := e.LoadConfig(source.FilePath, source.Name, source.Type); err != nil {
				return err
			}
		}

		if source.Type != nil && source.Type.Kind() == reflect.Map {
			e.LoadMapSource(source.Map)
		}
	}

	return nil
}

func (e *Environment) LoadMapSource(sourceMap AnyMap) {
	e.LoadMap(sourceMap)
}

func (e *Environment) StartWithSources(sources ...PropertySource) error {
	e.propertySources = append(e.propertySources, sources...)

	return e.Start()
}

func (e *Environment) Start() error {
	e.LoadSystemEnvVars()

	if err := e.LoadConfigSources(); err != nil {
		return err
	}

	e.configMap = convertIntsToInt64InMap(e.configMap)

	return nil
}

func (e *Environment) Get(key string) any {
	return e.getProperty(key)
}

func (e *Environment) SetProperty(key string, value any) {
	e.setProperty(key, value)
}

func (e *Environment) mergeMap(sourceMap AnyMap) {
	for key, value := range sourceMap {
		if existing, ok := e.configMap[key]; ok {
			if existingMap, ok1 := existing.(AnyMap); ok1 {
				if newMap, ok2 := value.(AnyMap); ok2 {
					mergedMap := make(AnyMap)
					for k, v := range existingMap {
						mergedMap[k] = v
					}
					for k, v := range newMap {
						mergedMap[k] = v
					}
					e.configMap[key] = mergedMap
					continue
				}
			}
		}
		e.configMap[key] = value
	}
}

func (e *Environment) mergeMaps(target AnyMap, source AnyMap) {
	for key, sourceValue := range source {
		targetValue, exists := target[key]

		if exists {
			if isMap(targetValue) && isMap(sourceValue) {
				e.mergeMaps(targetValue.(AnyMap), sourceValue.(AnyMap))
			} else {
				target[key] = sourceValue
			}
		} else {
			target[key] = sourceValue
		}
	}
}

func (e *Environment) getProperty(key string) any {
	keys := strings.Split(key, ".")
	current := e.configMap

	for idx, k := range keys {
		value, ok := current[k]
		if !ok {
			return nil
		}

		if v, ok := value.(map[any]any); ok {
			current = convertAnyMap(v).(AnyMap)
			continue
		}

		if current, ok = value.(AnyMap); !ok {
			if idx == len(keys)-1 {
				return value
			}
		}
	}

	return current
}

func (e *Environment) setProperty(key string, value any) {
	keys := strings.Split(key, ".")
	lastKey := keys[len(keys)-1]
	current := e.configMap
	for _, k := range keys[:len(keys)-1] {
		val, ok := current[k]
		if !ok {
			newMap := make(AnyMap)
			current[k] = newMap
			current = newMap
		} else {
			current = val.(AnyMap)
		}
	}
	current[lastKey] = value
}

func convertAnyMap(source any) any {
	switch v := source.(type) {
	case map[any]any:
		result := make(AnyMap)
		for key, value := range v {
			result[fmt.Sprintf("%v", key)] = convertAnyMap(value)
		}
		return result
	case []any:
		for i, value := range v {
			v[i] = convertAnyMap(value)
		}
	}
	return source
}

func convertIntsToInt64InMap(inputMap AnyMap) AnyMap {
	ctx := make(AnyMap)

	for key, val := range inputMap {
		strKey := fmt.Sprintf("%v", key)

		switch v := val.(type) {
		case AnyMap:
			ctx[strKey] = convertIntsToInt64InMap(v)
		case map[any]any:
			tmpMap := convertAnyMap(v).(AnyMap)
			ctx[strKey] = convertIntsToInt64InMap(tmpMap)
		case []any:
			ctx[strKey] = convertIntsToInt64InSlice(v)
		case int:
			ctx[strKey] = int64(v)
		default:
			ctx[strKey] = val
		}
	}

	return ctx
}

func convertIntsToInt64InSlice(inputSlice []any) []any {
	result := make([]any, len(inputSlice))

	for i, val := range inputSlice {
		switch v := val.(type) {
		case AnyMap:
			result[i] = convertIntsToInt64InMap(v)
		case map[any]any:
			tmpMap := convertAnyMap(v).(AnyMap)
			result[i] = convertIntsToInt64InMap(tmpMap)
		case []any:
			result[i] = convertIntsToInt64InSlice(v)
		case int:
			result[i] = int64(v)
		default:
			result[i] = val
		}
	}

	return result
}

func isMap(value any) bool {
	_, ok := value.(AnyMap)
	return ok
}
