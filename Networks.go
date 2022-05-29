package networkmanager

import (
	"fmt"

	"github.com/godbus/dbus/v5"
)

func AddNetwork(ssid, password string) error {

	settings := Settings{
		"connection": {
			"id":   ssid,
			"type": "802-11-wireless",
		},
		"802-11-wireless": {
			"ssid": []byte(ssid),
			"mode": "infrastructure",
		},
		"802-11-wireless-security": {
			"key-mgmt": "wpa-psk",
			"psk":      password,
		},
		"ipv4": {
			"method": "auto",
		},
		"ipv6": {
			"method": "auto",
		},
	}

	obj := c.Object("org.freedesktop.NetworkManager", "/org/freedesktop/NetworkManager/Settings")

	view := obj.Call("org.freedesktop.NetworkManager.Settings.AddConnection", 0, settings) //.Store(&id)

	if view.Err != nil {
		return view.Err
	}

	// view.Body[0]

	return nil
}

func ForgetNetwork(path dbus.ObjectPath) (interface{}, error) {
	obj := c.Object("org.freedesktop.NetworkManager", path)
	state := obj.Call("org.freedesktop.NetworkManager.Settings.Connection.Delete", 0)

	if state.Err != nil {
		return nil, fmt.Errorf("Network does not exist")
	}

	return state.Body, nil
}

func GetAccessPoints(path dbus.ObjectPath) (interface{}, error) {

	obj := c.Object("org.freedesktop.NetworkManager", path)

	aps, err := obj.GetProperty("org.freedesktop.NetworkManager.Device.Wireless.AccessPoints")

	if err != nil {
		return nil, err
	}

	var rv []interface{}

	for _, ap := range aps.Value().([]dbus.ObjectPath) {
		settings, err := GetAccessPointInfo(ap)
		if err != nil {
			fmt.Println(err)
		}
		rv = append(rv, settings)
	}

	if err != nil {
		return nil, err
	}

	return rv, nil

}

func InternetSharingOverEthernet(device dbus.ObjectPath) {

	obj := c.Object("org.freedesktop.NetworkManager", "/org/freedesktop/NetworkManager/Settings")

	settings := &Settings{
		"connection": map[string]interface{}{
			"id": "Internet Sharing over Ethernet",
			// "uuid": "uuid",
			"type": "802-3-ethernet",
		},
		"802-3-ethernet": map[string]interface{}{
			// "wake-on-lan": true,
			"auto-negotiate": false,
		},
		"ipv4": map[string]interface{}{
			"method": "shared",
		},
		"ipv6": map[string]interface{}{
			"method": "ignore",
		},
	}
	var call *dbus.Call

	call = obj.Call("org.freedesktop.NetworkManager.Settings.AddConnection", 0, settings)

	if call.Err != nil {
		fmt.Println(call.Err)
	}
	fmt.Println(call.Body)

	call2, err := ActivateConnection(call.Body[0].(dbus.ObjectPath), dbus.ObjectPath(device))
	if err != nil {
		panic(err)
	}

	fmt.Println(call2)

}
