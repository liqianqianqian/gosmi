package gosmi

import (
	"fmt"

	"github.com/sleepinggenius2/gosmi/models"
	"github.com/sleepinggenius2/gosmi/smi"
	"github.com/sleepinggenius2/gosmi/types"
)

// 保留兼容性定义
type SmiModule struct {
	models.Module
	smiModule *types.SmiModule
}

// 对外暴露的最终模块结构体
type Module struct {
	Module models.Module
}

// CreateModule: 内部辅助函数，把 types.SmiModule 转换为 models.Module
func CreateModule(smiModule *types.SmiModule) Module {
	return Module{
		Module: models.CreateModule(smiModule),
	}
}

// LoadModule: 你私有 fork 最核心的兼容性接口
func LoadModule(modulePath string) (*Module, error) {
	moduleName := smi.LoadModule(modulePath)
	if moduleName == "" {
		return nil, fmt.Errorf("Could not load module at %s", modulePath)
	}

	smiModule := smi.GetModule(moduleName)
	if smiModule == nil {
		return nil, fmt.Errorf("Could not get module %s", moduleName)
	}

	module := &Module{
		Module: models.CreateModule(smiModule),
	}

	return module, nil
}

// GetLoadedModules: 获取所有已加载模块 (如果需要可扩展)
func GetLoadedModules() []Module {
	var modules []Module
	for smiModule := smi.GetFirstModule(); smiModule != nil; smiModule = smi.GetNextModule(smiModule) {
		modules = append(modules, CreateModule(smiModule))
	}
	return modules
}

// 兼容性辅助函数（保留，可选）
func IsLoaded(moduleName string) bool {
	return smi.IsLoaded(moduleName)
}

func GetModule(moduleName string) (*Module, error) {
	smiModule := smi.GetModule(moduleName)
	if smiModule == nil {
		return nil, fmt.Errorf("Could not get module %s", moduleName)
	}
	module := &Module{
		Module: models.CreateModule(smiModule),
	}
	return module, nil
}
