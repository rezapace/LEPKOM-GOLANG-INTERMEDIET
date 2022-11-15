package main

import "testing"

var (
	segitiga           Segitiga = Segitiga{5, 4, 2}
	luasSeharusnya     float64  = 4
	kelilingSeharusnya float64  = 15
)

func TestHitungLuas(t *testing.T) {
	t.Logf("Luas : %.2f", segitiga.Luas())
	if segitiga.Luas() != luasSeharusnya {
		t.Errorf("SALAH!!! HARUSNYA %.2f", luasSeharusnya)
	}
}

func TestHitungKeliling(t *testing.T) {
	t.Logf("Luas : %.2f", segitiga.Keliling())
	if segitiga.Keliling() != kelilingSeharusnya {
		t.Errorf("SALAH!!! HARUSNYA %.2f", kelilingSeharusnya)
	}
}
