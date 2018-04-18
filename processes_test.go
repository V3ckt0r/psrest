package main

import (
	"testing"
	"fmt"
	"os"
	"strconv"
)

func TestFailedGetEnvForPid(t *testing.T) {
	ret, _ := getEnvForPid("1")
	fmt.Println("Test1 Success: ", ret, "\n")
}

func TestSuccessGetEnvForPid(t *testing.T) {
	ps := strconv.Itoa(os.Getpid())
	ret, err := getEnvForPid(ps)
	if err != nil {
		t.Errorf("Failed to get env for pid...")
	}
	fmt.Println("Test2 Success: ", ret, "\n")
}

func TestGetCmdForPid(t *testing.T) {
	ret, err := getCmdForPid("1")
	if err != nil {
		t.Errorf("Failed to get cmd for pid...")
	}
	fmt.Println("Test2: ", ret, "\n")
}

func TestGetProcPath(t *testing.T) {
	file, err := getProcPath("1", "stat")
	if err != nil {
		t.Error("Failed to open file")
	}
	fmt.Println("Test3:", file, "\n")
}

func TestGetProcesses(t *testing.T) {
	get_processes, err := processes()
	if err != nil {
		t.Error("Test failed: ",err)
	}
	fmt.Println("Test4: ", get_processes, "\n")
}