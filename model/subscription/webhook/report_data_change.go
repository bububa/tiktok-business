package webhook

import (
	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model/report"
)

type ReportDataChangeEntity struct {
	// The data level of the subscription.
	DataLevel string `json:"data_level,omitempty"`
	// Message An additional message about this subscription.
	// For example, you will receive a message when your access to some of your subscribed ad accounts is revoked.
	// By default, message is not returned.
	Message string `json:"message,omitempty"`
	// ChangedMetricsValues A map showing the details of the reporting metric data changes for each ad account.
	// Key (string): The ID of the ad account (advertiser_id) that you subscribe to.
	// Value (object): The changed metrics for the ad account within the last subscription notification timeframe and the values for the metrics. The metric values are queried based on the day in UTC+0 time when the metric changed.
	// To retrieve the notification timeframe of a subscription, use /subscription/get/ and check the returned notify_frequency.
	ChangedMetricsValues map[string]report.Metrics `json:"changed_metrics_values,omitempty"`
}

func (e ReportDataChangeEntity) Entity() enum.SubscribeEntity {
	return enum.SubscribeEntity_REPORT_DATA_CHANGE
}
