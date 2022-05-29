# NetworkManager

## Installation

```
go get github.com/seanvelasco/networkmanager
```

## Methods

## Common issues

### Wired device is strictly unmanaged

This problem typically occurs in systems running Ubuntu Server where NetworkManager is not originally installed.

Navigate to the below directory and make sure that the netplan configuration file is present.


```
cd /etc/netplan
```

```
sudo nano 50-cloud-init.yaml
```


If it does not exist in the given directory, run the following

```
sudo netplan generate && sudo netplan apply
```

Edit `/etc/netplan/50-cloud-init.yaml` and add **NetworkManager** as the renderer

```
network:
    renderer: NetworkManager
    ethernets:
        eth0:
            dhcp4: true
    version: 2
```

## Compatability

This software can be installed on any Linux system that uses NetWorkManager.

NetworkManager is bundled with most Debian-based Linux distributions. However, it can be installed on any Linux distribution.

This software, unfortunately, is not compatible with Windows and MacOS which do not use NetworkManager.

This software can be run in a Docker container, but a mapping to the host's DBUS is required in the Dockerfile, thus the same limitation applies if the Docker image is not Linux-based.

Successfully tested on Ubuntu Desktop, Ubuntu Server, & BalenaOS.

## Author

### [Sean Velasco](https://seanvelasco.com)