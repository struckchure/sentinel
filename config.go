package sentinel

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/invopop/jsonschema"
	"gopkg.in/yaml.v3"
)

type IConfigLoader interface {
	Load(path string, configType ConfigType) (*Config, error)
	LoadJson(path string) (*Config, error)
	LoadYaml(path string) (*Config, error)
	Schema(save bool, indentation int, output string) error
}

type ConfigLoader struct{}

type ConfigType string

const (
	ConfigTypeYaml ConfigType = "yaml"
	ConfigTypeJson ConfigType = "json"
)

func (c *ConfigLoader) LoadYaml(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var config Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, err
	}
	return &config, nil
}

func (c *ConfigLoader) LoadJson(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var config Config
	if err := json.Unmarshal(data, &config); err != nil {
		return nil, err
	}
	return &config, nil
}

func (c *ConfigLoader) Load(path string, configType ConfigType) (*Config, error) {
	switch configType {
	case ConfigTypeYaml:
		return c.LoadYaml(path)
	case ConfigTypeJson:
		return c.LoadJson(path)
	}

	return nil, errors.New("config type not supported")
}

func (c *ConfigLoader) Schema(save bool, indentation int, output string) error {
	r := new(jsonschema.Reflector)
	r.RequiredFromJSONSchemaTags = true
	schema := r.Reflect(&Config{})
	schema.Version = "https://json-schema.org/draft-07/schema"
	schema.ID = ""
	schema.Ref = ""

	jsonSchemaBytes, err := schema.MarshalJSON()

	if err != nil {
		log.Panic(err)
	}

	// Unmarshal and re-marshal with indentation (4 tab spaces)
	var jsonObj any
	if err := json.Unmarshal(jsonSchemaBytes, &jsonObj); err != nil {
		log.Panic(err)
	}
	jsonSchemaBytes, err = json.MarshalIndent(jsonObj, "", strings.Repeat(" ", indentation))
	if err != nil {
		log.Panic(err)
	}

	if save {
		err := os.WriteFile(output, jsonSchemaBytes, 0644)
		if err != nil {
			return err
		}
		return err
	}

	jsonSchemaString := string(jsonSchemaBytes)

	fmt.Println(jsonSchemaString)

	return nil
}

func NewConfigLoader() IConfigLoader {
	return &ConfigLoader{}
}
