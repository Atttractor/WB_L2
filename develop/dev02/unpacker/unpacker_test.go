package unpacker

import "testing"

func TestUnpack(t *testing.T) {
	testCases := []struct {
		desc  string
		input string
		want  string
	}{
		{
			desc:  "Цифры и буквы",
			input: "a4df5cvs1",
			want:  "aaaadfffffcvs",
		},
		{
			desc:  "Только буквы",
			input: "abcd",
			want:  "abcd",
		},
		{
			desc:  "Только цифры",
			input: "12",
			want:  "",
		},
		{
			desc:  "Пустая строка",
			input: "",
			want:  "",
		},
		{
			desc:  "Число больше 10",
			input: "a10df5cvs1",
			want:  "aaaaaaaaaadfffffcvs",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got, err := unpack(tC.input)
			if err != nil {
				t.Errorf("Error: %v", err)
			}
			if got != tC.want {
				t.Errorf("got: %s, want: %s", got, tC.want)
			}
		})
	}
}
