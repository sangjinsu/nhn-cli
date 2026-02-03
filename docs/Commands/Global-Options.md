# 전역 옵션

모든 NHN Cloud CLI 명령어에서 사용할 수 있는 공통 옵션입니다.

---

## 옵션 목록

| 옵션 | 설명 | 기본값 |
|------|------|--------|
| `--profile <name>` | 사용할 프로필 이름 | `default` |
| `--region <region>` | 리전 지정 (프로필 설정 오버라이드) | 프로필 설정값 |
| `--output <format>` | 출력 형식 (`table`, `json`) | `table` |
| `--debug` | 디버그 모드 활성화 | `false` |
| `--help`, `-h` | 도움말 표시 | - |

---

## --profile

프로필을 사용하여 여러 환경의 인증 정보를 관리합니다.

### 사용법

```bash
nhn --profile <profile-name> <command>
```

### 예시

```bash
# production 프로필로 인스턴스 조회
nhn --profile production compute instance list

# dev 프로필로 VPC 조회
nhn --profile dev vpc list
```

### 프로필 설정

```bash
# 프로필 생성
nhn configure --profile production

# 프로필 목록 확인
nhn configure list
```

---

## --region

프로필에 설정된 리전을 임시로 오버라이드합니다.

### 사용법

```bash
nhn --region <region-code> <command>
```

### 지원 리전

| 리전 코드 | 위치 |
|-----------|------|
| `KR1` | 한국 (판교) |
| `KR2` | 한국 (평촌) |
| `JP1` | 일본 (도쿄) |

### 예시

```bash
# KR2 리전의 VPC 조회 (프로필이 KR1이어도)
nhn --region KR2 vpc list

# JP1 리전의 인스턴스 조회
nhn --region JP1 compute instance list
```

---

## --output

출력 형식을 지정합니다.

### 사용법

```bash
nhn --output <format> <command>
```

### 지원 형식

| 형식 | 설명 |
|------|------|
| `table` | 표 형식 (기본값, 사람이 읽기 쉬움) |
| `json` | JSON 형식 (프로그래밍 처리용) |

### 예시

#### Table 출력 (기본)

```bash
nhn vpc list
```

```
ID                                      NAME            CIDR            STATE
8a5f3e2c-1234-5678-9abc-def012345678    my-vpc          192.168.0.0/16  available
```

#### JSON 출력

```bash
nhn --output json vpc list
```

```json
[
  {
    "id": "8a5f3e2c-1234-5678-9abc-def012345678",
    "name": "my-vpc",
    "cidr": "192.168.0.0/16",
    "state": "available"
  }
]
```

### jq와 함께 사용

```bash
# VPC 이름만 추출
nhn --output json vpc list | jq '.[].name'

# 특정 상태의 인스턴스 필터링
nhn --output json compute instance list | jq '.[] | select(.status == "ACTIVE")'

# 인스턴스 이름과 IP 추출
nhn --output json compute instance list | jq '.[] | {name: .name, ip: .addresses}'
```

---

## --debug

HTTP 요청/응답 정보를 출력하여 문제를 진단합니다.

### 사용법

```bash
nhn --debug <command>
```

### 예시

```bash
nhn --debug vpc list
```

### 출력 내용

- HTTP 요청 URL 및 메서드
- 요청 헤더
- 응답 상태 코드
- 응답 본문 (일부)

```
DEBUG: GET https://kr1-api-network-infrastructure.nhncloudservice.com/v2.0/vpcs
DEBUG: Request Headers:
  Authorization: Bearer eyJhbG...
  Content-Type: application/json
DEBUG: Response Status: 200 OK
DEBUG: Response Body: {"vpcs": [...]}

ID                                      NAME            CIDR            STATE
...
```

---

## --help

명령어 도움말을 표시합니다.

### 사용법

```bash
nhn --help
nhn <command> --help
nhn <command> <subcommand> --help
```

### 예시

```bash
# 전체 도움말
nhn --help

# VPC 명령어 도움말
nhn vpc --help

# 인스턴스 생성 도움말
nhn compute instance create --help
```

---

## 서비스별 옵션

AppKey가 필요한 서비스에서 사용할 수 있는 추가 옵션입니다.

### --app-key

프로필에 설정된 AppKey를 오버라이드합니다.

| 대상 서비스 | 설명 |
|-------------|------|
| `dns` | DNS Plus AppKey |
| `pipeline` | Pipeline AppKey |
| `deploy` | Deploy AppKey |
| `cdn` | CDN AppKey |
| `appguard` | AppGuard AppKey |
| `gamebase` | Gamebase App ID |

### --secret-key

프로필에 설정된 Secret Key를 오버라이드합니다.

| 대상 서비스 | 설명 |
|-------------|------|
| `cdn` | CDN Secret Key |
| `gamebase` | Gamebase Secret Key |

### 예시

```bash
# DNS Plus - AppKey만 필요
nhn dns zone list --app-key your-dns-appkey

# Pipeline - AppKey만 필요
nhn pipeline execute my-pipeline --app-key your-pipeline-appkey

# Deploy - AppKey만 필요
nhn deploy execute --artifact-id 123 --server-group-id 456 --app-key your-deploy-appkey

# CDN - AppKey + Secret Key 필요
nhn cdn service list --app-key your-cdn-appkey --secret-key your-cdn-secret-key

# AppGuard - AppKey만 필요
nhn appguard dashboard --target-date 2024-01-15 --app-key your-appguard-appkey

# Gamebase - AppKey(App ID) + Secret Key 필요
nhn gamebase member describe user123 --app-key your-app-id --secret-key your-secret-key
```

---

## 옵션 조합

여러 전역 옵션을 함께 사용할 수 있습니다.

### 예시

```bash
# production 프로필로 KR2 리전의 인스턴스를 JSON으로 출력
nhn --profile production --region KR2 --output json compute instance list

# 디버그 모드로 VPC 목록 조회
nhn --debug --output json vpc list

# 프로필과 AppKey 오버라이드 조합
nhn --profile prod dns zone list --app-key custom-appkey
```

---

## 환경 변수

전역 옵션 대신 환경 변수를 사용할 수 있습니다.

| 환경 변수 | 대응 옵션 |
|-----------|----------|
| `NHN_PROFILE` | `--profile` |
| `NHN_REGION` | `--region` |
| `NHN_DEBUG` | `--debug` |

### 예시

```bash
# 환경 변수로 기본값 설정
export NHN_PROFILE=production
export NHN_REGION=KR2

# 이후 명령은 production 프로필, KR2 리전 사용
nhn vpc list
nhn compute instance list

# 명령줄 옵션으로 오버라이드 가능
nhn --region KR1 vpc list  # KR1 리전 사용
```

---

## 우선순위

옵션 값 결정 우선순위 (높은 것이 우선):

1. **명령줄 옵션** - `--profile`, `--region`, `--app-key` 등
2. **환경 변수** - `NHN_PROFILE`, `NHN_REGION` 등
3. **프로필 설정** - `~/.nhn/config.json`
4. **기본값** - `default` 프로필, `KR1` 리전

---

## 참고

- [설정 가이드](../Configuration.md)
- [VPC 명령어](VPC.md)
- [Compute 명령어](Compute.md)
- [DNS Plus 명령어](DNS.md)
