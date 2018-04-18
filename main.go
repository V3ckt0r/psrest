package main

import (
	"fmt"
	"net/http"
	"encoding/json"
)

func main() {
	fmt.Println("Starting server...")
	http.Handle("/ps", pshander())
	http.ListenAndServe(":8080", nil)
}

/*
http handler for /ps endpoint
 */
func pshander() http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		procret := processResults()
		rw.Header().Set("Content-Type", "application/json")
		fmt.Fprint(rw, string(procret))
	})
}

/*
wrapper for getting processes
 */
func processResults() []byte {
	results,e := processes()
	if e != nil {
		// error if no procfs
		fmt.Println("There was an error getting process data, no procfs: ", e)
		errResponse := &JsonError{"There was an error getting process data"}
		errJson,_ := json.Marshal(errResponse)
		return errJson

	}
	json_data := jsonify(results)
	return json_data
}
