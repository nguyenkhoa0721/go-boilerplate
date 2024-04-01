package util

import (
	"database/sql"
	"errors"
	"fmt"
	"math/big"
	"strconv"
	"unicode"
)

func CapitalizeFirstLetter(input string) string {
	if len(input) == 0 {
		return input
	}
	firstChar := unicode.ToUpper(rune(input[0]))
	result := string(firstChar) + input[1:]

	return result
}

func NullStringToString(val sql.NullString) string {
	if val.Valid {
		return val.String
	} else {
		return ""
	}
}

func StrToInt(val string) (int, error) {
	i, err := strconv.Atoi(val)
	if err != nil {
		return 0, errors.New("can not parse int")
	}

	return i, nil
}

func StrToInt64(val string) (int64, error) {
	i, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		return 0, errors.New("can not parse int")
	}

	return i, nil
}

func StrToUInt64(val string) (uint64, error) {
	i, err := strconv.ParseUint(val, 10, 64)
	if err != nil {
		return 0, errors.New("can not parse int")
	}

	return i, nil
}

func StrToInt32(val string) (int32, error) {
	i, err := strconv.ParseInt(val, 10, 32)
	if err != nil {
		return 0, errors.New("can not parse int")
	}

	return int32(i), nil
}

func StrToFloat64(val string) (float64, error) {
	f, err := strconv.ParseFloat(val, 64)
	if err != nil {
		return 0, errors.New("can not parse float")
	}

	return f, nil
}

func UInt64ToStr(val uint64) string {
	return strconv.FormatUint(val, 10)
}

func ParseBigFloat(value string) (*big.Float, error) {
	f := new(big.Float)
	f.SetPrec(236) //  IEEE 754 octuple-precision binary floating-point format: binary256
	f.SetMode(big.ToNearestEven)
	_, err := fmt.Sscan(value, f)
	return f, err
}

func ConvertEthToWei(value string, d int) (*big.Int, bool) {
	eth, err := ParseBigFloat(value)
	if err != nil {
		return nil, false
	}
	decimal := new(big.Int)
	decimal.Exp(big.NewInt(10), big.NewInt(int64(d)), nil)
	wei := new(big.Float)
	wei.Mul(eth, new(big.Float).SetInt(decimal))

	weiInt := new(big.Int)
	wei.Int(weiInt)
	return weiInt, true
}

func ConvertGWeiToWei(value string) (*big.Int, bool) {
	gwei, err := ParseBigFloat(value)
	if err != nil {
		return nil, false
	}
	wei := new(big.Float)
	wei.Mul(gwei, big.NewFloat(1000000000))

	weiInt := new(big.Int)
	wei.Int(weiInt)
	return weiInt, true
}

func ConvertWeiToEth(value string, decimal uint) (string, error) {
	wei, err := ParseBigFloat(value)
	if err != nil {
		return "", errors.New("can not parse float")
	}
	decimalBig := new(big.Int)
	decimalBig.Exp(big.NewInt(10), big.NewInt(int64(decimal)), nil)
	eth := new(big.Float)
	eth.Quo(wei, new(big.Float).SetInt(decimalBig))

	return fmt.Sprintf("%.2f", eth), nil
}

func TextOverflow(text string, length int) string {
	if len(text) > length*2 {
		return text[:length-1] + "..." + text[len(text)-length:]
	}

	return text
}
