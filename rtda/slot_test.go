package rtda

import (
	"reflect"
	"testing"
)

func Test_inOne(t *testing.T) {
	vars := newLocalVars(100)
	vars.SetInt(0, 100)
	vars.SetInt(1, -100)
	vars.SetLong(2, 2997924580)
	vars.SetLong(4, -2997924580)
	vars.SetFloat(6, 3.1415926)
	vars.SetDouble(7, 2.71828182845)
	vars.SetRef(9, nil)
	println(vars.GetInt(0))
	println(vars.GetInt(1))
	println(vars.GetLong(2))
	println(vars.GetLong(4))
	println(vars.GetFloat(6))
	println(vars.GetDouble(7))
	println(vars.GetRef(9))
}

func Test_newLocalVars(t *testing.T) {
	type args struct {
		maxLocals uint
	}
	tests := []struct {
		name string
		args args
		want LocalVars
	}{
		{
			name: "normal",
			args: args{
				maxLocals: 100,
			},
			want: make([]Slot, 100),
		},
		{
			name: "zero",
			args: args{
				maxLocals: 0,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newLocalVars(tt.args.maxLocals); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newLocalVars() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLocalVars_SetInt(t *testing.T) {
	localVars := newLocalVars(100)
	type args struct {
		index uint
		val   int32
	}
	tests := []struct {
		name string
		self LocalVars
		args args
	}{
		{
			name: "positive",
			self: localVars,
			args: args{
				index: 1,
				val:   100,
			},
		},
		{
			name: "negative",
			self: localVars,
			args: args{
				index: 2,
				val:   -100,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.self.SetInt(tt.args.index, tt.args.val)
		})
	}
}

func TestLocalVars_GetInt(t *testing.T) {
	localVars := newLocalVars(100)
	localVars.SetInt(1, 100)
	localVars.SetInt(2, -100)
	type args struct {
		index uint
	}
	tests := []struct {
		name string
		self LocalVars
		args args
		want int32
	}{
		{
			name: "positive",
			self: localVars,
			args: args{
				1,
			},
			want: 100,
		},
		{
			name: "negative",
			self: localVars,
			args: args{
				2,
			},
			want: -100,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.self.GetInt(tt.args.index); got != tt.want {
				t.Errorf("LocalVars.GetInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLocalVars_SetFloat(t *testing.T) {
	type args struct {
		index uint
		val   float32
	}
	tests := []struct {
		name string
		self LocalVars
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.self.SetFloat(tt.args.index, tt.args.val)
		})
	}
}

func TestLocalVars_GetFloat(t *testing.T) {
	type args struct {
		index uint
	}
	tests := []struct {
		name string
		self LocalVars
		args args
		want float32
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.self.GetFloat(tt.args.index); got != tt.want {
				t.Errorf("LocalVars.GetFloat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLocalVars_SetLong(t *testing.T) {
	type args struct {
		index uint
		val   int64
	}
	tests := []struct {
		name string
		self LocalVars
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.self.SetLong(tt.args.index, tt.args.val)
		})
	}
}

func TestLocalVars_GetLong(t *testing.T) {
	type args struct {
		index uint
	}
	tests := []struct {
		name string
		self LocalVars
		args args
		want int64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.self.GetLong(tt.args.index); got != tt.want {
				t.Errorf("LocalVars.GetLong() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLocalVars_SetDouble(t *testing.T) {
	type args struct {
		index uint
		val   float64
	}
	tests := []struct {
		name string
		self LocalVars
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.self.SetDouble(tt.args.index, tt.args.val)
		})
	}
}

func TestLocalVars_GetDouble(t *testing.T) {
	type args struct {
		index uint
	}
	tests := []struct {
		name string
		self LocalVars
		args args
		want float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.self.GetDouble(tt.args.index); got != tt.want {
				t.Errorf("LocalVars.GetDouble() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLocalVars_SetRef(t *testing.T) {
	type args struct {
		index uint
		ref   *Object
	}
	tests := []struct {
		name string
		self LocalVars
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.self.SetRef(tt.args.index, tt.args.ref)
		})
	}
}

func TestLocalVars_GetRef(t *testing.T) {
	type args struct {
		index uint
	}
	tests := []struct {
		name string
		self LocalVars
		args args
		want *Object
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.self.GetRef(tt.args.index); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LocalVars.GetRef() = %v, want %v", got, tt.want)
			}
		})
	}
}
