package mfi

import (
	"github.com/huin/goupnp"
	"log"
)

func TestFindDevices() {
	var err error
	var devices []goupnp.MaybeRootDevice
	if devices, err = goupnp.DiscoverDevices("urn:schemas-upnp-org:device:Basic:1"); err != nil {
		return
	}

	for _, device := range devices {
		log.Println(device.Root.Device.Manufacturer)
		log.Println(device.Root.Device.ModelName)
		log.Println(device.Root.URLBase.Host)
		log.Println(device.Root.Device.PresentationURL.URL.Host)
		log.Println(device.Location)
		log.Println(device.Root.Device.UDN)
	}
}
