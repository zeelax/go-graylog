package graylog

// InputType is the implementation of the InputAttributes interface.
func (attrs *InputSyslogUDPAttrs) InputType() string {
	return InputTypeSyslogUDP
}

// InputSyslogUDPAttrs
type InputSyslogUDPAttrs struct {
	BindAddress            string `json:"bind_address,omitempty" v-create:"required" v-update:"required"`
	Port                   int    `json:"port,omitempty" v-create:"required" v-update:"required"`
	RecvBufferSize         int    `json:"recv_buffer_size,omitempty" v-create:"required" v-update:"required"`
	TCPKeepAlive           bool   `json:"tcp_keepalive,omitempty"`
	TLSEnable              bool   `json:"tls_enable,omitempty"`
	ThrottlingAllowed      bool   `json:"throttling_allowed,omitempty"`
	EnableCORS             bool   `json:"enable_cors,omitempty"`
	UseNullDelimiter       bool   `json:"use_null_delimiter,omitempty"`
	ExchangeBind           bool   `json:"exchange_bind,omitempty"`
	ForceRDNS              bool   `json:"force_rdns,omitempty"`
	StoreFullMessage       bool   `json:"store_full_message,omitempty"`
	ExpandStructuredData   bool   `json:"expand_structured_data,omitempty"`
	AllowOverrideDate      bool   `json:"allow_override_date,omitempty"`
	RequeueInvalidMessages bool   `json:"requeue_invalid_messages,omitempty"`
	UseFullNames           bool   `json:"use_full_names,omitempty"`
	TLS                    bool   `json:"tls,omitempty"`
}
