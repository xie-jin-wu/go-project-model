package program

import (
	"errors"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// GetProgramInfo 获取当前程序信息
func GetProgramInfo() (string, error) {
	file, err := exec.LookPath(os.Args[0])
	if err != nil {
		return "", err
	}
	path, err := filepath.Abs(file)
	if err != nil {
		return "", err
	}
	return path, nil
}

// GetProgramName 获取当前程序名
func GetProgramName() (string, error) {
	path, err := GetProgramInfo()
	if err != nil {
		return "", err
	}
	i := strings.LastIndex(path, "/")
	if i < 0 {
		i = strings.LastIndex(path, "\\")
	}
	if i < 0 {
		return "", errors.New(`error: Can't find "/" or "\"`)
	}
	return path[i+1:], nil
}

// GetProgramPath 获取当前程序路径
func GetProgramPath() (string, error) {
	path, err := GetProgramInfo()
	if err != nil {
		return "", err
	}
	i := strings.LastIndex(path, "/")
	if i < 0 {
		i = strings.LastIndex(path, "\\")
	}
	if i < 0 {
		return "", errors.New(`error: Can't find "/" or "\"`)
	}
	return path[0 : i+1], nil
}
