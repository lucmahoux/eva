package open

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"

	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "open",
		Short: "Open the current GitHub branch in your browser",
		Run: func(cmd *cobra.Command, args []string) {
			currentBranch := getCurrentBranch()
			repoURL := getRepoURL()
			branchURL := fmt.Sprintf("%s/tree/%s", repoURL, currentBranch)

			fmt.Printf("🌐 Opening GitHub branch: %s\n", branchURL)
			openBrowser(branchURL)
		},
	}
}

func getCurrentBranch() string {
	cmd := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD")
	out, err := cmd.Output()
	if err != nil {
		fmt.Println("❌ Failed to get current branch")
		os.Exit(1)
	}
	return strings.TrimSpace(string(out))
}

func getRepoURL() string {
	cmd := exec.Command("git", "config", "--get", "remote.origin.url")
	out, err := cmd.Output()
	if err != nil {
		fmt.Println("❌ Failed to get remote URL")
		os.Exit(1)
	}

	remote := strings.TrimSpace(string(out))

	// Convert SSH to HTTPS
	if strings.HasPrefix(remote, "git@") {
		re := regexp.MustCompile(`git@([^:]+):(.+)\.git`)
		matches := re.FindStringSubmatch(remote)
		if len(matches) == 3 {
			return fmt.Sprintf("https://github.com/%s", matches[2])
		}
	}

	// Convert HTTPS ending with .git
	if strings.HasPrefix(remote, "https://") && strings.HasSuffix(remote, ".git") {
		return strings.TrimSuffix(remote, ".git")
	}

	return remote
}

func openBrowser(url string) {
	commands := []string{"open", "xdg-open", "start"} // macOS, Linux, Windows
	for _, cmdName := range commands {
		cmd := exec.Command(cmdName, url)
		if err := cmd.Start(); err == nil {
			return
		}
	}
	fmt.Printf("🔗 Could not open browser. Here's the URL: %s\n", url)
}
