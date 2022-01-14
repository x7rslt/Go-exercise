package Command

import (
	"testing"
)

func TestNewPerson(t *testing.T) {
	xiaokong := NewPerson("xiaokong",NewCommand(nil,nil))
	xiaoli := NewPerson("xiaoli",NewCommand(&xiaokong,xiaokong.Listen))
	xiaoni := NewPerson("xiaoni",NewCommand(&xiaoli,xiaoli.Buy))
	xiaozheng := NewPerson("xiaozheng",NewCommand(&xiaoni,xiaoni.Cook))
	xiaozheng.Talk()
}
