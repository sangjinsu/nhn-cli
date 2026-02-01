# NHN Cloud CLI

AWS CLI 스타일의 NHN Cloud 명령줄 인터페이스입니다.

## 개요

NHN Cloud CLI는 NHN Cloud 서비스를 명령줄에서 관리할 수 있는 도구입니다. AWS CLI와 유사한 사용법을 제공하여 친숙하게 사용할 수 있습니다.

## 지원 기능

### 인증 (Authentication)

| 기능 | 설명 |
|------|------|
| OAuth 인증 | User Access Key ID + Secret Access Key |
| Identity 인증 | Tenant ID + Username + Password |
| 토큰 캐싱 | 자동 토큰 갱신 및 캐싱 |
| 다중 프로필 | 여러 계정/환경 프로필 관리 |

### VPC (Virtual Private Cloud)

| 기능 | 명령어 |
|------|--------|
| VPC 목록 조회 | `nhn vpc list` |
| VPC 상세 조회 | `nhn vpc describe <vpc-id>` |
| VPC 생성 | `nhn vpc create` |
| VPC 수정 | `nhn vpc update <vpc-id>` |
| VPC 삭제 | `nhn vpc delete <vpc-id>` |
| 서브넷 목록 조회 | `nhn vpc subnet list` |
| 서브넷 상세 조회 | `nhn vpc subnet describe <subnet-id>` |
| 서브넷 생성 | `nhn vpc subnet create` |
| 서브넷 삭제 | `nhn vpc subnet delete <subnet-id>` |
| 라우팅 테이블 목록 | `nhn vpc routingtable list` |
| 보안 그룹 목록 | `nhn vpc securitygroup list` |
| 보안 그룹 생성 | `nhn vpc securitygroup create` |
| 보안 그룹 규칙 추가 | `nhn vpc securitygroup add-rule` |
| 플로팅 IP 목록 | `nhn vpc floatingip list` |
| 플로팅 IP 생성 | `nhn vpc floatingip create` |
| 플로팅 IP 연결 | `nhn vpc floatingip associate` |

### Compute (Instance)

| 기능 | 명령어 |
|------|--------|
| 인스턴스 목록 조회 | `nhn compute instance list` |
| 인스턴스 상세 조회 | `nhn compute instance describe <id>` |
| 인스턴스 생성 | `nhn compute instance create` |
| 인스턴스 삭제 | `nhn compute instance delete <id>` |
| 인스턴스 시작 | `nhn compute instance start <id>` |
| 인스턴스 중지 | `nhn compute instance stop <id>` |
| 인스턴스 재부팅 | `nhn compute instance reboot <id>` |
| 인스턴스 타입 목록 | `nhn compute flavor list` |
| 이미지 목록 | `nhn compute image list` |
| 키페어 목록 | `nhn compute keypair list` |
| 키페어 생성 | `nhn compute keypair create` |
| 가용성 영역 목록 | `nhn compute az list` |

### Block Storage

| 기능 | 명령어 |
|------|--------|
| 볼륨 목록 조회 | `nhn blockstorage volume list` |
| 볼륨 상세 조회 | `nhn blockstorage volume describe <id>` |
| 볼륨 생성 | `nhn blockstorage volume create` |
| 볼륨 삭제 | `nhn blockstorage volume delete <id>` |
| 스냅샷 목록 조회 | `nhn blockstorage snapshot list` |
| 스냅샷 상세 조회 | `nhn blockstorage snapshot describe <id>` |
| 스냅샷 생성 | `nhn blockstorage snapshot create` |
| 스냅샷 삭제 | `nhn blockstorage snapshot delete <id>` |
| 볼륨 타입 목록 | `nhn blockstorage type list` |

### Load Balancer

| 기능 | 명령어 |
|------|--------|
| 로드 밸런서 목록 조회 | `nhn loadbalancer list` (별칭: `nhn lb list`) |
| 로드 밸런서 상세 조회 | `nhn loadbalancer describe <id>` |
| 로드 밸런서 생성 | `nhn loadbalancer create` |
| 로드 밸런서 수정 | `nhn loadbalancer update <id>` |
| 로드 밸런서 삭제 | `nhn loadbalancer delete <id>` |
| 리스너 목록 조회 | `nhn loadbalancer listener list` |
| 리스너 상세 조회 | `nhn loadbalancer listener describe <id>` |
| 리스너 생성 | `nhn loadbalancer listener create` |
| 리스너 삭제 | `nhn loadbalancer listener delete <id>` |

