package termui

import (
	gtermui "github.com/gizak/termui"
)

//Row is a widget for table rows
type Row struct {
	X, Y    int
	Width   int
	Height  int
	Columns []gtermui.GridBufferer
	Table   Table
}

//AddColumn adds the given column to this row
func (row *Row) AddColumn(c gtermui.GridBufferer) {
	row.Columns = append(row.Columns, c)
}

//GetHeight returns this Row heigth
func (row *Row) GetHeight() int {
	return row.Height
}

//SetX sets the x position of this Row
func (row *Row) SetX(x int) {
	row.X = x
}

//SetY sets the y position of this Row
func (row *Row) SetY(y int) {
	if y == row.Y {
		return
	}
	for _, col := range row.Columns {
		col.SetY(y)
	}
	row.Y = y
}

//SetWidth sets the width of this Row
func (row *Row) SetWidth(width int) {
	x := row.X
	//Setting the width of a row does very little, since
	//column widths come from the table
	//It might be worthy to add checks in the case of a width
	//being set that is different than the sum of row.Table.ColumnWidths()
	if row.Table != nil {
		for i, width := range row.Table.ColumnWidths() {
			col := row.Columns[i]
			col.SetX(x)
			col.SetWidth(width)
			x += width + DefaultColumnSpacing
		}
	} else {
		for _, col := range row.Columns {
			col.SetX(x)
			col.SetWidth(width)
			x += width + DefaultColumnSpacing
		}
	}
	row.Width = width
}

//Buffer returns this Row data as a gtermui.Buffer
func (row *Row) Buffer() gtermui.Buffer {
	buf := gtermui.NewBuffer()
	for _, col := range row.Columns {
		buf.Merge(col.Buffer())
	}
	return buf
}
