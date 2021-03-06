package message

import maybe "github.com/seoester/adcl/protocol/maybe"

// Code generated by adcl/protocol/generator. DO NOT EDIT.

type RESFlag string

const (
	RESFlagFN RESFlag = "FN"
	RESFlagSI         = "SI"
	RESFlagSL         = "SL"
	RESFlagTO         = "TO"
	RESFlagTR         = "TR"
	RESFlagTD         = "TD"
)

var _ ParamAccessor = &RESContent{}

type RESContent struct {
	FN    string
	fnStr string

	SI    int
	siStr string

	SL    maybe.Int
	slStr string

	TO    string
	toStr string

	TR    maybe.Base32Value
	trStr string

	TD    maybe.Int
	tdStr string

	Flags map[string]string

	// FI, FO, DA; EXT § 3.27 ASCH - Extended searching capability (EXT v1.0.8)
}

func (r *RESContent) Positional() []string {
	return []string{}
}

func (r *RESContent) PosLen() int {
	return 0
}

func (r *RESContent) PosAt(i int) string {
	panic("index out of range")
}

func (r *RESContent) Named() map[string]string {
	params := make(map[string]string)

	for key, val := range r.Flags {
		params[key] = val
	}

	params[r.fnStr[:2]] = r.fnStr[2:]
	params[r.siStr[:2]] = r.siStr[2:]
	if r.SL.IsSet {
		params[r.slStr[:2]] = r.slStr[2:]
	}
	params[r.toStr[:2]] = r.toStr[2:]
	if r.TR.IsSet {
		params[r.trStr[:2]] = r.trStr[2:]
	}
	if r.TD.IsSet {
		params[r.tdStr[:2]] = r.tdStr[2:]
	}

	return params
}

func (r *RESContent) NamedGet(key string) (string, bool) {
	switch RESFlag(key) {
	case RESFlagFN:
		return r.fnStr[2:], true
	case RESFlagSI:
		return r.siStr[2:], true
	case RESFlagSL:
		if r.SL.IsSet {
			return r.slStr[2:], true
		} else {
			return "", false
		}
	case RESFlagTO:
		return r.toStr[2:], true
	case RESFlagTR:
		if r.TR.IsSet {
			return r.trStr[2:], true
		} else {
			return "", false
		}
	case RESFlagTD:
		if r.TD.IsSet {
			return r.tdStr[2:], true
		} else {
			return "", false
		}
	}

	key, val := r.Flags[key]
	return key, val
}

