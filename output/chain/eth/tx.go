package eth

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
)

func SendTX(ctx context.Context, endpoint, privateKey, toStr string, data []byte) (string, error) {
	cli, err := ethclient.Dial(endpoint)
	if err != nil {
		return "", errors.Wrapf(err, "dial eth endpoint %s failed", endpoint)
	}

	pk := crypto.ToECDSAUnsafe(common.FromHex(privateKey))
	sender := crypto.PubkeyToAddress(pk.PublicKey)
	to := common.HexToAddress(toStr)

	gasPrice, err := cli.SuggestGasPrice(ctx)
	if err != nil {
		return "", errors.Wrap(err, "get suggest gas price failed")
	}

	chainid, err := cli.ChainID(ctx)
	if err != nil {
		return "", errors.Wrap(err, "get chain id failed")
	}

	nonce, err := cli.PendingNonceAt(ctx, sender)
	if err != nil {
		return "", errors.Wrap(err, "get pending nonce failed")
	}

	tx := types.NewTx(
		&types.LegacyTx{
			Nonce:    nonce,
			GasPrice: gasPrice,
			Gas:      uint64(3000000000), // TODO get gas limit by cli.EstimateGas
			To:       &to,
			Data:     data,
		})

	signedTx, err := types.SignTx(tx, types.NewLondonSigner(chainid), pk)
	if err != nil {
		return "", errors.Wrap(err, "sign tx failed")
	}

	if err = cli.SendTransaction(ctx, signedTx); err != nil {
		return "", errors.Wrap(err, "send transaction failed")
	}

	return tx.Hash().Hex(), nil
}
