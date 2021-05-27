package k8stest

import (
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
)

// .conf 文件测试
func ConfTest() {
	file_path := "/Users/cwb/Documents/work/beeee/go/src/go-test/k8stest/config"
	//file_path := "/var/run/secrets/kubernetes.io/serviceaccount/token"
	tokenBytes, err := ioutil.ReadFile(file_path)
	if os.IsNotExist(err) {
		logrus.Infof("Running in single server mode, will not peer connections")
		return
	} else if err != nil {
		return
	}

	logrus.Infof("Running in clustered mode with ID %s", string(tokenBytes))
}
