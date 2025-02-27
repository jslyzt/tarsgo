package cmake

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/jslyzt/tarsgo/tars/tools/tarsgo/internal/base"
	"github.com/jslyzt/tarsgo/tars/tools/tarsgo/internal/consts"

	"github.com/spf13/cobra"
)

// CmdNew represents the new command.
var CmdNew = &cobra.Command{
	Use:   "cmake App Server Servant GoModuleName",
	Short: "Create a service cmake template",
	Long: `Create a service cmake project using the repository template. Example:
tarsgo cmake TeleSafe PhonenumSogouServer SogouInfo github.com/TeleSafe/PhonenumSogouServer`,
	Run: run,
}

var (
	repoUrl string
	branch  string
	timeout string
)

func init() {
	timeout = "60s"
	CmdNew.Flags().StringVarP(&repoUrl, "repo-url", "r", consts.RepoURL, "layout repo")
	CmdNew.Flags().StringVarP(&branch, "branch", "b", branch, "repo branch")
	CmdNew.Flags().StringVarP(&timeout, "timeout", "t", timeout, "time out")
}

func run(cmd *cobra.Command, args []string) {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	t, err := time.ParseDuration(timeout)
	if err != nil {
		panic(err)
	}
	app, server, servant, goModuleName, err := base.GetArgs(cmd, args)
	if err != nil {
		return
	}
	p := base.NewProject(app, server, servant, goModuleName)
	ctx, cancel := context.WithTimeout(context.Background(), t)
	defer cancel()
	done := make(chan error, 1)
	go func() {
		done <- p.Create(ctx, wd, repoUrl, branch, consts.CMakeDemoDir)
	}()
	select {
	case <-ctx.Done():
		if errors.Is(ctx.Err(), context.DeadlineExceeded) {
			fmt.Fprint(os.Stderr, "\033[31mERROR: project creation timed out\033[m\n")
		} else {
			fmt.Fprintf(os.Stderr, "\033[31mERROR: failed to create project(%+v)\033[m\n", ctx.Err().Error())
		}
	case err = <-done:
		if err != nil {
			fmt.Fprintf(os.Stderr, "\033[31mERROR: Failed to create project(%+v)\033[m\n", err.Error())
		}
	}
}
