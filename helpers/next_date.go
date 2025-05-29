package helpers

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Antonpractic/go_final_project/constants"
)

func NextDate(now time.Time, date string, repeat string) (string, error) {
	parsedDate, err := time.Parse(constants.FormatDate, date)
	if err != nil {
		return "", fmt.Errorf("неверный формат даты: %v", err)
	}

	parts := strings.Split(repeat, " ")

	switch parts[0] {
	case "y":
		if len(parts) > 1 {
			return "", fmt.Errorf("неподдерживаемый формат повтора")
		}
		parsedDate = parsedDate.AddDate(1, 0, 0)
		for !parsedDate.After(now) {
			parsedDate = parsedDate.AddDate(1, 0, 0)
		}
	case "d":
		if len(parts) > 1 {
			days, err := strconv.Atoi(parts[1])
			if err != nil || days <= 0 || days > 400 {
				return "", fmt.Errorf("неправильный формат для дней: %v", err)
			}

			parsedDate = parsedDate.AddDate(0, 0, days)
			for !parsedDate.After(now) {
				parsedDate = parsedDate.AddDate(0, 0, days)
			}
		} else {
			return "", fmt.Errorf("количество дней не указано")
		}
	case "m", "w":
		return "", fmt.Errorf("неподдерживаемый формат повтора")
	default:
		return "", fmt.Errorf("неподдерживаемый формат повтора")
	}

	return parsedDate.Format(constants.FormatDate), nil
}