### DNS Plus

| 기능 | 명령어 |
|------|--------|
| Zone 목록 조회 | `nhn dns zone list` |
| Zone 상세 조회 | `nhn dns zone describe <zone-id>` |
| Zone 생성 | `nhn dns zone create --name <fqdn>` |
| Zone 수정 | `nhn dns zone update <zone-id> --description <desc>` |
| Zone 삭제 | `nhn dns zone delete <zone-id>` |
| Record Set 목록 조회 | `nhn dns recordset list --zone-id <id>` |
| Record Set 상세 조회 | `nhn dns recordset describe <rs-id> --zone-id <id>` |
| Record Set 생성 | `nhn dns recordset create --zone-id <id> --name <name> --type <type> --data <data>` |
| Record Set 수정 | `nhn dns recordset update <rs-id> --zone-id <id>` |
| Record Set 삭제 | `nhn dns recordset delete <rs-id> --zone-id <id>` |

### Object Storage

| 기능 | 명령어 |
|------|--------|
| 컨테이너 목록 조회 | `nhn objectstorage container list` (별칭: `nhn os container list`) |
| 컨테이너 메타데이터 조회 | `nhn objectstorage container describe <name>` |
| 컨테이너 생성 | `nhn objectstorage container create <name>` |
| 컨테이너 삭제 | `nhn objectstorage container delete <name>` |
| 오브젝트 목록 조회 | `nhn objectstorage object list --container <name>` |
| 오브젝트 업로드 | `nhn objectstorage object upload --container <name> --file <path>` |
| 오브젝트 다운로드 | `nhn objectstorage object download <object> --container <name>` |
| 오브젝트 삭제 | `nhn objectstorage object delete <object> --container <name>` |
| 오브젝트 메타데이터 조회 | `nhn objectstorage object describe <object> --container <name>` |

---

## 설치

### 빌드 요구사항

- Go 1.25 이상

### 빌드 및 설치

```bash
git clone https://github.com/nhn-cli/nhn.git
cd nhn-cli
go build -o nhn main.go

# Linux/macOS
sudo mv nhn /usr/local/bin/

# Windows
# nhn.exe를 PATH에 포함된 디렉토리로 이동
```

---

## 초기 설정

### 인증 정보 설정

```bash
nhn configure
```

대화형 프롬프트에서 인증 방식을 선택합니다:

**방식 1: OAuth 인증 (권장)**
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
```

**방식 2: Identity 인증**
```
선택 [1]: 2

=== Identity 인증 설정 ===
Tenant ID: your-tenant-id
Username (NHN Cloud ID): your-email@example.com
API Password: your-api-password

=== 리전 설정 ===
기본 리전 [KR1]: KR1
```

### 프로필 관리

```bash
# 프로필 목록 보기
nhn configure list

# 특정 프로필로 설정
nhn configure --profile production
```

---

## 전역 옵션

모든 명령어에서 사용 가능한 옵션:

| 옵션 | 설명 | 기본값 |
|------|------|--------|
| `--profile <name>` | 사용할 프로필 | default |
| `--region <region>` | 리전 지정 | 프로필 설정값 |
| `--output <format>` | 출력 형식 | table |
| `--help` | 도움말 표시 | - |

```bash
# 예시: production 프로필로 KR2 리전의 인스턴스 조회
nhn --profile production --region KR2 compute instance list

# JSON 형식으로 출력
nhn --output json vpc list
```

---

## VPC 명령어

### VPC 관리

```bash
# VPC 목록 조회
nhn vpc list

# VPC 상세 조회
nhn vpc describe <vpc-id>

# VPC 생성
nhn vpc create \
  --name my-vpc \
  --cidr 192.168.0.0/16

# VPC 이름/CIDR 수정
nhn vpc update <vpc-id> \
  --name new-vpc-name \
  --cidr 192.168.0.0/20

# VPC 삭제
nhn vpc delete <vpc-id>
```

### 서브넷 관리

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

### 라우팅 테이블

```bash
# 라우팅 테이블 목록
nhn vpc routingtable list

# 라우팅 테이블 상세 조회
nhn vpc routingtable describe <routingtable-id>

