# AppGuard 명령어

AppGuard 서비스를 관리하는 명령어입니다.

> **참고**: AppGuard는 AppKey 기반 인증을 사용합니다. 사전에 `nhn configure service appguard`로 AppKey를 설정하거나, `--app-key` 플래그를 사용하세요.

---

## 대시보드

### 비정상 행위 탐지 현황 조회

```bash
nhn appguard dashboard --target-date <date> [--os <os-type>] [--target-type <type>]
```

| 옵션 | 설명 | 필수 |
|------|------|------|
| `--target-date` | 조회 날짜 (YYYY-MM-DD) | O |
| `--os` | OS 타입 (1: Android, 2: iOS, 기본값: 1) | X |
| `--target-type` | 대상 타입 (0: 전체, 1: 앱, 기본값: 0) | X |

**예시:**
```bash
# 특정 날짜의 Android 탐지 현황
nhn appguard dashboard --target-date 2024-01-15

# iOS 탐지 현황
nhn appguard dashboard --target-date 2024-01-15 --os 2

# 앱별 탐지 현황
nhn appguard dashboard --target-date 2024-01-15 --target-type 1
```

출력 예시:
```
날짜          탐지 수     차단 수
----          ------      ------
2024-01-15    150         45
```

---

## AppKey 설정

### 프로필에 AppKey 저장

```bash
nhn configure service appguard
```

대화형 프롬프트:
```
프로필 이름 [default]:

=== APPGUARD 서비스 설정 ===
AppGuard AppKey: your-appguard-appkey
```

### 명령줄에서 AppKey 지정

```bash
# --app-key 플래그로 직접 지정
nhn appguard dashboard --target-date 2024-01-15 --app-key your-appguard-appkey
```

---

## 실전 예제

### 일일 보안 리포트 자동화

```bash
#!/bin/bash

# 어제 날짜 계산
YESTERDAY=$(date -d "yesterday" +%Y-%m-%d)

echo "=== AppGuard 일일 보안 리포트 ($YESTERDAY) ==="

# Android 탐지 현황
echo "## Android"
nhn appguard dashboard --target-date $YESTERDAY --os 1

# iOS 탐지 현황
echo "## iOS"
nhn appguard dashboard --target-date $YESTERDAY --os 2
```

### JSON 출력 활용

```bash
# JSON으로 데이터 추출
nhn --output json appguard dashboard --target-date 2024-01-15 | \
  jq '.[] | {date: .detectedDate, detected: .detectedCnt, blocked: .blockedCnt}'

# 탐지 수 합계 계산
nhn --output json appguard dashboard --target-date 2024-01-15 | \
  jq '[.[].detectedCnt] | add'
```

---

## 참고

- [설정 가이드](../Configuration.md)
- [전역 옵션](Global-Options.md)
