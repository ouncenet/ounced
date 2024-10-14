package util_test

import (
	"fmt"
	"github.com/ouncenet/ounced/util/difficulty"
	"math"
	"math/big"

	"github.com/ouncenet/ounced/util"
)

func ExampleAmount() {

	a := util.Amount(0)
	fmt.Println("Zero Grain:", a)

	a = util.Amount(1e8)
	fmt.Println("100,000,000 Grain:", a)

	a = util.Amount(1e5)
	fmt.Println("100,000 Grain:", a)
	// Output:
	// Zero Grain: 0 OZ
	// 100,000,000 Grain: 1 OZ
	// 100,000 Grain: 0.001 OZ
}

func ExampleNewAmount() {
	amountOne, err := util.NewAmount(1)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(amountOne) //Output 1

	amountFraction, err := util.NewAmount(0.01234567)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(amountFraction) //Output 2

	amountZero, err := util.NewAmount(0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(amountZero) //Output 3

	amountNaN, err := util.NewAmount(math.NaN())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(amountNaN) //Output 4

	// Output: 1 OZ
	// 0.01234567 OZ
	// 0 OZ
	// invalid ounce amount
}

func ExampleAmount_unitConversions() {
	amount := util.Amount(44433322211100)

	fmt.Println("Grain to kOZ:", amount.Format(util.AmountKiloOZ))
	fmt.Println("Grain to OZ:", amount)
	fmt.Println("Grain to MilliOZ:", amount.Format(util.AmountMilliOZ))
	fmt.Println("Grain to MicroμOZ:", amount.Format(util.AmountMicroμOZ))
	fmt.Println("Grain to Grain:", amount.Format(util.AmountGrain))

	// Output:
	// Grain to kOZ: 444.333222111 kOZ
	// Grain to OZ: 444333.222111 OZ
	// Grain to MilliOZ: 444333222.111 mOZ
	// Grain to MicroμOZ: 444333222111 μOZ
	// Grain to Grain: 44433322211100 Grain
}

// This example demonstrates how to convert the compact "bits" in a block header
// which represent the target difficulty to a big integer and display it using
// the typical hex notation.
func ExampleCompactToBig() {
	bits := uint32(419465580)
	targetDifficulty := difficulty.CompactToBig(bits)

	// Display it in hex.
	fmt.Printf("%064x\n", targetDifficulty.Bytes())

	// Output:
	// 0000000000000000896c00000000000000000000000000000000000000000000
}

// This example demonstrates how to convert a target difficulty into the compact
// "bits" in a block header which represent that target difficulty .
func ExampleBigToCompact() {
	// Convert the target difficulty from block 300000 in the bitcoin
	// main chain to compact form.
	t := "0000000000000000896c00000000000000000000000000000000000000000000"
	targetDifficulty, success := new(big.Int).SetString(t, 16)
	if !success {
		fmt.Println("invalid target difficulty")
		return
	}
	bits := difficulty.BigToCompact(targetDifficulty)

	fmt.Println(bits)

	// Output:
	// 419465580
}
