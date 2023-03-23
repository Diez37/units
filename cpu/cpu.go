package cpu

import (
	"fmt"
	"github.com/spf13/cast"
	"math"
	"strings"
)

const (
	milliSuffix = "m"

	MilliCore = 1
	Core      = MilliCore * 1000
)

func Parse(cpuSecond string) (uint32, error) {
	if strings.HasSuffix(cpuSecond, milliSuffix) {
		return ParseMilli(cpuSecond)
	}

	return ParseCore(cpuSecond)
}

func ParseMilli(cpuSecond string) (uint32, error) {
	seconds, err := cast.ToFloat64E(strings.TrimSuffix(cpuSecond, milliSuffix))
	if err != nil {
		return 0, fmt.Errorf("cpu: parse millicores failed: %w", err)
	}

	if _, float := math.Modf(seconds); float > 0 {
		return 0, fmt.Errorf("cpu: fractional parts are not allowed when specifying millicores")
	}

	return cast.ToUint32E(seconds)
}

func ParseCore(cpuSecond string) (uint32, error) {
	seconds, err := cast.ToFloat32E(cpuSecond)
	if err != nil {
		return 0, err
	}

	return ToSeconds(seconds), nil
}

func ToSeconds(cpu float32) uint32 {
	return uint32(cpu * float32(Core))
}

func ToCpu(seconds uint32) float32 {
	return float32(seconds) / Core
}

func Ceil(cpu float32) uint8 {
	cores, float := math.Modf(float64(cpu))

	if float == 0 {
		return uint8(cores)
	}

	return uint8(cores + 1)
}
