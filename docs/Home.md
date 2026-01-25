# NHN Cloud CLI Wiki

AWS CLI 스타일의 NHN Cloud 명령줄 인터페이스 문서입니다.

---

## 소개

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

## 주요 기능

| 카테고리 | 기능 |
|----------|------|
| **인증** | OAuth 인증, Identity 인증, 토큰 캐싱, 다중 프로필 |
| **VPC** | VPC, 서브넷, 보안 그룹, 플로팅 IP, 라우팅 테이블, 포트 관리 |
| **Compute** | 인스턴스, Flavor, 이미지, 키페어, 가용성 영역 관리 |

---

## 문서 목차

### 시작하기

- [빠른 시작 가이드](Getting-Started.md) - 5분 안에 CLI 설정하기
- [설치 가이드](Installation.md) - 상세 설치 방법
- [설정 가이드](Configuration.md) - 인증 및 프로필 설정

### 명령어 참조

- [VPC 명령어](Commands/VPC.md) - VPC, 서브넷, 보안 그룹 등
- [Compute 명령어](Commands/Compute.md) - 인스턴스, 이미지, 키페어 등
- [전역 옵션](Commands/Global-Options.md) - 모든 명령어 공통 옵션

### 실전 예제

- [기본 인프라 구성](Examples/Basic-Infrastructure.md) - VPC에서 인스턴스까지
- [다중 환경 관리](Examples/Multi-Environment.md) - 개발/운영 환경 분리
- [자동화 스크립트](Examples/Automation-Scripts.md) - JSON 출력과 jq 활용

### 참조

- [아키텍처](Architecture.md) - 시스템 구조 및 설계
- [API 레퍼런스](API-Reference.md) - API 엔드포인트 정보
- [문제 해결](Troubleshooting.md) - FAQ 및 오류 해결
- [기여 가이드](Contributing.md) - 프로젝트 기여 방법

---

## 지원 리전

| 리전 코드 | 위치 | 설명 |
|-----------|------|------|
| KR1 | 한국 (판교) | 기본 리전 |
| KR2 | 한국 (평촌) | - |
| JP1 | 일본 (도쿄) | - |

---

## 빠른 시작

```bash
# 1. 빌드 및 설치
git clone https://github.com/sangjinsu/nhn-cli.git
cd nhn-cli
go build -o nhn main.go
sudo mv nhn /usr/local/bin/

# 2. 인증 설정
nhn configure

# 3. 리소스 조회
nhn vpc list
nhn compute instance list
```

자세한 내용은 [빠른 시작 가이드](Getting-Started.md)를 참조하세요.

---

## 참고 문서

- [NHN Cloud 공식 문서](https://docs.nhncloud.com/)
- [NHN Cloud VPC API](https://docs.nhncloud.com/ko/Network/VPC/ko/public-api/)
- [NHN Cloud Instance API](https://docs.nhncloud.com/ko/Compute/Instance/ko/public-api/)
- [NHN Cloud 인증 API](https://docs.nhncloud.com/ko/nhncloud/ko/public-api/api-authentication/)

---

## 라이선스

MIT License
