package cdn

import (
	"fmt"

	"nhncli/cmd"
	"nhncli/internal/cdn"

	"github.com/spf13/cobra"
)

var authTokenCmd = &cobra.Command{
	Use:   "auth-token",
	Short: "인증 토큰 관리",
}

var authTokenCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "인증 토큰 생성",
	RunE:  runAuthTokenCreate,
}

var (
	sessionID          string
	singlePath         string
	singleWildcardPath string
	durationSeconds    int
)

func init() {
	CDNCmd.AddCommand(authTokenCmd)
	authTokenCmd.AddCommand(authTokenCreateCmd)

	authTokenCreateCmd.Flags().StringVar(&sessionID, "session-id", "", "세션 ID")
	authTokenCreateCmd.Flags().StringVar(&singlePath, "single-path", "", "단일 경로")
	authTokenCreateCmd.Flags().StringVar(&singleWildcardPath, "single-wildcard-path", "", "와일드카드 경로")
	authTokenCreateCmd.Flags().IntVar(&durationSeconds, "duration", 0, "유효 시간 (초)")
}

func runAuthTokenCreate(c *cobra.Command, args []string) error {
	appKey, _ := c.Flags().GetString("app-key")
	secretKey, _ := c.Flags().GetString("secret-key")
	opts := cdn.ClientOption{AppKey: appKey, SecretKey: secretKey}
	cdnClient, err := cdn.NewClient(cmd.GetProfile(), cmd.GetDebug(), opts)
	if err != nil {
		return err
	}

	req := &cdn.AuthTokenCreateRequest{
		SessionID:          sessionID,
		SinglePath:         singlePath,
		SingleWildcardPath: singleWildcardPath,
		DurationSeconds:    durationSeconds,
	}

	token, err := cdnClient.CreateAuthToken(req)
	if err != nil {
		return err
	}

	fmt.Printf("인증 토큰: %s\n", token)
	return nil
}
