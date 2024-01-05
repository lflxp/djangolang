package addcolumns

import (
	"errors"
	"strings"

	log "github.com/go-eden/slf4go"
)

// string转html
type AddColumns struct {
	Raw string
}

func (a *AddColumns) Transfer() interface{} {
	return func(data map[string]string) string {
		if err := a.Check(data); err != nil {
			log.Error(err)
			return err.Error()
		}
		col := strings.TrimSpace(data["List"]) + " 操作"
		result, err := DirectJson(strings.Split(col, " ")...)
		if err != nil {
			log.Error(err)
			return err.Error()
		}
		return result
	}
}

// 检查并赋值参数
func (a *AddColumns) Check(data map[string]string) error {
	if _, ok := data["List"]; !ok {
		return errors.New("List 参数缺失")
	}

	return nil
}
