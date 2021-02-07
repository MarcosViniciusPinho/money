package money

import (
	"github.com/wadey/go-rounding"
	"testing"
)

func TestMustAcceptAllTypesOfRoundingPossibleToSum(t *testing.T) {
	tests := []Test{
		{input: NewFromFloat(-515.4554).Add(NewFromFloat(-200.9876), rounding.Floor, 2), expected: NewFromFloat(-716.45),
			message: "Arredondamento do resultado final, pois o mesmo é um número menor que zero(0)"},
		{input: fromString("-515.4554").Add(fromString("-200.9876"), rounding.Floor, 2), expected: fromString("-716.45"),
			message: "Arredondamento do resultado final, pois o mesmo é um número menor que zero(0)"},

		{input: NewFromFloat(515.4554).Add(NewFromFloat(200.9876), rounding.Ceil, 2), expected: NewFromFloat(716.45),
			message: "Arredondamento do resultado final, pois o mesmo é um número maior que zero(0)"},
		{input: fromString("515.4554").Add(fromString("200.9876"), rounding.Ceil, 2), expected: fromString("716.45"),
			message: "Arredondamento do resultado final, pois o mesmo é um número maior que zero(0)"},
		{input: NewFromFloat(515.50001).Add(NewFromFloat(200.510501), rounding.Ceil, 2), expected: NewFromFloat(716.02),
			message: "Arredondamento do resultado final, pois o mesmo é um número igual à zero(0)"},
		{input: fromString("515.50001").Add(fromString("200.510501"), rounding.Ceil, 2), expected: fromString("716.02"),
			message: "Arredondamento do resultado final, pois o mesmo é um número igual à zero(0)"},

		{input: NewFromFloat(515.500).Add(NewFromFloat(200.9876), rounding.HalfDown, 2), expected: NewFromFloat(716.49),
			message: "Arredondamento do resultado final, pois o mesmo é um número maior que cinco(5)"},
		{input: fromString("515.500").Add(fromString("200.9876"), rounding.HalfDown, 2), expected: fromString("716.49"),
			message: "Arredondamento do resultado final, pois o mesmo é um número maior que cinco(5)"},

		{input: NewFromFloat(515.500).Add(NewFromFloat(200.9876), rounding.HalfEven, 2), expected: NewFromFloat(716.49),
			message: "Arredondamento do resultado final, pois o mesmo é um número maior que cinco(5)"},
		{input: fromString("515.500").Add(fromString("200.9876"), rounding.HalfEven, 2), expected: fromString("716.49"),
			message: "Arredondamento do resultado final, pois o mesmo é um número maior que cinco(5)"},

		{input: NewFromFloat(515.500).Add(NewFromFloat(200.9876), rounding.HalfUp, 2), expected: NewFromFloat(716.49),
			message: "Arredondamento do resultado final, pois o mesmo é um número maior que cinco(5)"},
		{input: fromString("515.500").Add(fromString("200.9876"), rounding.HalfUp, 2), expected: fromString("716.49"),
			message: "Arredondamento do resultado final, pois o mesmo é um número maior que cinco(5)"},
		{input: NewFromFloat(515.590).Add(NewFromFloat(200.555), rounding.HalfUp, 2), expected: NewFromFloat(716.15),
			message: "Arredondamento do resultado final, pois o mesmo é um número igual à cinco(5)"},
		{input: fromString("515.590").Add(fromString("200.555"), rounding.HalfUp, 2), expected: fromString("716.15"),
			message: "Arredondamento do resultado final, pois o mesmo é um número igual à cinco(5)"},

		{input: NewFromFloat(-515.4554).Add(NewFromFloat(-200.9876), rounding.Up, 2), expected: NewFromFloat(-716.45),
			message: "Arredondamento do resultado final quando numero for negativo"},
		{input: fromString("-515.4554").Add(fromString("-200.9876"), rounding.Up, 2), expected: fromString("-716.45"),
			message: "Arredondamento do resultado final quando numero for negativo"},
		{input: NewFromFloat(515.500).Add(NewFromFloat(200.9876), rounding.Up, 2), expected: NewFromFloat(716.49),
			message: "Arredondamento do resultado final quando numero for positivo e maior que cinco(5)"},
		{input: fromString("515.500").Add(fromString("200.9876"), rounding.Up, 2), expected: fromString("716.49"),
			message: "Arredondamento do resultado final quando numero for positivo e maior que cinco(5)"},
		{input: NewFromFloat(515.590).Add(NewFromFloat(200.555), rounding.Up, 2), expected: NewFromFloat(716.15),
			message: "Arredondamento do resultado final quando numero for positivo igual à cinco(5)"},
		{input: fromString("515.590").Add(fromString("200.555"), rounding.Up, 2), expected: fromString("716.15"),
			message: "Arredondamento do resultado final quando numero for positivo igual à cinco(5)"},
		{input: NewFromFloat(515.50001).Add(NewFromFloat(200.510501), rounding.Up, 2), expected: NewFromFloat(716.02),
			message: "Arredondamento do resultado final, pois o mesmo é um número igual à zero(0)"},
		{input: fromString("515.50001").Add(fromString("200.510501"), rounding.Up, 2), expected: fromString("716.02"),
			message: "Arredondamento do resultado final, pois o mesmo é um número igual à zero(0)"},
	}
	for _, test := range tests {
		assertEquals(t, test.expected.ToFloat64(), test.input.ToFloat64())
	}
}

