package anchor

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/bububa/tiktok-business/enum"
)

func TestCreateRequestEncodeByURL(t *testing.T) {
	req := CreateRequest{AccountID: "account", AnchorType: enum.TTOAnchorType_WEB_ANCHOR, AnchorSubType: enum.TTOAnchorSubType_DESTINATION, UploadType: enum.UPLOAD_BY_URL, ThumbnailURL: "https://example.com/image.png"}
	var body map[string]any
	if err := json.Unmarshal(req.Encode(), &body); err != nil {
		t.Fatal(err)
	}
	if got := body["tto_tcm_account_id"]; got != "account" {
		t.Fatalf("account ID = %v", got)
	}
	if got := body["upload_type"]; got != "UPLOAD_BY_URL" {
		t.Fatalf("upload type = %v", got)
	}
}

func TestCreateByFileRequestEncode(t *testing.T) {
	req := CreateByFileRequest{AccountID: "account", ThumbnailFile: bytes.NewBufferString("image"), ThumbnailFileName: "image.png"}
	fields := req.Encode()
	values := make(map[string]string, len(fields))
	var hasFile bool
	for _, field := range fields {
		values[field.Key] = field.Value
		if field.Key == "thumbnail_file" && field.Reader != nil {
			hasFile = true
		}
	}
	if values["upload_type"] != "UPLOAD_BY_FILE" {
		t.Fatalf("upload type = %q", values["upload_type"])
	}
	if !hasFile {
		t.Fatal("thumbnail file field missing")
	}
}
