# punydecode

**punydecode**는 파일 또는 표준 입력으로부터 퓨니코드(Punycode) 문자열을 읽어  
해당 유니코드 문자열로 디코딩하는 간단한 Go CLI 도구입니다.

예) `xn--fsq.com` → `例.com`

---

## ✨ 기능

- 파일 또는 파이프(stdin) 입력 처리
- 한 줄에 여러 개의 푸니코드도 모두 디코드
- `-h`, `--help` : 도움말 출력
- `-v`, `--version` : 버전 출력

---

## 🔧 설치

Go 환경이 있다면 다음 명령어로 설치할 수 있습니다:

```bash
go install github.com/yourname/punydecode@latest
```

또는 로컴에서 직접 빌드:

```bash
git clone https://github.com/yourname/punydecode.git
cd punydecode
go build -o punydecode
```

---

## 🛠️ 사용법

```bash
punydecode [옵션] [파일]
```

- 파일을 인자로 넘기거나
- 파이프로 입력 받을 수 있습니다.

---

## 💡 예시

### 파일 직접 입력

```bash
punydecode domains.txt
```

### 표준 입력 (리디렉션)

```bash
punydecode < domains.txt
```

### 표준 입력 (파이프)

```bash
cat domains.txt | punydecode
```

> ⚠️ PowerShell에서는 `< domains.txt` 방식이 동작하지 않을 수 있습니다.  
> 대신 `Get-Content domains.txt | punydecode`를 사용하세요.

---

## 🔪 출력 예시

입력 (`domains.txt`):

```
xn--fsq.com
hello xn--eckwd4c7c.com world
```

출력:

```
例.com
hello 日本.com world
```

---

## 📄 옵션

| 옵션              | 설명           |
| ----------------- | -------------- |
| `-h`, `--help`    | 도움말 출력    |
| `-v`, `--version` | 버전 정보 출력 |

---

## 📜 라이센스

MIT License

---

## ✍️ 작성자

by [박영기] – 사이버 보안 & 디지털 포렌식 전문가
