package main

import (
	"fmt"

	"vk-auto.com/vk-parser/vk"
	"vk-auto.com/vk-parser/writer"
)

func main() {
	fmt.Println("Getting posts...")

	result, err := vk.GetVkPublicPosts("-40867069")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Posts received")

	cars := GetCars(result)
	for _, c := range cars {
		if c.Id == 0 {
			continue
		}
		err := writer.WriteInfo(&c)
		if err != nil {
			fmt.Println(err)
		}
	}
	fmt.Println("Write success")
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
		}
	}
	return cars
}
