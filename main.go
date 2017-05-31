package main

import (
	"flag"
	"os"
	"strconv"
	
	"github.com/go-ini/ini"
)

const (
	emptyConfigSection   = ""
	defaultConfigSection = "*"

	rKey = "root"
	cKey = "charset"
	iKey = "indent_size"
	eKey = "end_of_line"
	sKey = "indent_style"
	fKey = "insert_final_newline"
	wKey = "trim_trailing_whitespace"
)

var (
	r = flag.Bool("r", true, "The root config value.")
	c = flag.String("c", "UTF-8", "The charset config value.")
	i = flag.Int("i", 4, "The indent-size config value.")
	e = flag.String("e", "lf", "The end of line config value.")
	s = flag.String("s", "space", "The indent_style config value.")
	f = flag.Bool("t", true, "The `insert_final_newline` config value.")
	w = flag.Bool("w", true, "The `trim_trailing_whitespace` config value")
)

func main() {
	flag.Parse()

	cfg := ini.Empty()

	defineRootConfiguration(cfg)

	defineOtherConfigurations(cfg)

	cfg.WriteTo(os.Stdout)
}

func defineRootConfiguration(cfg *ini.File) {
	if *r == true {
		cfg.Section(emptyConfigSection).NewKey(rKey, strconv.FormatBool(*r))
	}
}

func defineOtherConfigurations(cfg *ini.File) {
	o := cfg.Section(defaultConfigSection)

	o.NewKey(cKey, *c)

	o.NewKey(iKey, strconv.Itoa(*i))

	o.NewKey(eKey, *e)

	o.NewKey(sKey, *s)

	o.NewKey(fKey, strconv.FormatBool(*f))

	o.NewKey(wKey, strconv.FormatBool(*w))
}
