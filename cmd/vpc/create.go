package vpc

import (
	"fmt"

	"nhncli/internal/output"
	"nhncli/internal/vpc"

	"github.com/spf13/cobra"
)

var (
	createName string
	createCIDR string
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "VPC 생성",
	Long: `새로운 VPC를 생성합니다.

예시:
  nhn vpc create --name my-vpc --cidr 192.168.0.0/16`,
	RunE: runCreate,
}

func init() {
	VpcCmd.AddCommand(createCmd)
	createCmd.Flags().StringVar(&createName, "name", "", "VPC 이름 (필수)")
	createCmd.Flags().StringVar(&createCIDR, "cidr", "", "VPC CIDR (필수, 예: 192.168.0.0/16)")
	createCmd.MarkFlagRequired("name")
	createCmd.MarkFlagRequired("cidr")
}

func runCreate(cmd *cobra.Command, args []string) error {
	client, err := vpc.NewClient(GetProfile(), GetRegion(), GetDebug())
	if err != nil {
		return err
	}

	v, err := client.CreateVPC(createName, createCIDR)
	if err != nil {
		return err
	}

	if GetOutput() == "json" {
		return output.PrintJSON(v)
	}

	fmt.Printf("✅ VPC가 생성되었습니다.\n")
	fmt.Printf("VPC ID:  %s\n", v.ID)
	fmt.Printf("Name:    %s\n", v.Name)
	fmt.Printf("CIDR:    %s\n", v.CIDRv4)

	return nil
}
