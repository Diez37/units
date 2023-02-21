package size

import (
	"fmt"
	"testing"
)

// Tests

func TestCalculateSizeAndSuffix(t *testing.T) {
	type args struct {
		size     float64
		suffixes Suffixes
	}
	tests := []struct {
		name  string
		args  args
		want  float64
		want1 Suffix
	}{
		{
			name: "binary/Byte",
			args: args{
				size:     512,
				suffixes: binarySuffixes,
			},
			want:  512,
			want1: Byte,
		},
		{
			name: "binary/KibiByte",
			args: args{
				size:     512 * 1024,
				suffixes: binarySuffixes,
			},
			want:  512,
			want1: KibiByte,
		},
		{
			name: "binary/MebiByte",
			args: args{
				size:     512 * (1024 * 1024),
				suffixes: binarySuffixes,
			},
			want:  512,
			want1: MebiByte,
		},
		{
			name: "binary/GibiByte",
			args: args{
				size:     512 * (1024 * 1024 * 1024),
				suffixes: binarySuffixes,
			},
			want:  512,
			want1: GibiByte,
		},
		{
			name: "binary/TebiByte",
			args: args{
				size:     512 * (1024 * 1024 * 1024 * 1024),
				suffixes: binarySuffixes,
			},
			want:  512,
			want1: TebiByte,
		},
		{
			name: "binary/PebiByte",
			args: args{
				size:     512 * (1024 * 1024 * 1024 * 1024 * 1024),
				suffixes: binarySuffixes,
			},
			want:  512,
			want1: PebiByte,
		},
		{
			name: "decimal/Byte",
			args: args{
				size:     512,
				suffixes: decimalSuffixes,
			},
			want:  512,
			want1: Byte,
		},
		{
			name: "decimal/KiloByte",
			args: args{
				size:     512 * 1000,
				suffixes: decimalSuffixes,
			},
			want:  512,
			want1: KiloByte,
		},
		{
			name: "decimal/MegaByte",
			args: args{
				size:     512 * (1000 * 1000),
				suffixes: decimalSuffixes,
			},
			want:  512,
			want1: MegaByte,
		},
		{
			name: "decimal/GigaByte",
			args: args{
				size:     512 * (1000 * 1000 * 1000),
				suffixes: decimalSuffixes,
			},
			want:  512,
			want1: GigaByte,
		},
		{
			name: "decimal/TeraByte",
			args: args{
				size:     512 * (1000 * 1000 * 1000 * 1000),
				suffixes: decimalSuffixes,
			},
			want:  512,
			want1: TeraByte,
		},
		{
			name: "decimal/PetaByte",
			args: args{
				size:     512 * (1000 * 1000 * 1000 * 1000 * 1000),
				suffixes: decimalSuffixes,
			},
			want:  512,
			want1: PetaByte,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := calculateSizeAndSuffix(tt.args.size, tt.args.suffixes)
			if got != tt.want {
				t.Errorf("calculateSizeAndSuffix() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("calculateSizeAndSuffix() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestFormatBinary(t *testing.T) {
	type args struct {
		unit uint64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Byte",
			args: args{unit: 512},
			want: "512B",
		},
		{
			name: "KibiByte",
			args: args{unit: 512 * 1024},
			want: "512KiB",
		},
		{
			name: "MebiByte",
			args: args{unit: 512 * (1024 * 1024)},
			want: "512MiB",
		},
		{
			name: "GibiByte",
			args: args{unit: 512 * (1024 * 1024 * 1024)},
			want: "512GiB",
		},
		{
			name: "TebiByte",
			args: args{unit: 512 * (1024 * 1024 * 1024 * 1024)},
			want: "512TiB",
		},
		{
			name: "PebiByte",
			args: args{unit: 512 * (1024 * 1024 * 1024 * 1024 * 1024)},
			want: "512PiB",
		},
		{
			name: "KiloByteToKibiByte",
			args: args{unit: 512 * 1000},
			want: "500KiB",
		},
		{
			name: "MegaByteToMebiByte",
			args: args{unit: 512 * (1000 * 1000)},
			want: "488.3MiB",
		},
		{
			name: "GigaByteToGibiByte",
			args: args{unit: 512 * (1000 * 1000 * 1000)},
			want: "476.8GiB",
		},
		{
			name: "TeraByteToTebiByte",
			args: args{unit: 512 * (1000 * 1000 * 1000 * 1000)},
			want: "465.7TiB",
		},
		{
			name: "PetaByteToPebiByte",
			args: args{unit: 512 * (1000 * 1000 * 1000 * 1000 * 1000)},
			want: "454.7PiB",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FormatBinary(tt.args.unit); got != tt.want {
				t.Errorf("FormatBinary() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFormatHuman(t *testing.T) {
	type args struct {
		unit uint64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Byte",
			args: args{unit: 512},
			want: "512B",
		},
		{
			name: "KiloByte",
			args: args{unit: 512 * 1000},
			want: "512kB",
		},
		{
			name: "MegaByte",
			args: args{unit: 512 * (1000 * 1000)},
			want: "512MB",
		},
		{
			name: "GigaByte",
			args: args{unit: 512 * (1000 * 1000 * 1000)},
			want: "512GB",
		},
		{
			name: "TeraByte",
			args: args{unit: 512 * (1000 * 1000 * 1000 * 1000)},
			want: "512TB",
		},
		{
			name: "PetaByte",
			args: args{unit: 512 * (1000 * 1000 * 1000 * 1000 * 1000)},
			want: "512PB",
		},
		{
			name: "KibiByteToKiloByte",
			args: args{unit: 512 * 1024},
			want: "524.3kB",
		},
		{
			name: "MebiByteToMegaByte",
			args: args{unit: 512 * (1024 * 1024)},
			want: "536.9MB",
		},
		{
			name: "GibiByteToGigaByte",
			args: args{unit: 512 * (1024 * 1024 * 1024)},
			want: "549.8GB",
		},
		{
			name: "TebiByteToTeraByte",
			args: args{unit: 512 * (1024 * 1024 * 1024 * 1024)},
			want: "562.9TB",
		},
		{
			name: "PebiByteToPetaByte",
			args: args{unit: 512 * (1024 * 1024 * 1024 * 1024 * 1024)},
			want: "576.5PB",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FormatHuman(tt.args.unit); got != tt.want {
				t.Errorf("FormatHuman() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFormatSize(t *testing.T) {
	type args struct {
		format   string
		unit     uint64
		suffixes Suffixes
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "binary/Byte",
			args: args{
				format:   "%.4g%s",
				unit:     512,
				suffixes: binarySuffixes,
			},
			want: "512B",
		},
		{
			name: "binary/KibiByte",
			args: args{
				format:   "%.4g%s",
				unit:     512 * 1024,
				suffixes: binarySuffixes,
			},
			want: "512KiB",
		},
		{
			name: "binary/MebiByte",
			args: args{
				format:   "%.4g%s",
				unit:     512 * (1024 * 1024),
				suffixes: binarySuffixes,
			},
			want: "512MiB",
		},
		{
			name: "binary/GibiByte",
			args: args{
				format:   "%.4g%s",
				unit:     512 * (1024 * 1024 * 1024),
				suffixes: binarySuffixes,
			},
			want: "512GiB",
		},
		{
			name: "binary/TebiByte",
			args: args{
				format:   "%.4g%s",
				unit:     512 * (1024 * 1024 * 1024 * 1024),
				suffixes: binarySuffixes,
			},
			want: "512TiB",
		},
		{
			name: "binary/PebiByte",
			args: args{
				format:   "%.4g%s",
				unit:     512 * (1024 * 1024 * 1024 * 1024 * 1024),
				suffixes: binarySuffixes,
			},
			want: "512PiB",
		},
		{
			name: "decimal/Byte",
			args: args{
				format:   "%.4g%s",
				unit:     512,
				suffixes: decimalSuffixes,
			},
			want: "512B",
		},
		{
			name: "decimal/KiloByte",
			args: args{
				format:   "%.4g%s",
				unit:     512 * 1000,
				suffixes: decimalSuffixes,
			},
			want: "512kB",
		},
		{
			name: "decimal/MegaByte",
			args: args{
				format:   "%.4g%s",
				unit:     512 * (1000 * 1000),
				suffixes: decimalSuffixes,
			},
			want: "512MB",
		},
		{
			name: "decimal/GigaByte",
			args: args{
				format:   "%.4g%s",
				unit:     512 * (1000 * 1000 * 1000),
				suffixes: decimalSuffixes,
			},
			want: "512GB",
		},
		{
			name: "decimal/TeraByte",
			args: args{
				format:   "%.4g%s",
				unit:     512 * (1000 * 1000 * 1000 * 1000),
				suffixes: decimalSuffixes,
			},
			want: "512TB",
		},
		{
			name: "decimal/PetaByte",
			args: args{
				format:   "%.4g%s",
				unit:     512 * (1000 * 1000 * 1000 * 1000 * 1000),
				suffixes: decimalSuffixes,
			},
			want: "512PB",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FormatSize(tt.args.format, tt.args.unit, tt.args.suffixes); got != tt.want {
				t.Errorf("FormatSize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFromBinarySize(t *testing.T) {
	type args struct {
		size string
	}
	tests := []struct {
		name    string
		args    args
		want    uint64
		wantErr bool
	}{
		{
			name: "Byte",
			args: args{size: "512B"},
			want: 512,
		},
		{
			name: "Byte/NoSuffix",
			args: args{size: "512"},
			want: 512,
		},
		{
			name: "KibiByte",
			args: args{size: "512KiB"},
			want: 512 * 1024,
		},
		{
			name: "MebiByte",
			args: args{size: "512MiB"},
			want: 512 * (1024 * 1024),
		},
		{
			name: "GibiByte",
			args: args{size: "512GiB"},
			want: 512 * (1024 * 1024 * 1024),
		},
		{
			name: "TebiByte",
			args: args{size: "512TiB"},
			want: 512 * (1024 * 1024 * 1024 * 1024),
		},
		{
			name: "PebiByte",
			args: args{size: "512PiB"},
			want: 512 * (1024 * 1024 * 1024 * 1024 * 1024),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FromBinarySize(tt.args.size)
			if (err != nil) != tt.wantErr {
				t.Errorf("FromBinarySize() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("FromBinarySize() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFromHumanSize(t *testing.T) {
	type args struct {
		size string
	}
	tests := []struct {
		name    string
		args    args
		want    uint64
		wantErr bool
	}{
		{
			name: "Byte",
			args: args{size: "512B"},
			want: 512,
		},
		{
			name: "Byte/NoSuffix",
			args: args{size: "512"},
			want: 512,
		},
		{
			name: "KiloByte",
			args: args{size: "512kB"},
			want: 512 * 1000,
		},
		{
			name: "MegaByte",
			args: args{size: "512MB"},
			want: 512 * (1000 * 1000),
		},
		{
			name: "MegaByte/NoSuffixByte",
			args: args{size: "512M"},
			want: 512 * (1000 * 1000),
		},
		{
			name: "GigaByte",
			args: args{size: "512GB"},
			want: 512 * (1000 * 1000 * 1000),
		},
		{
			name: "TeraByte",
			args: args{size: "512TB"},
			want: 512 * (1000 * 1000 * 1000 * 1000),
		},
		{
			name: "PetaByte",
			args: args{size: "512PB"},
			want: 512 * (1000 * 1000 * 1000 * 1000 * 1000),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FromHumanSize(tt.args.size)
			if (err != nil) != tt.wantErr {
				t.Errorf("FromHumanSize() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("FromHumanSize() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFromSize(t *testing.T) {
	type args struct {
		size  string
		units Units
	}
	tests := []struct {
		name    string
		args    args
		want    uint64
		wantErr bool
	}{
		{
			name: "binary/Byte",
			args: args{
				size:  "512B",
				units: binaryUnits,
			},
			want: 512,
		},
		{
			name: "binary/KibiByte",
			args: args{
				size:  "512KiB",
				units: binaryUnits,
			},
			want: 512 * 1024,
		},
		{
			name: "binary/MebiByte",
			args: args{
				size:  "512MiB",
				units: binaryUnits,
			},
			want: 512 * (1024 * 1024),
		},
		{
			name: "binary/GibiByte",
			args: args{
				size:  "512GiB",
				units: binaryUnits,
			},
			want: 512 * (1024 * 1024 * 1024),
		},
		{
			name: "binary/TebiByte",
			args: args{
				size:  "512TiB",
				units: binaryUnits,
			},
			want: 512 * (1024 * 1024 * 1024 * 1024),
		},
		{
			name: "binary/PebiByte",
			args: args{
				size:  "512PiB",
				units: binaryUnits,
			},
			want: 512 * (1024 * 1024 * 1024 * 1024 * 1024),
		},
		{
			name: "decimal/Byte",
			args: args{
				size:  "512B",
				units: decimalUnits,
			},
			want: 512,
		},
		{
			name: "decimal/KiloByte",
			args: args{
				size:  "512kB",
				units: decimalUnits,
			},
			want: 512 * 1000,
		},
		{
			name: "decimal/MegaByte",
			args: args{
				size:  "512MB",
				units: decimalUnits,
			},
			want: 512 * (1000 * 1000),
		},
		{
			name: "decimal/GigaByte",
			args: args{
				size:  "512GB",
				units: decimalUnits,
			},
			want: 512 * (1000 * 1000 * 1000),
		},
		{
			name: "decimal/TeraByte",
			args: args{
				size:  "512TB",
				units: decimalUnits,
			},
			want: 512 * (1000 * 1000 * 1000 * 1000),
		},
		{
			name: "decimal/PetaByte",
			args: args{
				size:  "512PB",
				units: decimalUnits,
			},
			want: 512 * (1000 * 1000 * 1000 * 1000 * 1000),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FromSize(tt.args.size, tt.args.units)
			if (err != nil) != tt.wantErr {
				t.Errorf("FromSize() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("FromSize() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseSize(t *testing.T) {
	type args struct {
		size string
	}
	tests := []struct {
		name    string
		args    args
		want    uint64
		wantErr bool
	}{
		{
			name: "binary/Byte",
			args: args{size: "512B"},
			want: 512,
		},
		{
			name: "binary/Byte/LowerCase",
			args: args{size: "512b"},
			want: 512,
		},
		{
			name: "binary/KibiByte",
			args: args{size: "512KiB"},
			want: 512 * 1024,
		},
		{
			name: "binary/KibiByte/UpperCase",
			args: args{size: "512KIB"},
			want: 512 * 1024,
		},
		{
			name: "binary/KibiByte/LowerCase",
			args: args{size: "512kib"},
			want: 512 * 1024,
		},
		{
			name: "binary/MebiByte",
			args: args{size: "512MiB"},
			want: 512 * (1024 * 1024),
		},
		{
			name: "binary/MebiByte/UpperCase",
			args: args{size: "512MIB"},
			want: 512 * (1024 * 1024),
		},
		{
			name: "binary/MebiByte/LowerCase",
			args: args{size: "512mib"},
			want: 512 * (1024 * 1024),
		},
		{
			name: "binary/GibiByte",
			args: args{size: "512GiB"},
			want: 512 * (1024 * 1024 * 1024),
		},
		{
			name: "binary/GibiByte/UpperCase",
			args: args{size: "512GIB"},
			want: 512 * (1024 * 1024 * 1024),
		},
		{
			name: "binary/GibiByte/LowerCase",
			args: args{size: "512gib"},
			want: 512 * (1024 * 1024 * 1024),
		},
		{
			name: "binary/TebiByte",
			args: args{size: "512TiB"},
			want: 512 * (1024 * 1024 * 1024 * 1024),
		},
		{
			name: "binary/TebiByte/UpperCase",
			args: args{size: "512TIB"},
			want: 512 * (1024 * 1024 * 1024 * 1024),
		},
		{
			name: "binary/TebiByte/LowerCase",
			args: args{size: "512tib"},
			want: 512 * (1024 * 1024 * 1024 * 1024),
		},
		{
			name: "binary/PebiByte",
			args: args{size: "512PiB"},
			want: 512 * (1024 * 1024 * 1024 * 1024 * 1024),
		},
		{
			name: "binary/PebiByte/UpperCase",
			args: args{size: "512PIB"},
			want: 512 * (1024 * 1024 * 1024 * 1024 * 1024),
		},
		{
			name: "binary/PebiByte/LowerCase",
			args: args{size: "512pib"},
			want: 512 * (1024 * 1024 * 1024 * 1024 * 1024),
		},

		{
			name: "decimal/Byte",
			args: args{size: "512B"},
			want: 512,
		},
		{
			name: "decimal/Byte/LowerCase",
			args: args{size: "512b"},
			want: 512,
		},
		{
			name: "decimal/KiloByte",
			args: args{size: "512kB"},
			want: 512 * 1000,
		},
		{
			name: "decimal/KiloByte/UpperCase",
			args: args{size: "512KB"},
			want: 512 * 1000,
		},
		{
			name: "decimal/KiloByte/LowerCase",
			args: args{size: "512kb"},
			want: 512 * 1000,
		},
		{
			name: "decimal/MegaByte",
			args: args{size: "512MB"},
			want: 512 * (1000 * 1000),
		},
		{
			name: "decimal/MegaByte/LowerCase",
			args: args{size: "512mb"},
			want: 512 * (1000 * 1000),
		},
		{
			name: "decimal/GigaByte",
			args: args{size: "512GB"},
			want: 512 * (1000 * 1000 * 1000),
		},
		{
			name: "decimal/GigaByte/LowerCase",
			args: args{size: "512gb"},
			want: 512 * (1000 * 1000 * 1000),
		},
		{
			name: "decimal/TeraByte",
			args: args{size: "512TB"},
			want: 512 * (1000 * 1000 * 1000 * 1000),
		},
		{
			name: "decimal/TeraByte/LowerCase",
			args: args{size: "512tb"},
			want: 512 * (1000 * 1000 * 1000 * 1000),
		},
		{
			name: "decimal/PetaByte",
			args: args{size: "512PB"},
			want: 512 * (1000 * 1000 * 1000 * 1000 * 1000),
		},
		{
			name: "decimal/PetaByte/LowerCase",
			args: args{size: "512pb"},
			want: 512 * (1000 * 1000 * 1000 * 1000 * 1000),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseSize(tt.args.size)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseSize() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ParseSize() got = %v, want %v", got, tt.want)
			}
		})
	}
}

// Examples

func ExampleFormatBinary() {
	MiB5_6 := 5.6 * MiB
	GiB8_6 := 8.6 * GiB
	PiB2_7 := 2.7 * PiB

	fmt.Println(FormatBinary(1))
	fmt.Println(FormatBinary(1024))
	fmt.Println(FormatBinary(1024 * 1024))
	fmt.Println(FormatBinary(1048576))
	fmt.Println(FormatBinary(1024 * 1024 * 1024))
	fmt.Println(FormatBinary(uint64(MiB5_6)))
	fmt.Println(FormatBinary(uint64(GiB8_6)))
	fmt.Println(FormatBinary(uint64(PiB2_7)))
}

func ExampleFormatHuman() {
	MB5_6 := 5.6 * MB
	GB8_6 := 8.6 * GB
	PB2_7 := 2.7 * PB

	fmt.Println(FormatHuman(1))
	fmt.Println(FormatHuman(1000))
	fmt.Println(FormatHuman(1000 * 1000))
	fmt.Println(FormatHuman(1000000))
	fmt.Println(FormatHuman(1048576))
	fmt.Println(FormatHuman(1000 * 1000 * 1000))
	fmt.Println(FormatHuman(uint64(MB5_6)))
	fmt.Println(FormatHuman(uint64(GB8_6)))
	fmt.Println(FormatHuman(uint64(PB2_7)))
}

func ExampleFromBinarySize() {
	fmt.Println(FromBinarySize("512"))
	fmt.Println(FromBinarySize("512B"))
	fmt.Println(FromBinarySize("512KiB"))
	fmt.Println(FromBinarySize("512Ki"))
	fmt.Println(FromBinarySize("512kib"))
	fmt.Println(FromBinarySize("512MiB"))
	fmt.Println(FromBinarySize("512Mi"))
	fmt.Println(FromBinarySize("512mi"))
	fmt.Println(FromBinarySize("512GiB"))
	fmt.Println(FromBinarySize("512Gi"))
	fmt.Println(FromBinarySize("512TiB"))
	fmt.Println(FromBinarySize("512Ti"))
	fmt.Println(FromBinarySize("512PiB"))
	fmt.Println(FromBinarySize("512Pi"))
}

func ExampleFromHumanSize() {
	fmt.Println(FromHumanSize("512"))
	fmt.Println(FromHumanSize("512B"))
	fmt.Println(FromHumanSize("512kB"))
	fmt.Println(FromHumanSize("512KB"))
	fmt.Println(FromHumanSize("512K"))
	fmt.Println(FromHumanSize("512MB"))
	fmt.Println(FromHumanSize("512M"))
	fmt.Println(FromHumanSize("512m"))
	fmt.Println(FromHumanSize("512GB"))
	fmt.Println(FromHumanSize("512gB"))
	fmt.Println(FromHumanSize("512TB"))
	fmt.Println(FromHumanSize("512T"))
	fmt.Println(FromHumanSize("512PB"))
	fmt.Println(FromHumanSize("512pb"))
}

func ExampleParseSize() {
	fmt.Println(ParseSize("512"))
	fmt.Println(ParseSize("512B"))
	fmt.Println(ParseSize("512KiB"))
	fmt.Println(ParseSize("512Ki"))
	fmt.Println(ParseSize("512kib"))
	fmt.Println(ParseSize("512MiB"))
	fmt.Println(ParseSize("512Mi"))
	fmt.Println(ParseSize("512mi"))
	fmt.Println(ParseSize("512GiB"))
	fmt.Println(ParseSize("512Gi"))
	fmt.Println(ParseSize("512TiB"))
	fmt.Println(ParseSize("512Ti"))
	fmt.Println(ParseSize("512PiB"))
	fmt.Println(ParseSize("512Pi"))

	fmt.Println(ParseSize("512"))
	fmt.Println(ParseSize("512B"))
	fmt.Println(ParseSize("512kB"))
	fmt.Println(ParseSize("512KB"))
	fmt.Println(ParseSize("512K"))
	fmt.Println(ParseSize("512MB"))
	fmt.Println(ParseSize("512M"))
	fmt.Println(ParseSize("512m"))
	fmt.Println(ParseSize("512GB"))
	fmt.Println(ParseSize("512gB"))
	fmt.Println(ParseSize("512TB"))
	fmt.Println(ParseSize("512T"))
	fmt.Println(ParseSize("512PB"))
	fmt.Println(ParseSize("512pb"))
}

// Benchmark

func BenchmarkParseSize(b *testing.B) {
	units := []string{
		"", "32", "32b", "32 B", "32k", "32.5 K", "32kb", "32 Kb",
		"32.8Mb", "32.9Gb", "32.777Tb", "32Pb", "0.3Mb", "-1",
	}

	for i := 0; i < b.N; i++ {
		for _, s := range units {
			_, _ = ParseSize(s)
		}
	}
}
