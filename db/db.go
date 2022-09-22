package db

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func WriteFile(nameFile string, timer time.Duration, value string) {

	time.Sleep(timer * time.Second)
	content, err := os.ReadFile(nameFile)
	if err != nil {
		_, err := os.Create(nameFile)
		if err != nil {
			log.Fatal(err)
		}
	}

	a := strings.Split(string(content), `\n`)
	fmt.Println(a)
	data := value
	a = append(a, data)
	file, err := os.OpenFile(nameFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	datawriter := bufio.NewWriter(file)
	datawriter.WriteString(data + "\n")
	datawriter.Flush()
	defer file.Close()

}
