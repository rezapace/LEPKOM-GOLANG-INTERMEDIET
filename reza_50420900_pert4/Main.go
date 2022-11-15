package main

//Muhammad reza hidayat
import (
	"log"
	"net/http"
	"reza_50420900_pert4/handler" // Ganti dengan nama folder kalian masingmasing
)

func main() {
	http.HandleFunc("/api/", handler.API)
	log.Println("localhost : 8050")   //Ganti 2 digit akhir port dengan 2 digit akhir NPM anda
	http.ListenAndServe(":8050", nil) //Ganti 2 digit akhir port dengan 2 digit akhir NPM anda
}

// Windows PowerShell
// Copyright (C) Microsoft Corporation. All rights reserved.

// Try the new cross-platform PowerShell https://aka.ms/pscore6

// PS C:\Program Files\Go\src\reza_50420900_pert4> go run main.go
// 2022/11/04 03:56:47 localhost : 8050
