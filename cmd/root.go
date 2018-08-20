package cmd

import "github.com/spf13/cobra"

var root *cobra.Command

func init() {
	root = &cobra.Command{
		Use:          "beargo",
		SilenceUsage: true,

		Short: "Author blog content in Bear. Generate a static site using Hugo",
		Long:  "Author blog content in Bear. Generate a static site using Hugo",
	}

	opts := &BuildOptions{}

	// flags
	root.Flags().StringVar(&opts.ContentDirectory, "content-dir", "content", "Name of Hugo content directory")
	root.Flags().StringVar(&opts.ContentSectionDirectory, "content-section", "post", "Name of content section directory")
	root.Flags().StringVar(&opts.BearSQLiteDBPath, "bear-db", "", "Path to Bear SQLite database")
	root.Flags().BoolVar(&opts.UseHugoExec, "hugo-exec", false, "Call the version of Hugo installed on the host machine rather than use the latest version of Hugo as a library")

	// commands
	root.AddCommand(initWatchCmd())

	// run
	root.RunE = func(cmd *cobra.Command, args []string) error {
		return Build(opts)
	}

}

// Execute ...
func Execute() error {
	return root.Execute()
}
