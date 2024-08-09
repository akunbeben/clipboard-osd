package main

import (
	"fmt"
	"log"

	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

func main() {
	gtk.Init(nil)

	win, err := gtk.WindowNew(gtk.WINDOW_POPUP)
	if err != nil {
		log.Fatal("Failed to create window:", err)
	}

	win.SetDefaultSize(150, 70)
	win.SetPosition(gtk.WIN_POS_CENTER)

	screen, err := gdk.ScreenGetDefault()
	if err != nil {
		log.Fatal("Failed to get screen:", err)
	}

	visual, err := screen.GetRGBAVisual()
	if err != nil {
		log.Fatal("Failed to get visual:", err)
	}

	win.SetVisual(visual)

	label, err := gtk.LabelNew("Copied to clipboard")
	if err != nil {
		log.Fatal("Failed to create label:", err)
	}
    
	win.Add(label)
    
	win.SetKeepAbove(true)
	win.SetDecorated(false)

	cssProvider, err := gtk.CssProviderNew()
	if err != nil {
		log.Fatal("Failed to create CSS provider:", err)
	}

	css := `
		window {
			border-radius: 15px;
			background-color: rgba(0, 0, 0, 0.5);
			color: #fff;
		}
	`

	err = cssProvider.LoadFromData(css)
	if err != nil {
		log.Fatal("Failed to load CSS:", err)
	}

	gtk.AddProviderForScreen(screen, cssProvider, gtk.STYLE_PROVIDER_PRIORITY_USER)

	clipboard, err := gtk.ClipboardGet(gdk.SELECTION_CLIPBOARD)
	if err != nil {
		log.Fatal("Failed to get clipboard:", err)
	}

	var lastText string

	glib.TimeoutAdd(250, func() bool {
		text, err := clipboard.WaitForText()
		if err != nil {
			log.Println("Error reading clipboard:", err)
		}

		if text != "" && text != lastText {
			lastText = text
			fmt.Println("Clipboard changed:", text)
			win.ShowAll()

			glib.TimeoutAdd(3000, func() bool {
				win.Hide()
				return false
			})
		}

		return true
	})

	gtk.Main()
}
