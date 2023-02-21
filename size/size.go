package size

import (
	"fmt"
	"github.com/spf13/cast"
	"github.com/thoas/go-funk"
	"regexp"
	"strings"
)

// See: http://en.wikipedia.org/wiki/Binary_prefix
const (
	ByteBase    = 1
	DecimalBase = 1000
	BinaryBase  = 1024

	KB = DecimalBase
	MB = DecimalBase * KB
	GB = DecimalBase * MB
	TB = DecimalBase * GB
	PB = DecimalBase * TB

	KiB = BinaryBase
	MiB = BinaryBase * KiB
	GiB = BinaryBase * MiB
	TiB = BinaryBase * GiB
	PiB = BinaryBase * TiB

	Byte Suffix = "B"

	KiloByte Suffix = "kB"
	MegaByte        = "MB"
	GigaByte        = "GB"
	TeraByte        = "TB"
	PetaByte        = "PB"

	KibiByte Suffix = "KiB"
	MebiByte        = "MiB"
	GibiByte        = "GiB"
	TebiByte        = "TiB"
	PebiByte        = "PiB"

	FormatDefault = "%.4g%s"
)

type Suffix string

// Units provides a specification of the relationship of the suffix to the size.
type Units map[string]uint64

// Suffixes represents the specification of the ratio of size to suffix.
type Suffixes []*struct {
	Unit   uint64
	Suffix Suffix
}

var (
	decimalUnits = Units{
		strings.ToLower(string(Byte)):     ByteBase,
		strings.ToLower(string(KiloByte)): KB,
		strings.ToLower(string(MegaByte)): MB,
		strings.ToLower(string(GigaByte)): GB,
		strings.ToLower(string(TeraByte)): TB,
		strings.ToLower(string(PetaByte)): PB,
	}

	binaryUnits = Units{
		strings.ToLower(string(Byte)):     ByteBase,
		strings.ToLower(string(KibiByte)): KiB,
		strings.ToLower(string(MebiByte)): MiB,
		strings.ToLower(string(GibiByte)): GiB,
		strings.ToLower(string(TebiByte)): TiB,
		strings.ToLower(string(PebiByte)): PiB,
	}

	decimalSuffixes = Suffixes{
		{
			Unit:   ByteBase,
			Suffix: Byte,
		},
		{
			Unit:   KB,
			Suffix: KiloByte,
		},
		{
			Unit:   MB,
			Suffix: MegaByte,
		},
		{
			Unit:   GB,
			Suffix: GigaByte,
		},
		{
			Unit:   TB,
			Suffix: TeraByte,
		},
		{
			Unit:   PB,
			Suffix: PetaByte,
		},
	}

	binarySuffixes = Suffixes{
		{
			Unit:   ByteBase,
			Suffix: Byte,
		},
		{
			Unit:   KiB,
			Suffix: KibiByte,
		},
		{
			Unit:   MiB,
			Suffix: MebiByte,
		},
		{
			Unit:   GiB,
			Suffix: GibiByte,
		},
		{
			Unit:   TiB,
			Suffix: TebiByte,
		},
		{
			Unit:   PiB,
			Suffix: PebiByte,
		},
	}

	DecimalSizeRegexp = regexp.MustCompile(`(?m)^(\d+[\d\.]+?) ?([kKmMgGtTpPbB][bB]?)$`)

	BinarySizeRegexp = regexp.MustCompile(`(?m)^(\d+[\d\.]+?) ?([kKmMgGtTpPbB][iI][bB]?)$`)

	splitRegexp = regexp.MustCompile(`(?m)(\d+[\d\.]+?) ?([A-Za-z]+)$`)
)

// ParseSize defines the IEC/SI prefix and returns int64 as an integer or returns an error if it fails,
// units are case-insensitive, and the 'b' suffix is optional.
func ParseSize(size string) (uint64, error) {
	if DecimalSizeRegexp.MatchString(size) {
		return FromHumanSize(size)
	}

	if BinarySizeRegexp.MatchString(size) {
		return FromBinarySize(size)
	}

	return 0, fmt.Errorf("size: format size '%s' unknown", size)
}

// FromHumanSize returns an integer from a human-readable specification of a
// size using SI standard (eg. "512kB", "20MB") or returns an error if it fails,
// units are case-insensitive, and the 'b' suffix is optional.
func FromHumanSize(size string) (uint64, error) {
	size = strings.TrimSpace(size)

	if !strings.HasSuffix(size, "b") && !strings.HasSuffix(size, "B") {
		size += string(Byte)
	}

	return FromSize(size, decimalUnits)
}

// FromBinarySize parses a human-readable string representing an amount of RAM
// in bytes, kibibytes, mebibytes, gibibytes, or tebibytes and
// returns the number of bytes or returns an error if it fails,
// units are case-insensitive, and the 'b' suffix is optional.
func FromBinarySize(size string) (uint64, error) {
	size = strings.TrimSpace(size)

	if !strings.HasSuffix(size, "b") && !strings.HasSuffix(size, "B") {
		size += string(Byte)
	}

	return FromSize(size, binaryUnits)
}

// FromSize parses the human-readable size string into the amount it represents,
// according to the given specification or returns an error if it fails.
func FromSize(size string, units Units) (uint64, error) {
	matches := splitRegexp.FindAllStringSubmatch(size, -1)
	if len(matches) == 0 {
		return 0, fmt.Errorf("size: invalid format size '%s'", size)
	}

	if len(matches[0]) != 3 {
		return 0, fmt.Errorf("size: invalid format size '%s'", size)
	}

	unit, exist := units[strings.ToLower(strings.TrimSpace(matches[0][2]))]
	if !exist {
		return 0, fmt.Errorf("size: unit '%s' unknown, available units [%s]",
			matches[0][2],
			strings.Join(cast.ToStringSlice(funk.Keys(units)), ", "),
		)
	}

	unitSize, err := cast.ToFloat64E(matches[0][1])
	if err != nil {
		return 0, fmt.Errorf("size: cast invalid: %w", err)
	}

	return uint64(unitSize * float64(unit)), nil
}

// FormatHuman returns a human-readable approximation of a size
// capped at 4 valid numbers (eg. "25MB", "22GB").
func FormatHuman(unit uint64) string {
	return FormatSize(FormatDefault, unit, decimalSuffixes)
}

// FormatBinary returns a human-readable size in bytes, kibibytes,
// mebibytes, gibibytes, or tebibytes (eg. "512kiB", "4PiB").
func FormatBinary(unit uint64) string {
	return FormatSize(FormatDefault, unit, binarySuffixes)
}

// FormatSize returns a human-readable approximation of the size
// using the given format and the given Suffixes specification
func FormatSize(format string, unit uint64, suffixes Suffixes) string {
	size, suffix := calculateSizeAndSuffix(float64(unit), suffixes)
	return fmt.Sprintf(format, size, suffix)
}

func calculateSizeAndSuffix(size float64, suffixes Suffixes) (float64, Suffix) {
	for i := len(suffixes) - 1; i >= 0; i-- {
		if size >= float64(suffixes[i].Unit) {
			return size / float64(suffixes[i].Unit), suffixes[i].Suffix
		}
	}

	return size / float64(suffixes[0].Unit), suffixes[0].Suffix
}
