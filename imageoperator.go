package main

import (
	"flag"
	"os"

	"github.com/LQBing/ImageOperator/action"
	"github.com/LQBing/ImageOperator/runconfig"
	"github.com/LQBing/ImageOperator/storedconfig"
	"github.com/LQBing/ImageOperator/utils"
)

var sc storedconfig.Config
var rc runconfig.Config

func main() {
	isDryRun := flag.Bool("dry-run", false, "dry run")
	configFilePath := flag.String("c", "image.json", "config file path, default \"image.json\"")
	outputPackagePath := flag.String("o", "image.tar", "output package path, default \"image.tar\"")
	inputPackagePath := flag.String("i", "image.tar", "input package path, default \"image.tar\"")
	utils.GenHelpInfo()
	// parse flags
	flag.Parse()
	if *isDryRun {
		utils.DRYRUN = true
	}
	utils.TOOL_NAME = os.Args[0]
	_args := flag.Args()
	var f_action string
	// load config
	utils.GetContainerCtl()
	sc = storedconfig.Load(*configFilePath)
	utils.GenHelpInfo(storedconfig.GetRegistryGroupListInImages(sc))
	// process config
	rc = runconfig.Load(sc)
	// parse flags
	if len(_args) > 0 {
		f_action = flag.Args()[0]
	} else {
		utils.EchoHelp()
	}

	// do actions
	switch f_action {
	case "config_help":
		utils.EchoExampleJson()
	case "pull":
		checkArg(_args, 2)
		action.PullImages(rc, _args[1])
	case "push":
		checkArg(_args, 2)
		action.PushImages(rc, _args[1])
	case "rmi":
		checkArg(_args, 2)
		action.RmiImages(rc, _args[1])
	case "list":
		checkArg(_args, 2)
		action.ListImages(rc, _args[1])
	case "save":
		checkArg(_args, 2)
		action.SaveImages(rc, _args[1], *outputPackagePath)
	case "load":
		checkArg(_args, 2)
		action.SaveImages(rc, _args[1], *inputPackagePath)
	case "tag":
		checkArg(_args, 3)
		action.TagImages(rc, _args[1], _args[2])
	default:
		utils.EchoHelp()
	}
}
func checkArg(args []string, moreThen int) {
	if len(args) < moreThen {
		utils.EchoHelp()
	}
}
