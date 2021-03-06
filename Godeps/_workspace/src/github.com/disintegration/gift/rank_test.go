package gift

import (
	"image"
	"testing"
)

func TestMedian(t *testing.T) {
	testData := []struct {
		desc           string
		ksize          int
		disk           bool
		srcb, dstb     image.Rectangle
		srcPix, dstPix []uint8
	}{
		{
			"median (0, false)",
			0, false,
			image.Rect(-1, -1, 4, 2),
			image.Rect(0, 0, 5, 3),
			[]uint8{
				0x11, 0x99, 0x55, 0x22, 0x66,
				0xFF, 0x44, 0xFF, 0xCC, 0x00,
				0x33, 0x77, 0xBB, 0x88, 0xEE,
			},
			[]uint8{
				0x11, 0x99, 0x55, 0x22, 0x66,
				0xFF, 0x44, 0xFF, 0xCC, 0x00,
				0x33, 0x77, 0xBB, 0x88, 0xEE,
			},
		},
		{
			"median (1, false)",
			1, false,
			image.Rect(-1, -1, 4, 2),
			image.Rect(0, 0, 5, 3),
			[]uint8{
				0x11, 0x99, 0x55, 0x22, 0x66,
				0xFF, 0x44, 0xFF, 0xCC, 0x00,
				0x33, 0x77, 0xBB, 0x88, 0xEE,
			},
			[]uint8{
				0x11, 0x99, 0x55, 0x22, 0x66,
				0xFF, 0x44, 0xFF, 0xCC, 0x00,
				0x33, 0x77, 0xBB, 0x88, 0xEE,
			},
		},
		{
			"median (2, false)",
			2, false,
			image.Rect(-1, -1, 4, 2),
			image.Rect(0, 0, 5, 3),
			[]uint8{
				0x11, 0x99, 0x55, 0x22, 0x66,
				0xFF, 0x44, 0xFF, 0xCC, 0x00,
				0x33, 0x77, 0xBB, 0x88, 0xEE,
			},
			[]uint8{
				0x11, 0x99, 0x55, 0x22, 0x66,
				0xFF, 0x44, 0xFF, 0xCC, 0x00,
				0x33, 0x77, 0xBB, 0x88, 0xEE,
			},
		},
		{
			"median (3, false)",
			3, false,
			image.Rect(-1, -1, 4, 2),
			image.Rect(0, 0, 5, 3),
			[]uint8{
				0x11, 0x99, 0x55, 0x22, 0x66,
				0xFF, 0x44, 0xFF, 0xCC, 0x00,
				0x33, 0x77, 0xBB, 0x88, 0xEE,
			},
			[]uint8{
				0x44, 0x55, 0x55, 0x55, 0x66,
				0x44, 0x77, 0x88, 0x88, 0x66,
				0x44, 0x77, 0x88, 0xBB, 0xCC,
			},
		},
		{
			"median (3, true)",
			3, true,
			image.Rect(-1, -1, 4, 2),
			image.Rect(0, 0, 5, 3),
			[]uint8{
				0x11, 0x99, 0x55, 0x22, 0x66,
				0xFF, 0x44, 0xFF, 0xCC, 0x00,
				0x33, 0x77, 0xBB, 0x88, 0xEE,
			},
			[]uint8{
				0x11, 0x55, 0x55, 0x55, 0x66,
				0x44, 0x99, 0xBB, 0x88, 0x66,
				0x33, 0x77, 0xBB, 0xBB, 0xEE,
			},
		},
		{
			"median (4, true)",
			4, true,
			image.Rect(-1, -1, 4, 2),
			image.Rect(0, 0, 5, 3),
			[]uint8{
				0x11, 0x99, 0x55, 0x22, 0x66,
				0xFF, 0x44, 0xFF, 0xCC, 0x00,
				0x33, 0x77, 0xBB, 0x88, 0xEE,
			},
			[]uint8{
				0x11, 0x55, 0x55, 0x55, 0x66,
				0x44, 0x99, 0xBB, 0x88, 0x66,
				0x33, 0x77, 0xBB, 0xBB, 0xEE,
			},
		},
	}

	for _, d := range testData {
		src := image.NewGray(d.srcb)
		src.Pix = d.srcPix

		f := Median(d.ksize, d.disk)
		dst := image.NewGray(f.Bounds(src.Bounds()))
		f.Draw(dst, src, nil)

		if !checkBoundsAndPix(dst.Bounds(), d.dstb, dst.Pix, d.dstPix) {
			t.Errorf("test [%s] failed: %#v, %#v", d.desc, dst.Bounds(), dst.Pix)
		}
	}

	testDataNRGBA := []struct {
		desc           string
		ksize          int
		disk           bool
		srcb, dstb     image.Rectangle
		srcPix, dstPix []uint8
	}{
		{
			"median nrgba (3, true)",
			3, true,
			image.Rect(-1, -1, 4, 2),
			image.Rect(0, 0, 5, 3),
			[]uint8{
				0x00, 0x00, 0x00, 0x11, 0x00, 0x00, 0x00, 0x99, 0x00, 0x00, 0x00, 0x55, 0x00, 0x00, 0x00, 0x22, 0x00, 0x00, 0x00, 0x66,
				0x00, 0x00, 0x00, 0xFF, 0x00, 0x00, 0x00, 0x44, 0x00, 0x00, 0x00, 0xFF, 0x00, 0x00, 0x00, 0xCC, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x33, 0x00, 0x00, 0x00, 0x77, 0x00, 0x00, 0x00, 0xBB, 0x00, 0x00, 0x00, 0x88, 0x00, 0x00, 0x00, 0xEE,
			},
			[]uint8{
				0x00, 0x00, 0x00, 0x11, 0x00, 0x00, 0x00, 0x55, 0x00, 0x00, 0x00, 0x55, 0x00, 0x00, 0x00, 0x55, 0x00, 0x00, 0x00, 0x66,
				0x00, 0x00, 0x00, 0x44, 0x00, 0x00, 0x00, 0x99, 0x00, 0x00, 0x00, 0xBB, 0x00, 0x00, 0x00, 0x88, 0x00, 0x00, 0x00, 0x66,
				0x00, 0x00, 0x00, 0x33, 0x00, 0x00, 0x00, 0x77, 0x00, 0x00, 0x00, 0xBB, 0x00, 0x00, 0x00, 0xBB, 0x00, 0x00, 0x00, 0xEE,
			},
		},
	}

	for _, d := range testDataNRGBA {
		src := image.NewNRGBA(d.srcb)
		src.Pix = d.srcPix

		f := Median(d.ksize, d.disk)
		dst := image.NewNRGBA(f.Bounds(src.Bounds()))
		f.Draw(dst, src, nil)

		if !checkBoundsAndPix(dst.Bounds(), d.dstb, dst.Pix, d.dstPix) {
			t.Errorf("test [%s] failed: %#v, %#v", d.desc, dst.Bounds(), dst.Pix)
		}
	}

	// check no panics
	Median(5, false).Draw(image.NewGray(image.Rect(0, 0, 1, 1)), image.NewGray(image.Rect(0, 0, 1, 1)), nil)
	Median(5, false).Draw(image.NewGray(image.Rect(0, 0, 0, 0)), image.NewGray(image.Rect(0, 0, 0, 0)), nil)
}

