package output

import (
	"bytes"
	"encoding/json"
	"os"
	"strings"
	"testing"
)

func TestPrintJSON(t *testing.T) {
	// Capture stdout
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	data := map[string]string{"id": "123", "name": "test"}
	err := PrintJSON(data)

	w.Close()
	os.Stdout = old

	if err != nil {
		t.Fatalf("PrintJSON() error = %v", err)
	}

	var buf bytes.Buffer
	buf.ReadFrom(r)
	output := buf.String()

	var result map[string]string
	if err := json.Unmarshal([]byte(strings.TrimSpace(output)), &result); err != nil {
		t.Fatalf("output is not valid JSON: %v", err)
	}
	if result["id"] != "123" {
		t.Errorf("id = %q, want %q", result["id"], "123")
	}
}

func TestTableFormatterEmpty(t *testing.T) {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	f := &TableFormatter{}
	err := f.Format(&TableData{
		Headers: []string{"ID", "NAME"},
		Rows:    [][]string{},
	})

	w.Close()
	os.Stdout = old

	if err != nil {
		t.Fatalf("Format() error = %v", err)
	}

	var buf bytes.Buffer
	buf.ReadFrom(r)
	if !strings.Contains(buf.String(), "조회된 데이터가 없습니다") {
		t.Error("expected empty data message")
	}
}

func TestTableFormatterWithData(t *testing.T) {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	f := &TableFormatter{}
	err := f.Format(&TableData{
		Headers: []string{"ID", "NAME"},
		Rows: [][]string{
			{"1", "vpc-a"},
			{"2", "vpc-b"},
		},
	})

	w.Close()
	os.Stdout = old

	if err != nil {
		t.Fatalf("Format() error = %v", err)
	}

	var buf bytes.Buffer
	buf.ReadFrom(r)
	output := buf.String()

	if !strings.Contains(output, "ID") {
		t.Error("output should contain header ID")
	}
	if !strings.Contains(output, "vpc-a") {
		t.Error("output should contain vpc-a")
	}
	if !strings.Contains(output, "vpc-b") {
		t.Error("output should contain vpc-b")
	}
}

func TestTableFormatterWrongType(t *testing.T) {
	f := &TableFormatter{}
	err := f.Format("not a table")
	if err == nil {
		t.Fatal("expected error for wrong type")
	}
}

func TestPrint(t *testing.T) {
	// Test JSON mode
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	data := []string{"a", "b"}
	err := Print("json", nil, data)

	w.Close()
	os.Stdout = old

	if err != nil {
		t.Fatalf("Print(json) error = %v", err)
	}

	var buf bytes.Buffer
	buf.ReadFrom(r)
	if !strings.Contains(buf.String(), `"a"`) {
		t.Error("JSON output should contain 'a'")
	}
}

func TestNewFormatter(t *testing.T) {
	jf := NewFormatter("json")
	if _, ok := jf.(*JSONFormatter); !ok {
		t.Error("expected JSONFormatter")
	}

	tf := NewFormatter("table")
	if _, ok := tf.(*TableFormatter); !ok {
		t.Error("expected TableFormatter")
	}

	df := NewFormatter("unknown")
	if _, ok := df.(*TableFormatter); !ok {
		t.Error("expected TableFormatter for unknown format")
	}
}
