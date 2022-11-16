package service

import (
	"context"
	"fmt"
	"getblock/pkg/client"
	"getblock/pkg/utils"
	"strconv"
)

type GetBlockService struct {
	getBlockClient client.GetBlockClienter
}

func NewGetBlockService(gtc client.GetBlockClienter) *GetBlockService {
	return &GetBlockService{getBlockClient: gtc}
}

func (svc *GetBlockService) FindMaxChanged(ctx context.Context) (addr string, err error) {
	blockNumber, err := svc.getBlockClient.GetBlockNumber(ctx)
	if err != nil {
		return "", fmt.Errorf("couldn't get block number err: %w", err)
	}

	blockTrns, err := svc.getBlockClient.GetBlockByNumber(ctx, blockNumber, true)
	if err != nil {
		return "", fmt.Errorf("couldn't get block transactions err: %w", err)
	}

	var maxBlockHash string
	var maxAmount, spent float64
	for i := len(blockTrns.Transaction) - 101; i < len(blockTrns.Transaction); i++ {
		// getting sent amount.
		spent, err = svc.getSpentAmount(
			blockTrns.Transaction[i].Value, blockTrns.Transaction[i].Gas, blockTrns.Transaction[i].GasPrice)
		if err != nil {
			return "", err
		}

		if spent > maxAmount {
			maxBlockHash = blockTrns.Transaction[i].From
			maxAmount = spent
		}
	}

	return maxBlockHash, nil
}

// getSpentAmount return sum of sending coin with gas.
func (svc *GetBlockService) getSpentAmount(amount, gas, gasPrice string) (spent float64, err error) {
	amnt, err := strconv.ParseInt(amount[2:], 16, 64)
	if err != nil {
		return 0, fmt.Errorf("couldn't parse amount err: %w", err)
	}

	gasUnit, err := strconv.ParseInt(gas[2:], 16, 64)
	if err != nil {
		return 0, fmt.Errorf("couldn't parse gas err: %w", err)
	}

	gasPrc, err := strconv.ParseInt(gasPrice[2:], 16, 64)
	if err != nil {
		return 0, fmt.Errorf("couldn't parse gas price err: %w", err)
	}

	gasOver, err := utils.CalculateGas(gasUnit, gasPrc)
	if err != nil {
		return 0, fmt.Errorf("couldn't calculate gas err: %w", err)
	}

	return gasOver + float64(amnt), nil
}
