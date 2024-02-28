package automanager

import "time"

func (a *AutoManager) ActivateToolLength(toolId int) error {

	angle, err := a.CalculateAngleForToolLength(toolId)

	if err != nil {
		a.log.Error("unable to activate tool length")
		return err
	}

	if a.devicemanager.Servo1.IsAuto {
		a.devicemanager.Servo1.SetAngle(*angle)
	} else {
		a.log.Debug("skipped setting angle for servo1. servo is in manual mode")
	}

	if a.devicemanager.Servo2.IsAuto {
		a.devicemanager.Servo2.SetAngle(*angle)
	} else {
		a.log.Debug("skipped setting angle for servo1. servo is in manual mode")
	}
	return nil
}

func (a *AutoManager) AddToolToQueue(toolId int) {
	a.ToolQueue = append(a.ToolQueue, toolId)

	//If we just added a tool to the queue when there was nothing in the queue on program start
	if a.CurrentToolQueuePosition == 0 && len(a.ToolQueue) == 1 {
		a.activateCurrentTool()
	}
}

func (a *AutoManager) activateCurrentTool() {
	currentTool := a.ToolQueue[a.CurrentToolQueuePosition]
	//activate the tool
	a.ActivateToolLength(currentTool)
}

func (a *AutoManager) HandleEndOfProgramEvent() {

	a.ResetToolQueue()
	a.ResetToolQueuePosition()
	a.IsProgramRunning = false
}

func (a *AutoManager) HandleProgramStartEvent() {
	a.programStart = time.Now()
	a.IsProgramRunning = true

}

func (a *AutoManager) TimeSinceProgramStart() float64 {
	return time.Since(a.programStart).Seconds()
}

func (a *AutoManager) isProgramInIgnoreTime() bool {

	if !a.config.Viper.IsSet("automanager.programstartignoretime") {
		a.log.Error("error config key automanager.programstartignoretime not set")
	}

	ignoreTime := a.config.Viper.GetFloat64("automanager.programstartignoretime")

	return a.TimeSinceProgramStart() < ignoreTime

}

func (a *AutoManager) handleToolSenseEvent(isToolSenseHigh bool) {

	if !a.config.Viper.IsSet("automanager.istooladvanceonhigh") {
		a.log.Error("error finding config key automanager.istooladvanceonhig")
	}

	isToolAdvanceOnHigh := a.config.Viper.GetBool("automanager.istooladvanceonhigh")

	if isToolAdvanceOnHigh {

		if isToolSenseHigh {
			a.HandleNextToolEvent()
			a.log.Debug("tool advance event triggered on high change")
		}
	} else {
		if !isToolSenseHigh {
			a.HandleNextToolEvent()
			a.log.Debug("tool advance event triggered on low change")
		}
	}
}

func (a *AutoManager) HandleNextToolEvent() {

	if a.IsProgramRunning {

		if a.isProgramInIgnoreTime() {
			a.log.Debug("tool change ignored since program within ignore time during program start")
			return
		}
		a.AdvanceToolQueuePosition()
	} else {
		a.log.Debug("next tool event ignored since program not running")
	}
}
