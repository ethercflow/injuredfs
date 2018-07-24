package main

import (
	"strings"
	"sync"
	"syscall"
	"time"

	hookfs "github.com/osrg/hookfs/hookfs"
	log "github.com/sirupsen/logrus"
)

// implements hookfs.HookContext
type InjuredHookContext struct {
	path string
}

// implements hookfs.Hook
type InjuredHook struct {
	read map[string]time.Duration
	rl   *sync.RWMutex

	fsync map[string]time.Duration
	fl    *sync.RWMutex
}

// implements hookfs.HookWithInit
func (this *InjuredHook) Init() error {
	log.WithFields(log.Fields{
		"this": this,
	}).Info("InjuredHook Init: initializing")
	return nil
}

// implements hookfs.HookOnOpen
func (this *InjuredHook) PreOpen(path string, flags uint32) (bool, hookfs.HookContext, error) {
	ctx := InjuredHookContext{path: path}
	return false, ctx, nil
}

// implements hookfs.HookOnOpen
func (this *InjuredHook) PostOpen(realRetCode int32, ctx hookfs.HookContext) (bool, error) {
	return false, nil
}

// implements hookfs.HookOnRead
func (this *InjuredHook) PreRead(path string, length int64, offset int64) ([]byte, bool, hookfs.HookContext, error) {
	ctx := InjuredHookContext{path: path}

	this.rl.RLock()
	defer this.rl.RUnlock()

	t, ok := this.read[strings.Split(path, "/")[0]]
	if ok && t != 0 {
		sleep := t * time.Millisecond
		log.WithFields(log.Fields{
			"this":  this,
			"ctx":   ctx,
			"sleep": sleep,
		}).Info("InjuredHook PreRead: sleeping")
		time.Sleep(sleep)
	}

	return nil, false, ctx, nil
}

// implements hookfs.HookOnRead
func (this *InjuredHook) PostRead(realRetCode int32, realBuf []byte, ctx hookfs.HookContext) ([]byte, bool, error) {
	return nil, false, nil
}

// implements hookfs.HookOnWrite
func (this *InjuredHook) PreWrite(path string, buf []byte, offset int64) (hooked bool, ctx hookfs.HookContext, err error) {
	return false, nil, nil
}

// implements hookfs.HookOnWrite
func (this *InjuredHook) PostWrite(realRetCode int32, prehookCtx hookfs.HookContext) (hooked bool, err error) {
	return false, nil
}

// implements hookfs.HookOnMkdir
func (this *InjuredHook) PreMkdir(path string, mode uint32) (bool, hookfs.HookContext, error) {
	ctx := InjuredHookContext{path: path}
	return false, ctx, nil
}

// implements hookfs.HookOnMkdir
func (this *InjuredHook) PostMkdir(realRetCode int32, ctx hookfs.HookContext) (bool, error) {
	return false, nil
}

// implements hookfs.HookOnRmdir
func (this *InjuredHook) PreRmdir(path string) (bool, hookfs.HookContext, error) {
	ctx := InjuredHookContext{path: path}
	return false, ctx, nil
}

// implements hookfs.HookOnRmdir
func (this *InjuredHook) PostRmdir(realRetCode int32, ctx hookfs.HookContext) (bool, error) {
	return false, nil
}

// implements hookfs.HookOnOpenDir
func (this *InjuredHook) PreOpenDir(path string) (bool, hookfs.HookContext, error) {
	ctx := InjuredHookContext{path: path}
	return false, ctx, nil
}

// implements hookfs.HookOnOpenDir
func (this *InjuredHook) PostOpenDir(realRetCode int32, ctx hookfs.HookContext) (bool, error) {
	return false, nil
}

// implements hookfs.HookOnFsync
func (this *InjuredHook) PreFsync(path string, flags uint32) (bool, hookfs.HookContext, error) {
	ctx := InjuredHookContext{path: path}

	this.fl.RLock()
	defer this.fl.RUnlock()

	t, ok := this.fsync[strings.Split(path, "/")[0]]
	if ok && t != 0 && path != "" {
		sleep := t * time.Millisecond
		log.WithFields(log.Fields{
			"this":  this,
			"ctx":   ctx,
			"sleep": sleep,
		}).Info("InjuredHook PreFsync: sleeping")
		// time.Sleep(sleep)
	}

	return true, ctx, syscall.ENOMEM
}

// implements hookfs.HookOnFsync
func (this *InjuredHook) PostFsync(realRetCode int32, ctx hookfs.HookContext) (bool, error) {
	return false, nil
}

func (this *InjuredHook) SetReadLatency(path string, latency time.Duration) {
	this.rl.Lock()
	defer this.rl.Unlock()

	this.read[path] = latency
	log.WithFields(log.Fields{
		"this":    this,
		"path":    path,
		"latency": latency,
	}).Info("SetReadLatency")
}

func (this *InjuredHook) SetFsyncLatency(path string, latency time.Duration) {
	this.fl.Lock()
	defer this.fl.Unlock()

	this.fsync[path] = latency
	log.WithFields(log.Fields{
		"this":    this,
		"path":    path,
		"latency": latency,
	}).Info("SetFsyncLatency")
}
