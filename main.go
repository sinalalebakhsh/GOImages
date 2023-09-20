package main

func main() {
	
	for _, p := range Products {
		Printfln("Product: %v, Category: %v, Price: $%.2f",
			p.Name, p.Category, p.Price)
	}
	
}

