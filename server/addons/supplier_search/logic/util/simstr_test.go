package util

import (
	"github.com/antlabs/strsim"
	"testing"
)

func TestStrSim(t *testing.T) {
	result := strsim.FindBestMatch("一层  ", []string{"白日依山尽", "黄河入海流", "欲穷千里目", "更上一层楼"})
	t.Log(result.Match.S)
}
