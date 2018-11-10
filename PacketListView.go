package main

import "github.com/gotk3/gotk3/gtk"
import "github.com/gskartwii/roblox-dissector/peer"

const (
    PACKETLIST_COLUMN_ID = iota
    PACKETLIST_COLUMN_NAME
    PACKETLIST_COLUMN_DIRECTION
    PACKETLIST_COLUMN_SIZE
)

func createTreeColumn(header string, id int) *gtk.TreeViewColumn {
    render, err := gtk.CellRendererTextNew()
    if err != nil {
        panic(err)
    }

    column, err := gtk.TreeViewColumnNewWithAttribute(header, render, "text", id)
    if err != nil {
        panic(err)
    }
    return column
}

type packetViewWindow struct {
    *gtk.Window
    packetTree *gtk.TreeView
    packetTreeModel *gtk.ListStore
    packetInfoStack *gtk.Stack

    packets []*peer.PacketLayers
}

func (win *packetViewWindow) addRow(id uint32, name string, direction string, size uint32) {
    iter := win.packetTreeModel.Append()
    err := win.packetTreeModel.Set(iter, []int{PACKETLIST_COLUMN_ID, PACKETLIST_COLUMN_NAME, PACKETLIST_COLUMN_DIRECTION, PACKETLIST_COLUMN_SIZE}, []interface{}{id, name, direction, size})
    if err != nil {
        panic(err)
    }
}
func (win *packetViewWindow) reset() {
    packetTreeModel.Clear()
    win.packets = []*peer.PacketLayers{}
}

func newPacketViewWindow(title string) *packetViewWindow {
    win := &packetViewWindow{newMainWindow(title)}
    win.packetTree = gtk.TreeViewNew()
    win.packetTree.AppendColumn(createTreeColumn("ID"), PACKETLIST_COLUMN_ID)
    win.packetTree.AppendColumn(createTreeColumn("Name"), PACKETLIST_COLUMN_NAME)
    win.packetTree.AppendColumn(createTreeColumn("Direction"), PACKETLIST_COLUMN_DIRECTION)
    win.packetTree.AppendColumn(createTreeColumn("Size"), PACKETLIST_COLUMN_SIZE)
    win.packetTreeModel = gtk.ListStoreNew(glib.TYPE_INT, glib.TYPE_STRING, glib.TYPE_STRING, glib.TYPE_INT)
    win.packetTree.SetModel(win.packetTreeModel)

    return win
}
