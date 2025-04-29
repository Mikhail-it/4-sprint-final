package spentcalories

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	lenStep                    = 0.65 // средняя длина шага.
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе
)

func parseTraining(data string) (int, string, time.Duration, error) {
	// TODO: реализовать функцию
	parts := strings.Split(data, ",")
	if len(parts) != 3 {
		return 0, "", 0, errors.New("некорректный формат строки")
	}

	steps, err := strconv.Atoi(strings.TrimSpace(parts[0]))
	if err != nil || steps <= 0 {
		return 0, "", 0, errors.New("некорректное количество шагов")
	}

	activityType := strings.TrimSpace(parts[1])
	if activityType == "" {
		return 0, "", 0, fmt.Errorf("вид активности не должен быть пустым")
	}

	duration, err := time.ParseDuration(strings.TrimSpace(parts[2]))
	if err != nil || duration <= 0 {
		return 0, "", 0, errors.New("некорректная продолжительность")
	}

	return steps, activityType, duration, nil
}

func distance(steps int, height float64) float64 {
	// TODO: реализовать функцию
	stepLength := height * stepLengthCoefficient
	return float64(steps) * stepLength / mInKm
}

func meanSpeed(steps int, height float64, duration time.Duration) float64 {
	// TODO: реализовать функцию
	if duration <= 0 {
		return 0
	}
	dist := distance(steps, height)
	return dist / duration.Hours()
}

func TrainingInfo(data string, weight, height float64) (string, error) {
	// TODO: реализовать функцию
	steps, activity, duration, err := parseTraining(data)
	if err != nil {
		return "", err
	}
	dist := distance(steps, height)
	speed := meanSpeed(steps, height, duration)
	var calories float64
	switch activity {
	case "Ходьба":
		calories, err = WalkingSpentCalories(steps, weight, height, duration)
	case "Бег":
		calories, err = RunningSpentCalories(steps, weight, height, duration)
	default:
		return "", fmt.Errorf("неизвестный тип тренировки: %s", activity)
	}
	if err != nil {
		return "", err
	}
	result := fmt.Sprintf(
		"Вы занимались: %s\nКоличество шагов: %d\nДистанция: %.2f км\nСредняя скорость: %.2f км/ч\nСожжено калорий: %.2f ккал\n",
		activity, steps, dist, speed, calories,
	)
	return result, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, fmt.Errorf("некорректные входные параметры")
	}
	speed := meanSpeed(steps, height, duration)
	minutes := duration.Minutes()
	calories := (weight * speed * minutes) / minInH
	return calories, nil
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, fmt.Errorf("некорректные входные параметры")
	}
	speed := meanSpeed(steps, height, duration)
	minutes := duration.Minutes()
	calories := ((weight * speed * minutes) / minInH) * walkingCaloriesCoefficient
	return calories, nil
}
