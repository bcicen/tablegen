package tablewriter

import (
	"github.com/bcicen/color"
)

func format(s string, c color.Color) string { return c.Sprint(s) }

// Adding header colors
func (t *Table) SetHeaderColor(colors ...color.Color) {
	if t.colSize != len(colors) {
		panic("Number of header colors must be equal to number of headers.")
	}
	for i := 0; i < len(colors); i++ {
		t.headerParams = append(t.headerParams, colors[i])
	}
}

// Adding column colors
func (t *Table) SetColumnColor(colors ...color.Color) {
	if t.colSize != len(colors) {
		panic("Number of column colors must be equal to number of headers.")
	}
	for i := 0; i < len(colors); i++ {
		t.columnsParams = append(t.columnsParams, colors[i])
	}
}

// Adding border colors
func (t *Table) SetBorderColor(color color.Color) {
	t.borderColor = color
}

// Adding column colors
func (t *Table) SetFooterColor(colors ...color.Color) {
	if len(t.footers) != len(colors) {
		panic("Number of footer colors must be equal to number of footer.")
	}
	for i := 0; i < len(colors); i++ {
		t.footerParams = append(t.footerParams, colors[i])
	}
}
