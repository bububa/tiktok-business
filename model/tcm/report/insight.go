package report

type Insight struct {
	// Dimensions 维度数据
	Dimensions *Dimensions `json:"dimensions,omitempty"`
	// Metrics 指标数据
	Metrics *Metrics `json:"metrics,omitempty"`
}
