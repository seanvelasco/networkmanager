package networkmanager

func CheckConnectivity() (interface{}, error) {

	obj := c.Object("org.freedesktop.NetworkManager", "/org/freedesktop/NetworkManager")

	var state uint32
	obj.Call("org.freedesktop.NetworkManager.CheckConnectivity", 0).Store(&state)

	return state, nil
}

func ListSavedConnections() (interface{}, error) {

	obj := c.Object("org.freedesktop.NetworkManager", "/org/freedesktop/NetworkManager/Settings")

	settings := obj.Call("org.freedesktop.NetworkManager.Settings.ListConnections", 0)

	if settings.Err != nil {
		panic(settings.Err)
	}

	return settings.Body, nil

}