func TestNotMustAcceptRoundingToSum(t *testing.T) {
	tests := []Test{
		{input: NewFromFloat(515.4554).Add(NewFromFloat(200.9876), rounding.Floor, 2), expected: NewFromFloat(716.44),
			message: "Não deve existir arredondamento do resultado final, pois o mesmo não é um número menor que zero(0)"},
		{input: fromString("515.4554").Add(fromString("200.9876"), rounding.Floor, 2), expected: fromString("716.44"),
			message: "Não deve existir arredondamento do resultado final, pois o mesmo não é um número menor que zero(0)"},

		{input: NewFromFloat(-515.4554).Add(NewFromFloat(-200.9876), rounding.Ceil, 2), expected: NewFromFloat(-716.44),
			message: "Não deve existir arredondamento do resultado final, pois o mesmo não é um número maior que zero(0) nem igual"},
		{input: fromString("-515.4554").Add(fromString("-200.9876"), rounding.Ceil, 2), expected: fromString("-716.44"),
			message: "Não deve existir arredondamento do resultado final, pois o mesmo não é um número maior que zero(0) nem igual"},

		{input: NewFromFloat(515.50001).Add(NewFromFloat(200.510501), rounding.HalfDown, 2), expected: NewFromFloat(716.01),
			message: "Não faz arredondamento do resultado final, pois o mesmo não é um número maior que cinco(5)"},
		{input: fromString("515.50001").Add(fromString("200.510501"), rounding.HalfDown, 2), expected: fromString("716.01"),
			message: "Não faz arredondamento do resultado final, pois o mesmo não é um número maior que cinco(5)"},

		{input: NewFromFloat(515.50001).Add(NewFromFloat(200.510501), rounding.HalfEven, 2), expected: NewFromFloat(716.01),
			message: "Não faz arredondamento do resultado final, pois o mesmo não é um número maior que cinco(5)"},
		{input: fromString("515.50001").Add(fromString("200.510501"), rounding.HalfEven, 2), expected: fromString("716.01"),
			message: "Não faz arredondamento do resultado final, pois o mesmo não é um número maior que cinco(5)"},

		{input: NewFromFloat(515.50001).Add(NewFromFloat(200.510501), rounding.HalfUp, 2), expected: NewFromFloat(716.01),
			message: "Não faz arredondamento do resultado final, pois o mesmo não é um número maior que cinco(5)"},
		{input: fromString("515.50001").Add(fromString("200.510501"), rounding.HalfUp, 2), expected: fromString("716.01"),
			message: "Não faz arredondamento do resultado final, pois o mesmo não é um número maior que cinco(5)"},
		{input: NewFromFloat(515.50001).Add(NewFromFloat(200.510501), rounding.HalfUp, 2), expected: NewFromFloat(716.01),
			message: "Não faz arredondamento do resultado final, pois o mesmo não é um número igual à cinco(5)"},
		{input: fromString("515.50001").Add(fromString("200.510501"), rounding.HalfUp, 2), expected: fromString("716.01"),
			message: "Não faz arredondamento do resultado final, pois o mesmo não é um número igual à cinco(5)"},

		{input: NewFromFloat(-515.4554).Add(NewFromFloat(-200.9876), rounding.Down, 2), expected: NewFromFloat(-716.44),
			message: "Não faz arredondamento do resultado final quando numero for negativo"},
		{input: fromString("-515.4554").Add(fromString("-200.9876"), rounding.Down, 2), expected: fromString("-716.44"),
			message: "Não faz arredondamento do resultado final quando numero for negativo"},
		{input: NewFromFloat(515.500).Add(NewFromFloat(200.9876), rounding.Down, 2), expected: NewFromFloat(716.48),
			message: "Não faz arredondamento do resultado final quando numero for positivo e maior que cinco(5)"},
		{input: fromString("515.500").Add(fromString("200.9876"), rounding.Down, 2), expected: fromString("716.48"),
			message: "Não faz arredondamento do resultado final quando numero for positivo e maior que cinco(5)"},
		{input: NewFromFloat(515.590).Add(NewFromFloat(200.555), rounding.Down, 2), expected: NewFromFloat(716.14),
			message: "Não faz arredondamento do resultado final quando numero for positivo igual à cinco(5)"},
		{input: fromString("515.590").Add(fromString("200.555"), rounding.Down, 2), expected: fromString("716.14"),
			message: "Não faz arredondamento do resultado final quando numero for positivo igual à cinco(5)"},
		{input: NewFromFloat(515.50001).Add(NewFromFloat(200.510501), rounding.Down, 2), expected: NewFromFloat(716.01),
			message: "Não faz arredondamento do resultado final, pois o mesmo é um número igual à zero(0)"},
		{input: fromString("515.50001").Add(fromString("200.510501"), rounding.Down, 2), expected: fromString("716.01"),
			message: "Não faz arredondamento do resultado final, pois o mesmo é um número igual à zero(0)"},
		{input: NewFromInt(515).Add(NewFromInt(200), rounding.Down, 0), expected: NewFromInt(715),
			message: "Não existe arredondamento para numeros inteiros"},
	}
	for _, test := range tests {
		assertEquals(t, test.expected.ToFloat64(), test.input.ToFloat64())
	}
}

