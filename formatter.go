package main

import (
	"encoding/json"
	"fmt"
)

// remap struct to json output
func jsonify(processes []FedoraProcess) []byte {

	// error if no procfs
	if len(processes) <= 1 {
		errResponse := JsonError{"There was an error getting process data"}
		erroJson,_ := json.Marshal(errResponse)
		return erroJson
	}

	resultsmap := make(map[int]interface{})
	for _, process := range processes {
		pidStruct1 := Procinfo{}
		pidStruct2 := ProcinfoBool{}
		if len(process.cmdline) == 0 {
			pidStruct2 = ProcinfoBool{false, process.ppid,
				process.name, process.environment}
			resultsmap[process.pid] = pidStruct2
		} else {
			pidStruct1 = Procinfo{process.cmdline, process.ppid,
				process.name, process.environment}
			resultsmap[process.pid] = pidStruct1
		}
	}

	procj,e :=json.MarshalIndent(resultsmap, "", "    ")
	if e != nil {
		fmt.Println("Could not marshal json")
	}
	return procj
}

// struct for process information
type Procinfo struct {
	Cmdline	string
	Ppid	int
	Name	string
	Environment	[]string
}

// struct for process information. Boolean
// version of cmdline
type ProcinfoBool struct {
	Cmdline	bool
	Ppid	int
	Name	string
	Environment	[]string
}

//Simple error response to client
type JsonError struct {
	Err		string
}
