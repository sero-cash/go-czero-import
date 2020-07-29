package c_superzk

/*

#include "csuperzk.h"

*/
import "C"
import (
	"errors"
	"fmt"
	"unsafe"

	"github.com/sero-cash/go-czero-import/c_type"
)

func Tk2Pk(tk *c_type.Tk) (pk c_type.Uint512, e error) {
	tk = ClearTk(tk)
	ret := C.superzk_tk2pk(
		(*C.uchar)(unsafe.Pointer(&tk[0])),
		(*C.uchar)(unsafe.Pointer(&pk[0])),
	)
	SetFlag(pk[:])
	if ret != C.int(0) {
		e = errors.New("tk2pk error")
		return
	}
	return
}

func Pk2PKr(addr *c_type.Uint512, r *c_type.Uint256) (pkr c_type.PKr, e error) {
	assertPK(addr)
	addr = ClearPK(addr)
	if r == nil {
		t := c_type.RandUint256()
		r = &t
	} else {
		if (*r) == c_type.Empty_Uint256 {
			panic("gen pkr, but r is empty")
		}
	}
	ret := C.superzk_pk2pkr(
		(*C.uchar)(unsafe.Pointer(&addr[0])),
		(*C.uchar)(unsafe.Pointer(&r[0])),
		(*C.uchar)(unsafe.Pointer(&pkr[0])),
	)
	if ret != C.int(0) {
		e = errors.New("pk2pkr error")
		return
	}
	SetFlag(pkr[:])
	return
}

func IsPKValid(pk *c_type.Uint512) bool {
	if !IsSzkPK(pk) {
		return false
	}
	pk = ClearPK(pk)
	ret := C.superzk_pk_valid(
		(*C.uchar)(unsafe.Pointer(&pk[0])),
	)
	if ret != C.int(0) {
		return false
	}
	return true
}

func IsPKrValid(pkr *c_type.PKr) bool {
	if !IsSzkPKr(pkr) {
		return false
	}
	pkr = ClearPKr(pkr)
	ret := C.superzk_pkr_valid(
		(*C.uchar)(unsafe.Pointer(&pkr[0])),
	)
	if ret != C.int(0) {
		return false
	}
	return true
}

func IsMyPKr(tk *c_type.Tk, pkr *c_type.PKr) bool {
	assertPKr(pkr)
	tk = ClearTk(tk)
	pkr = ClearPKr(pkr)
	ret := C.superzk_my_pkr(
		(*C.uchar)(unsafe.Pointer(&tk[0])),
		(*C.uchar)(unsafe.Pointer(&pkr[0])),
	)
	if ret != C.int(0) {
		return false
	}
	return true
}

func FetchRootCM(tk *c_type.Tk, nl *c_type.Uint256, baser *c_type.Uint256) (root_cm c_type.Uint256, e error) {
	if !IsSzkTk(tk) {
		e = errors.New("csuperzk fetch rootcm error: tk is not szk")
		return
	}
	if !IsSzkNil(nl) {
		e = errors.New("csuperzk fetch rootcm error: nil is not szk")
		return
	}
	tk = ClearTk(tk)
	nl = ClearNil(nl)
	baser = ClearNil(baser)
	ret := C.superzk_nil2cm(
		(*C.uchar)(unsafe.Pointer(&tk[0])),
		(*C.uchar)(unsafe.Pointer(&nl[0])),
		(*C.uchar)(unsafe.Pointer(&baser[0])),
		(*C.uchar)(unsafe.Pointer(&root_cm[0])),
	)
	if ret != C.int(0) {
		e = fmt.Errorf("csuperzk fetch rootcm error: %d", int(ret))
		return
	}
	return
}

func PkrCrypte(data *c_type.Uint256, myPKr *c_type.PKr, myTk *c_type.Tk, youPKr *c_type.PKr) (dataEnc c_type.Uint256, e error) {
	if !IsSzkPKr(myPKr) {
		e = errors.New("csuperzk pkr crypte error: myPKr is not szk")
		return
	}
	if !IsSzkTk(myTk) {
		e = errors.New("csuperzk pkr crypte error: myTk is not szk")
		return
	}
	if !IsSzkPKr(youPKr) {
		e = errors.New("csuperzk pkr crypte error: youPKr is not szk")
		return
	}
	myPKr = ClearPKr(myPKr)
	myTk = ClearTk(myTk)
	youPKr = ClearPKr(youPKr)
	ret := C.superzk_pkr_crypte(
		(*C.uchar)(unsafe.Pointer(&data[0])),
		(*C.uchar)(unsafe.Pointer(&myPKr[0])),
		(*C.uchar)(unsafe.Pointer(&myTk[0])),
		(*C.uchar)(unsafe.Pointer(&youPKr[0])),
		(*C.uchar)(unsafe.Pointer(&dataEnc[0])),
	)
	if ret != C.int(0) {
		e = fmt.Errorf("csuperzk pkr crypte error: %d", int(ret))
		return
	}
	return
}