func TestMinimum(t *testing.T) {
	testData := []struct {
		desc           string
		ksize          int
		disk           bool
		srcb, dstb     image.Rectangle
		srcPix, dstPix []uint8
	}{
		{
			"minimum (0, false)",
			0, false,
			image.Rect(-1, -1, 4, 2),
			image.Rect(0, 0, 5, 3),
			[]uint8{
				0x11, 0x99, 0x55, 0x22, 0x66,
				0xFF, 0x44, 0xFF, 0xCC, 0x00,
				0x33, 0x77, 0xBB, 0x88, 0xEE,
			},
			[]uint8{
				0x11, 0x99, 0x55, 0x22, 0x66,
				0xFF, 0x44, 0xFF, 0xCC, 0x00,
				0x33, 0x77, 0xBB, 0x88, 0xEE,
			},
		},
		{
			"minimum (1, false)",
			1, false,
			image.Rect(-1, -1, 4, 2),
			image.Rect(0, 0, 5, 3),
			[]uint8{
				0x11, 0x99, 0x55, 0x22, 0x66,
				0xFF, 0x44, 0xFF, 0xCC, 0x00,
				0x33, 0x77, 0xBB, 0x88, 0xEE,
			},
			[]uint8{
				0x11, 0x99, 0x55, 0x22, 0x66,
				0xFF, 0x44, 0xFF, 0xCC, 0x00,
				0x33, 0x77, 0xBB, 0x88, 0xEE,
			},
		},
		{
			"minimum (2, false)",
			2, false,
			image.Rect(-1, -1, 4, 2),
			image.Rect(0, 0, 5, 3),
			[]uint8{
				0x11, 0x99, 0x55, 0x22, 0x66,
				0xFF, 0x44, 0xFF, 0xCC, 0x00,
				0x33, 0x77, 0xBB, 0x88, 0xEE,
			},
			[]uint8{
				0x11, 0x99, 0x55, 0x22, 0x66,
				0xFF, 0x44, 0xFF, 0xCC, 0x00,
				0x33, 0x77, 0xBB, 0x88, 0xEE,
			},
		},
		{
			"minimum (3, false)",
			3, false,
			image.Rect(-1, -1, 4, 2),
			image.Rect(0, 0, 5, 3),
			[]uint8{
				0x11, 0x99, 0x55, 0x22, 0x66,
				0xFF, 0x44, 0xFF, 0xCC, 0x00,
				0x33, 0x77, 0xBB, 0x88, 0xEE,
			},
			[]uint8{
				0x11, 0x11, 0x22, 0x00, 0x00,
				0x11, 0x11, 0x22, 0x00, 0x00,
				0x33, 0x33, 0x44, 0x00, 0x00,
			},
		},
		{
			"minimum (3, true)",
			3, true,
			image.Rect(-1, -1, 4, 2),
			image.Rect(0, 0, 5, 3),
			[]uint8{
				0x11, 0x99, 0x55, 0x22, 0x66,
				0xFF, 0x44, 0xFF, 0xCC, 0x00,
				0x33, 0x77, 0xBB, 0x88, 0xEE,
			},
			[]uint8{
				0x11, 0x11, 0x22, 0x22, 0x00,
				0x11, 0x44, 0x44, 0x00, 0x00,
				0x33, 0x33, 0x77, 0x88, 0x00,
			},
		},
		{
			"minimum (4, true)",
			4, true,
			image.Rect(-1, -1, 4, 2),
			image.Rect(0, 0, 5, 3),
			[]uint8{
				0x11, 0x99, 0x55, 0x22, 0x66,
				0xFF, 0x44, 0xFF, 0xCC, 0x00,
				0x33, 0x77, 0xBB, 0x88, 0xEE,
			},
			[]uint8{
				0x11, 0x11, 0x22, 0x22, 0x00,
				0x11, 0x44, 0x44, 0x00, 0x00,
				0x33, 0x33, 0x77, 0x88, 0x00,
			},
		},
	}

	for _, d := range testData {
		src := image.NewGray(d.srcb)
		src.Pix = d.srcPix

		f := Minimum(d.ksize, d.disk)
		dst := image.NewGray(f.Bounds(src.Bounds()))
		f.Draw(dst, src, nil)

		if !checkBoundsAndPix(dst.Bounds(), d.dstb, dst.Pix, d.dstPix) {
			t.Errorf("test [%s] failed: %#v, %#v", d.desc, dst.Bounds(), dst.Pix)
		}
	}

	testDataNRGBA := []struct {
		desc           string
		ksize          int
		disk           bool
		srcb, dstb     image.Rectangle
		srcPix, dstPix []uint8
	}{
		{
			"minimum nrgba (3, true)",
			3, true,
			image.Rect(-1, -1, 4, 2),
			image.Rect(0, 0, 5, 3),
			[]uint8{
				0x00, 0x00, 0x00, 0x11, 0x00, 0x00, 0x00, 0x99, 0x00, 0x00, 0x00, 0x55, 0x00, 0x00, 0x00, 0x22, 0x00, 0x00, 0x00, 0x66,
				0x00, 0x00, 0x00, 0xFF, 0x00, 0x00, 0x00, 0x44, 0x00, 0x00, 0x00, 0xFF, 0x00, 0x00, 0x00, 0xCC, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x33, 0x00, 0x00, 0x00, 0x77, 0x00, 0x00, 0x00, 0xBB, 0x00, 0x00, 0x00, 0x88, 0x00, 0x00, 0x00, 0xEE,
			},
			[]uint8{
				0x00, 0x00, 0x00, 0x11, 0x00, 0x00, 0x00, 0x11, 0x00, 0x00, 0x00, 0x22, 0x00, 0x00, 0x00, 0x22, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x11, 0x00, 0x00, 0x00, 0x44, 0x00, 0x00, 0x00, 0x44, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x33, 0x00, 0x00, 0x00, 0x33, 0x00, 0x00, 0x00, 0x77, 0x00, 0x00, 0x00, 0x88, 0x00, 0x00, 0x00, 0x00,
			},
		},
	}

	for _, d := range testDataNRGBA {
		src := image.NewNRGBA(d.srcb)
		src.Pix = d.srcPix

		f := Minimum(d.ksize, d.disk)
		dst := image.NewNRGBA(f.Bounds(src.Bounds()))
		f.Draw(dst, src, nil)

		if !checkBoundsAndPix(dst.Bounds(), d.dstb, dst.Pix, d.dstPix) {
			t.Errorf("test [%s] failed: %#v, %#v", d.desc, dst.Bounds(), dst.Pix)
		}
	}
}

