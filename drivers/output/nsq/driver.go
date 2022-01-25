package nsq

import (
	"context"
	"github.com/eolinker/eosc"
	"github.com/eolinker/eosc/formatter"
	"github.com/nsqio/go-nsq"
	"reflect"
	"sync"
)

type Driver struct {
	configType reflect.Type
}

func (d *Driver) ConfigType() reflect.Type {
	return d.configType
}

func (d *Driver) Check(v interface{}) (*NsqConf, error) {
	conf, ok := v.(*Config)
	if !ok {
		return nil, errConfigType
	}

	nsqConf := conf.Config
	if nsqConf == nil {
		return nil, errNsqConfNull
	}
	if nsqConf.Topic == "" {
		return nil, errTopicNull
	}
	if len(nsqConf.Address) == 0 {
		return nil, errAddressNull
	}
	if nsqConf.Type == "" {
		nsqConf.Type = "line"
	}
	switch nsqConf.Type {
	case "line", "json":
	default:
		return nil, errFormatterType
	}

	if len(nsqConf.Formatter) == 0 {
		return nil, errFormatterConf
	}

	return nsqConf, nil
}

func (d *Driver) Create(id, name string, v interface{}, workers map[eosc.RequireId]interface{}) (eosc.IWorker, error) {
	worker := &NsqOutput{
		Driver: d,
		id:     id,
		lock:   sync.Mutex{},
	}

	conf, err := d.Check(v)
	if err != nil {
		return nil, err
	}
	worker.topic = conf.Topic
	//创建producerTransation通道
	worker.ptChannel = make(chan *nsq.ProducerTransaction, 100)

	//创建生产者pool
	worker.pool, err = CreateProducerPool(conf.Address, conf.ClientConf)
	if err != nil {
		return nil, err
	}

	//创建formatter
	factory, has := formatter.GetFormatterFactory(conf.Type)
	if !has {
		return nil, errFormatterType
	}
	worker.formatter, err = factory.Create(conf.Formatter)

	//开始监听返回的异步信息
	ctx, cancelFunc := context.WithCancel(context.Background())
	worker.cancelFunc = cancelFunc
	go worker.listenAsycInfomation(worker.ptChannel, ctx)

	return worker, err
}
