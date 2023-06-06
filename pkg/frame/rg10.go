package frame

import "image"

func decodeRG10(frame []byte, width, height int) (image.Image, func(), error) {
	return image.NewYCbCr(image.Rect(0, 0, 0, 0), image.YCbCrSubsampleRatio420), nil, nil
}
