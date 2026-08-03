package app

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/bububa/tiktok-business/enum"
)

func TestTrackRequest_Encode(t *testing.T) {
	t.Parallel()

	timestamp := time.Date(2020, 9, 17, 19, 49, 27, 0, time.UTC)
	req := TrackRequest{
		Event: Event{
			Event:     enum.AppEventTypeAddToCart,
			Timestamp: timestamp,
			Context: &EventContext{
				Device: &Device{
					ATTStatus: enum.AppEventATTStatusDenied,
					Platform:  enum.AppEventPlatformIOS,
				},
				Ad: &Ad{
					IsRetargeting:   "false",
					AttributionType: "click_through",
				},
			},
		},
		TikTokAppID: "123",
		EventSource: enum.AppEventSourceAPI,
		PartnerName: "Segment",
	}

	var got map[string]any
	if err := json.Unmarshal(req.Encode(), &got); err != nil {
		t.Fatalf("Unmarshal(Encode()) error = %v", err)
	}
	if got["timestamp"] != "2020-09-17T19:49:27Z" {
		t.Errorf("timestamp = %v, want RFC3339 timestamp", got["timestamp"])
	}
	if got["event"] != string(enum.AppEventTypeAddToCart) {
		t.Errorf("event = %v, want %q", got["event"], enum.AppEventTypeAddToCart)
	}
	context, ok := got["context"].(map[string]any)
	if !ok {
		t.Fatalf("context = %T, want object", got["context"])
	}
	ad, ok := context["ad"].(map[string]any)
	if !ok {
		t.Fatalf("context.ad = %T, want object", context["ad"])
	}
	if ad["isRetargeting"] != "false" || ad["attributionType"] != "click_through" {
		t.Errorf("context.ad = %#v, want legacy attribution keys", ad)
	}
}

func TestBatchRequest_Encode(t *testing.T) {
	t.Parallel()

	req := BatchRequest{
		TikTokAppID: "123",
		EventSource: enum.AppEventSourceAPI,
		Batch: []Event{{
			Type:      enum.AppEventActionTrack,
			Event:     enum.AppEventTypePurchase,
			Timestamp: time.Date(2020, 9, 17, 19, 49, 27, 0, time.UTC),
		}},
	}

	var got struct {
		Batch []struct {
			Type enum.AppEventAction `json:"type"`
		} `json:"batch"`
	}
	if err := json.Unmarshal(req.Encode(), &got); err != nil {
		t.Fatalf("Unmarshal(Encode()) error = %v", err)
	}
	if len(got.Batch) != 1 || got.Batch[0].Type != enum.AppEventActionTrack {
		t.Fatalf("batch = %#v, want one track event", got.Batch)
	}
}
