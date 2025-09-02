package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/voduybaokhanh/go-cli-todo/task"
)

var doneCmd = &cobra.Command{
	Use:   "done [task id]",
	Short: "Mark a task as done",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			return fmt.Errorf("invalid task ID")
		}

		tasks, err := task.LoadTasks()
		if err != nil {
			return err
		}

		found := false
		for i, t := range tasks {
			if t.ID == id {
				tasks[i].Done = true
				found = true
				break
			}
		}

		if !found {
			fmt.Println("Task not found")
			return nil
		}

		if err := task.SaveTasks(tasks); err != nil {
			return err
		}
		fmt.Println("Task", id, "marked as done.")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(doneCmd)
}
