package usecase

import dto "VendingMachineTracker/adapter/domain"

func NewVmtUseCase(repo dto.Repository) *VmtUseCase {
	return &VmtUseCase{repo: repo}
}
