# Gamebase 명령어

Gamebase 서비스를 관리하는 명령어입니다.

> **참고**: Gamebase는 App ID(AppKey)와 Secret Key 기반 인증을 사용합니다. 사전에 `nhn configure service gamebase`로 설정하거나, `--app-key` 및 `--secret-key` 플래그를 사용하세요.

---

## 회원 관리

### 회원 조회

```bash
nhn gamebase member describe <user-id>
```

출력 예시:
```
User ID: user123
Valid: Y
App ID: app-12345
등록일: 2024-01-01T00:00:00Z
최근 로그인: 2024-01-15T10:30:00Z
```

### 회원 일괄 조회

```bash
nhn gamebase member list --user-ids <user-ids>
```

| 옵션 | 설명 | 필수 |
|------|------|------|
| `--user-ids` | 사용자 ID 목록 (쉼표 구분) | O |

**예시:**
```bash
nhn gamebase member list --user-ids "user1,user2,user3"
```

출력 예시:
```
User ID     상태    등록일                  최근 로그인
-------     ----    -----                   ---------
user1       Y       2024-01-01T00:00:00Z    2024-01-15T10:30:00Z
user2       Y       2024-01-02T00:00:00Z    2024-01-14T15:00:00Z
user3       N       2024-01-03T00:00:00Z    2024-01-10T08:00:00Z
```

### 회원 탈퇴

```bash
nhn gamebase member withdraw <user-id>
```

**예시:**
```bash
nhn gamebase member withdraw user123
```

출력 예시:
```
회원 'user123'이(가) 탈퇴 처리되었습니다.
```

---

## 이용 정지 관리

### 이용 정지

```bash
nhn gamebase ban create --user-id <id> --begin-date <date> --end-date <date> [options]
```

| 옵션 | 설명 | 필수 |
|------|------|------|
| `--user-id` | 사용자 ID | O |
| `--begin-date` | 시작일 | O |
| `--end-date` | 종료일 | O |
| `--type` | 정지 타입 | X |
| `--reason` | 사유 | X |
| `--message` | 메시지 | X |

**예시:**
```bash
# 기본 이용 정지
nhn gamebase ban create \
  --user-id user123 \
  --begin-date "2024-01-15T00:00:00Z" \
  --end-date "2024-01-22T00:00:00Z" \
  --reason "비정상 행위" \
  --message "7일간 이용이 정지됩니다."
```

출력 예시:
```
사용자 'user123'이(가) 이용 정지되었습니다.
```

### 이용 정지 목록 조회

```bash
nhn gamebase ban list --user-id <user-id>
```

| 옵션 | 설명 | 필수 |
|------|------|------|
| `--user-id` | 사용자 ID | O |

**예시:**
```bash
nhn gamebase ban list --user-id user123
```

출력 예시:
```
User ID     타입    시작일                  종료일                  사유
-------     ----    -----                   -----                   ----
user123     BAN     2024-01-15T00:00:00Z    2024-01-22T00:00:00Z    비정상 행위
```

### 이용 정지 해제

```bash
nhn gamebase ban release --user-id <user-id>
```

| 옵션 | 설명 | 필수 |
|------|------|------|
| `--user-id` | 사용자 ID | O |

**예시:**
```bash
nhn gamebase ban release --user-id user123
```

출력 예시:
```
사용자 'user123'의 이용 정지가 해제되었습니다.
```

---

## AppKey 설정

### 프로필에 AppKey 저장

```bash
nhn configure service gamebase
```

대화형 프롬프트:
```
프로필 이름 [default]:

=== GAMEBASE 서비스 설정 ===
Gamebase App ID: your-gamebase-app-id
Gamebase Secret Key: your-gamebase-secret-key
```

### 명령줄에서 AppKey 지정

```bash
# --app-key 및 --secret-key 플래그로 직접 지정
nhn gamebase member describe user123 --app-key your-app-id --secret-key your-secret-key
```

---

## 실전 예제

### 부정 행위 사용자 일괄 정지

```bash
#!/bin/bash

# 부정 행위 사용자 목록 (파일에서 읽기)
USERS=$(cat suspicious_users.txt)

# 현재 날짜 및 7일 후 날짜
BEGIN_DATE=$(date -u +%Y-%m-%dT%H:%M:%SZ)
END_DATE=$(date -u -d "+7 days" +%Y-%m-%dT%H:%M:%SZ)

for USER_ID in $USERS; do
  nhn gamebase ban create \
    --user-id $USER_ID \
    --begin-date "$BEGIN_DATE" \
    --end-date "$END_DATE" \
    --reason "자동 탐지: 부정 행위" \
    --message "부정 행위로 인해 7일간 이용이 정지됩니다."
  echo "정지 완료: $USER_ID"
done
```

### 회원 정보 모니터링

```bash
# JSON으로 회원 정보 추출
nhn --output json gamebase member describe user123 | \
  jq '{userId: .userId, valid: .valid, lastLogin: .lastLoginDate}'

# 여러 사용자 일괄 조회
nhn --output json gamebase member list --user-ids "user1,user2,user3" | \
  jq '.[] | {userId: .userId, status: .valid}'
```

---

## 참고

- [설정 가이드](../Configuration.md)
- [전역 옵션](Global-Options.md)
