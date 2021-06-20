package money

import (
	"github.com/wadey/go-rounding"
	"testing"
)

func TestMustAcceptAllTypesOfRoundingPossibleToSum(t *testing.T) {
	tests := []Test{
		createTest(NewFromFloat(-515.4554), NewFromFloat(-200.9876), NewFromFloat(-716.45), rounding.Floor,
			"Arredondamento do resultado final, pois o mesmo é um número menor que zero(0)"),
		createTest(fromString("-515.4554"), fromString("-200.9876"), fromString("-716.45"), rounding.Floor,
			"Arredondamento do resultado final, pois o mesmo é um número menor que zero(0)"),

		createTest(NewFromFloat(515.4554), NewFromFloat(200.9876), NewFromFloat(716.45), rounding.Ceil,
			"Arredondamento do resultado final, pois o mesmo é um número maior que zero(0)"),
		createTest(fromString("515.4554"), fromString("200.9876"), fromString("716.45"), rounding.Ceil,
			"Arredondamento do resultado final, pois o mesmo é um número maior que zero(0)"),
		createTest(NewFromFloat(515.50001), NewFromFloat(200.510501), NewFromFloat(716.02), rounding.Ceil,
			"Arredondamento do resultado final, pois o mesmo é um número igual à zero(0)"),
		createTest(fromString("515.50001"), fromString("200.510501"), fromString("716.02"), rounding.Ceil,
			"Arredondamento do resultado final, pois o mesmo é um número igual à zero(0)"),

		createTest(NewFromFloat(515.500), NewFromFloat(200.9876), NewFromFloat(716.49), rounding.HalfDown,
			"Arredondamento do resultado final, pois o mesmo é um número maior que cinco(5)"),
		createTest(fromString("515.500"), fromString("200.9876"), fromString("716.49"), rounding.HalfDown,
			"Arredondamento do resultado final, pois o mesmo é um número maior que cinco(5)"),

		createTest(NewFromFloat(515.500), NewFromFloat(200.9876), NewFromFloat(716.49), rounding.HalfEven,
			"Arredondamento do resultado final, pois o mesmo é um número maior que cinco(5)"),
		createTest(fromString("515.500"), fromString("200.9876"), fromString("716.49"), rounding.HalfEven,
			"Arredondamento do resultado final, pois o mesmo é um número maior que cinco(5)"),

		createTest(NewFromFloat(515.500), NewFromFloat(200.9876), NewFromFloat(716.49), rounding.HalfUp,
			"Arredondamento do resultado final, pois o mesmo é um número maior que cinco(5)"),
		createTest(fromString("515.500"), fromString("200.9876"), fromString("716.49"), rounding.HalfUp,
			"Arredondamento do resultado final, pois o mesmo é um número maior que cinco(5)"),
		createTest(NewFromFloat(515.590), NewFromFloat(200.555), NewFromFloat(716.15), rounding.HalfUp,
			"Arredondamento do resultado final, pois o mesmo é um número igual à cinco(5)"),
		createTest(fromString("515.590"), fromString("200.555"), fromString("716.15"), rounding.HalfUp,
			"Arredondamento do resultado final, pois o mesmo é um número igual à cinco(5)"),

		createTest(NewFromFloat(-515.4554), NewFromFloat(-200.9876), NewFromFloat(-716.45), rounding.Up,
			"Arredondamento do resultado final quando numero for negativo"),
		createTest(fromString("-515.4554"), fromString("-200.9876"), fromString("-716.45"), rounding.Up,
			"Arredondamento do resultado final quando numero for negativo"),
		createTest(NewFromFloat(515.500), NewFromFloat(200.9876), NewFromFloat(716.49), rounding.Up,
			"Arredondamento do resultado final quando numero for positivo e maior que cinco(5)"),
		createTest(fromString("515.500"), fromString("200.9876"), fromString("716.49"), rounding.Up,
			"Arredondamento do resultado final quando numero for positivo e maior que cinco(5)"),
		createTest(NewFromFloat(515.590), NewFromFloat(200.555), NewFromFloat(716.15), rounding.Up,
			"Arredondamento do resultado final quando numero for positivo igual à cinco(5)"),
		createTest(fromString("515.590"), fromString("200.555"), fromString("716.15"), rounding.Up,
			"Arredondamento do resultado final quando numero for positivo igual à cinco(5)"),
		createTest(NewFromFloat(515.50001), NewFromFloat(200.510501), NewFromFloat(716.02), rounding.Up,
			"Arredondamento do resultado final, pois o mesmo é um número igual à zero(0)"),
		createTest(fromString("515.50001"), fromString("200.510501"), fromString("716.02"), rounding.Up,
			"Arredondamento do resultado final, pois o mesmo é um número igual à zero(0)"),
	}
	for _, test := range tests {
		assertEquals(t, test.expected.ToFloat64(), test.input1.Add(test.input2, test.rounding, 2).ToFloat64())
	}
}

