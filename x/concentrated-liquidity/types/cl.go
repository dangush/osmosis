package types

import (
	fmt "fmt"
)

func OrderInitialPoolDenoms(denom0, denom1 string) (string, string, error) {
	if denom0 == denom1 {
		return "", "", fmt.Errorf("cannot have the same asset in a single pool")
	}
	if denom0 > denom1 {
		denom1, denom0 = denom0, denom1
	}

	return denom0, denom1, nil
}

// GetConcentratedLockupDenom returns the concentrated lockup denom for a given pool and position.
func GetConcentratedLockupDenom(poolId, positionId uint64) string {
	return fmt.Sprintf("%s/%d/%d", ClTokenPrefix, poolId, positionId)
}
