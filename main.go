// main.go
package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"

	"golang.org/x/net/idna"
)

const version = "1.0.0"

func decodePunycode(text string) string {
	re := regexp.MustCompile(`xn--[a-zA-Z0-9-]+`)
	return re.ReplaceAllStringFunc(text, func(match string) string {
		decoded, err := idna.ToUnicode(match)
		if err != nil {
			log.Printf("디코딩 실패: %s: %v\n", match, err)
			return match
		}
		return decoded
	})
}

func processScanner(scanner *bufio.Scanner) {
	for scanner.Scan() {
		line := scanner.Text()
		decodedLine := decodePunycode(line)
		fmt.Println(decodedLine)
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("입력 읽기 중 오류: %v", err)
	}
}

func printHelp() {
	fmt.Println(`이름  
punydecode - 파일 또는 표준 입력으로부터 퓨니코드 문자열을 디코딩

사용법  
punydecode [옵션] [파일]

설명  
punydecode는 지정된 파일 또는 표준 입력으로부터 한 줄씩 읽으며, 그 안에 포함된 퓨니코드 문자열(예: xn--fsq.com)을 해당 유니코드 문자열로 디코딩합니다.  
파일이 지정되지 않은 경우, 표준 입력으로부터 읽습니다.

옵션  

-h, --help  
    사용 방법을 보여주는 도움말 메시지를 출력하고 종료합니다.

-v, --version  
    현재 프로그램 버전을 출력하고 종료합니다.

예시  
punydecode domains.txt  
punydecode < domains.txt  
cat domains.txt | punydecode`)
}

func printVersion() {
	fmt.Printf("punydecode version %s\n", version)
}

func main() {
	flag.Usage = printHelp
	help := flag.Bool("h", false, "")
	helpLong := flag.Bool("help", false, "")
	showVersion := flag.Bool("v", false, "")
	showVersionLong := flag.Bool("version", false, "")
	flag.Parse()

	if *help || *helpLong {
		printHelp()
		return
	}

	if *showVersion || *showVersionLong {
		printVersion()
		return
	}

	args := flag.Args()
	var scanner *bufio.Scanner

	if len(args) >= 1 {
		file, err := os.Open(args[0])
		if err != nil {
			log.Fatalf("파일 열기 실패: %v", err)
		}
		defer file.Close()
		scanner = bufio.NewScanner(file)
	} else {
		scanner = bufio.NewScanner(os.Stdin)
	}

	processScanner(scanner)
}
