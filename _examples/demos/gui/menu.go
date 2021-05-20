package gui

import (
	"time"

	"github.com/kasworld/h4o/_examples/app"
	"github.com/kasworld/h4o/appbase/appwindow"
	"github.com/kasworld/h4o/eventtype"
	"github.com/kasworld/h4o/gui"
	"github.com/kasworld/h4o/gui/assets/icon"

	"strings"
)

func init() {
	app.DemoMap["gui.menu"] = &GuiMenu{}
}

type GuiMenu struct{}

// Start is called once at the start of the demo.
func (t *GuiMenu) Start(a *app.App) {

	// Show and enable demo panel
	a.DemoPanel().SetRenderable(true)
	a.DemoPanel().SetEnabled(true)

	// Label
	mbText := "Selected: "
	mbOption := gui.NewLabel(mbText)
	mbOption.SetPosition(300, 10)
	mbOption.SetPaddings(2, 2, 2, 2)
	mbOption.SetBorders(1, 1, 1, 1)
	a.DemoPanel().Add(mbOption)

	// Event handler for menu clicks
	onClick := func(evname eventtype.EventType, ev interface{}) {
		path := strings.Join(ev.(*gui.MenuItem).IdPath(), "/")
		mbOption.SetText(mbText + path)
	}

	// Create menu bar
	mb := gui.NewMenuBar()
	mb.Subscribe(eventtype.OnClick, onClick)
	mb.SetPosition(10, 10)

	// Create Menu1 and adds it to the menu bar
	m1 := gui.NewMenu()
	m1.AddOption("Menu1/Option1").
		SetId("option1")
	m1.AddOption("Menu1/Option2").
		SetId("option2")
	m1.AddOption("Menu1/Option3").
		SetId("option3").
		SetEnabled(false)
	m1.AddSeparator()
	m1.AddOption("Menu1/Option4").
		SetId("option4")
	mb.AddMenu("Menu1", m1).
		SetId("menu1").
		SetShortcut(appwindow.ModAlt, appwindow.Key1)

	// Create Menu2 and adds it to the menu bar
	m2 := gui.NewMenu()
	m2.AddOption("Menu2/Option1").
		SetId("option1").
		SetIcon(icon.Build).
		SetShortcut(appwindow.ModControl, appwindow.KeyA)
	m2.AddOption("Menu2/Option two").
		SetId("option2").
		SetIcon(icon.Cached).
		SetShortcut(appwindow.ModShift, appwindow.KeyB)
	m2.AddSeparator()
	m2.AddOption("Menu2/Option three").
		SetId("option3").
		SetIcon(icon.Print).
		SetShortcut(appwindow.ModAlt, appwindow.KeyC)
	m2.AddOption("Menu2/Option four").
		SetId("option4").
		SetIcon(icon.Settings).
		SetShortcut(appwindow.ModAlt|appwindow.ModShift, appwindow.KeyD)
	m2.AddOption("Menu2/Option five").
		SetId("option5").
		SetIcon(icon.Search).
		SetShortcut(appwindow.ModAlt|appwindow.ModShift|appwindow.ModControl, appwindow.KeyE)
	mb.AddMenu("Menu2", m2).
		SetId("menu2").
		SetShortcut(appwindow.ModAlt, appwindow.Key2)

	// Create Menu3 and adds it to the menu bar
	m3 := gui.NewMenu()
	m3.AddOption("Menu3 Option1").
		SetId("option1").
		SetIcon(icon.Star).
		SetShortcut(0, appwindow.KeyF1)
	m3.AddOption("Menu3 Option2").
		SetId("option2").
		SetIcon(icon.StarBorder).
		SetShortcut(appwindow.ModControl, appwindow.KeyF2)
	// Creates Menu3/Menu1
	m3m1 := gui.NewMenu()
	m3m1.AddOption("Menu3/Menu1/Option1").
		SetId("option1").
		SetIcon(icon.StarHalf).
		SetShortcut(appwindow.ModAlt, appwindow.KeyF3)
	m3m1.AddOption("Menu3/Menu1/Option2").
		SetId("option2").
		SetIcon(icon.Opacity).
		SetShortcut(appwindow.ModAlt|appwindow.ModControl, appwindow.KeyF4)
	m3m1.AddSeparator()
	// Creates Menu3/Menu1/Menu2
	m3m1m2 := gui.NewMenu()
	m3m1m2.AddOption("Menu3/Menu1/Menu2/Option1").
		SetId("option1").
		SetIcon(icon.HourglassFull).
		SetShortcut(appwindow.ModAlt|appwindow.ModControl|appwindow.ModShift, appwindow.KeyF5)
	m3m1m2.AddOption("Menu3/Menu1/Menu2/Option2").
		SetId("option2").
		SetIcon(icon.HourglassEmpty).
		SetShortcut(0, appwindow.KeyF6)
	m3m1.AddMenu("Menu3/Menu1/Menu2", m3m1m2).
		SetId("menu2")
	m3.AddSeparator()
	m3.AddMenu("Menu3/Menu1", m3m1).
		SetId("menu1").
		SetIcon(icon.Home)
	m3.AddOption("Menu3/Option3").
		SetId("option3")
	mb.AddMenu("Menu3", m3).
		SetId("menu3").
		SetShortcut(appwindow.ModAlt, appwindow.Key3)

	// Add separators and options to the menu bar
	mb.AddSeparator()
	mb.AddOption("OptionA").
		SetId("optionA").
		SetShortcut(appwindow.ModAlt, appwindow.KeyA)
	mb.AddOption("OptionB").
		SetId("optionB").
		SetShortcut(appwindow.ModAlt, appwindow.KeyB)

	a.DemoPanel().Add(mb)
	gui.Manager().SetKeyFocus(mb)
}

// Update is called every frame.
func (t *GuiMenu) Update(a *app.App, deltaTime time.Duration) {}

// Cleanup is called once at the end of the demo.
func (t *GuiMenu) Cleanup(a *app.App) {}
