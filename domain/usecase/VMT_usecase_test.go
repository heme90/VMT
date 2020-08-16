package usecase

import (
	"VendingMachineTracker/adapter/domain"
	"VendingMachineTracker/domain/repository"
	"testing"
)


func TestVmtUseCase_SaveSalesRecord(t *testing.T) {
	type fields struct {
		repo repository.VmtRepository
	}
	type args struct {
		record domain.ISalesRecords
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &VmtUseCase{
				repo: tt.fields.repo,
			}
			if err := v.SaveSalesRecord(tt.args.record); (err != nil) != tt.wantErr {
				t.Errorf("SaveSalesRecord() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}