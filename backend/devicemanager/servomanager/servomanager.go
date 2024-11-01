package servomanager

import (
	"fmt"
	"math"
	"time"

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
	wiggleTime      time.Time
	SoftLimitMin    int
	SoftLimitMax    int
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

	if !config.Viper.IsSet("devicemanager.servo.softlimitmin") {
		log.Error("cannot find key devicemanager.servo.softlimitmin")
	}

	if !config.Viper.IsSet("devicemanager.servo.softlimitmax") {
		log.Error("cannot find key devicemanager.servo.softlimitmax")
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
		id:              id,
		log:             log,
		config:          config,
		minDuty:         config.Viper.GetInt("devicemanager.servo.minduty"),
		maxDuty:         config.Viper.GetInt("devicemanager.servo.maxduty"),
		cycleLength:     config.Viper.GetInt("devicemanager.servo.cyclelength"),
		TravelRange:     config.Viper.GetInt("devicemanager.servo.travelrange"),
		pinNumber:       pinNumber,
		Offset:          offset,
		wiggleTime:      time.Now(),
		IsWiggle:        false,
		IsAuto:          true,
		WiggleAmplitude: 0,
		WiggleFrequency: 0,
		SoftLimitMin:    config.Viper.GetInt("devicemanager.servo.softlimitmin"),
		SoftLimitMax:    config.Viper.GetInt("devicemanager.servo.softlimitmax"),
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

	s.SetAngle( s.SoftLimitMin)
	go s.maintainServoAngle()
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

func (s *ServoManager) maintainServoAngle() {
	for {
		time.Sleep(time.Millisecond * 50)

		if !s.IsWiggle {
			s.moveServoToAngle(float64(s.Angle))
		} else {

			timeEllapsed := float64(time.Since(s.wiggleTime).Seconds())
			amplitudeShiftFactor := math.Sin(2 * math.Pi * float64(s.WiggleFrequency) * timeEllapsed)
			amplitudeShift := amplitudeShiftFactor * float64(s.WiggleAmplitude)
			s.moveServoToAngle(float64(s.Angle) + amplitudeShift)
		}
	}
}

func (s *ServoManager) TestServoWiggle() {

	s.Angle = 90
	s.IsWiggle = true
	s.WiggleAmplitude = 5
	s.WiggleFrequency = 1

	time.Sleep(time.Second * 5)

	s.Angle = 20
	s.IsWiggle = true
	s.WiggleAmplitude = 10
	s.WiggleFrequency = 1
	time.Sleep(time.Second * 5)

	s.Angle = 0
	s.IsWiggle = false

}

func (s *ServoManager) moveServoToAngle(angle float64) {
	// servo has 0 to 270 degree range
	// pretend servo is tilted back 15 degrees
	// real world range is -15 degrees to 255 degrees
	// offset is set to 15 so if i specify 0 degrees, servo actually goes an extra 15 degrees so it looks like 0

	adjustedAngle := angle + float64(s.Offset)

	dutyRange := s.maxDuty - s.minDuty

	anglePct := float32(adjustedAngle) / float32(s.TravelRange)
	dutyResult := s.minDuty + int(anglePct*float32(dutyRange))

	s.SetDutyCycle(dutyResult)

}

func (s *ServoManager) SetAngle(angle int) error {
	if angle > s.TravelRange-s.Offset || angle < 0-s.Offset {
		s.log.Error(fmt.Sprintf("invalid angle specified exceeds travel range: %v", angle))
		return fmt.Errorf("invalid angle specified  exceeds travel range: %v", angle)
	}

	if angle > s.SoftLimitMax || angle < s.SoftLimitMin {
		s.log.Error(fmt.Sprintf("invalid angle specified exceeds soft limits: %v", angle))
		return fmt.Errorf("invalid angle specified exceeds soft limits: %v", angle)
	}


	s.Angle = angle
	return nil
}
