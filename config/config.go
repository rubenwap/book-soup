package config

import (
	"fmt"
	"github.com/kylelemons/go-gypsy/yaml"
)

func Config() *yaml.File {
	config, err := yaml.ReadFile("config/config.yaml")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(config.Get("db.host"))
	// fmt.Println(config.GetBool("enabled"))
	return config
}
