package main

import (
	"os"
	"io/ioutil"
	"strconv"
	"fmt"
	"io"
	"strings"
	"bytes"
	"path/filepath"
)

/* This function finds process on host and gets the
 various details. Makes extensive use of procfs. stat,
 cmdline and environ files are all used to gather info.
*/
func processes() ([]FedoraProcess, error) {
	d, err := os.Open("/proc")
	if err != nil {
		return nil, err
	}
	defer d.Close()

	results := make([]FedoraProcess, 0, 50)
	for {
		fstat, err := d.Readdir(10)
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		for _, fi := range fstat {
			// The dirs are the PIDs so ignore files
			if !fi.IsDir() {
				continue
			}

			// Of the dirs, need to get those which are numerical.
			name := fi.Name()
			if name[0] < '0' || name[0] > '9' {
				continue
			}


			//If process ID is this program then skip
			procName := strings.Trim(name, "()")
			n ,err := strconv.Atoi(procName)
			if err != nil {
				fmt.Println(n, err)
			}

			if n == os.Getpid() {
				continue
			}

			// From this point forward, any errors we just ignore, because
			// it might simply be that the process doesn't exist anymore.
			pid, err := strconv.ParseInt(name, 10, 0)
			if err != nil {
				continue
			}

			processPathfile, err := getProcPath(name, "stat")
			if err != nil {
				fmt.Println("error occured opening file:", err)
				return nil, err
			}

			//split stat file and capture key data, ppid, name, cmd
			stat := strings.Split(string(processPathfile), " ")

			ppid, err := strconv.Atoi(stat[3])
			if err != nil {
				fmt.Println("An error with getting commands", err)
			}

			cmds, err := getCmdForPid(stat[0])
			if err != nil {
				fmt.Println("An error with getting commands", err)
			}

			env, err := getEnvForPid(name)
			words := strings.Fields(env)

			// create FedoraProcess struct
			p, err := newFedoraProcess(int(pid), strings.Trim(stat[1], "()"), ppid, cmds, words)
			if err != nil {
				continue
			}

			results = append(results, *p)
		}
	}

	return results, nil
}

// Generic function that handles getting files from procfs
func getProcPath(pid string, file string) ([]byte, error) {
	processPath := filepath.Join("/proc", pid, file)
	dat, err := ioutil.ReadFile(processPath)
	if err != nil {
		//fmt.Println("error occured opening file:", err)
		return nil, err
	}
	return dat, nil
}


// This function is used to work out the content of
// cmdline which holds the command line arguments.
// Some sanitisation is required due to NULL chars
func getCmdForPid(pid string) (string, error) {
	processPathfile, err := getProcPath(pid, "cmdline")
	if err != nil {
		return "", err
	}
	ret := bytes.Replace(processPathfile, []byte("\x00"), []byte(""), -1)
	return string(ret), nil
}

//search for environment variables and sanities
func getEnvForPid(pid string) (string, error) {
	processPathfile, err := getProcPath(pid, "environ")
	if err != nil {
		return "", err
	}
	environ := bytes.Replace(processPathfile, []byte("\x00"), []byte(" "), -1)
	return string(environ), nil
}
