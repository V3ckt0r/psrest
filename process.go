package main

// A Process data structure that will hold the relevant
// process information
type FedoraProcess struct {
	pid			int
	cmdline     string
	ppid        int
	name        string
	environment []string
}

// used to create a new FedoraProcess struct
func newFedoraProcess(pid int, name string, ppid int, cmdline string, env []string) (*FedoraProcess, error) {
	p := &FedoraProcess{pid: pid, name: name, ppid: ppid, cmdline: cmdline, environment: env}
	return p, nil
}
