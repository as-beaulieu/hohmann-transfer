package user

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Input() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Make a selection and press enter --> ")
	text, _ := reader.ReadString('\n')
	// remove the delimeter from the string
	input := strings.TrimSuffix(text, "\n")
	fmt.Println(input)
	return input
}

func Pause() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Press enter to continue-> ")
	reader.ReadString('\n')
}
