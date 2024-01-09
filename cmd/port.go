/*
Copyright © 2024 OnlyPiglet <jackwuchenghao4@gmail.com>
*/

package cmd

import (
	"fmt"
	"github.com/OnlyPiglet/fly/tcptools"
	"github.com/spf13/cobra"
	"github.com/vishvananda/netlink"
	"log"
	"strconv"
	"syscall"
	"time"
)

const (
	begin  = "\n*************************Listen Socket Info**********************\n"
	end    = "***********************************************\n"
	header = "LocalAddr   | Port | Recv-Q | Send-Q \n"
)

// portCmd represents the port command
var portCmd = &cobra.Command{
	Use:   "port",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		interval := 2 * time.Second
		timeoutInterval := 30 * time.Second

		if len(args) > 0 {
			ine, err := strconv.Atoi(args[len(args)-1])
			if err == nil {
				interval = time.Duration(ine) * time.Second
			}
		}

		ticker := time.NewTicker(interval)

		timeoutTicker := time.NewTicker(timeoutInterval)

		for {
			ticker.Reset(interval)
			timeoutTicker.Reset(timeoutInterval)
			select {
			case <-ticker.C:
				println(SFormatString())
			case <-timeoutTicker.C:
				log.Println(fmt.Sprintf("execute timeout\n"))
			}
		}

	},
}

func SFormatString() string {

	output := begin

	output = output + header

	tcpInfoResps, _ := netlink.SocketDiagTCPInfo(uint8(syscall.AF_INET))

	for _, tcpInfoResp := range tcpInfoResps {

		if tcptools.Listen(tcpInfoResp.InetDiagMsg.State) {

			output = output + fmt.Sprintf("%s       |   %d  |  %d  |  %d \n", tcpInfoResp.InetDiagMsg.ID.Source.String(), tcpInfoResp.InetDiagMsg.ID.SourcePort, tcpInfoResp.InetDiagMsg.RQueue, tcpInfoResp.InetDiagMsg.WQueue)

		}

	}

	output = output + end

	return output
}

func init() {
	ssCmd.AddCommand(portCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// portCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// portCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
