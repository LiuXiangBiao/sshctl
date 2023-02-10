package service

import (
	"github.com/spf13/cobra"
	"io/ioutil"
	"log"
	"strings"
)

type ConRes struct {
	Host   string
	result string
}

func SplitString(str string) (strList []string) {
	if str == "" {
		return
	}
	if strings.Contains(str, ",") {
		strList = strings.Split(str, ",")
	} else {
		strList = strings.Split(str, ";")
	}
	return
}

func MustFlag(name, t string, cmd *cobra.Command) interface{} {
	switch t {
	case "string":
		if v, err := cmd.Flags().GetString(name); err == nil && v != "" {
			return v
		}
	case "int":
		if v, err := cmd.Flags().GetInt(name); err == nil && v != 0 {
			return v
		}
	}
	log.Fatal(name, " is required")
	return nil
}

func Getfile(filePath string) ([]string, error) {
	result := []string{}
	b, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Println("read file ", filePath, err)
		return result, err
	}
	s := string(b)
	for _, lineStr := range strings.Split(s, "\n") {
		lineStr = strings.TrimSpace(lineStr)
		if lineStr == "" {
			continue
		}
		result = append(result, lineStr)
	}
	return result, nil
}
