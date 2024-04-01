package util

import "github.com/gofiber/fiber/v2"

// get device info from fiber request
type DeviceInfo struct {
	Id       string `json:"device_id"`
	Name     string `json:"device_name"`
	Ip       string `json:"device_ip"`
	Agent    string `json:"device_agent"`
	Location string `json:"device_location"`
}

func GetDeviceInfo(c *fiber.Ctx) DeviceInfo {
	return DeviceInfo{
		Id:       c.Get("X-Device-Id"),
		Name:     c.Get("X-Device-Name"),
		Ip:       c.IP(),
		Agent:    c.Get("User-Agent"),
		Location: c.Get("X-Device-Location"),
	}
}
