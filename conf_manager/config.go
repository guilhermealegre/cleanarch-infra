package config_manager

import (
	"auth/infra/http"
	"auth/infra/postgres"
	"auth/infra/redis"
	"fmt"
	"github.com/go-yaml/yaml"
	"os"
	"reflect"
	"strconv"
	"strings"
)

type Config struct {
	configFile string
	Http       http.Config     `yaml:"http" json:"http"`
	Postgres   postgres.Config `yaml:"postgres" json:"postgres"`
	Redis      redis.Config    `yaml:"redis" json:"redis"`
}

func NewConfig(configFile string) *Config {
	return &Config{configFile: configFile}
}

func (c *Config) Load() (*Config, error) {

	file, err := os.ReadFile(c.configFile)
	if err != nil {
		panic(err)
	}

	var mapConfigs map[any]any

	err = yaml.Unmarshal(file, &mapConfigs)
	if err != nil {
		panic(err)
	}

	// find all environment variable and get the value
	// but is very complex
	f := func(input any) (canUpdate bool, output any, err error) {
		v := input.(string)
		if strings.HasPrefix(v, "${") && strings.HasSuffix(v, "}") {
			output = os.Getenv(strings.TrimPrefix(strings.TrimSuffix(v, "}"), "${"))
			canUpdate = true
		}

		return canUpdate, output, err
	}

	if err := walkMap(mapConfigs, f, reflect.String); err != nil {
		return c, err
	}

	byteConfigs, err := yaml.Marshal(mapConfigs)
	if err != nil {
		panic(err.Error())
	}

	if err := yaml.Unmarshal(byteConfigs, c); err != nil {
		panic(err)
	}

	return c, nil

}

// I need review this function
func walkMap(m any, f func(any) (bool, any, error), targetKind reflect.Kind) error {
	inputMap, ok := m.(map[any]any)
	if !ok {
		return fmt.Errorf("Error in cast map")
	}
	for k, v := range inputMap {
		if reflect.ValueOf(v).Kind() == reflect.Map {
			if err := walkMap(v, f, targetKind); err != nil {
				return err
			}
		} else {
			if targetKind == reflect.ValueOf(v).Kind() {
				canUpdate, newValue, err := f(v)
				if err != nil {
					return err
				}

				if canUpdate {

					nv, ok := newValue.(string)
					if ok {
						if nvInt, err := strconv.Atoi(nv); err == nil {
							inputMap[k] = nvInt

						} else {
							inputMap[k] = nv
						}
					}
				}

			}
		}
	}

	return nil
}
