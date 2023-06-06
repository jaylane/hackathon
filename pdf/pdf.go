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

func BuildPDF(d *structs.Downloadable) {
	pdf := fpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 20)
	pdf.Cell(50, 30, d.Title)
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(50, 40, d.AuthorName)
	pdf.AddPage()
	pdf.Cell(20, 20, string(d.ChatGPTResponse.Choices[0].Message.Content))
	pdf.OutputFileAndClose(d.Filename + ".pdf")
}
