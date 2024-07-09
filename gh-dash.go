package main

import (
   "os"
   "log"
   "runtime/pprof"
	"github.com/dlvhdr/gh-dash/v4/cmd"
)

func main() {
   f, err := os.Create("cpu.profile")
   if err != nil {
      log.Fatal("could not create CPU profile: ", err)
   }
   defer f.Close()
   if err := pprof.StartCPUProfile(f); err != nil {
      log.Fatal("could not start CPU profile: ", err)
   }
   defer pprof.StopCPUProfile()

	cmd.Execute()
}
