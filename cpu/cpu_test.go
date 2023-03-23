package cpu

import "testing"

func Test_parseCore(t *testing.T) {
	type args struct {
		cpuSecond string
	}
	tests := []struct {
		name    string
		args    args
		want    uint32
		wantErr bool
	}{
		{
			name: "0.5",
			args: args{cpuSecond: "0.5"},
			want: 500,
		},
		{
			name: "1",
			args: args{cpuSecond: "1"},
			want: 1000,
		},
		{
			name: "2.5",
			args: args{cpuSecond: "2.5"},
			want: 2500,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseCore(tt.args.cpuSecond)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseCore() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ParseCore() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseMilli(t *testing.T) {
	type args struct {
		cpuSecond string
	}
	tests := []struct {
		name    string
		args    args
		want    uint32
		wantErr bool
	}{
		{
			name: "1m",
			args: args{cpuSecond: "1m"},
			want: 1,
		},
		{
			name: "2500m",
			args: args{cpuSecond: "2500m"},
			want: 2500,
		},
		{
			name:    "2.5m_error",
			args:    args{cpuSecond: "2.5m"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseMilli(tt.args.cpuSecond)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseMilli() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ParseMilli() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParse(t *testing.T) {
	type args struct {
		cpuSecond string
	}
	tests := []struct {
		name    string
		args    args
		want    uint32
		wantErr bool
	}{
		{
			name: "1",
			args: args{cpuSecond: "1"},
			want: 1000,
		},
		{
			name: "2.5",
			args: args{cpuSecond: "2.5"},
			want: 2500,
		},
		{
			name: "1000m",
			args: args{cpuSecond: "1000m"},
			want: 1000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Parse(tt.args.cpuSecond)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Parse() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToSeconds(t *testing.T) {
	type args struct {
		cpu float32
	}
	tests := []struct {
		name string
		args args
		want uint32
	}{
		{
			name: "2",
			args: args{cpu: 2},
			want: 2000,
		},
		{
			name: "1.5",
			args: args{cpu: 1.5},
			want: 1500,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToSeconds(tt.args.cpu); got != tt.want {
				t.Errorf("ToSeconds() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToCpu(t *testing.T) {
	type args struct {
		seconds uint32
	}
	tests := []struct {
		name string
		args args
		want float32
	}{
		{
			name: "1000",
			args: args{seconds: 1000},
			want: 1,
		},
		{
			name: "2500",
			args: args{seconds: 2500},
			want: 2.5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToCpu(tt.args.seconds); got != tt.want {
				t.Errorf("ToCpu() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCeil(t *testing.T) {
	type args struct {
		cpu float32
	}
	tests := []struct {
		name string
		args args
		want uint8
	}{
		{
			name: "0.5->1",
			args: args{cpu: 0.5},
			want: 1,
		},
		{
			name: "2",
			args: args{cpu: 2},
			want: 2,
		},
		{
			name: "1.2->2",
			args: args{cpu: 1.2},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Ceil(tt.args.cpu); got != tt.want {
				t.Errorf("Ceil() = %v, want %v", got, tt.want)
			}
		})
	}
}
