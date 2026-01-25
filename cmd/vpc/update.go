package vpc

import (
	"fmt"

	"nhncli/internal/output"
	"nhncli/internal/vpc"

	"github.com/spf13/cobra"
)

var (
	updateName string
	updateCIDR string
)

var updateCmd = &cobra.Command{
	Use:   "update <vpc-id>",
	Short: "VPC 수정",
	Long: `VPC의 이름 또는 CIDR을 수정합니다.

예시:
  nhn vpc update <vpc-id> --name new-name
  nhn vpc update <vpc-id> --cidr 10.0.0.0/16`,
	Args: cobra.ExactArgs(1),
	RunE: runUpdate,
}

func init() {
	VpcCmd.AddCommand(updateCmd)
	updateCmd.Flags().StringVar(&updateName, "name", "", "새 VPC 이름")
	updateCmd.Flags().StringVar(&updateCIDR, "cidr", "", "새 CIDR")
}

func runUpdate(cmd *cobra.Command, args []string) error {
	vpcID := args[0]

	if updateName == "" && updateCIDR == "" {
		return fmt.Errorf("--name 또는 --cidr 중 하나 이상을 지정해야 합니다")
	}

	client, err := vpc.NewClient(GetProfile(), GetRegion(), GetDebug())
	if err != nil {
		return err
	}

	v, err := client.UpdateVPC(vpcID, updateName, updateCIDR)
	if err != nil {
		return err
	}

	if GetOutput() == "json" {
		return output.PrintJSON(v)
	}

	fmt.Printf("✅ VPC가 수정되었습니다.\n")
	fmt.Printf("VPC ID:  %s\n", v.ID)
	fmt.Printf("Name:    %s\n", v.Name)
	fmt.Printf("CIDR:    %s\n", v.CIDRv4)

	return nil
}
