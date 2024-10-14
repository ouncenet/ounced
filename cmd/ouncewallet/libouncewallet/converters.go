package libouncewallet

import (
	"encoding/hex"

	"github.com/ouncenet/ounced/app/appmessage"
	"github.com/ouncenet/ounced/cmd/ouncewallet/daemon/pb"
	"github.com/ouncenet/ounced/domain/consensus/model/externalapi"
	"github.com/ouncenet/ounced/domain/consensus/utils/transactionid"
	"github.com/ouncenet/ounced/domain/consensus/utils/utxo"
)

// OuncewalletdUTXOsTolibouncewalletUTXOs converts a  []*pb.UtxosByAddressesEntry to a []*libouncewallet.UTXO
func OuncewalletdUTXOsTolibouncewalletUTXOs(ouncewalletdUtxoEntires []*pb.UtxosByAddressesEntry) ([]*UTXO, error) {
	UTXOs := make([]*UTXO, len(ouncewalletdUtxoEntires))
	for i, entry := range ouncewalletdUtxoEntires {
		script, err := hex.DecodeString(entry.UtxoEntry.ScriptPublicKey.ScriptPublicKey)
		if err != nil {
			return nil, err
		}
		transactionID, err := transactionid.FromString(entry.Outpoint.TransactionId)
		if err != nil {
			return nil, err
		}
		UTXOs[i] = &UTXO{
			UTXOEntry: utxo.NewUTXOEntry(
				entry.UtxoEntry.Amount,
				&externalapi.ScriptPublicKey{
					Script:  script,
					Version: uint16(entry.UtxoEntry.ScriptPublicKey.Version),
				},
				entry.UtxoEntry.IsCoinbase,
				entry.UtxoEntry.BlockDaaScore,
			),
			Outpoint: &externalapi.DomainOutpoint{
				TransactionID: *transactionID,
				Index:         entry.Outpoint.Index,
			},
		}
	}
	return UTXOs, nil
}

// AppMessageUTXOToOuncewalletdUTXO converts an appmessage.UTXOsByAddressesEntry to a  pb.UtxosByAddressesEntry
func AppMessageUTXOToOuncewalletdUTXO(appUTXOsByAddressesEntry *appmessage.UTXOsByAddressesEntry) *pb.UtxosByAddressesEntry {
	return &pb.UtxosByAddressesEntry{
		Outpoint: &pb.Outpoint{
			TransactionId: appUTXOsByAddressesEntry.Outpoint.TransactionID,
			Index:         appUTXOsByAddressesEntry.Outpoint.Index,
		},
		UtxoEntry: &pb.UtxoEntry{
			Amount: appUTXOsByAddressesEntry.UTXOEntry.Amount,
			ScriptPublicKey: &pb.ScriptPublicKey{
				Version:         uint32(appUTXOsByAddressesEntry.UTXOEntry.ScriptPublicKey.Version),
				ScriptPublicKey: appUTXOsByAddressesEntry.UTXOEntry.ScriptPublicKey.Script,
			},
			BlockDaaScore: appUTXOsByAddressesEntry.UTXOEntry.BlockDAAScore,
			IsCoinbase:    appUTXOsByAddressesEntry.UTXOEntry.IsCoinbase,
		},
	}
}
