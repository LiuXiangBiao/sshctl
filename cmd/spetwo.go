/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"log"
	"sshctl/service"

	"github.com/spf13/cobra"
)

var spetwoExample = `
》》》======《《《
sshctl spetwo -h [host file path] -c [cmd file path] -u [user] -p [password] 
ps:sshctl spetwo -n /root/hostfile -u root -p 123456 -c /root/cmdfile
》》》======《《《
`

// spetwoCmd represents the spetwo command
var spetwoCmd = &cobra.Command{
	Use:     "spetwo",
	Short:   "Host and Cmd",
	Long:    `Specifies host and cmd file execution`,
	Example: spetwoExample,
	Run: func(cmd *cobra.Command, args []string) {
		var cmdlist []string
		var hostlist []string

		hostpath := service.MustFlag("hp", "string", cmd).(string)
		cmdpath := service.MustFlag("cp", "string", cmd).(string)

		hostlist, err := service.Getfile(hostpath)
		if err != nil {
			log.Fatal("hostlist faield")
		}

		cmdlist, err = service.Getfile(cmdpath)
		if err != nil {
			log.Fatal("cmdlist faield")
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
	rootCmd.AddCommand(spetwoCmd)

	spetwoCmd.Flags().StringP("hp", "h", "", "set hostpath")
	spetwoCmd.Flags().StringP("cp", "c", "", "set cmdpath")
	spetwoCmd.Flags().StringP("user", "u", "", "set user")
	spetwoCmd.Flags().StringP("password", "p", "", "set password")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// spetwoCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// spetwoCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
