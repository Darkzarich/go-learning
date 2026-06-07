package sum

import "testing"

func TestSumPositive(t *testing.T) {
	sum, err := Sum(1, 2)
	if err != nil {
		t.Error("unexpected error")
	}
	if sum != 3 {
		t.Errorf("sum expected to be 3; got %d", sum)
	}
}

func TestSumNegative(t *testing.T) {
	_, err := Sum(-1, 2)
	if err == nil {
		t.Error("first arg negative - expected error not be nil")
	}
	_, err = Sum(1, -2)
	if err == nil {
		t.Error("second arg negative - expected error not be nil")
	}
	_, err = Sum(-1, -2)
	if err == nil {
		t.Error("all arg negative - expected error not be nil")
	}
}

func TestSumZero(t *testing.T) {
	_, err := Sum(0, 0)
	if err == nil {
		t.Error("both args zero - expected error not be nil")
	}
}

// Table-driven tests generated with gotests
func TestSum(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		// Cases
		{
			name: "positive",
			args: args{
				a: 1,
				b: 2,
			},
			want:    3,
			wantErr: false,
		},
		{
			name: "negative",
			args: args{
				a: -1,
				b: 2,
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "negative all",
			args: args{
				a: -1,
				b: -3,
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "zero",
			args: args{
				a: 0,
				b: 0,
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Sum(tt.args.a, tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("Sum() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Sum() = %v, want %v", got, tt.want)
			}
		})
	}
}
