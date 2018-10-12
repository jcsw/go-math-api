package properties

import (
	"io/ioutil"
	"os"
	"time"

	"github.com/jcsw/go-math-api/pkg/driver/syslog"
	"gopkg.in/yaml.v2"
)

// Properties define the properties values
type Properties struct {
	ServerPort int               `yaml:"serverPort"`
	MongoDB    MongoDBProperties `yaml:"mongodb"`
}

// MongoDBProperties define the mongoDB properties values
type MongoDBProperties struct {
	Hosts     []string      `yaml:"hosts"`
	Username  string        `yaml:"username"`
	Password  string        `yaml:"password"`
	Database  string        `yaml:"database"`
	Timeout   time.Duration `yaml:"timeout"`
	PoolLimit uint16        `yaml:"poolLimit"`
}

// AppProperties the loaded properties values
var AppProperties Properties

// LoadProperties load properties in AppProperties
func LoadProperties(env string) {

	pwd, _ := os.Getwd()
	fileProperties, err := ioutil.ReadFile(pwd + "/properties/" + env + ".yaml")
	if err != nil {
		syslog.Fatal("p=properties f=LoadProperties \n%v", err)
	}

	err = yaml.UnmarshalStrict(fileProperties, &AppProperties)
	if err != nil {
		syslog.Fatal("p=properties f=LoadProperties \n%v", err)
	}

	syslog.Info("p=properties f=LoadProperties \n%+v", AppProperties)
}
