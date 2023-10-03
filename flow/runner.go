package flow

import (
	"log"

	v8 "rogchap.com/v8go"
)

func RunFunc(function string, globalFunctions ...string) interface{} {
	ctx := v8.NewContext()

	var contextFunctions string

	if len(globalFunctions) == 0 || len(globalFunctions) > 1 {
		contextFunctions = ""
	}

	ctx.RunScript(contextFunctions, "main.js")
	ctx.RunScript(function, "main.js")

	data, err := ctx.RunScript("data", "main.js")

	if err != nil {
		log.Println("Error in running data script", err)
		return nil
	}

	log.Println(data)
	return data
}
