package ui

import (
	"github.com/gotk3/gotk3/gtk"
)

type SplashPanel struct {
	CommonPanel
	Label       *gtk.Label
	RetryButton *gtk.Button
}

func NewSplashPanel(ui *UI) *SplashPanel {
	m := &SplashPanel{CommonPanel: NewCommonPanel(ui, nil)}
	m.initialize()
	return m
}

func (m *SplashPanel) initialize() {
	logo := MustImageFromFile("logo.png")
	m.Label = MustLabel("...")
	m.Label.SetHExpand(true)
	m.Label.SetLineWrap(true)
	m.Label.SetMaxWidthChars(30)
	m.Label.SetText("Initializing printer...")

	main := MustBox(gtk.ORIENTATION_VERTICAL, 15)
	main.SetVAlign(gtk.ALIGN_END)
	main.SetVExpand(true)
	main.SetHExpand(true)

	main.Add(logo)
	main.Add(m.Label)

	box := MustBox(gtk.ORIENTATION_VERTICAL, 0)
	box.Add(main)
	box.Add(m.createActionBar())

	m.Grid().Add(box)
}

func (m *SplashPanel) createActionBar() gtk.IWidget {
	bar := MustBox(gtk.ORIENTATION_HORIZONTAL, 5)
	bar.SetHAlign(gtk.ALIGN_END)

	m.RetryButton = MustButtonImageStyle("Retry", "refresh.svg", "color2", m.releaseFromHold)
	m.RetryButton.SetProperty("width-request", m.Scaled(100))
	m.RetryButton.SetProperty("visible", true)
	bar.Add(m.RetryButton)
	ctx, _ := m.RetryButton.GetStyleContext()
	ctx.AddClass("hidden")

	sys := MustButtonImageStyle("System", "info.svg", "color3", m.showSystem)
	sys.SetProperty("width-request", m.Scaled(100))
	bar.Add(sys)

	net := MustButtonImageStyle("Network", "network.svg", "color4", m.showNetwork)
	net.SetProperty("width-request", m.Scaled(100))
	bar.Add(net)

	return bar
}

func (m *SplashPanel) putOnHold() {
	m.RetryButton.Show()
	ctx, _ := m.RetryButton.GetStyleContext()
	ctx.RemoveClass("hidden")
	m.Label.SetText("Cannot connect initialize the printer. Tap \"Retry\" to try again.")
}

func (m *SplashPanel) releaseFromHold() {
	m.RetryButton.Hide()
	ctx, _ := m.RetryButton.GetStyleContext()
	ctx.AddClass("hidden")

	m.Label.SetText("Loading...")
	m.UI.connectionAttempts = 0
}

func (m *SplashPanel) showNetwork() {
	m.UI.Add(NetworkPanel(m.UI, m))
}

func (m *SplashPanel) showSystem() {
	m.UI.Add(SystemPanel(m.UI, m))
}
