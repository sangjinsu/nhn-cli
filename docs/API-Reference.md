# API 레퍼런스

NHN Cloud CLI가 사용하는 API 엔드포인트 정보입니다.

---

## 인증 API

### OAuth API

| 항목 | 값 |
|------|-----|
| **Base URL** | `https://oauth.api.nhncloudservice.com` |
| **토큰 발급** | `POST /oauth2/token/create` |
| **토큰 유효기간** | 12시간 |

#### 토큰 발급 요청

```http
POST /oauth2/token/create
Content-Type: application/json

{
  "auth": {
    "user_access_key_id": "your-access-key-id",
    "secret_access_key": "your-secret-access-key"
  }
}
```

#### 토큰 발급 응답

```json
{
  "access": {
    "token": {
      "id": "token-string...",
      "expires_at": "2024-01-16T10:30:00Z"
    }
  }
}
```

### Identity API

| 항목 | 값 |
|------|-----|
| **Base URL** | `https://api-identity-infrastructure.nhncloudservice.com` |
| **토큰 발급** | `POST /v2.0/tokens` |
| **토큰 유효기간** | 12시간 |

#### 토큰 발급 요청

```http
POST /v2.0/tokens
Content-Type: application/json

{
  "auth": {
    "tenantId": "your-tenant-id",
    "passwordCredentials": {
      "username": "your-username",
      "password": "your-password"
    }
  }
}
```

---

## VPC API

### 엔드포인트

| 리전 | 엔드포인트 |
|------|-----------|
| KR1 (판교) | `https://kr1-api-network-infrastructure.nhncloudservice.com` |
| KR2 (평촌) | `https://kr2-api-network-infrastructure.nhncloudservice.com` |
| JP1 (도쿄) | `https://jp1-api-network-infrastructure.nhncloudservice.com` |

### VPC 작업

| 작업 | Method | 경로 |
|------|--------|------|
| VPC 목록 | `GET` | `/v2.0/vpcs` |
| VPC 조회 | `GET` | `/v2.0/vpcs/{vpcId}` |
| VPC 생성 | `POST` | `/v2.0/vpcs` |
| VPC 수정 | `PUT` | `/v2.0/vpcs/{vpcId}` |
| VPC 삭제 | `DELETE` | `/v2.0/vpcs/{vpcId}` |

#### VPC 생성 요청

```http
POST /v2.0/vpcs
Content-Type: application/json
X-Auth-Token: {token}

{
  "vpc": {
    "name": "my-vpc",
    "cidrv4": "192.168.0.0/16"
  }
}
```

### 서브넷 작업

| 작업 | Method | 경로 |
|------|--------|------|
| 서브넷 목록 | `GET` | `/v2.0/vpcsubnets` |
| 서브넷 조회 | `GET` | `/v2.0/vpcsubnets/{subnetId}` |
| 서브넷 생성 | `POST` | `/v2.0/vpcsubnets` |
| 서브넷 삭제 | `DELETE` | `/v2.0/vpcsubnets/{subnetId}` |

#### 서브넷 생성 요청

```http
POST /v2.0/vpcsubnets
Content-Type: application/json
X-Auth-Token: {token}

{
  "vpcsubnet": {
    "vpc_id": "vpc-id",
    "name": "my-subnet",
    "cidr": "192.168.1.0/24"
  }
}
```

### 보안 그룹 작업

| 작업 | Method | 경로 |
|------|--------|------|
| 보안 그룹 목록 | `GET` | `/v2.0/security-groups` |
| 보안 그룹 조회 | `GET` | `/v2.0/security-groups/{sgId}` |
| 보안 그룹 생성 | `POST` | `/v2.0/security-groups` |
| 보안 그룹 삭제 | `DELETE` | `/v2.0/security-groups/{sgId}` |
| 규칙 생성 | `POST` | `/v2.0/security-group-rules` |
| 규칙 삭제 | `DELETE` | `/v2.0/security-group-rules/{ruleId}` |

#### 보안 그룹 규칙 생성 요청

```http
POST /v2.0/security-group-rules
Content-Type: application/json
X-Auth-Token: {token}

{
  "security_group_rule": {
    "security_group_id": "sg-id",
    "direction": "ingress",
    "protocol": "tcp",
    "port_range_min": 22,
    "port_range_max": 22,
    "remote_ip_prefix": "0.0.0.0/0"
  }
}
```

### 플로팅 IP 작업

| 작업 | Method | 경로 |
|------|--------|------|
| 플로팅 IP 목록 | `GET` | `/v2.0/floatingips` |
| 플로팅 IP 조회 | `GET` | `/v2.0/floatingips/{fipId}` |
| 플로팅 IP 생성 | `POST` | `/v2.0/floatingips` |
| 플로팅 IP 수정 | `PUT` | `/v2.0/floatingips/{fipId}` |
| 플로팅 IP 삭제 | `DELETE` | `/v2.0/floatingips/{fipId}` |

#### 플로팅 IP 연결 요청

```http
PUT /v2.0/floatingips/{fipId}
Content-Type: application/json
X-Auth-Token: {token}

{
  "floatingip": {
    "port_id": "port-id"
  }
}
```

### 라우팅 테이블 작업

| 작업 | Method | 경로 |
|------|--------|------|
| 라우팅 테이블 목록 | `GET` | `/v2.0/routingtables` |
| 라우팅 테이블 조회 | `GET` | `/v2.0/routingtables/{rtId}` |

