/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"sshctl/service"
)

var spehostExample = `
》》》======《《《
sshctl spehost -f [host file path] -u [user] -p [password] -c [commands]
ps: sshctl spehost -u root -p 123456 -c "free -h,lsblk" -f /root/hostfile
》》》======《《《
`

// spehostCmd represents the spehost command
var spehostCmd = &cobra.Command{
	Use:     "spehost",
	Short:   "HostFilePath ",
	Long:    `Specify the host file path to execute`,
	Example: spehostExample,
	Run: func(cmd *cobra.Command, args []string) {
		var cmdlist []string
		var hostlist []string

		filepath := service.MustFlag("filepathd", "string", cmd).(string)

		cmds, err := cmd.Flags().GetString("cmds")
		if err != nil {
			log.Fatal("commands get faield")
		}
		if cmds != "" {
			cmdlist = service.SplitString(cmds)
		}

		hostlist, err = service.Getfile(filepath)
		if err != nil {
			log.Fatal("get host faield")
		}

		username := service.MustFlag("user", "string", cmd).(string)
		pwd, err := cmd.Flags().GetString("password")
		if err != nil {
			log.Fatal("password error")
		}

		service.SshExec(username, pwd, hostlist, cmdlist, nil)
	},
}

func init() {
	rootCmd.AddCommand(spehostCmd)

	spehostCmd.Flags().StringP("filepathd", "f", "", "set hostfile path")
	spehostCmd.Flags().StringP("cmds", "c", "", "set many command")
	spehostCmd.Flags().StringP("user", "u", "", "set user")
	spehostCmd.Flags().StringP("password", "p", "", "set password")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// spehostCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// spehostCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
