package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"nhncli/internal/config"

	"github.com/spf13/cobra"
)

var configureCmd = &cobra.Command{
	Use:   "configure",
	Short: "NHN Cloud 인증 정보 설정",
	Long: `NHN Cloud CLI 인증 정보를 설정합니다.

OAuth 인증 (권장):
  User Access Key ID와 Secret Access Key를 사용합니다.
  NHN Cloud 콘솔 > API 보안 설정에서 발급받을 수 있습니다.

Identity 인증:
  Tenant ID, Username, Password를 사용합니다.
  기존 OpenStack API 방식의 인증입니다.`,
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

	fmt.Println("\n=== 인증 방식 선택 ===")
	fmt.Println("1. OAuth 인증 (User Access Key ID) - 권장")
	fmt.Println("2. Identity 인증 (Tenant ID + Username)")
	authChoice := readInput(reader, "선택 [1]: ", "1")

	var profileConfig *config.ProfileConfig

	switch authChoice {
	case "1":
		profileConfig = configureOAuth(reader)
	case "2":
		profileConfig = configureIdentity(reader)
	default:
		return fmt.Errorf("잘못된 선택입니다")
	}

	fmt.Println("\n=== 리전 설정 ===")
	fmt.Println("사용 가능한 리전: KR1 (판교), KR2 (평촌), JP1 (도쿄)")
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
	return nil
}

func configureOAuth(reader *bufio.Reader) *config.ProfileConfig {
	fmt.Println("\n=== OAuth 인증 설정 ===")
	userAccessKeyID := readInput(reader, "User Access Key ID: ", "")
	secretAccessKey := readSecretInput(reader, "Secret Access Key: ")

	return &config.ProfileConfig{
		AuthType:        config.AuthTypeOAuth,
		UserAccessKeyID: userAccessKeyID,
		SecretAccessKey: secretAccessKey,
	}
}

func configureIdentity(reader *bufio.Reader) *config.ProfileConfig {
	fmt.Println("\n=== Identity 인증 설정 ===")
	tenantID := readInput(reader, "Tenant ID: ", "")
	username := readInput(reader, "Username (NHN Cloud ID): ", "")
	password := readSecretInput(reader, "API Password: ")

	return &config.ProfileConfig{
		AuthType: config.AuthTypeIdentity,
		TenantID: tenantID,
		Username: username,
		Password: password,
	}
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
