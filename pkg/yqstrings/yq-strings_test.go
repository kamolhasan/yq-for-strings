package yqstrings

import (
	"testing"

	"github.com/mikefarah/yq/v3/pkg/yqlib"
)

func TestMerge(t *testing.T) {
	type args struct {
		arrayMergeStrategy    yqlib.ArrayMergeStrategy
		commentsMergeStrategy yqlib.CommentsMergeStrategy
		overwrite             bool
		inputs                []string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "empty string",
			args: args{
				arrayMergeStrategy:    0,
				commentsMergeStrategy: 0,
				overwrite:             false,
				inputs:                []string{},
			},
			want:    "",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Merge(tt.args.arrayMergeStrategy, tt.args.commentsMergeStrategy, tt.args.overwrite, tt.args.inputs...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Merge() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Merge() got = %v, want %v", got, tt.want)
			}
		})
	}
}
