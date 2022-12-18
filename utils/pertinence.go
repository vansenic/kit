package utils

import (
	"fmt"
	"path"
	"runtime"
	"time"
)

const Skip = 1

func CompareTime(s1 string, s2 string) bool {
	t1, err := time.Parse("2006-01-02 15:04:05", s1)
	t2, err := time.Parse("2006-01-02 15:04:05", s2)
	if err != nil {
		return false
	}
	if t1.After(t2) {
		return true
	}
	return false
}

func TimeDiff(s1, s2 string, diff int64) bool {
	t1, err := time.ParseInLocation("2006-01-02 15:04:05", s1, time.Local)
	t2, err := time.ParseInLocation("2006-01-02 15:04:05", s2, time.Local)
	if err != nil {
		return false
	}
	if t2.After(t1) {
		diffTime := (t2.Unix() - t1.Unix()) / diff
		if diffTime > 1 {
			return true
		}
	}
	return false
}

func Line() string {
	pc, file, line, ok := runtime.Caller(Skip)
	if !ok {
		return "null"
	}
	funcName := runtime.FuncForPC(pc).Name()
	filename := path.Base(file)
	return fmt.Sprintf("file:%s, function:%s, line:%d", filename, funcName, line)
}
