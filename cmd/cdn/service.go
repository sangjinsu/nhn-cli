package cdn

import (
	"encoding/json"
	"fmt"
	"os"
	"text/tabwriter"

	"nhncli/cmd"
	"nhncli/internal/cdn"

	"github.com/spf13/cobra"
)

var serviceCmd = &cobra.Command{
	Use:   "service",
	Short: "CDN 서비스 관리",
}

var serviceListCmd = &cobra.Command{
	Use:   "list",
	Short: "CDN 서비스 목록 조회",
	RunE:  runServiceList,
}

var serviceCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "CDN 서비스 생성",
	RunE:  runServiceCreate,
}

var serviceUpdateCmd = &cobra.Command{
	Use:   "update [domain]",
	Short: "CDN 서비스 수정",
	Args:  cobra.ExactArgs(1),
	RunE:  runServiceUpdate,
}

var serviceDeleteCmd = &cobra.Command{
	Use:   "delete [domain]",
	Short: "CDN 서비스 삭제",
	Args:  cobra.ExactArgs(1),
	RunE:  runServiceDelete,
}

var (
	originURL   string
	domainAlias string
	description string
)

func init() {
	CDNCmd.AddCommand(serviceCmd)
	serviceCmd.AddCommand(serviceListCmd)
	serviceCmd.AddCommand(serviceCreateCmd)
	serviceCmd.AddCommand(serviceUpdateCmd)
	serviceCmd.AddCommand(serviceDeleteCmd)

	serviceCreateCmd.Flags().StringVar(&originURL, "origin-url", "", "원본 서버 URL (필수)")
	serviceCreateCmd.Flags().StringVar(&domainAlias, "domain-alias", "", "도메인 별칭")
	serviceCreateCmd.Flags().StringVar(&description, "description", "", "설명")
	serviceCreateCmd.MarkFlagRequired("origin-url")

	serviceUpdateCmd.Flags().StringVar(&originURL, "origin-url", "", "원본 서버 URL")
	serviceUpdateCmd.Flags().StringVar(&domainAlias, "domain-alias", "", "도메인 별칭")
	serviceUpdateCmd.Flags().StringVar(&description, "description", "", "설명")
}

func runServiceList(c *cobra.Command, args []string) error {
	cdnClient, err := cdn.NewClient(cmd.GetProfile(), cmd.GetDebug())
	if err != nil {
		return err
	}

	services, err := cdnClient.ListServices()
	if err != nil {
		return err
	}

	if cmd.GetOutput() == "json" {
		enc := json.NewEncoder(os.Stdout)
		enc.SetIndent("", "  ")
		return enc.Encode(services)
	}

	if len(services) == 0 {
		fmt.Println("CDN 서비스가 없습니다.")
		return nil
	}

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	fmt.Fprintln(w, "도메인\t상태\t원본 URL\t설명")
	fmt.Fprintln(w, "------\t----\t-------\t----")
	for _, s := range services {
		fmt.Fprintf(w, "%s\t%s\t%s\t%s\n", s.Domain, s.Status, s.OriginURL, s.Description)
	}
	return w.Flush()
}

func runServiceCreate(c *cobra.Command, args []string) error {
	cdnClient, err := cdn.NewClient(cmd.GetProfile(), cmd.GetDebug())
	if err != nil {
		return err
	}

	req := &cdn.ServiceCreateRequest{
		OriginURL:   originURL,
		DomainAlias: domainAlias,
		Description: description,
	}

	svc, err := cdnClient.CreateService(req)
	if err != nil {
		return err
	}

	fmt.Printf("CDN 서비스가 생성되었습니다. (도메인: %s)\n", svc.Domain)
	return nil
}

func runServiceUpdate(c *cobra.Command, args []string) error {
	domain := args[0]

	cdnClient, err := cdn.NewClient(cmd.GetProfile(), cmd.GetDebug())
	if err != nil {
		return err
	}

	req := &cdn.ServiceCreateRequest{
		OriginURL:   originURL,
		DomainAlias: domainAlias,
		Description: description,
	}

	svc, err := cdnClient.UpdateService(domain, req)
	if err != nil {
		return err
	}

	fmt.Printf("CDN 서비스가 수정되었습니다. (도메인: %s)\n", svc.Domain)
	return nil
}

func runServiceDelete(c *cobra.Command, args []string) error {
	domain := args[0]

	cdnClient, err := cdn.NewClient(cmd.GetProfile(), cmd.GetDebug())
	if err != nil {
		return err
	}

	if err := cdnClient.DeleteService(domain); err != nil {
		return err
	}

	fmt.Printf("CDN 서비스가 삭제되었습니다. (도메인: %s)\n", domain)
	return nil
}
