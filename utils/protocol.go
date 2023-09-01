package utils

import (
	"bytes"
	"encoding/binary"
)

const (
	Header       = "#Begin#"
	HeaderLength = 7
	DataLength   = 4
)

func IntToBytes(n int) []byte {
	x := int32(n)

	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, x)
	return bytesBuffer.Bytes()
}

func BytesToInt(b []byte) int {
	bytesBuffer := bytes.NewBuffer(b)

	var x uint32
	binary.Read(bytesBuffer, binary.BigEndian, &x)
	return int(x)
}

func Depack(buffer []byte, readerChan chan []byte) []byte {
	length := len(buffer)
	var i int
	for i = 0; i < length; i++ {
		if length < i+HeaderLength+DataLength {
			break
		}

		if string(buffer[i:i+HeaderLength]) == Header {
			messageLength := BytesToInt(buffer[i+HeaderLength : i+HeaderLength+DataLength])
			if length < i+HeaderLength+DataLength+messageLength {
				break
			}
			data := buffer[i+HeaderLength+DataLength : i+HeaderLength+DataLength+messageLength]
			readerChan <- data

			i += HeaderLength + DataLength + messageLength - 1
		}
	}
	if i == length {
		return make([]byte, 0)
	}
	return buffer[i:]
}

func Enpack(data []byte) []byte {
	return append(append([]byte(Header), IntToBytes(len(data))...), data...)
}
