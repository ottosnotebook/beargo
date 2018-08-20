package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// WatchOptions ...
type WatchOptions struct {
	BuildOptions
	Addr string
}

func initWatchCmd() *cobra.Command {

	opts := &WatchOptions{BuildOptions: BuildOptions{}}

	cmd := &cobra.Command{
		Use:   "watch",
		Short: "Watch for changes to Bear notes, regenerate the site using Hugo, and serve the contents",
		Long:  "Watch for changes to Bear notes, regenerate the site using Hugo, and serve the contents",

		RunE: func(cmd *cobra.Command, args []string) error {
			return Watch(opts)
		},
	}

	cmd.Flags().StringVar(&opts.BearSQLiteDBPath, "bear-db", "", "Path to Bear SQLite database.")
	cmd.Flags().BoolVar(&opts.UseHugoExec, "hugo-exec", false, "Call the version of Hugo installed on the host machine rather than use the latest version of Hugo as a library.")

	return cmd
}

// Watch ...
func Watch(opts *WatchOptions) error {
	return fmt.Errorf("watch: not implemented")
}
