package kcp

import (
	"encoding/json"

	kcp_util "github.com/go-gost/gost/pkg/common/util/kcp"
	md "github.com/go-gost/gost/pkg/metadata"
)

const (
	defaultBacklog = 128
)

type metadata struct {
	config  *kcp_util.Config
	backlog int
}

func (l *kcpListener) parseMetadata(md md.Metadata) (err error) {
	const (
		backlog = "backlog"
		config  = "config"
	)

	if mm, _ := md.Get(config).(map[interface{}]interface{}); len(mm) > 0 {
		m := make(map[string]interface{})
		for k, v := range mm {
			if sk, ok := k.(string); ok {
				m[sk] = v
			}
		}
		b, err := json.Marshal(m)
		if err != nil {
			return err
		}
		cfg := &kcp_util.Config{}
		if err := json.Unmarshal(b, cfg); err != nil {
			return err
		}
		l.md.config = cfg
	}

	if l.md.config == nil {
		l.md.config = kcp_util.DefaultConfig
	}

	l.md.backlog = md.GetInt(backlog)
	if l.md.backlog <= 0 {
		l.md.backlog = defaultBacklog
	}

	return
}