package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"sync"
	"time"
)

const (
	kernelSize = 3 // Размер ядра свёртки
)

var gaussianKernel = [kernelSize][kernelSize]float64{
	{1 / 16.0, 2 / 16.0, 1 / 16.0},
	{2 / 16.0, 4 / 16.0, 2 / 16.0},
	{1 / 16.0, 2 / 16.0, 1 / 16.0},
}

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

	bounds := img.Bounds()
	newImg := image.NewRGBA(bounds)

	start := time.Now()

	filterGaussian(img, newImg)

	duration := time.Since(start)
	fmt.Printf("Время выполнения: %v\n", duration)

	outFile, err := os.Create("output_gaussian.png")
	if err != nil {
		fmt.Println("Ошибка при создании файла:", err)
		return
	}
	defer outFile.Close()

	err = png.Encode(outFile, newImg)
	if err != nil {
		fmt.Println("Ошибка при сохранении изображения:", err)
		return
	}

	fmt.Println("Изображение успешно обработано и сохранено.")
}

// filterGaussian применяет матричный фильтр к изображению
func filterGaussian(src image.Image, dst *image.RGBA) {
	var wg sync.WaitGroup
	bounds := src.Bounds()

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		wg.Add(1)
		go func(y int) {
			defer wg.Done()
			for x := bounds.Min.X; x < bounds.Max.X; x++ {
				dst.Set(x, y, applyKernel(src, x, y))
			}
		}(y) // Передаем значение y в анонимную функцию
	}

	wg.Wait()
}

//применяет ядро свёртки к пикселю
func applyKernel(img image.Image, x, y int) color.Color {
	var r, g, b, a float64
	kernelRadius := kernelSize / 2

	for ky := -kernelRadius; ky <= kernelRadius; ky++ {
		for kx := -kernelRadius; kx <= kernelRadius; kx++ {
			nx, ny := x+kx, y+ky
			if nx >= 0 && nx < img.Bounds().Max.X && ny >= 0 && ny < img.Bounds().Max.Y {
				c := img.At(nx, ny)
				cr, cg, cb, ca := c.RGBA()
				weight := gaussianKernel[ky+kernelRadius][kx+kernelRadius]
				r += float64(cr>>8) * weight
				g += float64(cg>>8) * weight
				b += float64(cb>>8) * weight
				a += float64(ca>>8) * weight
			}
		}
	}

	return color.RGBA{
		R: uint8(clamp(r)),
		G: uint8(clamp(g)),
		B: uint8(clamp(b)),
		A: uint8(clamp(a)),
	}
}

func clamp(value float64) float64 {
	if value < 0 {
		return 0
	}
	if value > 255 {
		return 255
	}
	return value
}
