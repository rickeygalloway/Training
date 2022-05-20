package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	if err := killServer("server.pid"); err != nil {
		log.Fatal(err)
	}
}

func killServer(pidFile string) error {
	file, err := os.Open(pidFile)

	if err != nil {
		return err
	}
	//defer ==> closed when exit kill server
	defer file.Close()

	var pid int
	if _, err := fmt.Fscanf(file, "%d", &pid); err != nil {
		return fmt.Errorf("%#v: bad PID (%w)", pidFile, err)
	}

	fmt.Printf("Killing %d\n", pid) //simulate kill

	//what to do if os.Remove fails?
	//os.Remove(pidFile)

	if err := os.Remove(pidFile); err != nil {
		log.Println("Warning....")
	}

	return nil
}
