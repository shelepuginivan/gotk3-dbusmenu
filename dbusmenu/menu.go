// Package dbusmenu provides Go bindings for dbusmenu-gtk3.
package dbusmenu

/*
#cgo pkg-config: dbusmenu-gtk3-0.4 gtk+-3.0
#include <libdbusmenu-gtk/menu.h>
*/
import "C"
import (
	"errors"
	"unsafe"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

// DBusMenu is a representation of libdbusmenu-gtk3's DbusmenuGtkMenu.
type DBusMenu struct {
	gtk.Menu
}

// New is a wrapper around dbusmenu_gtkmenu_new().
func New(dbusName, dbusObject string) (*DBusMenu, error) {
	cDBusName := C.CString(dbusName)
	cDBusObject := C.CString(dbusObject)
	defer C.free(unsafe.Pointer(cDBusName))
	defer C.free(unsafe.Pointer(cDBusObject))

	ptr := C.dbusmenu_gtkmenu_new(cDBusName, cDBusObject)
	if ptr == nil {
		return nil, errors.New("dbusmenu_gtkmenu_new returned nil pointer")
	}

	cGObject := glib.ToGObject(unsafe.Pointer(ptr))

	gObject := &glib.Object{
		GObject: cGObject,
	}

	initiallyUnowned := glib.InitiallyUnowned{
		Object: gObject,
	}

	widget := gtk.Widget{
		InitiallyUnowned: initiallyUnowned,
	}

	container := gtk.Container{
		Widget: widget,
	}

	menushell := gtk.MenuShell{
		Container: container,
	}

	menu := gtk.Menu{
		MenuShell: menushell,
	}

	return &DBusMenu{menu}, nil
}
