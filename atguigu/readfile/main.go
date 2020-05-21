package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	fDir := "D:/haha1.txt"
	file, _ := os.OpenFile(fDir, os.O_APPEND, 0666)
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		fmt.Print(str)
	}

	str := "hello abc abcg\r\n"
	writer := bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		writer.WriteString(str)
		fmt.Print("********")
	}
	writer.Flush() //很重要，写的时候一定要flush

}
