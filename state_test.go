/**
 * Date: 14/02/14
 * Time: 9:08 PM
 */
package uuid

import (
	"bytes"
	"testing"
	"time"
)

func TestUUID_StateSeed(t *testing.T) {
	if state.past < Timestamp((1391463463*10000000) + (100*10) + gregorianToUNIXOffset) {
		t.Errorf("Expected a value greater than 02/03/2014 @ 9:37pm in UTC but got %d", state.past)
	}
	if state.node == nil {
		t.Errorf("Expected a non nil node")
	}
	if state.sequence <= 0 {
		t.Errorf("Expected a value greater than but got %d", state.sequence)
	}
}

func TestUUIDState_read(t *testing.T) {
	s := new(State)
	s.past = Timestamp((1391463463*10000000) + (100*10) + gregorianToUNIXOffset)
	s.node = uuidBytes

	now := Timestamp((1391463463*10000000) + (100*10))
	s.read(now + (100*10), make([]byte, length))

	if s.sequence != 1 {
		t.Error("The sequence should increment when the time is" +
					"older than the state past time and the node" +
					"id are not the same.", s.sequence)
	}
	s.read(now, uuidBytes)

	if s.sequence == 1 {
		t.Error("The sequence should be randomly generated when" +
					" the nodes are equal.", s.sequence)
	}

	s = new(State)
	s.past = Timestamp((1391463463*10000000) + (100*10) + gregorianToUNIXOffset)
	s.node = uuidBytes
	s.randomSequence = true
	s.read(now, make([]byte, length))

	if s.sequence == 0 {
		t.Error("The sequence should be randomly generated when" +
					" the randomSequence flag is set.", s.sequence)
	}

	if s.past != now {
		t.Error("The past time should equal the time passed in" +
				" the method.")
	}

	if !bytes.Equal(s.node, make([]byte, length)) {
		t.Error("The node id should equal the node passed in" +
				" the method.")
	}
}

func TestUUIDState_init(t *testing.T) {


}

// Tests that the schedule is run approx every ten seconds
// takes 90 seconds to complete on my machine at 90000000 UUIDs
func TestUUIDState_saveSchedule(t *testing.T) {
	count := 0
	now := time.Now()
	state.next = timestamp()
	for i := 0; i < 90000000; i++ {
		now = time.Now()
		NewV1()
		if lastTimestamp >= state.next {
			count++
		}
	}
	d := time.Since(now)
	tenSec := int(d.Seconds())/10
	if count != tenSec {
		t.Error("Should be as many saves as ten second increments but got: %s instead of %s", count, tenSec)
	}
}

func TestUUID_encode(t *testing.T) {

}

func TestUUID_decode(t *testing.T) {

}


