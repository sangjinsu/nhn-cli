package vpc

import (
	"fmt"

	"nhncli/internal/output"
	"nhncli/internal/vpc"

	"github.com/spf13/cobra"
)

var describeCmd = &cobra.Command{
	Use:   "describe <vpc-id>",
	Short: "VPC 상세 조회",
	Args:  cobra.ExactArgs(1),
	RunE:  runDescribe,
}

func init() {
	VpcCmd.AddCommand(describeCmd)
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
