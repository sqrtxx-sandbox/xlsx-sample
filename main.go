package main

import (
	"fmt"
	"github.com/tealeg/xlsx"
	"os"
)

func main() {
	excelFileName := os.Args[1]
	xlFile, error := xlsx.OpenFile(excelFileName)
	if error != nil {
		fmt.Println(error)
	}
	// sheet
	st := xlFile.Sheet["ｼｰﾄｽﾘｰ"]
	// cell
	fmt.Println("-----", "cell", "-----")
	cl := st.Cell(0, 0)
	fmt.Println(cl.String())
	// all sheets
	sheets := xlFile.Sheets
	// all cells
	fmt.Println("-----", "all cells", "-----")
	for _, sheet := range sheets {
		fmt.Printf("\n---\nsheet name: %s\n---\n", sheet.Name)
		for _, row := range sheet.Rows {
			for _, cell := range row.Cells {
				fmt.Println(cell.String())
			}
		}
	}
}
