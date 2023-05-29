package main

import (
	"errors"
	"fmt"

	"github.com/google/uuid"

	"github.com/chopper-c2-framework/c2-chopper/core/domain/entity"
	"github.com/chopper-c2-framework/c2-chopper/core/plugins"
	"github.com/chopper-c2-framework/c2-chopper/core/services"
)

type EvilPlugin struct {
	plugins.Plugin
	TaskService services.ITaskService
	targetIp    string
	agentId     uuid.UUID
	waitingTask *entity.TaskModel
}

func New(service services.ITaskService) plugins.IPlugin {
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
					"target":  "string",
					"agentId": "string",
				},
				ReturnType: "string",
			},
		},
		TaskService: service,
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
	// Outdated

	arg1, ok := args[0].(string)
	if !ok {
		return errors.New("Bad first argument")
	}

	arg2, ok := args[1].(string)
	if !ok {
		return errors.New("Bad second argument")
	}
	agentId, err := uuid.Parse(arg2)
	if err != nil {
		return errors.New("Invalid id")
	}
	p.agentId = agentId

	fmt.Println("Setting targetIp to", arg1)
	p.targetIp = fmt.Sprint(arg1)
	return nil
}

func (p *EvilPlugin) Exploit(Channel chan *entity.TaskResultModel, args ...interface{}) []byte {
	// Do evil things with the args
	fmt.Println("Do evil things to", p.targetIp)
	p.waitingTask = &entity.TaskModel{
		Name:    "Evil Plugin Task",
		AgentId: p.agentId,
		Args:    "ls",
		Type:    entity.TASK_TYPE_SHELL,
	}
	p.TaskService.CreateTask(p.waitingTask)

	fmt.Println("Waiting for task id", p.waitingTask.ID.String())
	result := <-Channel

	fmt.Println("Got msg:", result.Output)
	return []byte(fmt.Sprint("EvilPlugin attacking", p.targetIp, "\nOutput:", result.Output))
}

func (p EvilPlugin) IsWaitingForTaskResult() (bool, string) {
	if p.waitingTask == nil {
		return false, ""
	}
	return true, p.waitingTask.ID.String()
}
