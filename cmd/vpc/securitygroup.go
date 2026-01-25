package vpc

import (
	"fmt"
	"strconv"
	"strings"

	"nhncli/internal/output"
	"nhncli/internal/vpc"

	"github.com/spf13/cobra"
)

var (
	sgName        string
	sgDescription string
	sgDirection   string
	sgProtocol    string
	sgPort        int
	sgPortRange   string
	sgRemoteIP    string
	sgRemoteGroup string
	sgEtherType   string
	sgRuleDesc    string
)

var securitygroupCmd = &cobra.Command{
	Use:     "securitygroup",
	Aliases: []string{"sg"},
	Short:   "보안 그룹 관리",
}

var sgListCmd = &cobra.Command{
	Use:   "list",
	Short: "보안 그룹 목록 조회",
	RunE:  runSGList,
}

var sgDescribeCmd = &cobra.Command{
	Use:   "describe <sg-id>",
	Short: "보안 그룹 상세 조회",
	Args:  cobra.ExactArgs(1),
	RunE:  runSGDescribe,
}

var sgCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "보안 그룹 생성",
	RunE:  runSGCreate,
}

var sgDeleteCmd = &cobra.Command{
	Use:   "delete <sg-id>",
	Short: "보안 그룹 삭제",
	Args:  cobra.ExactArgs(1),
	RunE:  runSGDelete,
}

var sgAddRuleCmd = &cobra.Command{
	Use:   "add-rule <sg-id>",
	Short: "보안 그룹 규칙 추가",
	Long: `보안 그룹에 규칙을 추가합니다.

예시:
  # SSH 허용
  nhn vpc securitygroup add-rule <sg-id> --direction ingress --protocol tcp --port 22 --remote-ip 0.0.0.0/0

  # HTTP/HTTPS 허용
  nhn vpc securitygroup add-rule <sg-id> --direction ingress --protocol tcp --port-range 80-443 --remote-ip 0.0.0.0/0

  # 모든 ICMP 허용
  nhn vpc securitygroup add-rule <sg-id> --direction ingress --protocol icmp --remote-ip 0.0.0.0/0`,
	Args: cobra.ExactArgs(1),
	RunE: runSGAddRule,
}

var sgDeleteRuleCmd = &cobra.Command{
	Use:   "delete-rule <rule-id>",
	Short: "보안 그룹 규칙 삭제",
	Args:  cobra.ExactArgs(1),
	RunE:  runSGDeleteRule,
}

func init() {
	VpcCmd.AddCommand(securitygroupCmd)
	securitygroupCmd.AddCommand(sgListCmd)
	securitygroupCmd.AddCommand(sgDescribeCmd)
	securitygroupCmd.AddCommand(sgCreateCmd)
	securitygroupCmd.AddCommand(sgDeleteCmd)
	securitygroupCmd.AddCommand(sgAddRuleCmd)
	securitygroupCmd.AddCommand(sgDeleteRuleCmd)

	sgCreateCmd.Flags().StringVar(&sgName, "name", "", "보안 그룹 이름 (필수)")
	sgCreateCmd.Flags().StringVar(&sgDescription, "description", "", "보안 그룹 설명")
	sgCreateCmd.MarkFlagRequired("name")

	sgAddRuleCmd.Flags().StringVar(&sgDirection, "direction", "ingress", "방향 (ingress/egress)")
	sgAddRuleCmd.Flags().StringVar(&sgProtocol, "protocol", "", "프로토콜 (tcp/udp/icmp)")
	sgAddRuleCmd.Flags().IntVar(&sgPort, "port", 0, "단일 포트")
	sgAddRuleCmd.Flags().StringVar(&sgPortRange, "port-range", "", "포트 범위 (예: 80-443)")
	sgAddRuleCmd.Flags().StringVar(&sgRemoteIP, "remote-ip", "", "원격 IP CIDR")
	sgAddRuleCmd.Flags().StringVar(&sgRemoteGroup, "remote-group", "", "원격 보안 그룹 ID")
	sgAddRuleCmd.Flags().StringVar(&sgEtherType, "ether-type", "IPv4", "Ether 타입 (IPv4/IPv6)")
	sgAddRuleCmd.Flags().StringVar(&sgRuleDesc, "description", "", "규칙 설명")
}

func runSGList(cmd *cobra.Command, args []string) error {
	client, err := vpc.NewClient(GetProfile(), GetRegion(), GetDebug())
	if err != nil {
		return err
	}

	sgs, err := client.ListSecurityGroups()
	if err != nil {
		return err
	}

	if GetOutput() == "json" {
		return output.PrintJSON(sgs)
	}

	headers := []string{"ID", "NAME", "DESCRIPTION", "RULES"}
	rows := make([][]string, len(sgs))
	for i, sg := range sgs {
		rows[i] = []string{sg.ID, sg.Name, sg.Description, fmt.Sprintf("%d", len(sg.Rules))}
	}

	return output.Print(GetOutput(), &output.TableData{
		Headers: headers,
		Rows:    rows,
	}, sgs)
}

