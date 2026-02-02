package changelog

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

var (
	commaSep    = []byte{','}
	newLine     = []byte{'\n'}
	singleQuote = []byte{'"'}
)

// TaskDownloadRequest Get the downloaded file API Request
type TaskDownloadRequest struct {
	// AdvertiserID Advertiser ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// TaskID Task ID
	TaskID string `json:"task_id,omitempty"`
}

// Encode implements GetRequest interface
func (r TaskDownloadRequest) Encode() string {
	values := util.NewURLValues()
	values.Set("advertiser_id", r.AdvertiserID)
	values.Set("task_id", r.TaskID)
	ret := values.Encode()
	util.ReleaseURLValues(values)
	return ret
}

// TaskDownloadResponse Get the downloaded file API Response
type TaskDownloadResponse struct {
	model.BaseResponse
	Data *TaskDownloadResult `json:"data,omitempty"`
}

type TaskDownloadResult struct {
	// Status Task status
	Status enum.ChangelogTaskStatus `json:"status,omitempty"`
	// Changelog CSV file data of the log
	Changelog *Changelog `json:"changelog,omitempty"`
}

func (r *TaskDownloadResult) UnmarshalJSON(data []byte) error {
	var ret struct {
		// Status Task status
		Status enum.ChangelogTaskStatus `json:"status,omitempty"`
		// Changelog CSV file data of the log
		Changelog string `json:"changelog,omitempty"`
	}
	if err := json.Unmarshal(data, &ret); err != nil {
		return err
	}
	var changelog Changelog
	if err := json.Unmarshal([]byte(ret.Changelog), &changelog); err != nil {
		return err
	}
	r.Status = ret.Status
	r.Changelog = &changelog
	return nil
}

// Changelog need custom JSON Unmarshaler
type Changelog struct {
	FileData *FileData `json:"file_data,omitempty"`
	FileName string    `json:"file_name,omitempty"`
}

// UnmarshalJSON handles the custom unmarshaling for TaskDownloadResponse
// data may like: "{"file_data":"b'Budget changes,0\\r\\nBid changes,0\\…}]\\"\\r\\n'","file_name":"7602119211673993224.csv"}"
func (c *Changelog) UnmarshalJSON(data []byte) error {
	if data[0] == '"' && data[len(data)-1] == '"' {
		data = data[1 : len(data)-1]
	}
	var tmp struct {
		FileData *FileData `json:"file_data,omitempty"`
		FileName string    `json:"file_name,omitempty"`
	}
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	c.FileData = tmp.FileData
	c.FileName = tmp.FileName
	return nil
}

// FileData need custom JSON Unmarshaler
type FileData struct {
	Changes *ChangesCount `json:"changes,omitempty"`
	Logs    []Log         `json:"logs,omitempty"`
}

