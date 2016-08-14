package menu

import (
	"github.com/jroimartin/gocui"
	"log"
	"fmt"
	"strings"
)


type gui struct {
	gui *gocui.Gui
	list *[]ListItem
	curr int
	count int
	selected bool
	selection int
}

func InitGui() *gui {
	g := new(gui)
	g.selected = false
	g.gui = gocui.NewGui()

	if err := g.gui.Init(); err != nil{
		log.Panicln(err)
	}

	return g
}

func (g *gui) keyBindings() error {
	if err := g.gui.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, g.quit); err != nil { return err }
	if err := g.gui.SetKeybinding("", 'q', gocui.ModNone, g.quit); err != nil { return err }
	if err := g.gui.SetKeybinding("", gocui.KeyEnter, gocui.ModNone, g.enter); err != nil { return err }
	if err := g.gui.SetKeybinding("", gocui.KeyArrowUp, gocui.ModNone, g.prev); err != nil { return err }
	if err := g.gui.SetKeybinding("", gocui.KeyArrowDown, gocui.ModNone, g.next); err != nil { return err }

	return nil
}



func (g *gui) Run(lst *List) (*ListItem, bool){
	g.list = &lst.List
	g.curr = 0
	g.count = len(lst.List)

	g.gui.SelBgColor = gocui.ColorGreen
	g.gui.SelFgColor = gocui.ColorWhite
	//g.gui.Cursor = true

	g.gui.SetLayout(g.layout)
	if err := g.keyBindings(); err != nil {
		log.Panicln(err)
	}

	if err := g.gui.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}

	defer g.terminate()

	if g.selected {
		return &(*g.list)[g.curr], true
	} else {
		return nil, false
	}
}

func (g *gui) terminate(){
	g.gui.Close()
}

func (g *gui) next(gui *gocui.Gui, v *gocui.View) error {
	if v != nil {
		cx, cy := v.Cursor()

		next := g.curr < g.count - 1
		g.curr++

		if g.curr >= g.count {
			g.curr = g.count - 1
			return nil
		}

		if err := v.SetCursor(cx, cy + 1); err != nil && next {
			ox, oy := v.Origin()
			if err := v.SetOrigin(ox, oy + 1); err != nil {
				return err
			}
		}
	} else {
		return gocui.ErrQuit
	}
	return nil
}

func (g *gui) prev(gui *gocui.Gui, v *gocui.View) error {
	if v != nil {
		cx, cy := v.Cursor()
		ox, oy := v.Origin()

		g.curr--
		if(g.curr < 0) {
			g.curr = 0
		}

		if err := v.SetCursor(cx, cy - 1); err != nil && oy > 0 {
			if err := v.SetOrigin(ox, oy - 1); err != nil {
				return err
			}
		}
	} else {
		return gocui.ErrQuit
	}
	return nil
}

func (g *gui) enter(gui *gocui.Gui, v *gocui.View) error {

	g.selected = true

	return gocui.ErrQuit
}

func (g *gui) quit(gui *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

func (g *gui) layout(gui *gocui.Gui) error {
	maxx, maxy := gui.Size()
	v, err := gui.SetView("menu", 0, 0, maxx-1, maxy-4)
	if err != nil {
		if err != gocui.ErrUnknownView { return err }
		g.draw(v)
		if err := gui.SetCurrentView("menu"); err != nil {return err}
	} else {
		v.Clear()
		g.draw(v)
	}

	sv, err := gui.SetView("status", 0, maxy - 3, maxx-1, maxy - 1)
	if err != nil {
		if err != gocui.ErrUnknownView { return err }
		g.drawStatus(sv)
	} else {
		sv.Clear()
		g.drawStatus(sv)
	}

	return nil
}

func (g *gui) draw(v *gocui.View) {
	maxx, _ := v.Size()
	v.Highlight = true
	for idx, item := range *(g.list) {
		/*if idx == g.curr {
		}*/
		var line string
		if len(item.Desc) > 0 {
			line = fmt.Sprintf("[%2d] %s (%s)", idx, item.Title, item.Desc)
		} else {
			line = fmt.Sprintf("[%2d] %s", idx, item.Title)
		}

		lineLen := len([]rune(line))
		if lineLen < maxx {
			line += strings.Repeat(" ", maxx - lineLen)
		}

		fmt.Fprintln(v, line)
	}
}

func (g *gui) drawStatus(v *gocui.View) {

	str := fmt.Sprintf("Press 'q' to exit. [%2d / %d]:", g.curr+1, g.count)

	for _, c := range ((*g.list)[g.curr].Cmd) {
		if strings.ContainsAny(c, " \t") {
			str += " \"" + c + "\""
		} else {
			str += " " + c
		}
	}

	fmt.Fprintln(v, str)
}