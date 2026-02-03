# Object Storage 명령어

Object Storage 서비스를 관리하는 명령어입니다.

> **별칭**: `objectstorage` 대신 `os`를 사용할 수 있습니다.
> 예: `nhn os container list`

---

## 컨테이너 관리

### 컨테이너 목록 조회

```bash
nhn objectstorage container list
nhn os container list  # 별칭
```

출력 예시:
```
NAME            OBJECTS     BYTES
my-container    15          1048576
backup          100         52428800
```

### 컨테이너 메타데이터 조회

```bash
nhn os container describe <container-name>
```

출력 예시:
```
Container:    my-container
Object Count: 15
Bytes Used:   1048576
Read ACL:     .r:*
Write ACL:
```

### 컨테이너 생성

```bash
nhn os container create <container-name>
```

**예시:**
```bash
nhn os container create my-new-container
```

### 컨테이너 삭제

```bash
nhn os container delete <container-name>
```

> **주의**: 컨테이너 내에 오브젝트가 있으면 삭제할 수 없습니다. 먼저 모든 오브젝트를 삭제하세요.

---

## 오브젝트 관리

### 오브젝트 목록 조회

```bash
nhn os object list --container <container-name>
```

| 옵션 | 설명 | 필수 |
|------|------|------|
| `--container` | 컨테이너 이름 | O |

출력 예시:
```
NAME            BYTES       CONTENT-TYPE        LAST-MODIFIED
test.txt        1024        text/plain          2024-01-15T10:30:00Z
image.png       51200       image/png           2024-01-15T11:00:00Z
```

### 오브젝트 업로드

```bash
nhn os object upload --container <container-name> --file <file-path> [--name <object-name>]
```

| 옵션 | 설명 | 필수 |
|------|------|------|
| `--container` | 컨테이너 이름 | O |
| `--file` | 업로드할 파일 경로 | O |
| `--name` | 오브젝트 이름 (기본: 파일명) | X |

**예시:**
```bash
# 파일 업로드 (파일명 그대로)
nhn os object upload --container my-container --file ./test.txt

# 커스텀 이름으로 업로드
nhn os object upload --container my-container --file ./test.txt --name documents/test.txt
```

### 오브젝트 다운로드

```bash
nhn os object download <object-name> --container <container-name> [--output-file <path>]
```

| 옵션 | 설명 | 필수 |
|------|------|------|
| `--container` | 컨테이너 이름 | O |
| `--output-file` | 저장 경로 (기본: 현재 디렉토리에 오브젝트 이름) | X |

**예시:**
```bash
# 기본 다운로드
nhn os object download test.txt --container my-container

# 저장 경로 지정
nhn os object download test.txt --container my-container --output-file ./downloads/test.txt
```

### 오브젝트 메타데이터 조회

```bash
nhn os object describe <object-name> --container <container-name>
```

출력 예시:
```
Object:         test.txt
Container:      my-container
Content-Length: 1024
Content-Type:   text/plain
ETag:           d41d8cd98f00b204e9800998ecf8427e
Last-Modified:  2024-01-15T10:30:00Z
```

### 오브젝트 삭제

```bash
nhn os object delete <object-name> --container <container-name>
```

**예시:**
```bash
nhn os object delete test.txt --container my-container
```

---

## 실전 예제

### 백업 워크플로우

```bash
# 1. 백업용 컨테이너 생성
nhn os container create backups

# 2. 파일 업로드
nhn os object upload --container backups --file ./database-dump.sql

# 3. 업로드 확인
nhn os object list --container backups

# 4. 필요시 다운로드
nhn os object download database-dump.sql --container backups
```

### 디렉토리 구조 업로드

```bash
# 폴더 구조를 유지하면서 업로드
for file in ./data/*; do
  name=$(basename "$file")
  nhn os object upload --container my-container --file "$file" --name "data/$name"
done
```

### JSON 출력 활용

```bash
# 컨테이너별 사용량 확인
nhn --output json os container list | jq '.[] | {name: .name, bytes: .bytes}'

# 특정 확장자 오브젝트만 필터링
nhn --output json os object list --container my-container | \
  jq '.[] | select(.name | endswith(".txt"))'
```

---

## 참고

- [설정 가이드](../Configuration.md)
- [전역 옵션](Global-Options.md)