// UnmarshalJSON unmarshals the FileData from a JSON string that contains escaped CSV data
// data may like: "b'Budget changes,0\r\nBid changes,0\r\nStatus changes,1\r\n\r\nTime,log_object_type,Object ID,Object,Operator,Source,App ID,Activity details\r\n2026-02-02 04:05,Ad group,1846211451413538,On/Off Status,y***********n@gmail.com,Marketing API,7536109600076988432,\"[{\"\"action\"\": \"\"Change \"\", \"\"name\"\": \"\"Ad group\"\", \"\"before_after\"\": [{\"\"before\"\": \"\"b'Paused'\"\", \"\"after\"\": \"\"b'Enabled'\"\"}]}]\"\r\n'"
func (f *FileData) UnmarshalJSON(data []byte) error {
	// The fileDataStr looks like: "b'Budget changes,0\r\nBid changes,0\r\nStatus changes,1\r\n..."
	// We need to decode this CSV-like format
	if data[0] == '"' && data[len(data)-1] == '"' {
		data = data[1 : len(data)-1]
		data = cleanPythonBytes(data)

		// Replace the literal \r\n sequences with actual newlines for proper parsing
		data = fixEscapedNewlines(data)

		dirtyLines := bytes.Split(data, newLine)
		headerIndex := -1
		lines := make([][]byte, 0, len(dirtyLines))
		// Find where the counts section ends and the CSV section begins
		for _, line := range dirtyLines {
			line = bytes.TrimSpace(line)
			if len(line) == 0 {
				continue
			}
			lines = append(lines, line)
		}
		if len(lines) < 2 {
			return nil
		}
		for i, line := range lines {
			// Look for the header line to determine when the counts section ends
			if bytes.HasPrefix(line, []byte("Time,")) {
				headerIndex = i
				break
			}
		}

		countLines := headerIndex
		if countLines < 0 {
			countLines = len(lines)
		}
		for i := range countLines {
			countLine := bytes.TrimSpace(lines[i])
			if len(countLine) == 0 {
				continue
			}
			parts := bytes.SplitN(countLine, commaSep, 2) // Split into max 2 parts to avoid splitting values with commas
			if len(parts) == 2 {
				key := string(bytes.TrimSpace(parts[0]))
				value := string(bytes.TrimSpace(parts[1]))

				switch key {
				case "Budget changes":
					if val, err := strconv.Atoi(value); err == nil {
						if f.Changes == nil {
							f.Changes = new(ChangesCount)
						}
						f.Changes.Budget = val
					}
				case "Bid changes":
					if val, err := strconv.Atoi(value); err == nil {
						if f.Changes == nil {
							f.Changes = new(ChangesCount)
						}
						f.Changes.Bid = val
					}
				case "Status changes":
					if val, err := strconv.Atoi(value); err == nil {
						if f.Changes == nil {
							f.Changes = new(ChangesCount)
						}
						f.Changes.Status = val
					}
				case "Targeting changes":
					if val, err := strconv.Atoi(value); err == nil {
						if f.Changes == nil {
							f.Changes = new(ChangesCount)
						}
						f.Changes.Targeting = val
					}
				}
			}
		}
		// If we found a header, parse the counts before it
		if headerIndex < 0 {
			return nil
		}
		if logs, err := parseCSVLog(lines[headerIndex:]); err != nil {
			return fmt.Errorf("parse csv log failed: %w", err)
		} else {
			f.Logs = logs
		}
		return nil
	}
	var tmp struct {
		Changes *ChangesCount `json:"changes,omitempty"`
		Logs    []Log         `json:"logs,omitempty"`
	}
	if err := json.Unmarshal(data, &tmp); err != nil {
		return fmt.Errorf("unmarshal FileData failed: %w", err)
	}
	f.Changes = tmp.Changes
	f.Logs = tmp.Logs
	return nil
}

// ChangesCount parsed from file_data first changes count part
// Need custom JSON Unmarshaler
type ChangesCount struct {
	Budget    int `json:"budget,omitempty" csv:"Budget changes"`
	Bid       int `json:"bid,omitempty" csv:"Bid changes"`
	Status    int `json:"status,omitempty" csv:"Status changes"`
	Targeting int `json:"targeting,omitempty" csv:"Targeting changes"`
}

// Log parsed from file_data csv part
// Need custom JSON Unmarshaler
type Log struct {
	Time            model.DateTime   `json:"time,omitzero"`
	ObjectType      string           `json:"log_object_type,omitempty"`
	ObjectID        string           `json:"object_id,omitempty"`
	Object          string           `json:"object,omitempty"`
	Operator        string           `json:"operator,omitempty"`
	Source          string           `json:"source,omitempty"`
	AppID           string           `json:"app_id,omitempty"`
	ActivityDetails []ActivityDetail `json:"activity_details,omitempty"`
}

