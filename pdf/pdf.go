package pdf

import (
	"github.com/go-pdf/fpdf"
	"jay-lane.com/hackathon/structs"
)

func TestPage() {
	pdf := fpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetAutoPageBreak(true, 2)
	pdf.SetMargins(10, 20, 10)
	pdf.SetFont("Arial", "B", 26)
	pdf.Text(75, 100, "TITLE OF DOC")
	// pdf.CellFormat(200, 100, "Title of document", "", 1, "CM", false, 0, "")
	pdf.SetFont("Arial", "B", 16)
	// pdf.CellFormat(200, 110, "by Jay Lane", "", 1, "CM", false, 0, "")
	pdf.Text(75, 110, "by Jay Lane")
	pdf.OutputFileAndClose("hello.pdf")
}

func BuildPDF(d *structs.Downloadable) {
	pdf := fpdf.New("P", "mm", "A4", "")
	pdf.SetAutoPageBreak(true, 2)
	pdf.SetMargins(10, 20, 10)
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 26)
	pdf.CellFormat(200, 200, d.Title, "", 1, "CM", false, 0, "")
	pdf.SetFont("Arial", "B", 16)
	pdf.CellFormat(-200, -200, "by "+d.AuthorName, "", 2, "CM", false, 0, "")
	pdf.AddPage()
	pdf.Cell(20, 20, string(d.ChatGPTResponse.Choices[0].Message.Content))
	pdf.OutputFileAndClose(d.Filename + ".pdf")
}
