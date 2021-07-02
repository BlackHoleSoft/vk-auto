package main

import (
	"fmt"

	"vk-auto.com/vk-parser/vk"
	"vk-auto.com/vk-parser/writer"
)

func main() {
	fmt.Println("Getting posts...")

	postList := []string{
		"-40867069",
		"-101275644",
		"-174283946",
		"-124405175",
	}

	for _, p := range postList {
		HandlePublic(p)
	}

	fmt.Println("Write success")
}

func HandlePublic(pub string) {
	result, err := vk.GetVkPublicPosts(pub)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Posts received")

	cars := GetCars(result)
	err = writer.WriteJson(cars)
	if err != nil {
		fmt.Println(err)
	}
	for _, c := range cars {
		if c.Id == 0 {
			continue
		}
		err := writer.WriteInfo(&c)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func GetCars(posts []vk.VkPost) []writer.CarInfo {
	brands, models, err := writer.GetBrandsAndModels()
	if err != nil {
		panic(err)
	}

	cars := make([]writer.CarInfo, len(posts))
	for i, p := range posts {
		brand, model, err := writer.FindNames(p.Text, brands, models)
		if err != nil {
			continue
		}
		year := writer.GetYear(p.Text)
		price := writer.GetPrice(p.Text)
		photos := make([]string, len(p.Attachments))
		for i, a := range p.Attachments {
			photos[i] = a.Photo.Photo_1280
			if len(photos[i]) == 0 {
				photos[i] = a.Photo.Photo_807
			}
		}
		cars[i] = writer.CarInfo{
			Id:          p.Id,
			Brand:       brand,
			Model:       model,
			Description: p.Text,
			Photos:      photos,
			Year:        year,
			Price:       price,
		}
	}
	return cars
}
