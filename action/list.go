package action

import (
	"fmt"
	"log"

	"github.com/LQBing/ImageOperator/runconfig"
)

func ListImages(config runconfig.Config, registry_group_type string) {
	var _count int = 0
	for i := range config.Images {
		_img, ok := config.Images[i][registry_group_type]
		if ok {
			_count += 1
		} else {
			continue
		}
		if _img.ImageUrl != "" {
			fmt.Println("    - " + _img.ImageUrl)
		}
	}
	if _count == 0 {
		log.Fatalln("There is no image with registry group type \"" + registry_group_type + "\"")
	}
}
