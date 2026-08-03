package app

import (
	"encoding/json"
	"net/url"
	"testing"

	"github.com/bububa/tiktok-business/enum"
)

func TestOptimizationEventRequest_Encode(t *testing.T) {
	t.Parallel()

	availableOnly := false
	isSkan := true
	req := OptimizationEventRequest{
		AppID:            "222",
		AdvertiserID:     "111",
		Placements:       []enum.Placement{enum.PLACEMENT_TIKTOK, enum.PLACEMENT_PANGLE},
		PlacementType:    enum.PLACEMENT_TYPE_NORMAL,
		OptimizationGoal: enum.OptimizationGoal_IN_APP_EVENT,
		Objective:        enum.ObjectiveType_APP_PROMOTION,
		AvailableOnly:    &availableOnly,
		IsSkan:           &isSkan,
		AppPromotionType: enum.AppPromotionType_APP_INSTALL,
	}

	got, err := url.ParseQuery(req.Encode())
	if err != nil {
		t.Fatalf("ParseQuery() error = %v", err)
	}
	if got.Get("placement") != `["PLACEMENT_TIKTOK","PLACEMENT_PANGLE"]` {
		t.Errorf("placement = %q, want JSON enum array", got.Get("placement"))
	}
	if got.Get("available_only") != "false" {
		t.Errorf("available_only = %q, want explicit false", got.Get("available_only"))
	}
	if got.Get("is_skan") != "true" {
		t.Errorf("is_skan = %q, want true", got.Get("is_skan"))
	}
	if got.Get("optimization_goal") != string(enum.OptimizationGoal_IN_APP_EVENT) {
		t.Errorf("optimization_goal = %q", got.Get("optimization_goal"))
	}
}

func TestConversionEvent_UnmarshalJSON(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		input string
		want  enum.OptimizationEvent
	}{
		{
			name:  "string response",
			input: `"PURCHASE"`,
			want:  enum.OptimizationEvent("PURCHASE"),
		},
		{
			name:  "object response",
			input: `{"optimization_event":"PURCHASE","event_is_available":true}`,
			want:  enum.OptimizationEvent("PURCHASE"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			var got ConversionEvent
			if err := json.Unmarshal([]byte(tt.input), &got); err != nil {
				t.Fatalf("Unmarshal() error = %v", err)
			}
			if got.OptimizationEvent != tt.want {
				t.Errorf("OptimizationEvent = %q, want %q", got.OptimizationEvent, tt.want)
			}
		})
	}
}
