package vpc

import (
	"fmt"

	"nhncli/internal/output"
	"nhncli/internal/vpc"

	"github.com/spf13/cobra"
)

var (
	subnetVpcID string
	subnetName  string
	subnetCIDR  string
)

var subnetCmd = &cobra.Command{
	Use:   "subnet",
	Short: "서브넷 관리",
}

var subnetListCmd = &cobra.Command{
	Use:   "list",
	Short: "서브넷 목록 조회",
	RunE:  runSubnetList,
}

var subnetDescribeCmd = &cobra.Command{
	Use:   "describe <subnet-id>",
	Short: "서브넷 상세 조회",
	Args:  cobra.ExactArgs(1),
	RunE:  runSubnetDescribe,
}

var subnetCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "서브넷 생성",
	Long: `새로운 서브넷을 생성합니다.

예시:
  nhn vpc subnet create --vpc-id <vpc-id> --name my-subnet --cidr 192.168.1.0/24`,
	RunE: runSubnetCreate,
}

var subnetDeleteCmd = &cobra.Command{
	Use:   "delete <subnet-id>",
	Short: "서브넷 삭제",
	Args:  cobra.ExactArgs(1),
	RunE:  runSubnetDelete,
}

func init() {
	VpcCmd.AddCommand(subnetCmd)
	subnetCmd.AddCommand(subnetListCmd)
	subnetCmd.AddCommand(subnetDescribeCmd)
	subnetCmd.AddCommand(subnetCreateCmd)
	subnetCmd.AddCommand(subnetDeleteCmd)

	subnetListCmd.Flags().StringVar(&subnetVpcID, "vpc-id", "", "특정 VPC의 서브넷만 조회")

	subnetCreateCmd.Flags().StringVar(&subnetVpcID, "vpc-id", "", "VPC ID (필수)")
	subnetCreateCmd.Flags().StringVar(&subnetName, "name", "", "서브넷 이름 (필수)")
	subnetCreateCmd.Flags().StringVar(&subnetCIDR, "cidr", "", "서브넷 CIDR (필수)")
	subnetCreateCmd.MarkFlagRequired("vpc-id")
	subnetCreateCmd.MarkFlagRequired("name")
	subnetCreateCmd.MarkFlagRequired("cidr")
}

func runSubnetList(cmd *cobra.Command, args []string) error {
	client, err := vpc.NewClient(GetProfile(), GetRegion(), GetDebug())
	if err != nil {
		return err
	}

	subnets, err := client.ListSubnets(subnetVpcID)
	if err != nil {
		return err
	}

	if GetOutput() == "json" {
		return output.PrintJSON(subnets)
	}

	headers := []string{"ID", "NAME", "VPC ID", "CIDR", "STATE", "AVAILABLE IPs"}
	rows := make([][]string, len(subnets))
	for i, s := range subnets {
		rows[i] = []string{s.ID, s.Name, s.VPCID, s.CIDR, s.State, fmt.Sprintf("%d", s.AvailableIPCount)}
	}

	return output.Print(GetOutput(), &output.TableData{
		Headers: headers,
		Rows:    rows,
	}, subnets)
}

func runSubnetDescribe(cmd *cobra.Command, args []string) error {
	subnetID := args[0]

	client, err := vpc.NewClient(GetProfile(), GetRegion(), GetDebug())
	if err != nil {
		return err
	}

	s, err := client.GetSubnet(subnetID)
	if err != nil {
		return err
	}

	if GetOutput() == "json" {
		return output.PrintJSON(s)
	}

	fmt.Printf("Subnet ID:       %s\n", s.ID)
	fmt.Printf("Name:            %s\n", s.Name)
	fmt.Printf("VPC ID:          %s\n", s.VPCID)
	fmt.Printf("CIDR:            %s\n", s.CIDR)
	fmt.Printf("Gateway:         %s\n", s.Gateway)
	fmt.Printf("State:           %s\n", s.State)
	fmt.Printf("Available IPs:   %d\n", s.AvailableIPCount)
	fmt.Printf("Routing Table:   %s\n", s.RoutingTableID)
	fmt.Printf("Created:         %s\n", s.CreateTime.Format("2006-01-02 15:04:05"))

	return nil
}

func runSubnetCreate(cmd *cobra.Command, args []string) error {
	client, err := vpc.NewClient(GetProfile(), GetRegion(), GetDebug())
	if err != nil {
		return err
	}

	s, err := client.CreateSubnet(subnetVpcID, subnetName, subnetCIDR)
	if err != nil {
		return err
	}

	if GetOutput() == "json" {
		return output.PrintJSON(s)
	}

	fmt.Printf("✅ 서브넷이 생성되었습니다.\n")
	fmt.Printf("Subnet ID:  %s\n", s.ID)
	fmt.Printf("Name:       %s\n", s.Name)
	fmt.Printf("CIDR:       %s\n", s.CIDR)

	return nil
}

func runSubnetDelete(cmd *cobra.Command, args []string) error {
	subnetID := args[0]

	client, err := vpc.NewClient(GetProfile(), GetRegion(), GetDebug())
	if err != nil {
		return err
	}

	if err := client.DeleteSubnet(subnetID); err != nil {
		return err
	}

	fmt.Printf("✅ 서브넷 %s가 삭제되었습니다.\n", subnetID)
	return nil
}
