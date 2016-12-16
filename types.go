package openbaton

type AutoScalePolicy struct {
	Id                 string          `json:"id"`
	Version            int             `json:"version"`
	Name               string          `json:"name"`
	Threshold          float64         `json:"threshold"`
	ComparisonOperator string          `json:"comparisonOperator"`
	Period             int             `json:"period"`
	Cooldown           int             `json:"cooldown"`
	Mode               ScalingMode     `json:"mode"`
	Type               ScalingType     `json:"type"`
	Alarms             []ScalingAlarm  `json:"alarms"`
	Actions            []ScalingAction `json:"actions"`
}

type ScalingAction struct {
	Id      string            `json:"id"`
	Version int               `json:"version"`
	Type    ScalingActionType `json:"type"`
	Value   string            `json:"value"`
	Target  string            `json:"target"`
}

type ScalingActionType string

const (
	ScaleOut          ScalingActionType = "SCALE_OUT"
	ScaleOutTo        ScalingActionType = "SCALE_OUT_TO"
	ScaleOutToFlavour ScalingActionType = "SCALE_OUT_TO_FLAVOUR"
	ScaleIn           ScalingActionType = "SCALE_IN"
	ScaleInTo         ScalingActionType = "SCALE_IN_TO"
	ScaleInToFlavour  ScalingActionType = "SCALE_IN_TO_FLAVOUR"
)

type ScalingMode string

const (
	ScaleModeReactive   ScalingMode = "REACTIVE"
	ScaleModeProactive  ScalingMode = "PROACTIVE"
	ScaleModePredictive ScalingMode = "PREDICTIVE"
)

type ScalingType string

const (
	ScaleTypeSingle   ScalingType = "SINGLE"
	ScaleTypeVoted    ScalingType = "VOTED"
	ScaleTypeWeighted ScalingType = "WEIGHTED"
)

// A Virtual Network Function Record as described by ETSI GS NFV-MAN 001 V1.1.1
type VNFRecord struct {
	Id                           string                  `json:"id"`
	HbVersion                    int                     `json:"hb_version"`
	AutoScalePolicy              []AutoScalePolicy       `json:"auto_scale_policy"`
	ConnectionPoint              []ConnectionPoint       `json:"connection_point"`
	ProjectId                    string                  `json:"projectId"`
	DeploymentFlavourKey         string                  `json:"deployment_flavour_key"`
	Configurations               Configuration           `json:"configurations"`
	LifecycleEvent               []LifecycleEvent        `json:"lifecycle_event"`
	LifecycleEventHistory        []HistoryLifecycleEvent `json:"lifecycle_event_history"`
	Localization                 string                  `json:"localization"`
	MonitoringParameter          []string                `json:"monitoring_parameter"`
	Vdu                          []VirtualDeploymentUnit `json:"vdu"`
	Vendor                       string                  `json:"vendor"`
	Version                      string                  `json:"version"`
	VirtualLink                  []InternalVirtualLink   `json:"virtual_link"`
	ParentNsId                   string                  `json:"parent_ns_id"`
	DescriptorReference          string                  `json:"descriptor_reference"`
	VnfmId                       string                  `json:"vnfm_id"`
	ConnectedExternalVirtualLink []VirtualLinkRecord     `json:"connected_external_virtual_link"`
	VnfAddress                   []string                `json:"vnf_address"`
	Status                       Status                  `json:"status"`
	Notification                 []string                `json:"notification"`
	AuditLog                     string                  `json:"audit_log"`
	RuntimePolicyInfo            []string                `json:"runtime_policy_info"`
	Name                         string                  `json:"name"`
	Type                         string                  `json:"type"`
	Endpoint                     string                  `json:"endpoint"`
	Task                         string                  `json:"task"`
	Requires                     Configuration           `json:"requires"`
	Provides                     Configuration           `json:"provides"`
	CyclicDependency             bool                    `json:"cyclic_dependency"`
	PackageId                    string                  `json:"packageId"`
}