func runSGDescribe(cmd *cobra.Command, args []string) error {
	sgID := args[0]

	client, err := vpc.NewClient(GetProfile(), GetRegion(), GetDebug())
	if err != nil {
		return err
	}

	sg, err := client.GetSecurityGroup(sgID)
	if err != nil {
		return err
	}

	if GetOutput() == "json" {
		return output.PrintJSON(sg)
	}

	fmt.Printf("Security Group ID:  %s\n", sg.ID)
	fmt.Printf("Name:               %s\n", sg.Name)
	fmt.Printf("Description:        %s\n", sg.Description)
	fmt.Printf("Tenant ID:          %s\n", sg.TenantID)
	fmt.Printf("\nRules (%d):\n", len(sg.Rules))

	for _, rule := range sg.Rules {
		portInfo := "all"
		if rule.PortRangeMin != nil && rule.PortRangeMax != nil {
			if *rule.PortRangeMin == *rule.PortRangeMax {
				portInfo = fmt.Sprintf("%d", *rule.PortRangeMin)
			} else {
				portInfo = fmt.Sprintf("%d-%d", *rule.PortRangeMin, *rule.PortRangeMax)
			}
		}

		remote := rule.RemoteIPPrefix
		if remote == "" && rule.RemoteGroupID != "" {
			remote = rule.RemoteGroupID
		}
		if remote == "" {
			remote = "any"
		}

		protocol := rule.Protocol
		if protocol == "" {
			protocol = "all"
		}

		fmt.Printf("  - %s %s/%s port=%s remote=%s\n",
			rule.Direction, rule.EtherType, protocol, portInfo, remote)
	}

	return nil
}

func runSGCreate(cmd *cobra.Command, args []string) error {
	client, err := vpc.NewClient(GetProfile(), GetRegion(), GetDebug())
	if err != nil {
		return err
	}

	sg, err := client.CreateSecurityGroup(sgName, sgDescription)
	if err != nil {
		return err
	}

	if GetOutput() == "json" {
		return output.PrintJSON(sg)
	}

	fmt.Printf("✅ 보안 그룹이 생성되었습니다.\n")
	fmt.Printf("Security Group ID:  %s\n", sg.ID)
	fmt.Printf("Name:               %s\n", sg.Name)

	return nil
}

func runSGDelete(cmd *cobra.Command, args []string) error {
	sgID := args[0]

	client, err := vpc.NewClient(GetProfile(), GetRegion(), GetDebug())
	if err != nil {
		return err
	}

	if err := client.DeleteSecurityGroup(sgID); err != nil {
		return err
	}

	fmt.Printf("✅ 보안 그룹 %s가 삭제되었습니다.\n", sgID)
	return nil
}

func runSGAddRule(cmd *cobra.Command, args []string) error {
	sgID := args[0]

	var portMin, portMax *int

	if sgPort > 0 {
		portMin = &sgPort
		portMax = &sgPort
	} else if sgPortRange != "" {
		parts := strings.Split(sgPortRange, "-")
		if len(parts) == 2 {
			min, _ := strconv.Atoi(parts[0])
			max, _ := strconv.Atoi(parts[1])
			portMin = &min
			portMax = &max
		}
	}

	client, err := vpc.NewClient(GetProfile(), GetRegion(), GetDebug())
	if err != nil {
		return err
	}

	rule, err := client.AddSecurityGroupRule(
		sgID, sgDirection, sgProtocol, sgEtherType,
		portMin, portMax, sgRemoteIP, sgRemoteGroup, sgRuleDesc,
	)
	if err != nil {
		return err
	}

	if GetOutput() == "json" {
		return output.PrintJSON(rule)
	}

	fmt.Printf("✅ 보안 그룹 규칙이 추가되었습니다.\n")
	fmt.Printf("Rule ID:  %s\n", rule.ID)

	return nil
}

func runSGDeleteRule(cmd *cobra.Command, args []string) error {
	ruleID := args[0]

	client, err := vpc.NewClient(GetProfile(), GetRegion(), GetDebug())
	if err != nil {
		return err
	}

	if err := client.DeleteSecurityGroupRule(ruleID); err != nil {
		return err
	}

	fmt.Printf("✅ 보안 그룹 규칙 %s가 삭제되었습니다.\n", ruleID)
	return nil
}
