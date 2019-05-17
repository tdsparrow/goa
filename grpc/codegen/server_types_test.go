package codegen

import (
	"bytes"
	"testing"

	"goa.design/goa/codegen"
	"goa.design/goa/expr"
	"goa.design/goa/grpc/codegen/testdata"
)

func TestServerTypeFiles(t *testing.T) {
	cases := []struct {
		Name string
		DSL  func()
		Code string
	}{
		{"payload-with-nested-types", testdata.PayloadWithNestedTypesDSL, testdata.PayloadWithNestedTypesServerTypeCode},
		{"result-collection", testdata.ResultWithCollectionDSL, testdata.ResultWithCollectionServerTypeCode},
		{"with-errors", testdata.UnaryRPCWithErrorsDSL, testdata.WithErrorsServerTypeCode},
	}
	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			RunGRPCDSL(t, c.DSL)
			fs := ServerTypeFiles("", expr.Root)
			if len(fs) != 1 {
				t.Fatalf("got %d files, expected one", len(fs))
			}
			var buf bytes.Buffer
			for _, s := range fs[0].SectionTemplates[1:] {
				if err := s.Write(&buf); err != nil {
					t.Fatal(err)
				}
			}
			code := codegen.FormatTestCode(t, "package foo\n"+buf.String())
			if code != c.Code {
				t.Errorf("invalid code, got:\n%s\ngot vs. expected:\n%s", code, codegen.Diff(t, code, c.Code))
			}
		})
	}
}