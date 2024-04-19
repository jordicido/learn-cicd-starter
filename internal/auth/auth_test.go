package auth

import (
	"reflect"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name    string
		headers map[string][]string
		want    string
		wantErr error
	}{
		{
			name:    "no auth header",
			headers: map[string][]string{},
			want:    "",
			wantErr: ErrNoAuthHeaderIncluded,
		},
		{
			name:    "malformed auth header",
			headers: map[string][]string{"Authorization": {"blah"}},
			want:    "",
			wantErr: ErrMalformedAuthHeader,
		},
		{
			name:    "correct header",
			headers: map[string][]string{"Authorization": {"ApiKey 1234"}},
			want:    "1234",
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetAPIKey(tt.headers)
			if !reflect.DeepEqual(err, tt.wantErr) {
				t.Errorf("GetAPIKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetAPIKey() = %v, want %v", got, tt.want)
			}
		})
	}
}
