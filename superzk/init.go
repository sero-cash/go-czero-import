package superzk

import "C"
import (
	"github.com/sero-cash/go-czero-import/c_superzk"
	"github.com/sero-cash/go-czero-import/c_type"
)

func ZeroInit(account_dir string, netType c_type.NetType) error {
	c_superzk.InitParams()
	return nil
}

func ZeroInit_NoCircuit() error {
	c_superzk.InitParams_NoCircuit()
	return nil
}

func ZeroInit_OnlyInOuts() error {
	c_superzk.InitParams()
	return nil
}
