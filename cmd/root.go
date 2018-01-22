package cmd

import (
	"os"

	"github.com/astronomerio/commander/api"
	"github.com/astronomerio/commander/api/v1"
	"github.com/astronomerio/commander/kubernetes"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	log = logrus.WithField("package", "cmd")
)

// RootCmd is the commander root command.
var RootCmd = &cobra.Command{
	Use: "commander",
	Run: func(cmd *cobra.Command, args []string) {
		logrus.SetOutput(os.Stdout)
		logrus.SetLevel(logrus.DebugLevel)
		start()
	},
}

func start() {
	client := api.NewClient()
	initAirflowRouteHandler(client)
	client.Serve("8881")
}

func initAirflowRouteHandler(c *api.Client) {
	logger := log.WithField("function", "initAirflowRouteHandler")
	logger.Debug("Entered initAirflowRouteHandler")

	kubernetesProvisioner, err := kubernetes.NewKubeProvisioner()
	if err != nil {
		logger.Panic(err)
	}

	// Alternate provisioners can be swapped here
	airflowRouteHandler := v1.NewAirflowRouteHandler(kubernetesProvisioner)
	c.AppendRouteHandler(airflowRouteHandler)
}
