/**
 * Date: 15/02/14
 * Time: 12:49 PM
 */
package uuid

import (
	"testing"
)

// Tests all possible variants with their respective bits set
// Tests whether the expected output comes out on each byte case
func TestUUIDArray_VarientBits(t *testing.T) {
	for _, v := range uuidVariants {
		for i := 0; i <= 255; i ++ {
			uuidBytes[variantIndex] = byte(i)

			uArray := createUUIDArray(uuidBytes, 4, v)
			b := uArray[variantIndex]>>4
			t_VariantConstraint(v, b, uArray, t)

			if uArray.Variant() != v {
				t.Errorf("%d does not resolve to %x", i, v)
			}
		}
	}
}

// Tests all possible version numbers and that
// each number returned is the same
func TestUUIDArray_VersionBits(t *testing.T) {
	uArray := new(UUIDArray)
	for v := 0; v < 16; v++ {
		for i := 0; i <= 255; i ++ {
			uuidBytes[versionIndex] = byte(i)
			uArray.Unmarshal(uuidBytes)
			uArray.setVersion(v)
			output(uArray)
			if uArray.Version() != v {
				t.Errorf("%x does not resolve to %x", byte(uArray.Version()), v)
			}
			output("\n")
		}
	}
}

func TestUUIDArray_UnmarshalBinary(t *testing.T) {
	u := new(UUIDArray)
	err := u.UnmarshalBinary([]byte{1, 2, 3, 4, 5})
	if err == nil {
		t.Errorf("Expected error due to invalid byte length")
	}
	err = u.UnmarshalBinary(uuidBytes)
	if err != nil {
		t.Errorf("Expected bytes")
	}
}

func createUUIDArray(pData []byte, pVersion int, pVariant byte) *UUIDArray {
	o := new(UUIDArray)
	o.Unmarshal(pData)
	o.setVersion(pVersion)
	o.setVariant(pVariant)
	return o
}


