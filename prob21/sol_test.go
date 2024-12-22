package main

import (
	"reflect"
	"testing"
)

func Test_shortestDoorCodePath(t *testing.T) {
	type args struct {
		a rune
		b rune
	}
	tests := []struct {
		name string
		args args
		want []rune
	}{
		{
			name: "1",
			args: args{
				a: 'A',
				b: '4',
			},
			want: []rune("^^<<A"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := shortestDoorCodePath(tt.args.a, tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("shortestDoorCodePath() = %v, want %v", got, tt.want)
			}
		})
	}
}
