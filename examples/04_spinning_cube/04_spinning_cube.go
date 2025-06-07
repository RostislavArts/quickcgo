package main

import (
	"fmt"
	"math"
    "time"

	"github.com/RostislavArts/quickcgo/quickcg"
)

const (
	renderW = 200
	renderH = 200
	screenW = 600
	screenH = 600
)

var (
	cubeWidth       float64 = 50
	distanceFromCam int = 150
	incrementSpeed  float64 = 0.5

    sinA, cosA float64 
    sinB, cosB float64 
    sinC, cosC float64 
)

func calculateX(i, j, k float64) float64 {
	return j * sinA * sinB * cosC -
		   k * cosA * sinB * cosC +
		   j * cosA * sinC +
		   k * sinA * sinC +
		   i * cosB * cosC
}

func calculateY(i, j, k float64) float64 {
	return j * cosA * cosC +
		   k * sinA * cosC -
		   j * sinA * sinB * sinC +
		   k * cosA * sinB * sinC -
		   i * cosB * sinC
}

func calculateZ(i, j, k float64) float64 {
	return k * cosA * cosB -
		   j * sinA * cosB +
		   i * sinB
}

func calculateForSurface(cubeX, cubeY, cubeZ float64, color quickcg.ColorRGB,
zBuffer *[renderW * renderH]float64, buffer *[renderW * renderH]quickcg.ColorRGB) {
	x := calculateX(cubeX, cubeY, cubeZ)
	y := calculateY(cubeX, cubeY, cubeZ)
	z := calculateZ(cubeX, cubeY, cubeZ) + float64(distanceFromCam)

	if z <= 0 {
		return
	}

    zp := 40.0

    ooz := 1 / z
    xp := int(float64(renderW) / 2 + zp * ooz * x * 2)
    yp := int(float64(renderH) / 2 + zp * ooz * y * 2)

	if xp < 0 || xp >= renderW || yp < 0 || yp >= renderH {
		return
	}

    idx := xp + yp * renderW
	if ooz > zBuffer[idx] {
		zBuffer[idx] = ooz
		buffer[idx] = color
	}
}

func main() {
	s, err := quickcg.NewScreen(screenW, screenH, false, "Spinning Pixel Cube")
	if err != nil {
		fmt.Println(err)
		return
	}

	var a, b, c float64
    scaleX := screenW / renderW
    scaleY := screenH / renderH

	var zBuffer [renderW * renderH]float64
	var buffer [renderW * renderH]quickcg.ColorRGB

	colorRed := quickcg.ColorRGB{R: 255, G: 0, B: 0}
	colorGreen := quickcg.ColorRGB{R: 0, G: 255, B: 0}
	colorBlue := quickcg.ColorRGB{R: 0, G: 0, B: 255}
	colorYellow := quickcg.ColorRGB{R: 255, G: 255, B: 0}
	colorOrange := quickcg.ColorRGB{R: 255, G: 100, B: 0}
	colorWhite := quickcg.ColorRGB{R: 255, G: 255, B: 255}

	colorBlack := quickcg.ColorRGB{R: 0, G: 0, B: 0}

    newTime := time.Now()
	for !quickcg.Done(16) {
		if quickcg.KeyDown(quickcg.KEY_ESCAPE) {
			return
		}

		for i := range zBuffer {
			zBuffer[i] = 0
			buffer[i] = quickcg.ColorRGB{R: 0, G: 0, B: 0}
		}

        sinA, cosA = math.Sin(a), math.Cos(a)
        sinB, cosB = math.Sin(b), math.Cos(b)
        sinC, cosC = math.Sin(c), math.Cos(c)

		for cubeX := -cubeWidth; cubeX < cubeWidth; cubeX += incrementSpeed {
			for cubeY := -cubeWidth; cubeY < cubeWidth; cubeY += incrementSpeed {
				calculateForSurface(cubeX, cubeY, -cubeWidth, colorRed, &zBuffer, &buffer)
				calculateForSurface(cubeWidth, cubeY, cubeX, colorGreen, &zBuffer, &buffer)
				calculateForSurface(-cubeWidth, cubeY, -cubeX, colorBlue, &zBuffer, &buffer)
				calculateForSurface(-cubeX, cubeY, cubeWidth, colorOrange, &zBuffer, &buffer)
				calculateForSurface(cubeX, -cubeWidth, -cubeY, colorWhite, &zBuffer, &buffer)
				calculateForSurface(cubeX, cubeWidth, cubeY, colorYellow, &zBuffer, &buffer)
			}
		}

		err := s.Fill(colorBlack)
        if err != nil {
            fmt.Println(err)
            return
        }

		for i, pixel := range buffer {
			x := (i % renderW) * scaleX
			y := (i / renderW) * scaleY
            if pixel != colorBlack {
                err := s.DrawRect(x, y, x + scaleX, y + scaleY, pixel)
                if err != nil {
                    fmt.Println(err)
                    return
                }
            }
		}

		a += 0.03
		b += 0.02
		c += 0.01

        oldTime := newTime
        newTime = time.Now()

        delta := newTime.Sub(oldTime).Seconds()
        fps := int(1.0 / delta)
        err = s.DrawText(10, 20, fmt.Sprintf("%v FPS", fps), colorWhite)
        if err != nil {
            fmt.Println(err)
            return
        }

		s.Redraw()
	}
}

