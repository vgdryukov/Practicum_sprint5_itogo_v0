package spentenergy

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе.
)

func DataSplit(data string) ([]string, error) {
	dataParts := strings.Split(data, ",")

	if len(dataParts) < 2 || len(dataParts) > 3 {
		return []string{}, errors.New("incorrect number of parts of the incoming string: expected 2 or 3 parts")
	}
	if len(dataParts[0]) == 0 {
		return []string{}, errors.New("steps data cannot be empty")
	}
	if (len(dataParts) == 2 && len(dataParts[1]) == 0) || (len(dataParts) == 3 && len(dataParts[2]) == 0) {
		return []string{}, errors.New("duration cannot be empty")
	}
	if len(dataParts) == 3 && len(dataParts[1]) == 0 {
		return []string{}, errors.New("activity type cannot be empty")
	}

	return dataParts, nil
}

func ParseSteps(stepsStr string) (int, error) {
	steps, err := strconv.Atoi(stepsStr)
	if err != nil {
		return 0, fmt.Errorf("incorrect steps format '%s': %w: ", stepsStr, err)
	}
	if steps <= 0 {
		return 0, errors.New("steps must be positive")
	}

	return steps, nil
}

func CheckingData(nameFunc string, steps int, weight, height float64, duration time.Duration) error {
	type ValidationConfig struct {
		MinSteps    int
		MinWeight   float64
		MaxWeight   float64
		MinHeight   float64
		MaxHeight   float64
		MinDuration time.Duration
	}
	var DefaultValidation = ValidationConfig{
		MinSteps:    1,
		MinWeight:   5.0,   //кг.
		MaxWeight:   300.0, //кг.
		MinHeight:   1.0,   //м.
		MaxHeight:   2.5,   //м.
		MinDuration: time.Second,
	}
	if steps < DefaultValidation.MinSteps {
		return fmt.Errorf("%s: steps must be at least %d", nameFunc, DefaultValidation.MinSteps)
	}
	if weight < DefaultValidation.MinWeight || weight > DefaultValidation.MaxWeight {
		return fmt.Errorf("%s: weight must be at least %.1f ang less then %.1f kg.", nameFunc, DefaultValidation.MinWeight, DefaultValidation.MaxWeight)
	}
	if height < DefaultValidation.MinHeight || height > DefaultValidation.MaxHeight {
		return fmt.Errorf("%s: height must be at least %.1f ang less then %.1f m.", nameFunc, DefaultValidation.MinHeight, DefaultValidation.MaxHeight)
	}
	if duration < DefaultValidation.MinDuration {
		return fmt.Errorf("%s: duration must be positive", nameFunc)
	}

	return nil
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if err := CheckingData("WalkingSpentCalories", steps, weight, height, duration); err != nil {
		return 0, err
	}
	spentCalories := walkingCaloriesCoefficient * (weight * MeanSpeed(steps, height, duration) * duration.Minutes()) / minInH

	return spentCalories, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if err := CheckingData("RunningSpentCalories", steps, weight, height, duration); err != nil {
		return 0, err
	}
	spentCalories := (weight * MeanSpeed(steps, height, duration) * duration.Minutes()) / minInH

	return spentCalories, nil
}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	// TODO: реализовать функцию
	if steps <= 0 || float64(duration) <= 0 {
		return 0
	}
	averageSpeed := Distance(steps, height) / float64(duration.Hours())

	return averageSpeed
}

func Distance(steps int, height float64) float64 {
	// TODO: реализовать функцию
	distanceInKm := (float64(steps) * height * stepLengthCoefficient) / mInKm
	if distanceInKm <= 0 {
		return 0
	}

	return distanceInKm
}
