package writer

import (
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type CarInfo struct {
	Brand       string
	Model       string
	Year        int
	Price       int
	Description string
	Photos      []string
	Id          int
}

func WriteJson(cars []CarInfo) error {
	path := fmt.Sprintf("cars/json/")

	content, err := json.Marshal(cars)
	if err != nil {
		return err
	}

	if err := os.MkdirAll(path, os.ModeDir); err != nil {
		return err
	}

	if err := os.WriteFile(path+"cars.json", []byte(content) /*os.ModeAppend*/, 0644); err != nil {
		return err
	}
	return nil
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
	filename := fmt.Sprintf("/%s_%s_%v_%v_%v.html", brand, model, car.Year, car.Price, car.Id)
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

func GetYear(text string) int {
	re := regexp.MustCompile(`\b\d{4}\b`)
	matches := re.FindAll([]byte(strings.ToLower(text)), 10)
	for _, m := range matches {
		val, err := strconv.Atoi(string(m))
		if err != nil {
			continue
		}
		if val >= 1900 && val <= 2030 {
			return val
		}
	}

	return 0
}

func GetPrice(text string) int {
	re := regexp.MustCompile(`(?i)(цена:?\s*\d+)|(\d+\s*к\b)|(\d+\s*₽)|(\d+\s*р)|(\d+\s*т\.?р)|(\d+\s*руб)`)
	matches := re.FindAll([]byte(strings.ToLower(text)), 10)
	for _, m := range matches {
		mre := regexp.MustCompile(`\d+`)
		val, err := strconv.Atoi(string(mre.Find(m)))
		if err != nil {
			continue
		}
		if val >= 10 && val < 1000 {
			val = val * 1000
		}
		if val >= 10 {
			return val
		}
	}

	return 0
}

func Clear() error {
	if err := os.Remove("cars"); err != nil {
		return err
	}
	return nil
}
