package other

import "testing"

func TestMergeString(t *testing.T) {
	type args struct {
		inputs  []string
		pattern string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "comma",
			args: args{
				inputs:  []string{"a", "b"},
				pattern: ",",
			},
			want: "a,b",
		},
		{
			name: "space",
			args: args{
				inputs:  []string{"a", "b"},
				pattern: " ",
			},
			want: "a b",
		},
	}
	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Logf("test %d: %s", i+1, tt.name)
			if got := MergeString(tt.args.inputs, tt.args.pattern); got != tt.want {
				t.Errorf("MergeString() = %v, want %v", got, tt.want)
			}
		})
	}
}
