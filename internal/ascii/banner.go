// Package ascii handles terminal banners / logos.
package ascii

import "fmt"

// ShowBanner prints the main ReconX-Go banner.
func ShowBanner() {
	fmt.Println(`
██████╗ ███████╗ ██████╗ ██████╗ ███╗   ██╗██╗  ██╗
██╔══██╗██╔════╝██╔════╝██╔═══██╗████╗  ██║██║ ██╔╝
██████╔╝█████╗  ██║     ██║   ██║██╔██╗ ██║█████╔╝ 
██╔══██╗██╔══╝  ██║     ██║   ██║██║╚██╗██║██╔═██╗ 
██║  ██║███████╗╚██████╗╚██████╔╝██║ ╚████║██║  ██╗
╚═╝  ╚═╝╚══════╝ ╚═════╝ ╚═════╝ ╚═╝  ╚═══╝╚═╝  ╚═╝

ReconX-Go Framework
Developed by fu44c
`)
}
