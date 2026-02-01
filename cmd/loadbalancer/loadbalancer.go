package loadbalancer

import (
	"fmt"
	"strings"

	"nhncli/cmd"
	"nhncli/internal/loadbalancer"
	"nhncli/internal/output"

	"github.com/spf13/cobra"
)

var (
	lbName       string
	lbSubnetID   string
	lbVipAddress string
	lbDesc       string
	lbType       string
	lbUpdateName string
	lbUpdateDesc string
)

var LoadbalancerCmd = &cobra.Command{
	Use:     "loadbalancer",
	Aliases: []string{"lb"},
	Short:   "Load Balancer 서비스 관리",
	Long: `Load Balancer 서비스 관련 리소스를 관리합니다.

지원 리소스:
  - loadbalancer: 로드 밸런서 관리
  - listener: 리스너 관리`,
}

var lbListCmd = &cobra.Command{
	Use:   "list",
	Short: "로드 밸런서 목록 조회",
	RunE:  runLBList,
}

var lbDescribeCmd = &cobra.Command{
	Use:               "describe <lb-id>",
	Short:             "로드 밸런서 상세 조회",
	Args:              cobra.ExactArgs(1),
	ValidArgsFunction: completeLBIDs,
	RunE:              runLBDescribe,
}

var lbCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "로드 밸런서 생성",
	Long: `새로운 로드 밸런서를 생성합니다.

예시:
  nhn loadbalancer create \
    --name my-lb \
    --vip-subnet-id <subnet-id>`,
	RunE: runLBCreate,
}

var lbUpdateCmd = &cobra.Command{
	Use:               "update <lb-id>",
	Short:             "로드 밸런서 수정",
	Args:              cobra.ExactArgs(1),
	ValidArgsFunction: completeLBIDs,
	RunE:              runLBUpdate,
}

var lbDeleteCmd = &cobra.Command{
	Use:               "delete <lb-id>",
	Short:             "로드 밸런서 삭제",
	Args:              cobra.ExactArgs(1),
	ValidArgsFunction: completeLBIDs,
	RunE:              runLBDelete,
}

func init() {
	cmd.GetRootCmd().AddCommand(LoadbalancerCmd)

	LoadbalancerCmd.AddCommand(lbListCmd)
	LoadbalancerCmd.AddCommand(lbDescribeCmd)
	LoadbalancerCmd.AddCommand(lbCreateCmd)
	LoadbalancerCmd.AddCommand(lbUpdateCmd)
	LoadbalancerCmd.AddCommand(lbDeleteCmd)

	lbCreateCmd.Flags().StringVar(&lbName, "name", "", "로드 밸런서 이름")
	lbCreateCmd.Flags().StringVar(&lbSubnetID, "vip-subnet-id", "", "VIP 서브넷 ID (필수)")
	lbCreateCmd.Flags().StringVar(&lbVipAddress, "vip-address", "", "VIP 주소")
	lbCreateCmd.Flags().StringVar(&lbDesc, "description", "", "설명")
	lbCreateCmd.Flags().StringVar(&lbType, "type", "", "타입 (shared/dedicated)")
	lbCreateCmd.MarkFlagRequired("vip-subnet-id")

	lbUpdateCmd.Flags().StringVar(&lbUpdateName, "name", "", "로드 밸런서 이름")
	lbUpdateCmd.Flags().StringVar(&lbUpdateDesc, "description", "", "설명")
}

func GetProfile() string {
	return cmd.GetProfile()
}

func GetRegion() string {
	return cmd.GetRegion()
}

func GetOutput() string {
	return cmd.GetOutput()
}

func GetDebug() bool {
	return cmd.GetDebug()
}

func runLBList(c *cobra.Command, args []string) error {
	client, err := loadbalancer.NewClient(GetProfile(), GetRegion(), GetDebug())
	if err != nil {
		return err
	}

	lbs, err := client.ListLoadBalancers()
	if err != nil {
		return err
	}

	if GetOutput() == "json" {
		return output.PrintJSON(lbs)
	}

	headers := []string{"ID", "NAME", "STATUS", "VIP ADDRESS", "TYPE"}
	rows := make([][]string, len(lbs))
	for i, lb := range lbs {
		rows[i] = []string{
			lb.ID,
			lb.Name,
			lb.ProvisioningStatus,
			lb.VipAddress,
			lb.LoadbalancerType,
		}
	}

	return output.Print(GetOutput(), &output.TableData{
		Headers: headers,
		Rows:    rows,
	}, lbs)
}

