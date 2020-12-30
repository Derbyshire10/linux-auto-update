package main

import (
	"log"
	"os/exec"
	"runtime"
)

func main() {
	os := runtime.GOOS
	switch os {
	case "linux":
		cmd := exec.Command("/bin/sh", "-c", "sudo apt-get update; sudo apt-get upgrade -y;")
		log.Printf("Running command and waiting for it to finish...")
		err := cmd.Run()
		if err != nil {
			log.Fatal(err)
			log.Printf("Command finished with error: %v", err)
		} else {
			err = cmd.Wait()
			log.Printf("Command ran with no errors")
		}
	}
}
