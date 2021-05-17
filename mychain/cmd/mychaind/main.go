package main

import (
	"os"

	"github.com/alirangwala/mychain/cmd/mychaind/cmd"
    svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"
	"github.com/alirangwala/mychain/app"
)

func main() {
	rootCmd, _ := cmd.NewRootCmd()
    if err := svrcmd.Execute(rootCmd, app.DefaultNodeHome); err != nil {
		os.Exit(1)
	}
}
