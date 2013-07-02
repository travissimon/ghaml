package main

import (
	"errors"
)

type ByteParser struct {
	cp, np int
	data   []byte
}

// NewByteParser returns an initialised ByteParser.
func NewByteParser() *ByteParser {
	return new(ByteParser)
}

// SetData sets the data to be parsed and resets 
// the pointers to the beginning of the string.
func (bp *ByteParser) SetData(data []byte) {
	bp.data = data
	bp.np = 0
	bp.ResetPointers()
}

func (bp *ByteParser) SetString(str string) {
	bp.SetData([]byte(str))
}

// ResetPointers sets the current position as the starting point for the next call
// to the GetString() function.
func (bp *ByteParser) ResetPointers() {
	bp.cp = bp.np
}

// HasMoreData checks the bounds of the string.
func (bp *ByteParser) HasMoreData() bool {
	return bp.np < len(bp.data)
}

// NextChar returns the next character in the string 
// and increments the current position.
func (bp *ByteParser) NextByte() (byte, error) {
	if !bp.HasMoreData() {
		return 0, errors.New("End of Data")
	}
	c := bp.data[bp.np]
	bp.np++
	return c, nil
}

// PeekChar returns the next character in the string 
// but does not increment the current position
func (bp *ByteParser) PeekByte() (byte, error) {
	return bp.PeekBytes(0)
}

// PeekChars returns a character 'charCount' characters ahead in the string
// but does not increment the current position
func (bp *ByteParser) PeekBytes(charCount int) (byte, error) {
	index := bp.np + charCount
	if index >= len(bp.data) {
		return 0, errors.New("Data boundary exceeded")
	}

	return bp.data[index], nil
}

// PokeChar moves the current position in the string backwards by 1 character.
func (bp *ByteParser) PokeByte() error {
	if bp.np <= 0 {
		return errors.New("Data boundary exceeded")
	}
	bp.np--
	return nil
}

func (bp *ByteParser) GetSlice() []byte {
	defer bp.ResetPointers()
	return bp.data[bp.cp:bp.np]
}

func (bp *ByteParser) SkipWhitespace() {
	if !bp.HasMoreData() {
		return
	}
}
