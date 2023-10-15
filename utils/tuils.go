package utils

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

var DRYRUN bool = false
var CONTAINER_CTL string = "docker"
var TOOL_NAME string = "imageoperator"
var HELP_INFO []string

func checkCommandExist(commandName string) bool {
	_, err := exec.LookPath(commandName)
	if err != nil {
		return false
	}
	return true

}
func GetContainerCtl() {
	if checkCommandExist("docker") {
		CONTAINER_CTL = "docker"
		return
	}
	if checkCommandExist("nerdctl") {
		CONTAINER_CTL = "nerdctl"
		return
	}
	if !DRYRUN {
		if CONTAINER_CTL == "" {
			log.Fatalln("There is no container ctl, please install docker or nerdctl")
		}
	}
}

func GenHelpInfo(args ...[]string) {
	var rgl_str string
	if len(args) > 0 && len(args[0]) > 0 {
		rgl_str = strings.Join(args[0], "|")
	} else {
		rgl_str = "registry_group_types"
	}
	HELP_INFO = []string{}
	HELP_INFO = append(HELP_INFO, "Usage:")
	HELP_INFO = append(HELP_INFO, "    "+TOOL_NAME+" config_help")
	HELP_INFO = append(HELP_INFO, "    "+TOOL_NAME+" [--dry-run] [-c <config_file_path>] pull <"+rgl_str+">")
	HELP_INFO = append(HELP_INFO, "    "+TOOL_NAME+" [--dry-run] [-c <config_file_path>] push <"+rgl_str+">")
	HELP_INFO = append(HELP_INFO, "    "+TOOL_NAME+" [--dry-run] [-c <config_file_path>] rmi <"+rgl_str+">")
	HELP_INFO = append(HELP_INFO, "    "+TOOL_NAME+" [--dry-run] [-c <config_file_path>] list <"+rgl_str+">")
	HELP_INFO = append(HELP_INFO, "    "+TOOL_NAME+" [--dry-run] [-c <config_file_path>] tag <"+rgl_str+"> <"+rgl_str+">")
	HELP_INFO = append(HELP_INFO, "    "+TOOL_NAME+" [--dry-run] [-c <config_file_path>] [-o <package_path>] save <"+rgl_str+">")
	HELP_INFO = append(HELP_INFO, "    "+TOOL_NAME+" [--dry-run] [-c <config_file_path>] [-i <package_path>] load <"+rgl_str+">")
}

var example_json = `{
	"registry_groups": {
		"origin": {
			"registry": "your.registry.com",
			"group": "yourgroup"
		},
		"private": {
			"registry": "dockerhub.kubekey.local",
			"group": "yourgroup"
		}
	},
	"images": [{
			"origin": {
				"image": "tools",
				"tag": "alpine"
			},
			"private": {
				"image": "tools",
				"tag": "alpine"
			}
		},
		{
			"origin": {
				"image": "docker.io/library/redis",
				"tag": "alpine3.18"
			},
			"private": {
				"image": "mirrors",
				"tag": "library_redis_alpine3.18"
			}
		},
		{
			"origin": {
				"image": "docker.io/library/mysql",
				"tag": "8.0"
			},
			"private": {
				"image": "library/mysql",
				"tag": "8.0"
			}
		}
	]
}`
var base_json = `{
	"registry_groups": {
		"origin": {
			"registry": "your.registry.com",
			"group": "yourgroup"
		}
	},
	"images": [{
		"origin": {
			"image": "docker.io/library/redis",
			"tag": "alpine3.18"
		},
		"private": {
			"image": "mirrors",
			"tag": "library_redis_alpine3.18"
		}
	}]
}`

func EchoExampleJson() {
	fmt.Println("Json config instruction:")
	fmt.Println("    Default config file is \"image.json\"")
	fmt.Println("    `.registry_groups[<registry_group_type>]` for storage registry_group registry and group")
	fmt.Println("    `.images[i][<registry_groups>]` for storage iamges in registry_group")
	fmt.Println("    `.images[i][<registry_groups>][\"image\"]` If you mean `nginx:latest`, do not fill it with `nginx`, please fill it with `docker.io/library/mysql`. Or this image will be combine with `.registry_groups[<registry_group_type>][\"registry\"]` and `.registry_groups[<registry_group_type>][\"group\"]`")
	fmt.Println("Base structure:")
	fmt.Println(base_json)
	fmt.Println("Example image.json:")
	fmt.Println(example_json)
	os.Exit(0)
}
func EchoHelp() {
	for i := range HELP_INFO {
		fmt.Println(HELP_INFO[i])
	}
	os.Exit(0)
}
func DoContainerCtlCommand(cmd_str string) {
	cmd := exec.Command(CONTAINER_CTL, cmd_str)
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(out)
}
