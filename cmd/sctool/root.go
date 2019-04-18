// Copyright (C) 2017 ScyllaDB

package main

import (
	"os"

	"github.com/scylladb/mermaid/mermaidclient"
	"github.com/spf13/cobra"
)

var (
	defaultURL = "https://127.0.0.1:56443/api/v1"

	cfgURL     string
	cfgCluster string

	client mermaidclient.Client
)

var rootCmd = &cobra.Command{
	Use:   "sctool",
	Short: "Scylla Manager " + docsVersion,

	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		if cmd.IsAdditionalHelpTopicCommand() || cmd.Hidden {
			return nil
		}

		// Init client
		c, err := mermaidclient.NewClient(cfgURL)
		if err != nil {
			return err
		}
		client = c

		// RequireFlags cluster
		if needsCluster(cmd) {
			if os.Getenv("SCYLLA_MANAGER_CLUSTER") == "" {
				if err := cmd.Root().MarkFlagRequired("cluster"); err != nil {
					return err
				}
			}
		}

		return nil
	},
}

func needsCluster(cmd *cobra.Command) bool {
	switch cmd {
	case clusterAddCmd, clusterListCmd, taskListCmd, versionCmd:
		return false
	}
	return true
}

func init() {
	url := os.Getenv("SCYLLA_MANAGER_API_URL")
	if url == "" {
		url = defaultURL
	}
	rootCmd.PersistentFlags().StringVar(&cfgURL, "api-url", url, "`URL` of Scylla Manager server")

	cluster := os.Getenv("SCYLLA_MANAGER_CLUSTER")
	rootCmd.PersistentFlags().StringVarP(&cfgCluster, "cluster", "c", cluster, "target cluster `name` or ID")
}
