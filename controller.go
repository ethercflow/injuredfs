package main

import (
	"math/rand"
	"regexp"
	"sync"
	"time"
)

var (
	faultMap map[string]*faultContext
	fml      sync.Mutex

	methods map[string]bool
)

func init() {
	faultMap = make(map[string]*faultContext)
	initMethods()
}

type faultContext struct {
	errno  error
	random bool
	pct    int
	path   string
	delay  time.Duration
}

func initMethods() {
	methods = make(map[string]bool)
	methods["open"] = true
	methods["read"] = true
	methods["write"] = true
	methods["mkdir"] = true
	methods["rmdir"] = true
	methods["opendir"] = true
	methods["fsync"] = true
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

func Methods() []string {
	ms := make([]string, 0)
	for k := range methods {
		ms = append(ms, k)
	}
	return ms
}

func RecoverALL() {
	fml.Lock()
	defer fml.Unlock()
	// The compiler(1.11) now optimizes map clearing operations of the form:
	for k := range faultMap {
		delete(faultMap, k)
	}
}
func RecoverMethod(method string) {
	fml.Lock()
	defer fml.Unlock()
	delete(faultMap, method)
}

func FaultInject(methods []string, errno error, random bool, pct int, path string, delay time.Duration) {
	f := &faultContext{errno: errno, random: random, pct: pct, path: path, delay: delay}

	fml.Lock()
	defer fml.Unlock()
	for _, v := range methods {
		faultMap[v] = f
	}
}
func FaultInjectAll(errno error, random bool, pct int, path string, delay time.Duration) {
	ms := Methods()
	FaultInject(ms, errno, random, pct, path, delay)
}
