package compute

import (
	"fmt"
	"strings"

	"nhncli/internal/compute"
	"nhncli/internal/output"

	"github.com/spf13/cobra"
)

var (
	instanceName          string
	instanceImageID       string
	instanceFlavorID      string
	instanceNetworkID     string
	instanceKeyName       string
	instanceSecurityGroup string
	instanceAZ            string
	instanceHardReboot    bool
)

var instanceCmd = &cobra.Command{
	Use:   "instance",
	Short: "인스턴스 관리",
}

var instanceListCmd = &cobra.Command{
	Use:   "list",
	Short: "인스턴스 목록 조회",
	RunE:  runInstanceList,
}

var instanceDescribeCmd = &cobra.Command{
	Use:   "describe <instance-id>",
	Short: "인스턴스 상세 조회",
	Args:  cobra.ExactArgs(1),
	RunE:  runInstanceDescribe,
}

var instanceCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "인스턴스 생성",
	Long: `새로운 인스턴스를 생성합니다.

예시:
  nhn compute instance create \
    --name my-server \
    --image-id <image-id> \
    --flavor-id <flavor-id> \
    --network-id <network-id> \
    --key-name my-keypair \
    --security-group default`,
	RunE: runInstanceCreate,
}

var instanceDeleteCmd = &cobra.Command{
	Use:   "delete <instance-id>",
	Short: "인스턴스 삭제",
	Args:  cobra.ExactArgs(1),
	RunE:  runInstanceDelete,
}

var instanceStartCmd = &cobra.Command{
	Use:   "start <instance-id>",
	Short: "인스턴스 시작",
	Args:  cobra.ExactArgs(1),
	RunE:  runInstanceStart,
}

var instanceStopCmd = &cobra.Command{
	Use:   "stop <instance-id>",
	Short: "인스턴스 중지",
	Args:  cobra.ExactArgs(1),
	RunE:  runInstanceStop,
}

var instanceRebootCmd = &cobra.Command{
	Use:   "reboot <instance-id>",
	Short: "인스턴스 재부팅",
	Args:  cobra.ExactArgs(1),
	RunE:  runInstanceReboot,
}

func init() {
	ComputeCmd.AddCommand(instanceCmd)
	instanceCmd.AddCommand(instanceListCmd)
	instanceCmd.AddCommand(instanceDescribeCmd)
	instanceCmd.AddCommand(instanceCreateCmd)
	instanceCmd.AddCommand(instanceDeleteCmd)
	instanceCmd.AddCommand(instanceStartCmd)
	instanceCmd.AddCommand(instanceStopCmd)
	instanceCmd.AddCommand(instanceRebootCmd)

	instanceCreateCmd.Flags().StringVar(&instanceName, "name", "", "인스턴스 이름 (필수)")
	instanceCreateCmd.Flags().StringVar(&instanceImageID, "image-id", "", "이미지 ID (필수)")
	instanceCreateCmd.Flags().StringVar(&instanceFlavorID, "flavor-id", "", "인스턴스 타입 ID (필수)")
	instanceCreateCmd.Flags().StringVar(&instanceNetworkID, "network-id", "", "네트워크 ID (필수)")
	instanceCreateCmd.Flags().StringVar(&instanceKeyName, "key-name", "", "키페어 이름")
	instanceCreateCmd.Flags().StringVar(&instanceSecurityGroup, "security-group", "", "보안 그룹 이름")
	instanceCreateCmd.Flags().StringVar(&instanceAZ, "availability-zone", "", "가용성 영역")
	instanceCreateCmd.MarkFlagRequired("name")
	instanceCreateCmd.MarkFlagRequired("image-id")
	instanceCreateCmd.MarkFlagRequired("flavor-id")
	instanceCreateCmd.MarkFlagRequired("network-id")

	instanceRebootCmd.Flags().BoolVar(&instanceHardReboot, "hard", false, "하드 리부트 (강제 재부팅)")
}

func runInstanceList(cmd *cobra.Command, args []string) error {
	client, err := compute.NewClient(GetProfile(), GetRegion(), GetDebug())
	if err != nil {
		return err
	}

	instances, err := client.ListInstances()
	if err != nil {
		return err
	}

	if GetOutput() == "json" {
		return output.PrintJSON(instances)
	}

	headers := []string{"ID", "NAME", "STATUS", "FLAVOR", "IP ADDRESSES", "AZ"}
	rows := make([][]string, len(instances))
	for i, inst := range instances {
		var ips []string
		for _, addrs := range inst.Addresses {
			for _, addr := range addrs {
				ips = append(ips, addr.Addr)
			}
		}
		ipStr := strings.Join(ips, ", ")
		if ipStr == "" {
			ipStr = "-"
		}

		rows[i] = []string{
			inst.ID,
			inst.Name,
			inst.Status,
			inst.Flavor.ID,
			ipStr,
			inst.AvailabilityZone,
		}
	}

	return output.Print(GetOutput(), &output.TableData{
		Headers: headers,
		Rows:    rows,
	}, instances)
}

