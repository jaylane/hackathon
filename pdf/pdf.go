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
	pdf.SetAutoPageBreak(true, 2)
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 26)
	pdf.CellFormat(200, 200, d.Title, "", 1, "CM", false, 0, "")
	pdf.SetFont("Arial", "B", 16)
	pdf.CellFormat(-200, -200, "by "+d.AuthorName, "", 2, "CM", false, 0, "")
	pdf.AddPage()
	pdf.Cell(20, 20, string(d.ChatGPTResponse.Choices[0].Message.Content))
	pdf.OutputFileAndClose(d.Filename + ".pdf")
}
