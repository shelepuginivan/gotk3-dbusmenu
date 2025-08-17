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

func wrapDBusMenu(obj *glib.Object) *DBusMenu {
	if obj == nil {
		return nil
	}

	return &DBusMenu{
		Menu: gtk.Menu{
			MenuShell: gtk.MenuShell{
				Container: gtk.Container{
					Widget: gtk.Widget{
						InitiallyUnowned: glib.InitiallyUnowned{
							Object: obj,
						},
					},
				},
			},
		},
	}
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

	return wrapDBusMenu(glib.Take(unsafe.Pointer(ptr))), nil
}
