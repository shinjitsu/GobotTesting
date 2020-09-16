package main

import (
	"fmt"
	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/aio"
	g "gobot.io/x/gobot/platforms/dexter/gopigo3"
	"gobot.io/x/gobot/platforms/raspi"
	"time"
)

func main(){
	//We create the adaptors to connect the GoPiGo3 board with the Raspberry Pi 3
	raspiAdaptor := raspi.NewAdaptor()
	gopigo3 := g.NewDriver(raspiAdaptor)
	lightSensor:= aio.NewGroveLightSensorDriver(gopigo3, "AD_1_1")

	sensorReader := func(){
		for {
			sensorVal, err := lightSensor.Read()
			if err!=nil{
				fmt.Errorf("Error reading sensor %+v", err)
			}
			fmt.Println(sensorVal)
			time.Sleep(time.Second)
		}
//		lightSensor.On(lightSensor.Event("Data"), func(sensorVal interface{}, ){
//			fmt.Println(sensorVal)
//		})
	}



	robot := gobot.NewRobot("gopigo3sensorChecker",
		[]gobot.Connection{raspiAdaptor},
		[]gobot.Device{gopigo3, lightSensor},
		sensorReader,
	)

	robot.Start()
}