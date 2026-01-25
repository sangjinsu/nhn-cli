# 빠른 시작 가이드

이 가이드에서는 NHN Cloud CLI를 설치하고 첫 번째 명령을 실행하는 방법을 설명합니다.

---

## 전제 조건

- **Go 1.22 이상** - 소스에서 빌드하기 위해 필요
- **NHN Cloud 계정** - [NHN Cloud 콘솔](https://console.nhncloud.com)에서 가입
- **인증 정보** - OAuth 키 또는 Identity 정보

---

## 1단계: 설치

```bash
# 저장소 클론
git clone https://github.com/your-repo/nhncli.git
cd nhncli

# 빌드
go build -o nhn main.go

# 실행 파일 이동 (Linux/macOS)
sudo mv nhn /usr/local/bin/

# 설치 확인
nhn version
```

> Windows의 경우 `nhn.exe`를 PATH에 포함된 디렉토리로 이동하세요.

---

## 2단계: 인증 설정

```bash
nhn configure
```

대화형 프롬프트가 시작됩니다:

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

### OAuth 키 발급 방법

1. [NHN Cloud 콘솔](https://console.nhncloud.com) 로그인
2. 오른쪽 상단의 이메일 주소 클릭
3. **API 보안 설정** 메뉴 선택
4. **User Access Key ID 생성** 버튼 클릭
5. User Access Key ID와 Secret Access Key 저장

---

## 3단계: 첫 명령 실행

### VPC 목록 조회

```bash
nhn vpc list
```

출력 예시:
```
ID                                      NAME            CIDR            STATE
8a5f3e2c-1234-5678-9abc-def012345678    my-vpc          192.168.0.0/16  available
```

### 인스턴스 목록 조회

```bash
nhn compute instance list
```

출력 예시:
```
ID                                      NAME        STATUS  FLAVOR      IP ADDRESSES    AZ
a1b2c3d4-5678-9abc-def0-123456789abc    web-server  ACTIVE  m2.c1m2     192.168.1.10    kr-pub-a
```

---

## 4단계: JSON 출력

데이터를 프로그래밍 방식으로 처리하려면 JSON 형식으로 출력하세요:

```bash
nhn --output json vpc list
```

`jq`와 함께 사용:
```bash
nhn --output json vpc list | jq '.[].name'
```

---

## 다음 단계

- [설치 가이드](Installation.md) - 상세 설치 옵션
- [설정 가이드](Configuration.md) - 다중 프로필 및 고급 설정
- [VPC 명령어](Commands/VPC.md) - VPC 리소스 관리
- [Compute 명령어](Commands/Compute.md) - 인스턴스 관리
- [기본 인프라 구성 예제](Examples/Basic-Infrastructure.md) - 전체 인프라 구성 워크플로우

---

## 도움말

명령어 도움말 확인:
```bash
nhn --help
nhn vpc --help
nhn compute instance --help
```

문제가 발생하면 [문제 해결](Troubleshooting.md) 가이드를 참조하세요.