# 라우트 추가
nhn vpc routingtable add-route <routingtable-id> \
  --destination 10.0.0.0/8 \
  --gateway <gateway-id>
```

### 보안 그룹

```bash
# 보안 그룹 목록
nhn vpc securitygroup list

# 보안 그룹 생성
nhn vpc securitygroup create \
  --name my-sg \
  --description "My security group"

# 인바운드 규칙 추가 (SSH)
nhn vpc securitygroup add-rule <sg-id> \
  --direction ingress \
  --protocol tcp \
  --port 22 \
  --remote-ip 0.0.0.0/0

# 인바운드 규칙 추가 (HTTP/HTTPS)
nhn vpc securitygroup add-rule <sg-id> \
  --direction ingress \
  --protocol tcp \
  --port-range 80-443 \
  --remote-ip 0.0.0.0/0

# 보안 그룹 삭제
nhn vpc securitygroup delete <sg-id>
```

### 플로팅 IP

```bash
# 플로팅 IP 목록
nhn vpc floatingip list

# 플로팅 IP 생성
nhn vpc floatingip create

# 인스턴스에 플로팅 IP 연결
nhn vpc floatingip associate <floatingip-id> \
  --instance-id <instance-id>

# 플로팅 IP 연결 해제
nhn vpc floatingip disassociate <floatingip-id>

# 플로팅 IP 삭제
nhn vpc floatingip delete <floatingip-id>
```

### 네트워크 인터페이스

```bash
# 네트워크 인터페이스 목록
nhn vpc port list

# 네트워크 인터페이스 생성
nhn vpc port create \
  --network-id <network-id> \
  --name my-port

# 네트워크 인터페이스 삭제
nhn vpc port delete <port-id>
```

---

## Compute 명령어

### 인스턴스 관리

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

# 인스턴스 시작
nhn compute instance start <instance-id>

# 인스턴스 중지
nhn compute instance stop <instance-id>

# 인스턴스 재부팅 (소프트)
nhn compute instance reboot <instance-id>

# 인스턴스 재부팅 (하드)
nhn compute instance reboot <instance-id> --hard
```

### 인스턴스 타입 (Flavor)

```bash
# 인스턴스 타입 목록
nhn compute flavor list

# 인스턴스 타입 상세 조회
nhn compute flavor describe <flavor-id>
```

### 이미지

```bash
# 이미지 목록
nhn compute image list

# 이미지 상세 조회
nhn compute image describe <image-id>
```

### 키페어

```bash
# 키페어 목록
nhn compute keypair list

# 키페어 생성 (새 키 생성)
nhn compute keypair create --name my-keypair

# 키페어 생성 (공개키 등록)
nhn compute keypair create --name my-keypair \
  --public-key "ssh-rsa AAAA..."

# 키페어 삭제
nhn compute keypair delete my-keypair
```

### 가용성 영역

```bash
# 가용성 영역 목록
nhn compute az list
```

---

## Object Storage 명령어

### 컨테이너 관리

```bash
# 컨테이너 목록 조회
nhn objectstorage container list
nhn os container list  # 별칭

# 컨테이너 메타데이터 조회
nhn os container describe <container-name>

# 컨테이너 생성
nhn os container create my-container

# 컨테이너 삭제
nhn os container delete my-container
```

### 오브젝트 관리

```bash
# 오브젝트 목록 조회
nhn os object list --container my-container

# 파일 업로드
nhn os object upload --container my-container --file ./test.txt

# 커스텀 이름으로 업로드
nhn os object upload --container my-container --file ./test.txt --name custom-name.txt

# 오브젝트 다운로드
nhn os object download test.txt --container my-container

# 저장 경로 지정
nhn os object download test.txt --container my-container --output-file ./downloads/test.txt

# 오브젝트 메타데이터 조회
nhn os object describe test.txt --container my-container

# 오브젝트 삭제
nhn os object delete test.txt --container my-container
```

---

## DNS Plus 명령어

### Zone 관리

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

### Record Set 관리

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

# 다중 레코드 생성
nhn dns recordset create --zone-id <zone-id> \
  --name www.example.com. \
  --type A \
  --ttl 300 \
  --data 1.2.3.4,5.6.7.8

# Record Set 수정
nhn dns recordset update <recordset-id> --zone-id <zone-id> \
  --ttl 600 \
  --data 1.2.3.4

