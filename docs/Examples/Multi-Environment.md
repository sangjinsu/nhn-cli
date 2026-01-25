# 다중 환경 관리 예제

개발, 스테이징, 운영 등 여러 환경을 CLI 프로필로 관리하는 방법입니다.

---

## 시나리오

다음과 같은 환경을 관리합니다:

| 환경 | 프로필 | 리전 | 용도 |
|------|--------|------|------|
| 개발 | `dev` | KR1 | 개발 및 테스트 |
| 스테이징 | `staging` | KR1 | QA 및 사전 검증 |
| 운영 | `prod` | KR2 | 실제 서비스 |

---

## 1. 프로필 설정

### 개발 환경 프로필

```bash
nhn configure --profile dev
```

```
프로필 이름 [dev]:
=== 인증 방식 선택 ===
선택 [1]: 1

=== OAuth 인증 설정 ===
User Access Key ID: dev-access-key-id
Secret Access Key: dev-secret-access-key

=== 리전 설정 ===
기본 리전 [KR1]: KR1
```

### 스테이징 환경 프로필

```bash
nhn configure --profile staging
```

### 운영 환경 프로필

```bash
nhn configure --profile prod
```

```
=== 리전 설정 ===
기본 리전 [KR1]: KR2
```

### 프로필 확인

```bash
nhn configure list
```

**출력:**
```
PROFILE     AUTH TYPE   REGION
default     oauth       KR1
dev         oauth       KR1
staging     oauth       KR1
prod        oauth       KR2
```

---

## 2. 환경별 리소스 조회

### 특정 환경의 리소스 조회

```bash
# 개발 환경 VPC
nhn --profile dev vpc list

# 스테이징 환경 인스턴스
nhn --profile staging compute instance list

# 운영 환경 보안 그룹
nhn --profile prod vpc sg list
```

### 모든 환경 비교

```bash
# 환경별 인스턴스 수 확인
for env in dev staging prod; do
  echo "=== $env 환경 ==="
  nhn --profile $env compute instance list
  echo ""
done
```

---

## 3. 환경 변수 활용

### 기본 프로필 설정

```bash
# 기본 프로필을 dev로 설정
export NHN_PROFILE=dev

# 이후 명령은 dev 프로필 사용
nhn vpc list
nhn compute instance list
```

### 셸 설정 파일에 추가

```bash
# ~/.bashrc 또는 ~/.zshrc에 추가
echo 'export NHN_PROFILE=dev' >> ~/.zshrc
source ~/.zshrc
```

### 일시적 변경

```bash
# 일회성으로 prod 사용
NHN_PROFILE=prod nhn vpc list

# 또는 명령줄 옵션 사용
nhn --profile prod vpc list
```

---

## 4. 환경 관리 스크립트

### 환경별 리소스 목록 스크립트

```bash
#!/bin/bash
# list-all-resources.sh

PROFILES=("dev" "staging" "prod")

for profile in "${PROFILES[@]}"; do
  echo "============================================"
  echo "환경: $profile"
  echo "============================================"

  echo ""
  echo "📦 VPCs:"
  nhn --profile $profile vpc list

  echo ""
  echo "🖥️ Instances:"
  nhn --profile $profile compute instance list

  echo ""
  echo "🔒 Security Groups:"
  nhn --profile $profile vpc sg list

  echo ""
done
```

### 환경별 인스턴스 상태 확인

```bash
#!/bin/bash
# check-instance-status.sh

echo "인스턴스 상태 요약"
echo "=================="

for profile in dev staging prod; do
  total=$(nhn --profile $profile --output json compute instance list | jq 'length')
  active=$(nhn --profile $profile --output json compute instance list | jq '[.[] | select(.status == "ACTIVE")] | length')
  shutoff=$(nhn --profile $profile --output json compute instance list | jq '[.[] | select(.status == "SHUTOFF")] | length')

  echo "$profile: 전체 $total / 실행 $active / 중지 $shutoff"
done
```

### 환경 전환 함수

```bash
# ~/.bashrc 또는 ~/.zshrc에 추가
nhn-env() {
  case $1 in
    dev|staging|prod)
      export NHN_PROFILE=$1
      echo "✅ NHN 환경이 '$1'(으)로 변경되었습니다."
      ;;
    *)
      echo "사용법: nhn-env [dev|staging|prod]"
      echo "현재: ${NHN_PROFILE:-default}"
      ;;
  esac
}
```

사용:
```bash
nhn-env dev      # dev 환경으로 전환
nhn-env prod     # prod 환경으로 전환
nhn-env          # 현재 환경 확인
```

---

## 5. CI/CD 환경 설정

### GitHub Actions

