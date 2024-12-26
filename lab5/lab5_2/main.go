package main

import (
	"fmt"
	"image/color"
	"image/draw"
	"image/png"
	"os"
	"time"
)

func main() {
	file, err := os.Open("input.png")
	if err != nil {
		fmt.Println("Ошибка при открытии файла:", err)
		return
	}
	defer file.Close()

	img, err := png.Decode(file)
	if err != nil {
		fmt.Println("Ошибка при декодировании изображения:", err)
		return
	}

	drawImg, ok := img.(draw.RGBA64Image)
	if !ok {
		fmt.Println("Преобразование не удалось")
		return
	}

	start := time.Now()

	filter(drawImg)

	duration := time.Since(start)
	fmt.Printf("Время выполнения: %v\n", duration)

	outFile, err := os.Create("output.png")
	if err != nil {
		fmt.Println("Ошибка при создании файла:", err)
		return
	}
	defer outFile.Close()

	err = png.Encode(outFile, drawImg)
	if err != nil {
		fmt.Println("Ошибка при сохранении изображения:", err)
		return
	}

	fmt.Println("Изображение успешно обработано и сохранено.")
}

func filter(img draw.RGBA64Image) {
	bounds := img.Bounds()
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			c := img.RGBA64At(x, y)
			grayValue := uint16((uint32(c.R) + uint32(c.G) + uint32(c.B)) / 3)
			img.SetRGBA64(x, y, color.RGBA64{R: grayValue, G: grayValue, B: grayValue, A: c.A})
		}
	}
}
