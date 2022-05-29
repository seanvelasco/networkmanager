package networkmanager

import (
	"github.com/godbus/dbus/v5"
)

var c *dbus.Conn
var cobj dbus.Object

type APNetwork struct {
	Ssid     string          `json:"ssid"`
	Password string          `json:"password"`
	Device   dbus.ObjectPath `json:"device"`
}

type Device interface {
	GetAccessPoints() (dbus.ObjectPath, error)
}

type AccessPoint interface {
	GetSSID() ([]byte, error)
	GetFrequency() (int32, error)
	GetStrength() (uint8, error)
	GetMode() (uint32, error)
}

type dbusBase struct {
	conn *dbus.Conn
	obj  dbus.BusObject
}

type device struct {
	dbusBase
}

type accessPoint struct {
	dbusBase
}

type connection struct {
	dbusBase
}

type Settings map[string]map[string]interface{}

type ConnectionSettings map[string]map[string]interface{}

func init() {
	conn, err := dbus.SystemBus()
	if err != nil {
		panic(err)
	}
	c = conn
}

// func main() {

// 	savedconnections, _ := ListSavedConnections(conn)

// 	fmt.Println(savedconnections)

// 	obj := conn.Object("org.freedesktop.NetworkManager", "/org/freedesktop/NetworkManager")
// 	devices2 := obj.Call("org.freedesktop.NetworkManager.GetDevices", 0)

// 	if devices2.Err != nil {
// 		panic(devices2.Err)
// 	}

// 	deviceslist := devices2.Body
// 	fmt.Println(deviceslist)

// 	newobj := conn.Object("org.freedesktop.NetworkManager", "/org/freedesktop/NetworkManager/Devices/2")

// 	neoobj := device{dbusBase{conn, newobj}}

// 	aps, err := neoobj.GetAccessPoints()
// 	if err != nil {
// 		panic(err)
// 	}

// 	fmt.Println(aps)

// 	devicetype := newobj.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.NetworkManager.Device", "DeviceType")
// 	if devicetype.Err != nil {
// 		panic(devicetype.Err)
// 	}

// 	fmt.Println(devicetype.Body)

// 	ssidscan := conn.Object("org.freedesktop.NetworkManager", "/org/freedesktop/NetworkManager/AccessPoint/3")

// 	new := accessPoint{dbusBase{conn, ssidscan}}
// 	newfreq, _ := new.GetFrequency()
// 	newstr, _ := new.GetSSID()
// 	fmt.Println(newfreq)
// 	fmt.Println(newstr)

// 	connset := conn.Object("org.freedesktop.NetworkManager", "/org/freedesktop/NetworkManager/Settings/71")

// 	cont := connection{dbusBase{conn, connset}}

// 	connid, _ := cont.GetConnectionSettings()

// 	fmt.Println(connid)

// 	settings := Settings{
// 		"connection": {
// 			"id":   "test",
// 			"type": "802-11-wireless",
// 		},
// 		"802-11-wireless": {
// 			"ssid": []byte("test"),
// 			"mode": "infrastructure",
// 		},
// 		"802-11-wireless-security": {
// 			"key-mgmt": "wpa-psk",
// 			"psk":      "PASSWORDSSSS",
// 		},
// 		"ipv4": {
// 			"method": "auto",
// 		},
// 		"ipv6": {
// 			"method": "auto",
// 		},
// 	}

// 	add, err := AddNetwork(settings)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Println(add)

// 	state, err := CheckConnectivity()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println(state)

// 	_, error := ForgetNetwork("/org/freedesktop/NetworkManager/Settings/41")
// 	if error != nil {
// 		log.Fatal(error)
// 	}

// }
