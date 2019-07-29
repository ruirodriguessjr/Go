/* app.go */

package main

import "fmt"

type app []string

func (a app) print() {
	for i, app := range a {
		fmt.Println(i, app)
	}
}

func main() {
	apps := app{"Facebook", "Instagram", "WhatsApp"}
	apps.print()
}
