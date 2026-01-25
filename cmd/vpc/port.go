package vpc

import (
	"fmt"
	"strings"

	"nhncli/internal/output"
	"nhncli/internal/vpc"

	"github.com/spf13/cobra"
)

var (
	portNetworkID string
	portName      string
)

var portCmd = &cobra.Command{
	Use:   "port",
	Short: "네트워크 인터페이스 관리",
}

var portListCmd = &cobra.Command{
	Use:   "list",
	Short: "네트워크 인터페이스 목록 조회",
	RunE:  runPortList,
}

var portDescribeCmd = &cobra.Command{
	Use:   "describe <port-id>",
	Short: "네트워크 인터페이스 상세 조회",
	Args:  cobra.ExactArgs(1),
	RunE:  runPortDescribe,
}

var portCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "네트워크 인터페이스 생성",
	RunE:  runPortCreate,
}

var portDeleteCmd = &cobra.Command{
	Use:   "delete <port-id>",
	Short: "네트워크 인터페이스 삭제",
	Args:  cobra.ExactArgs(1),
	RunE:  runPortDelete,
}

func init() {
	VpcCmd.AddCommand(portCmd)
	portCmd.AddCommand(portListCmd)
	portCmd.AddCommand(portDescribeCmd)
	portCmd.AddCommand(portCreateCmd)
	portCmd.AddCommand(portDeleteCmd)

	portCreateCmd.Flags().StringVar(&portNetworkID, "network-id", "", "네트워크 ID (필수)")
	portCreateCmd.Flags().StringVar(&portName, "name", "", "포트 이름")
	portCreateCmd.MarkFlagRequired("network-id")
}

func runPortList(cmd *cobra.Command, args []string) error {
	client, err := vpc.NewClient(GetProfile(), GetRegion(), GetDebug())
	if err != nil {
		return err
	}

	ports, err := client.ListPorts()
	if err != nil {
		return err
	}

	if GetOutput() == "json" {
		return output.PrintJSON(ports)
	}

	headers := []string{"ID", "NAME", "STATUS", "MAC ADDRESS", "FIXED IPs", "DEVICE ID"}
	rows := make([][]string, len(ports))
	for i, p := range ports {
		var ips []string
		for _, ip := range p.FixedIPs {
			ips = append(ips, ip.IPAddress)
		}
		ipStr := strings.Join(ips, ", ")
		if ipStr == "" {
			ipStr = "-"
		}

		deviceID := p.DeviceID
		if deviceID == "" {
			deviceID = "-"
		} else if len(deviceID) > 12 {
			deviceID = deviceID[:12] + "..."
		}

		name := p.Name
		if name == "" {
			name = "-"
		}

		rows[i] = []string{p.ID, name, p.Status, p.MACAddress, ipStr, deviceID}
	}

	return output.Print(GetOutput(), &output.TableData{
		Headers: headers,
		Rows:    rows,
	}, ports)
}

func runPortDescribe(cmd *cobra.Command, args []string) error {
	portID := args[0]

	client, err := vpc.NewClient(GetProfile(), GetRegion(), GetDebug())
	if err != nil {
		return err
	}

	p, err := client.GetPort(portID)
	if err != nil {
		return err
	}

	if GetOutput() == "json" {
		return output.PrintJSON(p)
	}

	fmt.Printf("Port ID:       %s\n", p.ID)
	fmt.Printf("Name:          %s\n", p.Name)
	fmt.Printf("Network ID:    %s\n", p.NetworkID)
	fmt.Printf("MAC Address:   %s\n", p.MACAddress)
	fmt.Printf("Status:        %s\n", p.Status)
	fmt.Printf("Admin State:   %v\n", p.AdminStateUp)
	fmt.Printf("Device ID:     %s\n", p.DeviceID)
	fmt.Printf("Device Owner:  %s\n", p.DeviceOwner)

	if len(p.FixedIPs) > 0 {
		fmt.Printf("Fixed IPs:\n")
		for _, ip := range p.FixedIPs {
			fmt.Printf("  - %s (subnet: %s)\n", ip.IPAddress, ip.SubnetID)
		}
	}

	return nil
}

func runPortCreate(cmd *cobra.Command, args []string) error {
	client, err := vpc.NewClient(GetProfile(), GetRegion(), GetDebug())
	if err != nil {
		return err
	}

	p, err := client.CreatePort(portNetworkID, portName)
	if err != nil {
		return err
	}

	if GetOutput() == "json" {
		return output.PrintJSON(p)
	}

	fmt.Printf("✅ 네트워크 인터페이스가 생성되었습니다.\n")
	fmt.Printf("Port ID:      %s\n", p.ID)
	fmt.Printf("MAC Address:  %s\n", p.MACAddress)

	return nil
}

func runPortDelete(cmd *cobra.Command, args []string) error {
	portID := args[0]

	client, err := vpc.NewClient(GetProfile(), GetRegion(), GetDebug())
	if err != nil {
		return err
	}

	if err := client.DeletePort(portID); err != nil {
		return err
	}

	fmt.Printf("✅ 네트워크 인터페이스 %s가 삭제되었습니다.\n", portID)
	return nil
}
