package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

func Command(cmd string) (output *exec.Cmd) {
	output = exec.Command("bash", "-c", cmd)
	return output
}

func Run(file string) {
	out, _ := Command(fmt.Sprintf(`go run %s`, file)).Output()
	fmt.Println(out)
}

func ExecuteCommands(commands string) {
	filename := "test-123.go"
	ioutil.WriteFile(filename, []byte(commands), 0x777)
	Run(filename)
}

func main() {
	commands := `package main;import "fmt"`

	var cmd string

	fmt.Println("Go in the Shell (type `exit` to finish)")
	reader := bufio.NewReader(os.Stdin)

	for cmd != "exit" {
		if cmd != "" {
			ExecuteCommands(commands)
		}
		fmt.Print(">>> ")
		cmd, _ = reader.ReadString('\n')
		cmd = strings.TrimSpace(cmd)
		commands += "\n" + cmd
	}
}
