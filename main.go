package main

import (
 "time"
 "os/signal"
 "syscall"
 "fmt"
 "os"
 "strconv"
 "runtime"
 "math/rand"
 
)

//func Print(a ...interface{}) (n int, err error)

func main() {

    // alert to any OS signals sent while running
    alertAppWhenOSInterrupts()

    // get number of Mb to allocate from param (default=1)
    nmb := ReadEnvOrArgs(1,"nmb","1")
    // get number of milliseconds to wait between 1Mb allocations (default=100)
    nms := ReadEnvOrArgs(2,"nms","100")


    // print the memory usage before the tests begins
    PrintMemoryUsage()
    fmt.Println("Asked to allocate %dMb\n\n",nmb)

    // allocate memory 1Mb at a time
    rand.Seed(time.Now().UTC().UnixNano())
    var resarr = make([][]byte,nmb)
    for i:=0; i<nmb; i++ {
      resarr[i] = make([]byte, 1024*1024)
      // populate array so it takes up memory and  if this is not done, it will not fill up memory space
      rand.Read(resarr[i])
      PrintMemoryUsage()
      time.Sleep( time.Duration(nms) * time.Millisecond)
      //fmt.Printf("Total allocated: %dMb\n",i+1)
    }
    fmt.Println("\n")
    // print the memoryusage after the tests complete
    PrintMemoryUsage()
    fmt.Println("successfully allocated memory %dMb\n",len(resarr))
}


// conversation of byte to megabyte
func bToMb(b uint64) uint64 {
    return b / 1024 / 1024
}


// For info on each, see: https://golang.org/pkg/runtime/#MemStats
func PrintMemoryUsage() {
        var m runtime.MemStats
        runtime.ReadMemStats(&m)
        fmt.Printf("Alloc = %v MiB", bToMb(m.Alloc))
        fmt.Printf("\tTotalAlloc = %v MiB", bToMb(m.TotalAlloc))
        fmt.Printf("\tSys = %v MiB", bToMb(m.Sys))
        fmt.Printf("\tNumGC = %v\n", m.NumGC)
}


// picks 'nmb' value from OS env or cmd defaults to 1 if no values found
// get number of Mb to allocate from param (default=1)
// get number of milliseconds to wait between 1Mb allocations (default=100)
func ReadEnvOrArgs(posIndex int,pname string,defaultString string) int {
    nmbstr := defaultString
    if len(os.Args)>posIndex {
      nmbstr = os.Args[posIndex]
    }else if len(os.Getenv(pname))>0 {
      nmbstr = os.Getenv(pname)
    }
    nmb,err := strconv.Atoi(nmbstr)
    if err != nil {
      fmt.Println(err)
      os.Exit(1)
    }
    return nmb
}

// proof to docker container for OOM is sudden, no there no interrupt from Operating Signals.
func alertAppWhenOSInterrupts() {
    sigc := make(chan os.Signal, 1)
    signal.Notify(sigc,
        syscall.SIGINT,
        syscall.SIGTERM,
        syscall.SIGHUP,
        syscall.SIGQUIT)
    go func() {
        s := <-sigc
        fmt.Printf("There's a signal sent from operating Systems")
        fmt.Println(s)
    }()
}

