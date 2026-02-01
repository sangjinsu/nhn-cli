# Block Storage 명령어

NHN Cloud Block Storage 리소스 관리 명령어입니다.

---

## 개요

```bash
nhn blockstorage <subcommand> [options]
```

### 서브 명령어

| 명령어 | 설명 |
|--------|------|
| `volume` | 볼륨 관리 |
| `snapshot` | 스냅샷 관리 |
| `type` | 볼륨 타입 조회 |

---

## 볼륨 관리

### 볼륨 목록 조회

```bash
nhn blockstorage volume list
```

**출력 예시:**
```
ID                                      NAME            SIZE(GB)  STATUS      TYPE    AZ          CREATED
vol-11111111-...                        my-volume       20        available   SSD     kr-pub-a    2024-01-15T10:30:00Z
vol-22222222-...                        db-volume       100       in-use      HDD     kr-pub-b    2024-01-16T08:00:00Z
```

### 볼륨 상세 조회

```bash
nhn blockstorage volume describe <volume-id>
```

**옵션:**

| 옵션 | 설명 |
|------|------|
| `<volume-id>` | 볼륨 ID (필수) |

### 볼륨 생성

```bash
nhn blockstorage volume create [options]
```

**옵션:**

| 옵션 | 설명 | 필수 |
|------|------|------|
| `--size` | 볼륨 크기 (GB) | O |
| `--name` | 볼륨 이름 | - |
| `--type` | 볼륨 타입 (SSD/HDD) | - |
| `--availability-zone` | 가용성 영역 | - |
| `--description` | 볼륨 설명 | - |
| `--snapshot-id` | 스냅샷 ID (스냅샷에서 생성) | - |

**예시:**
```bash
# 기본 볼륨 생성
nhn blockstorage volume create --size 20 --name my-volume

# SSD 타입 볼륨 생성
nhn blockstorage volume create \
  --size 100 \
  --name db-volume \
  --type SSD \
  --availability-zone kr-pub-a \
  --description "Database volume"

# 스냅샷에서 볼륨 생성
nhn blockstorage volume create \
  --size 20 \
  --name restored-volume \
  --snapshot-id snap-12345678
```

### 볼륨 삭제

```bash
nhn blockstorage volume delete <volume-id>
```

---

## 스냅샷 관리

### 스냅샷 목록 조회

```bash
nhn blockstorage snapshot list
```

**출력 예시:**
```
ID                                      NAME            VOLUME ID                               SIZE(GB)  STATUS      CREATED
snap-11111111-...                       my-snapshot     vol-11111111-...                        20        available   2024-01-15T12:00:00Z
snap-22222222-...                       db-backup       vol-22222222-...                        100       available   2024-01-16T00:00:00Z
```

### 스냅샷 상세 조회

```bash
nhn blockstorage snapshot describe <snapshot-id>
```

**옵션:**

| 옵션 | 설명 |
|------|------|
| `<snapshot-id>` | 스냅샷 ID (필수) |

### 스냅샷 생성

```bash
nhn blockstorage snapshot create [options]
```

**옵션:**

| 옵션 | 설명 | 필수 |
|------|------|------|
| `--volume-id` | 볼륨 ID | O |
| `--name` | 스냅샷 이름 | - |
| `--description` | 스냅샷 설명 | - |
| `--force` | 사용 중인 볼륨 강제 스냅샷 | - |

**예시:**
```bash
# 기본 스냅샷 생성
nhn blockstorage snapshot create --volume-id vol-11111111

# 이름과 설명 지정
nhn blockstorage snapshot create \
  --volume-id vol-11111111 \
  --name my-snapshot \
  --description "Daily backup"

# 사용 중인 볼륨 강제 스냅샷
nhn blockstorage snapshot create \
  --volume-id vol-22222222 \
  --name forced-snapshot \
  --force
```

### 스냅샷 삭제

```bash
nhn blockstorage snapshot delete <snapshot-id>
```

---

## 볼륨 타입 조회

### 볼륨 타입 목록

```bash
nhn blockstorage type list
```

**출력 예시:**
```
ID                                      NAME
type-11111111-...                       SSD
type-22222222-...                       HDD
```

---

## 참고

- [전역 옵션](Global-Options.md)
- [Compute 명령어](Compute.md)
- [Load Balancer 명령어](LoadBalancer.md)
