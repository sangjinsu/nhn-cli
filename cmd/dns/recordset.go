package dns

import (
	"fmt"
	"strings"

	"nhncli/internal/dns"
	"nhncli/internal/output"

	"github.com/spf13/cobra"
)

var (
	rsZoneID string
	rsName   string
	rsType   string
	rsTTL    int
	rsData   []string
)

var recordsetCmd = &cobra.Command{
	Use:   "recordset",
	Short: "Record Set 관리",
}

var recordsetListCmd = &cobra.Command{
	Use:   "list",
	Short: "Record Set 목록 조회",
	RunE:  runRecordSetList,
}

var recordsetDescribeCmd = &cobra.Command{
	Use:               "describe <recordset-id>",
	Short:             "Record Set 상세 조회",
	Args:              cobra.ExactArgs(1),
	ValidArgsFunction: completeRecordSetIDs,
	RunE:              runRecordSetDescribe,
}

var recordsetCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Record Set 생성",
	RunE:  runRecordSetCreate,
}

var recordsetUpdateCmd = &cobra.Command{
	Use:               "update <recordset-id>",
	Short:             "Record Set 수정",
	Args:              cobra.ExactArgs(1),
	ValidArgsFunction: completeRecordSetIDs,
	RunE:              runRecordSetUpdate,
}

var recordsetDeleteCmd = &cobra.Command{
	Use:               "delete <recordset-id>",
	Short:             "Record Set 삭제",
	Args:              cobra.ExactArgs(1),
	ValidArgsFunction: completeRecordSetIDs,
	RunE:              runRecordSetDelete,
}

func init() {
	DnsCmd.AddCommand(recordsetCmd)
	recordsetCmd.AddCommand(recordsetListCmd)
	recordsetCmd.AddCommand(recordsetDescribeCmd)
	recordsetCmd.AddCommand(recordsetCreateCmd)
	recordsetCmd.AddCommand(recordsetUpdateCmd)
	recordsetCmd.AddCommand(recordsetDeleteCmd)

	// --zone-id required on all subcommands
	for _, c := range []*cobra.Command{recordsetListCmd, recordsetDescribeCmd, recordsetCreateCmd, recordsetUpdateCmd, recordsetDeleteCmd} {
		c.Flags().StringVar(&rsZoneID, "zone-id", "", "Zone ID (필수)")
		c.MarkFlagRequired("zone-id")
		c.RegisterFlagCompletionFunc("zone-id", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			return completeZoneIDsForFlag(cmd, args, toComplete)
		})
	}

	recordsetCreateCmd.Flags().StringVar(&rsName, "name", "", "Record Set 이름")
	recordsetCreateCmd.Flags().StringVar(&rsType, "type", "", "Record Set 타입 (A, AAAA, CNAME, MX, TXT, etc.)")
	recordsetCreateCmd.Flags().IntVar(&rsTTL, "ttl", 300, "TTL (초)")
	recordsetCreateCmd.Flags().StringSliceVar(&rsData, "data", nil, "레코드 데이터 (반복 가능)")
	recordsetCreateCmd.MarkFlagRequired("name")
	recordsetCreateCmd.MarkFlagRequired("type")
	recordsetCreateCmd.MarkFlagRequired("data")

	recordsetUpdateCmd.Flags().StringVar(&rsType, "type", "", "Record Set 타입")
	recordsetUpdateCmd.Flags().IntVar(&rsTTL, "ttl", 0, "TTL (초)")
	recordsetUpdateCmd.Flags().StringSliceVar(&rsData, "data", nil, "레코드 데이터 (반복 가능)")
}

func runRecordSetList(cmd *cobra.Command, args []string) error {
	opts := dns.ClientOption{AppKey: GetAppKey(cmd)}
	c, err := dns.NewClient(GetProfile(), GetDebug(), opts)
	if err != nil {
		return err
	}

	recordsets, err := c.ListRecordSets(rsZoneID)
	if err != nil {
		return err
	}

	if GetOutput() == "json" {
		return output.PrintJSON(recordsets)
	}

	headers := []string{"RECORDSET ID", "NAME", "TYPE", "TTL", "STATUS", "RECORDS"}
	rows := make([][]string, len(recordsets))
	for i, rs := range recordsets {
		var records []string
		for _, r := range rs.RecordList {
			records = append(records, r.RecordContent)
		}
		rows[i] = []string{
			rs.RecordsetID,
			rs.RecordsetName,
			rs.RecordsetType,
			fmt.Sprintf("%d", rs.RecordsetTTL),
			rs.RecordsetStatus,
			strings.Join(records, ", "),
		}
	}

	return output.Print(GetOutput(), &output.TableData{
		Headers: headers,
		Rows:    rows,
	}, recordsets)
}

