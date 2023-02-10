/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"sshctl/service"
)

var remotyexample = `
》》》===《《《
sshctl remoty -r [远程文件] -l [本地文件] -n [hostip] -u [user] -p [password]
ps: sshctl remoty -n 192.168.0.61 -r /root/sysconfigure.sh -l /root/lxb.sh -u root -p123456
》》》===《《《
`

// remotyCmd represents the remoty command
var remotyCmd = &cobra.Command{
	Use:   "remoty",
	Short: "Remote files",
	Long:  `Copy the contents of the remote file`,
	Run: func(cmd *cobra.Command, args []string) {
		remotepath := service.MustFlag("rp", "string", cmd).(string)
		localpath := service.MustFlag("lp", "string", cmd).(string)
		host := service.MustFlag("hosts", "string", cmd).(string)
		username := service.MustFlag("user", "string", cmd).(string)
		pwd, err := cmd.Flags().GetString("password")
		if err != nil {
			log.Fatal("password error")
		}

		service.CopyExec(remotepath, localpath, host, username, pwd, nil)
	},
}

func init() {
	rootCmd.AddCommand(remotyCmd)

	remotyCmd.Flags().StringP("rp", "r", "", "set remote file")
	remotyCmd.Flags().StringP("lp", "l", "", "set local file")
	remotyCmd.Flags().StringP("hosts", "n", "", "set many host")
	remotyCmd.Flags().StringP("user", "u", "", "set user")
	remotyCmd.Flags().StringP("password", "p", "", "set password")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// remotyCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// remotyCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
