# DNS Plus 명령어

DNS Plus 서비스를 관리하는 명령어입니다.

> **참고**: DNS Plus는 AppKey 기반 인증을 사용합니다. 사전에 `nhn configure service dns`로 AppKey를 설정하거나, `--app-key` 플래그를 사용하세요.

---

## Zone 관리

### Zone 목록 조회

```bash
nhn dns zone list
```

출력 예시:
```
ZONE ID                                 NAME            STATUS  RECORDS DESCRIPTION
550e8400-e29b-41d4-a716-446655440000    example.com.    ACTIVE  5       My Zone
```

### Zone 상세 조회

```bash
nhn dns zone describe <zone-id>
```

출력 예시:
```
Zone ID:       550e8400-e29b-41d4-a716-446655440000
Name:          example.com.
Status:        ACTIVE
Record Sets:   5
Description:   My Zone
Created At:    2024-01-15T10:30:00Z
Updated At:    2024-01-15T10:30:00Z
```

### Zone 생성

```bash
nhn dns zone create --name <zone-name> [--description <desc>]
```

| 옵션 | 설명 | 필수 |
|------|------|------|
| `--name` | Zone 이름 (FQDN, 예: example.com.) | O |
| `--description` | Zone 설명 | X |

**예시:**
```bash
# Zone 생성
nhn dns zone create --name example.com. --description "My Domain Zone"
```

### Zone 수정

```bash
nhn dns zone update <zone-id> --description <desc>
```

| 옵션 | 설명 | 필수 |
|------|------|------|
| `--description` | Zone 설명 | X |

### Zone 삭제

```bash
nhn dns zone delete <zone-id>
```

---

## Record Set 관리

### Record Set 목록 조회

```bash
nhn dns recordset list --zone-id <zone-id>
```

| 옵션 | 설명 | 필수 |
|------|------|------|
| `--zone-id` | Zone ID | O |

출력 예시:
```
RECORDSET ID                            NAME            TYPE    TTL     STATUS  RECORDS
660e8400-e29b-41d4-a716-446655440001    www.example.com. A      300     ACTIVE  1.2.3.4, 5.6.7.8
```

### Record Set 상세 조회

```bash
nhn dns recordset describe <recordset-id> --zone-id <zone-id>
```

출력 예시:
```
Recordset ID:  660e8400-e29b-41d4-a716-446655440001
Name:          www.example.com.
Type:          A
TTL:           300
Status:        ACTIVE
Created At:    2024-01-15T11:00:00Z
Updated At:    2024-01-15T11:00:00Z
Records:
  - 1.2.3.4
  - 5.6.7.8
```

### Record Set 생성

```bash
nhn dns recordset create --zone-id <zone-id> --name <name> --type <type> --data <data> [--ttl <ttl>]
```

| 옵션 | 설명 | 필수 |
|------|------|------|
| `--zone-id` | Zone ID | O |
| `--name` | Record Set 이름 | O |
| `--type` | 레코드 타입 (A, AAAA, CNAME, MX, TXT 등) | O |
| `--data` | 레코드 데이터 (반복 가능) | O |
| `--ttl` | TTL (초, 기본값: 300) | X |

**예시:**
```bash
# A 레코드 생성
nhn dns recordset create --zone-id <zone-id> \
  --name www.example.com. \
  --type A \
  --data 1.2.3.4

# 다중 레코드 생성
nhn dns recordset create --zone-id <zone-id> \
  --name www.example.com. \
  --type A \
  --ttl 600 \
  --data 1.2.3.4 --data 5.6.7.8

# CNAME 레코드 생성
nhn dns recordset create --zone-id <zone-id> \
  --name blog.example.com. \
  --type CNAME \
  --data www.example.com.

# MX 레코드 생성
nhn dns recordset create --zone-id <zone-id> \
  --name example.com. \
  --type MX \
  --data "10 mail.example.com."
```

### Record Set 수정

```bash
nhn dns recordset update <recordset-id> --zone-id <zone-id> [--ttl <ttl>] [--data <data>]
```

| 옵션 | 설명 | 필수 |
|------|------|------|
| `--zone-id` | Zone ID | O |
| `--type` | 레코드 타입 | X |
| `--ttl` | TTL (초) | X |
| `--data` | 레코드 데이터 (반복 가능) | X |

**예시:**
```bash
# TTL 변경
nhn dns recordset update <recordset-id> --zone-id <zone-id> --ttl 600

# 레코드 데이터 변경
nhn dns recordset update <recordset-id> --zone-id <zone-id> --data 10.20.30.40
```

### Record Set 삭제

```bash
nhn dns recordset delete <recordset-id> --zone-id <zone-id>
```

---

## AppKey 설정

### 프로필에 AppKey 저장

```bash
nhn configure service dns
```

대화형 프롬프트:
```
프로필 이름 [default]:
DNS Plus AppKey: your-dns-appkey
```

### 명령줄에서 AppKey 지정

```bash
# --app-key 플래그로 직접 지정
nhn dns zone list --app-key your-dns-appkey
```

---

## 실전 예제

### 웹사이트 DNS 설정

```bash
# 1. Zone 생성
nhn dns zone create --name example.com. --description "Example Domain"

# 2. Zone ID 확인
ZONE_ID=$(nhn --output json dns zone list | jq -r '.[] | select(.zoneName=="example.com.") | .zoneId')

# 3. A 레코드 생성 (웹 서버)
nhn dns recordset create --zone-id $ZONE_ID \
  --name example.com. \
  --type A \
  --data 203.0.113.10

# 4. www CNAME 생성
nhn dns recordset create --zone-id $ZONE_ID \
  --name www.example.com. \
  --type CNAME \
  --data example.com.

# 5. MX 레코드 생성 (메일 서버)
nhn dns recordset create --zone-id $ZONE_ID \
  --name example.com. \
  --type MX \
  --data "10 mail.example.com."
```

---

## 참고

- [설정 가이드](../Configuration.md)
- [전역 옵션](Global-Options.md)
