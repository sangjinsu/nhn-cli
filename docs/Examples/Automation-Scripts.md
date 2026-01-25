# 자동화 스크립트 예제

JSON 출력과 쉘 스크립트를 활용한 NHN Cloud CLI 자동화 방법입니다.

---

## JSON 출력 기본

### jq 설치

```bash
# macOS
brew install jq

# Ubuntu/Debian
sudo apt-get install jq

# CentOS/RHEL
sudo yum install jq
```

### 기본 사용법

```bash
# JSON 형식으로 출력
nhn --output json vpc list

# jq로 필터링
nhn --output json vpc list | jq '.[].name'
```

---

## 데이터 추출

### 인스턴스 정보 추출

```bash
# 인스턴스 이름과 ID
nhn --output json compute instance list | \
  jq '.[] | {name: .name, id: .id}'

# 인스턴스 이름과 IP 주소
nhn --output json compute instance list | \
  jq '.[] | {name: .name, addresses: .addresses}'

# 특정 필드만 테이블 형식으로
nhn --output json compute instance list | \
  jq -r '.[] | [.name, .status, .flavor_name] | @tsv'
```

### 필터링

```bash
# 실행 중인 인스턴스만
nhn --output json compute instance list | \
  jq '.[] | select(.status == "ACTIVE")'

# 특정 Flavor 인스턴스
nhn --output json compute instance list | \
  jq '.[] | select(.flavor_name == "m2.c2m4")'

# 이름에 "web" 포함된 인스턴스
nhn --output json compute instance list | \
  jq '.[] | select(.name | contains("web"))'
```

### 집계

```bash
# 상태별 인스턴스 수
nhn --output json compute instance list | \
  jq 'group_by(.status) | .[] | {status: .[0].status, count: length}'

# Flavor별 인스턴스 수
nhn --output json compute instance list | \
  jq 'group_by(.flavor_name) | .[] | {flavor: .[0].flavor_name, count: length}'

# 총 인스턴스 수
nhn --output json compute instance list | jq 'length'
```

---

## 일괄 작업 스크립트

### 모든 인스턴스 중지

```bash
#!/bin/bash
# stop-all-instances.sh

echo "모든 인스턴스를 중지합니다..."

INSTANCES=$(nhn --output json compute instance list | \
  jq -r '.[] | select(.status == "ACTIVE") | .id')

for id in $INSTANCES; do
  name=$(nhn --output json compute instance describe $id | jq -r '.name')
  echo "중지: $name ($id)"
  nhn compute instance stop $id
done

echo "완료!"
```

### 특정 조건의 인스턴스만 시작

```bash
#!/bin/bash
# start-web-servers.sh

echo "웹 서버 인스턴스를 시작합니다..."

INSTANCES=$(nhn --output json compute instance list | \
  jq -r '.[] | select(.status == "SHUTOFF") | select(.name | contains("web")) | .id')

for id in $INSTANCES; do
  name=$(nhn --output json compute instance describe $id | jq -r '.name')
  echo "시작: $name ($id)"
  nhn compute instance start $id
done

echo "완료!"
```

### 오래된 플로팅 IP 정리

```bash
#!/bin/bash
# cleanup-unused-fips.sh

echo "사용하지 않는 플로팅 IP를 정리합니다..."

UNUSED_FIPS=$(nhn --output json vpc fip list | \
  jq -r '.[] | select(.status == "DOWN") | .id')

for id in $UNUSED_FIPS; do
  ip=$(nhn --output json vpc fip list | jq -r ".[] | select(.id == \"$id\") | .floating_ip_address")
  echo "삭제: $ip ($id)"
  nhn vpc fip delete $id
done

echo "완료!"
```

---

## 모니터링 스크립트

### 인스턴스 상태 대시보드

