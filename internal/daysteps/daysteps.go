package daysteps

import (
	"errors"
	"fmt"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type DaySteps struct {
	// TODO: добавить поля
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

func (ds *DaySteps) Parse(datastring string) (err error) {
	// TODO: реализовать функцию
	var errReturning error

	dataParts, err := spentenergy.DataSplit(datastring)
	if err != nil {
		errReturning = fmt.Errorf("invalid incoming data string ('%s') '%v': %w", datastring, dataParts, err)
		return errReturning
	}

	if len(dataParts) != 2 {
		errReturning = errors.New("incorrect number of data line parts")
		return errReturning
	}

	ds.Steps, err = spentenergy.ParseSteps(dataParts[0])
	if err != nil {
		errReturning = fmt.Errorf("error parsing the number of steps '%s': %w: ", dataParts[0], err)
		return errReturning
	}

	duration, err := time.ParseDuration(dataParts[1])
	if err != nil {
		errReturning = fmt.Errorf("duration parsing error '%s': %w", dataParts[1], err)
		return errReturning
	}
	if duration <= 0 {
		errReturning = errors.New("incorrect duration")
		return errReturning
	}
	ds.Duration = duration

	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	// TODO: реализовать функцию
	var err error = nil
	var calories float64 = 0

	distance := spentenergy.Distance(ds.Steps, ds.Height)
	if distance <= 0 {
		err = errors.New("incorrect distance calculation result")
		return "", err
	}

	calories, err = spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	if err != nil {
		err = errors.New("incorrect result of caloric calculation")
		return "", err
	}
	stringReturning := fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", ds.Steps, distance, calories)

	return stringReturning, err
}
