// %N0 CARR13R█

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/n0carr1er/sauce/version"
)

func init() {
	RootCmd.AddCommand(versionCmd)
}

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Prints the version of sauce",
	Run:   versionCommandFunc,
}

func versionCommandFunc(cmd *cobra.Command, args []string) {
	fmt.Println("sauce version:", version.Version)
	fmt.Println("specification version:", version.SpecificationVersion)
	fmt.Println("specification revision:", version.SpecificationRevision)
}
