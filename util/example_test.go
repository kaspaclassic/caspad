package util_test

import (
	"fmt"
	"math"
	"math/big"

	"github.com/kaspaclassic/caspad/util/difficulty"

	"github.com/kaspaclassic/caspad/util"
)

func ExampleAmount() {

	a := util.Amount(0)
	fmt.Println("Zero Leor:", a)

	a = util.Amount(1e8)
	fmt.Println("100,000,000 Leor:", a)

	a = util.Amount(1e5)
	fmt.Println("100,000 Leor:", a)
	// Output:
	// Zero Leor: 0 CAS
	// 100,000,000 Leor: 1 CAS
	// 100,000 Leor: 0.001 CAS
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

	// Output: 1 CAS
	// 0.01234567 CAS
	// 0 CAS
	// invalid caspa amount
}

func ExampleAmount_unitConversions() {
	amount := util.Amount(44433322211100)

	fmt.Println("Leor to kCAS:", amount.Format(util.AmountKiloCAS))
	fmt.Println("Leor to CAS:", amount)
	fmt.Println("Leor to MilliCAS:", amount.Format(util.AmountMilliCAS))
	fmt.Println("Leor to MicroCAS:", amount.Format(util.AmountMicroCAS))
	fmt.Println("Leor to Leor:", amount.Format(util.AmountLeor))

	// Output:
	// Leor to kCAS: 444.333222111 kCAS
	// Leor to CAS: 444333.222111 CAS
	// Leor to MilliCAS: 444333222.111 mCAS
	// Leor to MicroCAS: 444333222111 μCAS
	// Leor to Leor: 44433322211100 Leor
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