func runInstanceDescribe(cmd *cobra.Command, args []string) error {
	instanceID := args[0]

	client, err := compute.NewClient(GetProfile(), GetRegion(), GetDebug())
	if err != nil {
		return err
	}

	inst, err := client.GetInstance(instanceID)
	if err != nil {
		return err
	}

	if GetOutput() == "json" {
		return output.PrintJSON(inst)
	}

	fmt.Printf("Instance ID:       %s\n", inst.ID)
	fmt.Printf("Name:              %s\n", inst.Name)
	fmt.Printf("Status:            %s\n", inst.Status)
	fmt.Printf("VM State:          %s\n", inst.VMState)
	fmt.Printf("Power State:       %d\n", inst.PowerState)
	fmt.Printf("Flavor ID:         %s\n", inst.Flavor.ID)
	fmt.Printf("Image ID:          %s\n", inst.Image.ID)
	fmt.Printf("Key Name:          %s\n", inst.KeyName)
	fmt.Printf("Availability Zone: %s\n", inst.AvailabilityZone)
	fmt.Printf("Created:           %s\n", inst.Created.Format("2006-01-02 15:04:05"))
	fmt.Printf("Updated:           %s\n", inst.Updated.Format("2006-01-02 15:04:05"))

	fmt.Printf("\nSecurity Groups:\n")
	for _, sg := range inst.SecurityGroups {
		fmt.Printf("  - %s\n", sg.Name)
	}

	fmt.Printf("\nAddresses:\n")
	for network, addrs := range inst.Addresses {
		fmt.Printf("  %s:\n", network)
		for _, addr := range addrs {
			fmt.Printf("    - %s (%s)\n", addr.Addr, addr.Type)
		}
	}

	return nil
}

func runInstanceCreate(cmd *cobra.Command, args []string) error {
	client, err := compute.NewClient(GetProfile(), GetRegion(), GetDebug())
	if err != nil {
		return err
	}

	req := &compute.InstanceCreateRequest{
		Server: compute.InstanceCreateBody{
			Name:      instanceName,
			FlavorRef: instanceFlavorID,
			Networks: []compute.NetworkRef{
				{UUID: instanceNetworkID},
			},
			BlockDeviceMapping: []compute.BlockDeviceMapping{
				{
					BootIndex:           0,
					UUID:                instanceImageID,
					SourceType:          "image",
					DestinationType:     "volume",
					VolumeSize:          20,
					DeleteOnTermination: true,
				},
			},
		},
	}

	if instanceKeyName != "" {
		req.Server.KeyName = instanceKeyName
	}

	if instanceSecurityGroup != "" {
		req.Server.SecurityGroups = []compute.SecurityGroupRef{
			{Name: instanceSecurityGroup},
		}
	}

	if instanceAZ != "" {
		req.Server.AvailabilityZone = instanceAZ
	}

	inst, err := client.CreateInstance(req)
	if err != nil {
		return err
	}

	if GetOutput() == "json" {
		return output.PrintJSON(inst)
	}

	fmt.Printf("✅ 인스턴스가 생성되었습니다.\n")
	fmt.Printf("Instance ID:  %s\n", inst.ID)
	fmt.Printf("Name:         %s\n", inst.Name)
	fmt.Printf("Status:       %s\n", inst.Status)

	return nil
}

func runInstanceDelete(cmd *cobra.Command, args []string) error {
	instanceID := args[0]

	client, err := compute.NewClient(GetProfile(), GetRegion(), GetDebug())
	if err != nil {
		return err
	}

	if err := client.DeleteInstance(instanceID); err != nil {
		return err
	}

	fmt.Printf("✅ 인스턴스 %s가 삭제되었습니다.\n", instanceID)
	return nil
}

func runInstanceStart(cmd *cobra.Command, args []string) error {
	instanceID := args[0]

	client, err := compute.NewClient(GetProfile(), GetRegion(), GetDebug())
	if err != nil {
		return err
	}

	if err := client.StartInstance(instanceID); err != nil {
		return err
	}

	fmt.Printf("✅ 인스턴스 %s 시작 요청이 완료되었습니다.\n", instanceID)
	return nil
}

func runInstanceStop(cmd *cobra.Command, args []string) error {
	instanceID := args[0]

	client, err := compute.NewClient(GetProfile(), GetRegion(), GetDebug())
	if err != nil {
		return err
	}

	if err := client.StopInstance(instanceID); err != nil {
		return err
	}

	fmt.Printf("✅ 인스턴스 %s 중지 요청이 완료되었습니다.\n", instanceID)
	return nil
}

func runInstanceReboot(cmd *cobra.Command, args []string) error {
	instanceID := args[0]

	client, err := compute.NewClient(GetProfile(), GetRegion(), GetDebug())
	if err != nil {
		return err
	}

	if err := client.RebootInstance(instanceID, instanceHardReboot); err != nil {
		return err
	}

	rebootType := "소프트"
	if instanceHardReboot {
		rebootType = "하드"
	}
	fmt.Printf("✅ 인스턴스 %s %s 재부팅 요청이 완료되었습니다.\n", instanceID, rebootType)
	return nil
}
