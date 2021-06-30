package writer

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

type CarInfo struct {
	Brand       string
	Model       string
	Description string
	Photos      []string
	Id          int
}

func WriteInfo(car *CarInfo) error {
	brandNames := strings.Split(car.Brand, " ")
	modelNames := strings.Split(car.Model, " ")
	var brand string
	var model string

	if len(brandNames) > 0 {
		brand = brandNames[0]
	} else {
		brand = "Other"
	}

	if len(modelNames) > 0 {
		model = modelNames[0]
	} else {
		model = "Other"
	}

	fmt.Println(fmt.Sprintf("Found car: %s %s", brand, model))

	path := fmt.Sprintf("cars")
	var content strings.Builder
	content.WriteString(fmt.Sprintf("<p>%s</p>", car.Description))
	for _, s := range car.Photos {
		content.WriteString(fmt.Sprintf("<img src=\"%s\">", s))
	}

	if err := os.MkdirAll(path, os.ModeDir); err != nil {
		return err
	}
	filename := fmt.Sprintf("/%s_%s_%v.html", brand, model, car.Id)
	if err := os.WriteFile(path+filename, []byte(content.String()) /*os.ModeAppend*/, 0644); err != nil {
		return err
	}
	fmt.Println("Write File: " + path + filename)
	return nil
}

func GetBrandsAndModels() (string, string, error) {
	brands, err := os.ReadFile("data/brands.csv")
	if err != nil {
		return "", "", err
	}
	models, err := os.ReadFile("data/models.csv")
	if err != nil {
		return "", "", err
	}
	return strings.ToLower(string(brands)), strings.ToLower(string(models)), nil
}

func FindNames(text string, brands string, models string) (string, string, error) {
	brandsList := strings.Split(string(brands), "\r\n")
	modelsList := strings.Split(string(models), "\r\n")

	shortLen := 100
	if len(text) < shortLen {
		shortLen = len(text) / 2
	}
	brand := GetName(brandsList, text[:shortLen])
	if brand == "Other" {
		brand = GetName(brandsList, text)
	}
	model := GetName(modelsList, text[:shortLen])
	if model == "Other" {
		model = GetName(brandsList, text)
	}
	return brand, model, nil
}

func GetName(list []string, text string) string {
	if len(list) == 0 {
		return "Other"
	}

	for _, b := range list {
		splitted := strings.Split(b, " ")
		for _, s := range splitted {
			//fmt.Println("|" + s + "|")
			re := regexp.MustCompile(`(\A|\s)+` + s)
			if re.Match([]byte(strings.ToLower(text))) {
				return splitted[0]
			}
		}
	}
	return "Other"
}

func Clear() error {
	if err := os.Remove("cars"); err != nil {
		return err
	}
	return nil
}
