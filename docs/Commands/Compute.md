# Compute 명령어

NHN Cloud Compute 리소스 관리 명령어입니다.

---

## 개요

```bash
nhn compute <subcommand> [options]
```

### 서브 명령어

| 명령어 | 설명 |
|--------|------|
| `instance` | 인스턴스 관리 |
| `flavor` | 인스턴스 타입 조회 |
| `image` | 이미지 조회 |
| `keypair` | 키페어 관리 |
| `az` | 가용성 영역 조회 |

---

## 인스턴스 관리

### 인스턴스 목록 조회

```bash
nhn compute instance list
```

**출력 예시:**
```
ID                                      NAME        STATUS  FLAVOR      IP ADDRESSES    AZ
a1b2c3d4-5678-9abc-def0-123456789abc    web-server  ACTIVE  m2.c1m2     192.168.1.10    kr-pub-a
b2c3d4e5-6789-abcd-ef01-23456789abcd    db-server   ACTIVE  m2.c2m4     192.168.2.20    kr-pub-b
```

### 인스턴스 상세 조회

```bash
nhn compute instance describe <instance-id>
```

**출력 예시:**
```
ID:               a1b2c3d4-5678-9abc-def0-123456789abc
Name:             web-server
Status:           ACTIVE
Flavor:           m2.c1m2 (1 vCPU, 2GB RAM)
Image:            Ubuntu 22.04 (image-xxxxxxxx)
Key Name:         my-keypair
Availability Zone: kr-pub-a
Created:          2024-01-15T10:30:00Z

Networks:
  subnet-aaaaaaaa: 192.168.1.10 (MAC: fa:16:3e:xx:xx:xx)

Security Groups:
  - default
  - web-sg

Volumes:
  - vol-xxxxxxxx (20GB, /dev/vda)
```

### 인스턴스 생성

```bash
nhn compute instance create [options]
```

**옵션:**

| 옵션 | 설명 | 필수 |
|------|------|------|
| `--name` | 인스턴스 이름 | O |
| `--image-id` | 이미지 ID | O |
| `--flavor-id` | 인스턴스 타입 ID | O |
| `--network-id` | 네트워크 (서브넷) ID | O |
| `--key-name` | SSH 키페어 이름 | O |
| `--security-group` | 보안 그룹 (여러 개 지정 가능) | - |
| `--availability-zone` | 가용성 영역 | - |
| `--user-data` | 사용자 데이터 (cloud-init) | - |

**예시:**
```bash
nhn compute instance create \
  --name web-server \
  --image-id image-12345678 \
  --flavor-id m2.c1m2 \
  --network-id subnet-aaaaaaaa \
  --key-name my-keypair \
  --security-group default \
  --security-group web-sg \
  --availability-zone kr-pub-a
```

**사용자 데이터 예시:**
```bash
nhn compute instance create \
  --name web-server \
  --image-id image-12345678 \
  --flavor-id m2.c1m2 \
  --network-id subnet-aaaaaaaa \
  --key-name my-keypair \
  --user-data '#!/bin/bash
apt-get update
apt-get install -y nginx'
```

### 인스턴스 삭제

```bash
nhn compute instance delete <instance-id>
```

### 인스턴스 시작

중지된 인스턴스를 시작합니다.

```bash
nhn compute instance start <instance-id>
```

### 인스턴스 중지

실행 중인 인스턴스를 중지합니다.

```bash
nhn compute instance stop <instance-id>
```

### 인스턴스 재부팅

```bash
# 소프트 재부팅 (OS 레벨)
nhn compute instance reboot <instance-id>

# 하드 재부팅 (전원 재시작)
nhn compute instance reboot <instance-id> --hard
```

**옵션:**

| 옵션 | 설명 |
|------|------|
| `--hard` | 하드 재부팅 수행 |

---

## 인스턴스 타입 (Flavor)

### Flavor 목록 조회

```bash
nhn compute flavor list
```

**출력 예시:**
```
ID              NAME        VCPUS   RAM (MB)    DISK (GB)   DESCRIPTION
m2.c1m2         m2.c1m2     1       2048        -           1 vCPU, 2GB RAM
m2.c2m4         m2.c2m4     2       4096        -           2 vCPU, 4GB RAM
m2.c4m8         m2.c4m8     4       8192        -           4 vCPU, 8GB RAM
m2.c8m16        m2.c8m16    8       16384       -           8 vCPU, 16GB RAM
```

### Flavor 상세 조회

```bash
nhn compute flavor describe <flavor-id>
```

**출력 예시:**
```
ID:           m2.c4m8
Name:         m2.c4m8
vCPUs:        4
RAM:          8192 MB
Disk:         0 GB (별도 볼륨 사용)
Is Public:    true
Description:  4 vCPU, 8GB RAM
```

