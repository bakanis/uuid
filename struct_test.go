/**
 * Date: 15/02/14
 * Time: 12:26 PM
 */
package uuid

import (
	"testing"
)

// Tests all possible variants with their respective bits set
// Tests whether the expected output comes out on each byte case
func TestUUIDStruct_VarientBits(t *testing.T) {
	for _, v := range uuidVariants {
		for i := 0; i <= 255; i ++ {
			uuidBytes[variantIndex] = byte(i)

			uStruct := createUUIDStruct(uuidBytes, 4, v)
			b := uStruct.sequenceHiAndVariant>>4
			t_VariantConstraint(v, b, uStruct, t)

			if uStruct.Variant() != v {
				t.Errorf("%d does not resolve to %x: get %x", i, v, uStruct.Variant())
			}
		}
	}
}

// Tests all possible version numbers and that
// each number returned is the same
func TestUUIDStruct_VersionBits(t *testing.T) {
	uStruct := new(UUIDStruct)
	for v := 0; v < 16; v++ {
		for i := 0; i <= 255; i ++ {
			uuidBytes[versionIndex] = byte(i)
			uStruct.Unmarshal(uuidBytes)
			uStruct.setVersion(v)
			output(uStruct)
			if uStruct.Version() != v {
				t.Errorf("%x does not resolve to %x", byte(uStruct.Version()), v)
			}
			output("\n")
		}
	}
}

func TestUUIDStruct_UnmarshalBinary(t *testing.T) {
	u := new(UUIDStruct)
	err := u.UnmarshalBinary([]byte{1, 2, 3, 4, 5})
	if err == nil {
		t.Errorf("Expected error due to invalid byte length")
	}
	err = u.UnmarshalBinary(uuidBytes)
	if err != nil {
		t.Errorf("Expected bytes")
	}
}

func createUUIDStruct(pData []byte, pVersion int, pVariant byte) *UUIDStruct {
	o := new(UUIDStruct)
	o.Unmarshal(pData)
	o.setVersion(pVersion)
	o.setVariant(pVariant)
	return o
}
