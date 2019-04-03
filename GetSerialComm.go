package main

import (
	"fmt"
	"io"
	"log"

	"golang.org/x/sys/windows/registry"
)

func main() {
	fmt.Println(getCOMlist())
}

func getCOMlist() []string {
	k, err := registry.OpenKey(registry.LOCAL_MACHINE, `HARDWARE\DEVICEMAP\SERIALCOMM`, registry.QUERY_VALUE)
	if err != nil {
		log.Fatal(err)
	}
	defer k.Close()
	i := 0
	var readErr error
	var s []string
	var comlist []string
	list := make(map[string]string)
	for readErr != io.EOF {
		s, readErr = k.ReadValueNames(i)
		for _, val := range s {
			// fmt.Printf("Deivce name  is %s\n", val)
			if err == nil {
				if _, ok := list[val]; !ok {
					comport, _, err := k.GetStringValue(val)
					if err == nil {
						// fmt.Printf("Comport  is %s\n", comport)
						list[val] = comport
						comlist = append(comlist, comport)
					}
				}
			}
		}
		i++
	}
	return comlist
}