func TestMaximum(t *testing.T) {
	testData := []struct {
		desc           string
		ksize          int
		disk           bool
		srcb, dstb     image.Rectangle
		srcPix, dstPix []uint8
	}{
		{
			"maximum (0, false)",
			0, false,
			image.Rect(-1, -1, 4, 2),
			image.Rect(0, 0, 5, 3),
			[]uint8{
				0x11, 0x99, 0x55, 0x22, 0x66,
				0xFF, 0x44, 0xFF, 0xCC, 0x00,
				0x33, 0x77, 0xBB, 0x88, 0xEE,
			},
			[]uint8{
				0x11, 0x99, 0x55, 0x22, 0x66,
				0xFF, 0x44, 0xFF, 0xCC, 0x00,
				0x33, 0x77, 0xBB, 0x88, 0xEE,
			},
		},
		{
			"maximum (1, false)",
			1, false,
			image.Rect(-1, -1, 4, 2),
			image.Rect(0, 0, 5, 3),
			[]uint8{
				0x11, 0x99, 0x55, 0x22, 0x66,
				0xFF, 0x44, 0xFF, 0xCC, 0x00,
				0x33, 0x77, 0xBB, 0x88, 0xEE,
			},
			[]uint8{
				0x11, 0x99, 0x55, 0x22, 0x66,
				0xFF, 0x44, 0xFF, 0xCC, 0x00,
				0x33, 0x77, 0xBB, 0x88, 0xEE,
			},
		},
		{
			"maximum (2, false)",
			2, false,
			image.Rect(-1, -1, 4, 2),
			image.Rect(0, 0, 5, 3),
			[]uint8{
				0x11, 0x99, 0x55, 0x22, 0x66,
				0xFF, 0x44, 0xFF, 0xCC, 0x00,
				0x33, 0x77, 0xBB, 0x88, 0xEE,
			},
			[]uint8{
				0x11, 0x99, 0x55, 0x22, 0x66,
				0xFF, 0x44, 0xFF, 0xCC, 0x00,
				0x33, 0x77, 0xBB, 0x88, 0xEE,
			},
		},
		{
			"maximum (3, false)",
			3, false,
			image.Rect(-1, -1, 4, 2),
			image.Rect(0, 0, 5, 3),
			[]uint8{
				0x11, 0x99, 0x55, 0x22, 0x66,
				0xFF, 0x44, 0xFF, 0xCC, 0x00,
				0x33, 0x77, 0xBB, 0x88, 0xEE,
			},
			[]uint8{
				0xFF, 0xFF, 0xFF, 0xFF, 0xCC,
				0xFF, 0xFF, 0xFF, 0xFF, 0xEE,
				0xFF, 0xFF, 0xFF, 0xFF, 0xEE,
			},
		},
		{
			"maximum (3, true)",
			3, true,
			image.Rect(-1, -1, 4, 2),
			image.Rect(0, 0, 5, 3),
			[]uint8{
				0x11, 0x99, 0x55, 0x22, 0x66,
				0xFF, 0x44, 0xFF, 0xCC, 0x00,
				0x33, 0x77, 0xBB, 0x88, 0xEE,
			},
			[]uint8{
				0xFF, 0x99, 0xFF, 0xCC, 0x66,
				0xFF, 0xFF, 0xFF, 0xFF, 0xEE,
				0xFF, 0xBB, 0xFF, 0xEE, 0xEE,
			},
		},
		{
			"maximum (4, true)",
			4, true,
			image.Rect(-1, -1, 4, 2),
			image.Rect(0, 0, 5, 3),
			[]uint8{
				0x11, 0x99, 0x55, 0x22, 0x66,
				0xFF, 0x44, 0xFF, 0xCC, 0x00,
				0x33, 0x77, 0xBB, 0x88, 0xEE,
			},
			[]uint8{
				0xFF, 0x99, 0xFF, 0xCC, 0x66,
				0xFF, 0xFF, 0xFF, 0xFF, 0xEE,
				0xFF, 0xBB, 0xFF, 0xEE, 0xEE,
			},
		},
	}

	for _, d := range testData {
		src := image.NewGray(d.srcb)
		src.Pix = d.srcPix

		f := Maximum(d.ksize, d.disk)
		dst := image.NewGray(f.Bounds(src.Bounds()))
		f.Draw(dst, src, nil)

		if !checkBoundsAndPix(dst.Bounds(), d.dstb, dst.Pix, d.dstPix) {
			t.Errorf("test [%s] failed: %#v, %#v", d.desc, dst.Bounds(), dst.Pix)
		}
	}

	testDataNRGBA := []struct {
		desc           string
		ksize          int
		disk           bool
		srcb, dstb     image.Rectangle
		srcPix, dstPix []uint8
	}{
		{
			"maximum nrgba (3, true)",
			3, true,
			image.Rect(-1, -1, 4, 2),
			image.Rect(0, 0, 5, 3),
			[]uint8{
				0x00, 0x00, 0x00, 0x11, 0x00, 0x00, 0x00, 0x99, 0x00, 0x00, 0x00, 0x55, 0x00, 0x00, 0x00, 0x22, 0x00, 0x00, 0x00, 0x66,
				0x00, 0x00, 0x00, 0xFF, 0x00, 0x00, 0x00, 0x44, 0x00, 0x00, 0x00, 0xFF, 0x00, 0x00, 0x00, 0xCC, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x33, 0x00, 0x00, 0x00, 0x77, 0x00, 0x00, 0x00, 0xBB, 0x00, 0x00, 0x00, 0x88, 0x00, 0x00, 0x00, 0xEE,
			},
			[]uint8{
				0x00, 0x00, 0x00, 0xFF, 0x00, 0x00, 0x00, 0x99, 0x00, 0x00, 0x00, 0xFF, 0x00, 0x00, 0x00, 0xCC, 0x00, 0x00, 0x00, 0x66,
				0x00, 0x00, 0x00, 0xFF, 0x00, 0x00, 0x00, 0xFF, 0x00, 0x00, 0x00, 0xFF, 0x00, 0x00, 0x00, 0xFF, 0x00, 0x00, 0x00, 0xEE,
				0x00, 0x00, 0x00, 0xFF, 0x00, 0x00, 0x00, 0xBB, 0x00, 0x00, 0x00, 0xFF, 0x00, 0x00, 0x00, 0xEE, 0x00, 0x00, 0x00, 0xEE,
			},
		},
	}

	for _, d := range testDataNRGBA {
		src := image.NewNRGBA(d.srcb)
		src.Pix = d.srcPix

		f := Maximum(d.ksize, d.disk)
		dst := image.NewNRGBA(f.Bounds(src.Bounds()))
		f.Draw(dst, src, nil)

		if !checkBoundsAndPix(dst.Bounds(), d.dstb, dst.Pix, d.dstPix) {
			t.Errorf("test [%s] failed: %#v, %#v", d.desc, dst.Bounds(), dst.Pix)
		}
	}
}
