package main

import (
	"context"
	"fmt"

	"github.com/Kalibriumnet/Kalibrium/cmd/kalibriumwallet/daemon/client"
	"github.com/Kalibriumnet/Kalibrium/cmd/kalibriumwallet/daemon/pb"
	"github.com/Kalibriumnet/Kalibrium/cmd/kalibriumwallet/utils"
)

func balance(conf *balanceConfig) error {
	daemonClient, tearDown, err := client.Connect(conf.DaemonAddress)
	if err != nil {
		return err
	}
	defer tearDown()

	ctx, cancel := context.WithTimeout(context.Background(), daemonTimeout)
	defer cancel()
	response, err := daemonClient.GetBalance(ctx, &pb.GetBalanceRequest{})
	if err != nil {
		return err
	}

	pendingSuffix := ""
	if response.Pending > 0 {
		pendingSuffix = " (pending)"
	}
	if conf.Verbose {
		pendingSuffix = ""
		println("Address                                                                       Available             Pending")
		println("-----------------------------------------------------------------------------------------------------------")
		for _, addressBalance := range response.AddressBalances {
			fmt.Printf("%s %s %s\n", addressBalance.Address, utils.FormatKali(addressBalance.Available), utils.FormatKali(addressBalance.Pending))
		}
		println("-----------------------------------------------------------------------------------------------------------")
		print("                                                 ")
	}
	fmt.Printf("Total balance, KALI %s %s%s\n", utils.FormatKali(response.Available), utils.FormatKali(response.Pending), pendingSuffix)

	return nil
}
