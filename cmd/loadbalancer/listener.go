package loadbalancer

import (
	"fmt"

	"nhncli/internal/loadbalancer"
	"nhncli/internal/output"

	"github.com/spf13/cobra"
)

var (
	listenerLBID     string
	listenerProtocol string
	listenerPort     int
	listenerName     string
	listenerDesc     string
	listenerPoolID   string
)

var listenerCmd = &cobra.Command{
	Use:   "listener",
	Short: "리스너 관리",
}

var listenerListCmd = &cobra.Command{
	Use:   "list",
	Short: "리스너 목록 조회",
	RunE:  runListenerList,
}

var listenerDescribeCmd = &cobra.Command{
	Use:               "describe <listener-id>",
	Short:             "리스너 상세 조회",
	Args:              cobra.ExactArgs(1),
	ValidArgsFunction: completeListenerIDs,
	RunE:              runListenerDescribe,
}

var listenerCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "리스너 생성",
	Long: `로드 밸런서에 리스너를 생성합니다.

예시:
  nhn loadbalancer listener create \
    --loadbalancer-id <lb-id> \
    --protocol HTTP \
    --port 80`,
	RunE: runListenerCreate,
}

var listenerDeleteCmd = &cobra.Command{
	Use:               "delete <listener-id>",
	Short:             "리스너 삭제",
	Args:              cobra.ExactArgs(1),
	ValidArgsFunction: completeListenerIDs,
	RunE:              runListenerDelete,
}

func init() {
	LoadbalancerCmd.AddCommand(listenerCmd)
	listenerCmd.AddCommand(listenerListCmd)
	listenerCmd.AddCommand(listenerDescribeCmd)
	listenerCmd.AddCommand(listenerCreateCmd)
	listenerCmd.AddCommand(listenerDeleteCmd)

	listenerCreateCmd.Flags().StringVar(&listenerLBID, "loadbalancer-id", "", "로드 밸런서 ID (필수)")
	listenerCreateCmd.Flags().StringVar(&listenerProtocol, "protocol", "", "프로토콜 (TCP/HTTP/HTTPS, 필수)")
	listenerCreateCmd.Flags().IntVar(&listenerPort, "port", 0, "포트 번호 (필수)")
	listenerCreateCmd.Flags().StringVar(&listenerName, "name", "", "리스너 이름")
	listenerCreateCmd.Flags().StringVar(&listenerDesc, "description", "", "설명")
	listenerCreateCmd.Flags().StringVar(&listenerPoolID, "default-pool-id", "", "기본 풀 ID")
	listenerCreateCmd.MarkFlagRequired("loadbalancer-id")
	listenerCreateCmd.MarkFlagRequired("protocol")
	listenerCreateCmd.MarkFlagRequired("port")
}

func runListenerList(cmd *cobra.Command, args []string) error {
	client, err := loadbalancer.NewClient(GetProfile(), GetRegion(), GetDebug())
	if err != nil {
		return err
	}

	listeners, err := client.ListListeners()
	if err != nil {
		return err
	}

	if GetOutput() == "json" {
		return output.PrintJSON(listeners)
	}

	headers := []string{"ID", "NAME", "PROTOCOL", "PORT", "STATUS"}
	rows := make([][]string, len(listeners))
	for i, l := range listeners {
		rows[i] = []string{
			l.ID,
			l.Name,
			l.Protocol,
			fmt.Sprintf("%d", l.ProtocolPort),
			l.ProvisioningStatus,
		}
	}

	return output.Print(GetOutput(), &output.TableData{
		Headers: headers,
		Rows:    rows,
	}, listeners)
}

func runListenerDescribe(cmd *cobra.Command, args []string) error {
	listenerID := args[0]

	client, err := loadbalancer.NewClient(GetProfile(), GetRegion(), GetDebug())
	if err != nil {
		return err
	}

	l, err := client.GetListener(listenerID)
	if err != nil {
		return err
	}

	if GetOutput() == "json" {
		return output.PrintJSON(l)
	}

	fmt.Printf("Listener ID:        %s\n", l.ID)
	fmt.Printf("Name:               %s\n", l.Name)
	fmt.Printf("Description:        %s\n", l.Description)
	fmt.Printf("Protocol:           %s\n", l.Protocol)
	fmt.Printf("Port:               %d\n", l.ProtocolPort)
	fmt.Printf("Default Pool ID:    %s\n", l.DefaultPoolID)
	fmt.Printf("Provisioning Status:%s\n", l.ProvisioningStatus)
	fmt.Printf("Operating Status:   %s\n", l.OperatingStatus)

	return nil
}

func runListenerCreate(cmd *cobra.Command, args []string) error {
	client, err := loadbalancer.NewClient(GetProfile(), GetRegion(), GetDebug())
	if err != nil {
		return err
	}

	req := &loadbalancer.ListenerCreateRequest{
		Listener: loadbalancer.ListenerCreateBody{
			LoadbalancerID: listenerLBID,
			Protocol:       listenerProtocol,
			ProtocolPort:   listenerPort,
		},
	}

	if listenerName != "" {
		req.Listener.Name = listenerName
	}
	if listenerDesc != "" {
		req.Listener.Description = listenerDesc
	}
	if listenerPoolID != "" {
		req.Listener.DefaultPoolID = listenerPoolID
	}

	l, err := client.CreateListener(req)
	if err != nil {
		return err
	}

	if GetOutput() == "json" {
		return output.PrintJSON(l)
	}

	fmt.Printf("✅ 리스너가 생성되었습니다.\n")
	fmt.Printf("Listener ID: %s\n", l.ID)
	fmt.Printf("Name:        %s\n", l.Name)
	fmt.Printf("Protocol:    %s\n", l.Protocol)
	fmt.Printf("Port:        %d\n", l.ProtocolPort)
	fmt.Printf("Status:      %s\n", l.ProvisioningStatus)

	return nil
}

func runListenerDelete(cmd *cobra.Command, args []string) error {
	listenerID := args[0]

	client, err := loadbalancer.NewClient(GetProfile(), GetRegion(), GetDebug())
	if err != nil {
		return err
	}

	if err := client.DeleteListener(listenerID); err != nil {
		return err
	}

	fmt.Printf("✅ 리스너 %s가 삭제되었습니다.\n", listenerID)
	return nil
}

func completeListenerIDs(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	if len(args) != 0 {
		return nil, cobra.ShellCompDirectiveNoFileComp
	}
	client, err := loadbalancer.NewClient(GetProfile(), GetRegion(), GetDebug())
	if err != nil {
		return nil, cobra.ShellCompDirectiveNoFileComp
	}
	listeners, err := client.ListListeners()
	if err != nil {
		return nil, cobra.ShellCompDirectiveNoFileComp
	}
	var completions []string
	for _, l := range listeners {
		completions = append(completions, fmt.Sprintf("%s\t%s", l.ID, l.Name))
	}
	return completions, cobra.ShellCompDirectiveNoFileComp
}
