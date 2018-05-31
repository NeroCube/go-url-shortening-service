package encode

import "testing"

func TestTinyURL(t *testing.T) {
	type args struct {
		random_length int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "TinyURL length is 1", args: args{random_length: 1}, want: 1},
		{name: "TinyURL length is 2", args: args{random_length: 2}, want: 2},
		{name: "TinyURL length is 3", args: args{random_length: 3}, want: 3},
		{name: "TinyURL length is 4", args: args{random_length: 4}, want: 4},
		{name: "TinyURL length is 5", args: args{random_length: 5}, want: 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TinyURL(tt.args.random_length); len(got) != tt.want {
				t.Errorf("TinyURL() = %v, want %v", got, tt.want)
			}
		})
	}
}
