package dingtalk

import (
	"log"

	v "github.com/CatchZeng/gutils/version"
	"github.com/spf13/cobra"
)

const (
	version   = "1.4.6"
	buildTime = "2025/03/12"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "dingtalk version",
	Long:  `dingtalk version`,
	Run:   runVersionCmd,
}

func runVersionCmd(_ *cobra.Command, _ []string) {
	v1 := v.Stringify(version, buildTime)
	log.Println(v1)
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
