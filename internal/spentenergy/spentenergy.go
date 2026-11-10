package spentenergy

import (
	"fmt"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе.
)

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 {
		return 0, fmt.Errorf("количество шагов должно быть положительным")
	}

	if weight <= 0 {
		return 0, fmt.Errorf("вес должен быть положительным")
	}

	if height <= 0 {
		return 0, fmt.Errorf("рост должен быть положительным")
	}

	if duration <= 0 {
		return 0, fmt.Errorf("продолжительность должна быть положительной")
	}

	meanSpeed := MeanSpeed(steps, height, duration)
	durationInMinutes := duration.Minutes()

	calories := weight * meanSpeed * durationInMinutes / minInH
	calories *= walkingCaloriesCoefficient

	return calories, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 {
		return 0, fmt.Errorf("количество шагов должно быть положительным")
	}

	if weight <= 0 {
		return 0, fmt.Errorf("вес должен быть положительным")
	}

	if height <= 0 {
		return 0, fmt.Errorf("рост должен быть положительным")
	}

	if duration <= 0 {
		return 0, fmt.Errorf("продолжительность должна быть положительной")
	}

	meanSpeed := MeanSpeed(steps, height, duration)
	durationInMinutes := duration.Minutes()

	calories := weight * meanSpeed * durationInMinutes / minInH

	return calories, nil
}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	if duration <= 0 || steps < 0 {
		return 0
	}

	distance := Distance(steps, height)
	meanSpeed := distance / duration.Hours()

	return meanSpeed
}

func Distance(steps int, height float64) float64 {
	stepLength := height * stepLengthCoefficient
	distance := float64(steps) * stepLength / mInKm

	return distance
}
