package dns

import (
	"fmt"

	"nhncli/internal/dns"
	"nhncli/internal/output"

	"github.com/spf13/cobra"
)

var (
	zoneCreateName        string
	zoneCreateDescription string
	zoneUpdateDescription string
)

var zoneCmd = &cobra.Command{
	Use:   "zone",
	Short: "DNS Zone 관리",
}

var zoneListCmd = &cobra.Command{
	Use:   "list",
	Short: "Zone 목록 조회",
	RunE:  runZoneList,
}

var zoneDescribeCmd = &cobra.Command{
	Use:               "describe <zone-id>",
	Short:             "Zone 상세 조회",
	Args:              cobra.ExactArgs(1),
	ValidArgsFunction: completeZoneIDs,
	RunE:              runZoneDescribe,
}

var zoneCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Zone 생성",
	RunE:  runZoneCreate,
}

var zoneUpdateCmd = &cobra.Command{
	Use:               "update <zone-id>",
	Short:             "Zone 수정",
	Args:              cobra.ExactArgs(1),
	ValidArgsFunction: completeZoneIDs,
	RunE:              runZoneUpdate,
}

var zoneDeleteCmd = &cobra.Command{
	Use:               "delete <zone-id>",
	Short:             "Zone 삭제",
	Args:              cobra.ExactArgs(1),
	ValidArgsFunction: completeZoneIDs,
	RunE:              runZoneDelete,
}

func init() {
	DnsCmd.AddCommand(zoneCmd)
	zoneCmd.AddCommand(zoneListCmd)
	zoneCmd.AddCommand(zoneDescribeCmd)
	zoneCmd.AddCommand(zoneCreateCmd)
	zoneCmd.AddCommand(zoneUpdateCmd)
	zoneCmd.AddCommand(zoneDeleteCmd)

	zoneCreateCmd.Flags().StringVar(&zoneCreateName, "name", "", "Zone 이름 (FQDN, 예: example.com.)")
	zoneCreateCmd.Flags().StringVar(&zoneCreateDescription, "description", "", "Zone 설명")
	zoneCreateCmd.MarkFlagRequired("name")

	zoneUpdateCmd.Flags().StringVar(&zoneUpdateDescription, "description", "", "Zone 설명")
}

func runZoneList(cmd *cobra.Command, args []string) error {
	opts := dns.ClientOption{AppKey: GetAppKey(cmd)}
	c, err := dns.NewClient(GetProfile(), GetDebug(), opts)
	if err != nil {
		return err
	}

	zones, err := c.ListZones()
	if err != nil {
		return err
	}

	if GetOutput() == "json" {
		return output.PrintJSON(zones)
	}

	headers := []string{"ZONE ID", "NAME", "STATUS", "RECORDS", "DESCRIPTION"}
	rows := make([][]string, len(zones))
	for i, z := range zones {
		rows[i] = []string{
			z.ZoneID,
			z.ZoneName,
			z.ZoneStatus,
			fmt.Sprintf("%d", z.RecordsetCount),
			z.Description,
		}
	}

	return output.Print(GetOutput(), &output.TableData{
		Headers: headers,
		Rows:    rows,
	}, zones)
}

func runZoneDescribe(cmd *cobra.Command, args []string) error {
	opts := dns.ClientOption{AppKey: GetAppKey(cmd)}
	c, err := dns.NewClient(GetProfile(), GetDebug(), opts)
	if err != nil {
		return err
	}

	zone, err := c.GetZone(args[0])
	if err != nil {
		return err
	}

	if GetOutput() == "json" {
		return output.PrintJSON(zone)
	}

	fmt.Printf("Zone ID:       %s\n", zone.ZoneID)
	fmt.Printf("Name:          %s\n", zone.ZoneName)
	fmt.Printf("Status:        %s\n", zone.ZoneStatus)
	fmt.Printf("Record Sets:   %d\n", zone.RecordsetCount)
	fmt.Printf("Description:   %s\n", zone.Description)
	fmt.Printf("Created At:    %s\n", zone.CreatedAt)
	fmt.Printf("Updated At:    %s\n", zone.UpdatedAt)

	return nil
}

func runZoneCreate(cmd *cobra.Command, args []string) error {
	opts := dns.ClientOption{AppKey: GetAppKey(cmd)}
	c, err := dns.NewClient(GetProfile(), GetDebug(), opts)
	if err != nil {
		return err
	}

	req := &dns.ZoneCreateRequest{
		Zone: dns.ZoneCreateBody{
			ZoneName:    zoneCreateName,
			Description: zoneCreateDescription,
		},
	}

	zone, err := c.CreateZone(req)
	if err != nil {
		return err
	}

	fmt.Printf("Zone '%s'이(가) 생성되었습니다. (ID: %s)\n", zone.ZoneName, zone.ZoneID)
	return nil
}

func runZoneUpdate(cmd *cobra.Command, args []string) error {
	opts := dns.ClientOption{AppKey: GetAppKey(cmd)}
	c, err := dns.NewClient(GetProfile(), GetDebug(), opts)
	if err != nil {
		return err
	}

	req := &dns.ZoneUpdateRequest{
		Zone: dns.ZoneUpdateBody{
			Description: zoneUpdateDescription,
		},
	}

	zone, err := c.UpdateZone(args[0], req)
	if err != nil {
		return err
	}

	fmt.Printf("Zone '%s'이(가) 수정되었습니다.\n", zone.ZoneName)
	return nil
}

func runZoneDelete(cmd *cobra.Command, args []string) error {
	opts := dns.ClientOption{AppKey: GetAppKey(cmd)}
	c, err := dns.NewClient(GetProfile(), GetDebug(), opts)
	if err != nil {
		return err
	}

	if err := c.DeleteZone(args[0]); err != nil {
		return err
	}

	fmt.Printf("Zone '%s' 삭제가 요청되었습니다.\n", args[0])
	return nil
}

func completeZoneIDs(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	if len(args) != 0 {
		return nil, cobra.ShellCompDirectiveNoFileComp
	}
	opts := dns.ClientOption{AppKey: GetAppKey(cmd)}
	c, err := dns.NewClient(GetProfile(), GetDebug(), opts)
	if err != nil {
		return nil, cobra.ShellCompDirectiveNoFileComp
	}
	zones, err := c.ListZones()
	if err != nil {
		return nil, cobra.ShellCompDirectiveNoFileComp
	}
	var completions []string
	for _, z := range zones {
		completions = append(completions, fmt.Sprintf("%s\t%s", z.ZoneID, z.ZoneName))
	}
	return completions, cobra.ShellCompDirectiveNoFileComp
}
