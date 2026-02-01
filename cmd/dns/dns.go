package dns

import (
	"nhncli/cmd"

	"github.com/spf13/cobra"
)

var DnsCmd = &cobra.Command{
	Use:   "dns",
	Short: "DNS Plus 서비스 관리",
	Long: `DNS Plus 서비스 관련 리소스를 관리합니다.

지원 리소스:
  - zone: DNS Zone 관리
  - recordset: Record Set 관리`,
}

func init() {
	cmd.GetRootCmd().AddCommand(DnsCmd)
	DnsCmd.PersistentFlags().String("app-key", "", "DNS AppKey (프로필 설정 오버라이드)")
}

func GetProfile() string {
	return cmd.GetProfile()
}

func GetOutput() string {
	return cmd.GetOutput()
}

func GetDebug() bool {
	return cmd.GetDebug()
}

func GetAppKey(c *cobra.Command) string {
	v, _ := c.Flags().GetString("app-key")
	return v
}
