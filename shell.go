package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

type Func struct {
	Name string
	Body []string
}

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

func buildFile(packageName string, imports []string, funcs []Func) string {
	var output string = packageName + "\n"
	output += strings.Join(imports, "\n") + "\n"
	for f := range funcs {
		output += funcs[f].Name
		output += strings.Join(funcs[f].Body, "\n")
		output += "\n}"
	}
	return output
}

func main() {
	packageName := `package main`
	imports := []string{`import "fmt"`, `import "io/ioutil"`}
	funcs := []Func{Func{`func main() {`, []string{`ioutil.Discard.Write([]byte(fmt.Sprint("")))`}}}
	// commands := []string{}

	var cmd string

	fmt.Println("Go in the Shell (type `exit` to finish)")
	reader := bufio.NewReader(os.Stdin)

	for cmd != "exit" {
		if cmd != "" {
			if strings.Contains(cmd, ":=") {
				varName := strings.Split(cmd, " ")[0]
				cmd += fmt.Sprintf("\nioutil.Discard.Write([]byte(%s))", varName)
			}
			funcs[0].Body = append(funcs[0].Body, cmd)

			fileText := buildFile(packageName, imports, funcs)
			status := ExecuteCommands(fileText)
			if status != 0 {
				funcs[0].Body = funcs[0].Body[:len(funcs[0].Body)-1]
			}
		}
		fmt.Print(">>> ")
		cmd, _ = reader.ReadString('\n')
		cmd = strings.TrimSpace(cmd)
	}
}
