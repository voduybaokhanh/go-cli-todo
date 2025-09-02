package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/voduybaokhanh/go-cli-todo/task"
)

var addCmd = &cobra.Command{
	Use:   "add [task name]",
	Short: "Add a new task",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		tasks, err := task.LoadTasks()
		if err != nil {
			return err
		}

		newTask := task.Task{
			ID:    len(tasks) + 1,
			Title: args[0],
			Done:  false,
		}
		tasks = append(tasks, newTask)

		if err := task.SaveTasks(tasks); err != nil {
			return err
		}
		fmt.Println("Added task:", newTask.Title)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
