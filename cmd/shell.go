/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/ssh/terminal"
	"log"
	"sshctl/service"
)

var shellexample = `
》》》------《《《
sshctl shell -n [hostIP] -u [user]  -c [command]
password:
ps: sshctl shell -n 192.168.0.61 -u root -c "lsblk&&df -h"
》》》------《《《
`

// commandCmd represents the command command
var shellCmd = &cobra.Command{
	Use:     "shell",
	Short:   "One execution per host",
	Long:    `Connect to the server and execute the command`,
	Example: shellexample,
	Run: func(cmd *cobra.Command, args []string) {
		ip := service.MustFlag("hosts", "string", cmd).(string)
		user := service.MustFlag("user", "string", cmd).(string)
		command := service.MustFlag("command", "string", cmd).(string)

		var count = 0
		for count < 3 {
			fmt.Print("password：")
			pwd, err := terminal.ReadPassword(0)
			if err != nil {
				log.Fatal(err)
			}
			err = service.ExecSingleCmd(user, string(pwd), ip, command, nil)
			if err != nil {
				fmt.Println("error password,try again")
				count++
			} else {
				break
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(shellCmd)

	shellCmd.Flags().StringP("hosts", "n", "", "set hostIP")
	shellCmd.Flags().StringP("user", "u", "", "set user")
	shellCmd.Flags().StringP("command", "c", "", "set command")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// commandCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// commandCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
