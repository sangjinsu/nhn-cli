# 아키텍처

NHN Cloud CLI의 시스템 아키텍처 및 설계 구조입니다.

---

## 시스템 개요

```
┌─────────────────────────────────────────────────────────────────┐
│                         NHN Cloud CLI                           │
├─────────────────────────────────────────────────────────────────┤
│  ┌──────────┐  ┌──────────┐  ┌──────────┐  ┌──────────────────┐ │
│  │configure │  │   vpc    │  │ compute  │  │  (future cmds)   │ │
│  └────┬─────┘  └────┬─────┘  └────┬─────┘  └────────┬─────────┘ │
│       │             │             │                  │          │
│  ┌────▼─────────────▼─────────────▼──────────────────▼────────┐ │
│  │                    Internal Modules                         │ │
│  │  ┌─────────┐  ┌─────────┐  ┌─────────┐  ┌─────────────────┐│ │
│  │  │ config  │  │  auth   │  │   vpc   │  │    compute      ││ │
│  │  └────┬────┘  └────┬────┘  └────┬────┘  └────────┬────────┘│ │
│  └───────┼────────────┼────────────┼────────────────┼─────────┘ │
│          │            │            │                │           │
│  ┌───────▼────────────▼────────────▼────────────────▼─────────┐ │
│  │                     HTTP Client                             │ │
│  └─────────────────────────┬───────────────────────────────────┘ │
└────────────────────────────┼────────────────────────────────────┘
                             │
                             ▼
┌─────────────────────────────────────────────────────────────────┐
│                      NHN Cloud APIs                             │
│  ┌──────────┐  ┌──────────┐  ┌──────────┐  ┌──────────────────┐ │
│  │  OAuth   │  │ Identity │  │   VPC    │  │     Compute      │ │
│  │   API    │  │   API    │  │   API    │  │       API        │ │
│  └──────────┘  └──────────┘  └──────────┘  └──────────────────┘ │
└─────────────────────────────────────────────────────────────────┘
```

---

## 계층 구조

### 1. Command Layer (cmd/)

사용자 인터페이스를 제공하는 최상위 레이어입니다.

- **역할**: CLI 명령어 정의, 인자 파싱, 사용자 입출력
- **의존성**: Cobra 프레임워크
- **특징**: 각 명령어는 독립적인 서브 명령어로 구성

### 2. Internal Layer (internal/)

비즈니스 로직과 API 통신을 담당하는 핵심 레이어입니다.

- **config**: 설정 파일 관리
- **auth**: 인증 및 토큰 관리
- **vpc**: VPC API 클라이언트
- **compute**: Compute API 클라이언트
- **client**: 공통 HTTP 클라이언트
- **output**: 출력 포매터

### 3. External Layer

NHN Cloud API와의 통신을 담당합니다.

---

## 디렉토리 구조

```
nhncli/
├── main.go                    # 엔트리포인트
├── go.mod                     # Go 모듈 정의
├── go.sum                     # 의존성 체크섬
├── README.md                  # 프로젝트 문서
│
├── cmd/                       # 명령어 정의
│   ├── root.go                # 루트 명령어, 전역 플래그
│   ├── configure.go           # nhn configure
│   ├── version.go             # nhn version
│   │
│   ├── vpc/                   # VPC 명령어
│   │   ├── vpc.go             # nhn vpc
│   │   ├── list.go            # nhn vpc list
│   │   ├── describe.go        # nhn vpc describe
│   │   ├── create.go          # nhn vpc create
│   │   ├── update.go          # nhn vpc update
│   │   ├── delete.go          # nhn vpc delete
│   │   ├── subnet.go          # nhn vpc subnet *
│   │   ├── securitygroup.go   # nhn vpc securitygroup *
│   │   ├── floatingip.go      # nhn vpc floatingip *
│   │   ├── routingtable.go    # nhn vpc routingtable *
│   │   └── port.go            # nhn vpc port *
│   │
│   └── compute/               # Compute 명령어
│       ├── compute.go         # nhn compute
│       ├── instance.go        # nhn compute instance *
│       ├── flavor.go          # nhn compute flavor *
│       ├── image.go           # nhn compute image *
│       ├── keypair.go         # nhn compute keypair *
│       └── az.go              # nhn compute az *
│
└── internal/                  # 내부 모듈
    ├── config/                # 설정 관리
    │   ├── config.go          # 설정 로드/저장
    │   └── profile.go         # 프로필 관리
    │
    ├── auth/                  # 인증
    │   ├── auth.go            # Authenticator 인터페이스
    │   ├── oauth.go           # OAuth 인증
    │   ├── identity.go        # Identity 인증
    │   ├── cache.go           # 토큰 캐싱
    │   └── types.go           # 인증 타입 정의
    │
    ├── client/                # HTTP 클라이언트
    │   ├── client.go          # 공통 HTTP 클라이언트
    │   └── errors.go          # API 에러 처리
    │
    ├── vpc/                   # VPC API
    │   ├── client.go          # VPC API 클라이언트
    │   ├── types.go           # VPC 타입 정의
    │   ├── vpc.go             # VPC CRUD
    │   ├── subnet.go          # 서브넷 CRUD
    │   ├── securitygroup.go   # 보안 그룹
    │   ├── floatingip.go      # 플로팅 IP
    │   ├── routingtable.go    # 라우팅 테이블
    │   └── port.go            # 네트워크 인터페이스
    │
    ├── compute/               # Compute API
    │   ├── client.go          # Compute API 클라이언트
    │   ├── types.go           # Compute 타입 정의
    │   ├── instance.go        # 인스턴스 CRUD + 액션
    │   ├── flavor.go          # 인스턴스 타입
    │   ├── image.go           # 이미지
    │   ├── keypair.go         # 키페어
    │   └── az.go              # 가용성 영역
    │
    └── output/                # 출력 포매터
        └── output.go          # Table, JSON 포매터
```

