package spentenergy

import (
	"errors"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе.
)

func CheckingData(nameFunc string, steps int, weight, height float64, duration time.Duration) error {
	var err error
	if steps <= 0 {
		err = errors.New("error in " + nameFunc + " : incoming steps are <= 0")
		return err
	}
	if weight < 5 {
		err = errors.New("error in " + nameFunc + " : incoming weight are less 5kg")
		return err
	}
	if height < 1 {
		err = errors.New("error in " + nameFunc + " : incoming height are less 1m")
		return err
	}
	if duration <= 0 {
		err = errors.New("error in " + nameFunc + " : incoming duration are <= 0")
		return err
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
