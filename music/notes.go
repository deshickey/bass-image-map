package music

type Notes struct {
	A  []string
	As []string
	B  []string
	C  []string
	Cs []string
	D  []string
	Ds []string
	E  []string
	F  []string
	Fs []string
	G  []string
	Gs []string
}

func (n Notes) GetNotes() Notes {
	notes := Notes{
		A: []string{
			"s6-f10", "s6-f22",
			"s5-f5", "s5-f17",
			"s4-f0", "s4-f12", "s4-f24",
			"s3-f7", "s3-f19",
			"s2-f2", "s2-f14",
			"s1-f9", "s1-f21",
		},
		As: []string{
			"s6-f11", "s6-f23",
			"s5-f6", "s5-f18",
			"s4-f1", "s4-f13",
			"s3-f8", "s3-f20",
			"s2-f3", "s2-f15",
			"s1-f10", "s1-f22",
		},
		B: []string{
			"s6-f0", "s6-f12", "s6-f24",
			"s5-f7", "s5-f19",
			"s4-f2", "s4-f14",
			"s3-f9", "s3-f21",
			"s2-f4", "s2-f16",
			"s1-f11", "s1-f23",
		},
		C: []string{
			"s6-f1", "s6-f13",
			"s5-f8", "s5-f20",
			"s4-f3", "s4-f15",
			"s3-f10", "s3-f22",
			"s2-f5", "s2-f17",
			"s1-f0", "s1-f12", "s1-f24",
		},
		Cs: []string{
			"s6-f2", "s6-f14",
			"s5-f9", "s5-f21",
			"s4-f4", "s4-f16",
			"s3-f11", "s3-f23",
			"s2-f6", "s2-f18",
			"s1-f1", "s1-f13",
		},
		D: []string{
			"s6-f3", "s6-f15",
			"s5-f10", "s5-f22",
			"s4-f5", "s4-f17",
			"s3-f0", "s3-f12", "s3-f24",
			"s2-f7", "s2-f19",
			"s1-f2", "s1-f14",
		},
		Ds: []string{
			"s6-f4", "s6-f16",
			"s5-f11", "s5-f23",
			"s4-f6", "s4-f18",
			"s3-f1", "s3-f13",
			"s2-f8", "s2-f20",
			"s1-f3", "s1-f15",
		},
		E: []string{
			"s6-f5", "s6-f17",
			"s5-f0", "s5-f12", "s5-f24",
			"s4-f7", "s4-f19",
			"s3-f2", "s3-f14",
			"s2-f9", "s2-f21",
			"s1-f4", "s1-f16",
		},
		F: []string{
			"s6-f6", "s6-f18",
			"s5-f1", "s5-f13",
			"s4-f8", "s4-f20",
			"s3-f3", "s3-f15",
			"s2-f10", "s2-f22",
			"s1-f4", "s1-f17",
		},
		Fs: []string{
			"s6-f7", "s6-f19",
			"s5-f2", "s5-f14",
			"s4-f9", "s4-f21",
			"s3-f4", "s3-f16",
			"s2-f11", "s2-f23",
			"s1-f5", "s1-f18",
		},
		G: []string{
			"s6-f8", "s6-f20",
			"s5-f3", "s5-f15",
			"s4-f10", "s4-f22",
			"s3-f5", "s3-f17",
			"s2-f0", "s2-f12", "s2-f24",
			"s1-f6", "s1-f19",
		},
		Gs: []string{
			"s6-f9", "s6-f21",
			"s5-f4", "s5-f16",
			"s4-f11", "s4-f23",
			"s3-f6", "s3-f18",
			"s2-f1", "s2-f13",
			"s1-f7", "s1-f20",
		},
	}
	return notes
}
