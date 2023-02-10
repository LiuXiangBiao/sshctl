/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/spf13/cobra"
	"sshctl/service"
)

var pingcmdexampl = `
》》》====《《《
sshctl ping -n [hostip]
ps: sshctl ping -n 192.168.0.62
》》》====《《《
`

// pingCmd represents the ping command
var pingCmd = &cobra.Command{
	Use:   "ping",
	Short: "remote ping",
	Long:  `Supports remote ping of the host`,
	Run: func(cmd *cobra.Command, args []string) {
		host := service.MustFlag("host", "string", cmd).(string)
		service.PingExec(host)
	},
}

func init() {
	rootCmd.AddCommand(pingCmd)
	pingCmd.Flags().StringP("host", "n", "", "set hostip")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// pingCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// pingCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
