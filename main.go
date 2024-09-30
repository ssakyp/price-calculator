package main

import (
	"fmt"

	"github.com/ssakyp/price-calculator/filemanager"
	"github.com/ssakyp/price-calculator/prices"
)

func main() {

	taxRates := []float64{0, 0.07, 0.1, 0.15}
	doneChans := make([]chan bool, len(taxRates))

	for i, taxRate := range taxRates {
		doneChans[i] = make(chan bool)
		fm := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
		//cmdm := cmdmanager.New()
		priceJob := prices.NewTaxIncludedPriceJob(*fm, taxRate)
		go priceJob.Process(doneChans[i])

		// if err != nil {
		// 	fmt.Println("Could not process job")
		// 	fmt.Println(err)
		// }
	}

	for _, doneChan := range doneChans {
		<- doneChan
	}

}
