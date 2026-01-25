# 설치 가이드

NHN Cloud CLI 설치 방법을 상세히 설명합니다.

---

## 요구 사항

| 항목 | 최소 버전 |
|------|----------|
| Go | 1.22 이상 |
| Git | 2.0 이상 |

---

## 소스에서 빌드

### 1. 저장소 클론

```bash
git clone https://github.com/sangjinsu/nhn-cli.git
cd nhn-cli
```

### 2. 의존성 설치

```bash
go mod download
```

### 3. 빌드

```bash
# 기본 빌드
go build -o nhn main.go

# 버전 정보 포함 빌드
go build -ldflags "-X main.version=1.0.0 -X main.buildDate=$(date -u +%Y-%m-%dT%H:%M:%SZ)" -o nhn main.go
```

---

## OS별 설치

### Linux

```bash
# 빌드 후 시스템 경로로 이동
sudo mv nhn /usr/local/bin/

# 실행 권한 확인
chmod +x /usr/local/bin/nhn

# 설치 확인
nhn version
```

### macOS

```bash
# 빌드 후 시스템 경로로 이동
sudo mv nhn /usr/local/bin/

# 또는 사용자 로컬 bin 디렉토리 사용
mkdir -p ~/bin
mv nhn ~/bin/

# PATH에 추가 (~/.zshrc 또는 ~/.bashrc)
echo 'export PATH="$HOME/bin:$PATH"' >> ~/.zshrc
source ~/.zshrc

# 설치 확인
nhn version
```

### Windows

1. `go build -o nhn.exe main.go` 실행
2. `nhn.exe`를 PATH에 포함된 디렉토리로 이동
   - 예: `C:\Users\<username>\bin\`
3. 환경 변수에서 PATH 확인

```powershell
# PowerShell에서 PATH 확인
$env:PATH -split ';'

# 설치 확인
nhn.exe version
```

---

## 버전 확인

```bash
nhn version
```

출력 예시:
```
nhn version 0.1.0 (built: 2024-01-15T10:30:00Z)
```

---

## 업그레이드

새 버전으로 업그레이드하려면:

```bash
cd nhncli
git pull origin main
go build -o nhn main.go
sudo mv nhn /usr/local/bin/
```

---

## 설치 제거

```bash
# Linux/macOS
sudo rm /usr/local/bin/nhn

# 설정 파일도 삭제하려면
rm -rf ~/.nhn
```

---

## 문제 해결

### Go 버전 오류

```
go: go.mod file indicates go 1.22, but maximum supported version is 1.21
```

Go를 1.22 이상으로 업그레이드하세요:
```bash
# Go 공식 사이트에서 다운로드
# https://go.dev/dl/

# 또는 brew 사용 (macOS)
brew upgrade go
```

### 권한 오류

```
permission denied: /usr/local/bin/nhn
```

`sudo`를 사용하거나 사용자 디렉토리에 설치하세요:
```bash
mkdir -p ~/bin
mv nhn ~/bin/
```

### PATH 오류

명령을 찾을 수 없는 경우:
```bash
# 현재 PATH 확인
echo $PATH

# PATH에 추가
export PATH="/usr/local/bin:$PATH"
```

---

## 다음 단계

설치가 완료되면 [설정 가이드](Configuration.md)에서 인증을 구성하세요.
