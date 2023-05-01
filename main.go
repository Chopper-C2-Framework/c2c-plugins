package main

import "github.com/c2-chopper/core/plugins"

type EvilPlugin struct {
	args []interface{}
}

func (p *EvilPlugin) New() plugins.Plugin {
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
	return &p.New().Metadata
}

func (p *EvilPlugin) Info() *plugins.PluginInfo {
	return &p.New().PluginInfo
}

func (p *EvilPlugin) Options() map[string]string {
	return p.New().PluginInfo.Options
}

func (p *EvilPlugin) SetArgs(args ...interface{}) {
	p.args = args
}

func (p *EvilPlugin) Exploit(args ...interface{}) []byte {
	// Do evil things with the args
	return []byte("EvilPlugin has done its dirty work")
}