```bash
#!/bin/bash
# instance-dashboard.sh

clear
echo "=========================================="
echo "    NHN Cloud 인스턴스 대시보드"
echo "    $(date '+%Y-%m-%d %H:%M:%S')"
echo "=========================================="
echo ""

# 요약 통계
TOTAL=$(nhn --output json compute instance list | jq 'length')
ACTIVE=$(nhn --output json compute instance list | jq '[.[] | select(.status == "ACTIVE")] | length')
SHUTOFF=$(nhn --output json compute instance list | jq '[.[] | select(.status == "SHUTOFF")] | length')
ERROR=$(nhn --output json compute instance list | jq '[.[] | select(.status == "ERROR")] | length')

echo "📊 요약"
echo "   전체: $TOTAL | 실행: $ACTIVE | 중지: $SHUTOFF | 오류: $ERROR"
echo ""

# 상세 목록
echo "📋 인스턴스 목록"
printf "%-30s %-10s %-12s %-16s\n" "NAME" "STATUS" "FLAVOR" "IP"
echo "-------------------------------------------------------------------"

nhn --output json compute instance list | \
  jq -r '.[] | [.name, .status, .flavor_name, (.addresses | to_entries | .[0].value | .[0].addr // "N/A")] | @tsv' | \
  while IFS=$'\t' read name status flavor ip; do
    # 상태에 따른 이모지
    case $status in
      ACTIVE)  emoji="🟢" ;;
      SHUTOFF) emoji="🔴" ;;
      BUILD)   emoji="🟡" ;;
      ERROR)   emoji="⚠️ " ;;
      *)       emoji="⚪" ;;
    esac
    printf "%-30s %s %-10s %-12s %-16s\n" "$name" "$emoji" "$status" "$flavor" "$ip"
  done
```

### 리소스 사용량 모니터

```bash
#!/bin/bash
# resource-monitor.sh

echo "NHN Cloud 리소스 사용량"
echo "======================="
echo ""

# VPC
VPC_COUNT=$(nhn --output json vpc list | jq 'length')
echo "📦 VPC: $VPC_COUNT"

# 서브넷
SUBNET_COUNT=$(nhn --output json vpc subnet list | jq 'length')
echo "📍 서브넷: $SUBNET_COUNT"

# 보안 그룹
SG_COUNT=$(nhn --output json vpc sg list | jq 'length')
echo "🔒 보안 그룹: $SG_COUNT"

# 인스턴스
INSTANCE_COUNT=$(nhn --output json compute instance list | jq 'length')
echo "🖥️  인스턴스: $INSTANCE_COUNT"

# 플로팅 IP
FIP_TOTAL=$(nhn --output json vpc fip list | jq 'length')
FIP_USED=$(nhn --output json vpc fip list | jq '[.[] | select(.status == "ACTIVE")] | length')
echo "🌐 플로팅 IP: $FIP_USED / $FIP_TOTAL (사용/전체)"

# 키페어
KEYPAIR_COUNT=$(nhn --output json compute keypair list | jq 'length')
echo "🔑 키페어: $KEYPAIR_COUNT"
```

---

## 백업 및 문서화

### 리소스 목록 백업

```bash
#!/bin/bash
# backup-resource-list.sh

BACKUP_DIR="./nhn-backup-$(date +%Y%m%d)"
mkdir -p $BACKUP_DIR

echo "리소스 목록을 백업합니다: $BACKUP_DIR"

# VPC
nhn --output json vpc list > $BACKUP_DIR/vpcs.json
echo "✅ VPC 목록 저장"

# 서브넷
nhn --output json vpc subnet list > $BACKUP_DIR/subnets.json
echo "✅ 서브넷 목록 저장"

# 보안 그룹
nhn --output json vpc sg list > $BACKUP_DIR/security-groups.json
echo "✅ 보안 그룹 목록 저장"

# 인스턴스
nhn --output json compute instance list > $BACKUP_DIR/instances.json
echo "✅ 인스턴스 목록 저장"

# 키페어
nhn --output json compute keypair list > $BACKUP_DIR/keypairs.json
echo "✅ 키페어 목록 저장"

# 플로팅 IP
nhn --output json vpc fip list > $BACKUP_DIR/floating-ips.json
echo "✅ 플로팅 IP 목록 저장"

echo ""
echo "백업 완료: $BACKUP_DIR"
ls -la $BACKUP_DIR
```

### 보안 그룹 규칙 문서화

