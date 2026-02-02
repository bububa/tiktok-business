package changelog

import (
	"encoding/json"
	"testing"

	"github.com/bububa/tiktok-business/enum"
)

func TestTaskDownloadResponse_UnmarshalJSON_Complex(t *testing.T) {
	jsonData := `{"code":0,"message":"OK","request_id":"202602021251040C5E39C8A3618C74BB64","data":{"changelog":"{\"file_data\":\"b'Budget changes,0\\r\\nBid changes,0\\r\\nStatus changes,1\\r\\n\\r\\nTime,log_object_type,Object ID,Object,Operator,Source,App ID,Activity details\\r\\n2026-02-02 04:05,Ad group,1846211451413538,On/Off Status,y***********n@gmail.com,Marketing API,7536109600076988432,\\\"[{\\\"\\\"action\\\"\\\": \\\"\\\"Change \\\"\\\", \\\"\\\"name\\\"\\\": \\\"\\\"Ad group\\\"\\\", \\\"\\\"before_after\\\"\\\": [{\\\"\\\"before\\\"\\\": \\\"\\\"b'Paused'\\\"\\\", \\\"\\\"after\\\"\\\": \\\"\\\"b'Enabled'\\\"\\\"}]}]\\\"\\r\\n'\",\"file_name\":\"7602119211673993224.csv\"}","status":"SUCCESS"}}`

	var resp TaskDownloadResponse
	err := json.Unmarshal([]byte(jsonData), &resp)
	if err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}

	if resp.Code != 0 {
		t.Errorf("Expected code 0, got %d", resp.Code)
	}

	if resp.Message != "OK" {
		t.Errorf("Expected message 'OK', got '%s'", resp.Message)
	}
	ret := resp.Data
	if ret == nil {
		t.Fatal("Expected data to not be nil")
	}
	if ret.Status != enum.ChangelogTaskStatus_SUCCESS {
		t.Errorf("Expected status SUCCESS, got '%s'", ret.Status)
	}

	if ret.Changelog == nil {
		t.Fatal("Expected changelog to not be nil")
	}

	changelog := ret.Changelog
	if changelog.FileName != "7602119211673993224.csv" {
		t.Errorf("Expected filename '7602119211673993224.csv', got '%s'", changelog.FileName)
	}

	if changelog.FileData == nil {
		t.Fatal("Expected file data to not be nil")
	}

	if changelog.FileData.Changes == nil {
		t.Fatal("Expected changes to not be nil")
	}

	if changelog.FileData.Changes.Status != 1 {
		t.Errorf("Expected status changes to be 1, got %d", changelog.FileData.Changes.Status)
	}

	if len(changelog.FileData.Logs) == 0 {
		t.Error("Expected at least one log entry")
	} else {
		logEntry := changelog.FileData.Logs[0]
		if logEntry.ObjectType != "Ad group" {
			t.Errorf("Expected object type 'Ad group', got '%s'", logEntry.ObjectType)
		}
		if logEntry.ObjectID != "1846211451413538" {
			t.Errorf("Expected object ID '1846211451413538', got '%s'", logEntry.ObjectID)
		}
		if len(logEntry.ActivityDetails) == 0 {
			t.Error("Expected at least one activity detail")
		} else {
			detail := logEntry.ActivityDetails[0]
			if detail.Action != "Change " {
				t.Errorf("Expected action 'Change ', got '%s'", detail.Action)
			}
			if detail.Name != "Ad group" {
				t.Errorf("Expected name 'Ad group', got '%s'", detail.Name)
			}
			if len(detail.BeforeAfter) == 0 {
				t.Error("Expected at least one before/after pair")
			} else {
				ba := detail.BeforeAfter[0]
				if ba.Before != "Paused" {
					t.Errorf("Expected before 'Paused', got '%s'", ba.Before)
				}
				if ba.After != "Enabled" {
					t.Errorf("Expected after 'Enabled', got '%s'", ba.After)
				}
			}
		}
	}
}

func TestChangelog_UnmarshalJSON_Direct(t *testing.T) {
	jsonData := `{"file_data":{"changes":{"budget":5,"bid":3,"status":2},"logs":[{"time":"2026-02-02T04:05:00Z","object_type":"Ad group","object_id":"12345","operator":"test@example.com"}]},"file_name":"test.csv"}`

	var changelog Changelog
	err := json.Unmarshal([]byte(jsonData), &changelog)
	if err != nil {
		t.Fatalf("Failed to unmarshal changelog: %v", err)
	}

	if changelog.FileName != "test.csv" {
		t.Errorf("Expected file_name 'test.csv', got '%s'", changelog.FileName)
	}

	if changelog.FileData == nil {
		t.Fatal("Expected file_data to not be nil")
	}

	if changelog.FileData.Changes == nil {
		t.Fatal("Expected changes to not be nil")
	}

	if changelog.FileData.Changes.Budget != 5 {
		t.Errorf("Expected budget changes to be 5, got %d", changelog.FileData.Changes.Budget)
	}
}

func TestParseFileData_Simple(t *testing.T) {
	fileDataStr := `"b'Budget changes,2\\r\\nBid changes,1\\r\\nStatus changes,3\\r\\nTargeting changes,0\\r\\n\\r\\nTime,log_object_type,Object ID,Object,Operator\\r\\n2026-02-02 04:05,Ad,12345,Status,test@example.com'"`

	var result FileData
	err := json.Unmarshal([]byte(fileDataStr), &result)
	if err != nil {
		t.Fatalf("Failed to parse file data: %v", err)
	}

	if result.Changes == nil {
		t.Fatal("nil changes")
	}

	if result.Changes.Budget != 2 {
		t.Errorf("Expected budget changes to be 2, got %d", result.Changes.Budget)
	}
	if result.Changes.Bid != 1 {
		t.Errorf("Expected bid changes to be 1, got %d", result.Changes.Bid)
	}
	if result.Changes.Status != 3 {
		t.Errorf("Expected status changes to be 3, got %d", result.Changes.Status)
	}
	if result.Changes.Targeting != 0 {
		t.Errorf("Expected targeting changes to be 0, got %d", result.Changes.Targeting)
	}
}
