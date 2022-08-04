package utils

import (
	"fmt"
	"math"
	"math/big"
	"strconv"
	"time"

	"github.com/pkg/errors"
)

func TimeToString(time time.Time) string {
	timeStr := fmt.Sprintf("%s %d, %d, %s",
		time.Month().String()[:3],
		time.Day(),
		time.Year(),
		time.Format("15:04 PM"),
	)
	return timeStr
}

func StringToTime(timestamp string) (t time.Time, err error) {
	i, err := strconv.ParseInt(timestamp, 10, 64)
	if err != nil {
		err = errors.Wrap(err, "formatTime failed")
		return
	}
	t = time.Unix(i, 0)
	return
}

func BigIntToFloat64(amount *big.Int, decimals int) float64 {
	decimalsFloat := big.NewFloat(math.Pow10(decimals))
	amountFloat := big.NewFloat(0).SetInt(amount)
	bigFloat := big.NewFloat(0).Quo(amountFloat, decimalsFloat)
	float, _ := bigFloat.Float64()
	return float
}
