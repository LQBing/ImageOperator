package action

import (
	"log"
	"strings"

	"github.com/LQBing/ImageOperator/runconfig"
	"github.com/LQBing/ImageOperator/utils"
)

func PullImages(config runconfig.Config, registry_group_type string) {
	var _count int = 0
	for i := range config.Images {
		_img, ok := config.Images[i][registry_group_type]
		if ok {
			_count += 1
		} else {
			continue
		}
		if _img.ImageUrl != "" {
			cmd_str := "pull " + _img.ImageUrl
			utils.DoContainerCtlCommand(strings.Split(cmd_str, " ")...)
		}
	}
	if _count == 0 {
		log.Fatalln("There is no image with registry group type \"" + registry_group_type + "\"")
	}
}
