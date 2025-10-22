package daysteps

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
	"github.com/Yandex-Practicum/tracker/internal/trainings"
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

	dataParts, err := trainings.DataParts(datastring, 3, 2)

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
	ds.Steps = num

	duration, err := time.ParseDuration(dataParts[1])
	if err != nil {
		errReturning = fmt.Errorf("duration parsing error '%s': %w", dataParts[2], err)
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
	var calorics float64 = 0
	distance := spentenergy.Distance(ds.Steps, ds.Height)
	calorics, err = spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)

	stringReturning := fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", ds.Steps, distance, calorics)
	return stringReturning, err
}
