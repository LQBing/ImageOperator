package action

import (
	"fmt"
	"log"

	"github.com/LQBing/ImageOperator/runconfig"
	"github.com/LQBing/ImageOperator/utils"
)

func PushImages(config runconfig.Config, registry_group_type string) {
	var _count int = 0
	for i := range config.Images {
		_img, ok := config.Images[i][registry_group_type]
		if ok {
			_count += 1
		} else {
			continue
		}
		if _img.ImageUrl != "" {
			cmd_str := " push " + _img.ImageUrl
			if utils.DRYRUN {
				fmt.Println(utils.CONTAINER_CTL + cmd_str)
			} else {
				utils.DoContainerCtlCommand(cmd_str)
			}
		}
	}
	if _count == 0 {
		log.Fatalln("There is no image with registry group type \"" + registry_group_type + "\"")
	}
}
