# Load Balancer 명령어

NHN Cloud Load Balancer 리소스 관리 명령어입니다.

---

## 개요

```bash
nhn loadbalancer <subcommand> [options]
# 또는 별칭
nhn lb <subcommand> [options]
```

### 서브 명령어

| 명령어 | 설명 |
|--------|------|
| `list` | 로드 밸런서 목록 조회 |
| `describe` | 로드 밸런서 상세 조회 |
| `create` | 로드 밸런서 생성 |
| `update` | 로드 밸런서 수정 |
| `delete` | 로드 밸런서 삭제 |
| `listener` | 리스너 관리 |

---

## 로드 밸런서 관리

### 로드 밸런서 목록 조회

```bash
nhn lb list
```

**출력 예시:**
```
ID                                      NAME            STATUS              VIP ADDRESS     TYPE
lb-11111111-...                         web-lb          ACTIVE              192.168.1.100   shared
lb-22222222-...                         api-lb          ACTIVE              192.168.2.100   dedicated
```

### 로드 밸런서 상세 조회

```bash
nhn lb describe <lb-id>
```

**옵션:**

| 옵션 | 설명 |
|------|------|
| `<lb-id>` | 로드 밸런서 ID (필수) |

### 로드 밸런서 생성

```bash
nhn lb create [options]
```

**옵션:**

| 옵션 | 설명 | 필수 |
|------|------|------|
| `--vip-subnet-id` | VIP 서브넷 ID | O |
| `--name` | 로드 밸런서 이름 | - |
| `--vip-address` | VIP 주소 | - |
| `--description` | 설명 | - |
| `--type` | 타입 (shared/dedicated) | - |

**예시:**
```bash
# 기본 로드 밸런서 생성
nhn lb create \
  --vip-subnet-id subnet-aaaaaaaa \
  --name web-lb

# 전용 로드 밸런서 생성
nhn lb create \
  --vip-subnet-id subnet-aaaaaaaa \
  --name api-lb \
  --type dedicated \
  --description "API load balancer"

# VIP 주소 지정
nhn lb create \
  --vip-subnet-id subnet-aaaaaaaa \
  --name web-lb \
  --vip-address 192.168.1.100
```

### 로드 밸런서 수정

```bash
nhn lb update <lb-id> [options]
```

**옵션:**

| 옵션 | 설명 |
|------|------|
| `--name` | 로드 밸런서 이름 |
| `--description` | 설명 |

**예시:**
```bash
nhn lb update lb-11111111 --name new-lb-name
nhn lb update lb-11111111 --description "Updated description"
```

### 로드 밸런서 삭제

```bash
nhn lb delete <lb-id>
```

> 로드 밸런서 삭제 전 모든 리스너를 먼저 삭제해야 합니다.

---

## 리스너 관리

### 리스너 목록 조회

```bash
nhn lb listener list
```

**출력 예시:**
```
ID                                      NAME            PROTOCOL    PORT    LB ID
listener-11111111-...                   http-listener   HTTP        80      lb-11111111-...
listener-22222222-...                   https-listener  HTTPS       443     lb-11111111-...
```

### 리스너 상세 조회

```bash
nhn lb listener describe <listener-id>
```

**옵션:**

| 옵션 | 설명 |
|------|------|
| `<listener-id>` | 리스너 ID (필수) |

### 리스너 생성

```bash
nhn lb listener create [options]
```

**옵션:**

| 옵션 | 설명 | 필수 |
|------|------|------|
| `--loadbalancer-id` | 로드 밸런서 ID | O |
| `--protocol` | 프로토콜 (TCP/HTTP/HTTPS) | O |
| `--port` | 포트 번호 | O |
| `--name` | 리스너 이름 | - |
| `--description` | 설명 | - |
| `--default-pool-id` | 기본 풀 ID | - |

**예시:**
```bash
# HTTP 리스너 생성
nhn lb listener create \
  --loadbalancer-id lb-11111111 \
  --protocol HTTP \
  --port 80 \
  --name http-listener

# HTTPS 리스너 생성
nhn lb listener create \
  --loadbalancer-id lb-11111111 \
  --protocol HTTPS \
  --port 443 \
  --name https-listener

# TCP 리스너 생성
nhn lb listener create \
  --loadbalancer-id lb-11111111 \
  --protocol TCP \
  --port 3306 \
  --name db-listener
```

### 리스너 삭제

```bash
nhn lb listener delete <listener-id>
```

---

## 참고

- [전역 옵션](Global-Options.md)
- [VPC 명령어](VPC.md)
- [Block Storage 명령어](BlockStorage.md)
