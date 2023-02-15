package ethutils

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"math/big"
	"testing"
)

func TestEthConversion(t *testing.T) {
	// Eth -> Wei
	fmt.Println("EtherToWei(big.NewFloat(1.0))")
	v := EtherToWei(big.NewFloat(1.0))
	fmt.Println(v.String())
	assert.Equal(t, v.String(), big.NewInt(1000000000000000000).String(), "Error: Not equal | EtherToWei(big.NewFloat(1.0))")
	fmt.Println()

	fmt.Println("EtherToWei(big.NewFloat(0.351))")
	v = EtherToWei(big.NewFloat(0.351))
	fmt.Println(v.String())
	assert.Equal(t, v.String(), big.NewInt(351000000000000000).String(), "Error: Not equal | EtherToWei(big.NewFloat(0.351))")
	fmt.Println()

	// Eth -> Gwei
	fmt.Println("EtherToGwei(big.NewFloat(1.0))")
	v = EtherToGwei(big.NewFloat(1.0))
	fmt.Println(v.String())
	assert.Equal(t, v.String(), big.NewInt(1000000000).String(), "EtherToGwei(big.NewFloat(1.0))")
	fmt.Println()

	fmt.Println("EtherToGwei(big.NewFloat(0.351))")
	v = EtherToGwei(big.NewFloat(0.351))
	fmt.Println(v.String())
	assert.Equal(t, v.String(), big.NewInt(351000000).String(), "EtherToGwei(big.NewFloat(0.351))")
	fmt.Println()
}

func TestGweiConversion(t *testing.T) {
	// Gwei -> Eth
	fmt.Println("GweiToEther(big.NewFloat(1000000000))")
	vf := GweiToEther(big.NewInt(1000000000))
	fmt.Println(vf.String())
	assert.Equal(t, vf.String(), big.NewFloat(1).String(), "GweiToEther(big.NewFloat(1000000000))")
	fmt.Println()

	fmt.Println("GweiToEther(big.NewFloat(351000000))")
	vf = GweiToEther(big.NewInt(351000000))
	fmt.Println(vf.String())
	assert.Equal(t, vf.String(), big.NewFloat(0.351).String(), "GweiToEther(big.NewFloat(351000000))")
	fmt.Println()

	// Gwei -> Wei
	fmt.Println("GweiToWei(big.NewFloat(1000000000))")
	vi := GweiToWei(big.NewInt(1000000000))
	fmt.Println(vi.String())
	assert.Equal(t, vi.String(), big.NewInt(1000000000000000000).String(), "GweiToWei(big.NewFloat(1000000000))")
	fmt.Println()

	fmt.Println("GweiToWei(big.NewFloat(351000000))")
	vi = GweiToWei(big.NewInt(351000000))
	fmt.Println(vi.String())
	assert.Equal(t, vi.String(), big.NewInt(351000000000000000).String(), "GweiToWei(big.NewFloat(351000000))")
	fmt.Println()
}

func TestWeiConversion(t *testing.T) {
	// Wei -> Eth
	fmt.Println("WeiToEther(big.NewFloat(1000000000000000000))")
	vf := WeiToEther(big.NewInt(1000000000000000000))
	fmt.Println(vf.String())
	assert.Equal(t, vf.String(), big.NewFloat(1.0).String(), "WeiToEther(big.NewFloat(1000000000000000000))")
	fmt.Println()

	fmt.Println("WeiToEther(big.NewFloat(351000000000000000))")
	vf = WeiToEther(big.NewInt(351000000000000000))
	fmt.Println(vf.String())
	assert.Equal(t, vf.String(), big.NewFloat(0.351).String(), "WeiToEther(big.NewFloat(351000000000000000))")
	fmt.Println()

	// Wei -> Gwei
	fmt.Println("WeiToGwei(big.NewFloat(1000000000000000000))")
	vi := WeiToGwei(big.NewInt(1000000000000000000))
	fmt.Println(vi.String())
	assert.Equal(t, vi.String(), big.NewInt(1000000000).String(), "WeiToGwei(big.NewFloat(1000000000000000000))")
	fmt.Println()

	fmt.Println("WeiToGwei(big.NewFloat(351000000000000000))")
	vi = WeiToGwei(big.NewInt(351000000000000000))
	fmt.Println(vi.String())
	assert.Equal(t, vi.String(), big.NewInt(351000000).String(), "WeiToGwei(big.NewFloat(351000000000000000))")
	fmt.Println()
}
