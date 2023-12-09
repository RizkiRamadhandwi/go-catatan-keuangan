package main

import "enigmacamp.com/livecode-catatan-keuangan/delivery"

func main() {
	delivery.NewServer().Run()
	println("Livecode Catatan Keuangan")
}
