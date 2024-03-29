package grafana

type Alert struct {
	Tags        map[string]string `json:"tags"`
	EvalMatches []*EvalMatch      `json:"evalMatches"`
	Title       string            `json:"title"`
	ImageUrl    string            `json:"imageUrl"`
	DashboardId int               `json:"dashboardId"`
	Message     string            `json:"message"`
	OrgId       int               `json:"orgId"`
	PanelId     int               `json:"panelId"`
	RuleId      int               `json:"ruleId"`
	RuleName    string            `json:"ruleName"`
	RuleUrl     string            `json:"ruleUrl"`
	State       string            `json:"state"`
}

type EvalMatch struct {
	Value  float64           `json:"value"`
	Metric string            `json:"metric"`
	Tags   map[string]string `json:"tags"`
}

type StateEnum struct {
	PAUSED   string `enum:"paused"`
	ALERTING string `enum:"alerting"`
	PENDING  string `enum:"pending"`
	NODATA   string `enum:"no_data"`
}
