package networkmanager

import "fmt"

func (dbus *dbusBase) getSliceByteProperty(iface string) ([]byte, error) {

	variant, err := dbus.obj.GetProperty(iface)
	if err != nil {
		return nil, err
	}

	return variant.Value().([]byte), nil
}

func (dbus *dbusBase) getInt32Property(iface string) (uint32, error) {
	variant, err := dbus.obj.GetProperty(iface)

	if err != nil {
		return 0, err
	}
	return variant.Value().(uint32), nil
}

func (dbus *dbusBase) getUint8Property(iface string) (uint8, error) {
	variant, err := dbus.obj.GetProperty(iface)

	if err != nil {
		return 0, err
	}
	return variant.Value().(uint8), nil
}
func makeErrVariantType(iface string) error {
	return fmt.Errorf("unexpected variant type for '%s'", iface)
}
func (dbus *dbusBase) getUint32Property(iface string) (uint32, error) {
	value, err := dbus.obj.GetProperty(iface)
	if err != nil {
		return 0, makeErrVariantType(iface)
	}
	return value.Value().(uint32), nil
}

// Convert type s to type ay (byte array)
// func StringToByteArray(s string) ([]byte, error) {
// 	return []byte(s), nil
// }
