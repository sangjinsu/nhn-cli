# Deploy 명령어

Deploy 서비스를 관리하는 명령어입니다.

> **참고**: Deploy는 AppKey 기반 인증을 사용합니다. 사전에 `nhn configure service deploy`로 AppKey를 설정하거나, `--app-key` 플래그를 사용하세요.

---

## 배포 실행

### 배포 실행

```bash
nhn deploy execute --artifact-id <id> --server-group-id <id> [options]
```

| 옵션 | 설명 | 필수 |
|------|------|------|
| `--artifact-id` | 아티팩트 ID | O |
| `--server-group-id` | 서버 그룹 ID | O |
| `--concurrent-num` | 동시 실행 수 | X |
| `--next-when-fail` | 실패 시 다음 서버 진행 | X |
| `--deploy-note` | 배포 메모 | X |
| `--async` | 비동기 실행 | X |

**예시:**
```bash
# 기본 배포 실행
nhn deploy execute --artifact-id 123 --server-group-id 456

# 동시 실행 및 메모 추가
nhn deploy execute --artifact-id 123 --server-group-id 456 \
  --concurrent-num 2 \
  --deploy-note "v1.2.0 릴리스"

# 실패해도 다음 서버 진행
nhn deploy execute --artifact-id 123 --server-group-id 456 --next-when-fail
```

출력 예시:
```
배포가 실행되었습니다. (ID: 789, 상태: RUNNING)
```

---

## 바이너리 업로드

### 바이너리 업로드

```bash
nhn deploy upload --artifact-id <id> --binary-group-key <key> --type <type> --file <path> [options]
```

| 옵션 | 설명 | 필수 |
|------|------|------|
| `--artifact-id` | 아티팩트 ID | O |
| `--binary-group-key` | 바이너리 그룹 키 | O |
| `--type` | 애플리케이션 타입 (`client` 또는 `server`) | O |
| `--file` | 업로드할 파일 경로 | O |
| `--version` | 바이너리 버전 (최대 100자) | X |
| `--description` | 설명 | X |
| `--os-type` | OS 타입 (iOS, Android 등, client 타입일 때) | X |
| `--meta-file` | iOS plist 파일 경로 | X |
| `--fix` | client 바이너리 fix 플래그 | X |

**예시:**
```bash
# 서버 바이너리 업로드
nhn deploy upload \
  --artifact-id 123 \
  --binary-group-key 456 \
  --type server \
  --file ./build/app.jar \
  --version "1.2.0" \
  --description "새 기능 추가"

# 클라이언트 바이너리 업로드 (Android)
nhn deploy upload \
  --artifact-id 123 \
  --binary-group-key 456 \
  --type client \
  --file ./build/app.apk \
  --os-type Android \
  --version "1.2.0"

# iOS 앱 업로드 (plist 포함)
nhn deploy upload \
  --artifact-id 123 \
  --binary-group-key 456 \
  --type client \
  --file ./build/app.ipa \
  --os-type iOS \
  --meta-file ./manifest.plist \
  --version "1.2.0"
```

출력 예시:
```
바이너리가 업로드되었습니다.
  Binary Key:   abc123def456
  Download URL: https://...
```

---

## AppKey 설정

### 프로필에 AppKey 저장

```bash
nhn configure service deploy
```

대화형 프롬프트:
```
프로필 이름 [default]:
Deploy AppKey: your-deploy-appkey
```

### 명령줄에서 AppKey 지정

```bash
# --app-key 플래그로 직접 지정
nhn deploy execute --artifact-id 123 --server-group-id 456 --app-key your-deploy-appkey
```

---

## 실전 예제

### CI/CD 파이프라인 통합

```bash
#!/bin/bash

# 1. 빌드
./gradlew build

# 2. 바이너리 업로드
nhn deploy upload \
  --artifact-id $ARTIFACT_ID \
  --binary-group-key $BINARY_GROUP_KEY \
  --type server \
  --file ./build/libs/app.jar \
  --version "$(git describe --tags)"

# 3. 배포 실행
nhn deploy execute \
  --artifact-id $ARTIFACT_ID \
  --server-group-id $SERVER_GROUP_ID \
  --deploy-note "$(git log -1 --pretty=%B)"
```

### 롤링 배포

```bash
# 동시 실행 수를 1로 설정하여 롤링 배포
nhn deploy execute \
  --artifact-id 123 \
  --server-group-id 456 \
  --concurrent-num 1 \
  --next-when-fail
```

---

## 참고

- [설정 가이드](../Configuration.md)
- [전역 옵션](Global-Options.md)
- [Pipeline 명령어](Pipeline.md)
