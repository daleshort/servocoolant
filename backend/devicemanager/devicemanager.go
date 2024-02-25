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
	log      *log.Logger
	config   *config.Config
	servo1   *servomanager.ServoManager
	servo2   *servomanager.ServoManager
	sensePin rpio.Pin
}

func GetDeviceManager(log *log.Logger, config *config.Config) *DeviceManager {

	d := DeviceManager{
		log:    log,
		config: config,
		servo1: servomanager.GetServoManager(log, config,1),
		servo2: servomanager.GetServoManager(log, config,2),
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

	d.servo1.Init()
   d.servo2.Init()
}

func (d *DeviceManager) RunTest() {

	for i := 0; i <4; i++ {

		res := d.sensePin.Read()

		d.log.Debug(fmt.Sprintf("pin is %v", res))

		d.servo1.SetMinDuty()
		d.servo2.SetMaxDuty()

		time.Sleep(time.Millisecond * 2000)
		res = d.sensePin.Read()

		d.log.Debug(fmt.Sprintf("pin is %v", res))
		d.servo1.SetMaxDuty()
		d.servo2.SetMinDuty()

		time.Sleep(time.Millisecond * 2000)

		d.log.Debug("pwm cycle")

	}
	log.Debug("end")
}
