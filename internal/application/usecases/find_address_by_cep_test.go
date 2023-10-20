package usecases

import (
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/ismael3s/go-cep/internal/application/gateways"
	"github.com/ismael3s/go-cep/internal/domain"
)

type MockCEPGateway struct{}

func (m *MockCEPGateway) GetName() string {
	return "MockCEPGateway"
}

func (m *MockCEPGateway) FindAddressByCEP(cep string) (domain.Address, error) {
	if cep == "00000000" {
		return domain.Address{}, errors.New("Some external error")
	}
	return domain.Address{
		Cep:        cep,
		Logradouro: "Rua dos Bobos",
		Bairro:     "Vila do Chaves",
		Cidade:     "São Paulo",
		// Service:    "MockCEPGateway",
	}, nil
}

func TestFindAddressByCEPUseCase_Do(t *testing.T) {
	type fields struct {
		cepsGateways []gateways.ICEPGateway
		duration     time.Duration
	}
	type args struct {
		input FindAddressByCEPInput
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		want      FindAddressByCEPOutput
		wantErr   bool
		wantedErr error
	}{
		{
			name: "Given an invalid CEP, should return the error returned by the domain",
			fields: fields{
				cepsGateways: []gateways.ICEPGateway{},
			},
			args: args{
				input: FindAddressByCEPInput{
					Value: "40283-31",
				},
			},
			want:      FindAddressByCEPOutput{},
			wantErr:   true,
			wantedErr: domain.INVALID_CEP_FORMAT,
		},
		{
			name: "Given an valid CEP, and the response from the first gateway is fast enough, should return the address",
			fields: fields{
				cepsGateways: []gateways.ICEPGateway{&MockCEPGateway{}},
			},
			args: args{
				input: FindAddressByCEPInput{
					Value: "40283310",
				},
			},
			want: FindAddressByCEPOutput{
				Address: domain.Address{
					Cep:        "40283310",
					Logradouro: "Rua dos Bobos",
					Bairro:     "Vila do Chaves",
					Cidade:     "São Paulo",
					// Service:    "MockCEPGateway",
				},
			},
			wantErr: false,
		},
		{
			name: "Given an valid CEP, and the response from the gateways is too slow return an time out",
			fields: fields{
				cepsGateways: []gateways.ICEPGateway{},
				duration:     10 * time.Millisecond,
			},
			args: args{
				input: FindAddressByCEPInput{
					Value: "40283310",
				},
			},
			want:      FindAddressByCEPOutput{},
			wantErr:   true,
			wantedErr: errors.New("timeout"),
		},
		{
			name: "Given an valid CEP, and all gateways return an error, the request should handle a timeout",
			fields: fields{
				cepsGateways: []gateways.ICEPGateway{&MockCEPGateway{}},
				duration:     10 * time.Millisecond,
			},
			args: args{
				input: FindAddressByCEPInput{
					Value: "00000000",
				},
			},
			want:      FindAddressByCEPOutput{},
			wantErr:   true,
			wantedErr: errors.New("timeout"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := NewFindAddresByCEPUseCase(tt.fields.duration, tt.fields.cepsGateways...)
			got, err := u.Do(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindAddressByCEPUseCase.Do() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil && err.Error() != tt.wantedErr.Error() {
				t.Errorf("FindAddressByCEPUseCase.Do() error = %v, wantedErr %v", err, tt.wantedErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindAddressByCEPUseCase.Do() = %v, want %v", got, tt.want)
			}
		})
	}
}
