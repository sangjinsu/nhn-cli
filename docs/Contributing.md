# 기여 가이드

NHN Cloud CLI 프로젝트에 기여하는 방법입니다.

---

## 기여 방법

1. **버그 리포트** - 문제점 발견 시 Issue 등록
2. **기능 제안** - 새로운 기능 아이디어 제안
3. **코드 기여** - Pull Request 제출
4. **문서 개선** - 문서 오류 수정 및 개선

---

## 개발 환경 설정

### 요구 사항

- Go 1.22 이상
- Git
- Make (선택)

### 저장소 클론

```bash
# Fork 후 클론
git clone https://github.com/your-username/nhncli.git
cd nhncli

# 원본 저장소 추가
git remote add upstream https://github.com/original/nhncli.git
```

### 의존성 설치

```bash
go mod download
```

### 빌드 및 테스트

```bash
# 빌드
go build -o nhn main.go

# 테스트
go test ./...

# 린트 (golangci-lint 필요)
golangci-lint run
```

---

## 코드 스타일

### Go 코딩 컨벤션

- [Effective Go](https://go.dev/doc/effective_go) 준수
- [Go Code Review Comments](https://go.dev/wiki/CodeReviewComments) 참고
- `gofmt` 또는 `goimports`로 포맷팅

### 파일 구조

```
cmd/              # 명령어 정의 (사용자 인터페이스)
internal/         # 내부 로직 (비즈니스 로직)
  config/         # 설정 관리
  auth/           # 인증
  client/         # HTTP 클라이언트
  vpc/            # VPC API
  compute/        # Compute API
  output/         # 출력 포매터
```

### 네이밍 컨벤션

```go
// 파일명: snake_case
// security_group.go

// 패키지명: 짧고 소문자
package vpc

// 타입명: PascalCase
type SecurityGroup struct { ... }

// 함수/메서드: PascalCase (exported), camelCase (unexported)
func CreateSecurityGroup() { ... }
func parseResponse() { ... }

// 상수: PascalCase 또는 ALL_CAPS
const DefaultRegion = "KR1"
const MAX_RETRY_COUNT = 3
```

### 에러 처리

```go
// 에러는 항상 처리
result, err := doSomething()
if err != nil {
    return fmt.Errorf("failed to do something: %w", err)
}

// 사용자 친화적 에러 메시지
if err != nil {
    return fmt.Errorf("VPC를 생성할 수 없습니다: %w", err)
}
```

---

## 커밋 컨벤션

### 커밋 메시지 형식

```
<type>(<scope>): <subject>

<body>

<footer>
```

### Type

| Type | 설명 |
|------|------|
| `feat` | 새로운 기능 |
| `fix` | 버그 수정 |
| `docs` | 문서 변경 |
| `style` | 코드 포맷팅 (기능 변경 없음) |
| `refactor` | 리팩토링 |
| `test` | 테스트 추가/수정 |
| `chore` | 빌드, 설정 등 기타 변경 |

### 예시

```
feat(vpc): add security group rule deletion

- Add delete-rule subcommand to vpc sg
- Update help text and documentation

Closes #123
```

```
fix(auth): handle token refresh on 401 error

The previous implementation didn't retry after token refresh.
Now it automatically retries the request after refreshing the token.
```

---

## Pull Request

### PR 제출 전 체크리스트

- [ ] 코드가 빌드됨 (`go build`)
- [ ] 테스트 통과 (`go test ./...`)
- [ ] 린트 통과 (`golangci-lint run`)
- [ ] 관련 문서 업데이트
- [ ] 커밋 메시지가 컨벤션을 따름

### PR 템플릿

```markdown
## 설명
<!-- 변경 사항 설명 -->

## 변경 유형
- [ ] 버그 수정
- [ ] 새로운 기능
- [ ] 문서 변경
- [ ] 리팩토링

## 테스트
<!-- 테스트 방법 설명 -->

## 관련 Issue
<!-- Closes #123 -->
```

### PR 프로세스

1. Fork 저장소
2. 기능 브랜치 생성 (`git checkout -b feature/amazing-feature`)
3. 변경 사항 커밋 (`git commit -m 'feat: add amazing feature'`)
4. 브랜치 푸시 (`git push origin feature/amazing-feature`)
5. Pull Request 생성

---

## Issue 리포팅

### 버그 리포트

```markdown
## 버그 설명
<!-- 버그 내용 -->

## 재현 방법
1.
2.
3.

## 예상 동작
<!-- 예상되는 정상 동작 -->

## 실제 동작
<!-- 실제 발생한 동작 -->

## 환경
- OS:
- Go 버전:
- CLI 버전:

## 추가 정보
<!-- 스크린샷, 로그 등 -->
```

### 기능 요청

```markdown
## 기능 설명
<!-- 원하는 기능 -->

## 사용 사례
<!-- 이 기능이 필요한 이유 -->

## 제안 구현 방법
<!-- (선택) 구현 아이디어 -->
```

---

## 테스트 작성

### 단위 테스트

```go
// internal/vpc/vpc_test.go
package vpc

import (
    "testing"
)

func TestCreateVPC(t *testing.T) {
    tests := []struct {
        name    string
        input   CreateVPCRequest
        want    *VPC
        wantErr bool
    }{
        {
            name: "valid request",
            input: CreateVPCRequest{
                Name: "test-vpc",
                CIDR: "192.168.0.0/16",
            },
            want: &VPC{
                Name: "test-vpc",
                CIDR: "192.168.0.0/16",
            },
            wantErr: false,
        },
        {
            name: "invalid CIDR",
            input: CreateVPCRequest{
                Name: "test-vpc",
                CIDR: "invalid",
            },
            wantErr: true,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got, err := CreateVPC(tt.input)
            if (err != nil) != tt.wantErr {
                t.Errorf("CreateVPC() error = %v, wantErr %v", err, tt.wantErr)
                return
            }
            if !tt.wantErr && got.Name != tt.want.Name {
                t.Errorf("CreateVPC() = %v, want %v", got, tt.want)
            }
        })
    }
}
```

### 통합 테스트

```go
// +build integration

func TestVPCIntegration(t *testing.T) {
    // 실제 API를 사용하는 테스트
    // CI 환경에서 별도 실행
}
```

---

## 새 기능 추가

### 새 서비스 추가 예시

1. **타입 정의** (`internal/newservice/types.go`)

```go
package newservice

type Resource struct {
    ID   string `json:"id"`
    Name string `json:"name"`
}
```

2. **API 클라이언트** (`internal/newservice/client.go`)

```go
package newservice

type Client struct {
    BaseURL string
    Auth    auth.Authenticator
}

func NewClient(region string, auth auth.Authenticator) *Client {
    return &Client{
        BaseURL: getEndpoint(region),
        Auth:    auth,
    }
}

func (c *Client) ListResources() ([]Resource, error) {
    // API 호출 구현
}
```

3. **명령어 정의** (`cmd/newservice/newservice.go`)

```go
package newservice

import "github.com/spf13/cobra"

var Cmd = &cobra.Command{
    Use:   "newservice",
    Short: "Manage NewService resources",
}

func init() {
    Cmd.AddCommand(listCmd)
}
```

4. **루트에 등록** (`cmd/root.go`)

```go
import "nhncli/cmd/newservice"

func init() {
    rootCmd.AddCommand(newservice.Cmd)
}
```

---

## 릴리스 프로세스

1. 버전 업데이트
2. CHANGELOG 작성
3. 태그 생성 (`git tag v1.0.0`)
4. GitHub Release 생성
5. 바이너리 빌드 및 업로드

---

## 문의

- GitHub Issues: 버그 및 기능 요청
- Discussions: 일반 토론 및 질문

---

## 라이선스

이 프로젝트에 기여하면 MIT 라이선스에 동의하는 것입니다.
