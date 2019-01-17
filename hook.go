package main

import (
	"github.com/ethercflow/hookfs/hookfs"
)

// implements hookfs.HookContext
type InjuredHookContext struct {
}

type InjuredHook struct {
	addr string
}

// implements hookfs.HookWithInit
func (h *InjuredHook) Init() error {
	StartServer(h.addr)
	return nil
}

// if hooked is true, the real open() would not be called
func (h *InjuredHook) PreOpen(path string, flags uint32) (bool, hookfs.HookContext, error) {
	ctx := &InjuredHookContext{}
	err := faultInject(path, "open")
	if err != nil {
		return true, ctx, err
	}
	return false, ctx, nil
}

func (h *InjuredHook) PostOpen(int32, hookfs.HookContext) (bool, error) {
	return false, nil
}

// if hooked is true, the real read() would not be called
func (h *InjuredHook) PreRead(path string, length int64, offset int64) ([]byte, bool, hookfs.HookContext, error) {
	ctx := &InjuredHookContext{}
	err := faultInject(path, "read")
	if err != nil {
		return nil, true, ctx, err
	}
	return nil, false, ctx, nil
}

func (h *InjuredHook) PostRead(realRetCode int32, realBuf []byte, prehookCtx hookfs.HookContext) ([]byte, bool, error) {
	return nil, false, nil
}

// if hooked is true, the real write() would not be called
func (h *InjuredHook) PreWrite(path string, buf []byte, offset int64) (bool, hookfs.HookContext, error) {
	ctx := &InjuredHookContext{}
	err := faultInject(path, "write")
	if err != nil {
		return true, ctx, err
	}
	return false, ctx, nil
}

func (h *InjuredHook) PostWrite(realRetCode int32, prehookCtx hookfs.HookContext) (bool, error) {
	return false, nil
}

// if hooked is true, the real mkdir() would not be called
func (h *InjuredHook) PreMkdir(path string, mode uint32) (bool, hookfs.HookContext, error) {
	ctx := &InjuredHookContext{}
	err := faultInject(path, "mkdir")
	if err != nil {
		return true, ctx, err
	}
	return false, ctx, nil
}

func (h *InjuredHook) PostMkdir(realRetCode int32, prehookCtx hookfs.HookContext) (bool, error) {
	return false, nil
}

// if hooked is true, the real rmdir() would not be called
func (h *InjuredHook) PreRmdir(path string) (bool, hookfs.HookContext, error) {
	ctx := &InjuredHookContext{}
	err := faultInject(path, "rmdir")
	if err != nil {
		return true, ctx, err
	}
	return false, ctx, nil
}

func (h *InjuredHook) PostRmdir(realRetCode int32, prehookCtx hookfs.HookContext) (bool, error) {
	return false, nil
}

// if hooked is true, the real opendir() would not be called
func (h *InjuredHook) PreOpenDir(path string) (bool, hookfs.HookContext, error) {
	ctx := &InjuredHookContext{}
	err := faultInject(path, "opendir")
	if err != nil {
		return true, ctx, err
	}
	return false, ctx, nil
}

func (h *InjuredHook) PostOpenDir(realRetCode int32, prehookCtx hookfs.HookContext) (bool, error) {
	return false, nil
}

// if hooked is true, the real fsync() would not be called
func (h *InjuredHook) PreFsync(path string, flags uint32) (bool, hookfs.HookContext, error) {
	ctx := &InjuredHookContext{}
	err := faultInject(path, "fsync")
	if err != nil {
		return true, ctx, err
	}
	return false, ctx, nil
}

func (h *InjuredHook) PostFsync(realRetCode int32, prehookCtx hookfs.HookContext) (bool, error) {
	return false, nil
}