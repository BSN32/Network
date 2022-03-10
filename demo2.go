package jsontest

import (
	"encoding/json"
	"fmt"
	"os"
)

/*type Devices struct {
	Devices []Device `json:"devices"`
}
*/
type Device struct {
	DeviceTitle   string `json:"device_title"`
	Interface1     string `json:"interface"`
	InterfaceName string `json:"interface_name"`
	IpAddress     string `json:"ip_address"`
}

func Demo2() {
	jsonFile, err := os.Open("list.json")
	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()

	//byteValue, _ := ioutil.ReadAll(jsonFile)

	var device []Device

	var deviceDecoder *json.Decoder = json.NewDecoder(jsonFile)
	if err != nil {
		fmt.Println(err)
	}
	err = deviceDecoder.Decode(&device)
	if err != nil {
		fmt.Println(err)
	}
	for i, dev := range device {
		fmt.Println(i + 1)
		fmt.Println(dev.DeviceTitle)
		fmt.Println(dev.IpAddress)
	}
	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	//json.Unmarshal(byteValue, &devices)

	/*for i := 0; i < len(device.Device); i++ {
		fmt.Println("device_title: " + device.Device[i].Device_title)
		fmt.Println("interface: " + device.Device[i].Interface1)
		fmt.Println("interface_name: " + device.Device[i].Interface_name)
		fmt.Println("ip_address: " + device.Device[i].Ip_address)
	}
	*/
}
