# CDN 명령어

CDN 서비스를 관리하는 명령어입니다.

> **참고**: CDN은 AppKey와 Secret Key 기반 인증을 사용합니다. 사전에 `nhn configure service cdn`으로 설정하거나, `--app-key` 및 `--secret-key` 플래그를 사용하세요.

---

## CDN 서비스 관리

### 서비스 목록 조회

```bash
nhn cdn service list
```

출력 예시:
```
도메인                          상태      원본 URL                    설명
------                          ----      -------                     ----
abc123.toastcdn.net             ACTIVE    https://origin.example.com  My CDN
def456.toastcdn.net             ACTIVE    https://api.example.com     API CDN
```

### 서비스 생성

```bash
nhn cdn service create --origin-url <url> [--domain-alias <alias>] [--description <desc>]
```

| 옵션 | 설명 | 필수 |
|------|------|------|
| `--origin-url` | 원본 서버 URL | O |
| `--domain-alias` | 도메인 별칭 | X |
| `--description` | 설명 | X |

**예시:**
```bash
# CDN 서비스 생성
nhn cdn service create \
  --origin-url https://origin.example.com \
  --description "My Website CDN"

# 도메인 별칭 설정
nhn cdn service create \
  --origin-url https://origin.example.com \
  --domain-alias cdn.example.com \
  --description "Custom Domain CDN"
```

출력 예시:
```
CDN 서비스가 생성되었습니다. (도메인: abc123.toastcdn.net)
```

### 서비스 수정

```bash
nhn cdn service update <domain> [--origin-url <url>] [--domain-alias <alias>] [--description <desc>]
```

| 옵션 | 설명 | 필수 |
|------|------|------|
| `--origin-url` | 원본 서버 URL | X |
| `--domain-alias` | 도메인 별칭 | X |
| `--description` | 설명 | X |

**예시:**
```bash
# 원본 URL 변경
nhn cdn service update abc123.toastcdn.net --origin-url https://new-origin.example.com

# 설명 변경
nhn cdn service update abc123.toastcdn.net --description "Updated CDN"
```

### 서비스 삭제

```bash
nhn cdn service delete <domain>
```

**예시:**
```bash
nhn cdn service delete abc123.toastcdn.net
```

---

## 캐시 퍼지

### 캐시 퍼지 요청

```bash
nhn cdn purge <domain> [--type <type>] [--items <items>]
```

| 옵션 | 설명 | 필수 |
|------|------|------|
| `--type` | 퍼지 타입 (`ALL` 또는 `ITEM`, 기본값: ALL) | X |
| `--items` | 퍼지 대상 경로 (쉼표 구분) | X |

**예시:**
```bash
# 전체 캐시 퍼지
nhn cdn purge abc123.toastcdn.net

# 특정 파일만 퍼지
nhn cdn purge abc123.toastcdn.net --type ITEM --items "/images/logo.png,/css/style.css"

# 디렉토리 퍼지
nhn cdn purge abc123.toastcdn.net --type ITEM --items "/api/*"
```

출력 예시:
```
캐시 퍼지가 요청되었습니다. (도메인: abc123.toastcdn.net, 타입: ITEM)
```

---

## AppKey 설정

### 프로필에 AppKey 저장

```bash
nhn configure service cdn
```

대화형 프롬프트:
```
프로필 이름 [default]:

=== CDN 서비스 설정 ===
CDN AppKey: your-cdn-appkey
CDN Secret Key: your-cdn-secret-key
```

### 명령줄에서 AppKey 지정

```bash
# --app-key 및 --secret-key 플래그로 직접 지정
nhn cdn service list --app-key your-appkey --secret-key your-secret-key
```

---

## 실전 예제

### 배포 후 캐시 자동 퍼지

```bash
#!/bin/bash

# 1. 파일 업로드 (S3 또는 원본 서버로)
aws s3 sync ./dist s3://my-bucket/

# 2. CDN 캐시 퍼지
nhn cdn purge abc123.toastcdn.net --type ALL

echo "배포 및 캐시 퍼지 완료"
```

### 특정 경로만 퍼지

```bash
# JavaScript 파일만 퍼지
nhn cdn purge abc123.toastcdn.net --type ITEM --items "/js/*.js"

# CSS 파일만 퍼지
nhn cdn purge abc123.toastcdn.net --type ITEM --items "/css/*.css"
```

### JSON 출력 활용

```bash
# CDN 서비스 목록을 JSON으로 출력
nhn --output json cdn service list | jq '.[] | {domain: .domain, status: .status}'
```

---

## 참고

- [설정 가이드](../Configuration.md)
- [전역 옵션](Global-Options.md)
