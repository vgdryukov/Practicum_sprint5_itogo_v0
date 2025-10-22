package trainings

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type Training struct {
	// TODO: добавить поля
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

func (t *Training) Parse(datastring string) (err error) {
	// TODO: реализовать функцию
	var errReturning error
	var duration time.Duration

	dataParts, err := spentenergy.DataSplit(datastring)

	if err != nil {
		errReturning = fmt.Errorf("invalid incoming data string ('%s') '%v': %w", datastring, dataParts, err)
		return errReturning
	}

	num, err := strconv.Atoi(dataParts[0])
	if err != nil {
		errReturning = fmt.Errorf("incorrect input of the number of steps '%s': %w", dataParts[0], err)
		return errReturning
	}
	if num <= 0 {
		errReturning = errors.New("incorrect number of steps")
		return errReturning
	}
	t.Steps = num

	if len(dataParts) == 3 {
		t.TrainingType = dataParts[1]

		dur, err := time.ParseDuration(dataParts[2])
		if err != nil {
			errReturning = fmt.Errorf("duration parsing error '%s': %w", dataParts[2], err)
			return errReturning
		}
		duration = dur
	}

	if len(dataParts) == 2 {
		dur, err := time.ParseDuration(dataParts[1])
		if err != nil {
			errReturning = fmt.Errorf("duration parsing error '%s': %w", dataParts[2], err)
			return errReturning
		}
		duration = dur
	}

	if duration <= 0 {
		errReturning = errors.New("incorrect duration")
		return errReturning
	}
	t.Duration = duration

	return nil
}

func (t Training) ActionInfo() (string, error) {
	// TODO: реализовать функцию
	var err error = nil
	var calorics float64 = 0

	distance := spentenergy.Distance(t.Steps, t.Height)
	if distance <= 0 {
		err = errors.New("incorrect distance calculation result")
		return "", err
	}

	speed := spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration)
	if speed <= 0 {
		err = errors.New("incorrect speed calculation result")
		return "", err
	}

	if t.TrainingType == "Бег" {
		calorics, err = spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	} else if t.TrainingType == "Ходьба" {
		calorics, err = spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	} else {
		return "", errors.New("unknown type of training")
	}

	stringReturning := fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", t.TrainingType, t.Duration.Hours(), distance, speed, calorics)
	return stringReturning, err
}
