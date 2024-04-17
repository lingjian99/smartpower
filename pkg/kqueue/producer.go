package kqueue

import (
	"context"
	"github.com/segmentio/kafka-go"
	"github.com/zeromicro/go-zero/core/executors"
	"github.com/zeromicro/go-zero/core/logx"
	"strconv"
	"time"
)

type (
	PushOption func(options *chunkOptions)

	Producer struct {
		kq       *kafka.Writer
		topic    string
		executor *executors.ChunkExecutor
	}

	chunkOptions struct {
		chunkSize     int
		flushInterval time.Duration
	}
)

func NewProducer(addrs []string, topic string, opts ...PushOption) *Producer {
	producer := &kafka.Writer{
		Addr:        kafka.TCP(addrs...),
		Topic:       topic,
		Balancer:    &kafka.LeastBytes{},
		Compression: kafka.Snappy,
	}
	pusher := &Producer{
		kq:    producer,
		topic: topic,
	}
	pusher.executor = executors.NewChunkExecutor(func(tasks []interface{}) {
		chunk := make([]kafka.Message, len(tasks))
		for i := range tasks {
			chunk[i] = tasks[i].(kafka.Message)
		}
		if err := pusher.kq.WriteMessages(context.Background(), chunk...); err != nil {
			logx.Error(err)
		}
	}, newOptions(opts)...)

	return pusher
}

func (p *Producer) Close() error {
	if p.executor != nil {
		p.executor.Flush()
	}

	return p.kq.Close()
}

func (p *Producer) Push(topic string, v []byte) error {
	msg := kafka.Message{
		Topic: topic,
		Key:   []byte(strconv.FormatInt(time.Now().UnixNano(), 10)),
		Value: v,
	}
	if p.executor != nil {
		return p.executor.Add(msg, len(v))
	} else {
		return p.kq.WriteMessages(context.Background(), msg)
	}
}

func WithChunkSize(chunkSize int) PushOption {
	return func(options *chunkOptions) {
		options.chunkSize = chunkSize
	}
}

func WithFlushInterval(interval time.Duration) PushOption {
	return func(options *chunkOptions) {
		options.flushInterval = interval
	}
}

func newOptions(opts []PushOption) []executors.ChunkOption {
	var options chunkOptions
	for _, opt := range opts {
		opt(&options)
	}

	var chunkOpts []executors.ChunkOption
	if options.chunkSize > 0 {
		chunkOpts = append(chunkOpts, executors.WithChunkBytes(options.chunkSize))
	}
	if options.flushInterval > 0 {
		chunkOpts = append(chunkOpts, executors.WithFlushInterval(options.flushInterval))
	}
	return chunkOpts
}
