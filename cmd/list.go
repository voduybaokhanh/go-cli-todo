package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/voduybaokhanh/go-cli-todo/task"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all tasks",
	RunE: func(cmd *cobra.Command, args []string) error {
		tasks, err := task.LoadTasks()
		if err != nil {
			return err
		}

		if len(tasks) == 0 {
			fmt.Println("No tasks found.")
			return nil
		}

		for _, t := range tasks {
			status := " "
			if t.Done {
				status = "x"
			}
			fmt.Printf("[%s] %d: %s\n", status, t.ID, t.Title)
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
