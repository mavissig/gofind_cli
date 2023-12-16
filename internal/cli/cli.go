package cli

import (
	"flag"
	"log"
)

type utils struct {
	UFind      bool
	UCount     bool
	UDefault   bool
	Utils      map[string]bool
	FindFlags  map[string]bool
	CountFlags map[string]bool
}

type Find interface {
	Find(map[string]bool, []string)
}

type Cli struct {
	utils
	find Find
}

func (c *Cli) Run() {
	c.Utils = make(map[string]bool)

	c.parsingFlags()
	c.validationFlagsUFind()
	c.validationFlagsUCount()
	c.validationCollision()

	args := flag.Args()

	if c.UFind {
		c.find.Find(c.FindFlags, args)
	} else if c.UDefault {
		c.FindFlags["sl"] = true
		c.FindFlags["d"] = true
		c.FindFlags["f"] = true

		c.find.Find(c.FindFlags, args)
	}
}

func (c *Cli) parsingFlags() {
	sl := flag.Bool("sl", false, "")
	d := flag.Bool("d", false, "")
	f := flag.Bool("f", false, "")
	ext := flag.Bool("ext", false, "")

	l := flag.Bool("l", false, "")
	m := flag.Bool("m", false, "")
	v := flag.Bool("v", false, "")

	flag.Parse()

	c.FindFlags = make(map[string]bool)
	c.FindFlags["sl"] = *sl
	c.FindFlags["d"] = *d
	c.FindFlags["f"] = *f
	c.FindFlags["ext"] = *ext

	c.CountFlags = make(map[string]bool)
	c.CountFlags["l"] = *l
	c.CountFlags["m"] = *m
	c.CountFlags["v"] = *v
}

func (c *Cli) validationFlagsUFind() {
	for _, val := range c.FindFlags {
		if val {
			c.UFind = true
		}
	}

	if c.FindFlags["ext"] && !c.FindFlags["f"] {
		log.Fatalln("[ERROR]: it is not possible to use the -ext flag without the -f flag")
	}

	if c.UFind {
		c.UDefault = false
		c.Utils["UFind"] = true
	}
}

func (c *Cli) validationFlagsUCount() {

	for _, val := range c.CountFlags {
		if val && !c.UCount {
			c.UCount = true
		} else if val && c.UCount {
			log.Fatalf("[ERROR]: too many flag")
		}
	}

	if c.UCount {
		c.UDefault = false
		c.Utils["UCount"] = true
	}
}

func (c *Cli) validationCollision() {
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

func New(fuc Find) *Cli {
	return &Cli{
		find:  fuc,
		utils: utils{UDefault: true},
	}
}
