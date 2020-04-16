package log

import (
	"github.com/sirupsen/logrus"
)

var Log = logrus.New()

func CheckErr(content string, err error) bool {
	if err != nil {
		Log.Error(content + " " + err.Error())
		return true
	}
	return false
}
func SetLogFile() {
	Log.Formatter = new(logrus.TextFormatter)
	Log.Level = logrus.TraceLevel
}
