Objective of the Project: 

```
If the user wants to validate a container image memory constraints. 
```

Updated with a Network Command Line CLI tool 

```
NAME:
   Website Lookup CLI - Let's you query IPs, CNAMEs, MX records and Name Servers!

USAGE:
   cli [global options] command [command options] [arguments...]

COMMANDS:
     ns       Looks Up the NameServers for a Particular Host
     cname    Looks up the CNAME for a particular host
     ip       Looks up the IP addresses for a particular host
     mx       Looks up the MX records for a particular host
     help, h  Shows a list of commands or help for one command
```


Container Image Validation -
```
Steps - 
1. Create a channel to catch for interrupts. 
2. Before the tests begin we will print the memory usage, 
3. start the timer
4. create an array and populate the array that it takes up memory
5. print the remaning memory usage. 
```


Commands : 
```
allocate 4Mb of data ./main 4

allocate 8Mb of data ./main 8
```

Make Targets
```
go-compile (compiles go program, although I already compiled Linux binary as convenience)
docker-run (builds golang binary, docker image, runs)
docker-run-ok (puts 6Mb limit on docker container, allocates 4, should run OK)
docker-run-bigmem (puts 6Mb limit on docker container, allocates 12, should FAIL)
```
