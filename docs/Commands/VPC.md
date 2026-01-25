# VPC 명령어

NHN Cloud VPC 리소스 관리 명령어입니다.

---

## 개요

```bash
nhn vpc <subcommand> [options]
```

### 서브 명령어

| 명령어 | 설명 |
|--------|------|
| `list` | VPC 목록 조회 |
| `describe` | VPC 상세 조회 |
| `create` | VPC 생성 |
| `update` | VPC 수정 |
| `delete` | VPC 삭제 |
| `subnet` | 서브넷 관리 |
| `securitygroup`, `sg` | 보안 그룹 관리 |
| `floatingip`, `fip` | 플로팅 IP 관리 |
| `routingtable`, `rt` | 라우팅 테이블 조회 |
| `port` | 네트워크 인터페이스 관리 |

---

## VPC 관리

### VPC 목록 조회

```bash
nhn vpc list
```

**출력 예시:**
```
ID                                      NAME            CIDR            STATE
8a5f3e2c-1234-5678-9abc-def012345678    my-vpc          192.168.0.0/16  available
b2c4d6e8-2345-6789-abcd-ef0123456789    prod-vpc        10.0.0.0/16     available
```

### VPC 상세 조회

```bash
nhn vpc describe <vpc-id>
```

**옵션:**

| 옵션 | 설명 |
|------|------|
| `<vpc-id>` | VPC ID (필수) |

**출력 예시:**
```
ID:           8a5f3e2c-1234-5678-9abc-def012345678
Name:         my-vpc
CIDR:         192.168.0.0/16
State:        available
Tenant ID:    tenant-12345
Created:      2024-01-15T10:30:00Z

Subnets:
  - subnet-aaaaaaaa (192.168.1.0/24)
  - subnet-bbbbbbbb (192.168.2.0/24)
```

### VPC 생성

```bash
nhn vpc create --name <name> --cidr <cidr>
```

**옵션:**

| 옵션 | 설명 | 필수 |
|------|------|------|
| `--name` | VPC 이름 | O |
| `--cidr` | CIDR 블록 (예: 192.168.0.0/16) | O |

**예시:**
```bash
nhn vpc create --name my-vpc --cidr 192.168.0.0/16
```

### VPC 수정

```bash
nhn vpc update <vpc-id> [options]
```

**옵션:**

| 옵션 | 설명 |
|------|------|
| `--name` | 새 VPC 이름 |
| `--cidr` | 새 CIDR 블록 |

**예시:**
```bash
nhn vpc update 8a5f3e2c-... --name new-vpc-name
nhn vpc update 8a5f3e2c-... --cidr 192.168.0.0/20
```

### VPC 삭제

```bash
nhn vpc delete <vpc-id>
```

> VPC 삭제 전 모든 서브넷, 보안 그룹 등 종속 리소스를 먼저 삭제해야 합니다.

---

## 서브넷 관리

### 서브넷 목록 조회

```bash
nhn vpc subnet list [options]
```

**옵션:**

| 옵션 | 설명 |
|------|------|
| `--vpc-id` | 특정 VPC의 서브넷만 필터링 |

**예시:**
```bash
# 모든 서브넷 조회
nhn vpc subnet list

# 특정 VPC의 서브넷만 조회
nhn vpc subnet list --vpc-id 8a5f3e2c-...
```

**출력 예시:**
```
ID                                      NAME            CIDR            VPC ID                                  STATE
subnet-aaaaaaaa-...                     public-subnet   192.168.1.0/24  8a5f3e2c-...                            available
subnet-bbbbbbbb-...                     private-subnet  192.168.2.0/24  8a5f3e2c-...                            available
```

### 서브넷 상세 조회

```bash
nhn vpc subnet describe <subnet-id>
```

### 서브넷 생성

```bash
nhn vpc subnet create --vpc-id <vpc-id> --name <name> --cidr <cidr>
```

**옵션:**

| 옵션 | 설명 | 필수 |
|------|------|------|
| `--vpc-id` | VPC ID | O |
| `--name` | 서브넷 이름 | O |
| `--cidr` | CIDR 블록 (VPC CIDR 범위 내) | O |

**예시:**
```bash
nhn vpc subnet create \
  --vpc-id 8a5f3e2c-... \
  --name public-subnet \
  --cidr 192.168.1.0/24
```

### 서브넷 삭제

```bash
nhn vpc subnet delete <subnet-id>
```

---

## 보안 그룹 관리

### 보안 그룹 목록 조회

```bash
nhn vpc securitygroup list
# 또는 별칭
nhn vpc sg list
```

**출력 예시:**
```
ID                                      NAME        DESCRIPTION                         RULES
sg-11111111-...                         default     default security group              5
sg-22222222-...                         web-sg      Web server security group           8
```

### 보안 그룹 상세 조회

```bash
nhn vpc sg describe <sg-id>
```

**출력 예시:**
```
ID:           sg-22222222-...
Name:         web-sg
Description:  Web server security group
Tenant ID:    tenant-12345

Security Group Rules:
DIRECTION   PROTOCOL  PORT RANGE   REMOTE IP/GROUP        DESCRIPTION
ingress     tcp       22           0.0.0.0/0              SSH access
ingress     tcp       80           0.0.0.0/0              HTTP access
ingress     tcp       443          0.0.0.0/0              HTTPS access
egress      any       any          0.0.0.0/0              Allow all outbound
```

