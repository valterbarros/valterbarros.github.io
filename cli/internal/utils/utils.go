package utils

import (
	"encoding/json"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

func SetEnvFromFile() {
	env, _ := GetFromRoot(".env")
	data := GetFileData(env)
	vars := strings.SplitSeq(data, "\n")
	for variable := range vars {
		spt := strings.Split(variable, "=")
		os.Setenv(spt[0], spt[1])
	}
}

func WriteLog(line string) {
	file := "/tmp/go-api-log.txt"
	data := GetFileData(file)
	data += time.Now().Format("2006-01-02 15:04:05") + " " + line + "\n"
	SaveFile(file, data)
}

func Capitalize(s string) string {
	return strings.ToUpper(s[:1]) + s[1:]
}

func DecodeJson(data io.Reader, variable any) error {
	decoder := json.NewDecoder(data)

	if err := decoder.Decode(&variable); err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}

func EncodeJson(data io.Writer, variable any) error {
	decoder := json.NewEncoder(data)

	if err := decoder.Encode(&variable); err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}
