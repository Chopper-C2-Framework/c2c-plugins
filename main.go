package main

import (
	"fmt"

	"github.com/chopper-c2-framework/c2-chopper/core/plugins"
)

type EvilPlugin plugins.Plugin

func  New() plugins.Plugin {
	return plugins.Plugin{
		Name: "EvilPlugin",
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
			Options: map[string]string{
				"target": "string",
			},
			ReturnType: "bytes",
		},
	}
}

func (p *EvilPlugin) MetaInfo() *plugins.Metadata {
	return &p.Metadata
}

func (p *EvilPlugin) Info() *plugins.PluginInfo {
	return &p.PluginInfo
}

func (p *EvilPlugin) Options() map[string]string {
	return p.PluginInfo.Options
}

func (p *EvilPlugin) SetArgs(args ...interface{}) {
	fmt.Println("Setting args")

}

func (p *EvilPlugin) Exploit(args ...interface{}) []byte {
	// Do evil things with the args
	return []byte("EvilPlugin has done its dirty work")
}
