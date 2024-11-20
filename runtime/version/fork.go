package version

import (
	"github.com/pkg/errors"
	"github.com/prysmaticlabs/prysm/v5/config/params"
	"github.com/prysmaticlabs/prysm/v5/consensus-types/primitives"
)

const (
	Phase0 = iota
	Altair
	Bellatrix
	Capella
	Deneb
	Electra
)

var versionToString = map[int]string{
	Phase0:    "phase0",
	Altair:    "altair",
	Bellatrix: "bellatrix",
	Capella:   "capella",
	Deneb:     "deneb",
	Electra:   "electra",
}

// stringToVersion and allVersions are populated in init()
var stringToVersion = map[string]int{}
var allVersions []int

// ErrUnrecognizedVersionName means a string does not match the list of canonical version names.
var ErrUnrecognizedVersionName = errors.New("version name doesn't map to a known value in the enum")

// FromString translates a canonical version name to the version number.
func FromString(name string) (int, error) {
	v, ok := stringToVersion[name]
	if !ok {
		return 0, errors.Wrap(ErrUnrecognizedVersionName, name)
	}
	return v, nil
}

// FromEpoch translates an epoch into it's corresponding version.
func FromEpoch(epoch primitives.Epoch) int {
	switch {
	case epoch >= params.BeaconConfig().ElectraForkEpoch:
		return Electra
	case epoch >= params.BeaconConfig().DenebForkEpoch:
		return Deneb
	case epoch >= params.BeaconConfig().CapellaForkEpoch:
		return Capella
	case epoch >= params.BeaconConfig().BellatrixForkEpoch:
		return Bellatrix
	case epoch >= params.BeaconConfig().AltairForkEpoch:
		return Altair
	default:
		return Phase0
	}
}

// String returns the canonical string form of a version.
// Unrecognized versions won't generate an error and are represented by the string "unknown version".
func String(version int) string {
	name, ok := versionToString[version]
	if !ok {
		return "unknown version"
	}
	return name
}

// All returns a list of all known fork versions.
func All() []int {
	return allVersions
}

func init() {
	allVersions = make([]int, len(versionToString))
	i := 0
	for v, s := range versionToString {
		allVersions[i] = v
		stringToVersion[s] = v
		i++
	}
}
