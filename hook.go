package main

import (
	"fmt"

	log "github.com/Sirupsen/logrus"
	hookfs "github.com/osrg/hookfs/hookfs"
)

// implements hookfs.HookContext
type InjuredHookContext struct {
	path string
}

// implements hookfs.Hook
type InjuredHook struct{}

// implements hookfs.HookWithInit
func (this *InjuredHook) Init() error {
	log.WithFields(log.Fields{
		"this": this,
	}).Info("InjuredHook Init: initializing")
	return nil
}

// implements hookfs.HookOnOpen
func (this *InjuredHook) PreOpen(path string, flags uint32) (error, bool, hookfs.HookContext) {
	ctx := InjuredHookContext{path: path}
	return nil, false, ctx
}

// implements hookfs.HookOnOpen
func (this *InjuredHook) PostOpen(realRetCode int32, ctx hookfs.HookContext) (error, bool) {
	return nil, false
}

// implements hookfs.HookOnRead
func (this *InjuredHook) PreRead(path string, length int64, offset int64) ([]byte, error, bool, hookfs.HookContext) {
	ctx := InjuredHookContext{path: path}
	return nil, nil, false, ctx
}

// implements hookfs.HookOnRead
func (this *InjuredHook) PostRead(realRetCode int32, realBuf []byte, ctx hookfs.HookContext) ([]byte, error, bool) {
	return nil, nil, false
}

// implements hookfs.HookOnMkdir
func (this *InjuredHook) PreMkdir(path string, mode uint32) (error, bool, hookfs.HookContext) {
	ctx := InjuredHookContext{path: path}
	return nil, false, ctx
}

// implements hookfs.HookOnMkdir
func (this *InjuredHook) PostMkdir(realRetCode int32, ctx hookfs.HookContext) (error, bool) {
	return nil, false
}

// implements hookfs.HookOnRmdir
func (this *InjuredHook) PreRmdir(path string) (error, bool, hookfs.HookContext) {
	ctx := InjuredHookContext{path: path}
	return nil, false, ctx
}

// implements hookfs.HookOnRmdir
func (this *InjuredHook) PostRmdir(realRetCode int32, ctx hookfs.HookContext) (error, bool) {
	return nil, false
}

// implements hookfs.HookOnOpenDir
func (this *InjuredHook) PreOpenDir(path string) (error, bool, hookfs.HookContext) {
	ctx := InjuredHookContext{path: path}
	return nil, false, ctx
}

// implements hookfs.HookOnOpenDir
func (this *InjuredHook) PostOpenDir(realRetCode int32, ctx hookfs.HookContext) (error, bool) {
	return nil, false
}

// implements hookfs.HookOnFsync
func (this *InjuredHook) PreFsync(path string, flags uint32) (error, bool, hookfs.HookContext) {
	ctx := InjuredHookContext{path: path}
	return nil, false, ctx
}

// implements hookfs.HookOnFsync
func (this *InjuredHook) PostFsync(realRetCode int32, ctx hookfs.HookContext) (error, bool) {
	return nil, false
}

var injurers = map[string]InjuredHook{}

func RegisterInjuredHook(name string, hook InjuredHook) {
	_, ok := injurers[name]
	if ok {
		panic(fmt.Sprintf("duplicate register injurer %s", name))
	}

	injurers[name] = hook
}

func GetInjurerCreator(name string) InjuredHook {
	return injurers[name]
}