# Record Set 삭제
nhn dns recordset delete <recordset-id> --zone-id <zone-id>
```

---

## 실전 예제

### 예제 1: 기본 인프라 구성

VPC, 서브넷, 보안 그룹을 생성하고 인스턴스를 배포하는 전체 과정:

```bash
# 1. VPC 생성
nhn vpc create --name my-vpc --cidr 192.168.0.0/16

# 2. 서브넷 생성
nhn vpc subnet create \
  --vpc-id <vpc-id> \
  --name public-subnet \
  --cidr 192.168.1.0/24

# 3. 보안 그룹 생성 및 규칙 추가
nhn vpc securitygroup create --name web-sg
nhn vpc securitygroup add-rule <sg-id> \
  --direction ingress --protocol tcp --port 22 --remote-ip 0.0.0.0/0
nhn vpc securitygroup add-rule <sg-id> \
  --direction ingress --protocol tcp --port 80 --remote-ip 0.0.0.0/0

# 4. 키페어 생성
nhn compute keypair create --name my-keypair > my-keypair.pem
chmod 400 my-keypair.pem

# 5. 인스턴스 생성
nhn compute instance create \
  --name web-server \
  --image-id <image-id> \
  --flavor-id <flavor-id> \
  --network-id <network-id> \
  --key-name my-keypair \
  --security-group web-sg

# 6. 플로팅 IP 생성 및 연결
nhn vpc floatingip create
nhn vpc floatingip associate <floatingip-id> --instance-id <instance-id>

# 7. SSH 접속
ssh -i my-keypair.pem centos@<floating-ip>
```

### 예제 2: 다중 환경 관리

```bash
# 개발 환경 설정
nhn configure --profile dev
nhn --profile dev compute instance list

# 운영 환경 설정
nhn configure --profile prod
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
  jq '.[] | {name: .name, ip: .addresses[].addr}'

# 실행 중인 인스턴스만 필터링
nhn --output json compute instance list | \
  jq '.[] | select(.status == "ACTIVE")'

# VPC별 서브넷 개수
nhn --output json vpc list | \
  jq '.[] | {name: .name, subnet_count: (.subnets | length)}'
```

### 예제 4: 인스턴스 일괄 작업

```bash
# 모든 인스턴스 중지
for id in $(nhn --output json compute instance list | jq -r '.[].id'); do
  nhn compute instance stop $id
done

# 특정 태그의 인스턴스만 재부팅
nhn --output json compute instance list | \
  jq -r '.[] | select(.metadata.env == "test") | .id' | \
  xargs -I {} nhn compute instance reboot {}
```

---

## 설정 파일

설정 파일은 `~/.nhn/` 디렉토리에 저장됩니다:

### ~/.nhn/config.json

```json
{
  "profiles": {
    "default": {
      "user_access_key_id": "your-access-key-id",
      "secret_access_key": "your-secret-access-key",
      "region": "KR1"
    },
    "production": {
      "tenant_id": "your-tenant-id",
      "username": "your-email@example.com",
      "password": "your-api-password",
      "region": "KR2"
    }
  }
}
```

### ~/.nhn/credentials.json (자동 생성)

```json
{
  "profiles": {
    "default": {
      "access_token": "cached-token...",
      "expires_at": 1704067200
    }
  }
}
```

---

## API 엔드포인트

### 인증 API

| 서비스 | 엔드포인트 |
|--------|-----------|
| OAuth | `https://oauth.api.nhncloudservice.com` |
| Identity | `https://api-identity-infrastructure.nhncloudservice.com` |

### VPC API (network 타입)

| 리전 | 엔드포인트 |
|------|-----------|
| KR1 | `https://kr1-api-network-infrastructure.nhncloudservice.com` |
| KR2 | `https://kr2-api-network-infrastructure.nhncloudservice.com` |
| JP1 | `https://jp1-api-network-infrastructure.nhncloudservice.com` |

### Compute API

| 리전 | 엔드포인트 |
|------|-----------|
| KR1 | `https://kr1-api-instance-infrastructure.nhncloudservice.com` |
| KR2 | `https://kr2-api-instance-infrastructure.nhncloudservice.com` |
| JP1 | `https://jp1-api-instance-infrastructure.nhncloudservice.com` |

### Block Storage API

