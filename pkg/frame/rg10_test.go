package frame

import (
	"image"
	"reflect"
	"testing"
)

func TestDecodeRG10(t *testing.T) {
	const (
		width  = 4
		height = 1
	)
	input := []byte{0x00, 0x00, 0x00, 0x00}
	expected := image.NewYCbCr(image.Rect(0, 0, 0, 0), image.YCbCrSubsampleRatio420)
	decoder, err := NewDecoder(FormatRG10)
	if err != nil {
		t.Fatal(err)
	}
	img, _, err := decoder.Decode(input, width, height)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(expected, img) {
		t.Errorf("Wrong decode result,\nexpected:\n%+v\ngot:\n%+v", expected, img)
	}
}
