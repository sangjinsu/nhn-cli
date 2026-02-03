# Pipeline 명령어

Pipeline 서비스를 관리하는 명령어입니다.

> **참고**: Pipeline은 AppKey 기반 인증을 사용합니다. 사전에 `nhn configure service pipeline`으로 AppKey를 설정하거나, `--app-key` 플래그를 사용하세요.

---

## 파이프라인 실행

### 파이프라인 수동 실행

```bash
nhn pipeline execute <pipeline-name>
```

| 인자 | 설명 | 필수 |
|------|------|------|
| `pipeline-name` | 실행할 파이프라인 이름 | O |

**예시:**
```bash
# 파이프라인 실행
nhn pipeline execute my-build-pipeline
```

출력 예시:
```
파이프라인 'my-build-pipeline' 실행이 요청되었습니다.
```

---

## AppKey 설정

### 프로필에 AppKey 저장

```bash
nhn configure service pipeline
```

대화형 프롬프트:
```
프로필 이름 [default]:
Pipeline AppKey: your-pipeline-appkey
```

### 명령줄에서 AppKey 지정

```bash
# --app-key 플래그로 직접 지정
nhn pipeline execute my-pipeline --app-key your-pipeline-appkey
```

---

## 실전 예제

### CI/CD 자동화 스크립트

```bash
#!/bin/bash

# 빌드 파이프라인 실행
nhn pipeline execute build-pipeline

# 테스트 파이프라인 실행
nhn pipeline execute test-pipeline

# 배포 파이프라인 실행 (프로덕션)
nhn pipeline execute deploy-production
```

### 다중 프로필 사용

```bash
# 개발 환경 파이프라인 실행
nhn --profile dev pipeline execute dev-pipeline

# 운영 환경 파이프라인 실행
nhn --profile prod pipeline execute prod-pipeline
```

---

## 참고

- [설정 가이드](../Configuration.md)
- [전역 옵션](Global-Options.md)
- [Deploy 명령어](Deploy.md)
