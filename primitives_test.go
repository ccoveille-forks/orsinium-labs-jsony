package jsony_test

import (
	"testing"

	"github.com/orsinium-labs/jsony"
)

func TestNull(t *testing.T) {
	got := jsony.EncodeString(jsony.Null)
	want := `null`
	if got != want {
		t.Fatalf("got `%s`, want `%s`", got, want)
	}
}

func TestBool(t *testing.T) {
	cases := []struct {
		given    bool
		expected string
	}{
		{true, `true`},
		{false, `false`},
	}
	for _, tc := range cases {
		t.Run(tc.expected, func(t *testing.T) {
			got := jsony.EncodeString(jsony.Bool(tc.given))
			if got != tc.expected {
				t.Fatalf("got `%s`, want `%s`", got, tc.expected)
			}
		})
	}
}

func TestInt8(t *testing.T) {
	cases := []struct {
		given    int8
		expected string
	}{
		{13, `13`},
		{0, `0`},
		{1, `1`},
		{9, `9`},
		{-13, `-13`},
	}
	for _, tc := range cases {
		t.Run(tc.expected, func(t *testing.T) {
			got := jsony.EncodeString(jsony.Int8(tc.given))
			if got != tc.expected {
				t.Fatalf("got `%s`, want `%s`", got, tc.expected)
			}
		})
	}
}

func TestInt16(t *testing.T) {
	cases := []struct {
		given    int16
		expected string
	}{
		{13, `13`},
		{0, `0`},
		{1, `1`},
		{9, `9`},
		{156, `156`},
		{-13, `-13`},
	}
	for _, tc := range cases {
		t.Run(tc.expected, func(t *testing.T) {
			got := jsony.EncodeString(jsony.Int16(tc.given))
			if got != tc.expected {
				t.Fatalf("got `%s`, want `%s`", got, tc.expected)
			}
		})
	}
}

func TestInt32(t *testing.T) {
	cases := []struct {
		given    int32
		expected string
	}{
		{13, `13`},
		{0, `0`},
		{1, `1`},
		{9, `9`},
		{156, `156`},
		{156098, `156098`},
		{1560980, `1560980`},
		{-13, `-13`},
	}
	for _, tc := range cases {
		t.Run(tc.expected, func(t *testing.T) {
			got := jsony.EncodeString(jsony.Int32(tc.given))
			if got != tc.expected {
				t.Fatalf("got `%s`, want `%s`", got, tc.expected)
			}
		})
	}
}

func TestInt64(t *testing.T) {
	cases := []struct {
		given    int64
		expected string
	}{
		{13, `13`},
		{0, `0`},
		{1, `1`},
		{9, `9`},
		{156, `156`},
		{156098, `156098`},
		{1560980, `1560980`},
		{-13, `-13`},
	}
	for _, tc := range cases {
		t.Run(tc.expected, func(t *testing.T) {
			got := jsony.EncodeString(jsony.Int64(tc.given))
			if got != tc.expected {
				t.Fatalf("got `%s`, want `%s`", got, tc.expected)
			}
		})
	}
}

func TestInt(t *testing.T) {
	cases := []struct {
		given    int
		expected string
	}{
		{13, `13`},
		{0, `0`},
		{1, `1`},
		{9, `9`},
		{156, `156`},
		{156098, `156098`},
		{1560980, `1560980`},
		{-13, `-13`},
	}
	for _, tc := range cases {
		t.Run(tc.expected, func(t *testing.T) {
			got := jsony.EncodeString(jsony.Int(tc.given))
			if got != tc.expected {
				t.Fatalf("got `%s`, want `%s`", got, tc.expected)
			}
		})
	}
}

func TestUInt(t *testing.T) {
	cases := []struct {
		given    uint
		expected string
	}{
		{13, `13`},
		{0, `0`},
		{1, `1`},
		{9, `9`},
		{156, `156`},
		{156098, `156098`},
		{1560980, `1560980`},
	}
	for _, tc := range cases {
		t.Run(tc.expected, func(t *testing.T) {
			got := jsony.EncodeString(jsony.UInt(tc.given))
			if got != tc.expected {
				t.Fatalf("got `%s`, want `%s`", got, tc.expected)
			}
		})
	}
}

func TestUInt8(t *testing.T) {
	cases := []struct {
		given    uint8
		expected string
	}{
		{13, `13`},
		{0, `0`},
		{1, `1`},
		{9, `9`},
		{156, `156`},
	}
	for _, tc := range cases {
		t.Run(tc.expected, func(t *testing.T) {
			got := jsony.EncodeString(jsony.UInt8(tc.given))
			if got != tc.expected {
				t.Fatalf("got `%s`, want `%s`", got, tc.expected)
			}
		})
	}
}

func TestUInt16(t *testing.T) {
	cases := []struct {
		given    uint16
		expected string
	}{
		{13, `13`},
		{0, `0`},
		{1, `1`},
		{9, `9`},
		{156, `156`},
	}
	for _, tc := range cases {
		t.Run(tc.expected, func(t *testing.T) {
			got := jsony.EncodeString(jsony.UInt16(tc.given))
			if got != tc.expected {
				t.Fatalf("got `%s`, want `%s`", got, tc.expected)
			}
		})
	}
}

