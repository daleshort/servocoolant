package devicemanager

import (
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/stianeikeland/go-rpio/v4"
	config "mechied.com/servocoolant/config"
	"mechied.com/servocoolant/devicemanager/servomanager"
)

type DeviceManager struct {
	log             *log.Logger
	config          *config.Config
	Servo1          *servomanager.ServoManager
	Servo2          *servomanager.ServoManager
	sensePin        rpio.Pin
	IsToolsenseHigh bool
}

func GetDeviceManager(log *log.Logger, config *config.Config) *DeviceManager {

	d := DeviceManager{
		log:    log,
		config: config,
		Servo1: servomanager.GetServoManager(log, config, 1),
		Servo2: servomanager.GetServoManager(log, config, 2),
	}
	d.init()
	return &d
}

func (d *DeviceManager) init() {

	d.log.Info("initalizing devicemanager")

	err := rpio.Open()

	if err != nil {
		d.log.Error(fmt.Sprintf("error opening pin %v", err))
	}

	if !d.config.Viper.IsSet("devicemanager.toolchangepin") {
		d.log.Error("cannot find key devicemanager.toolchangepin")
	}
	pinToolChanger := d.config.Viper.GetInt("devicemanager.toolchangepin")

	d.sensePin = rpio.Pin(pinToolChanger)
	d.sensePin.PullDown()
	// rpio.PinMode(pin, rpio.Input)
	// rpio.PullMode(pin, rpio.PullDown)
	d.sensePin.Input()

	d.Servo1.Init()
	d.Servo2.Init()
	go d.monitorSensePin()
}

func (d *DeviceManager) isStateHigh() bool {

	return d.sensePin.Read() == rpio.High

}
func (d *DeviceManager) monitorSensePin() {

	d.IsToolsenseHigh = d.isStateHigh()
	lastRead := d.IsToolsenseHigh
	currentRead := d.IsToolsenseHigh
	for {
		time.Sleep(time.Millisecond * 10)
		currentRead = d.isStateHigh()
		if currentRead != lastRead {
			//do nothing
		} else {
			if currentRead != d.IsToolsenseHigh {
				// tool change state has changed with debounce
				d.IsToolsenseHigh = currentRead
				d.log.Info(fmt.Sprintf("toolchange pin status has changed to %v", d.IsToolsenseHigh))
			}
		}
		lastRead = currentRead
	}
}

func (d *DeviceManager) RunRangeTest() {
	d.log.Debug("start range of motion test")
	for i := 0; i < 4; i++ {

		res := d.sensePin.Read()

		d.log.Debug(fmt.Sprintf("pin is %v", res))

		d.Servo1.SetMinDuty()
		d.Servo2.SetMaxDuty()

		time.Sleep(time.Millisecond * 2000)
		res = d.sensePin.Read()

		d.log.Debug(fmt.Sprintf("pin is %v", res))
		d.Servo1.SetMaxDuty()
		d.Servo2.SetMinDuty()

		time.Sleep(time.Millisecond * 2000)

		d.log.Debug("pwm cycle")

	}

	d.log.Debug("end range of motion test")
}

func (d *DeviceManager) RunAngleTest() {
	d.log.Debug("start angle test")
	for i := 0; i < 4; i++ {

		res := d.sensePin.Read()

		d.log.Debug(fmt.Sprintf("pin is %v", res))

		d.Servo1.SetAngle(0)
		d.log.Debug("angle 0")

		time.Sleep(time.Millisecond * 2000)

		d.Servo1.SetAngle(180)
		d.log.Debug("angle 180")
		time.Sleep(time.Millisecond * 2000)

		d.Servo1.SetAngle(90)
		d.log.Debug("angle 90")
		time.Sleep(time.Millisecond * 2000)

	}

	d.log.Debug("end angle test")
}
