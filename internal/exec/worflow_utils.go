package exec

import (
	"fmt"
	c "github.com/cloudposse/atmos/pkg/config"
	u "github.com/cloudposse/atmos/pkg/utils"
	"strings"
)

func executeWorkflowSteps(workflowDefinition c.WorkflowDefinition, dryRun bool, commandLineStack string) error {
	var steps = workflowDefinition.Steps

	for _, step := range steps {
		var command = strings.TrimSpace(step.Command)
		var commandType = strings.TrimSpace(step.Type)

		u.PrintInfo(fmt.Sprintf("Executing workflow step: %s", command))

		if commandType == "" {
			commandType = "atmos"
		}

		if commandType == "shell" {
			args := strings.Fields(command)
			if err := ExecuteShellCommand(args[0], args[1:], ".", []string{}, dryRun); err != nil {
				return err
			}
		} else if commandType == "atmos" {
			args := strings.Fields(command)

			var workflowStack = strings.TrimSpace(workflowDefinition.Stack)
			var stepStack = strings.TrimSpace(step.Stack)
			var finalStack = ""

			// The workflow `stack` attribute overrides the stack in the `command` (if specified)
			// The step `stack` attribute overrides the stack in the `command` and the workflow `stack` attribute
			// The stack defined on the command line (`atmos workflow <name> -f <file> -s <stack>`) has the highest priority,
			// it overrides all other stacks attributes
			if workflowStack != "" {
				finalStack = workflowStack
			}
			if stepStack != "" {
				finalStack = stepStack
			}
			if commandLineStack != "" {
				finalStack = commandLineStack
			}

			if finalStack != "" {
				args = append(args, []string{"-s", finalStack}...)
				u.PrintInfo(fmt.Sprintf("Stack: %s", finalStack))
			}

			if err := ExecuteShellCommand("atmos", args, ".", []string{}, dryRun); err != nil {
				return err
			}
		} else {
			return fmt.Errorf("invalid workflow step type '%s'. Supported types are 'atmos' and 'shell'", commandType)
		}

		fmt.Println()
	}

	return nil
}
