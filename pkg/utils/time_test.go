package utils

import (
	"testing"
	"time"
)

func TestGetMonthDayCount(t *testing.T) {
	type args struct {
		t time.Time
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "month test",
			args: args{
				t: time.Date(2023, 2, 10, 0, 0, 0, 0, time.Local),
			},
			want: 28,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetMonthDayCount(tt.args.t); got != tt.want {
				t.Errorf("GetMonthDayCount() = %v, want %v", got, tt.want)
			}
		})
	}
}
