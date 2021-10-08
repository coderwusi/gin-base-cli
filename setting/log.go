package setting

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

var log = logrus.New()

func LogInit() {
	file, err := os.OpenFile("./logs/logon_"+time.Now().Format("2006-01-02")+".log", os.O_WRONLY|os.O_APPEND|os.O_SYNC|os.O_CREATE|os.O_RDWR, 0766)
	fmt.Println(err)
	if err != nil {
		logrus.Fatal("log init failed")

	}
	if err == nil {
		log.Out = file
	}
	//log.WithFields(logrus.Fields{
	//	"animal": "walrus",
	//	"size": 10,
	//}).Info("blog log")
	//file.Close()
}
