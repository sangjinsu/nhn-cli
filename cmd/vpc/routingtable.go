package vpc

import (
	"fmt"

	"nhncli/internal/output"
	"nhncli/internal/vpc"

	"github.com/spf13/cobra"
)

var routingtableCmd = &cobra.Command{
	Use:     "routingtable",
	Aliases: []string{"rt"},
	Short:   "라우팅 테이블 조회",
}

var rtListCmd = &cobra.Command{
	Use:   "list",
	Short: "라우팅 테이블 목록 조회",
	RunE:  runRTList,
}

var rtDescribeCmd = &cobra.Command{
	Use:   "describe <routingtable-id>",
	Short: "라우팅 테이블 상세 조회",
	Args:  cobra.ExactArgs(1),
	RunE:  runRTDescribe,
}

func init() {
	VpcCmd.AddCommand(routingtableCmd)
	routingtableCmd.AddCommand(rtListCmd)
	routingtableCmd.AddCommand(rtDescribeCmd)
}

func runRTList(cmd *cobra.Command, args []string) error {
	client, err := vpc.NewClient(GetProfile(), GetRegion(), GetDebug())
	if err != nil {
		return err
	}

	rts, err := client.ListRoutingTables()
	if err != nil {
		return err
	}

	if GetOutput() == "json" {
		return output.PrintJSON(rts)
	}

	headers := []string{"ID", "NAME", "VPC ID", "DEFAULT", "STATE", "SUBNETS"}
	rows := make([][]string, len(rts))
	for i, rt := range rts {
		defaultTable := "No"
		if rt.DefaultTable {
			defaultTable = "Yes"
		}
		rows[i] = []string{rt.ID, rt.Name, rt.VPCID, defaultTable, rt.State, fmt.Sprintf("%d", rt.SubnetCount)}
	}

	return output.Print(GetOutput(), &output.TableData{
		Headers: headers,
		Rows:    rows,
	}, rts)
}

func runRTDescribe(cmd *cobra.Command, args []string) error {
	rtID := args[0]

	client, err := vpc.NewClient(GetProfile(), GetRegion(), GetDebug())
	if err != nil {
		return err
	}

	rt, err := client.GetRoutingTable(rtID)
	if err != nil {
		return err
	}

	if GetOutput() == "json" {
		return output.PrintJSON(rt)
	}

	fmt.Printf("Routing Table ID:  %s\n", rt.ID)
	fmt.Printf("Name:              %s\n", rt.Name)
	fmt.Printf("VPC ID:            %s\n", rt.VPCID)
	fmt.Printf("Default Table:     %v\n", rt.DefaultTable)
	fmt.Printf("Distributed:       %v\n", rt.Distributed)
	fmt.Printf("State:             %s\n", rt.State)
	fmt.Printf("Subnet Count:      %d\n", rt.SubnetCount)

	if len(rt.Routes) > 0 {
		fmt.Printf("\nRoutes:\n")
		for _, route := range rt.Routes {
			fmt.Printf("  - %s -> %s\n", route.DestinationCIDR, route.GatewayID)
		}
	}

	return nil
}
