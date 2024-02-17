package main

import (
	"fmt"

	"github.com/GoSeoTaxi/olegMicroserviceChatServer/cmd"
	"github.com/fatih/color"
)

func main() {
	fmt.Println(color.HiRedString("Starting Chat Server ..."))
	cmd.RunService()
}
