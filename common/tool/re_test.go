package tool

import "testing"

func TestVerifyPhone(t *testing.T) {
	type args struct {
		phone string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "1",
			args: args{
				phone: "18274612897",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := VerifyPhone(tt.args.phone); got != tt.want {
				t.Errorf("VerifyPhone() = %v, want %v", got, tt.want)
			}
		})
	}
}