```bash
#!/bin/bash
# document-security-rules.sh

OUTPUT_FILE="security-rules-$(date +%Y%m%d).md"

echo "# 보안 그룹 규칙 문서" > $OUTPUT_FILE
echo "" >> $OUTPUT_FILE
echo "생성일: $(date '+%Y-%m-%d %H:%M:%S')" >> $OUTPUT_FILE
echo "" >> $OUTPUT_FILE

SG_IDS=$(nhn --output json vpc sg list | jq -r '.[].id')

for sg_id in $SG_IDS; do
  SG_INFO=$(nhn --output json vpc sg describe $sg_id)
  SG_NAME=$(echo $SG_INFO | jq -r '.name')
  SG_DESC=$(echo $SG_INFO | jq -r '.description // "N/A"')

  echo "## $SG_NAME" >> $OUTPUT_FILE
  echo "" >> $OUTPUT_FILE
  echo "- ID: $sg_id" >> $OUTPUT_FILE
  echo "- 설명: $SG_DESC" >> $OUTPUT_FILE
  echo "" >> $OUTPUT_FILE
  echo "### 인바운드 규칙" >> $OUTPUT_FILE
  echo "" >> $OUTPUT_FILE
  echo "| 프로토콜 | 포트 | 원격 IP |" >> $OUTPUT_FILE
  echo "|----------|------|---------|" >> $OUTPUT_FILE

  echo $SG_INFO | jq -r '.security_group_rules[] | select(.direction == "ingress") | "| \(.protocol // "any") | \(.port_range_min // "all")-\(.port_range_max // "all") | \(.remote_ip_prefix // "N/A") |"' >> $OUTPUT_FILE

  echo "" >> $OUTPUT_FILE
done

echo "문서 생성 완료: $OUTPUT_FILE"
```

---

## 알림 스크립트

### Slack 알림

```bash
#!/bin/bash
# notify-slack.sh

SLACK_WEBHOOK_URL="https://hooks.slack.com/services/xxx/yyy/zzz"

# 오류 상태 인스턴스 확인
ERROR_INSTANCES=$(nhn --output json compute instance list | \
  jq -r '.[] | select(.status == "ERROR") | .name')

if [ -n "$ERROR_INSTANCES" ]; then
  MESSAGE="🚨 *NHN Cloud 알림*\n오류 상태 인스턴스 발견:\n\`\`\`$ERROR_INSTANCES\`\`\`"

  curl -X POST -H 'Content-type: application/json' \
    --data "{\"text\": \"$MESSAGE\"}" \
    $SLACK_WEBHOOK_URL
fi
```

### 이메일 알림

```bash
#!/bin/bash
# notify-email.sh

EMAIL="admin@example.com"
SUBJECT="NHN Cloud 일일 리포트"

# 리포트 생성
REPORT=$(cat << EOF
NHN Cloud 일일 리포트
=====================
날짜: $(date '+%Y-%m-%d')

리소스 현황:
- VPC: $(nhn --output json vpc list | jq 'length')개
- 인스턴스: $(nhn --output json compute instance list | jq 'length')개
  - 실행 중: $(nhn --output json compute instance list | jq '[.[] | select(.status == "ACTIVE")] | length')개
  - 중지: $(nhn --output json compute instance list | jq '[.[] | select(.status == "SHUTOFF")] | length')개
- 플로팅 IP: $(nhn --output json vpc fip list | jq 'length')개

EOF
)

echo "$REPORT" | mail -s "$SUBJECT" $EMAIL
```

---

## Cron 작업 설정

### 정기 작업 등록

```bash
# crontab -e

# 매일 오전 9시에 상태 체크
0 9 * * * /path/to/instance-dashboard.sh > /var/log/nhn-dashboard.log

# 매주 일요일 자정에 리소스 백업
0 0 * * 0 /path/to/backup-resource-list.sh

# 매 시간 오류 체크 및 알림
0 * * * * /path/to/notify-slack.sh
```

---

## 참고

- [VPC 명령어](../Commands/VPC.md)
- [Compute 명령어](../Commands/Compute.md)
- [다중 환경 관리](Multi-Environment.md)
