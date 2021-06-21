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

// Money struct was created so that it is possible to have a type that can be used for monetary purposes where precision is required.
type Money struct {
	decimal.Decimal
}

func NewFromString(value string) (Money, error) {
	d, err := decimal.NewFromString(value)
	if err != nil {
		return Money{}, err
	}
	return Money{d}, nil
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

// Div 'm2' Attribute that represents the new value for money.
//'round' Attribute that represents the type of rounding desired.
//'prec' Attribute that represents the desired decimal place limit.
//Function that has the responsibility to division the values and use the type of rounding that was informed for the result. Example: m / m2
func (m Money) Div(m2 Money, round rounding.RoundingMode, prec int) Money {
	money := Money{m.Decimal.Div(m2.Decimal)}
	return money.scale(prec, round)
}

// Sub 'm2' Attribute that represents the new value for money.
//'round' Attribute that represents the type of rounding desired.
//'prec' Attribute that represents the desired decimal place limit.
//Function that has the responsibility to subtraction the values and use the type of rounding that was informed for the result. Example: m - m2
func (m Money) Sub(m2 Money, round rounding.RoundingMode, prec int) Money {
	money := Money{m.Decimal.Sub(m2.Decimal)}
	return money.scale(prec, round)
}

// Add 'm2' Attribute that represents the new value for money.
//'round' Attribute that represents the type of rounding desired.
//'prec' Attribute that represents the desired decimal place limit.
//Function that has the responsibility to add the values and use the type of rounding that was informed for the result. Example: m + m2
func (m Money) Add(m2 Money, round rounding.RoundingMode, prec int) Money {
	money := Money{m.Decimal.Add(m2.Decimal)}
	return money.scale(prec, round)
}

// Mul 'm2' Attribute that represents the new value for money.
//'round' Attribute that represents the type of rounding desired.
//'prec' Attribute that represents the desired decimal place limit.
//Function that has the responsibility to multiply the values and use the type of rounding that was informed for the result. Example: m * m2
func (m Money) Mul(m2 Money, round rounding.RoundingMode, prec int) Money {
	money := Money{m.Decimal.Mul(m2.Decimal)}
	return money.scale(prec, round)
}

func (m Money) Pow(m2 Money, round rounding.RoundingMode, prec int) Money {
	money := Money{m.Decimal.Pow(m2.Decimal)}
	return money.scale(prec, round)
}

func (m Money) GreaterThan(m2 Money) bool {
	return m.Decimal.GreaterThan(m2.Decimal)
}

func (m Money) GreaterThanOrEqual(m2 Money) bool {
	return m.Decimal.GreaterThanOrEqual(m2.Decimal)
}

func (m Money) LessThan(m2 Money) bool {
	return m.Decimal.LessThan(m2.Decimal)
}

func (m Money) LessThanOrEqual(m2 Money) bool {
	return m.Decimal.LessThanOrEqual(m2.Decimal)
}

func (m Money) Equal(m2 Money) bool {
	return m.Decimal.Equal(m2.Decimal)
}

func (m Money) ToFloat64() float64 {
	result, _ := m.Decimal.Float64()
	return result
}

func (m Money) ToInt() int64 {
	return m.Decimal.IntPart()
}

func (m Money) String() string {
	return m.Decimal.String()
}

func (m Money) scale(prec int, roundingMode rounding.RoundingMode) Money {
	number, _ := new(big.Rat).SetString(m.String())
	rounding.Round(number, prec, roundingMode)
	f, _ := number.Float64()
	return NewFromFloat(f)
}
