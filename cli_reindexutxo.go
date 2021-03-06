
package main 

import "fmt"


func (cli *CLI) reindexUTXO() {

	bc := NewBlockchain()
	UTXOSet := UTXOSet{bc}
	UTXOSet.Reindex()

	count := UTXOSet.CountTransactions()
	fmt.printf("Done! There are %d transactions in the UTXO Set .\n", count)
}