func parseCSVLog(lines [][]byte) ([]Log, error) {
	if len(lines) < 2 {
		return nil, nil
	}
	// 2. Map header names to slice indices
	// slice size is the number of Log structure fields
	headers := make([]string, 0, 8)
	for key := range bytes.SplitSeq(lines[0], commaSep) {
		headers = append(headers, string(bytes.ToLower(bytes.TrimSpace(key))))
	}
	headerCount := len(headers)
	list := make([]Log, 0, len(lines)-1)
	for i := 1; i < len(lines); i++ {
		parts := bytes.SplitN(lines[i], commaSep, headerCount)
		var l Log
		for idx, part := range parts {
			if idx >= headerCount {
				break
			}
			key := headers[idx]
			part := bytes.TrimSpace(part)
			switch key {
			case "time":
				if err := (&l.Time).UnmarshalJSON(part); err != nil {
					return nil, fmt.Errorf("parse csv log time failed: %w", err)
				}
			case "log_object_type":
				l.ObjectType = string(part)
			case "object id":
				l.ObjectID = string(part)
			case "object":
				l.Object = string(part)
			case "operator":
				l.Operator = string(part)
			case "source":
				l.Source = string(part)
			case "app id":
				l.AppID = string(part)
			case "activity details":
				if part[0] == '"' && part[len(part)-1] == '"' {
					part = part[1 : len(part)-1]
				}
				part = bytes.ReplaceAll(part, []byte{'"', '"'}, singleQuote)
				if err := json.Unmarshal(part, &l.ActivityDetails); err != nil {
					return nil, err
				}
			}
		}
		list = append(list, l)
	}
	return list, nil
}

// ActivityDetail parsed from csv:Activity details
type ActivityDetail struct {
	Action      string        `json:"action,omitempty"`
	Name        string        `json:"name,omitempty"`
	BeforeAfter []BeforeAfter `json:"before_after,omitempty"`
}

type BeforeAfter struct {
	Before String `json:"before,omitempty"`
	After  String `json:"after,omitempty"`
}

type String string

func (s *String) UnmarshalJSON(data []byte) error {
	if data[0] == '"' && data[len(data)-1] == '"' {
		data = data[1 : len(data)-1]
	}
	bs := cleanPythonBytes(data)
	*s = String(bs)
	return nil
}

func cleanPythonBytes(input []byte) []byte {
	// 1. 最小长度校验: b'' 至少 3 个字节
	if len(input) < 3 {
		return input
	}

	// 2. 检查前缀 b' 或 b" 以及对应的后缀
	// Python bytes 表现形式可能是 b'...' 或 b"..."
	hasPrefix := input[0] == 'b' && (input[1] == '\'' || input[1] == '"')
	if !hasPrefix {
		return input
	}

	quoteChar := input[1]
	if input[len(input)-1] != quoteChar {
		return input // 引号不闭合，不处理
	}

	// 3. 提取中间的内容
	content := input[2 : len(input)-1]

	// 4. 处理转义字符
	// strconv.Unquote 要求内容必须被双引号包裹
	// 我们构造一个临时 buffer: "内容"
	wrapper := util.NewBuffer()
	defer util.ReleaseBuffer(wrapper)
	wrapper.WriteByte('"')
	wrapper.Write(content)
	wrapper.WriteByte('"')

	// Unquote 会返回 string，这是处理转义最安全的方式
	unquoted, err := strconv.Unquote(wrapper.String())
	if err != nil {
		// 如果解析失败（例如内部有不合法的转义），则返回提取出的原始内容副本
		res := make([]byte, len(content))
		copy(res, content)
		return res
	}

	return []byte(unquoted)
}

func fixEscapedNewlines(input []byte) []byte {
	// 替换字面量的 \r\n 为真正的换行符
	input = bytes.ReplaceAll(input, []byte(`\r\n`), newLine)
	// 替换字面量的 \n 为真正的换行符 (以防万一)
	input = bytes.ReplaceAll(input, []byte(`\n`), newLine)
	return input
}
