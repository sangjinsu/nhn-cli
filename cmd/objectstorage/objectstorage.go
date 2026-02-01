package objectstorage

import (
	"nhncli/cmd"

	"github.com/spf13/cobra"
)

var ObjectstorageCmd = &cobra.Command{
	Use:     "objectstorage",
	Aliases: []string{"os"},
	Short:   "Object Storage 서비스 관리",
	Long: `Object Storage 서비스 관련 리소스를 관리합니다.

지원 리소스:
  - container: 컨테이너 관리
  - object: 오브젝트 관리`,
}

func init() {
	cmd.GetRootCmd().AddCommand(ObjectstorageCmd)
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
