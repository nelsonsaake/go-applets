package main

import (
	_ "embed"
	"encoding/json"
	"fmt"
)

var (
	//go:embed config.json
	config    string
	configMap map[string]any
)

func init() {
	if len(config) == 0 {
		return
	}

	err := json.Unmarshal([]byte(config), &configMap)
	if err != nil {
		panic(err)
	}

	initGlobalVars()
}

// global variables

var (
	dir                   = ""
	changes               = map[string]string{}
	changerFuncs          = []func(string) string{}
	mustBeDirectlyRelated = false
	skipFiles             = true
	skipDirs              = false
)

func initGlobalVars() {

	dir, _ = configMap["dir"].(string)
	if len(dir) == 0 {
		panic("no dir provided")
	}

	mustBeDirectlyRelated, _ = configMap["mustBeDirectlyRelated"].(bool)

	skipFiles, _ = configMap["skipFiles"].(bool)

	skipDirs, _ = configMap["skipDirs"].(bool)

	cs, ok := configMap["changes"].(map[string]interface{})
	if !ok {
		panic("changes cast not ok")
	}

	for k, v := range cs {
		changes[k] = fmt.Sprint(v)
	}

	cfs, ok := configMap["changerFuncs"].([]any)
	if !ok {
		fmt.Printf("%T", configMap["changerFuncs"])
		panic("changerFuncs cast not ok")
	}

	fmt.Println(cfs)
	for _, name := range cfs {
		name := name.(string)
		changerFunc, ok := changerFuncLookup[name]
		if ok {
			changerFuncs = append(changerFuncs, changerFunc)
		}
	}
}