func TestNotMustAcceptRoundingToSum(t *testing.T) {
	tests := []Test{
		createTest(NewFromFloat(515.4554), NewFromFloat(200.9876), NewFromFloat(716.44), rounding.Floor,
			"Não deve existir arredondamento do resultado final, pois o mesmo não é um número menor que zero(0)"),
		createTest(fromString("515.4554"), fromString("200.9876"), fromString("716.44"), rounding.Floor,
			"Não deve existir arredondamento do resultado final, pois o mesmo não é um número menor que zero(0)"),

		createTest(NewFromFloat(-515.4554), NewFromFloat(-200.9876), NewFromFloat(-716.44), rounding.Ceil,
			"Não deve existir arredondamento do resultado final, pois o mesmo não é um número maior que zero(0) nem igual"),
		createTest(fromString("-515.4554"), fromString("-200.9876"), fromString("-716.44"), rounding.Ceil,
			"Não deve existir arredondamento do resultado final, pois o mesmo não é um número maior que zero(0) nem igual"),

		createTest(NewFromFloat(515.50001), NewFromFloat(200.510501), NewFromFloat(716.01), rounding.HalfDown,
			"Não faz arredondamento do resultado final, pois o mesmo não é um número maior que cinco(5)"),
		createTest(fromString("515.50001"), fromString("200.510501"), fromString("716.01"), rounding.HalfDown,
			"Não faz arredondamento do resultado final, pois o mesmo não é um número maior que cinco(5)"),

		createTest(NewFromFloat(515.50001), NewFromFloat(200.510501), NewFromFloat(716.01), rounding.HalfEven,
			"Não faz arredondamento do resultado final, pois o mesmo não é um número maior que cinco(5)"),
		createTest(fromString("515.50001"), fromString("200.510501"), fromString("716.01"), rounding.HalfEven,
			"Não faz arredondamento do resultado final, pois o mesmo não é um número maior que cinco(5)"),

		createTest(NewFromFloat(515.50001), NewFromFloat(200.510501), NewFromFloat(716.01), rounding.HalfUp,
			"Não faz arredondamento do resultado final, pois o mesmo não é um número maior que cinco(5)"),
		createTest(fromString("515.50001"), fromString("200.510501"), fromString("716.01"), rounding.HalfUp,
			"Não faz arredondamento do resultado final, pois o mesmo não é um número maior que cinco(5)"),
		createTest(NewFromFloat(515.50001), NewFromFloat(200.510501), NewFromFloat(716.01), rounding.HalfUp,
			"Não faz arredondamento do resultado final, pois o mesmo não é um número igual à cinco(5)"),
		createTest(fromString("515.50001"), fromString("200.510501"), fromString("716.01"), rounding.HalfUp,
			"Não faz arredondamento do resultado final, pois o mesmo não é um número igual à cinco(5)"),

		createTest(NewFromFloat(-515.4554), NewFromFloat(-200.9876), NewFromFloat(-716.44), rounding.Down,
			"Não faz arredondamento do resultado final quando numero for negativo"),
		createTest(fromString("-515.4554"), fromString("-200.9876"), fromString("-716.44"), rounding.Down,
			"Não faz arredondamento do resultado final quando numero for negativo"),
		createTest(NewFromFloat(515.500), NewFromFloat(200.9876), NewFromFloat(716.48), rounding.Down,
			"Não faz arredondamento do resultado final quando numero for positivo e maior que cinco(5)"),
		createTest(fromString("515.500"), fromString("200.9876"), fromString("716.48"), rounding.Down,
			"Não faz arredondamento do resultado final quando numero for positivo e maior que cinco(5)"),
		createTest(NewFromFloat(515.590), NewFromFloat(200.555), NewFromFloat(716.14), rounding.Down,
			"Não faz arredondamento do resultado final quando numero for positivo igual à cinco(5)"),
		createTest(fromString("515.590"), fromString("200.555"), fromString("716.14"), rounding.Down,
			"Não faz arredondamento do resultado final quando numero for positivo igual à cinco(5)"),
		createTest(NewFromFloat(515.50001), NewFromFloat(200.510501), NewFromFloat(716.01), rounding.Down,
			"Não faz arredondamento do resultado final, pois o mesmo é um número igual à zero(0)"),
		createTest(fromString("515.50001"), fromString("200.510501"), fromString("716.01"), rounding.Down,
			"Não faz arredondamento do resultado final, pois o mesmo é um número igual à zero(0)"),
		createTest(NewFromInt(515), NewFromInt(200), NewFromInt(715), rounding.Down,
			"Não existe arredondamento para numeros inteiros"),
	}
	for _, test := range tests {
		assertEquals(t, test.expected.ToFloat64(), test.input1.Add(test.input2, test.rounding, 2).ToFloat64())
	}
}

