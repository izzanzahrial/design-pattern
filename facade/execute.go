package facade

import (
	"fmt"
	"log"
)

func execute() {
	walletFacade := newWalletFacade("izzan", 04)

	err := walletFacade.addMoneyToWallet("izzan", 04, 9999)
	if err != nil {
		log.Fatalf("Error: %s\n", err.Error())
	}

	fmt.Println()
	err = walletFacade.deductMoneyFromWallet("izzan", 04, 4444)
	if err != nil {
		log.Fatalf("Error: %s\n", err.Error())
	}
}
