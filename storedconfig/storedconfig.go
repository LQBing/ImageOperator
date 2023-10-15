package storedconfig

import (
	"encoding/json"
	"io"
	"log"
	"os"

	"golang.org/x/exp/slices"
)

type RegistryGroup struct {
	Registry string `json:"registry"`
	Group    string `json:"group"`
}
type Image struct {
	Image          string `json:"image"`
	Tag            string `json:"tag"`
	CustomRegistry bool   `json:"custom_registry"`
	CustomGroup    bool   `json:"custom_group"`
}

type Config struct {
	RegistryGroups map[string]RegistryGroup `json:"registry_groups"`
	Images         []map[string]Image       `json:"images"`
}

func Load(configFile string) Config {
	fileContent, err := os.Open(configFile)
	if err != nil {
		log.Fatal(err)
	}
	byteResult, _ := io.ReadAll(fileContent)
	defer fileContent.Close()
	var config Config
	json.Unmarshal([]byte(byteResult), &config)
	return config
}

func GetRegistryGroupListInImages(config Config) []string {
	var rgs []string
	for _ik := range config.Images {
		for _rg := range config.Images[_ik] {
			if !slices.Contains(rgs, _rg) {
				rgs = append(rgs, _rg)
			}
		}
	}
	return rgs
}
