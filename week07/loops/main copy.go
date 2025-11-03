package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func maina() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("이름 입력: ")
	input, _ := reader.ReadString('\n') // \n 포함해서 입력받음
	input = strings.TrimSpace(input)    // TrimSpace: 양쪽 공백/줄바꿈 제거
	fmt.Println("입력된 이름:", input)
}
