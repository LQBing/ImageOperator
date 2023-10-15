package action

import (
	"log"
	"strings"

	"github.com/LQBing/ImageOperator/runconfig"
	"github.com/LQBing/ImageOperator/utils"
)

func LoadImages(config runconfig.Config, registry_group_type string, package_path string) {
	var _count int = 0
	image_list := []string{}
	for i := range config.Images {
		_img, ok := config.Images[i][registry_group_type]
		if ok {
			_count += 1
		} else {
			continue
		}
		if _img.ImageUrl != "" {
			image_list = append(image_list, _img.ImageUrl)
		}
	}
	if _count == 0 {
		log.Fatalln("There is no image with registry group type \"" + registry_group_type + "\"")
	}
	cmd_str := "load -i " + package_path + " " + strings.Join(image_list, " ")
	utils.DoContainerCtlCommand(strings.Split(cmd_str, " ")...)
}
