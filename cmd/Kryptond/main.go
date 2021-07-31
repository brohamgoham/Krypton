package main

import (
	"os"

	"github.com/brohamgoham/Krypton/app"
	"github.com/brohamgoham/Krypton/cmd/Kryptond/cmd"
	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"
)

func main() {
	rootCmd, _ := cmd.NewRootCmd()
	if err := svrcmd.Execute(rootCmd, app.DefaultNodeHome); err != nil {
		os.Exit(1)
	}
}
