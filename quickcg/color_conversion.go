package quickcg

import (
	"math"
)

// RGBtoHSL converts an RGB color to the HSL color model.
func RGBtoHSL(rgb ColorRGB) ColorHSL {
	r := float64(rgb.R) / 255
	g := float64(rgb.G) / 255
	b := float64(rgb.B) / 255

	max := math.Max(r, math.Max(g, b))
	min := math.Min(r, math.Min(g, b))
	l := (max + min) / 2

	var h, s float64
	if max == min {
		h = 0
		s = 0
	} else {
		d := max - min
		s = d / (1 - math.Abs(2*l-1))
		switch max {
		case r:
			h = math.Mod((g-b)/d, 6)
		case g:
			h = (b - r) / d + 2
		case b:
			h = (r - g) / d + 4
		}
		h /= 6
		if h < 0 {
			h += 1
		}
	}
	return ColorHSL{H: h, S: s, L: l}
}

// HSLtoRGB converts an HSL color to the RGB color model.
func HSLtoRGB(hsl ColorHSL) ColorRGB {
	h, s, l := hsl.H, hsl.S, hsl.L
	var r, g, b float64
	if s == 0 {
		r, g, b = l, l, l
	} else {
		q := l * (1 + s)
		if l >= 0.5 {
			q = l + s - l*s
		}
		p := 2*l - q
		r = hueToRGB(p, q, h+1.0/3.0)
		g = hueToRGB(p, q, h)
		b = hueToRGB(p, q, h-1.0/3.0)
	}
	return ColorRGB{R: uint8(r * 255), G: uint8(g * 255), B: uint8(b * 255)}
}

func hueToRGB(p, q, t float64) float64 {
	if t < 0 {
		t += 1
	}
	if t > 1 {
		t -= 1
	}
	if t < 1.0/6.0 {
		return p + (q-p)*6*t
	}
	if t < 1.0/2.0 {
		return q
	}
	if t < 2.0/3.0 {
		return p + (q-p)*(2.0/3.0 - t)*6
	}
	return p
}

// RGBtoHSV converts an RGB color to the HSV color model.
func RGBtoHSV(rgb ColorRGB) ColorHSV {
	r := float64(rgb.R) / 255
	g := float64(rgb.G) / 255
	b := float64(rgb.B) / 255

	max := math.Max(r, math.Max(g, b))
	min := math.Min(r, math.Min(g, b))
	d := max - min

	v := max
	var h, s float64
	if max != 0 {
		s = d / max
	} else {
		s = 0
		h = 0
		return ColorHSV{H: h, S: s, V: v}
	}

	switch max {
	case r:
		h = (g - b) / d
		if g < b {
			h += 6
		}
	case g:
		h = (b - r) / d + 2
	case b:
		h = (r - g) / d + 4
	}
	h /= 6
	return ColorHSV{H: h, S: s, V: v}
}

// HSVtoRGB converts an HSV color to the RGB color model.
func HSVtoRGB(hsv ColorHSV) ColorRGB {
	h := hsv.H * 6
	s := hsv.S
	v := hsv.V

	i := int(math.Floor(h))
	f := h - float64(i)
	p := v * (1 - s)
	q := v * (1 - f*s)
	t := v * (1 - (1-f)*s)

	var r, g, b float64
	switch i % 6 {
	case 0:
		r, g, b = v, t, p
	case 1:
		r, g, b = q, v, p
	case 2:
		r, g, b = p, v, t
	case 3:
		r, g, b = p, q, v
	case 4:
		r, g, b = t, p, v
	case 5:
		r, g, b = v, p, q
	}

	return ColorRGB{R: uint8(r * 255), G: uint8(g * 255), B: uint8(b * 255)}
}

