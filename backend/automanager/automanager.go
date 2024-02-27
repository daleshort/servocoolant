package automanager

import (
	"fmt"
	"math"

	log "github.com/sirupsen/logrus"
	config "mechied.com/servocoolant/config"
	"mechied.com/servocoolant/devicemanager"
)

type AutoManager struct {
	log           *log.Logger
	config        *config.Config
	devicemanager *devicemanager.DeviceManager
}

func GetAutoManager(log *log.Logger, config *config.Config, deviceManager *devicemanager.DeviceManager) *AutoManager {

	a := AutoManager{
		log:           log,
		config:        config,
		devicemanager: deviceManager,
	}

	return &a
}

func (a *AutoManager) CalculateAngleForToolLength(toolId int) int {

	length, _ := a.config.GetToolLength(toolId)

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
	return int(angleDeg)
}