func TestMustAcceptAllTypesOfRoundingPossibleToSubtraction(t *testing.T) {
	tests := []Test{
		createTest(NewFromFloat(-515.4554), NewFromFloat(-200.9876), NewFromFloat(-314.47), rounding.Floor,
			"Arredondamento do resultado final, pois o mesmo é um número menor que zero(0)"),
		createTest(fromString("-515.4554"), fromString("-200.9876"), fromString("-314.47"), rounding.Floor,
			"Arredondamento do resultado final, pois o mesmo é um número menor que zero(0)"),

		createTest(NewFromFloat(515.4554), NewFromFloat(200.9876), NewFromFloat(314.47), rounding.Ceil,
			"Arredondamento do resultado final, pois o mesmo é um número maior que zero(0)"),
		createTest(fromString("515.4554"), fromString("200.9876"), fromString("314.47"), rounding.Ceil,
			"Arredondamento do resultado final, pois o mesmo é um número maior que zero(0)"),
		createTest(NewFromFloat(515.2021), NewFromFloat(200.522), NewFromFloat(314.69), rounding.Ceil,
			"Arredondamento do resultado final, pois o mesmo é um número igual à zero(0)"),
		createTest(fromString("515.2021"), fromString("200.522"), fromString("314.69"), rounding.Ceil,
			"Arredondamento do resultado final, pois o mesmo é um número igual à zero(0)"),

		createTest(NewFromFloat(515.500), NewFromFloat(200.7222), NewFromFloat(314.78), rounding.HalfDown,
			"Arredondamento do resultado final, pois o mesmo é um número maior que cinco(5)"),
		createTest(fromString("515.500"), fromString("200.7222"), fromString("314.78"), rounding.HalfDown,
			"Arredondamento do resultado final, pois o mesmo é um número maior que cinco(5)"),

		createTest(NewFromFloat(515.500), NewFromFloat(200.7222), NewFromFloat(314.78), rounding.HalfEven,
			"Arredondamento do resultado final, pois o mesmo é um número maior que cinco(5)"),
		createTest(fromString("515.500"), fromString("200.7222"), fromString("314.78"), rounding.HalfEven,
			"Arredondamento do resultado final, pois o mesmo é um número maior que cinco(5)"),

		createTest(NewFromFloat(515.500), NewFromFloat(200.7222), NewFromFloat(314.78), rounding.HalfUp,
			"Arredondamento do resultado final, pois o mesmo é um número maior que cinco(5)"),
		createTest(fromString("515.500"), fromString("200.7222"), fromString("314.78"), rounding.HalfUp,
			"Arredondamento do resultado final, pois o mesmo é um número maior que cinco(5)"),
		createTest(NewFromFloat(515.500), NewFromFloat(200.725), NewFromFloat(314.78), rounding.HalfUp,
			"Arredondamento do resultado final, pois o mesmo é um número igual à cinco(5)"),
		createTest(fromString("515.500"), fromString("200.725"), fromString("314.78"), rounding.HalfUp,
			"Arredondamento do resultado final, pois o mesmo é um número igual à cinco(5)"),

		createTest(NewFromFloat(-515.4554), NewFromFloat(-200.9876), NewFromFloat(-314.47), rounding.Up,
			"Arredondamento do resultado final quando numero for negativo"),
		createTest(fromString("-515.4554"), fromString("-200.9876"), fromString("-314.47"), rounding.Up,
			"Arredondamento do resultado final quando numero for negativo"),
		createTest(NewFromFloat(515.500), NewFromFloat(200.7222), NewFromFloat(314.78), rounding.Up,
			"Arredondamento do resultado final quando numero for positivo e maior que cinco(5)"),
		createTest(fromString("515.500"), fromString("200.7222"), fromString("314.78"), rounding.Up,
			"Arredondamento do resultado final quando numero for positivo e maior que cinco(5)"),
		createTest(NewFromFloat(515.500), NewFromFloat(200.725), NewFromFloat(314.78), rounding.Up,
			"Arredondamento do resultado final quando numero for positivo igual à cinco(5)"),
		createTest(fromString("515.500"), fromString("200.725"), fromString("314.78"), rounding.Up,
			"Arredondamento do resultado final quando numero for positivo igual à cinco(5)"),
		createTest(NewFromFloat(515.50001), NewFromFloat(200.510501), NewFromFloat(314.99), rounding.Up,
			"Arredondamento do resultado final, pois o mesmo é um número igual à zero(0)"),
		createTest(fromString("515.50001"), fromString("200.510501"), fromString("314.99"), rounding.Up,
			"Arredondamento do resultado final, pois o mesmo é um número igual à zero(0)"),
	}
	for _, test := range tests {
		assertEquals(t, test.expected.ToFloat64(), test.input1.Sub(test.input2, test.rounding, 2).ToFloat64())
	}
}

