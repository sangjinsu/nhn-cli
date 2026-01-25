package vpc

import (
	"fmt"

	"nhncli/internal/output"
	"nhncli/internal/vpc"

	"github.com/spf13/cobra"
)

var (
	fipNetworkID  string
	fipInstanceID string
	fipPortID     string
)

var floatingipCmd = &cobra.Command{
	Use:     "floatingip",
	Aliases: []string{"fip"},
	Short:   "플로팅 IP 관리",
}

var fipListCmd = &cobra.Command{
	Use:   "list",
	Short: "플로팅 IP 목록 조회",
	RunE:  runFIPList,
}

var fipDescribeCmd = &cobra.Command{
	Use:   "describe <floatingip-id>",
	Short: "플로팅 IP 상세 조회",
	Args:  cobra.ExactArgs(1),
	RunE:  runFIPDescribe,
}

var fipCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "플로팅 IP 생성",
	RunE:  runFIPCreate,
}

var fipAssociateCmd = &cobra.Command{
	Use:   "associate <floatingip-id>",
	Short: "플로팅 IP 연결",
	Long: `플로팅 IP를 인스턴스 또는 포트에 연결합니다.

예시:
  nhn vpc floatingip associate <floatingip-id> --instance-id <instance-id>
  nhn vpc floatingip associate <floatingip-id> --port-id <port-id>`,
	Args: cobra.ExactArgs(1),
	RunE: runFIPAssociate,
}

var fipDisassociateCmd = &cobra.Command{
	Use:   "disassociate <floatingip-id>",
	Short: "플로팅 IP 연결 해제",
	Args:  cobra.ExactArgs(1),
	RunE:  runFIPDisassociate,
}

var fipDeleteCmd = &cobra.Command{
	Use:   "delete <floatingip-id>",
	Short: "플로팅 IP 삭제",
	Args:  cobra.ExactArgs(1),
	RunE:  runFIPDelete,
}

func init() {
	VpcCmd.AddCommand(floatingipCmd)
	floatingipCmd.AddCommand(fipListCmd)
	floatingipCmd.AddCommand(fipDescribeCmd)
	floatingipCmd.AddCommand(fipCreateCmd)
	floatingipCmd.AddCommand(fipAssociateCmd)
	floatingipCmd.AddCommand(fipDisassociateCmd)
	floatingipCmd.AddCommand(fipDeleteCmd)

	fipCreateCmd.Flags().StringVar(&fipNetworkID, "network-id", "", "외부 네트워크 ID (기본값 사용)")

	fipAssociateCmd.Flags().StringVar(&fipInstanceID, "instance-id", "", "인스턴스 ID")
	fipAssociateCmd.Flags().StringVar(&fipPortID, "port-id", "", "포트 ID")
}

func runFIPList(cmd *cobra.Command, args []string) error {
	client, err := vpc.NewClient(GetProfile(), GetRegion(), GetDebug())
	if err != nil {
		return err
	}

	fips, err := client.ListFloatingIPs()
	if err != nil {
		return err
	}

	if GetOutput() == "json" {
		return output.PrintJSON(fips)
	}

	headers := []string{"ID", "FLOATING IP", "FIXED IP", "STATUS", "PORT ID"}
	rows := make([][]string, len(fips))
	for i, f := range fips {
		portID := f.PortID
		if portID == "" {
			portID = "-"
		}
		fixedIP := f.FixedIPAddress
		if fixedIP == "" {
			fixedIP = "-"
		}
		rows[i] = []string{f.ID, f.FloatingIPAddress, fixedIP, f.Status, portID}
	}

	return output.Print(GetOutput(), &output.TableData{
		Headers: headers,
		Rows:    rows,
	}, fips)
}

func runFIPDescribe(cmd *cobra.Command, args []string) error {
	fipID := args[0]

	client, err := vpc.NewClient(GetProfile(), GetRegion(), GetDebug())
	if err != nil {
		return err
	}

	f, err := client.GetFloatingIP(fipID)
	if err != nil {
		return err
	}

	if GetOutput() == "json" {
		return output.PrintJSON(f)
	}

	fmt.Printf("Floating IP ID:  %s\n", f.ID)
	fmt.Printf("Floating IP:     %s\n", f.FloatingIPAddress)
	fmt.Printf("Fixed IP:        %s\n", f.FixedIPAddress)
	fmt.Printf("Status:          %s\n", f.Status)
	fmt.Printf("Port ID:         %s\n", f.PortID)
	fmt.Printf("Network ID:      %s\n", f.FloatingNetworkID)

	return nil
}

func runFIPCreate(cmd *cobra.Command, args []string) error {
	client, err := vpc.NewClient(GetProfile(), GetRegion(), GetDebug())
	if err != nil {
		return err
	}

	f, err := client.CreateFloatingIP(fipNetworkID)
	if err != nil {
		return err
	}

	if GetOutput() == "json" {
		return output.PrintJSON(f)
	}

	fmt.Printf("✅ 플로팅 IP가 생성되었습니다.\n")
	fmt.Printf("Floating IP ID:  %s\n", f.ID)
	fmt.Printf("Floating IP:     %s\n", f.FloatingIPAddress)

	return nil
}

func runFIPAssociate(cmd *cobra.Command, args []string) error {
	fipID := args[0]

	if fipInstanceID == "" && fipPortID == "" {
		return fmt.Errorf("--instance-id 또는 --port-id를 지정해야 합니다")
	}

	client, err := vpc.NewClient(GetProfile(), GetRegion(), GetDebug())
	if err != nil {
		return err
	}

	portID := fipPortID
	if fipInstanceID != "" && portID == "" {
		port, err := client.GetPortByInstanceID(fipInstanceID)
		if err != nil {
			return err
		}
		portID = port.ID
	}

	f, err := client.AssociateFloatingIP(fipID, portID)
	if err != nil {
		return err
	}

	if GetOutput() == "json" {
		return output.PrintJSON(f)
	}

	fmt.Printf("✅ 플로팅 IP가 연결되었습니다.\n")
	fmt.Printf("Floating IP:  %s\n", f.FloatingIPAddress)
	fmt.Printf("Fixed IP:     %s\n", f.FixedIPAddress)

	return nil
}

func runFIPDisassociate(cmd *cobra.Command, args []string) error {
	fipID := args[0]

	client, err := vpc.NewClient(GetProfile(), GetRegion(), GetDebug())
	if err != nil {
		return err
	}

	f, err := client.DisassociateFloatingIP(fipID)
	if err != nil {
		return err
	}

	if GetOutput() == "json" {
		return output.PrintJSON(f)
	}

	fmt.Printf("✅ 플로팅 IP 연결이 해제되었습니다.\n")
	fmt.Printf("Floating IP:  %s\n", f.FloatingIPAddress)

	return nil
}

func runFIPDelete(cmd *cobra.Command, args []string) error {
	fipID := args[0]

	client, err := vpc.NewClient(GetProfile(), GetRegion(), GetDebug())
	if err != nil {
		return err
	}

	if err := client.DeleteFloatingIP(fipID); err != nil {
		return err
	}

	fmt.Printf("✅ 플로팅 IP %s가 삭제되었습니다.\n", fipID)
	return nil
}
