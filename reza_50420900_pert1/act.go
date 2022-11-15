package main

type Segitiga struct {
	Sisi   float64
	Alas   float64
	Tinggi float64
}

func (k Segitiga) Luas() float64 {
	return k.Alas * k.Tinggi * 0.5
}

func (k Segitiga) Keliling() float64 {
	return k.Sisi + k.Sisi + k.Sisi
}
