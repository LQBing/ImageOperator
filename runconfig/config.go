package runconfig

import (
	"fmt"
	"log"
	"strings"

	"github.com/LQBing/ImageOperator/storedconfig"
)

type RegistryGroup struct {
	Registry string
	Group    string
}
type Image struct {
	ImageUrl string
}

type Config struct {
	RegistryGroups map[string]RegistryGroup
	Images         []map[string]Image
}

func Load(sc storedconfig.Config) Config {
	var config Config
	// trans registry group config
	registryGroups := map[string]RegistryGroup{}
	for rg := range sc.RegistryGroups {
		registryGroups[rg] = RegistryGroup{sc.RegistryGroups[rg].Registry, sc.RegistryGroups[rg].Group}
	}
	config.RegistryGroups = registryGroups
	// trans images config
	images := []map[string]Image{}
	for _image_key := range sc.Images {
		_img := map[string]Image{}
		for _rg := range sc.Images[_image_key] {
			var _image_url string
			var _tag string
			if sc.Images[_image_key][_rg].Tag == "" {
				_tag = "latest"

			} else {
				_tag = sc.Images[_image_key][_rg].Tag
			}
			var _image_name = sc.Images[_image_key][_rg].Image + ":" + _tag
			var _t = strings.Split(sc.Images[_image_key][_rg].Image, "/")
			if len(_t) == 1 {
				_image_url = strings.Join([]string{config.RegistryGroups[_rg].Registry, config.RegistryGroups[_rg].Group, _image_name}, "/")
			} else if len(_t) == 2 {
				_image_url = strings.Join([]string{config.RegistryGroups[_rg].Registry, _image_name}, "/")
			} else if len(_t) == 3 {
				_image_url = _image_name
			} else {
				fmt.Println(_t)
				fmt.Println(len(_t))
				log.Fatalln(_image_name + " is not a valid image name")
			}
			_img[_rg] = Image{_image_url}
		}
		images = append(images, _img)
	}
	config.Images = images
	return config
}
