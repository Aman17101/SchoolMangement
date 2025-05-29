package util

import (
	"flag"
	"net/url"
	"os"

	"github.com/Aman17101/SchoolMangement/model"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

var Logger *logrus.Logger

func init() {
	Logger = logrus.New()

	Logger.Out = os.Stdout

	Logger.SetFormatter(&logrus.TextFormatter{
		ForceColors:     true,                //add for highlighting
		TimestampFormat: `json:"created_at"`, // Go's reference time
		FullTimestamp:   true,
	})

}
func SetLogger() {
	logLevel := flag.String(model.LogLevel, model.LogLevelInfo, "log-level ,(debug,info,error ,warning)")

	flag.Parse()
	switch *logLevel {

	case model.LogLevelDebug:
		Logger.SetLevel(logrus.DebugLevel)

	case model.LogLevelError:
		Logger.SetLevel(logrus.ErrorLevel)

	case model.LogLevelWarning:
		Logger.SetLevel(logrus.WarnLevel)

	default:
		Logger.SetLevel(logrus.InfoLevel)

	}
}
func Log(logLevel, packageLevel, functionName string, message, parameter interface{}) {
	switch logLevel {
	case model.LogLevelDebug:
		if parameter != nil {
			Logger.Debugf("packingLevel : %s,functionName :%s, message :%v,parameter :%v\n", packageLevel, functionName, message, parameter)
		} else {
			Logger.Debugf("packingLevel : %s,functionName :%s, message :%v\n", packageLevel, functionName, message)
		}

	case model.LogLevelError:
		if parameter != nil {
			Logger.Errorf("packingLevel : %s,functionName :%s, message :%v,parameter :%v\n", packageLevel, functionName, message, parameter)
		} else {
			Logger.Errorf("packingLevel : %s,functionName :%s, message :%v\n", packageLevel, functionName, message)
		}

	case model.LogLevelWarning:
		if parameter != nil {
			Logger.Warnf("packingLevel : %s,functionName :%s, message :%v,parameter :%v\n", packageLevel, functionName, message, parameter)
		} else {
			Logger.Warnf("packingLevel : %s,functionName :%s, message :%v\n", packageLevel, functionName, message)
		}

	default:
		if parameter != nil {
			Logger.Infof("packingLevel : %s,functionName :%s, message :%v,parameter :%v\n", packageLevel, functionName, message, parameter)
		} else {
			Logger.Infof("packingLevel : %s,functionName :%s, message :%v\n", packageLevel, functionName, message)
		}
	}

} // ConvertQueryParams converts url.Values to map[string]interface{}
func ConvertQueryParams(queryParams url.Values) map[string]interface{} {
	result := make(map[string]interface{})

	for key, values := range queryParams {
		if key == "id" {
			uuid, _ := uuid.Parse(values[0])
			result[key] = uuid
			continue
		}
		if len(values) == 1 {
			result[key] = values[0] // single value, add as string
		} else {
			result[key] = values // multiple values, add as []string
		}
	}

	return result
}