| 리전 | 엔드포인트 |
|------|-----------|
| KR1 | `https://kr1-api-block-storage-infrastructure.nhncloudservice.com` |
| KR2 | `https://kr2-api-block-storage-infrastructure.nhncloudservice.com` |
| JP1 | `https://jp1-api-block-storage-infrastructure.nhncloudservice.com` |

### Load Balancer API

Load Balancer API는 VPC API와 동일한 네트워크 엔드포인트를 사용합니다.

| 리전 | 엔드포인트 |
|------|-----------|
| KR1 | `https://kr1-api-network-infrastructure.nhncloudservice.com` |
| KR2 | `https://kr2-api-network-infrastructure.nhncloudservice.com` |
| JP1 | `https://jp1-api-network-infrastructure.nhncloudservice.com` |

### DNS Plus API (글로벌)

| 엔드포인트 |
|-----------|
| `https://dnsplus.api.nhncloudservice.com` |

DNS Plus는 글로벌 서비스로 리전 구분 없이 단일 엔드포인트를 사용합니다.
인증은 AppKey 기반으로, URL 경로에 AppKey를 포함합니다.

### Object Storage API

| 리전 | 엔드포인트 |
|------|-----------|
| KR1 | `https://kr1-api-object-storage.nhncloudservice.com` |
| KR2 | `https://kr2-api-object-storage.nhncloudservice.com` |
| JP1 | `https://jp1-api-object-storage.nhncloudservice.com` |

Base URL 패턴: `https://{region}-api-object-storage.nhncloudservice.com/v1/AUTH_{TenantID}`

---

## API 참조

### VPC API

| 작업 | Method | 경로 |
|------|--------|------|
| VPC 목록 | GET | `/v2.0/vpcs` |
| VPC 조회 | GET | `/v2.0/vpcs/{vpcId}` |
| VPC 생성 | POST | `/v2.0/vpcs` |
| VPC 수정 | PUT | `/v2.0/vpcs/{vpcId}` |
| VPC 삭제 | DELETE | `/v2.0/vpcs/{vpcId}` |
| 서브넷 목록 | GET | `/v2.0/vpcsubnets` |
| 서브넷 조회 | GET | `/v2.0/vpcsubnets/{subnetId}` |
| 서브넷 생성 | POST | `/v2.0/vpcsubnets` |
| 서브넷 삭제 | DELETE | `/v2.0/vpcsubnets/{subnetId}` |
| 보안 그룹 목록 | GET | `/v2.0/security-groups` |
| 보안 그룹 생성 | POST | `/v2.0/security-groups` |
| 보안 그룹 규칙 생성 | POST | `/v2.0/security-group-rules` |
| 플로팅 IP 목록 | GET | `/v2.0/floatingips` |
| 플로팅 IP 생성 | POST | `/v2.0/floatingips` |

### Compute API

| 작업 | Method | 경로 |
|------|--------|------|
| 인스턴스 목록 | GET | `/v2/{tenantId}/servers/detail` |
| 인스턴스 조회 | GET | `/v2/{tenantId}/servers/{serverId}` |
| 인스턴스 생성 | POST | `/v2/{tenantId}/servers` |
| 인스턴스 삭제 | DELETE | `/v2/{tenantId}/servers/{serverId}` |
| 인스턴스 액션 | POST | `/v2/{tenantId}/servers/{serverId}/action` |
| Flavor 목록 | GET | `/v2/{tenantId}/flavors/detail` |
| 이미지 목록 | GET | `/v2/{tenantId}/images/detail` |
| 키페어 목록 | GET | `/v2/{tenantId}/os-keypairs` |
| 키페어 생성 | POST | `/v2/{tenantId}/os-keypairs` |

### Block Storage API

| 작업 | Method | 경로 |
|------|--------|------|
| 볼륨 목록 | GET | `/v2/{tenantId}/volumes/detail` |
| 볼륨 조회 | GET | `/v2/{tenantId}/volumes/{volumeId}` |
| 볼륨 생성 | POST | `/v2/{tenantId}/volumes` |
| 볼륨 삭제 | DELETE | `/v2/{tenantId}/volumes/{volumeId}` |
| 스냅샷 목록 | GET | `/v2/{tenantId}/snapshots/detail` |
| 스냅샷 조회 | GET | `/v2/{tenantId}/snapshots/{snapshotId}` |
| 스냅샷 생성 | POST | `/v2/{tenantId}/snapshots` |
| 스냅샷 삭제 | DELETE | `/v2/{tenantId}/snapshots/{snapshotId}` |
| 볼륨 타입 목록 | GET | `/v2/{tenantId}/types` |

