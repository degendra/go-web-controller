package main
import (
	"fmt"
	"os/exec"
	"strings"
	"sync"
	"time"
)

func main()  {
	for {
		timer1 := time.NewTimer(time.Minute * 5)
		<-timer1.C
		wg := new(sync.WaitGroup)
		wg.Add(1)
		exe_cmd("killall -v go-web", wg)
		fmt.Println("go-web is dead. It will be up after 10 minutes.")

		timer1 = time.NewTimer(time.Minute * 10)
		<-timer1.C
		wg.Add(1)
		exe_cmd("./go-web", wg)
		fmt.Println("go-web is up. It will be down after 5 minutes.")
	}
}



func exe_cmd(cmd string, wg *sync.WaitGroup) {
	fmt.Println("command is ",cmd)
	// splitting head => g++ parts => rest of the command
	parts := strings.Fields(cmd)
	head := parts[0]
	parts = parts[1:len(parts)]

	out, err := exec.Command(head,parts...).Output()
	if err != nil {
		fmt.Printf("%s", err)
	}
	fmt.Printf("%s", out)
	wg.Done() // Need to signal to waitgroup that this goroutine is done
}