package amqp

import (
	"github.com/mcilloni/go-openbaton/vnfm"
	"github.com/mcilloni/go-openbaton/vnfm/channel"
	"github.com/mcilloni/go-openbaton/vnfm/config"
	log "github.com/sirupsen/logrus"
)

func init() {
	vnfm.Register("amqp", amqpDriver{})
}

type amqpDriver struct{}

func (amqpDriver) Init(cnf *config.Config, log *log.Logger) (channel.Channel, error) {
	ret, err := newChannel(cnf, log)
	if err != nil {
		return nil, err
	}

	return ret, nil
}
