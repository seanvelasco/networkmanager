package networkmanager

import (
	"fmt"

	"github.com/godbus/dbus/v5"
)

func GetNetworkByType(typeName string) (interface{}, error) {

	obj := c.Object("org.freedesktop.NetworkManager", "/org/freedesktop/NetworkManager")
	devices, err := obj.GetProperty("org.freedesktop.NetworkManager." + typeName)

	if err != nil {
		return nil, err
	}

	return devices.Value(), nil
}

func getDeviceByType(targetDeviceType uint32) ([]interface{}, error) {

	devices, err := ListDevices()

	if err != nil {
		return nil, err
	}

	var wiredDevices []interface{}

	for _, device := range devices.([]interface{}) {

		deviceType := device.(map[string]interface{})["DeviceType"].(uint32)

		if deviceType == targetDeviceType {
			wiredDevices = append(wiredDevices, device)
		}
	}

	return wiredDevices, nil

}

func GetWirelessDevicesDeprac() (interface{}, error) {
	devices, err := listDevicesByPath()

	if err != nil {
		return nil, err
	}

	var wirelessDevices []string

	for _, device := range devices.([]dbus.ObjectPath) {

		obj := c.Object("org.freedesktop.NetworkManager", device)

		wirelessCapabilities, err := obj.GetProperty("org.freedesktop.NetworkManager.Device.Wireless.WirelessCapabilities")
		if err != nil {
			continue
		}
		fmt.Println(wirelessCapabilities.Value().(uint32))
	}

	return wirelessDevices, nil

}

func getDeviceInterface(devicePath dbus.ObjectPath) (map[string]interface{}, error) {

	// ["Driver", "Ip4Connectivity"m "Ip6Connectivity", "Wired", "Wireless"]
	// Create slice

	attributes := []string{
		"Udi",
		"Path",
		"Interface",
		"IpInterface",
		"Driver",
		"DriverVersion",
		"FirmwareVersion",
		"Capabilities",
		"Ip4Address",
		"State",
		"StateReason",
		"ActiveConnection",
		"Ip4Config",
		"Dhcp4Config",
		"Ip6Config",
		"Dhcp6Config",
		"Managed",
		"Autoconnect",
		"FirmwareMissing",
		"NmPluginMissing",
		"DeviceType",
		"AvailableConnections",
		"PhysicalPortId",
		"Mtu",
		"Metered",
		"LldpNeighbors",
		"Real",
		"Ip4Connectivity",
		"Ip6Connectivity",
		"InterfaceFlags",
		"HwAddress",
	}

	deviceAttributes := make(map[string]interface{})

	for _, attribute := range attributes {
		obj := c.Object("org.freedesktop.NetworkManager", devicePath)
		property, err := obj.GetProperty("org.freedesktop.NetworkManager.Device." + attribute)
		if err != nil {
			continue
		}

		deviceAttributes[attribute] = property.Value()

	}

	deviceAttributes["DevicePath"] = devicePath

	return deviceAttributes, nil
}

func listDevicesByPath() (interface{}, error) {
	obj := c.Object("org.freedesktop.NetworkManager", "/org/freedesktop/NetworkManager")
	devices, err := obj.GetProperty("org.freedesktop.NetworkManager.Devices")

	if err != nil {
		return nil, err
	}

	return devices.Value().([]dbus.ObjectPath), nil
}

func ListDevices() (interface{}, error) {
	devices, err := listDevicesByPath()
	if err != nil {
		return nil, err
	}

	var namedDevices []interface{}

	for _, device := range devices.([]dbus.ObjectPath) {

		iface, err := getDeviceInterface(device)
		if err != nil {
			return nil, err
		}
		namedDevices = append(namedDevices, iface)
		// append device path to slice
	}

	return namedDevices, nil

}

func GetWirelessDevices() ([]interface{}, error) {
	wirelessDevices, err := getDeviceByType(NM_DEVICE_TYPE_WIFI)
	if err != nil {
		return nil, err
	}
	return wirelessDevices, nil
}

func GetWiredDevices() ([]interface{}, error) {
	wirelessDevices, err := getDeviceByType(NM_DEVICE_TYPE_ETHERNET)
	if err != nil {
		return nil, err
	}
	return wirelessDevices, nil
}

func GetBluetoothDevices() ([]interface{}, error) {
	bluetoothDevices, err := getDeviceByType(NM_DEVICE_TYPE_BT)
	if err != nil {
		return nil, err
	}
	return bluetoothDevices, nil
}
