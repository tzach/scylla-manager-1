// Copyright (C) 2017 ScyllaDB

package info

import (
	_ "embed"
	"fmt"

	"github.com/scylladb/scylla-manager/pkg/command/flag"
	"github.com/scylladb/scylla-manager/pkg/managerclient"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

//go:embed res.yaml
var res []byte

type command struct {
	cobra.Command
	client *managerclient.Client

	cluster string
	limit   int
}

func NewCommand(client *managerclient.Client) *cobra.Command {
	cmd := &command{
		client: client,
		Command: cobra.Command{
			Args: cobra.ExactArgs(1),
		},
	}
	if err := yaml.Unmarshal(res, &cmd.Command); err != nil {
		panic(err)
	}
	cmd.init()
	cmd.RunE = func(_ *cobra.Command, args []string) error {
		return cmd.run(args)
	}
	return &cmd.Command
}

func (cmd *command) init() {
	defer flag.MustSetUsages(&cmd.Command, res, "cluster")

	w := flag.Wrap(cmd.Flags())
	w.Cluster(&cmd.cluster)
	w.Unwrap().IntVar(&cmd.limit, "limit", 10, "")
}

func (cmd *command) run(args []string) error {
	taskType, taskID, err := cmd.client.TaskSplit(cmd.Context(), cmd.cluster, args[0])
	if err != nil {
		return err
	}

	w := cmd.OutOrStdout()

	task, err := cmd.client.GetTask(cmd.Context(), cmd.cluster, taskType, taskID)
	if err != nil {
		return err
	}

	ti := managerclient.TaskInfo{
		Task: task,
	}
	if err := ti.Render(w); err != nil {
		return err
	}
	fmt.Fprintln(w)

	runs, err := cmd.client.GetTaskHistory(cmd.Context(), cmd.cluster, taskType, taskID, int64(cmd.limit))
	if err != nil {
		return err
	}
	if len(runs) == 0 {
		fmt.Fprintln(w, "No runs yet.")
		return nil
	}

	return runs.Render(w)
}