func TestMustAcceptAllTypesOfRoundingPossibleToSubtraction(t *testing.T) {
	tests := []Test{
		{input: NewFromFloat(-515.4554).Sub(NewFromFloat(-200.9876), rounding.Floor, 2), expected: NewFromFloat(-314.47),
			message: "Arredondamento do resultado final, pois o mesmo é um número menor que zero(0)"},
		{input: fromString("-515.4554").Sub(fromString("-200.9876"), rounding.Floor, 2), expected: fromString("-314.47"),
			message: "Arredondamento do resultado final, pois o mesmo é um número menor que zero(0)"},

		{input: NewFromFloat(515.4554).Sub(NewFromFloat(200.9876), rounding.Ceil, 2), expected: NewFromFloat(314.47),
			message: "Arredondamento do resultado final, pois o mesmo é um número maior que zero(0)"},
		{input: fromString("515.4554").Sub(fromString("200.9876"), rounding.Ceil, 2), expected: fromString("314.47"),
			message: "Arredondamento do resultado final, pois o mesmo é um número maior que zero(0)"},
		{input: NewFromFloat(515.2021).Sub(NewFromFloat(200.522), rounding.Ceil, 2), expected: NewFromFloat(314.69),
			message: "Arredondamento do resultado final, pois o mesmo é um número igual à zero(0)"},
		{input: fromString("515.2021").Sub(fromString("200.522"), rounding.Ceil, 2), expected: fromString("314.69"),
			message: "Arredondamento do resultado final, pois o mesmo é um número igual à zero(0)"},

		{input: NewFromFloat(515.500).Sub(NewFromFloat(200.7222), rounding.HalfDown, 2), expected: NewFromFloat(314.78),
			message: "Arredondamento do resultado final, pois o mesmo é um número maior que cinco(5)"},
		{input: fromString("515.500").Sub(fromString("200.7222"), rounding.HalfDown, 2), expected: fromString("314.78"),
			message: "Arredondamento do resultado final, pois o mesmo é um número maior que cinco(5)"},

		{input: NewFromFloat(515.500).Sub(NewFromFloat(200.7222), rounding.HalfEven, 2), expected: NewFromFloat(314.78),
			message: "Arredondamento do resultado final, pois o mesmo é um número maior que cinco(5)"},
		{input: fromString("515.500").Sub(fromString("200.7222"), rounding.HalfEven, 2), expected: fromString("314.78"),
			message: "Arredondamento do resultado final, pois o mesmo é um número maior que cinco(5)"},

		{input: NewFromFloat(515.500).Sub(NewFromFloat(200.7222), rounding.HalfUp, 2), expected: NewFromFloat(314.78),
			message: "Arredondamento do resultado final, pois o mesmo é um número maior que cinco(5)"},
		{input: fromString("515.500").Sub(fromString("200.7222"), rounding.HalfUp, 2), expected: fromString("314.78"),
			message: "Arredondamento do resultado final, pois o mesmo é um número maior que cinco(5)"},
		{input: NewFromFloat(515.500).Sub(NewFromFloat(200.725), rounding.HalfUp, 2), expected: NewFromFloat(314.78),
			message: "Arredondamento do resultado final, pois o mesmo é um número igual à cinco(5)"},
		{input: fromString("515.500").Sub(fromString("200.725"), rounding.HalfUp, 2), expected: fromString("314.78"),
			message: "Arredondamento do resultado final, pois o mesmo é um número igual à cinco(5)"},

		{input: NewFromFloat(-515.4554).Sub(NewFromFloat(-200.9876), rounding.Up, 2), expected: NewFromFloat(-314.47),
			message: "Arredondamento do resultado final quando numero for negativo"},
		{input: fromString("-515.4554").Sub(fromString("-200.9876"), rounding.Up, 2), expected: fromString("-314.47"),
			message: "Arredondamento do resultado final quando numero for negativo"},
		{input: NewFromFloat(515.500).Sub(NewFromFloat(200.7222), rounding.Up, 2), expected: NewFromFloat(314.78),
			message: "Arredondamento do resultado final quando numero for positivo e maior que cinco(5)"},
		{input: fromString("515.500").Sub(fromString("200.7222"), rounding.Up, 2), expected: fromString("314.78"),
			message: "Arredondamento do resultado final quando numero for positivo e maior que cinco(5)"},
		{input: NewFromFloat(515.500).Sub(NewFromFloat(200.725), rounding.Up, 2), expected: NewFromFloat(314.78),
			message: "Arredondamento do resultado final quando numero for positivo igual à cinco(5)"},
		{input: fromString("515.500").Sub(fromString("200.725"), rounding.Up, 2), expected: fromString("314.78"),
			message: "Arredondamento do resultado final quando numero for positivo igual à cinco(5)"},
		{input: NewFromFloat(515.50001).Sub(NewFromFloat(200.510501), rounding.Up, 2), expected: NewFromFloat(314.99),
			message: "Arredondamento do resultado final, pois o mesmo é um número igual à zero(0)"},
		{input: fromString("515.50001").Sub(fromString("200.510501"), rounding.Up, 2), expected: fromString("314.99"),
			message: "Arredondamento do resultado final, pois o mesmo é um número igual à zero(0)"},
	}
	for _, test := range tests {
		assertEquals(t, test.expected.ToFloat64(), test.input.ToFloat64())
	}
}

