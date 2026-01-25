# 설정 가이드

NHN Cloud CLI의 인증 및 프로필 설정 방법을 설명합니다.

---

## 인증 방식 비교

NHN Cloud CLI는 두 가지 인증 방식을 지원합니다:

| 항목 | OAuth 인증 (권장) | Identity 인증 |
|------|------------------|---------------|
| **필요 정보** | User Access Key ID, Secret Access Key | Tenant ID, Username, Password |
| **토큰 유효기간** | 12시간 | 12시간 |
| **보안성** | 높음 (키 회전 가능) | 중간 |
| **권장 사용** | 프로덕션, 자동화 | 개발, 테스트 |

---

## OAuth 인증 설정 (권장)

### 1. API 키 발급

1. [NHN Cloud 콘솔](https://console.nhncloud.com) 로그인
2. 오른쪽 상단의 이메일 주소 클릭
3. **API 보안 설정** 메뉴 선택
4. **User Access Key ID 생성** 버튼 클릭
5. **User Access Key ID**와 **Secret Access Key** 저장

> Secret Access Key는 발급 시 한 번만 표시됩니다. 안전한 곳에 보관하세요.

### 2. CLI 설정

```bash
nhn configure
```

```
프로필 이름 [default]:
=== 인증 방식 선택 ===
1. OAuth 인증 (User Access Key ID) - 권장
2. Identity 인증 (Tenant ID + Username)
선택 [1]: 1

=== OAuth 인증 설정 ===
User Access Key ID: your-access-key-id
Secret Access Key: your-secret-access-key

=== 리전 설정 ===
기본 리전 [KR1]: KR1

✅ 프로필 'default' 설정이 저장되었습니다.
```

---

## Identity 인증 설정

### 1. Tenant ID 및 API 비밀번호 확인

1. [NHN Cloud 콘솔](https://console.nhncloud.com) 로그인
2. **Compute > Instance** 메뉴 이동
3. **API 엔드포인트 설정** 버튼 클릭
4. **Tenant ID** 확인
5. **API 비밀번호** 설정 (미설정 시 새로 생성)

### 2. CLI 설정

```bash
nhn configure
```

```
프로필 이름 [default]:
=== 인증 방식 선택 ===
1. OAuth 인증 (User Access Key ID) - 권장
2. Identity 인증 (Tenant ID + Username)
선택 [1]: 2

=== Identity 인증 설정 ===
Tenant ID: your-tenant-id
Username (NHN Cloud ID): your-email@example.com
API Password: your-api-password

=== 리전 설정 ===
기본 리전 [KR1]: KR1

✅ 프로필 'default' 설정이 저장되었습니다.
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
PROFILE     AUTH TYPE   REGION
default     oauth       KR1
dev         oauth       KR1
prod        identity    KR2
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
      "auth_type": "oauth",
      "user_access_key_id": "your-access-key-id",
      "secret_access_key": "your-secret-access-key",
      "region": "KR1"
    },
    "prod": {
      "auth_type": "identity",
      "tenant_id": "your-tenant-id",
      "username": "your-email@example.com",
      "password": "your-api-password",
      "region": "KR2"
    }
  }
}
```

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
- [기본 인프라 구성 예제](Examples/Basic-Infrastructure.md)
