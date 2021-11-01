package factory

import "fmt"

func getGun(gunType string) (iGun, error) {
	if gunType == "ak47" {
		return newAk47(), nil
	} else if gunType == "m4a1" {
		return newM4a1(), nil
	} else {
		return nil, fmt.Errorf("Wrong type of gun")
	}
}
