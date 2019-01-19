package main

import (
	"bufio"
	"flag"
	"fmt"
	"math/big"
	"os"
	"time"
)

// TimeEntry
type TimeEntry struct {
	Project  string    `json:"project"`
	Task     string    `json:"task"`
	Duration int       `json:"duration"`
	Start    time.Time `json:"start"`
	End      time.Time `json:"end"`
}

func NewTimer(project string, task string) *time.Ticker {

	fmt.Println(project)

	entry := TimeEntry{
		Task:    task,
		Project: project,
	}

	startTime := time.Now()
	fmt.Printf("Start:  \t %s \n", startTime.Format(time.ANSIC))
	fmt.Printf("Project:\t %s \n", project)
	fmt.Printf("Task:   \t %s \n", task)
	fmt.Println()
	fmt.Printf("\rElapsed Time:\t 00:00:00")

	ticker := time.NewTicker(1 * time.Second)

	if entry.Project == "" {

	}
	//timeEntryChannel := make(chan TimeEntry)

	go func() {
		for range ticker.C {

			h := big.NewInt(int64(time.Since(startTime).Hours()))
			m := big.NewInt(int64(time.Since(startTime).Minutes()))
			s := big.NewInt(int64(time.Since(startTime).Seconds()))
			s = s.Mod(s, big.NewInt(60))
			m = m.Mod(m, big.NewInt(60))
			fmt.Printf("\rElapsed Time:\t %02d:%02d:%02d", h, m, s)
		}
		fmt.Println("askldhjasjkhdjsakhdjkashdjskahdjksa")
	}()

	//output, err := json.Marshal(entry)
	//if err != nil {
	//	fmt.Print("could not marshal json")
	//}
	//fmt.Print(string(output))
	return ticker

}

func main() {

	listCommand := flag.NewFlagSet("list", flag.ExitOnError)
	timeCommand := flag.NewFlagSet("time", flag.ExitOnError)

	// time command flags
	timeProjectPtr := timeCommand.String("project", "default", "Select which project should be used for time tracking")
	timeTaskPtr := timeCommand.String("task", "default", "Select which task should be used for time tracking")

	if len(os.Args) < 2 {
		listCommand.PrintDefaults()
		timeCommand.PrintDefaults()
		os.Exit(1)
	}

	switch os.Args[1] {
	case "list":
		listCommand.Parse(os.Args[2:])
	case "time":
		timeCommand.Parse(os.Args[2:])
	default:
		flag.PrintDefaults()
		os.Exit(1)

	}

	if timeCommand.Parsed() {

		if *timeProjectPtr == "" {
			timeCommand.PrintDefaults()
			os.Exit(1)
		}

		if *timeTaskPtr == "" {
			timeCommand.PrintDefaults()
			os.Exit(1)
		}

		t := NewTimer(*timeProjectPtr, *timeTaskPtr)

		reader := bufio.NewReader(os.Stdin)
		var input string
		for input != "exit\n" {
			input, _ = reader.ReadString('\n')
			if input == "stop\n" {
				hello  <- &t.C;
				fmt.Println("sdkaslædasælkdlæas")
				t.Stop()
			}

		}

	}

}