### 보안 그룹 생성

```bash
nhn vpc sg create --name <name> [--description <desc>]
```

**옵션:**

| 옵션 | 설명 | 필수 |
|------|------|------|
| `--name` | 보안 그룹 이름 | O |
| `--description` | 설명 | - |

**예시:**
```bash
nhn vpc sg create --name web-sg --description "Web server security group"
```

### 보안 그룹 규칙 추가

```bash
nhn vpc sg add-rule <sg-id> [options]
```

**옵션:**

| 옵션 | 설명 | 필수 |
|------|------|------|
| `--direction` | 방향 (`ingress` 또는 `egress`) | O |
| `--protocol` | 프로토콜 (`tcp`, `udp`, `icmp`, `any`) | O |
| `--port` | 단일 포트 번호 | - |
| `--port-range` | 포트 범위 (예: `80-443`) | - |
| `--remote-ip` | 원격 IP CIDR (예: `0.0.0.0/0`) | - |
| `--remote-group` | 원격 보안 그룹 ID | - |
| `--description` | 규칙 설명 | - |

**예시:**
```bash
# SSH 허용
nhn vpc sg add-rule sg-22222222-... \
  --direction ingress \
  --protocol tcp \
  --port 22 \
  --remote-ip 0.0.0.0/0

# HTTP/HTTPS 포트 범위 허용
nhn vpc sg add-rule sg-22222222-... \
  --direction ingress \
  --protocol tcp \
  --port-range 80-443 \
  --remote-ip 0.0.0.0/0

# ICMP (ping) 허용
nhn vpc sg add-rule sg-22222222-... \
  --direction ingress \
  --protocol icmp \
  --remote-ip 0.0.0.0/0

# 다른 보안 그룹에서의 접근 허용
nhn vpc sg add-rule sg-22222222-... \
  --direction ingress \
  --protocol tcp \
  --port 3306 \
  --remote-group sg-33333333-...
```

### 보안 그룹 규칙 삭제

```bash
nhn vpc sg delete-rule <rule-id>
```

### 보안 그룹 삭제

```bash
nhn vpc sg delete <sg-id>
```

---

## 플로팅 IP 관리

### 플로팅 IP 목록 조회

```bash
nhn vpc floatingip list
# 또는 별칭
nhn vpc fip list
```

**출력 예시:**
```
ID                                      FLOATING IP     STATUS  INSTANCE ID                             FIXED IP
fip-11111111-...                        133.186.x.x     ACTIVE  instance-99999999-...                   192.168.1.10
fip-22222222-...                        133.186.x.y     DOWN    -                                       -
```

### 플로팅 IP 생성

```bash
nhn vpc fip create
```

### 플로팅 IP 연결

```bash
nhn vpc fip associate <floatingip-id> --instance-id <instance-id>
# 또는 포트에 직접 연결
nhn vpc fip associate <floatingip-id> --port-id <port-id>
```

**옵션:**

| 옵션 | 설명 |
|------|------|
| `--instance-id` | 인스턴스 ID |
| `--port-id` | 포트 ID (네트워크 인터페이스) |

### 플로팅 IP 연결 해제

```bash
nhn vpc fip disassociate <floatingip-id>
```

### 플로팅 IP 삭제

```bash
nhn vpc fip delete <floatingip-id>
```

---

## 라우팅 테이블 조회

### 라우팅 테이블 목록

```bash
nhn vpc routingtable list
# 또는 별칭
nhn vpc rt list
```

**출력 예시:**
```
ID                                      NAME            VPC ID                                  DEFAULT
rt-11111111-...                         default-rt      8a5f3e2c-...                            true
rt-22222222-...                         custom-rt       8a5f3e2c-...                            false
```

### 라우팅 테이블 상세 조회

```bash
nhn vpc rt describe <routingtable-id>
```

**출력 예시:**
```
ID:           rt-11111111-...
Name:         default-rt
VPC ID:       8a5f3e2c-...
Default:      true

Routes:
DESTINATION         TARGET                  TYPE
0.0.0.0/0           igw-xxxxxxxx-...        internet_gateway
192.168.0.0/16      local                   local
10.0.0.0/8          vpn-xxxxxxxx-...        vpn_gateway
```

---

## 네트워크 인터페이스 (포트) 관리

### 포트 목록 조회

```bash
nhn vpc port list
```

**출력 예시:**
```
ID                                      NAME            STATUS  MAC ADDRESS         IP ADDRESS      NETWORK ID
port-11111111-...                       my-port         ACTIVE  fa:16:3e:xx:xx:xx   192.168.1.10    subnet-aaaaaaaa-...
port-22222222-...                       db-port         ACTIVE  fa:16:3e:yy:yy:yy   192.168.2.20    subnet-bbbbbbbb-...
```

### 포트 상세 조회

```bash
nhn vpc port describe <port-id>
```

### 포트 생성

```bash
nhn vpc port create --network-id <network-id> --name <name>
```

**옵션:**

| 옵션 | 설명 | 필수 |
|------|------|------|
| `--network-id` | 네트워크 (서브넷) ID | O |
| `--name` | 포트 이름 | - |

### 포트 삭제

```bash
nhn vpc port delete <port-id>
```

---

## 참고

- [전역 옵션](Global-Options.md)
- [Compute 명령어](Compute.md)
- [기본 인프라 구성 예제](../Examples/Basic-Infrastructure.md)
