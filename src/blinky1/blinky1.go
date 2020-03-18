package main

import (
	"machine"
)

var buffer = make([]byte, 100)

func main() {

	var rxdData bool = false

	for {

		/*
		   println("\nThis is from 'println( hard-coded string)'")

		   machine.UART0.WriteByte('a')
		   println("\nThe 'a' above is machine.UART0.WriteByte('a')")

		   data := []byte("This is a generated byte[] array named 'data' prepopulated with these words.")
		   machine.UART0.Write(data)
		   println("\nThe words above machine.UART0.Write(data)")
		*/
		theArray := make([]byte, 100)
		a, _ := machine.UART0.Read(theArray)

		if a != 0 {
			rxdData = true
			str := string(theArray[:])
			println(str)
		} else {
			if rxdData {
				println("")
				rxdData = false
			}
		}

		/*
		   //theByte, _:= machine.UART0.ReadByte()
		   if theByte != 0 {
		       print(theByte)
		       rxdData = true
		   } else {
		       if rxdData {
		           println("")
		           rxdData = false
		       }
		   }
		*/

	}
}
