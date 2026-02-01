package blockstorage

import (
	"nhncli/cmd"

	"github.com/spf13/cobra"
)

var BlockstorageCmd = &cobra.Command{
	Use:   "blockstorage",
	Short: "Block Storage 서비스 관리",
	Long: `Block Storage 서비스 관련 리소스를 관리합니다.

지원 리소스:
  - volume: 볼륨 관리
  - snapshot: 스냅샷 관리
  - type: 볼륨 타입 조회`,
}

func init() {
	cmd.GetRootCmd().AddCommand(BlockstorageCmd)
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
