package main

import (
	"errors"
	"fmt"

	"github.com/chopper-c2-framework/c2-chopper/core/plugins"
)

type EvilPlugin struct {
	plugins.Plugin
	targetIp string
}

func New() plugins.IPlugin {
	return &EvilPlugin{
		Plugin: plugins.Plugin{
			Metadata: plugins.Metadata{
				Version:     "1.0",
				Author:      "Evil Corp",
				Tags:        []string{"evil", "malware"},
				ReleaseDate: "2023-05-01",
				Type:        1,
				SourceLink:  "https://evilcorp.com/plugins/evil",
				Description: "This plugin does evil things",
			},
			PluginInfo: plugins.PluginInfo{
				Name: "EvilPlugin",
				Options: map[string]string{
					"target": "string",
				},
				ReturnType: "bytes",
			},
		},
		targetIp: "",
	}
}

func (p EvilPlugin) MetaInfo() *plugins.Metadata {
	return &p.Metadata
}

func (p EvilPlugin) Info() *plugins.PluginInfo {
	return &p.PluginInfo
}

func (p EvilPlugin) Options() map[string]string {
	return p.PluginInfo.Options
}

func (p *EvilPlugin) SetArgs(args ...interface{}) error {
	fmt.Println("Setting args")

	arg1, ok := args[0].(string)
	if !ok {
		return errors.New("Bad first argument")
	}

	fmt.Println("Setting targetIp to", arg1)
	p.targetIp = fmt.Sprint(arg1)
	return nil
}

func (p EvilPlugin) Exploit(args ...interface{}) []byte {
	// Do evil things with the args
	fmt.Println("Do evil things to", p.targetIp)
	return []byte(fmt.Sprint("EvilPlugin attacking ", p.targetIp))
}
