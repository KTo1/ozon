package main

import (
	"github.com/jung-kurt/gofpdf"
	"io/ioutil"
	"log"
	"time"
)

func trace(s string) (string, time.Time) {
	log.Println("START:", s)
	return s, time.Now()
}

func un(s string, startTime time.Time) {
	endTime := time.Now()
	log.Println("END:", s, "ElapsedTime in seconds:", endTime.Sub(startTime))
}

func gopdf() {
	defer un(trace("gopdf"))

	file := "in.txt"
	content, err := ioutil.ReadFile(file)

	if err != nil {
		log.Fatalf("reading file error: %v", err)
	}

	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 12)
	pdf.MultiCell(190, 5, string(content), "0", "0", false)

	err = pdf.OutputFileAndClose("out.pdf")
	if err != nil {
		log.Fatalf("writing file error: %v", err)
	}
}

func main() {
	gopdf()
}
