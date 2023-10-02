package main

import "fmt"

type Flow struct {
	Name       string
	Trigger    string
	Collection string
	Steps      []Step
}

type Step struct {
	Name    string
	Handler func(data interface{})
}

var RegisteredFlows []Flow

func CreateNewFlow() {
	RegisteredFlows = append(RegisteredFlows, Flow{
		Name:       "TestFlow" + fmt.Sprintf("%d", len(RegisteredFlows)),
		Trigger:    "insert",
		Collection: "users",
		Steps: []Step{
			{
				Name: "TestStep_0",
				Handler: func(data interface{}) {
					fmt.Println("TestStep_0")
					fmt.Println("insert on users :: ", data)
				},
			},
		},
	})

	RegisteredFlows = append(RegisteredFlows, Flow{
		Name:       "TestFlow" + fmt.Sprintf("%d", len(RegisteredFlows)),
		Trigger:    "insert",
		Collection: "products",
		Steps: []Step{
			{
				Name: "TestStep_0",
				Handler: func(data interface{}) {
					fmt.Println("TestStep_0")
					fmt.Println("insert on products :: ", data)
				},
			},
		},
	})
}

func RunFlow(collection string, operationType interface{}, data interface{}) {
	for _, flow := range RegisteredFlows {
		if flow.Trigger == operationType && flow.Collection == collection {
			for _, step := range flow.Steps {
				step.Handler(data)
			}
		}
	}
}
