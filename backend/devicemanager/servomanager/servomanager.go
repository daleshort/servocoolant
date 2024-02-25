package servomanager

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/stianeikeland/go-rpio/v4"
)

	//500-2500μsec  pulse = .0005 seconds to .0025 = 50 cycles to 250 cycles
	//50hz cycle = .02 seconds per cycle = 2000 clock cycles
	//100000 hz clock =  .00001 second per cycle

type ServoManager struct {
	log *log.Logger
	pin rpio.Pin
}

func GetServoManager(log *log.Logger) *ServoManager {

	return &ServoManager{
		log: log,
	}
}

func (s *ServoManager) Init(pinNumber int) {

	s.pin = rpio.Pin(pinNumber)
	rpio.PinMode(s.pin, rpio.Pwm)
	s.pin.Freq(100000)
	rpio.StartPwm()

	log.Info(fmt.Sprintf("started servo manager with pin number %v", pinNumber))
}

func (s *ServoManager) SetDutyCycle(dutyLenght int, cycleLength int) {

	s.pin.DutyCycle(uint32(dutyLenght), uint32(cycleLength))

}
