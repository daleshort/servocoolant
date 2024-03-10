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
	log                    *log.Logger
	config                 *config.Config
	Servo1                 *servomanager.ServoManager
	Servo2                 *servomanager.ServoManager
	toolSensePin           rpio.Pin
	IsToolsenseHigh        bool
	IsProbesenseHigh		bool
	probeSensePin          rpio.Pin
	probeWritePin          rpio.Pin
	AutoToolsenseEventChan *chan bool
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

	d.toolSensePin = rpio.Pin(pinToolChanger)
	d.toolSensePin.PullDown()
	d.toolSensePin.Input()

	d.Servo1.Init()
	d.Servo2.Init()
	go d.monitorToolSensePin()
}

func (d *DeviceManager) initProbeInvert() {
	if !d.config.Viper.GetBool("devicemanager.probeinvert") {
		return
	}

	if !d.config.Viper.IsSet("devicemanager.probesensepin") {
		d.log.Error("cannot find key devicemanager.probesensepin")
		return
	}
	if !d.config.Viper.IsSet("devicemanager.probewritepin") {
		d.log.Error("cannot find key devicemanager.probewritepin")
		return
	}

	d.probeSensePin = rpio.Pin(d.config.Viper.GetInt("devicemanager.probesensepin"))
	d.probeSensePin.Input()

	d.probeWritePin = rpio.Pin(d.config.Viper.GetInt("devicemanager.probewritepin"))
	d.probeWritePin.Output()

	go d.monitorProbeSensePin()

}

func (d *DeviceManager) monitorProbeSensePin() {

	currentState := d.probeSensePin.Read()
	lastState := currentState

	
	for {
		time.Sleep(time.Millisecond * 10)
		currentState = d.probeSensePin.Read()
		if(currentState != lastState){
			d.log.Debug(fmt.Sprintf("probe state read change to %v", currentState))
			lastState = currentState
		}

		if currentState == rpio.High {
			d.IsProbesenseHigh  = true
			d.probeWritePin.Write(rpio.Low)
		} else {
			d.IsProbesenseHigh  = false
			d.probeWritePin.Write(rpio.High)
		}
	}

}

func (d *DeviceManager) isToolStateHigh() bool {

	return d.toolSensePin.Read() == rpio.High

}
func (d *DeviceManager) monitorToolSensePin() {

	d.IsToolsenseHigh = d.isToolStateHigh()
	lastRead := d.IsToolsenseHigh
	currentRead := d.IsToolsenseHigh
	for {
		time.Sleep(time.Millisecond * 10)
		currentRead = d.isToolStateHigh()
		if currentRead != lastRead {
			//do nothing
		} else {
			if currentRead != d.IsToolsenseHigh {
				// tool change state has changed with debounce
				d.IsToolsenseHigh = currentRead
				d.log.Info(fmt.Sprintf("toolchange pin status has changed to %v", d.IsToolsenseHigh))
				if d.AutoToolsenseEventChan != nil {
					*d.AutoToolsenseEventChan <- d.IsToolsenseHigh
				}

			}
		}
		lastRead = currentRead
	}
}

func (d *DeviceManager) RunRangeTest() {
	d.log.Debug("start range of motion test")
	for i := 0; i < 4; i++ {

		res := d.toolSensePin.Read()

		d.log.Debug(fmt.Sprintf("pin is %v", res))

		d.Servo1.SetMinDuty()
		d.Servo2.SetMaxDuty()

		time.Sleep(time.Millisecond * 2000)
		res = d.toolSensePin.Read()

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

		res := d.toolSensePin.Read()

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
