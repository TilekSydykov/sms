package i18n

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

const (
	folder      = "langs"
	fileName    = "i18n.json"
	lang        = "lang"
	defaultLang = "ru"
)

type Lang struct {
	Code        string
	Translation string
	Values      map[string]string
}

type I18n struct {
	Langs map[string]Lang
}

func addNames(path string, obj map[string]interface{}) map[string]string {
	var target = make(map[string]string)
	for key, value := range obj {
		keyUpper := strings.Title(strings.ToLower(key))
		switch value.(type) {
		case string:
			target[path+keyUpper] = value.(string)
		case map[string]interface{}:
			res := addNames(path+keyUpper, value.(map[string]interface{}))
			for k, v := range res {
				target[k] = v
			}
		}
	}
	return target
}

func NewI18n() *I18n {
	i := I18n{}
	i.Init()
	return &i
}

func (i *I18n) Init() {
	now := time.Now()
	folderName := "./i18n/" + folder + "/"
	files, err := ioutil.ReadDir(folderName)
	if err != nil {
		panic(err)
	}
	i.Langs = make(map[string]Lang)
	for _, f := range files {
		if f.IsDir() {
			dirName := f.Name()
			fileDir := folderName + dirName + "/" + fileName
			dat, err := os.ReadFile(fileDir)
			if err == nil {
				var langData map[string]interface{}
				err := json.Unmarshal(dat, &langData)
				if err != nil {
					panic(err)
				}
				tName := langData[lang]
				if tName == nil {
					fmt.Println("i18n: in " + dirName + " required key \"" + lang + "\" not found ")
					tName = "Unknown"
				}

				i.Langs[dirName] = Lang{
					dirName,
					tName.(string),
					addNames("", langData),
				}
			}
		}
	}
	fmt.Printf("[i18n] Initialised in %s\n", time.Since(now))
}

func (i *I18n) Resolve(code string, m map[string]interface{}) string {
	word, ok := i.Langs[m["Lang"].(string)].Values[code]
	if !ok {
		return "!-!"
	}
	return word
}
