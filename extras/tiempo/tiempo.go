package tiempo

import (
	"time"
)

func Tiempo() string {
	tiempo := time.Now()
	return determinarMomentoDelDia(tiempo.Hour())
}

func determinarMomentoDelDia(hora int) string {
	switch {
	case hora >= 6 && hora < 12:
		return ", buenos días!"
	case hora >= 12 && hora < 18:
		return ", buenas tardes!"
	default:
		return ", buenas noches!"
	}
}