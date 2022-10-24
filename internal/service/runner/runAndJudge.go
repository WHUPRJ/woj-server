package runner

import "github.com/WHUPRJ/woj-server/internal/e"

func (s *service) RunAndJudge(version uint, user string, lang string, config *Config) (JudgeStatus, int32, e.Status) {
	// run user program
	status := s.checkAndExecute(version, user, lang, "problem_run.sh", e.RunnerRunFailed)
	if status != e.Success {
		return JudgeStatus{Message: "run failed"}, 0, status
	}

	// run judger
	status = s.checkAndExecute(version, user, lang, "problem_judge.sh", e.RunnerJudgeFailed)
	if status != e.Success {
		return JudgeStatus{Message: "judge failed"}, 0, status
	}

	// check result
	result, pts := s.checkResults(user, config)

	return result, pts, e.Success
}
