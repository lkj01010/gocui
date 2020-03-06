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

type Css struct {
    X, Y           int
    W, H           int
    ML, MT, MR, MB int
}

func NewCss(x, y, w, h int) Css {
    return Css{x, y, w, h, 0, 0, 0, 0}
}

type Layout interface {
    Add(css Css) Rect
    //AddLayout(lay Setup) (x0, y0, x1, y1 int)
    Reset()
}

////////////////////////////////////////////////////////////////////////////////////////////////////

type FlowLayout struct {
    Css
    CurMaxY      int
    CurX, CurY int
}

func NewFlowLayout(css Css) *FlowLayout {
    l := &FlowLayout{
        Css:     css,
        CurX:    css.X,
        CurY:    css.Y,
        CurMaxY: css.Y,
    }
    return l
}

func (f *FlowLayout) Reset() {
    f.CurX = f.Css.X
    f.CurY = f.Css.Y
    f.CurMaxY = 1
}

func (f *FlowLayout) Add(css Css) Rect {
    maxX := f.CurX + css.W + css.ML + css.MR
    layoutMaxX := f.Css.X + f.Css.W

    if maxX <= layoutMaxX  {
    } else {
        f.CurY = f.CurMaxY + 1
        f.CurX = f.Css.X
    }

    x0, y0 := f.CurX + css.ML, f.CurY + css.MT
    x1, y1 := x0 + css.W - 1, y0 + css.H - 1
    f.CurX = x1 + css.MR + 1

    if y1 > f.CurMaxY {
        f.CurMaxY = y1
    }

    return Rect{x0, y0, x1, y1}
}

func (f *FlowLayout) NextLine(interval int) {
    f.CurMaxY += 1 + interval
    f.CurX = f.Css.X
    f.CurY = f.CurMaxY
}
//
//func (f *FlowLayout) AddLayout(lay Setup) (x0, y0, x1, y1 int) {
//    x0, y0, x1, y1 = f.Add(lay.Css)
//}