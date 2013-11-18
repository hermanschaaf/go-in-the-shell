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

func Run(file string) (exitStatus int) {
	cmd := fmt.Sprintf(`go run %s`, file)
	command := Command(cmd)
	out, err := command.CombinedOutput()

	fmt.Println(string(out))
	if err != nil {
		return 1
	}
	return 0
}

func ExecuteCommands(commands string) (status int) {
	filename := "test-123.go"
	ioutil.WriteFile(filename, []byte(commands), os.ModeTemporary|os.ModePerm)
	return Run(filename)
}

func main() {
	commands := []string{`package main`, `import "fmt"`, `func main() {`}

	var cmd string

	fmt.Println("Go in the Shell (type `exit` to finish)")
	reader := bufio.NewReader(os.Stdin)

	for cmd != "exit" {
		if cmd != "" {
			status := ExecuteCommands(strings.Join(commands, "\n") + "\n" + cmd + "\n}")
			if status == 0 {
				commands = append(commands, cmd)
			}
		}
		fmt.Print(">>> ")
		cmd, _ = reader.ReadString('\n')
		cmd = strings.TrimSpace(cmd)
	}
}
