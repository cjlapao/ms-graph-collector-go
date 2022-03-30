package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/cjlapao/common-go/execution_context"
	"github.com/cjlapao/common-go/helper"
	"github.com/cjlapao/common-go/version"
	"github.com/cjlapao/ms-graph-collector-go/startup"
)

var svc = execution_context.Get()

func main() {
	svc.Services.Version.Major = 0
	svc.Services.Version.Minor = 0
	svc.Services.Version.Build = 0
	svc.Services.Version.Rev = 4
	svc.Services.Version.Name = "MS Graph Collector POC"
	svc.Services.Version.Author = "Carlos Lapao"
	svc.Services.Version.License = "MIT"
	getVersion := helper.GetFlagSwitch("version", false)
	if getVersion {
		format := helper.GetFlagValue("o", "json")
		switch strings.ToLower(format) {
		case "json":
			fmt.Println(svc.Services.Version.PrintVersion(int(version.JSON)))
		case "yaml":
			fmt.Println(svc.Services.Version.PrintVersion(int(version.JSON)))
		default:
			fmt.Println("Please choose a valid format, this can be either json or yaml")
		}
		os.Exit(0)
	}

	svc.Services.Version.PrintAnsiHeader()

	configFile := helper.GetFlagValue("config", "")
	if configFile != "" {
		svc.Services.Logger.Command("Loading configuration from " + configFile)
		svc.Configuration.LoadFromFile(configFile)
	}

	defer func() {
	}()

	startup.Init()
}
