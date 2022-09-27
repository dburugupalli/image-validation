Objective of the Project: 
```
If the user wants to validate a container image memory constraints. 
```
Project: 
```
It allocates memory in 1mb blocks
```
Commands : 
```
allocate 4Mb of data ./golang-memtest 4

allocate 8Mb of data ./golang-memtest 8
```

Make Targets
```
go-compile (compiles go program, although I already compiled Linux binary as convenience)
docker-run (builds golang binary, docker image, runs)
docker-run-ok (puts 8Mb limit on docker container, allocates 4, should run OK)
docker-run-bigmem (puts 8Mb limit on docker container, allocates 12, should FAIL)
```
