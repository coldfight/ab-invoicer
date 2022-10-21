package main

import (
	"fmt"
	"github.com/johnfercher/maroto/pkg/color"
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/pdf"
	"github.com/johnfercher/maroto/pkg/props"
	"os"
	"time"
)

var (
	darkGrayColor color.Color = getDarkGrayColor()
	grayColor     color.Color = getGrayColor()
	whiteColor    color.Color = color.NewWhite()
	blueColor     color.Color = getBlueColor()
	redColor      color.Color = getRedColor()
)

func renderContactInfo(m pdf.Maroto) {
	m.Text("Jane Doe", props.Text{
		Size:        8,
		Align:       consts.Left,
		Extrapolate: false,
		Color:       redColor,
	})
	m.Text("555 John Street", props.Text{
		Top:         4,
		Size:        8,
		Align:       consts.Left,
		Extrapolate: false,
		Color:       redColor,
	})
	m.Text("Mississauga, ON, M4T2T1", props.Text{
		Top:         7.5,
		Size:        8,
		Align:       consts.Left,
		Extrapolate: false,
		Color:       redColor,
	})
	m.Text("Phone: 416.555.0555", props.Text{
		Top:   13,
		Style: consts.BoldItalic,
		Size:  8,
		Align: consts.Left,
		Color: blueColor,
	})
	m.Text("Email: jane.doe@gmail.com", props.Text{
		Top:   16,
		Style: consts.BoldItalic,
		Size:  8,
		Align: consts.Left,
		Color: blueColor,
	})
}

func renderInvoiceDateAndNumber(m pdf.Maroto) {

	m.Text("I N V O I C E", props.Text{
		Style: consts.Bold,
		Size:  26,
		Align: consts.Right,
	})

	var fromTop float64 = 12

	m.Text("INVOICE # 001 | DATE: 01/01/2022", props.Text{
		Top:   fromTop,
		Size:  8,
		Align: consts.Right,
		Style: consts.Bold,
	})
}

func renderHeader(m pdf.Maroto) {
	m.RegisterHeader(func() {
		m.Row(22, func() {
			m.Col(4, func() {
				renderContactInfo(m)
			})
			m.ColSpace(3)
			m.Col(5, func() {
				renderInvoiceDateAndNumber(m)
			})
		})
		m.Line(1.0, props.Line{
			Color: color.Color{
				Red:   0,
				Green: 0,
				Blue:  0,
			},
			Style: consts.Solid,
			Width: 0.5,
		})
		m.Row(3, func() {})
	})

}

func renderFooter(m pdf.Maroto) {
	m.RegisterFooter(func() {
		m.Row(20, func() {
			m.Col(12, func() {
				m.Text("Phone: 416.555.0555", props.Text{
					Top:   13,
					Style: consts.BoldItalic,
					Size:  8,
					Align: consts.Left,
					Color: blueColor,
				})
				m.Text("Email: jane.doe@gmail.com", props.Text{
					Top:   16,
					Style: consts.BoldItalic,
					Size:  8,
					Align: consts.Left,
					Color: blueColor,
				})
			})
		})
	})
}

func renderBilledTo(m pdf.Maroto) {
	m.Row(25, func() {
		m.Col(2, func() {
			m.Text("Billed To:", props.Text{
				Top:   8,
				Style: consts.Bold,
			})
		})
		m.Col(10, func() {
			m.Text("Some person out there", props.Text{
				Top:  1,
				Size: 9,
			})
			m.Text("5555 Somekindof Rd", props.Text{
				Top:  6,
				Size: 9,
			})

			m.Text("Mississauga, ON, M4T2T1", props.Text{
				Top:  11,
				Size: 9,
			})

			m.Text("905.555.0555", props.Text{
				Top:  16,
				Size: 9,
			})

		})
	})
}

