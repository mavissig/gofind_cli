package cli

import (
	"flag"
	"log"
)

type utils struct {
	UFind      bool
	UCount     bool
	Utils      map[string]bool
	FindFlags  map[string]bool
	CountFlags map[string]bool
}

type FindUC interface {
	Find()
}

type Cli struct {
	utils
	fuc FindUC
}

func (c *Cli) Run() {
	c.parsingFlagsUFind()
	c.parsingFlagsUCount()
	c.validation()
}

func (c *Cli) parsingFlagsUFind() {
	c.FindFlags["sl"] = *flag.Bool("sl", false, "")
	c.FindFlags["d"] = *flag.Bool("d", false, "")
	c.FindFlags["f"] = *flag.Bool("f", false, "")
	flag.Parse()

	for _, val := range c.FindFlags {
		if val {
			c.UFind = true
		}
	}

	if c.UFind {
		c.Utils["UFind"] = true
	}
}

func (c *Cli) parsingFlagsUCount() {
	c.CountFlags["l"] = *flag.Bool("l", false, "")
	c.CountFlags["m"] = *flag.Bool("m", false, "")
	c.CountFlags["v"] = *flag.Bool("v", false, "")
	flag.Parse()

	for _, val := range c.CountFlags {
		if val && !c.UCount {
			c.UCount = true
		} else if val && c.UCount {
			log.Fatalf("[ERROR]: too many flag")
		}
	}

	if c.UFind {
		c.Utils["UCount"] = true
	}
}

func (c *Cli) validation() {
	check := false
	for _, val := range c.Utils {
		if val && check {
			log.Fatalln("[ERROR]: incompatible flags")
		} else if val && !check {
			check = true
		} else {
			log.Fatalln("[ERROR]: no flags")
		}
	}
}

func New(fuc FindUC) *Cli {
	return &Cli{
		fuc: fuc,
	}
}
