package config

import (
	"bufio"
	"io"
	"os"
	"regexp"
)

func Properties(file *os.File) *map[string]string {
	properties := make(map[string]string)
	reader := bufio.NewReader(file)
	matcher, _ := regexp.Compile("(^\\w+)(\\s*=\\s*)(.*)")
	var buff []byte
	var err error
	var isPrefix bool
	var content string
	for {
		content = ""
		for {
			buff, isPrefix, err = reader.ReadLine()
			content += string(buff)

			if !isPrefix {
				break
			}
		}
		kv := matcher.FindAllStringSubmatch(content, 1)
		if len(kv) > 0 {
			properties[kv[0][1]] = kv[0][3]
		}
		if io.EOF == err {
			break
		}
	}

	return &properties
}