func TestNotMustAcceptRoundingToSubtraction(t *testing.T) {
	tests := []Test{
		{input: NewFromFloat(515.4554).Sub(NewFromFloat(200.9876), rounding.Floor, 2), expected: NewFromFloat(314.46),
			message: "Não deve existir arredondamento do resultado final, pois o mesmo não é um número menor que zero(0)"},
		{input: fromString("515.4554").Sub(fromString("200.9876"), rounding.Floor, 2), expected: fromString("314.46"),
			message: "Não deve existir arredondamento do resultado final, pois o mesmo não é um número menor que zero(0)"},

		{input: NewFromFloat(-515.4554).Sub(NewFromFloat(-200.9876), rounding.Ceil, 2), expected: NewFromFloat(-314.46),
			message: "Não deve existir arredondamento do resultado final, pois o mesmo não é um número maior que zero(0) nem igual"},
		{input: fromString("-515.4554").Sub(fromString("-200.9876"), rounding.Ceil, 2), expected: fromString("-314.46"),
			message: "Não deve existir arredondamento do resultado final, pois o mesmo não é um número maior que zero(0) nem igual"},

		{input: NewFromFloat(515.50001).Sub(NewFromFloat(200.510501), rounding.HalfDown, 2), expected: NewFromFloat(314.99),
			message: "Não faz arredondamento do resultado final, pois o mesmo não é um número maior que cinco(5)"},
		{input: fromString("515.50001").Sub(fromString("200.510501"), rounding.HalfDown, 2), expected: fromString("314.99"),
			message: "Não faz arredondamento do resultado final, pois o mesmo não é um número maior que cinco(5)"},

		{input: NewFromFloat(515.50001).Sub(NewFromFloat(200.510501), rounding.HalfEven, 2), expected: NewFromFloat(314.99),
			message: "Não faz arredondamento do resultado final, pois o mesmo não é um número maior que cinco(5)"},
		{input: fromString("515.50001").Sub(fromString("200.510501"), rounding.HalfEven, 2), expected: fromString("314.99"),
			message: "Não faz arredondamento do resultado final, pois o mesmo não é um número maior que cinco(5)"},

		{input: NewFromFloat(515.50001).Sub(NewFromFloat(200.510501), rounding.HalfUp, 2), expected: NewFromFloat(314.99),
			message: "Não faz arredondamento do resultado final, pois o mesmo não é um número maior que cinco(5)"},
		{input: fromString("515.50001").Sub(fromString("200.510501"), rounding.HalfUp, 2), expected: fromString("314.99"),
			message: "Não faz arredondamento do resultado final, pois o mesmo não é um número maior que cinco(5)"},
		{input: NewFromFloat(515.50001).Sub(NewFromFloat(200.510501), rounding.HalfUp, 2), expected: NewFromFloat(314.99),
			message: "Não faz arredondamento do resultado final, pois o mesmo não é um número igual à cinco(5)"},
		{input: fromString("515.50001").Sub(fromString("200.510501"), rounding.HalfUp, 2), expected: fromString("314.99"),
			message: "Não faz arredondamento do resultado final, pois o mesmo não é um número igual à cinco(5)"},

		{input: NewFromFloat(-515.4554).Sub(NewFromFloat(-200.9876), rounding.Down, 2), expected: NewFromFloat(-314.46),
			message: "Não faz arredondamento do resultado final quando numero for negativo"},
		{input: fromString("-515.4554").Sub(fromString("-200.9876"), rounding.Down, 2), expected: fromString("-314.46"),
			message: "Não faz arredondamento do resultado final quando numero for negativo"},
		{input: NewFromFloat(515.500).Sub(NewFromFloat(200.9876), rounding.Down, 2), expected: NewFromFloat(314.51),
			message: "Não faz arredondamento do resultado final quando numero for positivo e maior que cinco(5)"},
		{input: fromString("515.500").Sub(fromString("200.9876"), rounding.Down, 2), expected: fromString("314.51"),
			message: "Não faz arredondamento do resultado final quando numero for positivo e maior que cinco(5)"},
		{input: NewFromFloat(515.590).Sub(NewFromFloat(200.555), rounding.Down, 2), expected: NewFromFloat(315.03),
			message: "Não faz arredondamento do resultado final quando numero for positivo igual à cinco(5)"},
		{input: fromString("515.590").Sub(fromString("200.555"), rounding.Down, 2), expected: fromString("315.03"),
			message: "Não faz arredondamento do resultado final quando numero for positivo igual à cinco(5)"},
		{input: NewFromFloat(515.50001).Sub(NewFromFloat(200.510501), rounding.Down, 2), expected: NewFromFloat(314.98),
			message: "Não faz arredondamento do resultado final, pois o mesmo é um número igual à zero(0)"},
		{input: fromString("515.50001").Sub(fromString("200.510501"), rounding.Down, 2), expected: fromString("314.98"),
			message: "Não faz arredondamento do resultado final, pois o mesmo é um número igual à zero(0)"},
		{input: NewFromInt(515).Sub(NewFromInt(200), rounding.Down, 0), expected: NewFromInt(315),
			message: "Não existe arredondamento para numeros inteiros"},
	}
	for _, test := range tests {
		assertEquals(t, test.expected.ToFloat64(), test.input.ToFloat64())
	}
}

