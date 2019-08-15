package main

import "testing"

func Test_DeliveryImg(t *testing.T) {
	name, src := ImgConvertToB64Str("testData/verify.jpg")
	B64StrSaveToImg(name, src)
}