### 포트 작업

| 작업 | Method | 경로 |
|------|--------|------|
| 포트 목록 | `GET` | `/v2.0/ports` |
| 포트 조회 | `GET` | `/v2.0/ports/{portId}` |
| 포트 생성 | `POST` | `/v2.0/ports` |
| 포트 삭제 | `DELETE` | `/v2.0/ports/{portId}` |

---

## Compute API

### 엔드포인트

| 리전 | 엔드포인트 |
|------|-----------|
| KR1 (판교) | `https://kr1-api-instance-infrastructure.nhncloudservice.com` |
| KR2 (평촌) | `https://kr2-api-instance-infrastructure.nhncloudservice.com` |
| JP1 (도쿄) | `https://jp1-api-instance-infrastructure.nhncloudservice.com` |

### 인스턴스 작업

| 작업 | Method | 경로 |
|------|--------|------|
| 인스턴스 목록 | `GET` | `/v2/{tenantId}/servers/detail` |
| 인스턴스 조회 | `GET` | `/v2/{tenantId}/servers/{serverId}` |
| 인스턴스 생성 | `POST` | `/v2/{tenantId}/servers` |
| 인스턴스 삭제 | `DELETE` | `/v2/{tenantId}/servers/{serverId}` |
| 인스턴스 액션 | `POST` | `/v2/{tenantId}/servers/{serverId}/action` |

#### 인스턴스 생성 요청

```http
POST /v2/{tenantId}/servers
Content-Type: application/json
X-Auth-Token: {token}

{
  "server": {
    "name": "my-server",
    "imageRef": "image-id",
    "flavorRef": "flavor-id",
    "networks": [
      {
        "uuid": "network-id"
      }
    ],
    "key_name": "my-keypair",
    "security_groups": [
      { "name": "default" }
    ],
    "availability_zone": "kr-pub-a"
  }
}
```

#### 인스턴스 액션 (시작/중지/재부팅)

```http
POST /v2/{tenantId}/servers/{serverId}/action
Content-Type: application/json
X-Auth-Token: {token}

# 시작
{ "os-start": null }

# 중지
{ "os-stop": null }

# 소프트 재부팅
{ "reboot": { "type": "SOFT" } }

# 하드 재부팅
{ "reboot": { "type": "HARD" } }
```

### Flavor 작업

| 작업 | Method | 경로 |
|------|--------|------|
| Flavor 목록 | `GET` | `/v2/{tenantId}/flavors/detail` |
| Flavor 조회 | `GET` | `/v2/{tenantId}/flavors/{flavorId}` |

### 이미지 작업

| 작업 | Method | 경로 |
|------|--------|------|
| 이미지 목록 | `GET` | `/v2/{tenantId}/images/detail` |
| 이미지 조회 | `GET` | `/v2/{tenantId}/images/{imageId}` |

### 키페어 작업

| 작업 | Method | 경로 |
|------|--------|------|
| 키페어 목록 | `GET` | `/v2/{tenantId}/os-keypairs` |
| 키페어 생성 | `POST` | `/v2/{tenantId}/os-keypairs` |
| 키페어 삭제 | `DELETE` | `/v2/{tenantId}/os-keypairs/{keypairName}` |

#### 키페어 생성 요청 (새 키)

```http
POST /v2/{tenantId}/os-keypairs
Content-Type: application/json
X-Auth-Token: {token}

{
  "keypair": {
    "name": "my-keypair"
  }
}
```

#### 키페어 생성 요청 (공개키 등록)

```http
POST /v2/{tenantId}/os-keypairs
Content-Type: application/json
X-Auth-Token: {token}

{
  "keypair": {
    "name": "my-keypair",
    "public_key": "ssh-rsa AAAA..."
  }
}
```

### 가용성 영역 작업

| 작업 | Method | 경로 |
|------|--------|------|
| 가용성 영역 목록 | `GET` | `/v2/{tenantId}/os-availability-zone` |

---

## 공통 헤더

모든 API 요청에 필요한 헤더:

```http
X-Auth-Token: {access_token}
Content-Type: application/json
```

---

## 에러 응답

### 일반 에러 형식

```json
{
  "error": {
    "message": "에러 메시지",
    "code": 400
  }
}
```

### 주요 상태 코드

| 코드 | 의미 |
|------|------|
| 200 | 성공 |
| 201 | 생성 성공 |
| 204 | 삭제 성공 |
| 400 | 잘못된 요청 |
| 401 | 인증 실패 |
| 403 | 권한 없음 |
| 404 | 리소스 없음 |
| 409 | 충돌 (중복 등) |
| 500 | 서버 오류 |

---

## 리전 정보

| 리전 코드 | 위치 | 설명 |
|-----------|------|------|
| KR1 | 한국 (판교) | 기본 리전 |
| KR2 | 한국 (평촌) | - |
| JP1 | 일본 (도쿄) | - |

---

## 참고 문서

- [NHN Cloud VPC API](https://docs.nhncloud.com/ko/Network/VPC/ko/public-api/)
- [NHN Cloud Instance API](https://docs.nhncloud.com/ko/Compute/Instance/ko/public-api/)
- [NHN Cloud 인증 API](https://docs.nhncloud.com/ko/nhncloud/ko/public-api/api-authentication/)