---

## 인증 플로우

### OAuth 인증

```
┌─────────┐     ┌─────────┐     ┌─────────────┐     ┌─────────┐
│   CLI   │     │  Cache  │     │  OAuth API  │     │   API   │
└────┬────┘     └────┬────┘     └──────┬──────┘     └────┬────┘
     │               │                 │                 │
     │  Check Token  │                 │                 │
     │──────────────>│                 │                 │
     │               │                 │                 │
     │  Token Valid  │                 │                 │
     │<──────────────│                 │                 │
     │               │                 │                 │
     │         API Request with Token                    │
     │──────────────────────────────────────────────────>│
     │               │                 │                 │
     │         API Response                              │
     │<──────────────────────────────────────────────────│
```

### Token 만료 시

```
┌─────────┐     ┌─────────┐     ┌─────────────┐     ┌─────────┐
│   CLI   │     │  Cache  │     │  OAuth API  │     │   API   │
└────┬────┘     └────┬────┘     └──────┬──────┘     └────┬────┘
     │               │                 │                 │
     │  Check Token  │                 │                 │
     │──────────────>│                 │                 │
     │               │                 │                 │
     │ Token Expired │                 │                 │
     │<──────────────│                 │                 │
     │               │                 │                 │
     │     Request New Token           │                 │
     │────────────────────────────────>│                 │
     │               │                 │                 │
     │     New Token                   │                 │
     │<────────────────────────────────│                 │
     │               │                 │                 │
     │  Save Token   │                 │                 │
     │──────────────>│                 │                 │
     │               │                 │                 │
     │         API Request with Token                    │
     │──────────────────────────────────────────────────>│
```

---

## API 호출 플로우

```
┌─────────────────────────────────────────────────────────────┐
│                        Command                               │
│  nhn vpc list                                                │
└───────────────────────────┬─────────────────────────────────┘
                            │
                            ▼
┌─────────────────────────────────────────────────────────────┐
│                    Config Loader                             │
│  1. Load profile from ~/.nhn/config.json                     │
│  2. Apply --profile, --region overrides                      │
└───────────────────────────┬─────────────────────────────────┘
                            │
                            ▼
┌─────────────────────────────────────────────────────────────┐
│                    Authenticator                             │
│  1. Check token cache                                        │
│  2. Request new token if expired                             │
│  3. Return valid access token                                │
└───────────────────────────┬─────────────────────────────────┘
                            │
                            ▼
┌─────────────────────────────────────────────────────────────┐
│                    VPC Client                                │
│  1. Build API endpoint URL                                   │
│  2. Set headers (Authorization, Content-Type)                │
│  3. Make HTTP request                                        │
│  4. Parse response                                           │
└───────────────────────────┬─────────────────────────────────┘
                            │
                            ▼
┌─────────────────────────────────────────────────────────────┐
│                    Output Formatter                          │
│  1. Format data (table or json)                              │
│  2. Print to stdout                                          │
└─────────────────────────────────────────────────────────────┘
```

---

## 주요 컴포넌트

### Config Manager

```go
// 설정 로드
config, err := config.Load()

// 프로필 가져오기
profile, err := config.GetProfile("default")

// 프로필 저장
config.SetProfile("default", profile)
config.Save()
```

### Authenticator Interface

```go
type Authenticator interface {
    GetToken() (string, error)
    GetTenantID() (string, error)
}

// OAuth 구현
type OAuthAuthenticator struct { ... }

// Identity 구현
type IdentityAuthenticator struct { ... }
```

### API Client

```go
// VPC Client
type VPCClient struct {
    BaseURL string
    Auth    auth.Authenticator
    HTTP    *http.Client
}

func (c *VPCClient) ListVPCs() ([]VPC, error) { ... }
func (c *VPCClient) GetVPC(id string) (*VPC, error) { ... }
func (c *VPCClient) CreateVPC(req CreateVPCRequest) (*VPC, error) { ... }
```

### Output Formatter

```go
type Formatter interface {
    Format(data interface{}) string
}

// Table 포매터
type TableFormatter struct { ... }

// JSON 포매터
type JSONFormatter struct { ... }
```

---

## 설정 파일 구조

### ~/.nhn/config.json

```json
{
  "profiles": {
    "default": {
      "tenant_id": "...",
      "username": "...",
      "password": "...",
      "user_access_key_id": "...",
      "secret_access_key": "...",
      "region": "KR1"
    }
  }
}
```

> **참고**: Identity 인증(tenant_id, username, password)과 OAuth 인증(user_access_key_id, secret_access_key) 모두 필수입니다.

### ~/.nhn/credentials.json

```json
{
  "profiles": {
    "default": {
      "access_token": "...",
      "expires_at": 1704067200,
      "tenant_id": "..."
    }
  }
}
```

---

## 확장성

### 새 서비스 추가

1. `internal/<service>/` 디렉토리 생성
2. API 클라이언트 및 타입 정의
3. `cmd/<service>/` 명령어 추가
4. `cmd/root.go`에 서브 명령어 등록

### 새 리전 지원

`internal/client/endpoints.go`에 엔드포인트 추가:

```go
var endpoints = map[string]map[string]string{
    "KR1": {
        "vpc":     "https://kr1-api-network-infrastructure...",
        "compute": "https://kr1-api-instance-infrastructure...",
    },
    "KR2": { ... },
    "JP1": { ... },
    "NEW_REGION": { ... },  // 새 리전 추가
}
```

---

## 참고

- [API 레퍼런스](API-Reference.md)
- [기여 가이드](Contributing.md)
