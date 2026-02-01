package cdn

import (
	"nhncli/cmd"

	"github.com/spf13/cobra"
)

var CDNCmd = &cobra.Command{
	Use:   "cdn",
	Short: "CDN 서비스 관리",
	Long: `CDN 서비스 관련 리소스를 관리합니다.

지원 리소스:
  - service: CDN 서비스 관리
  - purge: 캐시 퍼지
  - auth-token: 인증 토큰 생성`,
}

func init() {
	cmd.GetRootCmd().AddCommand(CDNCmd)
}