func runRecordSetDescribe(cmd *cobra.Command, args []string) error {
	opts := dns.ClientOption{AppKey: GetAppKey(cmd)}
	c, err := dns.NewClient(GetProfile(), GetDebug(), opts)
	if err != nil {
		return err
	}

	rs, err := c.GetRecordSet(rsZoneID, args[0])
	if err != nil {
		return err
	}

	if GetOutput() == "json" {
		return output.PrintJSON(rs)
	}

	fmt.Printf("Recordset ID:  %s\n", rs.RecordsetID)
	fmt.Printf("Name:          %s\n", rs.RecordsetName)
	fmt.Printf("Type:          %s\n", rs.RecordsetType)
	fmt.Printf("TTL:           %d\n", rs.RecordsetTTL)
	fmt.Printf("Status:        %s\n", rs.RecordsetStatus)
	fmt.Printf("Created At:    %s\n", rs.CreatedAt)
	fmt.Printf("Updated At:    %s\n", rs.UpdatedAt)
	fmt.Println("Records:")
	for _, r := range rs.RecordList {
		disabled := ""
		if r.RecordDisabled {
			disabled = " (disabled)"
		}
		fmt.Printf("  - %s%s\n", r.RecordContent, disabled)
	}

	return nil
}

func runRecordSetCreate(cmd *cobra.Command, args []string) error {
	opts := dns.ClientOption{AppKey: GetAppKey(cmd)}
	c, err := dns.NewClient(GetProfile(), GetDebug(), opts)
	if err != nil {
		return err
	}

	records := make([]dns.Record, len(rsData))
	for i, d := range rsData {
		records[i] = dns.Record{RecordContent: d}
	}

	req := &dns.RecordSetCreateRequest{
		Recordset: dns.RecordSetCreateBody{
			RecordsetName: rsName,
			RecordsetType: rsType,
			RecordsetTTL:  rsTTL,
			RecordList:    records,
		},
	}

	rs, err := c.CreateRecordSet(rsZoneID, req)
	if err != nil {
		return err
	}

	fmt.Printf("Record Set '%s' (%s)이(가) 생성되었습니다. (ID: %s)\n", rs.RecordsetName, rs.RecordsetType, rs.RecordsetID)
	return nil
}

func runRecordSetUpdate(cmd *cobra.Command, args []string) error {
	opts := dns.ClientOption{AppKey: GetAppKey(cmd)}
	c, err := dns.NewClient(GetProfile(), GetDebug(), opts)
	if err != nil {
		return err
	}

	body := dns.RecordSetUpdateBody{
		RecordsetType: rsType,
		RecordsetTTL:  rsTTL,
	}

	if len(rsData) > 0 {
		records := make([]dns.Record, len(rsData))
		for i, d := range rsData {
			records[i] = dns.Record{RecordContent: d}
		}
		body.RecordList = records
	}

	req := &dns.RecordSetUpdateRequest{Recordset: body}

	rs, err := c.UpdateRecordSet(rsZoneID, args[0], req)
	if err != nil {
		return err
	}

	fmt.Printf("Record Set '%s'이(가) 수정되었습니다.\n", rs.RecordsetName)
	return nil
}

func runRecordSetDelete(cmd *cobra.Command, args []string) error {
	opts := dns.ClientOption{AppKey: GetAppKey(cmd)}
	c, err := dns.NewClient(GetProfile(), GetDebug(), opts)
	if err != nil {
		return err
	}

	if err := c.DeleteRecordSet(rsZoneID, args[0]); err != nil {
		return err
	}

	fmt.Printf("Record Set '%s' 삭제가 요청되었습니다.\n", args[0])
	return nil
}

func completeZoneIDsForFlag(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
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

func completeRecordSetIDs(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	if len(args) != 0 {
		return nil, cobra.ShellCompDirectiveNoFileComp
	}
	if rsZoneID == "" {
		return nil, cobra.ShellCompDirectiveNoFileComp
	}
	opts := dns.ClientOption{AppKey: GetAppKey(cmd)}
	c, err := dns.NewClient(GetProfile(), GetDebug(), opts)
	if err != nil {
		return nil, cobra.ShellCompDirectiveNoFileComp
	}
	recordsets, err := c.ListRecordSets(rsZoneID)
	if err != nil {
		return nil, cobra.ShellCompDirectiveNoFileComp
	}
	var completions []string
	for _, rs := range recordsets {
		completions = append(completions, fmt.Sprintf("%s\t%s (%s)", rs.RecordsetID, rs.RecordsetName, rs.RecordsetType))
	}
	return completions, cobra.ShellCompDirectiveNoFileComp
}