func TestMustAcceptAllTypesOfRoundingPossibleToMultiply(t *testing.T) {
	tests := []Test{
		{input: NewFromFloat(-515.4554).Mul(NewFromFloat(200.9876), rounding.Floor, 2), expected: NewFromFloat(-103600.15),
			message: "Arredondamento do resultado final, pois o mesmo é um número menor que zero(0)"},
		{input: fromString("-515.4554").Mul(fromString("200.9876"), rounding.Floor, 2), expected: fromString("-103600.15"),
			message: "Arredondamento do resultado final, pois o mesmo é um número menor que zero(0)"},

		{input: NewFromFloat(515.4554).Mul(NewFromFloat(200.9876), rounding.Ceil, 2), expected: NewFromFloat(103600.15),
			message: "Arredondamento do resultado final, pois o mesmo é um número maior que zero(0)"},
		{input: fromString("515.4554").Mul(fromString("200.9876"), rounding.Ceil, 2), expected: fromString("103600.15"),
			message: "Arredondamento do resultado final, pois o mesmo é um número maior que zero(0)"},
		{input: NewFromFloat(515.01).Mul(NewFromFloat(200.01), rounding.Ceil, 2), expected: NewFromFloat(103007.16),
			message: "Arredondamento do resultado final, pois o mesmo é um número igual à zero(0)"},
		{input: fromString("515.01").Mul(fromString("200.01"), rounding.Ceil, 2), expected: fromString("103007.16"),
			message: "Arredondamento do resultado final, pois o mesmo é um número igual à zero(0)"},

		{input: NewFromFloat(515.60).Mul(NewFromFloat(200.98), rounding.HalfDown, 2), expected: NewFromFloat(103625.29),
			message: "Arredondamento do resultado final, pois o mesmo é um número maior que cinco(5)"},
		{input: fromString("515.60").Mul(fromString("200.98"), rounding.HalfDown, 2), expected: fromString("103625.29"),
			message: "Arredondamento do resultado final, pois o mesmo é um número maior que cinco(5)"},

		{input: NewFromFloat(515.60).Mul(NewFromFloat(200.98), rounding.HalfEven, 2), expected: NewFromFloat(103625.29),
			message: "Arredondamento do resultado final, pois o mesmo é um número maior que cinco(5)"},
		{input: fromString("515.60").Mul(fromString("200.98"), rounding.HalfEven, 2), expected: fromString("103625.29"),
			message: "Arredondamento do resultado final, pois o mesmo é um número maior que cinco(5)"},

		{input: NewFromFloat(515.60).Mul(NewFromFloat(200.98), rounding.HalfUp, 2), expected: NewFromFloat(103625.29),
			message: "Arredondamento do resultado final, pois o mesmo é um número maior que cinco(5)"},
		{input: fromString("515.60").Mul(fromString("200.98"), rounding.HalfUp, 2), expected: fromString("103625.29"),
			message: "Arredondamento do resultado final, pois o mesmo é um número maior que cinco(5)"},
		{input: NewFromFloat(515.57).Mul(NewFromFloat(200.500), rounding.HalfUp, 2), expected: NewFromFloat(103371.79),
			message: "Arredondamento do resultado final, pois o mesmo é um número igual à cinco(5)"},
		{input: fromString("515.57").Mul(fromString("200.500"), rounding.HalfUp, 2), expected: fromString("103371.79"),
			message: "Arredondamento do resultado final, pois o mesmo é um número igual à cinco(5)"},

		{input: NewFromFloat(-515.4554).Mul(NewFromFloat(200.9876), rounding.Up, 2), expected: NewFromFloat(-103600.15),
			message: "Arredondamento do resultado final quando numero for negativo"},
		{input: fromString("-515.4554").Mul(fromString("200.9876"), rounding.Up, 2), expected: fromString("-103600.15"),
			message: "Arredondamento do resultado final quando numero for negativo"},
		{input: NewFromFloat(515.60).Mul(NewFromFloat(200.98), rounding.Up, 2), expected: NewFromFloat(103625.29),
			message: "Arredondamento do resultado final quando numero for positivo e maior que cinco(5)"},
		{input: fromString("515.60").Mul(fromString("200.98"), rounding.Up, 2), expected: fromString("103625.29"),
			message: "Arredondamento do resultado final quando numero for positivo e maior que cinco(5)"},
		{input: NewFromFloat(515.57).Mul(NewFromFloat(200.500), rounding.Up, 2), expected: NewFromFloat(103371.79),
			message: "Arredondamento do resultado final quando numero for positivo igual à cinco(5)"},
		{input: fromString("515.57").Mul(fromString("200.500"), rounding.Up, 2), expected: fromString("103371.79"),
			message: "Arredondamento do resultado final quando numero for positivo igual à cinco(5)"},
		{input: NewFromFloat(515.01).Mul(NewFromFloat(200.01), rounding.Up, 2), expected: NewFromFloat(103007.16),
			message: "Arredondamento do resultado final, pois o mesmo é um número igual à zero(0)"},
		{input: fromString("515.01").Mul(fromString("200.01"), rounding.Up, 2), expected: fromString("103007.16"),
			message: "Arredondamento do resultado final, pois o mesmo é um número igual à zero(0)"},
	}
	for _, test := range tests {
		assertEquals(t, test.expected.ToFloat64(), test.input.ToFloat64())
	}
}