func TestUInt32(t *testing.T) {
	cases := []struct {
		given    uint32
		expected string
	}{
		{13, `13`},
		{0, `0`},
		{1, `1`},
		{9, `9`},
		{156, `156`},
		{156098, `156098`},
		{1560980, `1560980`},
	}
	for _, tc := range cases {
		t.Run(tc.expected, func(t *testing.T) {
			got := jsony.EncodeString(jsony.UInt32(tc.given))
			if got != tc.expected {
				t.Fatalf("got `%s`, want `%s`", got, tc.expected)
			}
		})
	}
}

func TestUInt64(t *testing.T) {
	cases := []struct {
		given    uint64
		expected string
	}{
		{13, `13`},
		{0, `0`},
		{1, `1`},
		{9, `9`},
		{156, `156`},
		{156098, `156098`},
		{1560980, `1560980`},
	}
	for _, tc := range cases {
		t.Run(tc.expected, func(t *testing.T) {
			got := jsony.EncodeString(jsony.UInt64(tc.given))
			if got != tc.expected {
				t.Fatalf("got `%s`, want `%s`", got, tc.expected)
			}
		})
	}
}

func TestUIntPtr(t *testing.T) {
	cases := []struct {
		given    uintptr
		expected string
	}{
		{13, `13`},
		{0, `0`},
		{1, `1`},
		{9, `9`},
		{156, `156`},
		{156098, `156098`},
		{1560980, `1560980`},
	}
	for _, tc := range cases {
		t.Run(tc.expected, func(t *testing.T) {
			got := jsony.EncodeString(jsony.UIntPtr(tc.given))
			if got != tc.expected {
				t.Fatalf("got `%s`, want `%s`", got, tc.expected)
			}
		})
	}
}

func TestFloat32(t *testing.T) {
	cases := []struct {
		given    float32
		expected string
	}{
		{13.0, `13`},
		{13.1, `13.1`},
		{0.4, `0.4`},
		{-0.4, `-0.4`},
		{0.0, `0`},
		{1e-7, `1e-07`},
	}
	for _, tc := range cases {
		t.Run(tc.expected, func(t *testing.T) {
			got := jsony.EncodeString(jsony.Float32(tc.given))
			if got != tc.expected {
				t.Fatalf("got `%s`, want `%s`", got, tc.expected)
			}
		})
	}
}

func TestFloat64(t *testing.T) {
	cases := []struct {
		given    float64
		expected string
	}{
		{13.0, `13`},
		{13.1, `13.1`},
		{0.4, `0.4`},
		{-0.4, `-0.4`},
		{0.0, `0`},
		{1e-7, `1e-07`},
	}
	for _, tc := range cases {
		t.Run(tc.expected, func(t *testing.T) {
			got := jsony.EncodeString(jsony.Float64(tc.given))
			if got != tc.expected {
				t.Fatalf("got `%s`, want `%s`", got, tc.expected)
			}
		})
	}
}

func TestComment(t *testing.T) {
	cases := []struct {
		given    jsony.Encoder
		expected string
	}{
		{jsony.Comment("hi"), "//hi\n"},
	}
	for _, tc := range cases {
		t.Run(tc.expected, func(t *testing.T) {
			got := jsony.EncodeString(tc.given)
			if got != tc.expected {
				t.Fatalf("got `%s`, want `%s`", got, tc.expected)
			}
		})
	}
}

func TestString(t *testing.T) {
	cases := []struct {
		given    string
		expected string
	}{
		{"hi", `"hi"`},
		{"", `""`},
		{"hello world!", `"hello world!"`},
		{"hello\n", `"hello\n"`},
		{"hello\\n", `"hello\\n"`},
		{"hi\bsup", `"hi\bsup"`},
		{"hi\fsup", `"hi\fsup"`},
		{"hi\rsup", `"hi\rsup"`},
		{"hi\tsup", `"hi\tsup"`},
		{"привет", `"привет"`},
		{"<b>", `"\u003cb\u003e"`},
		{"\"", `"\""`},
		{"\uFFFD", "\"\uFFFD\""},
		{"\u2028", `"\u2028"`},
	}
	for _, tc := range cases {
		t.Run(tc.expected, func(t *testing.T) {
			got := jsony.EncodeString(jsony.String(tc.given))
			if got != tc.expected {
				t.Fatalf("got `%s`, want `%s`", got, tc.expected)
			}
		})
	}
}

func TestSafeString(t *testing.T) {
	cases := []struct {
		given    jsony.Encoder
		expected string
	}{
		{jsony.SafeString("hi"), `"hi"`},
		{jsony.SafeString(""), `""`},
		{jsony.SafeString("hello world!"), `"hello world!"`},
	}
	for _, tc := range cases {
		t.Run(tc.expected, func(t *testing.T) {
			got := jsony.EncodeString(tc.given)
			if got != tc.expected {
				t.Fatalf("got `%s`, want `%s`", got, tc.expected)
			}
		})
	}
}
