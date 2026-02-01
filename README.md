# NHN Cloud CLI

AWS CLI 스타일의 NHN Cloud 명령줄 인터페이스입니다.

[![Go Version](https://img.shields.io/badge/Go-1.22+-00ADD8?style=flat&logo=go)](https://go.dev/)
[![License](https://img.shields.io/badge/License-PolyForm%20Noncommercial-blue.svg)](LICENSE)

## 📋 목차

- [개요](#-개요)
- [주요 기능](#-주요-기능)
- [설치](#-설치)
- [초기 설정](#-초기-설정)
- [사용법](#-사용법)
  - [VPC 명령어](#vpc-명령어)
  - [Compute 명령어](#compute-명령어)
  - [Block Storage 명령어](#block-storage-명령어)
  - [Load Balancer 명령어](#load-balancer-명령어)
  - [DNS Plus 명령어](#dns-plus-명령어)
  - [Object Storage 명령어](#object-storage-명령어)
  - [Pipeline 명령어](#pipeline-명령어)
  - [Deploy 명령어](#deploy-명령어)
  - [CDN 명령어](#cdn-명령어)
  - [AppGuard 명령어](#appguard-명령어)
  - [Gamebase 명령어](#gamebase-명령어)
- [전역 옵션](#-전역-옵션)
- [실전 예제](#-실전-예제)
- [설정 파일](#-설정-파일)
- [아키텍처](#-아키텍처)
- [API 엔드포인트](#-api-엔드포인트)
- [문제 해결](#-문제-해결)
- [라이선스](#-라이선스)

---

## 🌟 개요

NHN Cloud CLI는 NHN Cloud 서비스를 명령줄에서 관리할 수 있는 도구입니다. AWS CLI와 유사한 사용법을 제공하여 친숙하게 사용할 수 있습니다.

```bash
# VPC 목록 조회
$ nhn vpc list
ID                                      NAME            CIDR            STATE
8a5f3e2c-...                            my-vpc          192.168.0.0/16  available

# 인스턴스 목록 조회
$ nhn compute instance list
ID                                      NAME        STATUS  FLAVOR      IP ADDRESSES    AZ
a1b2c3d4-...                            web-server  ACTIVE  m2.c1m2     192.168.1.10    kr-pub-a
```

---

## ✨ 주요 기능

### 인증 (Authentication)

| 기능 | 설명 |
|------|------|
| **Identity 인증** | Tenant ID + Username + Password (VPC, Compute API 필수) |
| **OAuth 인증** | User Access Key ID + Secret Access Key (필수) |
| **토큰 캐싱** | 자동 토큰 갱신 및 캐싱 |
| **다중 프로필** | 여러 계정/환경 프로필 관리 |
| **서비스별 AppKey 설정** | `nhn configure service <name>` (dns, pipeline, deploy, cdn, appguard, gamebase) |
| **서비스별 AppKey 오버라이드** | `--app-key` / `--secret-key` 플래그로 프로필 설정 오버라이드 |

> **참고**: Identity와 OAuth 인증 모두 필수입니다. 각 인증 방식은 다른 API에 사용됩니다.

### VPC (Virtual Private Cloud)

| 기능 | 명령어 |
|------|--------|
| VPC 관리 | `nhn vpc list/describe/create/update/delete` |
| 서브넷 관리 | `nhn vpc subnet list/describe/create/delete` |
| 보안 그룹 관리 | `nhn vpc securitygroup list/create/delete/add-rule` |
| 플로팅 IP 관리 | `nhn vpc floatingip list/create/associate/delete` |
| 라우팅 테이블 조회 | `nhn vpc routingtable list/describe` |
| 네트워크 인터페이스 | `nhn vpc port list/describe/create/delete` |

### Compute (Instance)

| 기능 | 명령어 |
|------|--------|
| 인스턴스 관리 | `nhn compute instance list/describe/create/delete` |
| 인스턴스 제어 | `nhn compute instance start/stop/reboot` |
| 인스턴스 타입 조회 | `nhn compute flavor list/describe` |
| 이미지 조회 | `nhn compute image list/describe` |
| 키페어 관리 | `nhn compute keypair list/create/delete` |
| 가용성 영역 조회 | `nhn compute az list` |

### Block Storage

| 기능 | 명령어 |
|------|--------|
| 볼륨 관리 | `nhn blockstorage volume list/describe/create/delete` |
| 스냅샷 관리 | `nhn blockstorage snapshot list/describe/create/delete` |
| 볼륨 타입 조회 | `nhn blockstorage type list` |

### Load Balancer

| 기능 | 명령어 |
|------|--------|
| 로드 밸런서 관리 | `nhn loadbalancer list/describe/create/update/delete` |
| 리스너 관리 | `nhn loadbalancer listener list/describe/create/delete` |

> **참고**: `loadbalancer`는 `lb`로 축약하여 사용할 수 있습니다. (예: `nhn lb list`)

### DNS Plus

| 기능 | 명령어 |
|------|--------|
| Zone 관리 | `nhn dns zone list/describe/create/update/delete` |
| Record Set 관리 | `nhn dns recordset list/describe/create/update/delete` |

### Object Storage

| 기능 | 명령어 |
|------|--------|
| 컨테이너 관리 | `nhn objectstorage container list/describe/create/delete` |
| 오브젝트 관리 | `nhn objectstorage object list/upload/download/describe/delete` |

> **참고**: `objectstorage`는 `os`로 축약하여 사용할 수 있습니다. (예: `nhn os container list`)

### Pipeline

| 기능 | 명령어 |
|------|--------|
| 파이프라인 실행 | `nhn pipeline execute [pipeline-name]` |

### Deploy

| 기능 | 명령어 |
|------|--------|
| 배포 실행 | `nhn deploy execute` |
| 바이너리 업로드 | `nhn deploy upload` |

### CDN

| 기능 | 명령어 |
|------|--------|
| CDN 서비스 관리 | `nhn cdn service list/create/update/delete` |
| 캐시 퍼지 | `nhn cdn purge [domain]` |
| 인증 토큰 생성 | `nhn cdn auth-token create` |

### AppGuard

| 기능 | 명령어 |
|------|--------|
| 탐지 현황 조회 | `nhn appguard dashboard` |

### Gamebase

| 기능 | 명령어 |
|------|--------|
| 회원 관리 | `nhn gamebase member list/describe/withdraw` |
| 이용 정지 관리 | `nhn gamebase ban list/create/release` |
| 론칭 상태 조회 | `nhn gamebase launching` |
| 인증 토큰 검증 | `nhn gamebase auth validate` |

---

## 📦 설치

### 요구사항

- Go 1.22 이상

### 소스에서 빌드

```bash
# 저장소 클론
git clone https://github.com/sangjinsu/nhn-cli.git
cd nhn-cli

# 빌드
go build -o nhn main.go

# 실행 파일 이동 (Linux/macOS)
sudo mv nhn /usr/local/bin/

# Windows의 경우 nhn.exe를 PATH에 포함된 디렉토리로 이동
```

### 설치 확인

```bash
nhn version
# 출력: nhn version 0.1.0 (built: unknown)
```

---

## 🔐 초기 설정

### 인증 정보 설정

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

### NHN Cloud 인증 정보 발급 방법

#### OAuth 인증 키 발급

1. [NHN Cloud 콘솔](https://console.nhncloud.com) 로그인
2. 오른쪽 상단의 이메일 주소 클릭
3. **API 보안 설정** 메뉴 선택
4. **User Access Key ID 생성** 버튼 클릭
5. User Access Key ID와 Secret Access Key 발급

#### Identity 인증 정보 확인

1. [NHN Cloud 콘솔](https://console.nhncloud.com) 로그인
2. **Compute > Instance** 메뉴 이동
3. **API 엔드포인트 설정** 버튼 클릭
4. Tenant ID 확인 및 API 비밀번호 설정

### 프로필 관리

```bash
# 프로필 목록 보기
nhn configure list

# 특정 프로필로 설정
nhn configure --profile production

# 특정 프로필 사용
nhn --profile production vpc list
```

### 서비스별 AppKey 설정

AppKey가 필요한 서비스는 별도로 설정합니다:

```bash
# DNS Plus AppKey 설정
nhn configure service dns

# CDN AppKey + Secret Key 설정
nhn configure service cdn

# 지원 서비스: dns, pipeline, deploy, cdn, appguard, gamebase
```

---

## 📖 사용법

### VPC 명령어

#### VPC 관리

```bash
# VPC 목록 조회
nhn vpc list

# VPC 상세 조회
nhn vpc describe <vpc-id>

# VPC 생성
nhn vpc create --name my-vpc --cidr 192.168.0.0/16

# VPC 이름/CIDR 수정
nhn vpc update <vpc-id> --name new-name --cidr 192.168.0.0/20

# VPC 삭제
nhn vpc delete <vpc-id>
```

#### 서브넷 관리

```bash
# 서브넷 목록 조회
nhn vpc subnet list

# 특정 VPC의 서브넷만 조회
nhn vpc subnet list --vpc-id <vpc-id>

# 서브넷 상세 조회
nhn vpc subnet describe <subnet-id>

# 서브넷 생성
nhn vpc subnet create \
  --vpc-id <vpc-id> \
  --name my-subnet \
  --cidr 192.168.1.0/24

# 서브넷 삭제
nhn vpc subnet delete <subnet-id>
```

#### 보안 그룹 관리

```bash
# 보안 그룹 목록
nhn vpc securitygroup list
# 또는 별칭 사용
nhn vpc sg list

# 보안 그룹 상세 조회 (규칙 포함)
nhn vpc sg describe <sg-id>

# 보안 그룹 생성
nhn vpc sg create --name my-sg --description "My security group"

# 인바운드 규칙 추가 (SSH)
nhn vpc sg add-rule <sg-id> \
  --direction ingress \
  --protocol tcp \
  --port 22 \
  --remote-ip 0.0.0.0/0

# 인바운드 규칙 추가 (HTTP/HTTPS 포트 범위)
nhn vpc sg add-rule <sg-id> \
  --direction ingress \
  --protocol tcp \
  --port-range 80-443 \
  --remote-ip 0.0.0.0/0

# 인바운드 규칙 추가 (ICMP)
nhn vpc sg add-rule <sg-id> \
  --direction ingress \
  --protocol icmp \
  --remote-ip 0.0.0.0/0

# 보안 그룹 규칙 삭제
nhn vpc sg delete-rule <rule-id>

# 보안 그룹 삭제
nhn vpc sg delete <sg-id>
```

#### 플로팅 IP 관리

```bash
# 플로팅 IP 목록
nhn vpc floatingip list
# 또는 별칭 사용
nhn vpc fip list

# 플로팅 IP 생성
nhn vpc fip create

# 인스턴스에 플로팅 IP 연결
nhn vpc fip associate <floatingip-id> --instance-id <instance-id>

# 포트에 직접 연결
nhn vpc fip associate <floatingip-id> --port-id <port-id>

# 플로팅 IP 연결 해제
nhn vpc fip disassociate <floatingip-id>

# 플로팅 IP 삭제
nhn vpc fip delete <floatingip-id>
```

#### 라우팅 테이블 조회

```bash
# 라우팅 테이블 목록
nhn vpc routingtable list
# 또는 별칭 사용
nhn vpc rt list

# 라우팅 테이블 상세 조회
nhn vpc rt describe <routingtable-id>
```

#### 네트워크 인터페이스 (포트) 관리

```bash
# 포트 목록
nhn vpc port list

# 포트 상세 조회
nhn vpc port describe <port-id>

# 포트 생성
nhn vpc port create --network-id <network-id> --name my-port

# 포트 삭제
nhn vpc port delete <port-id>
```

---

### Compute 명령어

#### 인스턴스 관리

```bash
# 인스턴스 목록 조회
nhn compute instance list

# 인스턴스 상세 조회
nhn compute instance describe <instance-id>

# 인스턴스 생성
nhn compute instance create \
  --name my-server \
  --image-id <image-id> \
  --flavor-id <flavor-id> \
  --network-id <network-id> \
  --key-name my-keypair \
  --security-group default \
  --availability-zone kr-pub-a

# 인스턴스 삭제
nhn compute instance delete <instance-id>
```

#### 인스턴스 제어

```bash
# 인스턴스 시작
nhn compute instance start <instance-id>

# 인스턴스 중지
nhn compute instance stop <instance-id>

# 인스턴스 재부팅 (소프트)
nhn compute instance reboot <instance-id>

# 인스턴스 재부팅 (하드)
nhn compute instance reboot <instance-id> --hard
```

#### 인스턴스 타입 (Flavor)

```bash
# 인스턴스 타입 목록
nhn compute flavor list

# 인스턴스 타입 상세 조회
nhn compute flavor describe <flavor-id>
```

#### 이미지

```bash
# 이미지 목록
nhn compute image list

# 이미지 상세 조회
nhn compute image describe <image-id>
```

#### 키페어

```bash
# 키페어 목록
nhn compute keypair list

# 키페어 생성 (새 키 생성 - 개인키 출력됨)
nhn compute keypair create --name my-keypair

# 키페어 생성 (기존 공개키 등록)
nhn compute keypair create --name my-keypair \
  --public-key "ssh-rsa AAAA..."

# 키페어 삭제
nhn compute keypair delete my-keypair
```

#### 가용성 영역

```bash
# 가용성 영역 목록
nhn compute az list
```

---

### Block Storage 명령어

#### 볼륨 관리

```bash
# 볼륨 목록 조회
nhn blockstorage volume list

# 볼륨 상세 조회
nhn blockstorage volume describe <volume-id>

# 볼륨 생성
nhn blockstorage volume create \
  --size 20 \
  --name my-volume \
  --type SSD \
  --availability-zone kr-pub-a

# 스냅샷에서 볼륨 생성
nhn blockstorage volume create \
  --size 20 \
  --name restored-volume \
  --snapshot-id <snapshot-id>

# 볼륨 삭제
nhn blockstorage volume delete <volume-id>
```

#### 스냅샷 관리

```bash
# 스냅샷 목록 조회
nhn blockstorage snapshot list

# 스냅샷 상세 조회
nhn blockstorage snapshot describe <snapshot-id>

# 스냅샷 생성
nhn blockstorage snapshot create \
  --volume-id <volume-id> \
  --name my-snapshot

# 사용 중인 볼륨 강제 스냅샷
nhn blockstorage snapshot create \
  --volume-id <volume-id> \
  --name forced-snapshot \
  --force

# 스냅샷 삭제
nhn blockstorage snapshot delete <snapshot-id>
```

#### 볼륨 타입 조회

```bash
# 볼륨 타입 목록
nhn blockstorage type list
```

---

### Load Balancer 명령어

#### 로드 밸런서 관리

```bash
# 로드 밸런서 목록 조회
nhn lb list

# 로드 밸런서 상세 조회
nhn lb describe <lb-id>

# 로드 밸런서 생성
nhn lb create \
  --vip-subnet-id <subnet-id> \
  --name web-lb \
  --type shared

# 로드 밸런서 수정
nhn lb update <lb-id> --name new-name --description "Updated"

# 로드 밸런서 삭제
nhn lb delete <lb-id>
```

#### 리스너 관리

```bash
# 리스너 목록 조회
nhn lb listener list

# 리스너 상세 조회
nhn lb listener describe <listener-id>

# 리스너 생성
nhn lb listener create \
  --loadbalancer-id <lb-id> \
  --protocol HTTP \
  --port 80 \
  --name http-listener

# 리스너 삭제
nhn lb listener delete <listener-id>
```

---

### DNS Plus 명령어

#### Zone 관리

```bash
# Zone 목록 조회
nhn dns zone list

# Zone 상세 조회
nhn dns zone describe <zone-id>

# Zone 생성
nhn dns zone create --name example.com. --description "My Zone"

# Zone 수정
nhn dns zone update <zone-id> --description "Updated description"

# Zone 삭제
nhn dns zone delete <zone-id>
```

#### Record Set 관리

```bash
# Record Set 목록 조회
nhn dns recordset list --zone-id <zone-id>

# Record Set 상세 조회
nhn dns recordset describe <recordset-id> --zone-id <zone-id>

# A 레코드 생성
nhn dns recordset create --zone-id <zone-id> \
  --name www.example.com. \
  --type A \
  --ttl 300 \
  --data 1.2.3.4

# Record Set 수정
nhn dns recordset update <recordset-id> --zone-id <zone-id> \
  --ttl 600 --data 1.2.3.4

# Record Set 삭제
nhn dns recordset delete <recordset-id> --zone-id <zone-id>
```

---

### Object Storage 명령어

#### 컨테이너 관리

```bash
# 컨테이너 목록 조회
nhn os container list

# 컨테이너 메타데이터 조회
nhn os container describe <container-name>

# 컨테이너 생성
nhn os container create my-container

# 컨테이너 삭제
nhn os container delete my-container
```

#### 오브젝트 관리

```bash
# 오브젝트 목록 조회
nhn os object list --container my-container

# 파일 업로드
nhn os object upload --container my-container --file ./test.txt

# 커스텀 이름으로 업로드
nhn os object upload --container my-container --file ./test.txt --name custom-name.txt

# 오브젝트 다운로드
nhn os object download test.txt --container my-container

# 오브젝트 메타데이터 조회
nhn os object describe test.txt --container my-container

# 오브젝트 삭제
nhn os object delete test.txt --container my-container
```

---

### Pipeline 명령어

```bash
# 파이프라인 수동 실행
nhn pipeline execute <pipeline-name>
```

> **참고**: Pipeline AppKey가 필요합니다. `nhn configure service pipeline`로 설정하거나 `--app-key` 플래그를 사용하세요.

---

### Deploy 명령어

```bash
# 배포 실행
nhn deploy execute \
  --server-group-id <id> \
  --artifact-id <id> \
  --deploy-note "배포 메모"

# 비동기 실행
nhn deploy execute \
  --server-group-id <id> \
  --artifact-id <id> \
  --async

# 바이너리 업로드
nhn deploy upload \
  --artifact-id <id> \
  --binary-group-key <key> \
  --type server \
  --file ./app.jar \
  --version "1.0.0" \
  --description "릴리즈"

# client 타입 바이너리 업로드 (iOS)
nhn deploy upload \
  --artifact-id <id> \
  --binary-group-key <key> \
  --type client \
  --file ./app.ipa \
  --os-type iOS \
  --meta-file ./app.plist
```

> **참고**: Deploy AppKey가 필요합니다. `nhn configure service deploy`로 설정하거나 `--app-key` 플래그를 사용하세요.

---

### CDN 명령어

#### CDN 서비스 관리

```bash
# CDN 서비스 목록 조회
nhn cdn service list

# CDN 서비스 생성
nhn cdn service create

# CDN 서비스 수정
nhn cdn service update

# CDN 서비스 삭제
nhn cdn service delete
```

#### 캐시 퍼지

```bash
# 전체 퍼지
nhn cdn purge <domain> --type ALL

# 특정 경로 퍼지
nhn cdn purge <domain> --type ITEM --items "/path1,/path2"
```

#### 인증 토큰

```bash
# 인증 토큰 생성
nhn cdn auth-token create
```

> **참고**: CDN AppKey와 Secret Key가 필요합니다. `nhn configure service cdn`으로 설정하거나 `--app-key`, `--secret-key` 플래그를 사용하세요.

---

### AppGuard 명령어

```bash
# 비정상 행위 탐지 현황 조회
nhn appguard dashboard --target-date 2025-01-01

# iOS 탐지 현황
nhn appguard dashboard --target-date 2025-01-01 --os 2
```

> **참고**: AppGuard AppKey가 필요합니다. `nhn configure service appguard`로 설정하거나 `--app-key` 플래그를 사용하세요.

---

### Gamebase 명령어

#### 회원 관리

```bash
# 회원 조회
nhn gamebase member describe <user-id>

# 회원 일괄 조회
nhn gamebase member list

# 회원 탈퇴
nhn gamebase member withdraw <user-id>
```

#### 이용 정지 관리

```bash
# 이용 정지 목록 조회
nhn gamebase ban list

# 이용 정지
nhn gamebase ban create

# 이용 정지 해제
nhn gamebase ban release
```

#### 론칭 상태 및 인증

```bash
# 론칭 상태 조회
nhn gamebase launching

# 인증 토큰 검증
nhn gamebase auth validate
```

> **참고**: Gamebase App ID와 Secret Key가 필요합니다. `nhn configure service gamebase`로 설정하거나 `--app-key`, `--secret-key` 플래그를 사용하세요.

---

## 🔧 전역 옵션

모든 명령어에서 사용 가능한 옵션:

| 옵션 | 설명 | 기본값 |
|------|------|--------|
| `--profile <name>` | 사용할 프로필 | default |
| `--region <region>` | 리전 지정 (프로필 설정 오버라이드) | 프로필 설정값 |
| `--output <format>` | 출력 형식 (table, json) | table |
| `--debug` | 디버그 모드 | false |
| `--help` | 도움말 표시 | - |

### 사용 예시

```bash
# production 프로필로 KR2 리전의 인스턴스 조회
nhn --profile production --region KR2 compute instance list

# JSON 형식으로 출력
nhn --output json vpc list

# 디버그 모드로 실행 (HTTP 요청/응답 출력)
nhn --debug vpc list
```

### 서비스별 옵션

AppKey가 필요한 서비스(dns, pipeline, deploy, cdn, appguard, gamebase)에서 사용 가능:

| 옵션 | 설명 | 대상 서비스 |
|------|------|-------------|
| `--app-key <key>` | 서비스 AppKey (프로필 설정 오버라이드) | 모든 AppKey 서비스 |
| `--secret-key <key>` | Secret Key (프로필 설정 오버라이드) | cdn, gamebase |

```bash
# 프로필 대신 플래그로 AppKey 지정
nhn dns zone list --app-key my-dns-appkey

# CDN은 AppKey + SecretKey 모두 지정 가능
nhn cdn service list --app-key my-cdn-appkey --secret-key my-secret

# Gamebase도 AppKey(App ID) + SecretKey 지정 가능
nhn gamebase member describe user123 --app-key my-app-id --secret-key my-secret
```

---

## 🎯 실전 예제

### 예제 1: 기본 인프라 구성

VPC, 서브넷, 보안 그룹을 생성하고 인스턴스를 배포하는 전체 과정:

```bash
# 1. VPC 생성
nhn vpc create --name my-vpc --cidr 192.168.0.0/16
# 출력된 VPC ID 기록: vpc-12345678

# 2. 서브넷 생성
nhn vpc subnet create \
  --vpc-id vpc-12345678 \
  --name public-subnet \
  --cidr 192.168.1.0/24
# 출력된 Subnet ID (Network ID로 사용): subnet-87654321

# 3. 보안 그룹 생성 및 규칙 추가
nhn vpc sg create --name web-sg
# 출력된 Security Group ID: sg-11111111

nhn vpc sg add-rule sg-11111111 \
  --direction ingress --protocol tcp --port 22 --remote-ip 0.0.0.0/0

nhn vpc sg add-rule sg-11111111 \
  --direction ingress --protocol tcp --port 80 --remote-ip 0.0.0.0/0

nhn vpc sg add-rule sg-11111111 \
  --direction ingress --protocol tcp --port 443 --remote-ip 0.0.0.0/0

# 4. 키페어 생성 및 저장
nhn compute keypair create --name my-keypair > my-keypair.pem
chmod 400 my-keypair.pem

# 5. 이미지 및 Flavor 확인
nhn compute image list
nhn compute flavor list

# 6. 인스턴스 생성
nhn compute instance create \
  --name web-server \
  --image-id <image-id> \
  --flavor-id m2.c1m2 \
  --network-id subnet-87654321 \
  --key-name my-keypair \
  --security-group web-sg \
  --availability-zone kr-pub-a
# 출력된 Instance ID: instance-99999999

# 7. 플로팅 IP 생성 및 연결
nhn vpc fip create
# 출력된 Floating IP ID: fip-44444444

nhn vpc fip associate fip-44444444 --instance-id instance-99999999

# 8. SSH 접속
ssh -i my-keypair.pem centos@<floating-ip>
```

### 예제 2: 다중 환경 관리

```bash
# 개발 환경 설정
nhn configure --profile dev
# ... 인증 정보 입력

# 운영 환경 설정
nhn configure --profile prod
# ... 인증 정보 입력

# 각 환경별 리소스 확인
nhn --profile dev compute instance list
nhn --profile prod compute instance list

# 스크립트에서 사용
for env in dev staging prod; do
  echo "=== $env 환경 인스턴스 ==="
  nhn --profile $env compute instance list
done
```

### 예제 3: JSON 출력과 jq 활용

```bash
# 모든 인스턴스의 이름과 IP 추출
nhn --output json compute instance list | \
  jq '.[] | {name: .name, addresses: .addresses}'

# 실행 중인 인스턴스만 필터링
nhn --output json compute instance list | \
  jq '.[] | select(.status == "ACTIVE")'

# VPC별 서브넷 개수
nhn --output json vpc subnet list | \
  jq 'group_by(.vpc_id) | .[] | {vpc_id: .[0].vpc_id, count: length}'

# 특정 보안 그룹의 인바운드 규칙만 추출
nhn --output json vpc sg describe <sg-id> | \
  jq '.security_group_rules[] | select(.direction == "ingress")'
```

### 예제 4: 인스턴스 일괄 작업

```bash
# 모든 인스턴스 ID 추출
INSTANCES=$(nhn --output json compute instance list | jq -r '.[].id')

# 모든 인스턴스 중지
for id in $INSTANCES; do
  echo "Stopping $id..."
  nhn compute instance stop $id
done

# 특정 상태의 인스턴스만 처리
nhn --output json compute instance list | \
  jq -r '.[] | select(.status == "SHUTOFF") | .id' | \
  xargs -I {} nhn compute instance start {}
```

---

## 📁 설정 파일

설정 파일은 `~/.nhn/` 디렉토리에 저장됩니다:

### ~/.nhn/config.json

프로필 및 인증 정보 저장:

```json
{
  "profiles": {
    "default": {
      "tenant_id": "your-tenant-id",
      "username": "your-email@example.com",
      "password": "your-api-password",
      "user_access_key_id": "your-access-key-id",
      "secret_access_key": "your-secret-access-key",
      "region": "KR1"
    },
    "production": {
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

> **참고**: Identity 인증(tenant_id, username, password)과 OAuth 인증(user_access_key_id, secret_access_key) 모두 필수입니다.

### ~/.nhn/credentials.json

토큰 캐시 (자동 생성/관리):

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

## 🏗 아키텍처

```
┌──────────────────────────────────────────────────────────────────────────────┐
│                              NHN Cloud CLI                                    │
├──────────────────────────────────────────────────────────────────────────────┤
│  configure │ vpc │ compute │ blockstorage │ lb │ dns │ objectstorage         │
│  pipeline  │ deploy │ cdn │ appguard │ gamebase                              │
├──────────────────────────────────────────────────────────────────────────────┤
│                          Internal Modules                                     │
│  config │ auth │ vpc │ compute │ blockstorage │ lb │ dns │ objectstorage     │
│  pipeline │ deploy │ cdn │ appguard │ gamebase │ output                      │
├──────────────────────────────────────────────────────────────────────────────┤
│                           HTTP Client                                         │
└──────────────────────────┬───────────────────────────────────────────────────┘
                           │
                           ▼
┌──────────────────────────────────────────────────────────────────────────────┐
│                           NHN Cloud APIs                                      │
│  OAuth │ Identity │ VPC │ Compute │ BlockStorage │ LB │ DNS Plus             │
│  ObjectStorage │ Pipeline │ Deploy │ CDN │ AppGuard │ Gamebase               │
└──────────────────────────────────────────────────────────────────────────────┘
```

### 디렉토리 구조

```
nhncli/
├── main.go                    # 엔트리포인트
├── cmd/
│   ├── root.go                # 루트 명령어, 전역 플래그
│   ├── configure.go           # nhn configure
│   ├── version.go             # nhn version
│   ├── vpc/
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
│   ├── compute/
│   │   ├── compute.go         # nhn compute
│   │   ├── instance.go        # nhn compute instance *
│   │   ├── flavor.go          # nhn compute flavor *
│   │   ├── image.go           # nhn compute image *
│   │   ├── keypair.go         # nhn compute keypair *
│   │   └── az.go              # nhn compute az *
│   ├── blockstorage/
│   │   ├── blockstorage.go    # nhn blockstorage
│   │   ├── volume.go          # nhn blockstorage volume *
│   │   ├── snapshot.go        # nhn blockstorage snapshot *
│   │   └── type.go            # nhn blockstorage type *
│   ├── loadbalancer/
│   │   ├── loadbalancer.go    # nhn loadbalancer (lb)
│   │   └── listener.go        # nhn loadbalancer listener *
│   ├── dns/                  # nhn dns *
│   ├── objectstorage/        # nhn objectstorage (os) *
│   ├── pipeline/             # nhn pipeline *
│   ├── deploy/               # nhn deploy *
│   ├── cdn/                  # nhn cdn *
│   ├── appguard/             # nhn appguard *
│   └── gamebase/             # nhn gamebase *
└── internal/
    ├── config/
    │   ├── config.go          # 설정 로드/저장
    │   └── profile.go         # 프로필 관리
    ├── auth/
    │   ├── auth.go            # Authenticator 인터페이스
    │   ├── oauth.go           # OAuth 인증
    │   ├── identity.go        # Identity 인증
    │   ├── cache.go           # 토큰 캐싱
    │   └── types.go           # 인증 타입 정의
    ├── client/
    │   ├── client.go          # 공통 HTTP 클라이언트
    │   └── errors.go          # API 에러 처리
    ├── vpc/
    │   ├── client.go          # VPC API 클라이언트
    │   ├── types.go           # VPC 타입 정의
    │   ├── vpc.go             # VPC CRUD
    │   ├── subnet.go          # 서브넷 CRUD
    │   ├── securitygroup.go   # 보안 그룹
    │   ├── floatingip.go      # 플로팅 IP
    │   ├── routingtable.go    # 라우팅 테이블
    │   └── port.go            # 네트워크 인터페이스
    ├── compute/
    │   ├── client.go          # Compute API 클라이언트
    │   ├── types.go           # Compute 타입 정의
    │   ├── instance.go        # 인스턴스 CRUD + 액션
    │   ├── flavor.go          # 인스턴스 타입
    │   ├── image.go           # 이미지
    │   ├── keypair.go         # 키페어
    │   └── az.go              # 가용성 영역
    ├── blockstorage/
    │   ├── client.go          # Block Storage API 클라이언트
    │   ├── types.go           # Block Storage 타입 정의
    │   ├── volume.go          # 볼륨 CRUD
    │   └── snapshot.go        # 스냅샷 CRUD
    ├── loadbalancer/
    │   ├── client.go          # Load Balancer API 클라이언트
    │   ├── types.go           # Load Balancer 타입 정의
    │   ├── lb.go              # 로드 밸런서 CRUD
    │   └── listener.go        # 리스너 CRUD
    ├── dns/                  # DNS Plus API 클라이언트
    ├── objectstorage/        # Object Storage API 클라이언트
    ├── pipeline/             # Pipeline API 클라이언트
    ├── deploy/               # Deploy API 클라이언트
    ├── cdn/                  # CDN API 클라이언트
    ├── appguard/             # AppGuard API 클라이언트
    ├── gamebase/             # Gamebase API 클라이언트
    └── output/
        └── output.go          # 출력 포매터 (table, json)
```

---

## 🌐 API 엔드포인트

### 인증 API

| 서비스 | 엔드포인트 |
|--------|-----------|
| OAuth | `https://oauth.api.nhncloudservice.com` |
| Identity | `https://api-identity-infrastructure.nhncloudservice.com` |

### VPC API

| 리전 | 엔드포인트 |
|------|-----------|
| KR1 (판교) | `https://kr1-api-network-infrastructure.nhncloudservice.com` |
| KR2 (평촌) | `https://kr2-api-network-infrastructure.nhncloudservice.com` |
| JP1 (도쿄) | `https://jp1-api-network-infrastructure.nhncloudservice.com` |

### Compute API

| 리전 | 엔드포인트 |
|------|-----------|
| KR1 (판교) | `https://kr1-api-instance-infrastructure.nhncloudservice.com` |
| KR2 (평촌) | `https://kr2-api-instance-infrastructure.nhncloudservice.com` |
| JP1 (도쿄) | `https://jp1-api-instance-infrastructure.nhncloudservice.com` |

### Block Storage API

| 리전 | 엔드포인트 |
|------|-----------|
| KR1 (판교) | `https://kr1-api-block-storage-infrastructure.nhncloudservice.com` |
| KR2 (평촌) | `https://kr2-api-block-storage-infrastructure.nhncloudservice.com` |
| JP1 (도쿄) | `https://jp1-api-block-storage-infrastructure.nhncloudservice.com` |

### Load Balancer API

Load Balancer API는 VPC API와 동일한 네트워크 엔드포인트를 사용합니다.

| 리전 | 엔드포인트 |
|------|-----------|
| KR1 (판교) | `https://kr1-api-network-infrastructure.nhncloudservice.com` |
| KR2 (평촌) | `https://kr2-api-network-infrastructure.nhncloudservice.com` |
| JP1 (도쿄) | `https://jp1-api-network-infrastructure.nhncloudservice.com` |

### DNS Plus API

DNS Plus는 글로벌 서비스로 리전 구분 없이 단일 엔드포인트를 사용합니다.

| 엔드포인트 |
|-----------|
| `https://dnsplus.api.nhncloudservice.com` |

### Object Storage API

| 리전 | 엔드포인트 |
|------|-----------|
| KR1 (판교) | `https://kr1-api-object-storage.nhncloudservice.com` |
| KR2 (평촌) | `https://kr2-api-object-storage.nhncloudservice.com` |
| JP1 (도쿄) | `https://jp1-api-object-storage.nhncloudservice.com` |

### 리전 정보

| 리전 코드 | 위치 | 설명 |
|-----------|------|------|
| KR1 | 한국 (판교) | 기본 리전 |
| KR2 | 한국 (평촌) | - |
| JP1 | 일본 (도쿄) | - |

---

## 🔧 문제 해결

### 인증 오류

```bash
# 토큰 캐시 삭제 후 재시도
rm ~/.nhn/credentials.json
nhn compute instance list
```

### "Tenant ID가 필요합니다" 오류

Compute API는 Tenant ID가 필요합니다. 해결 방법:
1. Identity 인증 사용 (Tenant ID 포함)
2. 또는 설정 파일에 직접 tenant_id 추가

### 네트워크 오류

```bash
# 리전 엔드포인트 확인
nhn --region KR1 compute instance list

# 디버그 모드로 HTTP 요청 확인
nhn --debug compute instance list
```

### 권한 오류

- Tenant ID가 올바른지 확인
- API 비밀번호가 만료되지 않았는지 확인
- 프로젝트 멤버 권한 확인

### 프로필 문제

```bash
# 프로필 목록 확인
nhn configure list

# 특정 프로필로 테스트
nhn --profile <profile-name> vpc list
```

---

## 📜 라이선스

PolyForm Noncommercial License 1.0.0

이 소프트웨어는 비상업적 목적으로만 사용할 수 있습니다. 자세한 내용은 [LICENSE](LICENSE) 파일을 참조하세요.

---

## 📚 참고 문서

- [NHN Cloud 공식 문서](https://docs.nhncloud.com/)
- [NHN Cloud VPC API](https://docs.nhncloud.com/ko/Network/VPC/ko/public-api/)
- [NHN Cloud Instance API](https://docs.nhncloud.com/ko/Compute/Instance/ko/public-api/)
- [NHN Cloud 인증 API](https://docs.nhncloud.com/ko/nhncloud/ko/public-api/api-authentication/)

---

## 🚀 향후 개발 계획

- [x] Block Storage 관리
- [x] Load Balancer 관리
- [x] Object Storage 관리
- [x] DNS 관리
- [x] Pipeline 관리
- [x] Deploy 관리
- [x] CDN 관리
- [x] AppGuard 관리
- [x] Gamebase 관리
- [ ] Auto Scale 관리
- [ ] 자동완성 지원 (bash, zsh, fish)
- [ ] 설정 파일 암호화
- [ ] 대화형 모드

---

## 🤝 기여하기

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request