---

## 이미지

### 이미지 목록 조회

```bash
nhn compute image list
```

**출력 예시:**
```
ID                                      NAME                            STATUS  SIZE (GB)   CREATED
image-11111111-...                      Ubuntu 22.04                    active  20          2024-01-01
image-22222222-...                      CentOS 8                        active  20          2024-01-01
image-33333333-...                      Windows Server 2022             active  50          2024-01-01
image-44444444-...                      Rocky Linux 9                   active  20          2024-01-01
```

### 이미지 상세 조회

```bash
nhn compute image describe <image-id>
```

**출력 예시:**
```
ID:             image-11111111-...
Name:           Ubuntu 22.04
Status:         active
Size:           20 GB
Min Disk:       20 GB
Min RAM:        1024 MB
OS Distro:      ubuntu
OS Version:     22.04
Created:        2024-01-01T00:00:00Z
Updated:        2024-01-15T00:00:00Z

Properties:
  architecture: x86_64
  hypervisor_type: qemu
  os_type: linux
```

---

## 키페어

### 키페어 목록 조회

```bash
nhn compute keypair list
```

**출력 예시:**
```
NAME            FINGERPRINT                                     TYPE
my-keypair      ab:cd:ef:12:34:56:78:90:ab:cd:ef:12:34:56:78:90  ssh
prod-keypair    12:34:56:78:90:ab:cd:ef:12:34:56:78:90:ab:cd:ef  ssh
```

### 키페어 생성 (새 키 생성)

새로운 SSH 키를 생성하고 개인키를 출력합니다.

```bash
nhn compute keypair create --name <name>
```

**예시:**
```bash
# 개인키를 파일로 저장
nhn compute keypair create --name my-keypair > my-keypair.pem
chmod 400 my-keypair.pem
```

**출력 (개인키):**
```
-----BEGIN RSA PRIVATE KEY-----
MIIEpAIBAAKCAQEA...
...
-----END RSA PRIVATE KEY-----
```

### 키페어 생성 (기존 공개키 등록)

기존 SSH 공개키를 등록합니다.

```bash
nhn compute keypair create --name <name> --public-key <public-key>
```

**예시:**
```bash
# 공개키 직접 입력
nhn compute keypair create --name my-keypair \
  --public-key "ssh-rsa AAAAB3NzaC1yc2E... user@host"

# 파일에서 읽기
nhn compute keypair create --name my-keypair \
  --public-key "$(cat ~/.ssh/id_rsa.pub)"
```

### 키페어 삭제

```bash
nhn compute keypair delete <keypair-name>
```

---

## 가용성 영역

### 가용성 영역 목록 조회

```bash
nhn compute az list
```

**출력 예시:**
```
NAME            STATE       HOSTS
kr-pub-a        available   -
kr-pub-b        available   -
```

---

## 인스턴스 상태

| 상태 | 설명 |
|------|------|
| `BUILD` | 인스턴스 생성 중 |
| `ACTIVE` | 실행 중 |
| `SHUTOFF` | 중지됨 |
| `PAUSED` | 일시 중지됨 |
| `SUSPENDED` | 정지됨 |
| `REBOOT` | 재부팅 중 |
| `ERROR` | 오류 상태 |

---

## 일반적인 워크플로우

### 인스턴스 생성 전 정보 수집

```bash
# 1. 사용 가능한 이미지 확인
nhn compute image list

# 2. 인스턴스 타입 확인
nhn compute flavor list

# 3. 가용성 영역 확인
nhn compute az list

# 4. 네트워크 (서브넷) 확인
nhn vpc subnet list

# 5. 보안 그룹 확인
nhn vpc sg list

# 6. 키페어 확인 (없으면 생성)
nhn compute keypair list
```

### 인스턴스 생성 및 접속

```bash
# 1. 키페어 생성
nhn compute keypair create --name my-keypair > my-keypair.pem
chmod 400 my-keypair.pem

# 2. 인스턴스 생성
nhn compute instance create \
  --name web-server \
  --image-id <ubuntu-image-id> \
  --flavor-id m2.c1m2 \
  --network-id <subnet-id> \
  --key-name my-keypair \
  --security-group web-sg

# 3. 플로팅 IP 연결
nhn vpc fip create
nhn vpc fip associate <fip-id> --instance-id <instance-id>

# 4. SSH 접속
ssh -i my-keypair.pem ubuntu@<floating-ip>
```

---

## 참고

- [전역 옵션](Global-Options.md)
- [VPC 명령어](VPC.md)
- [기본 인프라 구성 예제](../Examples/Basic-Infrastructure.md)
