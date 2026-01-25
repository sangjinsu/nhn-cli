package output

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
)

type Formatter interface {
	Format(data interface{}) error
}

type OutputFormat string

const (
	FormatTable OutputFormat = "table"
	FormatJSON  OutputFormat = "json"
)

func NewFormatter(format string) Formatter {
	switch OutputFormat(format) {
	case FormatJSON:
		return &JSONFormatter{}
	default:
		return &TableFormatter{}
	}
}

type TableData struct {
	Headers []string
	Rows    [][]string
}

type TableFormatter struct{}

func (f *TableFormatter) Format(data interface{}) error {
	tableData, ok := data.(*TableData)
	if !ok {
		return fmt.Errorf("테이블 포맷터는 TableData 타입만 지원합니다")
	}

	if len(tableData.Rows) == 0 {
		fmt.Println("조회된 데이터가 없습니다.")
		return nil
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(tableData.Headers)
	table.SetAutoWrapText(false)
	table.SetAutoFormatHeaders(true)
	table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetCenterSeparator("")
	table.SetColumnSeparator("")
	table.SetRowSeparator("")
	table.SetHeaderLine(false)
	table.SetBorder(false)
	table.SetTablePadding("\t")
	table.SetNoWhiteSpace(true)
	table.AppendBulk(tableData.Rows)
	table.Render()

	return nil
}

type JSONFormatter struct{}

func (f *JSONFormatter) Format(data interface{}) error {
	output, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return fmt.Errorf("JSON 직렬화 실패: %w", err)
	}
	fmt.Println(string(output))
	return nil
}

func PrintTable(headers []string, rows [][]string) {
	formatter := &TableFormatter{}
	formatter.Format(&TableData{
		Headers: headers,
		Rows:    rows,
	})
}

func PrintJSON(data interface{}) error {
	formatter := &JSONFormatter{}
	return formatter.Format(data)
}

func Print(format string, tableData *TableData, jsonData interface{}) error {
	switch OutputFormat(format) {
	case FormatJSON:
		return PrintJSON(jsonData)
	default:
		return (&TableFormatter{}).Format(tableData)
	}
}
