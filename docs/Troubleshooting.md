# 문제 해결

NHN Cloud CLI 사용 시 발생할 수 있는 문제와 해결 방법입니다.

---

## 인증 오류

### "토큰이 만료되었습니다" 오류

**증상:**
```
Error: 토큰이 만료되었습니다. 다시 인증하세요.
```

**해결:**
```bash
# 토큰 캐시 삭제
rm ~/.nhn/credentials.json

# 다시 명령 실행 (자동으로 새 토큰 발급)
nhn vpc list
```

### "인증 실패" 오류

**증상:**
```
Error: 인증에 실패했습니다. 인증 정보를 확인하세요.
```

**원인:**
- User Access Key ID 또는 Secret Access Key가 잘못됨
- Identity 인증의 경우 Tenant ID, Username, Password가 잘못됨

**해결:**
```bash
# 인증 정보 재설정
nhn configure

# 또는 특정 프로필 재설정
nhn configure --profile <profile-name>
```

### "Tenant ID가 필요합니다" 오류

**증상:**
```
Error: Tenant ID가 필요합니다.
```

**원인:**
- Compute API는 Tenant ID가 필요하지만, OAuth 인증만으로는 Tenant ID를 얻을 수 없는 경우

**해결:**

1. **Identity 인증 사용 (권장):**
   ```bash
   nhn configure
   # 인증 방식 선택: 2 (Identity)
   ```

2. **설정 파일에 tenant_id 직접 추가:**
   ```bash
   # ~/.nhn/config.json 편집
   {
     "profiles": {
       "default": {
         "auth_type": "oauth",
         "user_access_key_id": "...",
         "secret_access_key": "...",
         "region": "KR1",
         "tenant_id": "your-tenant-id"  // 추가
       }
     }
   }
   ```

---

## 네트워크 오류

### "연결할 수 없습니다" 오류

**증상:**
```
Error: API 서버에 연결할 수 없습니다.
```

**원인:**
- 네트워크 연결 문제
- 방화벽에서 HTTPS 차단
- API 서버 점검

**해결:**
```bash
# 네트워크 연결 확인
ping oauth.api.nhncloudservice.com

# HTTPS 연결 확인
curl -v https://oauth.api.nhncloudservice.com

# 디버그 모드로 상세 정보 확인
nhn --debug vpc list
```

### "타임아웃" 오류

**증상:**
```
Error: 요청 시간이 초과되었습니다.
```

**해결:**
- 네트워크 상태 확인
- 잠시 후 다시 시도
- API 서버 상태 확인

---

## 권한 오류

### "권한이 없습니다" 오류

**증상:**
```
Error: 이 작업을 수행할 권한이 없습니다.
```

**원인:**
- 프로젝트 멤버 권한 부족
- 다른 프로젝트의 리소스 접근 시도

**해결:**
1. NHN Cloud 콘솔에서 프로젝트 멤버 권한 확인
2. 올바른 프로필(프로젝트) 사용 중인지 확인:
   ```bash
   nhn configure list
   nhn --profile <correct-profile> vpc list
   ```

### "리소스를 찾을 수 없습니다" 오류

**증상:**
```
Error: 리소스를 찾을 수 없습니다: vpc-12345678
```

**원인:**
- 잘못된 리소스 ID
- 다른 리전의 리소스
- 삭제된 리소스

**해결:**
```bash
# 올바른 리전 확인
nhn --region KR1 vpc list
nhn --region KR2 vpc list

# 리소스 ID 재확인
nhn vpc list
```

---

## 프로필 오류

### "프로필을 찾을 수 없습니다" 오류

**증상:**
```
Error: 프로필 'production'을 찾을 수 없습니다.
```

**해결:**
```bash
# 프로필 목록 확인
nhn configure list

# 프로필 생성
nhn configure --profile production
```

### 설정 파일 손상

**증상:**
```
Error: 설정 파일을 파싱할 수 없습니다.
```

**해결:**
```bash
# 설정 파일 백업
mv ~/.nhn/config.json ~/.nhn/config.json.backup

# 새로 설정
nhn configure
```

---

