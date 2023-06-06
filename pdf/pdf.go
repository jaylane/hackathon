package pdf

import (
	"github.com/go-pdf/fpdf"
	"jay-lane.com/hackathon/structs"
)

func TestPage() {
	pdf := fpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(40, 10, "Hello, world")
	pdf.OutputFileAndClose("hello.pdf")
}

func BuildPDF(downloadable structs.Downloadable) {
	return
}
