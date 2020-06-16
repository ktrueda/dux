package visualization

import (
	"testing"
)

func TestShow(t *testing.T) {
	type args struct {
		data map[string]int64
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"1 file",
			args{
				map[string]int64{"/tmp/file1": 1},
			},
		},
		{
			"2 file",
			args{
				map[string]int64{
					"/tmp/file1": 1,
					"/tmp/file2": 2,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Show(tt.args.data)
		})
	}
}

func Test_showLine(t *testing.T) {
	type args struct {
		label      string
		val        int64
		length     int
		unit       string
		labelFomat string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"medium",
			args{"label1", 123, 10, "B", "%10s : "},
			"    label1 : ▇▇▇▇▇▇▇▇▇▇ 123 B",
		},
		{
			"small",
			args{"label1", 123456, 0, "B", "%7s : "},
			" label1 :  123,456 B",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := showLine(tt.args.label, tt.args.val, tt.args.length, tt.args.unit, tt.args.labelFomat); got != tt.want {
				t.Errorf("showLine() = %v, want %v", got, tt.want)
			}
		})
	}
}
