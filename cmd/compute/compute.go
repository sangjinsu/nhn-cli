package compute

import (
	"nhncli/cmd"

	"github.com/spf13/cobra"
)

var ComputeCmd = &cobra.Command{
	Use:   "compute",
	Short: "Compute 서비스 관리",
	Long: `Compute 서비스 관련 리소스를 관리합니다.

지원 리소스:
  - instance: 인스턴스 관리
  - flavor: 인스턴스 타입 조회
  - image: 이미지 조회
  - keypair: 키페어 관리
  - az: 가용성 영역 조회`,
}

func init() {
	cmd.GetRootCmd().AddCommand(ComputeCmd)
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
