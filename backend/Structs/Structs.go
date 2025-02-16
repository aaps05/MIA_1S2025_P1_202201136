package Structs

import (
	"fmt"
	
)


//MBR
type MBR struct {
	Mbrsize int32
	CreationDate [10]byte
	Signature int32
	Fit [1]byte
	//Partitions [4]Partition
}

func ImpresionMBR(data MBR) {
	fmt.Println(fmt.Sprintf("Creation Date: %s, Fit: %s, Size: %d, Signature: %d ",
		string(data.CreationDate[:]),
		string(data.Fit[:]),
		data.Mbrsize,
		data.Signature))

	//return "Mbrsize: " + string(m.Mbrsize) + "\nCreationDate: " + string(m.CreationDate[:]) + "\nSignature: " + string(m.Signature) + "\nFit: " + string(m.Fit[:])
}