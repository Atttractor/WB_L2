package mySort

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSort(t *testing.T) {
	testCases := []struct {
		desc string
		app  appEnv
		data []string
		want []string
	}{
		{
			desc: "normal",
			app:  appEnv{},
			data: []string{
				"Standing on one's head at job interviews forms a lasting impression.",
				"The chic gangster liked to start the day with a pink scarf.",
				"He kept telling himself that one day it would all somehow make sense.",
				"Cats are good pets, for they are clean and are not noisy.",
				"I am my aunt's sister's daughter.",
			},
			want: []string{
				"Cats are good pets, for they are clean and are not noisy.",
				"He kept telling himself that one day it would all somehow make sense.",
				"I am my aunt's sister's daughter.",
				"Standing on one's head at job interviews forms a lasting impression.",
				"The chic gangster liked to start the day with a pink scarf.",
			},
		},
		{
			desc: "not numeric order",
			app:  appEnv{},
			data: []string{
				"1",
				"5",
				"13",
				"23",
				"11",
				"21",
				"31",
			},
			want: []string{
				"1",
				"11",
				"13",
				"21",
				"23",
				"31",
				"5",
			},
		},
		{
			desc: "reverse order",
			app: appEnv{
				reverse: true,
			},
			data: []string{
				"1",
				"5",
				"13",
				"23",
				"11",
				"21",
				"31",
			},
			want: []string{
				"5",
				"31",
				"23",
				"21",
				"13",
				"11",
				"1",
			},
		},
		{
			desc: "delete duplicate",
			app: appEnv{
				noCopy: true,
			},
			data: []string{
				"1",
				"1",
				"5",
				"13",
				"23",
				"11",
				"11",
				"21",
				"31",
				"31",
				"31",
			},
			want: []string{
				"1",
				"11",
				"13",
				"21",
				"23",
				"31",
				"5",
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := tC.app.sort(tC.data)
			assert.Equal(t, tC.want, got)
		})
	}
}
