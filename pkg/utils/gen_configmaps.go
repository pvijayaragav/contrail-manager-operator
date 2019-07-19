package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"text/template"

	"github.com/ghodss/yaml"
	//"k8s.io/apimachinery/pkg/runtime"
	//contrailv1alpha1 "github.com/michaelhenkel/contrail-manager/pkg/apis/contrail/v1alpha1"
	//apiextensionsv1beta1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"

	corev1 "k8s.io/api/core/v1"
)

var directories = [...]string{
	"cassandra",
}

var cmRegexp = regexp.MustCompile("manifests/configmaps/")

const (
	srcConfigMapDirectory = "./"
	dstConfigMapDirectory = "../configs/"
)

//go:generate go run gen_configmaps.go
func main() {
	err := createCms()
	if err != nil {
		fmt.Println(err)
	}
}

func createCms() error {
	var files []string
	var utilPath string
	var goFile string

	for _, utilDirName := range directories {
		utilPath = srcConfigMapDirectory + utilDirName
		err := filepath.Walk(utilPath, func(path string, info os.FileInfo, err error) error {
			files = append(files, path)
			return nil
		})
		if err != nil {
			panic(err)
		}
		for _, file := range files {
			switch {
			case cmRegexp.MatchString(file):
				goFile = getGoFileName(file)
				generateConfigMap(utilDirName, file, dstConfigMapDirectory+utilDirName+"/"+goFile+".go")
			}
		}
	}
	return nil
}

func getGoFileName(file string) string {
	tempList := strings.Split(file, "/")
	retFile := strings.Split(tempList[len(tempList)-1], ".yaml")[0]
	return retFile
}

func generateConfigMap(dirName string, srcPath string, dstPath string) {
	var packageTemplate = template.Must(template.New("").Parse(`package {{ .Kind }}
	
import (
	"github.com/ghodss/yaml"
	corev1 "k8s.io/api/core/v1"
)

var yamlData{{ .Kind }}= {{ .YamlData }}

func GetConfigMap() *corev1.ConfigMap {
	fileData := configMap
	yamlData := string(fileData)
	cm := corev1.ConfigMap{}
	err := yaml.Unmarshal([]byte(yamlData), &cm)
	if err != nil {
		panic(err)
	}
	jsonData, err := yaml.YAMLToJSON([]byte(yamlData))
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal([]byte(jsonData), &cm)
	if err != nil {
		panic(err)
	}
	return &cm
}
	`))

	yamlData, err := ioutil.ReadFile(srcPath)
	if err != nil {
		panic(err)
	}

	jsonData, err := yaml.YAMLToJSON([]byte(yamlData))
	if err != nil {
		panic(err)
	}
	var configMap corev1.ConfigMap
	err = yaml.Unmarshal([]byte(jsonData), &configMap)
	if err != nil {
		panic(err)
	}
	f, err := os.Create(dstPath)
	if err != nil {
		panic(err)
	}

	yamlDataQuoted := fmt.Sprintf("`\n")
	yamlDataQuoted = yamlDataQuoted + string(yamlData)
	yamlDataQuoted = yamlDataQuoted + "`"
	packageTemplate.Execute(f, struct {
		YamlData string
		Kind     string
	}{
		YamlData: yamlDataQuoted,
		Kind:     dirName,
	})
}
