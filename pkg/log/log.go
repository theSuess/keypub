/*
 Taken from kubernetes-sigs/controller-runtime
*/
package log

import (
	"github.com/go-logr/logr"
)

// SetLogger sets a concrete logging implementation for all deferred Loggers.
func SetLogger(l logr.Logger) {
	Log.Fulfill(l)
}

// Log is the base logger used by kubebuilder.  It delegates
// to another logr.Logger.  You *must* call SetLogger to
// get any actual logging.
var Log = NewDelegatingLogger(NullLogger{})

type GormLogr struct {
	logr.Logger
}

func WrapGorm(l logr.Logger) *GormLogr {
	return &GormLogr{l}
}

func (l *GormLogr) Print(v ...interface{}) {
	l.V(1).Info(v[3].(string), "caller", v[1], "parameters", v[4])
}
