package main

import (
	"github.com/osrg/hookfs/hookfs"
	log "github.com/sirupsen/logrus"
	"math/rand"
	"regexp"
	"strings"
	"sync"
	"time"
)

var (
	faultMap map[string]*faultContext
	fml sync.Mutex
)

func init() {
	faultMap = make(map[string]*faultContext)
}

type faultContext struct {
	errno error
	random bool
	pct int
	path string
	delay time.Duration
}

func randomErrno() error {
	return nil
}

func probab(percentage int) bool {
	return rand.Intn(99) < percentage
}

func faultInject(path, method string) error {
	fml.Lock()
	fc, ok := faultMap[method]
	if !ok {
		fml.Unlock()
		return nil
	}
	fml.Unlock()

	if !probab(fc.pct) {
		return nil
	}

	if len(fc.path) > 0 {
		re, err := regexp.Compile(fc.path)
		if err != nil || !re.MatchString(path) {
			return nil
		}
	}

	var errno error = nil
	if fc.errno != nil {
		errno = fc.errno
	} else if fc.random {
		errno = randomErrno()
	}

	if fc.delay > 0 {
		time.Sleep(fc.delay)
	}

	return errno
}

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

	t, ok := this.fsync[strings.Split(path, "/")[0]]
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