func renderExpenses(m pdf.Maroto) {
	header := getExpensesHeader()
	contents := getExpensesContents()

	m.SetBackgroundColor(darkGrayColor)

	m.Row(7, func() {
		m.Col(3, func() {
			m.Text("Expenses", props.Text{
				Top:   1.5,
				Size:  9,
				Style: consts.Bold,
				Align: consts.Left,
				Color: color.NewWhite(),
			})
		})
		m.ColSpace(9)
	})

	m.SetBackgroundColor(whiteColor)

	m.TableList(header, contents, props.TableList{
		HeaderProp: props.TableListContent{
			Size:      9,
			GridSizes: []uint{2, 6, 2, 2},
		},
		ContentProp: props.TableListContent{
			Size:      8,
			GridSizes: []uint{2, 6, 2, 2},
		},
		Align:                consts.Left,
		AlternatedBackground: &grayColor,
		HeaderContentSpace:   1,
		Line:                 false,
	})

	m.Row(20, func() {
		m.ColSpace(7)
		m.Col(2, func() {
			m.Text("Total Expenses:", props.Text{
				Top:   2,
				Style: consts.Bold,
				Size:  8,
				Align: consts.Right,
			})
		})
		m.Col(3, func() {
			m.Text("R$ 2.567,00", props.Text{
				Top:   2,
				Style: consts.Bold,
				Size:  8,
				Align: consts.Center,
			})
		})
	})

}

func renderLabour(m pdf.Maroto) {
	header := getLabourHeader()
	contents := getLabourContents()

	m.SetBackgroundColor(darkGrayColor)

	m.Row(7, func() {
		m.Col(3, func() {
			m.Text("Labour", props.Text{
				Top:   1.5,
				Size:  9,
				Style: consts.Bold,
				Align: consts.Center,
				Color: color.NewWhite(),
			})
		})
		m.ColSpace(9)
	})

	m.SetBackgroundColor(whiteColor)

	m.TableList(header, contents, props.TableList{
		HeaderProp: props.TableListContent{
			Size:      9,
			GridSizes: []uint{2, 8, 2},
		},
		ContentProp: props.TableListContent{
			Size:      8,
			GridSizes: []uint{2, 8, 2},
		},
		Align:                consts.Left,
		AlternatedBackground: &grayColor,
		HeaderContentSpace:   1,
		Line:                 false,
	})

	m.Row(20, func() {
		m.ColSpace(7)
		m.Col(2, func() {
			m.Text("Total Labour:", props.Text{
				Top:   2,
				Style: consts.Bold,
				Size:  8,
				Align: consts.Right,
			})
		})
		m.Col(3, func() {
			m.Text("R$ 2.567,00", props.Text{
				Top:   2,
				Style: consts.Bold,
				Size:  8,
				Align: consts.Center,
			})
		})
	})

}

func main() {
	begin := time.Now()

	m := pdf.NewMaroto(consts.Portrait, consts.A4)
	m.SetPageMargins(10, 15, 10)
	//m.SetBorder(true)

	renderHeader(m)
	renderFooter(m)
	renderBilledTo(m)

	renderExpenses(m)
	renderLabour(m)

	err := m.OutputFileAndClose("pdf/invoice001.pdf")
	if err != nil {
		fmt.Println("Could not save PDF:", err)
		os.Exit(1)
	}

	end := time.Now()
	fmt.Println(end.Sub(begin))
}

func getExpensesHeader() []string {
	return []string{"Quantity", "Description", "Unit Price", "Amount"}
}

func getExpensesContents() [][]string {
	return [][]string{
		{"1", "Swamp", "$12.00", "$12.00"},
		{"1", "Sorin, A Planeswalker", "$12.00", "$12.00"},
		{"1", "Tassa", "$12.00", "$12.00"},
		{"1", "Skinrender", "$12.00", "$12.00"},
		{"1", "Island", "$12.00", "$12.00"},
		{"1", "Mountain", "$12.00", "$12.00"},
		{"1", "Plain", "$12.00", "$12.00"},
	}
}

func getLabourHeader() []string {
	return []string{"Date", "Description", "Amount"}
}

func getLabourContents() [][]string {
	return [][]string{
		{"Jan 1, 2022", "Swamp", "$12.00"},
		{"Jan 1, 2022", "Sorin, A Planeswalker", "$12.00"},
		{"Jan 1, 2022", "Tassa", "$12.00"},
		{"Jan 1, 2022", "Skinrender", "$12.00"},
		{"Jan 1, 2022", "Island", "$12.00"},
		{"Jan 1, 2022", "Mountain", "$12.00"},
		{"Jan 1, 2022", "Plain", "$12.00"},
	}
}

func getDarkGrayColor() color.Color {
	return color.Color{
		Red:   55,
		Green: 55,
		Blue:  55,
	}
}

func getGrayColor() color.Color {
	return color.Color{
		Red:   200,
		Green: 200,
		Blue:  200,
	}
}

func getBlueColor() color.Color {
	return color.Color{
		Red:   10,
		Green: 10,
		Blue:  150,
	}
}

func getRedColor() color.Color {
	return color.Color{
		Red:   150,
		Green: 10,
		Blue:  10,
	}
}
