package trainings

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
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

	dataParts, err := DataParts(datastring, 10, 3)

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

	t.TrainingType = dataParts[1]

	duration, err := time.ParseDuration(dataParts[2])
	if err != nil {
		errReturning = fmt.Errorf("duration parsing error '%s': %w", dataParts[2], err)
		return errReturning
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
	speed := spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration)

	if t.TrainingType == "Бег" {
		calorics, err = spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	} else if t.TrainingType == "Ходьба" {
		calorics, err = spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	} else {
		return "", errors.New("unknown type of training")
	}

	stringReturning := fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f", t.TrainingType, t.Duration.Hours(), distance, speed, calorics)
	return stringReturning, err
}

func DataParts(data string, minLength, numberParts int) ([]string, error) {
	if len(data) < minLength {
		return []string{}, errors.New("error in the incoming data string")
	}

	dataParts := strings.Split(data, ",")

	if len(dataParts) != numberParts {
		return []string{}, errors.New("error in the number of parts of the incoming string")
	}
	if len(dataParts[0]) == 0 {
		return []string{}, errors.New("error in the incoming number of steps")
	}
	if numberParts == 2 && len(dataParts[1]) == 0 {
		return []string{}, errors.New("error in incoming duration")
	}
	if numberParts == 3 && len(dataParts[1]) == 0 {
		return []string{}, errors.New("error in the incoming activity type")
	}
	if numberParts == 3 && len(dataParts[2]) == 0 {
		return []string{}, errors.New("error in incoming duration")
	}
	return dataParts, nil
}
