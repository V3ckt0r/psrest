# ps-rest

A quick hack to get process information of a Linux machine that implements procfs.

### Build Binary
To get up an running please compile the code via running the below command (while in the source dir):
```
go build
```
The binary file created will have the same name as the directory you were in.
Note: Make sure you have set your $GOPATH

### Run and Stop
To run the binary simply execute it with your preferred user:
```
./main
```
Stopping the program cleanly can be done via SIGINT, ctrl+c

### Get process info
The process information can be restrived by going to `http://localhost/ps`

### Tests
To run project tests please do the below in the project src
```
go test
```
All tests should be successful when running on a Fedora based system

### Organisation
The project files have been split into 3 main areas.
* main.go: The entry point containing the web server and handeler
* processes.go: Runs through the various logical checks to gather and structure ps data
* process.go: Underlying data structure and utility command
* formatter.go: Transforms the ps data into json

### Roadmap
* A logger - To help debugging
* Custom serialiser - For unique json behaviour, such as the cmdline attribute, this project
could benefit from a custom serialiser.
