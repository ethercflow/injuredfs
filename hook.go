package main

import (
	"sync"
	"time"

	log "github.com/Sirupsen/logrus"
	hookfs "github.com/osrg/hookfs/hookfs"
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

	this.rl.RLock()
	defer this.rl.RUnlock()

	t, ok := this.read[path]
	if ok && t != 0 {
		sleep := t * time.Millisecond
		log.WithFields(log.Fields{
			"this":  this,
			"ctx":   ctx,
			"sleep": sleep,
		}).Info("InjuredHook PreRead: sleeping")
		time.Sleep(sleep)
	}

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

	this.fl.RLock()
	defer this.fl.RUnlock()

	t, ok := this.fsync[path]
	if ok && t != 0 && path != "" {
		sleep := t * time.Millisecond
		log.WithFields(log.Fields{
			"this":  this,
			"ctx":   ctx,
			"sleep": sleep,
		}).Info("InjuredHook PreFsync: sleeping")
		time.Sleep(sleep)
	}

	return nil, false, ctx
}

// implements hookfs.HookOnFsync
func (this *InjuredHook) PostFsync(realRetCode int32, ctx hookfs.HookContext) (error, bool) {
	return nil, false
}

func (this *InjuredHook) SetReadLatency(path string, latency time.Duration) {
	this.rl.Lock()
	defer this.rl.Unlock()

	log.WithFields(log.Fields{
		"this":    this,
		"path":    path,
		"latency": latency,
	}).Info("SetReadLatency")
	this.read[path] = latency
}

func (this *InjuredHook) SetFsyncLatency(path string, latency time.Duration) {
	this.fl.Lock()
	defer this.fl.Unlock()

	log.WithFields(log.Fields{
		"this":    this,
		"path":    path,
		"latency": latency,
	}).Info("SetFsyncLatency")
	this.fsync[path] = latency
}
