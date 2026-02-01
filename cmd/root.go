package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"nhncli/internal/config"
)

var (
	profile   string
	region    string
	output    string
	debug     bool
	Version   = "0.1.0"
	BuildDate = "unknown"
)

var rootCmd = &cobra.Command{
	Use:   "nhn",
	Short: "NHN Cloud CLI - AWS CLI 스타일의 NHN Cloud 명령줄 인터페이스",
	Long: `NHN Cloud CLI는 NHN Cloud 서비스를 명령줄에서 관리할 수 있는 도구입니다.
AWS CLI와 유사한 사용법을 제공하여 친숙하게 사용할 수 있습니다.

지원 서비스:
  - VPC: Virtual Private Cloud 관리
  - Compute: 인스턴스 관리
  - Block Storage: 블록 스토리지 관리
  - Load Balancer: 로드 밸런서 관리
  - Object Storage: 오브젝트 스토리지 관리
  - DNS Plus: DNS Zone 및 Record Set 관리
  - Pipeline: 파이프라인 실행 관리
  - Deploy: 배포 실행 관리
  - CDN: CDN 서비스 관리
  - AppGuard: 앱 보안 탐지 현황
  - Gamebase: 게임 회원/정지/론칭 관리

사용 예시:
  nhn configure                    # 인증 정보 설정
  nhn vpc list                     # VPC 목록 조회
  nhn compute instance list        # 인스턴스 목록 조회`,
	SilenceUsage:  true,
	SilenceErrors: true,
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.PersistentFlags().StringVar(&profile, "profile", "default", "사용할 프로필")
	rootCmd.PersistentFlags().StringVar(&region, "region", "", "리전 지정 (프로필 설정 오버라이드)")
	rootCmd.PersistentFlags().StringVar(&output, "output", "table", "출력 형식 (table, json)")
	rootCmd.PersistentFlags().BoolVar(&debug, "debug", false, "디버그 모드")

	rootCmd.RegisterFlagCompletionFunc("profile", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		cfg, err := config.Load()
		if err != nil {
			return nil, cobra.ShellCompDirectiveNoFileComp
		}
		return cfg.ListProfiles(), cobra.ShellCompDirectiveNoFileComp
	})

	rootCmd.RegisterFlagCompletionFunc("region", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		regions := []string{
			"KR1\t한국 (판교)",
			"KR2\t한국 (평촌)",
			"JP1\t일본 (도쿄)",
		}
		return regions, cobra.ShellCompDirectiveNoFileComp
	})

	rootCmd.RegisterFlagCompletionFunc("output", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return []string{"table\t테이블 형식", "json\tJSON 형식"}, cobra.ShellCompDirectiveNoFileComp
	})
}

func GetProfile() string {
	return profile
}

func GetRegion() string {
	return region
}

func GetOutput() string {
	return output
}

func GetDebug() bool {
	return debug
}

func ExitWithError(format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, "Error: "+format+"\n", args...)
	os.Exit(1)
}

func GetRootCmd() *cobra.Command {
	return rootCmd
}
