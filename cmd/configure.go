package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"nhncli/internal/auth"
	"nhncli/internal/config"

	"github.com/spf13/cobra"
)

var configureCmd = &cobra.Command{
	Use:   "configure",
	Short: "NHN Cloud 인증 정보 설정",
	Long: `NHN Cloud CLI 인증 정보를 설정합니다.

Identity 인증:
  Tenant ID, Username, Password를 사용합니다.
  VPC, Compute 등 OpenStack 기반 API에 필요합니다.

OAuth 인증:
  User Access Key ID와 Secret Access Key를 사용합니다.
  NHN Cloud 고유 API에서 사용됩니다.`,
	RunE: runConfigure,
}

var configureListCmd = &cobra.Command{
	Use:   "list",
	Short: "설정된 프로필 목록 조회",
	RunE:  runConfigureList,
}

func init() {
	rootCmd.AddCommand(configureCmd)
	configureCmd.AddCommand(configureListCmd)
}

func runConfigure(cmd *cobra.Command, args []string) error {
	reader := bufio.NewReader(os.Stdin)

	profileName := readInput(reader, fmt.Sprintf("프로필 이름 [%s]: ", profile), profile)

	fmt.Println("\n=== NHN Cloud 인증 설정 ===")
	fmt.Println("")
	fmt.Println("📌 VPC, Compute 등 OpenStack 기반 API 사용을 위해 Identity 인증 정보가 필요합니다.")

	// Identity 인증 (필수)
	fmt.Println("\n--- Identity 인증 (필수) ---")
	fmt.Println("")
	fmt.Println("📌 Tenant ID 확인 방법:")
	fmt.Println("   1. NHN Cloud 콘솔 (https://console.nhncloud.com) 로그인")
	fmt.Println("   2. 프로젝트 선택 후 'Compute > Instance' 메뉴 이동")
	fmt.Println("   3. 'API 엔드포인트 설정' 버튼 클릭")
	fmt.Println("   4. Tenant ID 확인")
	fmt.Println("")
	fmt.Println("📌 API Password 설정 방법:")
	fmt.Println("   위 'API 엔드포인트 설정' 화면에서 'API 비밀번호 설정' 클릭")
	fmt.Println("")

	tenantID := readInput(reader, "Tenant ID: ", "")
	username := readInput(reader, "Username (이메일 주소): ", "")
	password := readSecretInput(reader, "API Password: ")

	profileConfig := &config.ProfileConfig{
		TenantID: tenantID,
		Username: username,
		Password: password,
	}

	// OAuth 인증 (필수)
	fmt.Println("\n--- OAuth 인증 (필수) ---")
	fmt.Println("")
	fmt.Println("📌 User Access Key ID 발급 방법:")
	fmt.Println("   1. NHN Cloud 콘솔 (https://console.nhncloud.com) 로그인")
	fmt.Println("   2. 오른쪽 상단의 이메일 주소 클릭")
	fmt.Println("   3. 'API 보안 설정' 메뉴 선택")
	fmt.Println("   4. 'User Access Key ID 생성' 버튼 클릭")
	fmt.Println("")

	userAccessKeyID := readInput(reader, "User Access Key ID: ", "")
	secretAccessKey := readSecretInput(reader, "Secret Access Key: ")
	profileConfig.UserAccessKeyID = userAccessKeyID
	profileConfig.SecretAccessKey = secretAccessKey

	// 리전 설정
	fmt.Println("\n=== 리전 설정 ===")
	fmt.Println("")
	fmt.Println("사용 가능한 리전:")
	fmt.Println("   KR1 - 한국 (판교) 리전")
	fmt.Println("   KR2 - 한국 (평촌) 리전")
	fmt.Println("   JP1 - 일본 (도쿄) 리전")
	fmt.Println("")
	profileConfig.Region = readInput(reader, "기본 리전 [KR1]: ", "KR1")

	if err := profileConfig.Validate(); err != nil {
		return fmt.Errorf("설정 검증 실패: %w", err)
	}

	cfg, err := config.Load()
	if err != nil {
		return err
	}

	cfg.SetProfile(profileName, profileConfig)

	if err := cfg.Save(); err != nil {
		return err
	}

	fmt.Printf("\n✅ 프로필 '%s' 설정이 저장되었습니다.\n", profileName)

	// Identity 토큰 발급으로 인증 정보 검증
	fmt.Println("\n🔐 Identity 인증 정보 검증 중...")
	token, tenantIDResp, err := auth.GetAuthenticatedToken(profileName, profileConfig, false)
	if err != nil {
		fmt.Printf("⚠️  인증 실패: %v\n", err)
		fmt.Println("   인증 정보를 다시 확인해주세요.")
		return nil // 설정은 저장되었으므로 에러 반환하지 않음
	}

	fmt.Println("✅ Identity 인증 성공!")
	if tenantIDResp != "" {
		fmt.Printf("   Tenant ID: %s\n", tenantIDResp)
	}
	fmt.Printf("   토큰이 캐시되었습니다. (유효기간: 12시간)\n")

	fmt.Println("   OAuth 인증 정보도 저장되었습니다.")

	_ = token // 사용하지 않는 변수 경고 방지

	return nil
}

func runConfigureList(cmd *cobra.Command, args []string) error {
	cfg, err := config.Load()
	if err != nil {
		return err
	}

	profiles := cfg.ListProfiles()
	if len(profiles) == 0 {
		fmt.Println("설정된 프로필이 없습니다. 'nhn configure'로 프로필을 추가하세요.")
		return nil
	}

	fmt.Println("=== 프로필 목록 ===")
	for _, name := range profiles {
		p, _ := cfg.GetProfile(name)
		fmt.Printf("  %s:\n", name)
		fmt.Printf("    인증 방식: %s\n", p.GetAuthTypeDisplay())
		fmt.Printf("    자격 증명: %s\n", p.GetMaskedCredentials())
		fmt.Printf("    리전: %s\n", p.Region)
	}

	return nil
}

func readInput(reader *bufio.Reader, prompt, defaultVal string) string {
	fmt.Print(prompt)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	if input == "" {
		return defaultVal
	}
	return input
}

func readSecretInput(reader *bufio.Reader, prompt string) string {
	fmt.Print(prompt)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}
