package trainings

import (
	"errors"
	"fmt"
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

func (t *Training) Parse(dataString string) (err error) {
	// TODO: реализовать функцию
	var errReturning error
	var duration time.Duration

	dataParts, err := spentenergy.DataSplit(dataString)
	if err != nil {
		return fmt.Errorf("invalid incoming data string ('%s') '%v': %w: ", dataString, dataParts, err)
	}

	if len(dataParts) != 3 {
		return errors.New("incorrect number of data line parts, expected 3")
	}

	t.Steps, err = spentenergy.ParseSteps(dataParts[0])
	if err != nil {
		errReturning = fmt.Errorf("error parsing the number of steps '%s': %w: ", dataParts[0], err)
		return errReturning
	}

	t.TrainingType = dataParts[1]

	dur, err := time.ParseDuration(dataParts[2])
	if err != nil {
		errReturning = fmt.Errorf("duration parsing error '%s': %w: ", dataParts[2], err)
		return errReturning
	}
	duration = dur

	if duration <= 0 {
		errReturning = errors.New("incorrect duration: must be positive")
		return errReturning
	}
	t.Duration = duration

	return nil
}

func (t Training) ActionInfo() (string, error) {
	// TODO: реализовать функцию
	var err error
	distance := spentenergy.Distance(t.Steps, t.Height)
	if distance <= 0 {
		err = errors.New("invalid distance calculation result")
		return "", err
	}

	speed := spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration)
	if speed <= 0 {
		err = errors.New("invalid speed calculation result")
		return "", err
	}

	var calories float64 = 0
	switch t.TrainingType {
	case "Бег":
		calories, err = spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	case "Ходьба":
		calories, err = spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	default:
		return "", fmt.Errorf("unknown training type: %s", t.TrainingType)
	}
	if err != nil {
		return "", fmt.Errorf("calories calculation failed: %w", err)
	}

	stringReturning := fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", t.TrainingType, t.Duration.Hours(), distance, speed, calories)

	return stringReturning, nil
}
