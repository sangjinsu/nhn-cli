# 기본 인프라 구성 예제

VPC부터 인스턴스까지 전체 인프라를 구성하는 워크플로우입니다.

---

## 목표

다음 인프라를 구성합니다:

```
┌─────────────────────────────────────────────────────┐
│  VPC: my-vpc (192.168.0.0/16)                       │
│  ┌───────────────────────────────────────────────┐  │
│  │  Subnet: public-subnet (192.168.1.0/24)       │  │
│  │  ┌─────────────────────────────────────────┐  │  │
│  │  │  Instance: web-server                   │  │  │
│  │  │  - Flavor: m2.c1m2                      │  │  │
│  │  │  - Private IP: 192.168.1.10             │  │  │
│  │  │  - Floating IP: 133.186.x.x             │  │  │
│  │  │  - Security Group: web-sg               │  │  │
│  │  │    (SSH:22, HTTP:80, HTTPS:443)         │  │  │
│  │  └─────────────────────────────────────────┘  │  │
│  └───────────────────────────────────────────────┘  │
└─────────────────────────────────────────────────────┘
```

---

## 사전 준비

CLI가 설치되고 인증이 설정되어 있어야 합니다:

```bash
# 설치 확인
nhn version

# 인증 설정
nhn configure
```

---

## 1단계: VPC 생성

```bash
nhn vpc create --name my-vpc --cidr 192.168.0.0/16
```

**출력:**
```
✅ VPC created successfully
ID:     vpc-8a5f3e2c-1234-5678-9abc-def012345678
Name:   my-vpc
CIDR:   192.168.0.0/16
```

VPC ID를 기록합니다: `vpc-8a5f3e2c-...`

---

## 2단계: 서브넷 생성

```bash
nhn vpc subnet create \
  --vpc-id vpc-8a5f3e2c-1234-5678-9abc-def012345678 \
  --name public-subnet \
  --cidr 192.168.1.0/24
```

**출력:**
```
✅ Subnet created successfully
ID:       subnet-aaaaaaaa-1111-2222-3333-444444444444
Name:     public-subnet
CIDR:     192.168.1.0/24
VPC ID:   vpc-8a5f3e2c-...
```

서브넷 ID를 기록합니다: `subnet-aaaaaaaa-...`

---

## 3단계: 보안 그룹 생성

```bash
# 보안 그룹 생성
nhn vpc sg create --name web-sg --description "Web server security group"
```

**출력:**
```
✅ Security group created successfully
ID:   sg-22222222-5555-6666-7777-888888888888
Name: web-sg
```

보안 그룹 ID를 기록합니다: `sg-22222222-...`

### 보안 그룹 규칙 추가

```bash
# SSH 허용 (포트 22)
nhn vpc sg add-rule sg-22222222-5555-6666-7777-888888888888 \
  --direction ingress \
  --protocol tcp \
  --port 22 \
  --remote-ip 0.0.0.0/0

# HTTP 허용 (포트 80)
nhn vpc sg add-rule sg-22222222-5555-6666-7777-888888888888 \
  --direction ingress \
  --protocol tcp \
  --port 80 \
  --remote-ip 0.0.0.0/0

# HTTPS 허용 (포트 443)
nhn vpc sg add-rule sg-22222222-5555-6666-7777-888888888888 \
  --direction ingress \
  --protocol tcp \
  --port 443 \
  --remote-ip 0.0.0.0/0
```

### 규칙 확인

```bash
nhn vpc sg describe sg-22222222-5555-6666-7777-888888888888
```

---

## 4단계: 키페어 생성

```bash
# 키페어 생성 및 개인키 저장
nhn compute keypair create --name my-keypair > my-keypair.pem

# 파일 권한 설정 (Linux/macOS)
chmod 400 my-keypair.pem
```

**확인:**
```bash
nhn compute keypair list
```

---

## 5단계: 이미지 및 Flavor 확인

```bash
# 사용 가능한 이미지 확인
nhn compute image list
```

**출력:**
```
ID                                      NAME                STATUS  SIZE (GB)
image-11111111-...                      Ubuntu 22.04        active  20
image-22222222-...                      CentOS 8            active  20
...
```

Ubuntu 이미지 ID를 기록합니다: `image-11111111-...`

```bash
# 사용 가능한 Flavor 확인
nhn compute flavor list
```

**출력:**
```
ID              NAME        VCPUS   RAM (MB)    DISK (GB)
m2.c1m2         m2.c1m2     1       2048        -
m2.c2m4         m2.c2m4     2       4096        -
...
```

---

## 6단계: 인스턴스 생성

```bash
nhn compute instance create \
  --name web-server \
  --image-id image-11111111-aaaa-bbbb-cccc-dddddddddddd \
  --flavor-id m2.c1m2 \
  --network-id subnet-aaaaaaaa-1111-2222-3333-444444444444 \
  --key-name my-keypair \
  --security-group web-sg \
  --availability-zone kr-pub-a
```

**출력:**
```
✅ Instance created successfully
ID:       instance-99999999-eeee-ffff-0000-111111111111
Name:     web-server
Status:   BUILD
Flavor:   m2.c1m2
```

인스턴스 ID를 기록합니다: `instance-99999999-...`

### 생성 완료 대기

인스턴스가 `ACTIVE` 상태가 될 때까지 기다립니다:

```bash
# 상태 확인
nhn compute instance describe instance-99999999-eeee-ffff-0000-111111111111
```

---

## 7단계: 플로팅 IP 생성 및 연결

