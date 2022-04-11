package uci

type Option struct {
	Name string
	Type OptionType

	Default string
	Min     int
	Max     int
	Options []string
}

func (o Option) DefaultValue() string {
	if o.Default == "" {
		return "<empty>"
	}
	return o.Default
}
