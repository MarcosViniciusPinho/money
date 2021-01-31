package money

import (
	"math/big"

	"github.com/shopspring/decimal"
	"github.com/wadey/go-rounding"
)

/**
* File that has the responsibility of defining a new data type to work with high precision floating point values.
* Objective: The idea is to encapsulate the decimal and go-rounding dependencies here, to be used together in order to achieve the desired.
* PS: It has the same effect as a BigDecimal in Java for example.
 */

var OneHundred = NewFromInt(100)
var Ten = NewFromInt(10)
var One = NewFromInt(1)
var Zero = NewFromInt(0)

//struct was created so that it is possible to have a type that can be used for monetary purposes where precision is required.
type Money struct {
	decimal.Decimal
}

func NewFromString(value string) (Money, error) {
	m, err := decimal.NewFromString(value)
	if err != nil {
		return Money{}, err
	}
	return Money{m}, nil
}

func NewFromFloat(value float64) Money {
	return Money{decimal.NewFromFloat(value)}
}

func NewFromInt(value int) Money {
	return Money{decimal.NewFromInt(int64(value))}
}

/**
 * NOTE ON ROUNDING TYPES:
 * rounding.Up -> Always adds rounding regardless of the result
 * rounding.Down -> Always ignores rounding regardless of result
 * rounding.Ceil -> It will always add rounding if the result is greater than zero; that is; the result needs to be positive
 * rounding.Floor -> It will always add rounding if the result is less than zero; that is; the result must be negative
 * rounding.HalfUp -> It will always add rounding when the decimal place is greater than or equal to five (5)
 * rounding.HalfDown -> It will always add rounding when the decimal place is greater than to five (5)
 * rounding.HalfEven -> It will always add rounding when the decimal place is greater than to five (5)
 */

//'d2' Attribute that represents the new value for money.
//'round' Attribute that represents the type of rounding desired.
//'prec' Attribute that represents the desired decimal place limit.
//Function that has the responsibility to division the values ​​and use the type of rounding that was informed for the result. Example: d / d2
func (d Money) Div(d2 Money, round rounding.RoundingMode, prec int) Money {
	money := Money{d.Decimal.Div(d2.Decimal)}
	return money.scale(prec, round)
}

//'d2' Attribute that represents the new value for money.
//'round' Attribute that represents the type of rounding desired.
//'prec' Attribute that represents the desired decimal place limit.
//Function that has the responsibility to subtraction the values ​​and use the type of rounding that was informed for the result. Example: d - d2
func (d Money) Sub(d2 Money, round rounding.RoundingMode, prec int) Money {
	money := Money{d.Decimal.Sub(d2.Decimal)}
	return money.scale(prec, round)
}

//'d2' Attribute that represents the new value for money.
//'round' Attribute that represents the type of rounding desired.
//'prec' Attribute that represents the desired decimal place limit.
//Function that has the responsibility to add the values ​​and use the type of rounding that was informed for the result. Example: d + d2
func (d Money) Add(d2 Money, round rounding.RoundingMode, prec int) Money {
	money := Money{d.Decimal.Add(d2.Decimal)}
	return money.scale(prec, round)
}

//'d2' Attribute that represents the new value for money.
//'round' Attribute that represents the type of rounding desired.
//'prec' Attribute that represents the desired decimal place limit.
//Function that has the responsibility to multiply the values ​​and use the type of rounding that was informed for the result. Example: d * d2
func (d Money) Mul(d2 Money, round rounding.RoundingMode, prec int) Money {
	money := Money{d.Decimal.Mul(d2.Decimal)}
	return money.scale(prec, round)
}

func (d Money) Pow(d2 Money, round rounding.RoundingMode, prec int) Money {
	money := Money{d.Decimal.Pow(d2.Decimal)}
	return money.scale(prec, round)
}

func (d Money) GreaterThan(d2 Money) bool {
	return d.Decimal.GreaterThan(d2.Decimal)
}

func (d Money) GreaterThanOrEqual(d2 Money) bool {
	return d.Decimal.GreaterThanOrEqual(d2.Decimal)
}

func (d Money) LessThan(d2 Money) bool {
	return d.Decimal.LessThan(d2.Decimal)
}

func (d Money) LessThanOrEqual(d2 Money) bool {
	return d.Decimal.LessThanOrEqual(d2.Decimal)
}

func (d Money) Equal(d2 Money) bool {
	return d.Decimal.Equal(d2.Decimal)
}

func (d Money) ToFloat64() float64 {
	result, _ := d.Decimal.Float64()
	return result
}

func (d Money) ToInt() int64 {
	return d.Decimal.IntPart()
}

func (d Money) String() string {
	return d.Decimal.String()
}

func (d Money) scale(prec int, roundingMode rounding.RoundingMode) Money {
	number, _ := new(big.Rat).SetString(d.String())
	rounding.Round(number, prec, roundingMode)
	f, _ := number.Float64()
	return NewFromFloat(f)
}