func TestNotMustAcceptRoundingToSubtraction(t *testing.T) {
	tests := []Test{
		createTest(NewFromFloat(515.4554), NewFromFloat(200.9876), NewFromFloat(314.46), rounding.Floor,
			"Não deve existir arredondamento do resultado final, pois o mesmo não é um número menor que zero(0)"),
		createTest(fromString("515.4554"), fromString("200.9876"), fromString("314.46"), rounding.Floor,
			"Não deve existir arredondamento do resultado final, pois o mesmo não é um número menor que zero(0)"),

		createTest(NewFromFloat(-515.4554), NewFromFloat(-200.9876), NewFromFloat(-314.46), rounding.Ceil,
			"Não deve existir arredondamento do resultado final, pois o mesmo não é um número maior que zero(0) nem igual"),
		createTest(fromString("-515.4554"), fromString("-200.9876"), fromString("-314.46"), rounding.Ceil,
			"Não deve existir arredondamento do resultado final, pois o mesmo não é um número maior que zero(0) nem igual"),

		createTest(NewFromFloat(515.50001), NewFromFloat(200.510501), NewFromFloat(314.99), rounding.HalfDown,
			"Não faz arredondamento do resultado final, pois o mesmo não é um número maior que cinco(5)"),
		createTest(fromString("515.50001"), fromString("200.510501"), fromString("314.99"), rounding.HalfDown,
			"Não faz arredondamento do resultado final, pois o mesmo não é um número maior que cinco(5)"),

		createTest(NewFromFloat(515.50001), NewFromFloat(200.510501), NewFromFloat(314.99), rounding.HalfEven,
			"Não faz arredondamento do resultado final, pois o mesmo não é um número maior que cinco(5)"),
		createTest(fromString("515.50001"), fromString("200.510501"), fromString("314.99"), rounding.HalfEven,
			"Não faz arredondamento do resultado final, pois o mesmo não é um número maior que cinco(5)"),

		createTest(NewFromFloat(515.50001), NewFromFloat(200.510501), NewFromFloat(314.99), rounding.HalfUp,
			"Não faz arredondamento do resultado final, pois o mesmo não é um número maior que cinco(5)"),
		createTest(fromString("515.50001"), fromString("200.510501"), fromString("314.99"), rounding.HalfUp,
			"Não faz arredondamento do resultado final, pois o mesmo não é um número maior que cinco(5)"),
		createTest(NewFromFloat(515.50001), NewFromFloat(200.510501), NewFromFloat(314.99), rounding.HalfUp,
			"Não faz arredondamento do resultado final, pois o mesmo não é um número igual à cinco(5)"),
		createTest(fromString("515.50001"), fromString("200.510501"), fromString("314.99"), rounding.HalfUp,
			"Não faz arredondamento do resultado final, pois o mesmo não é um número igual à cinco(5)"),

		createTest(NewFromFloat(-515.4554), NewFromFloat(-200.9876), NewFromFloat(-314.46), rounding.Down,
			"Não faz arredondamento do resultado final quando numero for negativo"),
		createTest(fromString("-515.4554"), fromString("-200.9876"), fromString("-314.46"), rounding.Down,
			"Não faz arredondamento do resultado final quando numero for negativo"),
		createTest(NewFromFloat(515.500), NewFromFloat(200.9876), NewFromFloat(314.51), rounding.Down,
			"Não faz arredondamento do resultado final quando numero for positivo e maior que cinco(5)"),
		createTest(fromString("515.500"), fromString("200.9876"), fromString("314.51"), rounding.Down,
			"Não faz arredondamento do resultado final quando numero for positivo e maior que cinco(5)"),
		createTest(NewFromFloat(515.590), NewFromFloat(200.555), NewFromFloat(315.03), rounding.Down,
			"Não faz arredondamento do resultado final quando numero for positivo igual à cinco(5)"),
		createTest(fromString("515.590"), fromString("200.555"), fromString("315.03"), rounding.Down,
			"Não faz arredondamento do resultado final quando numero for positivo igual à cinco(5)"),
		createTest(NewFromFloat(515.50001), NewFromFloat(200.510501), NewFromFloat(314.98), rounding.Down,
			"Não faz arredondamento do resultado final, pois o mesmo é um número igual à zero(0)"),
		createTest(fromString("515.50001"), fromString("200.510501"), fromString("314.98"), rounding.Down,
			"Não faz arredondamento do resultado final, pois o mesmo é um número igual à zero(0)"),
		createTest(NewFromInt(515), NewFromInt(200), NewFromInt(315), rounding.Down,
			"Não existe arredondamento para numeros inteiros"),
	}
	for _, test := range tests {
		assertEquals(t, test.expected.ToFloat64(), test.input1.Sub(test.input2, test.rounding, 2).ToFloat64())
	}
}

