package repository

import (
	"testing"
)

func TestVmtRepository_GetInstance(t *testing.T) {
	vmtRepo := VmtRepository{}
	vmtRepo.InitRepository()

}
