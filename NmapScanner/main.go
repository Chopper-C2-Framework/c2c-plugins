package main

import (
	"errors"
	"fmt"
	"strings"

	"github.com/Ullaakut/nmap"

	"github.com/google/uuid"

	"github.com/chopper-c2-framework/c2-chopper/core/domain/entity"
	"github.com/chopper-c2-framework/c2-chopper/core/plugins"
	"github.com/chopper-c2-framework/c2-chopper/core/services"
)

type NmapScanner struct {
	plugins.Plugin
	TaskService services.ITaskService
	targetIp    string
	port_range  string
	agentId     uuid.UUID
	waitingTask *entity.TaskModel
}

func New(service services.ITaskService) plugins.IPlugin {
	return &NmapScanner{
		Plugin: plugins.Plugin{
			Metadata: plugins.Metadata{
				Version:     "1.0",
				Author:      "C2-Chopper",
				Tags:        []string{"scanner", "info gathering"},
				ReleaseDate: "2023-05-01",
				Type:        plugins.InfoRetriever,
				SourceLink:  "https://github.com/Chopper-C2-Framework",
				Description: "Scan a given IP address.",
			},
			PluginInfo: plugins.PluginInfo{
				Name: "NmapScanner",
				Options: map[string]string{
					"target":     "string",
					"port_range": "string",
				},
				ReturnType: "string",
			},
		},
		TaskService: service,
	}
}

func (p NmapScanner) MetaInfo() *plugins.Metadata {
	return &p.Metadata
}

func (p NmapScanner) Info() *plugins.PluginInfo {
	return &p.PluginInfo
}

func (p NmapScanner) Options() map[string]string {
	return p.PluginInfo.Options
}

func (p *NmapScanner) SetArgs(args ...interface{}) error {
	fmt.Println("Setting args")

	arg1, ok := args[0].(string)
	if !ok {
		return errors.New("Bad first argument")
	}

	arg2, ok := args[1].(string)
	if !ok {
		p.port_range = "1-1000"
	}

	fmt.Println("[NmapScanner] Setting targetIp to", arg1)
	p.targetIp = fmt.Sprint(arg1)
	fmt.Println("[NmapScanner] Setting port_range to", arg2)
	p.port_range = fmt.Sprint(arg2)
	return nil
}

func (p *NmapScanner) Exploit(Channel chan *entity.TaskResultModel, args ...interface{}) []byte {
	scanner, err := nmap.NewScanner(
		nmap.WithTargets(p.targetIp),
		nmap.WithPorts(p.port_range),
	)
	if err != nil {
		return []byte(err.Error())
	}

	result, warnings, err := scanner.Run()
	output := ""
	if len(warnings) > 0 {
		output = "Warnings:\n" + strings.Join(warnings, "\n") + "\n"
	}

	if err != nil {
		return []byte(output + err.Error())
	}

	// Use the results to print an example output
	for _, host := range result.Hosts {
		if len(host.Ports) == 0 || len(host.Addresses) == 0 {
			continue
		}

		output += fmt.Sprintf("Host %q:\n", host.Addresses[0])

		for _, port := range host.Ports {
			output += fmt.Sprintf("\tPort %d/%s %s %s\n", port.ID, port.Protocol, port.State, port.Service.Name)
		}
	}

	return []byte(output)
}

func (p NmapScanner) IsWaitingForTaskResult() (bool, string) {
	if p.waitingTask == nil {
		return false, ""
	}
	return true, p.waitingTask.ID.String()
}
