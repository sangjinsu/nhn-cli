package gamebase

import (
	"nhncli/cmd"

	"github.com/spf13/cobra"
)

var GamebaseCmd = &cobra.Command{
	Use:   "gamebase",
	Short: "Gamebase 서비스 관리",
	Long: `Gamebase 서비스 관련 리소스를 관리합니다.

지원 리소스:
  - member: 회원 관리
  - ban: 이용 정지 관리
  - launching: 론칭 상태 조회
  - auth: 인증 토큰 검증`,
}

func init() {
	cmd.GetRootCmd().AddCommand(GamebaseCmd)
}
