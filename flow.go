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
	Handler func(data *map[string]interface{}) error
}

var RegisteredFlows []Flow

func CreateNewFlow() {
	RegisteredFlows = append(RegisteredFlows, Flow{
		Name:       "TestFlow" + fmt.Sprintf("%d", len(RegisteredFlows)),
		Trigger:    "insert",
		Collection: "users",
		Steps: []Step{
			{
				Name: "Step 1",
				Handler: func(data *map[string]interface{}) error {
					(*data)["step1"] = "step 1"
					return nil
				},
			},
			{
				Name: "Step 2",
				Handler: func(data *map[string]interface{}) error {
					(*data)["step2"] = "step2"
					return nil
				},
			},
			{
				Name: "Step 3",
				Handler: func(data *map[string]interface{}) error {
					(*data)["step3"] = "step3"
					return nil
				},
			},
		},
	})

	RegisteredFlows = append(RegisteredFlows, Flow{
		Name:       "TestFlow" + fmt.Sprintf("%d", len(RegisteredFlows)),
		Trigger:    "delete",
		Collection: "users",
		Steps: []Step{
			{
				Name: "Step 1",
				Handler: func(data *map[string]interface{}) error {
					return nil
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
				Handler: func(data *map[string]interface{}) error {
					(*data)["prod"] = "hello"
					return nil
				},
			},
		},
	})

	RegisteredFlows = append(RegisteredFlows, Flow{
		Name:       "TestFlow" + fmt.Sprintf("%d", len(RegisteredFlows)),
		Trigger:    "delete",
		Collection: "products",
		Steps: []Step{
			{
				Name: "TestStep_0",
				Handler: func(data *map[string]interface{}) error {
					return nil
				},
			},
		},
	})
}

func RunFlow(collection string, operationType interface{}, data map[string]interface{}) {
	for _, flow := range RegisteredFlows {
		if flow.Trigger == operationType && flow.Collection == collection {
			if len(flow.Steps) == 0 {
				break
			}

			for _, step := range flow.Steps {
				err := step.Handler(&data)
				if err != nil {
					fmt.Println(err)
					break
				}
			}

			fmt.Print("\n===================== Final ", collection, " Data =====================\n")
			for k, v := range data {
				fmt.Printf("%s: %v\n", k, v)
			}
			fmt.Print("======================================================\n\n")
		}
	}
}
