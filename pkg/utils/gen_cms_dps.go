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
	"github.com/iancoleman/strcase"

	//"k8s.io/apimachinery/pkg/runtime"
	//contrailv1alpha1 "github.com/michaelhenkel/contrail-manager/pkg/apis/contrail/v1alpha1"
	//apiextensionsv1beta1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"

	managerinfra "github.com/operators/contrail-manager-test-1/pkg/controller/contrailmanager"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	apiextensionsv1beta1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
)

var crPrefix = "Contrail"

var version = managerinfra.Version

var directories = [...]string{
	"cassandra",
}

var cmRegexp = regexp.MustCompile("manifests/configmaps/")
var dpRegexp = regexp.MustCompile("manifests/deployments/")

var (
	srcConfigMapDirectory  = "./"
	dstConfigMapDirectory  = "../configs/"
	srcDeploymentDirectory = "./"
	dstDeploymentDirectory = "../configs/"
	srcCRDDirectory        = "../../deploy/crds/"
	dstCRDDirectory        = "../apis/contrail/v1alpha1/"
	srcCRDirectory         = "../../deploy/crds/"
	dstCRDirectory         = "../apis/contrail/v1alpha1/"
)

//go:generate go run gen_cms_dps.go
func main() {
	err := createCms()
	if err != nil {
		fmt.Println(err)
	}
	err = createDeployments()
	if err != nil {
		fmt.Println(err)
	}
	err = createCRDs()
	if err != nil {
		fmt.Println(err)
	}
	err = createCRs()
	if err != nil {
		fmt.Println(err)
	}
}

func getGoFileName(file string) string {
	tempList := strings.Split(file, "/")
	retFile := strings.Split(tempList[len(tempList)-1], ".yaml")[0]
	return strings.ToLower(retFile)
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
				generateConfigMap(utilDirName, file, goFile)
			}
		}
	}
	return nil
}

func createDeployments() error {
	var files []string
	var utilPath string
	var goFile string

	for _, utilDirName := range directories {
		utilPath = srcDeploymentDirectory + utilDirName
		err := filepath.Walk(utilPath, func(path string, info os.FileInfo, err error) error {
			files = append(files, path)
			return nil
		})
		if err != nil {
			panic(err)
		}
		for _, file := range files {
			switch {
			case dpRegexp.MatchString(file):
				goFile = getGoFileName(file)
				generateDeployment(utilDirName, file, goFile)
			}
		}
	}
	return nil
}

func createCRDs() error {
	var files []string
	var utilPath string
	var goFile string

	for _, utilDirName := range directories {
		utilPath = srcCRDDirectory
		err := filepath.Walk(utilPath, func(path string, info os.FileInfo, err error) error {
			files = append(files, path)
			return nil
		})
		if err != nil {
			panic(err)
		}
		var crRegexp1 = regexp.MustCompile(utilDirName)
		var crRegexp2 = regexp.MustCompile("_crd.yaml")
		for _, file := range files {
			switch {
			case crRegexp1.MatchString(file) && crRegexp2.MatchString(file):
				goFile = getGoFileName(file)
				generateCRD(utilDirName, file, goFile)
			}
		}
	}
	return nil
}

func createCRs() error {
	var files []string
	var utilPath string
	var goFile string

	for _, utilDirName := range directories {
		utilPath = srcCRDirectory
		err := filepath.Walk(utilPath, func(path string, info os.FileInfo, err error) error {
			files = append(files, path)
			return nil
		})
		if err != nil {
			panic(err)
		}
		var crRegexp1 = regexp.MustCompile(utilDirName)
		var crRegexp2 = regexp.MustCompile("_cr.yaml")
		for _, file := range files {
			switch {
			case crRegexp1.MatchString(file) && crRegexp2.MatchString(file):
				goFile = getGoFileName(file)
				generateCR(utilDirName, file, goFile)
			}
		}
	}
	return nil
}

