package action

import (
	"log"
	"strings"

	"github.com/LQBing/ImageOperator/runconfig"
	"github.com/LQBing/ImageOperator/utils"
)

func TagImages(config runconfig.Config, source_registry_group_type string, target_registry_group_type string) {
	var _count int = 0
	for i := range config.Images {
		_img_s, ok_s := config.Images[i][source_registry_group_type]
		_img_t, ok_t := config.Images[i][target_registry_group_type]
		if ok_s && ok_t {
			_count += 1
		} else {
			continue
		}
		if _img_s.ImageUrl != "" || _img_t.ImageUrl != "" {
			cmd_str := "tag " + _img_s.ImageUrl + " " + _img_t.ImageUrl
			utils.DoContainerCtlCommand(strings.Split(cmd_str, " ")...)
		}
	}
	if _count == 0 {
		log.Fatalln("There is no image with registry group type both \"" + source_registry_group_type + "\" and \"" + target_registry_group_type + "\"")
	}
}