func runLBDescribe(c *cobra.Command, args []string) error {
	lbID := args[0]

	client, err := loadbalancer.NewClient(GetProfile(), GetRegion(), GetDebug())
	if err != nil {
		return err
	}

	lb, err := client.GetLoadBalancer(lbID)
	if err != nil {
		return err
	}

	if GetOutput() == "json" {
		return output.PrintJSON(lb)
	}

	fmt.Printf("LB ID:              %s\n", lb.ID)
	fmt.Printf("Name:               %s\n", lb.Name)
	fmt.Printf("Description:        %s\n", lb.Description)
	fmt.Printf("Provisioning Status:%s\n", lb.ProvisioningStatus)
	fmt.Printf("Operating Status:   %s\n", lb.OperatingStatus)
	fmt.Printf("VIP Address:        %s\n", lb.VipAddress)
	fmt.Printf("VIP Subnet ID:      %s\n", lb.VipSubnetID)
	fmt.Printf("Type:               %s\n", lb.LoadbalancerType)

	if len(lb.Listeners) > 0 {
		ids := make([]string, len(lb.Listeners))
		for i, l := range lb.Listeners {
			ids[i] = l.ID
		}
		fmt.Printf("Listeners:          %s\n", strings.Join(ids, ", "))
	}

	return nil
}

func runLBCreate(c *cobra.Command, args []string) error {
	client, err := loadbalancer.NewClient(GetProfile(), GetRegion(), GetDebug())
	if err != nil {
		return err
	}

	req := &loadbalancer.LoadBalancerCreateRequest{
		Loadbalancer: loadbalancer.LoadBalancerCreateBody{
			VipSubnetID: lbSubnetID,
		},
	}

	if lbName != "" {
		req.Loadbalancer.Name = lbName
	}
	if lbVipAddress != "" {
		req.Loadbalancer.VipAddress = lbVipAddress
	}
	if lbDesc != "" {
		req.Loadbalancer.Description = lbDesc
	}
	if lbType != "" {
		req.Loadbalancer.LoadbalancerType = lbType
	}

	lb, err := client.CreateLoadBalancer(req)
	if err != nil {
		return err
	}

	if GetOutput() == "json" {
		return output.PrintJSON(lb)
	}

	fmt.Printf("✅ 로드 밸런서가 생성되었습니다.\n")
	fmt.Printf("LB ID:       %s\n", lb.ID)
	fmt.Printf("Name:        %s\n", lb.Name)
	fmt.Printf("VIP Address: %s\n", lb.VipAddress)
	fmt.Printf("Status:      %s\n", lb.ProvisioningStatus)

	return nil
}

func runLBUpdate(c *cobra.Command, args []string) error {
	lbID := args[0]

	client, err := loadbalancer.NewClient(GetProfile(), GetRegion(), GetDebug())
	if err != nil {
		return err
	}

	req := &loadbalancer.LoadBalancerUpdateRequest{
		Loadbalancer: loadbalancer.LoadBalancerUpdateBody{},
	}

	if lbUpdateName != "" {
		req.Loadbalancer.Name = lbUpdateName
	}
	if lbUpdateDesc != "" {
		req.Loadbalancer.Description = lbUpdateDesc
	}

	lb, err := client.UpdateLoadBalancer(lbID, req)
	if err != nil {
		return err
	}

	if GetOutput() == "json" {
		return output.PrintJSON(lb)
	}

	fmt.Printf("✅ 로드 밸런서가 수정되었습니다.\n")
	fmt.Printf("LB ID:  %s\n", lb.ID)
	fmt.Printf("Name:   %s\n", lb.Name)

	return nil
}

func runLBDelete(c *cobra.Command, args []string) error {
	lbID := args[0]

	client, err := loadbalancer.NewClient(GetProfile(), GetRegion(), GetDebug())
	if err != nil {
		return err
	}

	if err := client.DeleteLoadBalancer(lbID); err != nil {
		return err
	}

	fmt.Printf("✅ 로드 밸런서 %s가 삭제되었습니다.\n", lbID)
	return nil
}

func completeLBIDs(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	if len(args) != 0 {
		return nil, cobra.ShellCompDirectiveNoFileComp
	}
	client, err := loadbalancer.NewClient(GetProfile(), GetRegion(), GetDebug())
	if err != nil {
		return nil, cobra.ShellCompDirectiveNoFileComp
	}
	lbs, err := client.ListLoadBalancers()
	if err != nil {
		return nil, cobra.ShellCompDirectiveNoFileComp
	}
	var completions []string
	for _, lb := range lbs {
		completions = append(completions, fmt.Sprintf("%s\t%s", lb.ID, lb.Name))
	}
	return completions, cobra.ShellCompDirectiveNoFileComp
}