func generateConfigMap(dirName string, srcPath string, fileName string) {
	dstPath := dstConfigMapDirectory + dirName + "/" + fileName + ".go"
	var packageTemplate = template.Must(template.New("").Parse(`package {{ .Kind }}
	
import (
	"github.com/ghodss/yaml"
	corev1 "k8s.io/api/core/v1"
)

var yamlData{{ .DataKind }}= {{ .YamlData }}

func {{ .FileName }}() *corev1.ConfigMap {
	fileData := yamlData{{ .DataKind }}
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
		FileName string
		DataKind string
	}{
		YamlData: yamlDataQuoted,
		Kind:     dirName,
		DataKind: dirName + strcase.ToCamel(fileName),
		FileName: strcase.ToCamel(fileName),
	})
}

func generateDeployment(dirName string, srcPath string, fileName string) {
	dstPath := dstConfigMapDirectory + dirName + "/" + fileName + ".go"
	var packageTemplate = template.Must(template.New("").Parse(`package {{ .Kind }}
	
import (
	"github.com/ghodss/yaml"
	appsv1 "k8s.io/api/apps/v1"
)

var yamlData{{ .DataKind }}= {{ .YamlData }}
func {{ .FileName }}() *appsv1.Deployment{
	deployment := appsv1.Deployment{}
	err := yaml.Unmarshal([]byte(yamlData{{ .DataKind }}), &deployment)
	if err != nil {
		panic(err)
	}
	jsonData, err := yaml.YAMLToJSON([]byte(yamlData{{ .DataKind }}))
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal([]byte(jsonData), &deployment)
	if err != nil {
		panic(err)
	}
	return &deployment
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
	var configMap appsv1.Deployment
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
		FileName string
		DataKind string
	}{
		YamlData: yamlDataQuoted,
		Kind:     dirName,
		FileName: strcase.ToCamel(fileName),
		DataKind: dirName + strcase.ToCamel(fileName),
	})

}

func generateCRD(dirName string, srcPath string, fileName string) {
	// dstPath := dstCRDDirectory + dirName + "/crds/" + fileName + ".go"
	dstPath := dstCRDDirectory + fileName + ".go"
	var packageTemplate = template.Must(template.New("").Parse(`package {{ .Kind}}
	
import (
	"github.com/ghodss/yaml"
	apiextensionsv1beta1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
)

var yamlData{{ .DataKind }} = {{ .YamlData }}
func {{ .DataKind }}() *apiextensionsv1beta1.CustomResourceDefinition{
	crd := apiextensionsv1beta1.CustomResourceDefinition{}
	err := yaml.Unmarshal([]byte(yamlData{{ .DataKind }}), &crd)
	if err != nil {
		panic(err)
	}
	jsonData, err := yaml.YAMLToJSON([]byte(yamlData{{ .DataKind }}))
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal([]byte(jsonData), &crd)
	if err != nil {
		panic(err)
	}
	return &crd
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
	var crd apiextensionsv1beta1.CustomResourceDefinition
	err = yaml.Unmarshal([]byte(jsonData), &crd)
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
		FileName string
		DataKind string
	}{
		YamlData: yamlDataQuoted,
		Kind:     version,
		DataKind: strcase.ToCamel(dirName) + strcase.ToCamel(fileName),
		FileName: strcase.ToCamel(fileName),
	})
}

func generateCR(dirName string, srcPath string, fileName string) {
	// dstPath := dstCRDDirectory + dirName + "/crds/" + fileName + ".go"
	dstPath := dstCRDDirectory + fileName + ".go"
	var packageTemplate = template.Must(template.New("").Parse(`package {{ .Kind }}
	
import (
	"github.com/ghodss/yaml"
)

var yamlData{{ .DataKind }} = {{ .YamlData }}
func {{ .DataKind }}() *{{ .Prefix }}{{ .CamelKind }}{
	cr := {{ .Prefix }}{{ .CamelKind }}{}
	err := yaml.Unmarshal([]byte(yamlData{{ .DataKind }}), &cr)
	if err != nil {
		panic(err)
	}
	jsonData, err := yaml.YAMLToJSON([]byte(yamlData{{ .DataKind }}))
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal([]byte(jsonData), &cr)
	if err != nil {
		panic(err)
	}
	return &cr
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
	var crd apiextensionsv1beta1.CustomResourceDefinition
	err = yaml.Unmarshal([]byte(jsonData), &crd)
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
		YamlData  string
		Kind      string
		FileName  string
		DataKind  string
		Prefix    string
		CamelKind string
	}{
		YamlData:  yamlDataQuoted,
		Kind:      version,
		DataKind:  strcase.ToCamel(dirName) + strcase.ToCamel(fileName),
		FileName:  strcase.ToCamel(fileName),
		CamelKind: strcase.ToCamel(dirName),
		Prefix:    crPrefix,
	})
}
