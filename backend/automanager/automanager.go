package automanager

import (
	"fmt"
	"math"
	"time"

	log "github.com/sirupsen/logrus"
	config "mechied.com/servocoolant/config"
	"mechied.com/servocoolant/devicemanager"
)

type AutoManager struct {
	log                      *log.Logger
	config                   *config.Config
	devicemanager            *devicemanager.DeviceManager
	CurrentToolQueuePosition int
	ToolQueue                []int
	programStart             time.Time
	IsProgramRunning         bool
	IsEndRequested			bool
	autoToolsenseEventChan   chan bool
}

func GetAutoManager(log *log.Logger, config *config.Config, deviceManager *devicemanager.DeviceManager) *AutoManager {

	a := AutoManager{
		log:                      log,
		config:                   config,
		devicemanager:            deviceManager,
		CurrentToolQueuePosition: 0,
		ToolQueue:                make([]int,0),
		IsProgramRunning:         false,
		IsEndRequested:           false,
		autoToolsenseEventChan:   make(chan bool),
	}
	a.init()
	return &a
}

func (a *AutoManager) init() {

	// register the channel with devicemanager.
	// think about doing this channel thing less adhoc in the future
	a.devicemanager.AutoToolsenseEventChan = &a.autoToolsenseEventChan

	go a.monitorToolChange()

}

func (a *AutoManager) monitorToolChange() {
	for {
		isToolSenseHigh := <-a.autoToolsenseEventChan

		a.handleToolSenseEvent(isToolSenseHigh)
	}
}

func (a *AutoManager) CalculateAngleForToolLength(toolId int) (*int, error) {

	length, err := a.config.GetToolLength(toolId)

	if err != nil {
		a.log.Error(fmt.Sprintf("error locating tool %v length", toolId))
		return nil, err
	}

	//    offset standoff
	// ___________
	// ||
	// ||
	//  |
	//  |
	// 0 t

	quadrantOffsetDeg := a.config.Viper.GetFloat64("automanager.quadrantoffset")

	baseToolAngleDeg := a.config.Viper.GetFloat64("automanager.0offsetangle") - quadrantOffsetDeg

	offsetStandoff := a.config.Viper.GetFloat64("automanager.offsetstandoff")

	// tan(baseToolAngle) =  baseToolHeight/offsetStandoff
	baseToolHeight := math.Tan(baseToolAngleDeg*2*math.Pi/360) * offsetStandoff

	a.log.WithFields(log.Fields{
		"basetToolHeight":  baseToolHeight,
		"baseToolAngleDeg": baseToolAngleDeg,
	}).Debug("computed angle")
	actualLength := float64(*length) + baseToolHeight

	angleRad := math.Atan(actualLength / offsetStandoff)

	angleDeg := ((angleRad * 360) / (2 * math.Pi)) + quadrantOffsetDeg

	if angleDeg > quadrantOffsetDeg+90 || angleDeg < quadrantOffsetDeg {
		a.log.Warn(fmt.Sprintf("angle seems suspicious %v", angleDeg))
	}

	a.log.WithFields(log.Fields{
		"toolId":       toolId,
		"length":       *length,
		"actualLength": actualLength,
		"angleDeg":     int(angleDeg),
	}).Debug("computed angle")
	angleInt := int(angleDeg)
	return &angleInt, nil
}

func (a *AutoManager) GetCurrentTool() int {
	if len(a.ToolQueue) == 0 {
		return -1
	}
	if a.CurrentToolQueuePosition >= len(a.ToolQueue) {
		return -1
	}
	return a.ToolQueue[a.CurrentToolQueuePosition]
}

func (a *AutoManager) ResetToolQueuePosition() {
	a.CurrentToolQueuePosition = 0
	a.log.Debug("tool queue position reset")
}

func (a *AutoManager) ResetToolQueue() {
	a.ToolQueue = make([]int, 0)
	a.log.Debug("tool queue reset")

}

func (a *AutoManager) AdvanceToolQueuePosition() {
	a.CurrentToolQueuePosition += 1

	if a.CurrentToolQueuePosition >= len(a.ToolQueue) {
		a.CurrentToolQueuePosition = len(a.ToolQueue) - 1
		a.log.Error(fmt.Sprintf("tool queue position advanced beyond length of queue. remaining at %v", a.CurrentToolQueuePosition))
	}
	a.log.Debug(fmt.Sprintf("tool queue advanced to %v", a.CurrentToolQueuePosition))
	//activate the tool. ie. set the servos to the correct angle
	a.activateCurrentTool()
	a.CheckShouldProgramEnd()
}
