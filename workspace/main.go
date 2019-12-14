package main

import (
	"sniper/cmd/job"
	"sniper/cmd/server"

	"github.com/spf13/cobra"
)

func main() {
	root := cobra.Command{Use: "sniper"}

	root.AddCommand(
		server.Cmd,
		job.Cmd,
	)

	root.Execute()
}