```yaml
# .github/workflows/deploy.yml
name: Deploy to NHN Cloud

on:
  push:
    branches: [main, develop]

jobs:
  deploy:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4

      - name: Set up Go
        uses: actions/setup-go@v5
        with:
          go-version: '1.22'

      - name: Install NHN CLI
        run: |
          git clone https://github.com/your-repo/nhncli.git
          cd nhncli
          go build -o nhn main.go
          sudo mv nhn /usr/local/bin/

      - name: Configure NHN CLI
        run: |
          mkdir -p ~/.nhn
          cat > ~/.nhn/config.json << EOF
          {
            "profiles": {
              "default": {
                "auth_type": "oauth",
                "user_access_key_id": "${{ secrets.NHN_ACCESS_KEY_ID }}",
                "secret_access_key": "${{ secrets.NHN_SECRET_ACCESS_KEY }}",
                "region": "KR1"
              }
            }
          }
          EOF

      - name: Deploy to Dev
        if: github.ref == 'refs/heads/develop'
        run: |
          nhn compute instance list
          # 배포 스크립트 실행

      - name: Deploy to Prod
        if: github.ref == 'refs/heads/main'
        env:
          NHN_REGION: KR2
        run: |
          nhn --region KR2 compute instance list
          # 배포 스크립트 실행
```

### GitLab CI

```yaml
# .gitlab-ci.yml
stages:
  - deploy

variables:
  NHN_ACCESS_KEY_ID: $NHN_ACCESS_KEY_ID
  NHN_SECRET_ACCESS_KEY: $NHN_SECRET_ACCESS_KEY

.nhn-setup: &nhn-setup
  before_script:
    - git clone https://github.com/your-repo/nhncli.git
    - cd nhncli && go build -o /usr/local/bin/nhn main.go && cd ..
    - mkdir -p ~/.nhn
    - |
      cat > ~/.nhn/config.json << EOF
      {
        "profiles": {
          "default": {
            "auth_type": "oauth",
            "user_access_key_id": "$NHN_ACCESS_KEY_ID",
            "secret_access_key": "$NHN_SECRET_ACCESS_KEY",
            "region": "$NHN_REGION"
          }
        }
      }
      EOF

deploy-dev:
  stage: deploy
  <<: *nhn-setup
  variables:
    NHN_REGION: KR1
  script:
    - nhn compute instance list
  only:
    - develop

deploy-prod:
  stage: deploy
  <<: *nhn-setup
  variables:
    NHN_REGION: KR2
  script:
    - nhn compute instance list
  only:
    - main
  when: manual
```

---

## 6. 환경별 리소스 네이밍

### 네이밍 컨벤션

| 리소스 | 개발 | 스테이징 | 운영 |
|--------|------|----------|------|
| VPC | `dev-vpc` | `staging-vpc` | `prod-vpc` |
| 서브넷 | `dev-public-subnet` | `staging-public-subnet` | `prod-public-subnet` |
| 보안 그룹 | `dev-web-sg` | `staging-web-sg` | `prod-web-sg` |
| 인스턴스 | `dev-web-01` | `staging-web-01` | `prod-web-01` |

### 환경별 리소스 검색

```bash
# 개발 환경 리소스만 필터링
nhn --profile dev --output json compute instance list | \
  jq '.[] | select(.name | startswith("dev-"))'

# 운영 환경 웹 서버만 조회
nhn --profile prod --output json compute instance list | \
  jq '.[] | select(.name | contains("web"))'
```

---

## 7. 환경 간 리소스 동기화 확인

```bash
#!/bin/bash
# compare-environments.sh

echo "환경 간 리소스 비교"
echo "===================="

# VPC 개수 비교
echo ""
echo "📦 VPC 개수:"
for profile in dev staging prod; do
  count=$(nhn --profile $profile --output json vpc list | jq 'length')
  echo "  $profile: $count"
done

# 인스턴스 타입 분포
echo ""
echo "🖥️ 인스턴스 타입 분포:"
for profile in dev staging prod; do
  echo "  [$profile]"
  nhn --profile $profile --output json compute instance list | \
    jq -r '.[].flavor_name' | sort | uniq -c | \
    while read count flavor; do
      echo "    $flavor: $count"
    done
done

# 보안 그룹 규칙 수
echo ""
echo "🔒 보안 그룹별 규칙 수:"
for profile in dev staging prod; do
  echo "  [$profile]"
  nhn --profile $profile --output json vpc sg list | \
    jq -r '.[] | "    \(.name): \(.rules_count // "N/A") rules"'
done
```

---

## 참고

- [설정 가이드](../Configuration.md)
- [자동화 스크립트](Automation-Scripts.md)
- [전역 옵션](../Commands/Global-Options.md)