### Load Balancer API

| 작업 | Method | 경로 |
|------|--------|------|
| 로드 밸런서 목록 | GET | `/v2.0/lbaas/loadbalancers` |
| 로드 밸런서 조회 | GET | `/v2.0/lbaas/loadbalancers/{lbId}` |
| 로드 밸런서 생성 | POST | `/v2.0/lbaas/loadbalancers` |
| 로드 밸런서 수정 | PUT | `/v2.0/lbaas/loadbalancers/{lbId}` |
| 로드 밸런서 삭제 | DELETE | `/v2.0/lbaas/loadbalancers/{lbId}` |
| 리스너 목록 | GET | `/v2.0/lbaas/listeners` |
| 리스너 조회 | GET | `/v2.0/lbaas/listeners/{listenerId}` |
| 리스너 생성 | POST | `/v2.0/lbaas/listeners` |
| 리스너 삭제 | DELETE | `/v2.0/lbaas/listeners/{listenerId}` |

### DNS Plus API

| 작업 | Method | 경로 |
|------|--------|------|
| Zone 목록 | GET | `/zones` |
| Zone 조회 | GET | `/zones?zoneIdList={zoneId}` |
| Zone 생성 | POST | `/zones` |
| Zone 수정 | PUT | `/zones/{zoneId}` |
| Zone 삭제 | DELETE | `/zones/async?zoneIdList={zoneId}` |
| Record Set 목록 | GET | `/zones/{zoneId}/recordsets` |
| Record Set 조회 | GET | `/zones/{zoneId}/recordsets?recordsetIdList={rsId}` |
| Record Set 생성 | POST | `/zones/{zoneId}/recordsets` |
| Record Set 수정 | PUT | `/zones/{zoneId}/recordsets/{rsId}` |
| Record Set 삭제 | DELETE | `/zones/{zoneId}/recordsets?recordsetIdList={rsId}` |

Base URL: `https://dnsplus.api.nhncloudservice.com/dnsplus/v1.0/appkeys/{appkey}`

### Object Storage API

| 작업 | Method | 경로 |
|------|--------|------|
| 컨테이너 목록 | GET | `/v1/AUTH_{tenantId}?format=json` |
| 컨테이너 메타데이터 | HEAD | `/v1/AUTH_{tenantId}/{container}` |
| 컨테이너 생성 | PUT | `/v1/AUTH_{tenantId}/{container}` |
| 컨테이너 삭제 | DELETE | `/v1/AUTH_{tenantId}/{container}` |
| 오브젝트 목록 | GET | `/v1/AUTH_{tenantId}/{container}?format=json` |
| 오브젝트 메타데이터 | HEAD | `/v1/AUTH_{tenantId}/{container}/{object}` |
| 오브젝트 업로드 | PUT | `/v1/AUTH_{tenantId}/{container}/{object}` |
| 오브젝트 다운로드 | GET | `/v1/AUTH_{tenantId}/{container}/{object}` |
| 오브젝트 삭제 | DELETE | `/v1/AUTH_{tenantId}/{container}/{object}` |

---

## CI/CD

### GitHub Actions

프로젝트는 GitHub Actions를 통해 자동화된 빌드 및 테스트를 수행합니다.

| 항목 | 설명 |
|------|------|
| 워크플로우 | `.github/workflows/ci.yml` |
| 트리거 | `main` 브랜치 push/PR |
| 플랫폼 | Ubuntu, Windows, macOS |
| Go 버전 | 1.21 |

**실행 단계:**
1. 코드 체크아웃
2. Go 환경 설정 (캐싱 포함)
3. 의존성 다운로드
4. 코드 포맷팅 검사 (`gofmt`, Linux만)
5. 유닛 테스트 실행 (`-race` 플래그)
6. 멀티 플랫폼 빌드

---

## 아키텍처

