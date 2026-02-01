package vpc

import (
	"fmt"

	"nhncli/internal/output"
	"nhncli/internal/vpc"

	"github.com/spf13/cobra"
)

var describeCmd = &cobra.Command{
	Use:               "describe <vpc-id>",
	Short:             "VPC 상세 조회",
	Args:              cobra.ExactArgs(1),
	RunE:              runDescribe,
	ValidArgsFunction: completeVPCIDs,
}

func init() {
	VpcCmd.AddCommand(describeCmd)
}

func completeVPCIDs(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	if len(args) != 0 {
		return nil, cobra.ShellCompDirectiveNoFileComp
	}
	client, err := vpc.NewClient(GetProfile(), GetRegion(), GetDebug())
	if err != nil {
		return nil, cobra.ShellCompDirectiveNoFileComp
	}
	vpcs, err := client.ListVPCs()
	if err != nil {
		return nil, cobra.ShellCompDirectiveNoFileComp
	}
	var completions []string
	for _, v := range vpcs {
		completions = append(completions, fmt.Sprintf("%s\t%s", v.ID, v.Name))
	}
	return completions, cobra.ShellCompDirectiveNoFileComp
}

func runDescribe(cmd *cobra.Command, args []string) error {
	vpcID := args[0]

	client, err := vpc.NewClient(GetProfile(), GetRegion(), GetDebug())
	if err != nil {
		return err
	}

	v, err := client.GetVPC(vpcID)
	if err != nil {
		return err
	}

	if GetOutput() == "json" {
		return output.PrintJSON(v)
	}

	fmt.Printf("VPC ID:      %s\n", v.ID)
	fmt.Printf("Name:        %s\n", v.Name)
	fmt.Printf("CIDR:        %s\n", v.CIDRv4)
	fmt.Printf("State:       %s\n", v.State)
	fmt.Printf("Tenant ID:   %s\n", v.TenantID)
	fmt.Printf("Created:     %s\n", v.CreateTime)

	return nil
}
