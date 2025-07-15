package git

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func BranchExists(branch string) bool {
	cmd := exec.Command("git", "show-ref", "--verify", "--quiet", fmt.Sprintf("refs/heads/%s", branch))
	return cmd.Run() == nil
}

func GetBaseBranch() string {
	if BranchExists("dev") {
		return "dev"
	} else if BranchExists("development") {
		return "development"
	}
	return "main"
}

func CreateAndPushBranch(branchName, baseBranch string) error {
	steps := [][]string{
		{"checkout", baseBranch},
		{"pull", "origin", baseBranch},
		{"checkout", "-b", branchName},
		{"push", "--set-upstream", "origin", branchName},
	}

	for _, args := range steps {
		cmd := exec.Command("git", args...)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			return fmt.Errorf("git command failed: git %s", strings.Join(args, " "))
		}
	}

	return nil
}
