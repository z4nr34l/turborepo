package context

import (
	"sync"
	"turbo/internal/util"

	"github.com/fatih/color"
)

type colorFn = func(format string, a ...interface{}) string

func getTerminalPackageColors() []colorFn {
	return []colorFn{color.CyanString, color.MagentaString, color.GreenString, color.YellowString, color.BlueString}
}

type ColorCache struct {
	mu         sync.Mutex
	index      int
	TermColors []colorFn
	Cache      map[interface{}]colorFn
}

func NewColorCache() *ColorCache {
	return &ColorCache{
		TermColors: getTerminalPackageColors(),
		index:      0,
		Cache:      make(map[interface{}]colorFn),
	}
}

// PrefixColor returns a color function for a given package name
func (c *ColorCache) PrefixColor(name string) colorFn {
	c.mu.Lock()
	defer c.mu.Unlock()
	colorFn, ok := c.Cache[name]
	if ok {
		return colorFn
	}
	c.index++
	colorFn = c.TermColors[util.PositiveMod(c.index, 5)] // 5 possible colors
	c.Cache[name] = colorFn
	return colorFn
}
