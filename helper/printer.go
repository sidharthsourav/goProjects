package helper

import "fmt"

func printHelper(message string) string {
	fmt.Println("This message is printed by the helper function.")
	return message + "-helperWorkedHere!!!"
}
