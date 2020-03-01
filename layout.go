package gocui

type Rect struct {
    X0, Y0, X1, Y1 int
}

func (r Rect) W() int {
    return r.X1 - r.X0 + 1
}

func (r Rect) H() int {
    return r.Y1 - r.Y0 + 1
}

type CSS struct {
    X, Y           int
    W, H           int
    ML, MT, MR, MB int
}

func NewCSS(x, y, w, h int) CSS {
    return CSS{x, y, w, h, 0, 0, 0, 0}
}

type Layout interface {
    Add(css CSS) (x0, y0, x1, y1 int)
    //AddLayout(lay Setup) (x0, y0, x1, y1 int)
    Reset()
}

////////////////////////////////////////////////////////////////////////////////////////////////////

type FlowLayout struct {
    CSS
    CurMaxY      int
    CurX, CurY int
}

func NewFlowLayout(css CSS) *FlowLayout {
    l := &FlowLayout{
        CSS: css,
        CurX: css.X,
        CurY: css.Y,
        CurMaxY: css.Y,
    }
    return l
}

func (f *FlowLayout) Reset() {
    f.CurX = f.CSS.X
    f.CurY = f.CSS.Y
    f.CurMaxY = 1
}

func (f *FlowLayout) Add(css CSS) Rect {
    maxX := f.CurX + css.W + css.ML + css.MR
    layoutMaxX := f.CSS.X + f.CSS.W

    if maxX <= layoutMaxX  {
    } else {
        f.CurY = f.CurMaxY + 1
        f.CurX = f.CSS.X
    }

    x0, y0 := f.CurX + css.ML, f.CurY + css.MT
    x1, y1 := x0 + css.W - 1, y0 + css.H - 1
    f.CurX = x1 + css.MR + 1

    if y1 > f.CurMaxY {
        f.CurMaxY = y1
    }

    return Rect{x0, y0, x1, y1}
}
//
//func (f *FlowLayout) AddLayout(lay Setup) (x0, y0, x1, y1 int) {
//    x0, y0, x1, y1 = f.Add(lay.CSS)
//}