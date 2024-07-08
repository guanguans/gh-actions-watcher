// Copyright (c) 2024 guanguans<ityaozm@gmail.com>
// For the full copyright and license information, please view
// the LICENSE file that was distributed with this source code.
// https://github.com/guanguans/gh-actions-watcher

package color

import "testing"

func TestColorString(t *testing.T) {
	tests := []struct {
		name string
		x    color
		want string
	}{
		{"empty", color(""), ""},
		{"red", color("#ff0000"), "#ff0000"},
		{"invalid", color("invalid"), "invalid"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.x.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestColorIsValid(t *testing.T) {
	tests := []struct {
		name string
		x    color
		want bool
	}{
		{"empty", color(""), false},
		{"red", color("#ff0000"), true},
		{"invalid", color("invalid"), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.x.IsValid(); got != tt.want {
				t.Errorf("IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParsecolor(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		args    args
		want    color
		wantErr bool
	}{
		{"empty", args{name: ""}, color(""), true},
		{"red", args{name: "#ff0000"}, color("#ff0000"), false},
		{"invalid", args{name: "invalid"}, color(""), true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Parsecolor(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parsecolor() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Parsecolor() got = %v, want %v", got, tt.want)
			}
		})
	}
}
