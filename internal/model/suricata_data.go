package model

type SuricataAlert struct {
	Action      string                 `json:"action"`
	GID         int                    `json:"gid"`
	SignatureID int                    `json:"signature_id"`
	Rev         int                    `json:"rev"`
	Signature   string                 `json:"signature"`
	Category    string                 `json:"category"`
	Severity    int                    `json:"severity"`
	Metadata    map[string]interface{} `json:"metadata"`
}

type SuricataMessageData struct {
	Timestamp   string                 `json:"timestamp"`
	FlowID      int                    `json:"flow_id"`
	InInterface string                 `json:"in_iface"`
	EventType   string                 `json:"event_type"`
	SrcIP       string                 `json:"src_ip"`
	SrcPort     int                    `json:"src_port"`
	DestIP      string                 `json:"dest_ip"`
	DestPort    int                    `json:"dest_port"`
	Proto       string                 `json:"proto"`
	PacketSrc   string                 `json:"pkt_src"`
	Alert       SuricataAlert          `json:"alert"`
	HTTP        map[string]interface{} `json:"http"`
	AppProto    string                 `json:"app_proto"`
	Direction   string                 `json:"direction"`
	Stream      int                    `json:"stream"`
	Packet      string                 `json:"packet"`
	Host        string                 `json:"host"`
}