## 리소스 오류

### "리소스가 사용 중입니다" 오류

**증상:**
```
Error: VPC를 삭제할 수 없습니다. 종속 리소스가 있습니다.
```

**해결:**
종속 리소스를 먼저 삭제:
```bash
# 1. 인스턴스 삭제
nhn compute instance list --vpc-id <vpc-id>
nhn compute instance delete <instance-id>

# 2. 플로팅 IP 연결 해제 및 삭제
nhn vpc fip list
nhn vpc fip disassociate <fip-id>
nhn vpc fip delete <fip-id>

# 3. 보안 그룹 삭제 (default 제외)
nhn vpc sg list
nhn vpc sg delete <sg-id>

# 4. 서브넷 삭제
nhn vpc subnet list --vpc-id <vpc-id>
nhn vpc subnet delete <subnet-id>

# 5. VPC 삭제
nhn vpc delete <vpc-id>
```

### "CIDR이 겹칩니다" 오류

**증상:**
```
Error: 지정한 CIDR이 기존 서브넷과 겹칩니다.
```

**해결:**
```bash
# 기존 서브넷 CIDR 확인
nhn vpc subnet list --vpc-id <vpc-id>

# 겹치지 않는 CIDR로 재시도
nhn vpc subnet create --vpc-id <vpc-id> --name new-subnet --cidr 192.168.2.0/24
```

---

## 출력 오류

### JSON 출력이 비어있음

**증상:**
```bash
nhn --output json vpc list
[]
```

**원인:**
- 해당 리전에 리소스가 없음
- 프로필/리전 설정 확인 필요

**해결:**
```bash
# 프로필 확인
nhn configure list

# 다른 리전 시도
nhn --region KR1 vpc list
nhn --region KR2 vpc list
```

### 출력 형식 깨짐

**증상:**
- 테이블 정렬이 맞지 않음
- 특수 문자가 깨짐

**해결:**
- 터미널 인코딩을 UTF-8로 설정
- 터미널 폭을 늘림
- JSON 출력 사용:
  ```bash
  nhn --output json vpc list | jq
  ```

---

## 디버깅

### 디버그 모드

```bash
nhn --debug vpc list
```

출력:
```
DEBUG: GET https://kr1-api-network-infrastructure.nhncloudservice.com/v2.0/vpcs
DEBUG: Request Headers:
  Authorization: Bearer eyJhbG...
  Content-Type: application/json
DEBUG: Response Status: 200 OK
DEBUG: Response Body: {"vpcs": [...]}
```

### 로그 확인

```bash
# 디버그 출력을 파일로 저장
nhn --debug vpc list 2>&1 | tee debug.log
```

---

## 자주 묻는 질문 (FAQ)

### Q: OAuth와 Identity 중 어떤 인증을 사용해야 하나요?

**A:** OAuth 인증을 권장합니다.
- OAuth: 보안성이 높고, API 키 회전이 쉬움
- Identity: 설정이 간단하지만, 비밀번호 관리 필요

### Q: 여러 프로젝트를 관리하려면?

**A:** 프로필을 사용하세요.
```bash
nhn configure --profile project-a
nhn configure --profile project-b

nhn --profile project-a vpc list
nhn --profile project-b vpc list
```

### Q: 토큰 유효 기간은 얼마인가요?

**A:** 12시간입니다. CLI가 자동으로 갱신합니다.

### Q: 디버그 정보는 어디에 저장되나요?

**A:** 디버그 정보는 stdout으로 출력되며, 별도 파일에 저장되지 않습니다.

### Q: API 호출 한도가 있나요?

**A:** NHN Cloud API 정책에 따릅니다. 과도한 호출 시 일시적으로 차단될 수 있습니다.

---

## 도움 요청

문제가 해결되지 않으면:

1. **GitHub Issues**: 버그 리포트 및 기능 요청
2. **NHN Cloud 고객센터**: 계정 및 서비스 관련 문의
3. **NHN Cloud 문서**: https://docs.nhncloud.com/

---

## 참고

- [설정 가이드](Configuration.md)
- [API 레퍼런스](API-Reference.md)
