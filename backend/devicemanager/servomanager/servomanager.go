package servomanager

import (
	"fmt"

	logrus "github.com/sirupsen/logrus"
	"github.com/stianeikeland/go-rpio/v4"
	config "mechied.com/servocoolant/config"
)

	//500-2500μsec  pulse = .0005 seconds to .0025 = 100 cycles to 500 cycles
	//50hz cycle = .02 seconds per cycle = 4000 clock cycles
	//200000 hz clock =  .000005 second per cycle

type ServoManager struct {
	log *logrus.Logger
	pin rpio.Pin
	config   *config.Config
	minDuty int
	maxDuty int
	cycleLength int
}

func GetServoManager(log *logrus.Logger,  config *config.Config) *ServoManager {


	if !config.Viper.IsSet("devicemanager.servo.minduty") {
		log.Error("cannot find key devicemanager.servo.minduty")
	}

	if !config.Viper.IsSet("devicemanager.servo.maxduty") {
		log.Error("cannot find key devicemanager.servo.maxduty")
	}

	if !config.Viper.IsSet("devicemanager.servo.cyclelength") {
		log.Error("cannot find key devicemanager.servo.cyclelength")
	}


	return &ServoManager{
		log: log,
		config: config,
		minDuty: config.Viper.GetInt("devicemanager.servo.minduty"),
		maxDuty: config.Viper.GetInt("devicemanager.servo.maxduty"),
		cycleLength: config.Viper.GetInt("devicemanager.servo.cyclelength"),
	}
}

func (s *ServoManager) Init(pinNumber int) {

	s.pin = rpio.Pin(pinNumber)
	rpio.PinMode(s.pin, rpio.Pwm)

	if !s.config.Viper.IsSet("devicemanager.servo.clockfreq") {
		s.log.Error("cannot find key devicemanager.servo.clockfreq")
	}

	s.pin.Freq(s.config.Viper.GetInt("devicemanager.servo.clockfreq"))
	rpio.StartPwm()

	s.log.Info(fmt.Sprintf("started servo manager with pin number %v", pinNumber))
}

func (s *ServoManager) SetDutyCycle(dutyLength int) {

	s.pin.DutyCycle(uint32(dutyLength), uint32(s.cycleLength))

}

func (s *ServoManager) SetMinDuty(){
	s.SetDutyCycle(s.minDuty)
}

func (s *ServoManager) SetMaxDuty(){
	s.SetDutyCycle(s.maxDuty)
}