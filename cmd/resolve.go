package cmd

import (
	"github.com/spf13/cobra"
)

var resolveCmd = &cobra.Command{
	Use:   "resolve",
	Short: "Resolve a domain",
	Long:  "Resolve records of a domain. Domain must be specified",
}
