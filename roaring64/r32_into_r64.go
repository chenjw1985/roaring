package roaring64

import "github.com/RoaringBitmap/roaring"

func CreateR64FromR32Slice(r32 ...*roaring.Bitmap) *Bitmap {
	if len(r32) == 0 {
		return NewBitmap()
	}
	size := len(r32)
	rb := NewBitmap()
	rb.highlowcontainer = roaringArray64{}
	rb.highlowcontainer.keys = make([]uint32, size)
	rb.highlowcontainer.containers = make([]*roaring.Bitmap, size)
	rb.highlowcontainer.needCopyOnWrite = make([]bool, size)
	for k, v := range r32 {
		rb.highlowcontainer.keys[k] = uint32(k)
		rb.highlowcontainer.containers[k] = v
		rb.highlowcontainer.needCopyOnWrite[k] = false
	}
	return rb
}
