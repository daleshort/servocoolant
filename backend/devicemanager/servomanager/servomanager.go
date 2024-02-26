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
	id              int //is this servo 1 or 2
	pinNumber       int
	log             *logrus.Logger
	pin             rpio.Pin
	config          *config.Config
	minDuty         int
	maxDuty         int
	cycleLength     int
	Offset          int
	TravelRange     int
	Angle           int
	IsAuto          bool
	IsWiggle        bool
	WiggleAmplitude int
	WiggleFrequency float32
}

func GetServoManager(log *logrus.Logger, config *config.Config, id int) *ServoManager {

	if !config.Viper.IsSet("devicemanager.servo.minduty") {
		log.Error("cannot find key devicemanager.servo.minduty")
	}

	if !config.Viper.IsSet("devicemanager.servo.maxduty") {
		log.Error("cannot find key devicemanager.servo.maxduty")
	}

	if !config.Viper.IsSet("devicemanager.servo.cyclelength") {
		log.Error("cannot find key devicemanager.servo.cyclelength")
	}

	if !config.Viper.IsSet("devicemanager.servo.travelrange") {
		log.Error("cannot find key devicemanager.servo.travelrange")
	}

	var pinNumber int
	var offset int
	if id == 1 {
		if !config.Viper.IsSet("devicemanager.servo1pin") {
			log.Error("cannot find key devicemanager.servo1pin")
		}
		if !config.Viper.IsSet("devicemanager.servo1offset") {
			log.Error("cannot find key devicemanager.servo1offset")
		}
		pinNumber = config.Viper.GetInt("devicemanager.servo1pin")
		offset = config.Viper.GetInt("devicemanager.servo1offset")
	} else if id == 2 {
		if !config.Viper.IsSet("devicemanager.servo2pin") {
			log.Error("cannot find key devicemanager.servo2pin")
		}
		if !config.Viper.IsSet("devicemanager.servo2offset") {
			log.Error("cannot find key devicemanager.servo2offset")
		}
		pinNumber = config.Viper.GetInt("devicemanager.servo2pin")
		offset = config.Viper.GetInt("devicemanager.servo2offset")
	} else {
		log.Error(fmt.Sprintf("invalid servo id:%v", id))
	}

	return &ServoManager{
		id:          id,
		log:         log,
		config:      config,
		minDuty:     config.Viper.GetInt("devicemanager.servo.minduty"),
		maxDuty:     config.Viper.GetInt("devicemanager.servo.maxduty"),
		cycleLength: config.Viper.GetInt("devicemanager.servo.cyclelength"),
		TravelRange: config.Viper.GetInt("devicemanager.servo.travelrange"),
		pinNumber:   pinNumber,
		Offset:      offset,
	}
}

func (s *ServoManager) Init() {

	s.pin = rpio.Pin(s.pinNumber)
	rpio.PinMode(s.pin, rpio.Pwm)

	if !s.config.Viper.IsSet("devicemanager.servo.clockfreq") {
		s.log.Error("cannot find key devicemanager.servo.clockfreq")
	}

	s.pin.Freq(s.config.Viper.GetInt("devicemanager.servo.clockfreq"))
	rpio.StartPwm()

	s.SetAngle(0)

}

func (s *ServoManager) SetDutyCycle(dutyLength int) {

	s.pin.DutyCycle(uint32(dutyLength), uint32(s.cycleLength))

}

func (s *ServoManager) SetMinDuty() {
	s.SetDutyCycle(s.minDuty)
}

func (s *ServoManager) SetMaxDuty() {
	s.SetDutyCycle(s.maxDuty)
}

func (s *ServoManager) SetAngle(angle int) error {

	// servo has 0 to 270 degree range
	// pretend servo is tilted back 15 degrees
	// real world range is -15 degrees to 255 degrees
	// offset is set to 15 so if i specify 0 degrees, servo actually goes an extra 15 degrees so it looks like 0

	if angle > s.TravelRange-s.Offset || angle < 0-s.Offset {
		s.log.Error(fmt.Sprintf("invalid angle specified: %v", angle))
		return fmt.Errorf("invalid angle specified: %v", angle)
	}

	s.Angle = angle
	adjustedAngle := angle + s.Offset

	dutyRange := s.maxDuty - s.minDuty

	anglePct := float32(adjustedAngle) / float32(s.TravelRange)
	dutyResult := s.minDuty + int(anglePct*float32(dutyRange))

	s.SetDutyCycle(dutyResult)

	return nil
}
