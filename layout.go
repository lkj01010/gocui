package gocui

type CSS struct {
    X, Y           int
    W, H           int
    ML, MT, MR, MB int
}

func NewCSS(x, y, w, h int) CSS {
    return CSS{x, y, w, h, 0, 0, 0, 0}
}

type Layouter interface {
    Add(css CSS) (x0, y0, x1, y1 int)
    Reset()
}

type FlowLayout struct {
    CSS
    CurMaxY      int
    CurX, CurY int
}

func NewFlowLayout(g *Gui, css CSS) *FlowLayout {
    l := &FlowLayout{
        CSS: css,
        CurX: css.X,
        CurY: css.Y,
    }
    g.AddLayouter(l)
    return l
}

func (f *FlowLayout) Reset() {
    f.CurX = f.CSS.X
    f.CurY = f.CSS.Y
    f.CurMaxY = 1
}

func (f *FlowLayout) Add(css CSS) (x0, y0, x1, y1 int) {
    maxX := f.CurX + css.W + css.ML + css.MR
    layoutMaxX := f.CSS.X + f.CSS.W

    if maxX <= layoutMaxX  {
    } else {
        f.CurY = f.CurMaxY + 1
    }

    x0, y0 = f.CurX + css.ML, f.CurY + css.MT
    x1, y1 = x0 + css.W - 1, y0 + css.H - 1
    f.CurX = x1 + css.MR + 1

    if y1 > f.CurMaxY {
        f.CurMaxY = y1
    }
    return
}
