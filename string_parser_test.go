package main

import "testing"

func Test_Hookup(t *testing.T) {
	t.Log("Hookup Succeeded")
}

func Test_NextByte(t *testing.T) {
	msg := "123"
	sp := NewByteParser()
	sp.SetString(msg)

	c, err := sp.NextByte()
	if c != '1' || err != nil {
		t.Error("Next Byte not as expected")
	}

	c, err = sp.NextByte()
	if c != '2' || err != nil {
		t.Error("Next Byte not as expedcted")
	}

	c, err = sp.NextByte()
	if c != '3' || err != nil {
		t.Error("Next Byte not as expected")
	}

	c, err = sp.NextByte()
	if err == nil || err.Error() != "End of Data" {
		t.Error("Next Byte (err) not as expected")
	}

	t.Log("Next Byte behaved as expected")
}

func Test_PeekByte(t *testing.T) {
	msg := "123"
	sp := NewByteParser()
	sp.SetString(msg)

	c, err := sp.PeekByte()
	if c != '1' || err != nil {
		t.Error("Next Byte not as expected")
	}

	c, err = sp.PeekByte()
	if c != '1' || err != nil {
		t.Error("Next Byte not as expected")
	}

	c, err = sp.PeekBytes(1)
	if c != '2' || err != nil {
		t.Error("Next Byte not as expedcted")
	}

	c, err = sp.PeekBytes(2)
	if c != '3' || err != nil {
		t.Error("Next Byte not as expected")
	}

	c, err = sp.PeekBytes(3)
	if err == nil || err.Error() != "Data boundary exceeded" {
		t.Error("Next Byte (err) not as expected")
	}

	t.Log("Next Byte behaved as expected")
}

func Test_PokeByte(t *testing.T) {
	msg := "123"
	sp := NewByteParser()
	sp.SetString(msg)

	sp.NextByte()

	if c, err := sp.PeekByte(); c != '2' && err != nil {
		t.Error("Peek Byte not as expected")
	}

	if err := sp.PokeByte(); err != nil {
		t.Error("Poke Byte error not as expeced")
	}

	if c, err := sp.PeekByte(); c != '1' || err != nil {
		t.Error("Peek Byte not as expected")
	}

	if err := sp.PokeByte(); err == nil {
		t.Error("Peek Byte (err) not as expected")
	}

	t.Log("Poke Byte functions as expected")
}

func Test_GetSlice(t *testing.T) {
	msg := "Test 123"
	sp := NewByteParser()
	sp.SetString(msg)

	for i := 0; i < 4; i++ {
		sp.NextByte()
	}

	str := sp.GetSlice()
	if string(str) != "Test" {
		t.Error("GetSlice did not behave as expected")
		t.Error("GetSlice returned: ", str)
	}

	t.Log("GetString behaved as expected")
}
