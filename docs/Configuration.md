# 설정 가이드

NHN Cloud CLI의 인증 및 프로필 설정 방법을 설명합니다.

---

## 인증 방식

NHN Cloud CLI는 **Identity 인증**과 **OAuth 인증** 모두를 필수로 요구합니다:

| 인증 방식 | 필요 정보 | 용도 |
|----------|----------|------|
| **Identity 인증** | Tenant ID, Username, Password | VPC, Compute 등 OpenStack 기반 API |
| **OAuth 인증** | User Access Key ID, Secret Access Key | 기타 NHN Cloud API |

> **참고**: 두 인증 방식 모두 필수입니다. `nhn configure` 명령 실행 시 순차적으로 입력합니다.

---

## 인증 정보 발급

### Identity 인증 정보

1. [NHN Cloud 콘솔](https://console.nhncloud.com) 로그인
2. **Compute > Instance** 메뉴 이동
3. **API 엔드포인트 설정** 버튼 클릭
4. **Tenant ID** 확인
5. **API 비밀번호** 설정 (미설정 시 새로 생성)

### OAuth 인증 정보

1. [NHN Cloud 콘솔](https://console.nhncloud.com) 로그인
2. 오른쪽 상단의 이메일 주소 클릭
3. **API 보안 설정** 메뉴 선택
4. **User Access Key ID 생성** 버튼 클릭
5. **User Access Key ID**와 **Secret Access Key** 저장

> Secret Access Key는 발급 시 한 번만 표시됩니다. 안전한 곳에 보관하세요.

---

## CLI 설정

```bash
nhn configure
```

대화형 프롬프트에서 Identity와 OAuth 인증 정보를 순차적으로 입력합니다:

```
프로필 이름 [default]:

=== NHN Cloud 인증 설정 ===

📌 VPC, Compute 등 OpenStack 기반 API 사용을 위해 Identity 인증 정보가 필요합니다.

--- Identity 인증 (필수) ---

📌 Tenant ID 확인 방법:
   1. NHN Cloud 콘솔 (https://console.nhncloud.com) 로그인
   2. 프로젝트 선택 후 'Compute > Instance' 메뉴 이동
   3. 'API 엔드포인트 설정' 버튼 클릭
   4. Tenant ID 확인

📌 API Password 설정 방법:
   위 'API 엔드포인트 설정' 화면에서 'API 비밀번호 설정' 클릭

Tenant ID: your-tenant-id
Username (이메일 주소): your-email@example.com
API Password: your-api-password

--- OAuth 인증 (필수) ---

📌 User Access Key ID 발급 방법:
   1. NHN Cloud 콘솔 (https://console.nhncloud.com) 로그인
   2. 오른쪽 상단의 이메일 주소 클릭
   3. 'API 보안 설정' 메뉴 선택
   4. 'User Access Key ID 생성' 버튼 클릭

User Access Key ID: your-access-key-id
Secret Access Key: your-secret-access-key

=== 리전 설정 ===

사용 가능한 리전:
   KR1 - 한국 (판교) 리전
   KR2 - 한국 (평촌) 리전
   JP1 - 일본 (도쿄) 리전

기본 리전 [KR1]: KR1

✅ 프로필 'default' 설정이 저장되었습니다.

🔐 Identity 인증 정보 검증 중...
✅ Identity 인증 성공!
   Tenant ID: your-tenant-id
   토큰이 캐시되었습니다. (유효기간: 12시간)
   OAuth 인증 정보도 저장되었습니다.
```

---

## 서비스별 AppKey 설정

일부 서비스는 별도의 AppKey가 필요합니다. `nhn configure service` 명령으로 설정합니다.

### 지원 서비스

| 서비스 | 명령어 | 필요 정보 |
|--------|--------|----------|
| DNS Plus | `nhn configure service dns` | AppKey |
| Pipeline | `nhn configure service pipeline` | AppKey |
| Deploy | `nhn configure service deploy` | AppKey |
| CDN | `nhn configure service cdn` | AppKey + Secret Key |
| AppGuard | `nhn configure service appguard` | AppKey |
| Gamebase | `nhn configure service gamebase` | App ID + Secret Key |

### 설정 예시

```bash
# DNS Plus AppKey 설정
nhn configure service dns
```

대화형 프롬프트:
```
프로필 이름 [default]:

=== DNS 서비스 설정 ===
DNS Plus AppKey: your-dns-appkey

✅ 프로필 'default'의 dns 서비스 설정이 저장되었습니다.
```

```bash
# CDN AppKey + Secret Key 설정
nhn configure service cdn
```

대화형 프롬프트:
```
프로필 이름 [default]:

=== CDN 서비스 설정 ===
CDN AppKey: your-cdn-appkey
CDN Secret Key: your-cdn-secret-key

✅ 프로필 'default'의 cdn 서비스 설정이 저장되었습니다.
```

### 명령줄에서 AppKey 오버라이드

프로필에 저장된 AppKey 대신 `--app-key` 플래그로 직접 지정할 수 있습니다:

```bash
# DNS Plus
nhn dns zone list --app-key your-dns-appkey

# CDN (AppKey + Secret Key)
nhn cdn service list --app-key your-cdn-appkey --secret-key your-cdn-secret-key

# Gamebase (App ID + Secret Key)
nhn gamebase member describe user123 --app-key your-app-id --secret-key your-secret-key
```

---

## 프로필 관리

여러 환경(개발, 스테이징, 운영)을 관리하려면 프로필을 사용하세요.

### 프로필 생성

```bash
# 개발 환경 프로필
nhn configure --profile dev

# 운영 환경 프로필
nhn configure --profile prod
```

### 프로필 목록 확인

```bash
nhn configure list
```

출력 예시:
```
PROFILE     IDENTITY    OAUTH    REGION
default     ✓           ✓        KR1
dev         ✓           ✓        KR1
prod        ✓           ✓        KR2
```

### 프로필 사용

```bash
# 특정 프로필로 명령 실행
nhn --profile prod compute instance list

# 환경별 리소스 비교
nhn --profile dev vpc list
nhn --profile prod vpc list
```

---

## 환경 변수

설정 파일 대신 환경 변수를 사용할 수 있습니다:

| 환경 변수 | 설명 | 예시 |
|-----------|------|------|
| `NHN_PROFILE` | 기본 프로필 | `export NHN_PROFILE=prod` |
| `NHN_REGION` | 기본 리전 | `export NHN_REGION=KR2` |
| `NHN_DEBUG` | 디버그 모드 | `export NHN_DEBUG=true` |

```bash
# 환경 변수로 프로필 설정
export NHN_PROFILE=prod
nhn vpc list  # prod 프로필 사용

# 일회성 환경 변수
NHN_REGION=KR2 nhn compute instance list
```

---

## 설정 파일 구조

설정 파일은 `~/.nhn/` 디렉토리에 저장됩니다.

### ~/.nhn/config.json

프로필 및 인증 정보:

```json
{
  "profiles": {
    "default": {
      "tenant_id": "your-tenant-id",
      "username": "your-email@example.com",
      "password": "your-api-password",
      "user_access_key_id": "your-access-key-id",
      "secret_access_key": "your-secret-access-key",
      "region": "KR1",
      "app_key": "dns-appkey",
      "pipeline_app_key": "pipeline-appkey",
      "deploy_app_key": "deploy-appkey",
      "cdn_app_key": "cdn-appkey",
      "cdn_secret_key": "cdn-secret-key",
      "appguard_app_key": "appguard-appkey",
      "gamebase_app_id": "gamebase-app-id",
      "gamebase_secret_key": "gamebase-secret-key"
    },
    "prod": {
      "tenant_id": "your-tenant-id",
      "username": "your-email@example.com",
      "password": "your-api-password",
      "user_access_key_id": "your-access-key-id",
      "secret_access_key": "your-secret-access-key",
      "region": "KR2"
    }
  }
}
```

> **참고**: Identity 인증(tenant_id, username, password)과 OAuth 인증(user_access_key_id, secret_access_key) 모두 필수입니다. 서비스별 AppKey는 해당 서비스 사용 시에만 필요합니다.

### ~/.nhn/credentials.json

토큰 캐시 (자동 관리):

```json
{
  "profiles": {
    "default": {
      "access_token": "cached-token...",
      "expires_at": 1704067200,
      "tenant_id": "tenant-id-from-token"
    }
  }
}
```

---

## 토큰 캐싱

CLI는 인증 토큰을 자동으로 캐싱합니다:

- **유효 기간**: 12시간
- **자동 갱신**: 토큰 만료 시 자동으로 새 토큰 발급
- **저장 위치**: `~/.nhn/credentials.json`

### 토큰 캐시 삭제

인증 문제 발생 시 캐시를 삭제하세요:

```bash
rm ~/.nhn/credentials.json
```

---

## 보안 권장 사항

### 파일 권한 설정

```bash
# 설정 디렉토리 권한 제한
chmod 700 ~/.nhn
chmod 600 ~/.nhn/config.json
chmod 600 ~/.nhn/credentials.json
```

### API 키 관리

- Secret Access Key는 안전한 곳에 보관
- 정기적으로 API 키 회전
- 사용하지 않는 키는 비활성화
- CI/CD 환경에서는 환경 변수 사용

### 프로필 분리

- 개발/운영 환경 프로필 분리
- 팀원별 개인 프로필 사용
- 자동화 작업용 별도 프로필 생성

---

## 다음 단계

- [VPC 명령어](Commands/VPC.md)
- [Compute 명령어](Commands/Compute.md)
- [DNS Plus 명령어](Commands/DNS.md)
- [기본 인프라 구성 예제](Examples/Basic-Infrastructure.md)