```
┌─────────────────────────────────────────────────────────────────┐
│                         NHN Cloud CLI                           │
├─────────────────────────────────────────────────────────────────┤
│  ┌──────────┐  ┌──────────┐  ┌──────────┐  ┌──────────────────┐ │
│  │configure │  │  vpc   │  │compute │  │blockstorage│  │loadbalancer│ │
│  └────┬─────┘  └───┬───┘  └───┬────┘  └─────┬──────┘  └─────┬──────┘ │
│       │            │          │              │               │        │
│  ┌────▼────────────▼──────────▼──────────────▼───────────────▼──────┐ │
│  │                    Internal Modules                               │ │
│  │  ┌───────┐ ┌──────┐ ┌─────┐ ┌───────┐ ┌────────────┐ ┌────────┐ │ │
│  │  │config │ │ auth │ │ vpc │ │compute│ │blockstorage│ │  lb    │ │ │
│  │  └───┬───┘ └──┬───┘ └──┬──┘ └───┬───┘ └─────┬──────┘ └───┬────┘ │ │
│  └───────┼────────────┼────────────┼───────────┼──────────────┼─────────┘ │
│          │            │            │           │              │           │
│  ┌───────▼────────────▼────────────▼───────────▼──────────────▼─────────┐ │
│  │                     HTTP Client                             │ │
│  └─────────────────────────┬───────────────────────────────────┘ │
└────────────────────────────┼────────────────────────────────────┘
                             │
                             ▼
┌─────────────────────────────────────────────────────────────────┐
│                      NHN Cloud APIs                             │
│  ┌──────────┐  ┌──────────┐  ┌──────────┐  ┌──────────────────┐ │
│  │ OAuth │ │Identity│ │  VPC  │ │Compute│ │BlockStorage│ │   LB   │ │
│  │  API  │ │  API   │ │  API  │ │  API  │ │    API     │ │  API   │ │
│  └───────┘ └────────┘ └───────┘ └───────┘ └────────────┘ └────────┘ │
└─────────────────────────────────────────────────────────────────┘
```

---

## 리전 정보

| 리전 코드 | 위치 | 설명 |
|-----------|------|------|
| KR1 | 한국 (판교) | 기본 리전 |
| KR2 | 한국 (평촌) | - |
| JP1 | 일본 (도쿄) | - |

---

## NHN Cloud 인증 정보 발급

### OAuth 인증 (권장)

1. [NHN Cloud 콘솔](https://console.nhncloud.com) 로그인
2. 오른쪽 상단의 이메일 주소 클릭
3. **API 보안 설정** 메뉴 선택
4. **User Access Key ID 생성** 버튼 클릭
5. User Access Key ID와 Secret Access Key 발급

### Identity 인증

1. [NHN Cloud 콘솔](https://console.nhncloud.com) 로그인
2. **Compute > Instance** 메뉴 이동
3. **API 엔드포인트 설정** 버튼 클릭
4. Tenant ID 확인 및 API 비밀번호 설정

---

## 문제 해결

### 인증 오류

```bash
# 토큰 캐시 삭제 후 재시도
rm ~/.nhn/credentials.json
nhn compute instance list
```

### 네트워크 오류

```bash
# 리전 엔드포인트 확인
nhn --region KR1 compute instance list

# 디버그 모드
nhn --debug compute instance list
```

### 권한 오류

- Tenant ID가 올바른지 확인
- API 비밀번호가 만료되지 않았는지 확인
- 프로젝트 멤버 권한 확인

---

## 향후 개발 계획

- [x] Block Storage 관리
- [x] Load Balancer 관리
- [x] Object Storage 관리
- [ ] Auto Scale 관리
- [x] DNS 관리
- [ ] 자동완성 지원 (bash, zsh, fish)
- [ ] 설정 파일 암호화
- [ ] 병렬 처리 지원
- [ ] 대화형 모드

---

## 라이선스

PolyForm Noncommercial License 1.0.0

이 소프트웨어는 비상업적 용도로만 사용할 수 있습니다. 자세한 내용은 [LICENSE](LICENSE) 파일을 참조하세요.

---

## 참고 문서

- [NHN Cloud 공식 문서](https://docs.nhncloud.com/)
- [NHN Cloud VPC API](https://docs.nhncloud.com/ko/Network/VPC/ko/public-api/)
- [NHN Cloud Instance API](https://docs.nhncloud.com/ko/Compute/Instance/ko/public-api/)
- [NHN Cloud 인증 API](https://docs.nhncloud.com/ko/nhncloud/ko/public-api/api-authentication/)