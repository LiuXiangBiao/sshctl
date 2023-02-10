/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"log"
	"sshctl/service"

	"github.com/spf13/cobra"
)

var specmdExample = `
》》》======《《《
sshctl specmd -f [cmd file path] -n [hosts] -u [user] -p [password] 
ps: sshctl specmd -n "192.168.0.61,192.168.0.62" -u root -p 123456  -f /root/cmdfile
》》》======《《《
`

// specmdCmd represents the specmd command
var specmdCmd = &cobra.Command{
	Use:     "specmd",
	Short:   "CmdFilePath",
	Long:    `Specify the cmd file path to execute`,
	Example: specmdExample,
	Run: func(cmd *cobra.Command, args []string) {
		var cmdlist []string
		var hostlist []string

		filepath := service.MustFlag("filepathd", "string", cmd).(string)

		hosts, err := cmd.Flags().GetString("hosts")
		if err != nil {
			log.Fatal("hosts get faield")
		}
		if hosts != "" {
			hostlist = service.SplitString(hosts)
		}

		cmdlist, err = service.Getfile(filepath)
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
	rootCmd.AddCommand(specmdCmd)

	specmdCmd.Flags().StringP("filepathd", "f", "", "set cmd file")
	specmdCmd.Flags().StringP("hosts", "n", "", "set many host")
	specmdCmd.Flags().StringP("user", "u", "", "set user")
	specmdCmd.Flags().StringP("password", "p", "", "set password")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// specmdCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// specmdCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
