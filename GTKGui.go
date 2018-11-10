package main

import "github.com/gotk3/gotk3/gtk"
import "log"

func newMainWindow(title string) *gtk.Window {
    win, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
    if err != nil {
        panic(err)
    }
    win.SetTitle(title)
    win.Connect("destroy", func() {
        gtk.MainQuit()
    }
    win.SetDefaultSize(800, 600)
    win.SetPosition(gtk.WIN_POS_CENTER)
    return win
}

func GUIMain() {
    gtk.Init(nil)
    window := newMainWindow("Roblox Dissector")

}
