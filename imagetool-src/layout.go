package main

// Cell 拼图模板中的一个格子
type Cell struct {
	Row, Col    int // 起始行/列
	RowSpan     int // 占行数
	ColSpan     int // 占列数
	HasImage    bool // 是否放置图片（false 表示留空，显示背景色）
}

// Layout 一个拼图布局
type Layout struct {
	ID    string
	Name  string
	Rows  int
	Cols  int
	Cells []Cell
}

// 拼图模式模板（拼图模式）
var collageTemplates = map[string][]Layout{
	"classic9": { // 经典九宫格
		{ID: "classic9-a", Name: "经典九宫格 3×3", Rows: 3, Cols: 3, Cells: gridCells(3, 3, true)},
	},
	"2": {
		{ID: "2-a", Name: "左右各一", Rows: 1, Cols: 2, Cells: gridCells(1, 2, true)},
		{ID: "2-b", Name: "上下各一", Rows: 2, Cols: 1, Cells: gridCells(2, 1, true)},
		{ID: "2-c", Name: "左大右小", Rows: 2, Cols: 2, Cells: []Cell{
			{Row: 0, Col: 0, RowSpan: 2, ColSpan: 1, HasImage: true},
			{Row: 0, Col: 1, RowSpan: 1, ColSpan: 1, HasImage: true},
			{Row: 1, Col: 1, RowSpan: 1, ColSpan: 1, HasImage: false},
		}},
		{ID: "2-d", Name: "上大下小", Rows: 2, Cols: 2, Cells: []Cell{
			{Row: 0, Col: 0, RowSpan: 1, ColSpan: 2, HasImage: true},
			{Row: 1, Col: 0, RowSpan: 1, ColSpan: 1, HasImage: true},
			{Row: 1, Col: 1, RowSpan: 1, ColSpan: 1, HasImage: false},
		}},
	},
	"3": {
		{ID: "3-a", Name: "横排三图", Rows: 1, Cols: 3, Cells: gridCells(1, 3, true)},
		{ID: "3-b", Name: "竖排三图", Rows: 3, Cols: 1, Cells: gridCells(3, 1, true)},
		{ID: "3-c", Name: "上1下2", Rows: 2, Cols: 2, Cells: []Cell{
			{Row: 0, Col: 0, RowSpan: 1, ColSpan: 2, HasImage: true},
			{Row: 1, Col: 0, RowSpan: 1, ColSpan: 1, HasImage: true},
			{Row: 1, Col: 1, RowSpan: 1, ColSpan: 1, HasImage: true},
		}},
		{ID: "3-d", Name: "上2下1", Rows: 2, Cols: 2, Cells: []Cell{
			{Row: 0, Col: 0, RowSpan: 1, ColSpan: 1, HasImage: true},
			{Row: 0, Col: 1, RowSpan: 1, ColSpan: 1, HasImage: true},
			{Row: 1, Col: 0, RowSpan: 1, ColSpan: 2, HasImage: true},
		}},
		{ID: "3-e", Name: "左1右2", Rows: 2, Cols: 2, Cells: []Cell{
			{Row: 0, Col: 0, RowSpan: 2, ColSpan: 1, HasImage: true},
			{Row: 0, Col: 1, RowSpan: 1, ColSpan: 1, HasImage: true},
			{Row: 1, Col: 1, RowSpan: 1, ColSpan: 1, HasImage: true},
		}},
		{ID: "3-f", Name: "左2右1", Rows: 2, Cols: 2, Cells: []Cell{
			{Row: 0, Col: 0, RowSpan: 1, ColSpan: 1, HasImage: true},
			{Row: 1, Col: 0, RowSpan: 1, ColSpan: 1, HasImage: true},
			{Row: 0, Col: 1, RowSpan: 2, ColSpan: 1, HasImage: true},
		}},
	},
	"4": {
		{ID: "4-a", Name: "田字格 2×2", Rows: 2, Cols: 2, Cells: gridCells(2, 2, true)},
		{ID: "4-b", Name: "横排四图", Rows: 1, Cols: 4, Cells: gridCells(1, 4, true)},
		{ID: "4-c", Name: "竖排四图", Rows: 4, Cols: 1, Cells: gridCells(4, 1, true)},
		{ID: "4-d", Name: "上1下3", Rows: 2, Cols: 3, Cells: []Cell{
			{Row: 0, Col: 0, RowSpan: 1, ColSpan: 3, HasImage: true},
			{Row: 1, Col: 0, RowSpan: 1, ColSpan: 1, HasImage: true},
			{Row: 1, Col: 1, RowSpan: 1, ColSpan: 1, HasImage: true},
			{Row: 1, Col: 2, RowSpan: 1, ColSpan: 1, HasImage: true},
		}},
		{ID: "4-e", Name: "左1右3", Rows: 3, Cols: 2, Cells: []Cell{
			{Row: 0, Col: 0, RowSpan: 3, ColSpan: 1, HasImage: true},
			{Row: 0, Col: 1, RowSpan: 1, ColSpan: 1, HasImage: true},
			{Row: 1, Col: 1, RowSpan: 1, ColSpan: 1, HasImage: true},
			{Row: 2, Col: 1, RowSpan: 1, ColSpan: 1, HasImage: true},
		}},
	},
	"5": {
		{ID: "5-a", Name: "上2下3", Rows: 2, Cols: 3, Cells: []Cell{
			{Row: 0, Col: 0, RowSpan: 1, ColSpan: 1, HasImage: true},
			{Row: 0, Col: 1, RowSpan: 1, ColSpan: 1, HasImage: true},
			{Row: 1, Col: 0, RowSpan: 1, ColSpan: 1, HasImage: true},
			{Row: 1, Col: 1, RowSpan: 1, ColSpan: 1, HasImage: true},
			{Row: 1, Col: 2, RowSpan: 1, ColSpan: 1, HasImage: true},
		}},
		{ID: "5-b", Name: "上3下2", Rows: 2, Cols: 3, Cells: []Cell{
			{Row: 0, Col: 0, RowSpan: 1, ColSpan: 1, HasImage: true},
			{Row: 0, Col: 1, RowSpan: 1, ColSpan: 1, HasImage: true},
			{Row: 0, Col: 2, RowSpan: 1, ColSpan: 1, HasImage: true},
			{Row: 1, Col: 0, RowSpan: 1, ColSpan: 1, HasImage: true},
			{Row: 1, Col: 1, RowSpan: 1, ColSpan: 1, HasImage: true},
		}},
		{ID: "5-c", Name: "上1下4", Rows: 2, Cols: 4, Cells: []Cell{
			{Row: 0, Col: 0, RowSpan: 1, ColSpan: 4, HasImage: true},
			{Row: 1, Col: 0, RowSpan: 1, ColSpan: 1, HasImage: true},
			{Row: 1, Col: 1, RowSpan: 1, ColSpan: 1, HasImage: true},
			{Row: 1, Col: 2, RowSpan: 1, ColSpan: 1, HasImage: true},
			{Row: 1, Col: 3, RowSpan: 1, ColSpan: 1, HasImage: true},
		}},
		{ID: "5-d", Name: "2-1-2 三段", Rows: 3, Cols: 2, Cells: []Cell{
			{Row: 0, Col: 0, RowSpan: 1, ColSpan: 1, HasImage: true},
			{Row: 0, Col: 1, RowSpan: 1, ColSpan: 1, HasImage: true},
			{Row: 1, Col: 0, RowSpan: 1, ColSpan: 2, HasImage: true},
			{Row: 2, Col: 0, RowSpan: 1, ColSpan: 1, HasImage: true},
			{Row: 2, Col: 1, RowSpan: 1, ColSpan: 1, HasImage: true},
		}},
	},
	"6": {
		{ID: "6-a", Name: "2行3列", Rows: 2, Cols: 3, Cells: gridCells(2, 3, true)},
		{ID: "6-b", Name: "3行2列", Rows: 3, Cols: 2, Cells: gridCells(3, 2, true)},
		{ID: "6-c", Name: "2-2-2", Rows: 3, Cols: 2, Cells: gridCells(3, 2, true)},
		{ID: "6-d", Name: "上3中1下2", Rows: 3, Cols: 3, Cells: []Cell{
			{Row: 0, Col: 0, RowSpan: 1, ColSpan: 1, HasImage: true},
			{Row: 0, Col: 1, RowSpan: 1, ColSpan: 1, HasImage: true},
			{Row: 0, Col: 2, RowSpan: 1, ColSpan: 1, HasImage: true},
			{Row: 1, Col: 0, RowSpan: 1, ColSpan: 3, HasImage: true},
			{Row: 2, Col: 0, RowSpan: 1, ColSpan: 1, HasImage: true},
			{Row: 2, Col: 1, RowSpan: 1, ColSpan: 1, HasImage: true},
			{Row: 2, Col: 2, RowSpan: 1, ColSpan: 1, HasImage: false},
		}},
	},
	"7": {
		{ID: "7-a", Name: "2-3-2", Rows: 3, Cols: 3, Cells: []Cell{
			{Row: 0, Col: 0, RowSpan: 1, ColSpan: 1, HasImage: true},
			{Row: 0, Col: 1, RowSpan: 1, ColSpan: 1, HasImage: true},
			{Row: 1, Col: 0, RowSpan: 1, ColSpan: 1, HasImage: true},
			{Row: 1, Col: 1, RowSpan: 1, ColSpan: 1, HasImage: true},
			{Row: 1, Col: 2, RowSpan: 1, ColSpan: 1, HasImage: true},
			{Row: 2, Col: 0, RowSpan: 1, ColSpan: 1, HasImage: true},
			{Row: 2, Col: 1, RowSpan: 1, ColSpan: 1, HasImage: true},
		}},
		{ID: "7-b", Name: "上1中3下3", Rows: 3, Cols: 3, Cells: []Cell{
			{Row: 0, Col: 0, RowSpan: 1, ColSpan: 3, HasImage: true},
			{Row: 1, Col: 0, RowSpan: 1, ColSpan: 1, HasImage: true},
			{Row: 1, Col: 1, RowSpan: 1, ColSpan: 1, HasImage: true},
			{Row: 1, Col: 2, RowSpan: 1, ColSpan: 1, HasImage: true},
			{Row: 2, Col: 0, RowSpan: 1, ColSpan: 1, HasImage: true},
			{Row: 2, Col: 1, RowSpan: 1, ColSpan: 1, HasImage: true},
			{Row: 2, Col: 2, RowSpan: 1, ColSpan: 1, HasImage: true},
		}},
		{ID: "7-c", Name: "3-1-3", Rows: 3, Cols: 3, Cells: []Cell{
			{Row: 0, Col: 0, RowSpan: 1, ColSpan: 1, HasImage: true},
			{Row: 0, Col: 1, RowSpan: 1, ColSpan: 1, HasImage: true},
			{Row: 0, Col: 2, RowSpan: 1, ColSpan: 1, HasImage: true},
			{Row: 1, Col: 0, RowSpan: 1, ColSpan: 3, HasImage: true},
			{Row: 2, Col: 0, RowSpan: 1, ColSpan: 1, HasImage: true},
			{Row: 2, Col: 1, RowSpan: 1, ColSpan: 1, HasImage: true},
			{Row: 2, Col: 2, RowSpan: 1, ColSpan: 1, HasImage: true},
		}},
	},
	"8": {
		{ID: "8-a", Name: "2行4列", Rows: 2, Cols: 4, Cells: gridCells(2, 4, true)},
		{ID: "8-b", Name: "4行2列", Rows: 4, Cols: 2, Cells: gridCells(4, 2, true)},
		{ID: "8-c", Name: "3-2-3", Rows: 3, Cols: 3, Cells: []Cell{
			{Row: 0, Col: 0, RowSpan: 1, ColSpan: 1, HasImage: true},
			{Row: 0, Col: 1, RowSpan: 1, ColSpan: 1, HasImage: true},
			{Row: 0, Col: 2, RowSpan: 1, ColSpan: 1, HasImage: true},
			{Row: 1, Col: 0, RowSpan: 1, ColSpan: 1, HasImage: true},
			{Row: 1, Col: 1, RowSpan: 1, ColSpan: 1, HasImage: true},
			{Row: 2, Col: 0, RowSpan: 1, ColSpan: 1, HasImage: true},
			{Row: 2, Col: 1, RowSpan: 1, ColSpan: 1, HasImage: true},
			{Row: 2, Col: 2, RowSpan: 1, ColSpan: 1, HasImage: true},
		}},
		{ID: "8-d", Name: "2-4-2", Rows: 3, Cols: 4, Cells: []Cell{
			{Row: 0, Col: 0, RowSpan: 1, ColSpan: 1, HasImage: true},
			{Row: 0, Col: 1, RowSpan: 1, ColSpan: 1, HasImage: true},
			{Row: 1, Col: 0, RowSpan: 1, ColSpan: 1, HasImage: true},
			{Row: 1, Col: 1, RowSpan: 1, ColSpan: 1, HasImage: true},
			{Row: 1, Col: 2, RowSpan: 1, ColSpan: 1, HasImage: true},
			{Row: 1, Col: 3, RowSpan: 1, ColSpan: 1, HasImage: true},
			{Row: 2, Col: 2, RowSpan: 1, ColSpan: 1, HasImage: true},
			{Row: 2, Col: 3, RowSpan: 1, ColSpan: 1, HasImage: true},
		}},
	},
	"9": {
		{ID: "9-a", Name: "三行三列", Rows: 3, Cols: 3, Cells: gridCells(3, 3, true)},
		{ID: "9-b", Name: "大图置顶", Rows: 3, Cols: 3, Cells: []Cell{
			{Row: 0, Col: 0, RowSpan: 1, ColSpan: 2, HasImage: true},
			{Row: 0, Col: 2, RowSpan: 1, ColSpan: 1, HasImage: true},
			{Row: 1, Col: 0, RowSpan: 1, ColSpan: 1, HasImage: true},
			{Row: 1, Col: 1, RowSpan: 1, ColSpan: 1, HasImage: true},
			{Row: 1, Col: 2, RowSpan: 1, ColSpan: 1, HasImage: true},
			{Row: 2, Col: 0, RowSpan: 1, ColSpan: 1, HasImage: true},
			{Row: 2, Col: 1, RowSpan: 1, ColSpan: 1, HasImage: true},
			{Row: 2, Col: 2, RowSpan: 1, ColSpan: 1, HasImage: true},
		}},
		{ID: "9-c", Name: "2-3-4", Rows: 3, Cols: 4, Cells: []Cell{
			{Row: 0, Col: 0, RowSpan: 1, ColSpan: 2, HasImage: true},
			{Row: 0, Col: 2, RowSpan: 1, ColSpan: 2, HasImage: true},
			{Row: 1, Col: 0, RowSpan: 1, ColSpan: 1, HasImage: true},
			{Row: 1, Col: 1, RowSpan: 1, ColSpan: 1, HasImage: true},
			{Row: 1, Col: 2, RowSpan: 1, ColSpan: 2, HasImage: true},
			{Row: 2, Col: 0, RowSpan: 1, ColSpan: 1, HasImage: true},
			{Row: 2, Col: 1, RowSpan: 1, ColSpan: 1, HasImage: true},
			{Row: 2, Col: 2, RowSpan: 1, ColSpan: 1, HasImage: true},
			{Row: 2, Col: 3, RowSpan: 1, ColSpan: 1, HasImage: true},
		}},
	},
}