func TestMustAcceptAllTypesOfRoundingPossibleToMultiply(t *testing.T) {
	tests := []Test{
		createTest(NewFromFloat(-515.4554), NewFromFloat(200.9876), NewFromFloat(-103600.15), rounding.Floor,
			"Arredondamento do resultado final, pois o mesmo é um número menor que zero(0)"),
		createTest(fromString("-515.4554"), fromString("200.9876"), fromString("-103600.15"), rounding.Floor,
			"Arredondamento do resultado final, pois o mesmo é um número menor que zero(0)"),

		createTest(NewFromFloat(515.4554), NewFromFloat(200.9876), NewFromFloat(103600.15), rounding.Ceil,
			"Arredondamento do resultado final, pois o mesmo é um número maior que zero(0)"),
		createTest(fromString("515.4554"), fromString("200.9876"), fromString("103600.15"), rounding.Ceil,
			"Arredondamento do resultado final, pois o mesmo é um número maior que zero(0)"),
		createTest(NewFromFloat(515.01), NewFromFloat(200.01), NewFromFloat(103007.16), rounding.Ceil,
			"Arredondamento do resultado final, pois o mesmo é um número igual à zero(0)"),
		createTest(fromString("515.01"), fromString("200.01"), fromString("103007.16"), rounding.Ceil,
			"Arredondamento do resultado final, pois o mesmo é um número igual à zero(0)"),

		createTest(NewFromFloat(515.60), NewFromFloat(200.98), NewFromFloat(103625.29), rounding.HalfDown,
			"Arredondamento do resultado final, pois o mesmo é um número maior que cinco(5)"),
		createTest(fromString("515.60"), fromString("200.98"), fromString("103625.29"), rounding.HalfDown,
			"Arredondamento do resultado final, pois o mesmo é um número maior que cinco(5)"),

		createTest(NewFromFloat(515.60), NewFromFloat(200.98), NewFromFloat(103625.29), rounding.HalfEven,
			"Arredondamento do resultado final, pois o mesmo é um número maior que cinco(5)"),
		createTest(fromString("515.60"), fromString("200.98"), fromString("103625.29"), rounding.HalfEven,
			"Arredondamento do resultado final, pois o mesmo é um número maior que cinco(5)"),

		createTest(NewFromFloat(515.60), NewFromFloat(200.98), NewFromFloat(103625.29), rounding.HalfUp,
			"Arredondamento do resultado final, pois o mesmo é um número maior que cinco(5)"),
		createTest(fromString("515.60"), fromString("200.98"), fromString("103625.29"), rounding.HalfUp,
			"Arredondamento do resultado final, pois o mesmo é um número maior que cinco(5)"),
		createTest(NewFromFloat(515.57), NewFromFloat(200.500), NewFromFloat(103371.79), rounding.HalfUp,
			"Arredondamento do resultado final, pois o mesmo é um número igual à cinco(5)"),
		createTest(fromString("515.57"), fromString("200.500"), fromString("103371.79"), rounding.HalfUp,
			"Arredondamento do resultado final, pois o mesmo é um número igual à cinco(5)"),

		createTest(NewFromFloat(-515.4554), NewFromFloat(200.9876), NewFromFloat(-103600.15), rounding.Up,
			"Arredondamento do resultado final quando numero for negativo"),
		createTest(fromString("-515.4554"), fromString("200.9876"), fromString("-103600.15"), rounding.Up,
			"Arredondamento do resultado final quando numero for negativo"),
		createTest(NewFromFloat(515.60), NewFromFloat(200.98), NewFromFloat(103625.29), rounding.Up,
			"Arredondamento do resultado final quando numero for positivo e maior que cinco(5)"),
		createTest(fromString("515.60"), fromString("200.98"), fromString("103625.29"), rounding.Up,
			"Arredondamento do resultado final quando numero for positivo e maior que cinco(5)"),
		createTest(NewFromFloat(515.57), NewFromFloat(200.500), NewFromFloat(103371.79), rounding.Up,
			"Arredondamento do resultado final quando numero for positivo igual à cinco(5)"),
		createTest(fromString("515.57"), fromString("200.500"), fromString("103371.79"), rounding.Up,
			"Arredondamento do resultado final quando numero for positivo igual à cinco(5)"),
		createTest(NewFromFloat(515.01), NewFromFloat(200.01), NewFromFloat(103007.16), rounding.Up,
			"Arredondamento do resultado final, pois o mesmo é um número igual à zero(0)"),
		createTest(fromString("515.01"), fromString("200.01"), fromString("103007.16"), rounding.Up,
			"Arredondamento do resultado final, pois o mesmo é um número igual à zero(0)"),
	}
	for _, test := range tests {
		assertEquals(t, test.expected.ToFloat64(), test.input1.Mul(test.input2, test.rounding, 2).ToFloat64())
	}
}