func TestNotMustAcceptRoundingToMultiply(t *testing.T) {
	tests := []Test{
		{input: NewFromFloat(515.4554).Mul(NewFromFloat(200.9876), rounding.Floor, 2), expected: NewFromFloat(103600.14),
			message: "Não deve existir arredondamento do resultado final, pois o mesmo não é um número menor que zero(0)"},
		{input: fromString("515.4554").Mul(fromString("200.9876"), rounding.Floor, 2), expected: fromString("103600.14"),
			message: "Não deve existir arredondamento do resultado final, pois o mesmo não é um número menor que zero(0)"},

		{input: NewFromFloat(-515.4554).Mul(NewFromFloat(200.9876), rounding.Ceil, 2), expected: NewFromFloat(-103600.14),
			message: "Não deve existir arredondamento do resultado final, pois o mesmo não é um número maior que zero(0) nem igual"},
		{input: fromString("-515.4554").Mul(fromString("200.9876"), rounding.Ceil, 2), expected: fromString("-103600.14"),
			message: "Não deve existir arredondamento do resultado final, pois o mesmo não é um número maior que zero(0) nem igual"},

		{input: NewFromFloat(515.65).Mul(NewFromFloat(200.100), rounding.HalfDown, 2), expected: NewFromFloat(103181.56),
			message: "Não faz arredondamento do resultado final, pois o mesmo não é um número maior que cinco(5)"},
		{input: fromString("515.65").Mul(fromString("200.100"), rounding.HalfDown, 2), expected: fromString("103181.56"),
			message: "Não faz arredondamento do resultado final, pois o mesmo não é um número maior que cinco(5)"},

		{input: NewFromFloat(515.60).Mul(NewFromFloat(200.115), rounding.HalfEven, 2), expected: NewFromFloat(103179.29),
			message: "Não faz arredondamento do resultado final, pois o mesmo não é um número maior que cinco(5)"},
		{input: fromString("515.60").Mul(fromString("200.115"), rounding.HalfEven, 2), expected: fromString("103179.29"),
			message: "Não faz arredondamento do resultado final, pois o mesmo não é um número maior que cinco(5)"},

		{input: NewFromFloat(515.60).Mul(NewFromFloat(200.115), rounding.HalfUp, 2), expected: NewFromFloat(103179.29),
			message: "Não faz arredondamento do resultado final, pois o mesmo não é um número maior que cinco(5)"},
		{input: fromString("515.60").Mul(fromString("200.115"), rounding.HalfUp, 2), expected: fromString("103179.29"),
			message: "Não faz arredondamento do resultado final, pois o mesmo não é um número maior que cinco(5)"},
		{input: NewFromFloat(515.60).Mul(NewFromFloat(200.115), rounding.HalfUp, 2), expected: NewFromFloat(103179.29),
			message: "Não faz arredondamento do resultado final, pois o mesmo não é um número igual à cinco(5)"},
		{input: fromString("515.60").Mul(fromString("200.115"), rounding.HalfUp, 2), expected: fromString("103179.29"),
			message: "Não faz arredondamento do resultado final, pois o mesmo não é um número igual à cinco(5)"},

		{input: NewFromFloat(-515.4554).Mul(NewFromFloat(200.9876), rounding.Down, 2), expected: NewFromFloat(-103600.14),
			message: "Não faz arredondamento do resultado final quando numero for negativo"},
		{input: fromString("-515.4554").Mul(fromString("200.9876"), rounding.Down, 2), expected: fromString("-103600.14"),
			message: "Não faz arredondamento do resultado final quando numero for negativo"},
		{input: NewFromFloat(515.60).Mul(NewFromFloat(200.98), rounding.Down, 2), expected: NewFromFloat(103625.28),
			message: "Não faz arredondamento do resultado final quando numero for positivo e maior que cinco(5)"},
		{input: fromString("515.60").Mul(fromString("200.98"), rounding.Down, 2), expected: fromString("103625.28"),
			message: "Não faz arredondamento do resultado final quando numero for positivo e maior que cinco(5)"},
		{input: NewFromFloat(515.57).Mul(NewFromFloat(200.500), rounding.Down, 2), expected: NewFromFloat(103371.78),
			message: "Não faz arredondamento do resultado final quando numero for positivo igual à cinco(5)"},
		{input: fromString("515.57").Mul(fromString("200.500"), rounding.Down, 2), expected: fromString("103371.78"),
			message: "Não faz arredondamento do resultado final quando numero for positivo igual à cinco(5)"},
		{input: NewFromFloat(515.01).Mul(NewFromFloat(200.01), rounding.Down, 2), expected: NewFromFloat(103007.15),
			message: "Não faz arredondamento do resultado final, pois o mesmo é um número igual à zero(0)"},
		{input: fromString("515.01").Mul(fromString("200.01"), rounding.Down, 2), expected: fromString("103007.15"),
			message: "Não faz arredondamento do resultado final, pois o mesmo é um número igual à zero(0)"},
		{input: NewFromInt(515).Mul(NewFromInt(200), rounding.Down, 0), expected: NewFromInt(103000),
			message: "Não existe arredondamento para numeros inteiros"},
	}
	for _, test := range tests {
		assertEquals(t, test.expected.ToFloat64(), test.input.ToFloat64())
	}
}

func assertEquals(t *testing.T, expected interface{}, actual interface{}) {
	if actual != expected {
		t.Fatalf("was expected '%s', got '%s'", expected, actual)
	}
}

type Test struct {
	input    Money
	expected Money
	message  string
}

func fromString(str string) Money {
	m, _ := NewFromString(str)
	return m
}