// gridCells 生成 uniform 网格布局
func gridCells(rows, cols int, allImage bool) []Cell {
	cells := make([]Cell, 0, rows*cols)
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			cells = append(cells, Cell{Row: r, Col: c, RowSpan: 1, ColSpan: 1, HasImage: true})
		}
	}
	return cells
}

// 分图模板（分图模式）：rows × cols
type SplitTemplate struct {
	ID   string
	Name string
	Rows int
	Cols int
}

var splitTemplates = []SplitTemplate{
	{ID: "split-9", Name: "1拆9 (3×3)", Rows: 3, Cols: 3},
	{ID: "split-6", Name: "1拆6 (2×3)", Rows: 2, Cols: 3},
	{ID: "split-4", Name: "1拆4 (2×2)", Rows: 2, Cols: 2},
}

// findCollageLayout 查找拼图布局
func findCollageLayout(mode, variant string) *Layout {
	layouts, ok := collageTemplates[mode]
	if !ok {
		return nil
	}
	if variant == "" {
		return &layouts[0]
	}
	for i := range layouts {
		if layouts[i].ID == variant {
			return &layouts[i]
		}
	}
	return &layouts[0]
}

// layoutImageCount 布局需要的图片数量
func layoutImageCount(l *Layout) int {
	n := 0
	for _, c := range l.Cells {
		if c.HasImage {
			n++
		}
	}
	return n
}
