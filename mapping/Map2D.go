package mapping

import "log"

func Map2D() map[string]map[string]float32 {
	books := make(map[string]map[string]float32)
	books["chinese"] = make(map[string]float32)
	books["english"] = make(map[string]float32)
	books["malay"] = make(map[string]float32)

	books["chinese"]["Water Margin"] = 29.00
	books["chinese"]["Journey to the West"] = 21.00
	books["english"]["A Tale of Two Cities"] = 9.90
	books["malay"]["Aku Anak Timur"] = 13.50

	for language, book := range books {
		for title, price := range book {
			log.Printf("The book title %v is written in %v language. Its price is RM %v", title, language, price)
		}
	}

	return books
}
