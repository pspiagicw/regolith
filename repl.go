package regolith

import (
	"fmt"
	"strings"

	"github.com/chzyer/readline"
)

type Regolith struct {
	config   *Config
	readline *readline.Instance
}

type Config struct {
	Prompt          string
	Name            string
	InterruptPrompt string
	StartWords      []string
	EndWords        []string
	ContinuePrompt  string
}

func (c *Config) Init() {
	if c.Prompt == "" {
		c.Prompt = ">>> "
	}

	if c.Name == "" {
		c.Name = "regolith"
	}

	if c.InterruptPrompt == "" {
		c.InterruptPrompt = "^D"
	}
	if len(c.StartWords) == 0 {
		c.StartWords = []string{"(", "{", "["}
	}

	if len(c.EndWords) == 0 {
		c.EndWords = []string{")", "}", "]"}
	}

	if c.ContinuePrompt == "" {
		c.ContinuePrompt = "..."
	}
}

func New(c *Config) (*Regolith, error) {
	c.Init()
	r, err := readline.NewEx(&readline.Config{
		Prompt:          c.Prompt,
		HistoryFile:     "/tmp/readline.tmp",
		InterruptPrompt: c.InterruptPrompt,
	})
	if err != nil {
		return nil, fmt.Errorf("Error initializing readline: %v", err)
	}
	rg := &Regolith{
		config:   c,
		readline: r,
	}

	return rg, nil
}
func (r *Regolith) read() (string, error) {
	return r.readline.Readline()
}
func (r *Regolith) setPrompt(p string) {
	r.readline.SetPrompt(p)
}
func (r *Regolith) countStart(s string) int {
	total := 0

	for _, word := range r.config.StartWords {
		total += strings.Count(s, word)
	}

	return total
}
func (r *Regolith) countEnd(s string) int {
	total := 0

	for _, word := range r.config.EndWords {
		total += strings.Count(s, word)
	}

	return total
}

func (r *Regolith) Input() (string, error) {
	var cmds []string

	parenScore := 0
	for {
		line, err := r.read()
		if err != nil {
			return "", err
		}
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			continue
		}
		cmds = append(cmds, line)
		parenScore += r.countStart(line) - r.countEnd(line)
		if parenScore == 0 {
			break
		}
		r.setPrompt("... ")
	}
	r.setPrompt(">>> ")

	return strings.Join(cmds, "\n"), nil
}
func (r *Regolith) Close() {
	r.readline.Close()
}
