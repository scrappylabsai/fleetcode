package main

import (
	"github.com/scrappylabsai/fleetcode/cmd"
	"github.com/scrappylabsai/fleetcode/internal/logging"
)

func main() {
	defer logging.RecoverPanic("main", func() {
		logging.ErrorPersist("Application terminated due to unhandled panic")
	})

	cmd.Execute()
}
