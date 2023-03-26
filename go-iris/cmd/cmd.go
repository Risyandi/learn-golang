package cmd

import (
	"go-iris/api"

	iris "github.com/kataras/iris/v12"
	"github.com/spf13/cobra"
)

const defaultConfigFilename = "server.yaml"

var serverConfig api.Configuration

// new return a new CLI app
// build with:
// $ go build -ldflags-"-s -w"

func New() *cobra.Command {
	configFile := defaultConfigFilename

	rootCmd := &cobra.Command{
		Use:                        "project",
		Short:                      "command line interface for project",
		Long:                       "the root command will start the http server",
		Version:                    "v0.0.1",
		SilenceError:               true,
		SilenceUsage:               true,
		TraverseChildren:           true,
		SuggestionsMinimumDistance: 1,
		PersistencePreRunE: func(cmd *cobra.Command, args []string) error {
			// read config from file before any of the commands run function
			return serverConfig.BindFile(configFile)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return startServer()
		},
	}

	helpTemplate := HelpTemplate{
		BuildRevision:        iris.BuildRevision,
		BuildTime:            iris.BuildTime,
		ShowGoRuntimeVersion: true,
	}

	rootCmd.setHelpTemplate(helpTemplate.String())

	// shared flags
	flags := rootCmd.PersistentFlags()
	flags.StringVar(&configFile, "config", configFile, "--config=server.yaml a file path which contains the YAML config format")

	// subcommand here.
	// rootCmd.AddCommand(...)

	return rootCmd
}

func startServer() error {
	srv := api.NewServer(serverConfig)
	return srv.Start()
}
