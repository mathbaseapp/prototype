package converter

import (
	"encoding/xml"
	"fmt"
	"os/exec"
)

// Run app
func Run() {
	latexStr := "$$x = 4$$"
	pandocCmd := "echo '" + latexStr + "'  | pandoc -f html+tex_math_dollars -t html --mathml"
	out, err := exec.Command("sh", "-c", pandocCmd).Output()
	str := Node{}
	fmt.Println(string(out))
	xml.Unmarshal(out, &str)

	if err != nil {
		fmt.Println("Command Exec Error.")
	}
	fmt.Println(str.Nodes[0].Nodes[0].Nodes[0].Value)
}
