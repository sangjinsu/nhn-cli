package vpc

import (
	"nhncli/cmd"

	"github.com/spf13/cobra"
)

var VpcCmd = &cobra.Command{
	Use:   "vpc",
	Short: "VPC 관리",
	Long: `Virtual Private Cloud (VPC) 관련 리소스를 관리합니다.

지원 리소스:
  - vpc: VPC 관리
  - subnet: 서브넷 관리
  - securitygroup: 보안 그룹 관리
  - floatingip: 플로팅 IP 관리
  - routingtable: 라우팅 테이블 조회
  - port: 네트워크 인터페이스 관리`,
}

func init() {
	cmd.GetRootCmd().AddCommand(VpcCmd)
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
