/**
 * Date: 3/02/14
 * Time: 10:59 PM
 */
package uuid

import (
	"crypto/md5"
	"crypto/sha1"
	"fmt"
	"testing"
)

const (
	secondsPerHour       = 60*60
	secondsPerDay        = 24*secondsPerHour
	unixToInternal int64 = (1969*365 + 1969/4 - 1969/100 + 1969/400)*secondsPerDay
	internalToUnix int64 = -unixToInternal
)

var (
	printer = true

	uuidBytes      = []byte{
	0xAA, 0xCF, 0xEE, 0x12,
	0xD4, 0x00,
	0x27, 0x23,
	0x00,
	0xD3,
	0x23, 0x12, 0x4A, 0x11, 0x89, 0xFF,
}
	uuidVariants   = []byte{
	ReservedNCS, ReservedRFC4122, ReservedMicrosoft, ReservedFuture,
}
	namespaceUuids = []UUID {
	NamespaceDNS, NamespaceURL, NamespaceOID, NamespaceX500,
}

	invalidHexStrings = [...]string{
	"foo",
	"6ba7b814-9dad-11d1-80b4-",
	"6ba7b814--9dad-11d1-80b4--00c04fd430c8",
	"6ba7b814-9dad7-11d1-80b4-00c04fd430c8999",
	"{6ba7b814-9dad-1180b4-00c04fd430c8",
	"{6ba7b814--11d1-80b4-00c04fd430c8}",
	"urn:uuid:6ba7b814-9dad-1666666680b4-00c04fd430c8",
}
	validHexStrings   = [...]string{
	"6ba7b8149dad-11d1-80b4-00c04fd430c8}",
	"{6ba7b8149dad-11d1-80b400c04fd430c8}",
	"{6ba7b814-9dad11d180b400c04fd430c8}",
	"6ba7b8149dad-11d1-80b4-00c04fd430c8",
	"6ba7b814-9dad11d1-80b4-00c04fd430c8",
	"6ba7b814-9dad-11d180b4-00c04fd430c8",
	"6ba7b814-9dad-11d1-80b400c04fd430c8",
	"6ba7b8149dad11d180b400c04fd430c8",
	"6ba7b814-9dad-11d1-80b4-00c04fd430c8",
	"{6ba7b814-9dad-11d1-80b4-00c04fd430c8}",
	"{6ba7b814-9dad-11d1-80b4-00c04fd430c8",
	"6ba7b814-9dad-11d1-80b4-00c04fd430c8}",
	"(6ba7b814-9dad-11d1-80b4-00c04fd430c8)",
	"urn:uuid:6ba7b814-9dad-11d1-80b4-00c04fd430c8",
}
)

func TestUUIDNew(t *testing.T) {
	u := New(uuidBytes)
	if u == nil {
		t.Error("Expected a valid UUID")
	}
	if u.Version() != 2 {
		t.Errorf("Expected correct version %d, but got %d", 2, u.Version())
	}
	if u.Variant() != ReservedNCS {
		t.Errorf("Expected ReservedNCS variant %x, but got %x", ReservedNCS, u.Variant())
	}
	if !parseUUIDRegex.MatchString(u.String()) {
		t.Errorf("Expected string representation to be valid, given: %s", u.String())
	}
}

func TestUUID_NewBulk(t *testing.T) {
	for i := 0; i < 1000000; i++ {
		New(uuidBytes)
	}
}

func TestUUID_GoIdBulk(t *testing.T) {
	for i := 0; i < 100; i++ {
		u := GoId(NamespaceX500, Name(""), md5.New())
		SwitchFormat(GoIdFormat, false)
		outputln(u)
	}
}

func TestUUIDNewHex(t *testing.T) {
	s := "f3593cffee9240df408687825b523f13"
	u := NewHex(s)
	if u == nil {
		t.Error("Expected a valid UUID")
	}
	if u.Version() != 4 {
		t.Errorf("Expected correct version %d, but got %d", 4, u.Version())
	}
	if u.Variant() != ReservedNCS {
		t.Errorf("Expected ReservedNCS variant %x, but got %x", ReservedNCS, u.Variant())
	}
	if !parseUUIDRegex.MatchString(u.String()) {
		t.Errorf("Expected string representation to be valid, given: %s", u.String())
	}
}

func TestUUID_NewHexBulk(t *testing.T) {
	for i := 0; i < 1000000; i++ {
		s := "f3593cffee9240df408687825b523f13"
		NewHex(s)
	}
}

// Tests PasreUUID
func TestUUIDParseUUID(t *testing.T) {
	for _, v := range invalidHexStrings {
		_, err := ParseUUID(v)
		if err == nil {
			t.Error("Expected error due to invalid UUID string:", v)
		}
	}
	for _, v := range validHexStrings {
		_, err := ParseUUID(v)
		if err != nil {
			t.Error("Expected valid UUID string but got error:", v)
		}
	}
	for _, id := range namespaceUuids {
		id, err := ParseUUID(id.String())
		if err != nil {
			t.Error("Expected valid UUID string but got error:", id)
		}
	}
}

func TestUUIDSum(t *testing.T) {
	u := new(UUIDArray)
	Digest(u, NamespaceDNS, goLang, md5.New())
	if u.Bytes() == nil {
		t.Error("Expected new data in bytes")
	}
	output(u.Bytes())
	u = new(UUIDArray)
	Digest(u, NamespaceDNS, goLang, sha1.New())
	if u.Bytes() == nil {
		t.Error("Expected new data in bytes")
	}
	output(u.Bytes())
}

func BenchmarkParseUUID(b *testing.B) {
	s := "f3593cff-ee92-40df-4086-87825b523f13"
	for i := 0; i < b.N; i++ {
		_, err := ParseUUID(s)
		if err != nil {
			b.Fatal(err)
		}
	}
	b.StopTimer()
	b.ReportAllocs()
}

// *******************************************************

func t_VariantConstraint(v byte, b byte, o UUID, t *testing.T) {
	output(o)
	switch v {
	case ReservedNCS:
		switch b {
		case 0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07:
			outputf(": %X ", b)
			break
		default: t.Errorf("%X most high bits do not resolve to 0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07", b)
		}
	case ReservedRFC4122:
		switch b {
		case 0x08, 0x09, 0x0A, 0x0B:
			outputf(": %X ", b)
			break
		default: t.Errorf("%X most high bits do not resolve to 0x08, 0x09, 0x0A, 0x0B", b)
		}
	case ReservedMicrosoft:
		switch b {
		case 0x0C, 0x0D:
			outputf(": %X ", b)
			break
		default: t.Errorf("%X most high bits do not resolve to 0x0C, 0x0D", b)
		}
	case ReservedFuture:
		switch b {
		case 0x0E, 0x0F:
			outputf(": %X ", b)
			break
		default: t.Errorf("%X most high bits do not resolve to 0x0E, 0x0F", b)
		}
	}
	output("\n")
}

func output(a... interface{}) {
	if printer {
		fmt.Print(a...)
	}
}

func outputln(a... interface{}) {
	if printer {
		fmt.Println(a...)
	}
}

func outputf(format string, a... interface{}) {
	if printer {
		fmt.Printf(format, a)
	}
}


