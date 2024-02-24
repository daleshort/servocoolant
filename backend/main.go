package main

import (
	"fmt"

	"github.com/stianeikeland/go-rpio/v4"
	config "mechied.com/servocoolant/config"
	slog "mechied.com/servocoolant/logger"
	"time"
)

func main() {
	log := slog.GetLog()
	log.Info("Initializing Servo Coolant Nozzle Application ")
	config := config.GetConfig(log)
	fmt.Println(config.GetVersion())
	
	err := rpio.Open()

	if err != nil {
		log.Error(fmt.Sprintf("error opening pin %v", err))
	}

	// pin :=rpio.Pin(13)  //BCM pin 13 is pin 33 on the actual physical pin headers
	// pin.Input()
	// pin.PullDown()

	//500-2500μsec pulse = .0005 seconds to .0025 = 50 cycles to 250 cycles
	//50hz cycle = .02 seconds per cycle = 2000 clock cycles
	//100000 hz clock =  .00001 second per cycle

	pwm1 := rpio.Pin(13)
	rpio.PinMode(pwm1,rpio.Pwm)
	rpio.SetDutyCycleWithPwmMode(pwm1,150,2000,true)
	pwm1.DutyCycle(150, 2000)
	pwm1.Freq(100000)
	rpio.StartPwm()

	for i := 0; i < 10; i++ {
		
	

		// res := pin.Read()

		// log.Debug(fmt.Sprintf("pin is %v", res))
		rpio.StopPwm()
		rpio.SetDutyCycleWithPwmMode(pwm1,150,2000,true)
		pwm1.DutyCycle(150, 2000)
		rpio.StartPwm()

		time.Sleep(time.Millisecond*1000)
		rpio.StopPwm()
		rpio.SetDutyCycleWithPwmMode(pwm1,50,2000,true)
		pwm1.DutyCycle(50, 2000)
		rpio.StartPwm()


		time.Sleep(time.Millisecond*1000)

		log.Debug("pwm cycle")



	}
 rpio.Close()
 log.Debug("end")
}
