package runner

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/WHUPRJ/woj-server/pkg/utils"
	"golang.org/x/text/encoding/charmap"
	"io"
	"path/filepath"
)

const (
	VerdictAccepted = iota
	VerdictWrongAnswer
	VerdictJuryFailed
	VerdictPartialCorrect
	VerdictTimeLimitExceeded
	VerdictMemoryLimitExceeded
	VerdictRuntimeError
	VerdictCompileError
	VerdictSystemError
)

type TestLibReport struct {
	XMLName xml.Name `xml:"result"`
	Outcome string   `xml:"outcome,attr"`
	PCType  int      `xml:"pctype,attr"`
	Points  float64  `xml:"points,attr"`
	Result  string   `xml:",chardata"`
}

type TaskStatus struct {
	Id       int    `json:"id"`
	Points   int32  `json:"points"`
	RealTime int    `json:"real_time"`
	CpuTime  int    `json:"cpu_time"`
	Memory   int    `json:"memory"`
	Verdict  int    `json:"verdict"`
	Message  string `json:"message"`

	infoText  []byte
	info      map[string]interface{}
	judgeText string
	judge     TestLibReport
}

type JudgeStatus struct {
	Message string       `json:"message"`
	Tasks   []TaskStatus `json:"tasks"`
}

func (t *TaskStatus) getInfoText(infoFile string) *TaskStatus {
	if t.Verdict != VerdictAccepted {
		return t
	}

	var err error
	t.infoText, err = utils.FileRead(infoFile)
	if err != nil {
		t.Verdict = VerdictSystemError
		t.Message = "cannot read info file"
	}

	return t
}

func (t *TaskStatus) getInfo() *TaskStatus {
	if t.Verdict != VerdictAccepted {
		return t
	}

	err := json.Unmarshal(t.infoText, &t.info)
	if err != nil {
		t.Verdict = VerdictSystemError
		t.Message = "cannot parse info file"
	} else {
		t.RealTime = int(t.info["real_time"].(float64))
		t.CpuTime = int(t.info["cpu_time"].(float64))
		t.Memory = int(t.info["memory"].(float64))
	}

	return t
}

func (t *TaskStatus) checkExit() *TaskStatus {
	if t.Verdict != VerdictAccepted {
		return t
	}

	if t.info["status"] != "exited" || t.info["code"] != 0.0 {
		t.Verdict = VerdictRuntimeError
		t.Message = fmt.Sprintf("status: %v, code: %v", t.info["status"], t.info["code"])
	}

	return t
}

func (t *TaskStatus) checkTime(config *Config) *TaskStatus {
	if t.Verdict != VerdictAccepted {
		return t
	}

	if t.info["real_time"].(float64) > float64(config.Runtime.TimeLimit)+5 {
		t.Verdict = VerdictTimeLimitExceeded
		t.Message = fmt.Sprintf("real_time: %v cpu_time: %v", t.info["real_time"], t.info["cpu_time"])
	}

	return t
}

func (t *TaskStatus) checkMemory(config *Config) *TaskStatus {
	if t.Verdict != VerdictAccepted {
		return t
	}

	if t.info["memory"].(float64) > float64((config.Runtime.MemoryLimit+1)*1024) {
		t.Verdict = VerdictMemoryLimitExceeded
		t.Message = fmt.Sprintf("memory: %v", t.info["memory"])
	}

	return t
}

func (t *TaskStatus) getJudgeText(judgeFile string) *TaskStatus {
	if t.Verdict != VerdictAccepted {
		return t
	}

	j, err := utils.FileRead(judgeFile)
	if err != nil {
		t.Verdict = VerdictSystemError
		t.Message = "cannot read judge file"
	} else {
		t.judgeText = string(j)
	}

	return t
}

func (t *TaskStatus) getJudge() *TaskStatus {
	if t.Verdict != VerdictAccepted {
		return t
	}

	b := bytes.NewReader([]byte(t.judgeText))
	d := xml.NewDecoder(b)
	d.CharsetReader = func(charset string, input io.Reader) (io.Reader, error) {
		switch charset {
		case "windows-1251":
			return charmap.Windows1251.NewDecoder().Reader(input), nil
		default:
			return nil, fmt.Errorf("unknown charset: %s", charset)
		}
	}

	err := d.Decode(&t.judge)
	if err != nil {
		t.Verdict = VerdictSystemError
		t.Message = "cannot parse judge file"
	}

	return t
}

func (t *TaskStatus) checkJudge(pts *map[int]int32) *TaskStatus {
	if t.Verdict != VerdictAccepted {
		return t
	}

	mp := map[string]int{
		"accepted":           VerdictAccepted,
		"wrong-answer":       VerdictWrongAnswer,
		"presentation-error": VerdictWrongAnswer,
		"points":             VerdictPartialCorrect,
		"relative-scoring":   VerdictPartialCorrect,
	}

	if v, ok := mp[t.judge.Outcome]; ok {
		t.Verdict = v
		t.Message = t.judge.Result
		if v == VerdictAccepted {
			t.Points = (*pts)[t.Id]
		} else if v == VerdictPartialCorrect {
			t.Points = int32(t.judge.Points) + int32(t.judge.PCType)
		}
	} else {
		t.Verdict = VerdictJuryFailed
		t.Message = fmt.Sprintf("unknown outcome: %v, result: %v", t.judge.Outcome, t.judge.Result)
	}

	return t
}

func (s *service) checkResults(user string, config *Config) (JudgeStatus, int32) {
	// CE will be processed in phase compile

	pts := map[int]int32{}
	for _, task := range config.Tasks {
		pts[task.Id] = task.Points
	}

	var results []TaskStatus
	dir := filepath.Join(UserDir, user)
	var sum int32 = 0

	for i := 1; i <= len(config.Tasks); i++ {
		result := TaskStatus{Id: i, Verdict: VerdictAccepted, Points: 0}

		info := filepath.Join(dir, fmt.Sprintf("%d.info", i))
		judge := filepath.Join(dir, fmt.Sprintf("%d.judge", i))

		result.getInfoText(info).
			getInfo().
			checkTime(config).
			checkMemory(config).
			checkExit().
			getJudgeText(judge).
			getJudge().
			checkJudge(&pts)

		sum += result.Points
		results = append(results, result)
	}

	return JudgeStatus{Message: "", Tasks: results}, sum
}