```bash
# 플로팅 IP 생성
nhn vpc fip create
```

**출력:**
```
✅ Floating IP created successfully
ID:          fip-44444444-9999-aaaa-bbbb-cccccccccccc
Floating IP: 133.186.xxx.xxx
Status:      DOWN
```

플로팅 IP ID를 기록합니다: `fip-44444444-...`

```bash
# 인스턴스에 연결
nhn vpc fip associate fip-44444444-9999-aaaa-bbbb-cccccccccccc \
  --instance-id instance-99999999-eeee-ffff-0000-111111111111
```

**출력:**
```
✅ Floating IP associated successfully
Floating IP: 133.186.xxx.xxx → instance-99999999-...
```

---

## 8단계: SSH 접속

```bash
ssh -i my-keypair.pem ubuntu@133.186.xxx.xxx
```

> Ubuntu 이미지의 기본 사용자: `ubuntu`
> CentOS 이미지의 기본 사용자: `centos`

---

## 전체 스크립트

모든 단계를 하나의 스크립트로 실행합니다:

```bash
#!/bin/bash
set -e

# 설정
VPC_NAME="my-vpc"
VPC_CIDR="192.168.0.0/16"
SUBNET_NAME="public-subnet"
SUBNET_CIDR="192.168.1.0/24"
SG_NAME="web-sg"
KEYPAIR_NAME="my-keypair"
INSTANCE_NAME="web-server"
FLAVOR_ID="m2.c1m2"
AZ="kr-pub-a"

echo "📦 Creating VPC..."
VPC_OUTPUT=$(nhn --output json vpc create --name $VPC_NAME --cidr $VPC_CIDR)
VPC_ID=$(echo $VPC_OUTPUT | jq -r '.id')
echo "   VPC ID: $VPC_ID"

echo "📦 Creating Subnet..."
SUBNET_OUTPUT=$(nhn --output json vpc subnet create \
  --vpc-id $VPC_ID --name $SUBNET_NAME --cidr $SUBNET_CIDR)
SUBNET_ID=$(echo $SUBNET_OUTPUT | jq -r '.id')
echo "   Subnet ID: $SUBNET_ID"

echo "🔒 Creating Security Group..."
SG_OUTPUT=$(nhn --output json vpc sg create --name $SG_NAME)
SG_ID=$(echo $SG_OUTPUT | jq -r '.id')
echo "   Security Group ID: $SG_ID"

echo "🔒 Adding Security Group Rules..."
nhn vpc sg add-rule $SG_ID --direction ingress --protocol tcp --port 22 --remote-ip 0.0.0.0/0
nhn vpc sg add-rule $SG_ID --direction ingress --protocol tcp --port 80 --remote-ip 0.0.0.0/0
nhn vpc sg add-rule $SG_ID --direction ingress --protocol tcp --port 443 --remote-ip 0.0.0.0/0

echo "🔑 Creating Keypair..."
nhn compute keypair create --name $KEYPAIR_NAME > ${KEYPAIR_NAME}.pem
chmod 400 ${KEYPAIR_NAME}.pem
echo "   Keypair saved to ${KEYPAIR_NAME}.pem"

echo "🔍 Finding Ubuntu image..."
IMAGE_ID=$(nhn --output json compute image list | jq -r '.[] | select(.name | contains("Ubuntu")) | .id' | head -1)
echo "   Image ID: $IMAGE_ID"

echo "🖥️ Creating Instance..."
INSTANCE_OUTPUT=$(nhn --output json compute instance create \
  --name $INSTANCE_NAME \
  --image-id $IMAGE_ID \
  --flavor-id $FLAVOR_ID \
  --network-id $SUBNET_ID \
  --key-name $KEYPAIR_NAME \
  --security-group $SG_NAME \
  --availability-zone $AZ)
INSTANCE_ID=$(echo $INSTANCE_OUTPUT | jq -r '.id')
echo "   Instance ID: $INSTANCE_ID"

echo "⏳ Waiting for instance to be ready..."
sleep 30

echo "🌐 Creating Floating IP..."
FIP_OUTPUT=$(nhn --output json vpc fip create)
FIP_ID=$(echo $FIP_OUTPUT | jq -r '.id')
FIP_ADDR=$(echo $FIP_OUTPUT | jq -r '.floating_ip_address')
echo "   Floating IP: $FIP_ADDR"

echo "🔗 Associating Floating IP..."
nhn vpc fip associate $FIP_ID --instance-id $INSTANCE_ID

echo ""
echo "✅ Infrastructure created successfully!"
echo "   SSH: ssh -i ${KEYPAIR_NAME}.pem ubuntu@${FIP_ADDR}"
```

---

## 리소스 정리

테스트 후 리소스를 정리합니다:

```bash
# 1. 플로팅 IP 연결 해제 및 삭제
nhn vpc fip disassociate <fip-id>
nhn vpc fip delete <fip-id>

# 2. 인스턴스 삭제
nhn compute instance delete <instance-id>

# 3. 키페어 삭제
nhn compute keypair delete my-keypair
rm my-keypair.pem

# 4. 보안 그룹 삭제
nhn vpc sg delete <sg-id>

# 5. 서브넷 삭제
nhn vpc subnet delete <subnet-id>

# 6. VPC 삭제
nhn vpc delete <vpc-id>
```

---

## 참고

- [VPC 명령어](../Commands/VPC.md)
- [Compute 명령어](../Commands/Compute.md)
- [다중 환경 관리](Multi-Environment.md)
