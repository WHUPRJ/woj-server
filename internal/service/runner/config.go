package runner

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

type Config struct {
	Runtime struct {
		TimeLimit   int `json:"TimeLimit"`
		MemoryLimit int `json:"MemoryLimit"`
		NProcLimit  int `json:"NProcLimit"`
	} `json:"Runtime"`
	Languages []struct {
		Lang   string `json:"Lang"`
		Type   string `json:"Type"`
		Script string `json:"Script"`
		Cmp    string `json:"Cmp"`
	} `json:"Languages"`
	Tasks []struct {
		Id     int   `json:"Id"`
		Points int32 `json:"Points"`
	} `json:"Tasks"`
}

func (s *service) ParseConfig(version uint, skipCheck bool) (Config, error) {
	base := filepath.Join(ProblemDir, fmt.Sprintf("%d", version))
	file := filepath.Join(base, "config.json")

	data, err := os.ReadFile(file)
	if err != nil {
		return Config{}, err
	}

	config := Config{}
	err = json.Unmarshal(data, &config)
	if err != nil {
		return Config{}, err
	}

	if skipCheck {
		return config, nil
	}

	err = s.checkConfig(&config, base)
	if err != nil {
		return Config{}, err
	}

	return config, nil
}

func (s *service) checkConfig(config *Config, base string) error {
	if config.Runtime.TimeLimit < 0 {
		return errors.New("time limit is negative")
	}
	if config.Runtime.MemoryLimit < 0 {
		return errors.New("memory limit is negative")
	}
	if config.Runtime.NProcLimit < 0 {
		return errors.New("nproc limit is negative")
	}

	allowedLang := map[string]struct{}{
		"c":   {},
		"cpp": {},
	}
	for _, lang := range config.Languages {
		if _, ok := allowedLang[lang.Lang]; !ok {
			return fmt.Errorf("language %s is not allowed", lang.Lang)
		}

		if lang.Type != "custom" && lang.Type != "default" {
			return fmt.Errorf("language %s has invalid type %s", lang.Lang, lang.Type)
		}

		if lang.Type == "custom" {
			if lang.Script == "" {
				return fmt.Errorf("language %s has empty script", lang.Lang)
			}

			file := filepath.Join(base, "judge", lang.Script)
			_, err := os.Stat(file)
			if err != nil {
				return fmt.Errorf("language %s has invalid script %s", lang.Lang, lang.Script)
			}
		}

		if lang.Type == "default" {
			if lang.Cmp == "" {
				return fmt.Errorf("language %s has empty cmp", lang.Lang)
			}
		}
	}

	if len(config.Tasks) == 0 {
		return errors.New("no tasks")
	}
	ids := map[int]struct{}{}
	total := (1 + len(config.Tasks)) * len(config.Tasks) / 2
	for _, task := range config.Tasks {
		if task.Id <= 0 {
			return fmt.Errorf("task %d has non-positive id", task.Id)
		}

		if task.Points < 0 {
			return fmt.Errorf("task %d has negative points", task.Id)
		}

		if _, ok := ids[task.Id]; ok {
			return fmt.Errorf("task %d has duplicate id", task.Id)
		}

		total -= task.Id
		ids[task.Id] = struct{}{}
	}
	if total != 0 {
		return errors.New("task ids are not continuous")
	}

	return nil
}