func TestNotMustAcceptRoundingToMultiply(t *testing.T) {
	tests := []Test{
		createTest(NewFromFloat(515.4554), NewFromFloat(200.9876), NewFromFloat(103600.14), rounding.Floor,
			"Não deve existir arredondamento do resultado final, pois o mesmo não é um número menor que zero(0)"),
		createTest(fromString("515.4554"), fromString("200.9876"), fromString("103600.14"), rounding.Floor,
			"Não deve existir arredondamento do resultado final, pois o mesmo não é um número menor que zero(0)"),

		createTest(NewFromFloat(-515.4554), NewFromFloat(200.9876), NewFromFloat(-103600.14), rounding.Ceil,
			"Não deve existir arredondamento do resultado final, pois o mesmo não é um número maior que zero(0) nem igual"),
		createTest(fromString("-515.4554"), fromString("200.9876"), fromString("-103600.14"), rounding.Ceil,
			"Não deve existir arredondamento do resultado final, pois o mesmo não é um número maior que zero(0) nem igual"),

		createTest(NewFromFloat(515.65), NewFromFloat(200.100), NewFromFloat(103181.56), rounding.HalfDown,
			"Não faz arredondamento do resultado final, pois o mesmo não é um número maior que cinco(5)"),
		createTest(fromString("515.65"), fromString("200.100"), fromString("103181.56"), rounding.HalfDown,
			"Não faz arredondamento do resultado final, pois o mesmo não é um número maior que cinco(5)"),

		createTest(NewFromFloat(515.60), NewFromFloat(200.115), NewFromFloat(103179.29), rounding.HalfEven,
			"Não faz arredondamento do resultado final, pois o mesmo não é um número maior que cinco(5)"),
		createTest(fromString("515.60"), fromString("200.115"), fromString("103179.29"), rounding.HalfEven,
			"Não faz arredondamento do resultado final, pois o mesmo não é um número maior que cinco(5)"),

		createTest(NewFromFloat(515.60), NewFromFloat(200.115), NewFromFloat(103179.29), rounding.HalfUp,
			"Não faz arredondamento do resultado final, pois o mesmo não é um número maior que cinco(5)"),
		createTest(fromString("515.60"), fromString("200.115"), fromString("103179.29"), rounding.HalfUp,
			"Não faz arredondamento do resultado final, pois o mesmo não é um número maior que cinco(5)"),
		createTest(NewFromFloat(515.60), NewFromFloat(200.115), NewFromFloat(103179.29), rounding.HalfUp,
			"Não faz arredondamento do resultado final, pois o mesmo não é um número igual à cinco(5)"),
		createTest(fromString("515.60"), fromString("200.115"), fromString("103179.29"), rounding.HalfUp,
			"Não faz arredondamento do resultado final, pois o mesmo não é um número igual à cinco(5)"),

		createTest(NewFromFloat(-515.4554), NewFromFloat(200.9876), NewFromFloat(-103600.14), rounding.Down,
			"Não faz arredondamento do resultado final quando numero for negativo"),
		createTest(fromString("-515.4554"), fromString("200.9876"), fromString("-103600.14"), rounding.Down,
			"Não faz arredondamento do resultado final quando numero for negativo"),
		createTest(NewFromFloat(515.60), NewFromFloat(200.98), NewFromFloat(103625.28), rounding.Down,
			"Não faz arredondamento do resultado final quando numero for positivo e maior que cinco(5)"),
		createTest(fromString("515.60"), fromString("200.98"), fromString("103625.28"), rounding.Down,
			"Não faz arredondamento do resultado final quando numero for positivo e maior que cinco(5)"),
		createTest(NewFromFloat(515.57), NewFromFloat(200.500), NewFromFloat(103371.78), rounding.Down,
			"Não faz arredondamento do resultado final quando numero for positivo igual à cinco(5)"),
		createTest(fromString("515.57"), fromString("200.500"), fromString("103371.78"), rounding.Down,
			"Não faz arredondamento do resultado final quando numero for positivo igual à cinco(5)"),
		createTest(NewFromFloat(515.01), NewFromFloat(200.01), NewFromFloat(103007.15), rounding.Down,
			"Não faz arredondamento do resultado final, pois o mesmo é um número igual à zero(0)"),
		createTest(fromString("515.01"), fromString("200.01"), fromString("103007.15"), rounding.Down,
			"Não faz arredondamento do resultado final, pois o mesmo é um número igual à zero(0)"),
		createTest(NewFromInt(515), NewFromInt(200), NewFromInt(103000), rounding.Down,
			"Não existe arredondamento para numeros inteiros"),
	}
	for _, test := range tests {
		assertEquals(t, test.expected.ToFloat64(), test.input1.Mul(test.input2, test.rounding, 2).ToFloat64())
	}
}

func assertEquals(t *testing.T, expected interface{}, actual interface{}) {
	if actual != expected {
		t.Fatalf("was expected '%s', got '%s'", expected, actual)
	}
}

type Test struct {
	input1   Money
	input2   Money
	rounding rounding.RoundingMode
	expected Money
	message  string
}

func createTest(input1, input2, expected Money, rounding rounding.RoundingMode, message string) Test {
	return Test{input1, input2, rounding, expected, message}
}

func fromString(str string) Money {
	m, _ := NewFromString(str)
	return m
}
