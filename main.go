package main

import (
	"fmt"
	"github.com/fu44c/reconx/cmd"
)

func main() {
	banner := `
	
	____                           
	|  _ \ ___  ___ ___  _ __ __  __
	| |_) / _ \/ __/ _ \| '_ \\ \/ /
	|  _ <  __/ (_| (_) | | | |>  < 
	|_| \_\___|\___\___/|_| |_/_/\_\
             

			ReconX - Bug Bounty & Recon Tool
                 Developer: Ahmed ibrahim
	`

	fmt.Println(banner)
	cmd.Execute()
}
