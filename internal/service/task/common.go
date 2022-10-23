package task

import (
	"github.com/WHUPRJ/woj-server/internal/e"
	"github.com/hibiken/asynq"
	"go.uber.org/zap"
)

func (s *service) submit(typename string, payload []byte, queue string) (*asynq.TaskInfo, e.Status) {
	task := asynq.NewTask(typename, payload)

	info, err := s.queue.Enqueue(task, asynq.Queue(queue))
	if err != nil {
		s.log.Warn("failed to enqueue task", zap.Error(err), zap.Any("task", task))
		return nil, e.TaskEnqueueFailed
	}

	return info, e.Success
}

func (s *service) GetTaskInfo(id string, queue string) (*asynq.TaskInfo, e.Status) {
	task, err := s.inspector.GetTaskInfo(queue, id)
	if err != nil {
		s.log.Debug("get task info failed", zap.Error(err), zap.String("id", id))
		return nil, e.TaskGetInfoFailed
	}
	return task, e.Success
}
