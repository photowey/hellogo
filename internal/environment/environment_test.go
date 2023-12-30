package environment

import (
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

type AppConfig struct {
	App struct {
		Name    string
		Version float64
		Details struct {
			Enabled     bool
			Description string
		}
	}
	Database struct {
		Host     string
		Port     int64
		Username string
		Password string
	}
}

func TestEnvironmentWithSources(t *testing.T) {
	testConfigDir := "testdata"
	err := os.Mkdir(testConfigDir, os.ModePerm)
	assert.NoError(t, err)
	defer os.RemoveAll(testConfigDir)

	yamlConfigContent := `
app:
  name: YAMLApp
  version: 3.0
  details:
    enabled: true
    description: "YAML application description"
database:
  host: yamlhost
  port: 5434
  username: yamluser
  password: yamlpassword
`
	yamlConfigFile := filepath.Join(testConfigDir, "test_yaml_config.yaml")
	err = os.WriteFile(yamlConfigFile, []byte(yamlConfigContent), os.ModePerm)
	assert.NoError(t, err)

	tomlConfigContent := `
[tomalapp]
  name = "TOMLApp"
  version = 4.0

[tomalapp.details]
  enabled = true
  description = "TOML application description"

[tomaldatabase]
  host = "tomlhost"
  port = 5435
  username = "tomluser"
  password = "tomlpassword"
`
	tomlConfigFile := filepath.Join(testConfigDir, "test_toml_config.toml")
	err = os.WriteFile(tomlConfigFile, []byte(tomlConfigContent), os.ModePerm)
	assert.NoError(t, err)

	mapConfig := AnyMap{
		"mapapp": AnyMap{
			"name":    "MapApp",
			"version": 6.0,
			"details": AnyMap{
				"enabled":     true,
				"description": "Map application description",
			},
		},
		"mapdatabase": AnyMap{
			"host":     "maphost",
			"port":     5436,
			"username": "mapuser",
			"password": "mappassword",
		},
	}

	env := NewEnvironment()
	err = env.StartWithSources(
		PropertySource{FilePath: testConfigDir, Name: "test_yaml_config.yaml"},
		PropertySource{FilePath: testConfigDir, Name: "test_toml_config.toml"},
		PropertySource{Type: reflect.TypeOf(AnyMap{}), Map: mapConfig},
	)
	assert.NoError(t, err)

	assert.Equal(t, "YAMLApp", env.Get("app.name"))
	assert.Equal(t, 3.0, env.Get("app.version"))
	assert.Equal(t, true, env.Get("app.details.enabled"))
	assert.Equal(t, "YAML application description", env.Get("app.details.description"))
	assert.Equal(t, "yamlhost", env.Get("database.host"))
	assert.Equal(t, int64(5434), env.Get("database.port"))
	assert.Equal(t, "yamluser", env.Get("database.username"))
	assert.Equal(t, "yamlpassword", env.Get("database.password"))

	assert.Equal(t, "TOMLApp", env.Get("tomalapp.name"))
	assert.Equal(t, 4.0, env.Get("tomalapp.version"))
	assert.Equal(t, true, env.Get("tomalapp.details.enabled"))
	assert.Equal(t, "TOML application description", env.Get("tomalapp.details.description"))
	assert.Equal(t, "tomlhost", env.Get("tomaldatabase.host"))
	assert.Equal(t, int64(5435), env.Get("tomaldatabase.port"))
	assert.Equal(t, "tomluser", env.Get("tomaldatabase.username"))
	assert.Equal(t, "tomlpassword", env.Get("tomaldatabase.password"))

	assert.Equal(t, "MapApp", env.Get("mapapp.name"))
	assert.Equal(t, 6.0, env.Get("mapapp.version"))
	assert.Equal(t, true, env.Get("mapapp.details.enabled"))
	assert.Equal(t, "Map application description", env.Get("mapapp.details.description"))
	assert.Equal(t, "maphost", env.Get("mapdatabase.host"))
	assert.Equal(t, int64(5436), env.Get("mapdatabase.port"))
	assert.Equal(t, "mapuser", env.Get("mapdatabase.username"))
	assert.Equal(t, "mappassword", env.Get("mapdatabase.password"))
}
