/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"log"
	"sshctl/service"

	"github.com/spf13/cobra"
)

var multiplexample = `
》》》------《《《
sshctl multiple -c [commands] -n [ips] -u [user] -p [password]
ps: sshctl multiple -c "lsblk,df -h" -n 192.168.0.61,192.168.0.62 -u root -p 123456
》》》------《《《
`

// shellCmd represents the shell command
var multipleCmd = &cobra.Command{
	Use:     "multiple",
	Short:   "Batch execution",
	Long:    `Execute commands on bulk SSH connections to the server`,
	Example: multiplexample,
	Run: func(cmd *cobra.Command, args []string) {
		var hostlist []string
		var cmdlist []string

		cmds, err := cmd.Flags().GetString("cmds")
		if err != nil {
			log.Fatal("cmds get faield")
		}
		if cmds != "" {
			cmdlist = service.SplitString(cmds)
		}

		hosts, err := cmd.Flags().GetString("hosts")
		if err != nil {
			log.Fatal("hosts get faield")
		}
		if hosts != "" {
			hostlist = service.SplitString(hosts)
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
	rootCmd.AddCommand(multipleCmd)
	multipleCmd.Flags().StringP("cmds", "c", "", "set many command")
	multipleCmd.Flags().StringP("hosts", "n", "", "set many host")
	multipleCmd.Flags().StringP("user", "u", "", "set user")
	multipleCmd.Flags().StringP("password", "p", "", "set password")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// shellCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// shellCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
