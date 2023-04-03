package types

func (resume Resume) UpdateBalances() {
	resume.BalanceTotal = int32(resume.CreditedTotal) - int32(resume.ReturnedTotal)
	resume.BalanceCurrent = int32(resume.CreditedCurrent) - int32(resume.ReturnedTotal)
}
