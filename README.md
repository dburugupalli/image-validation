Objective of the Project: 

```
If the user wants to validate a container image memory constraints. 
```


Container Image Validation - Steps - 
```
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
