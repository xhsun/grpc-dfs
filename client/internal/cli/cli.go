package cli

import (
	"os"

	"github.com/urfave/cli/v2"
	"github.com/xhsun/grpc-file-transfer/client/internal/config"
	"github.com/xhsun/grpc-file-transfer/client/internal/file"
)

// CLI the CLI starting point
type CLI struct {
	app *cli.App
}

// NewCLI method creates a new CLI
func NewCLI(config *config.Config, fileCommand *file.FileCommand) *CLI {
	return &CLI{
		app: &cli.App{
			Name:                 config.ProjectName,
			Usage:                "Distributed File System CLI",
			EnableBashCompletion: true,
			Commands: []*cli.Command{
				{
					Name:    "upload",
					Usage:   "Upload `LOCAL_FILE` to server",
					Aliases: []string{"u"},
					Action:  fileCommand.Upload,
				},
				{
					Name:    "copy",
					Aliases: []string{"cp"},
					Usage:   "Copy `REMOTE_FILE` from server to `LOCAL_FILE`",
					Action:  fileCommand.Download,
				},
				{
					Name:    "remove",
					Aliases: []string{"rm"},
					Usage:   "Remove `REMOTE_FILE` from server",
					Action:  fileCommand.Remove,
				},
				{
					Name:    "list",
					Aliases: []string{"ls"},
					Usage:   "Lists all files in server",
					Action:  fileCommand.ListFiles,
				},
			},
		},
	}

}

// Start method to start the CLI
func (c *CLI) Start() {
	c.app.Run(os.Args)
}
