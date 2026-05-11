package main

/*
#include <jni.h>
*/
import "C"
import (
	"unsafe"

	"github.com/timob/jnigi"
)

// Wrapper around some JNI stuff to make making objects easier

type JavaObject struct {
	className  string
	objPointer *jnigi.ObjectRef
	env        *C.JNIEnv
}

func NewJavaObjectNoConstructor(env *C.JNIEnv, className string) *JavaObject {
	// normalize . to /
	className = normalizeClassName(className)

	// attempt to create object and get java pointer
	var jniEnv = jnigi.WrapEnv(unsafe.Pointer(env))
	createdObj, _ := jniEnv.NewObject(
		className,
	)

	return &JavaObject{className: className, objPointer: createdObj, env: env}
}

func (jo *JavaObject) CallMethod(methodName string, returnType string, args ...interface{}) (interface{}, error) {
	var jniEnv = jnigi.WrapEnv(unsafe.Pointer(jo.env))
	returnType = normalizeClassName(returnType)

	// is return type void? then  nil is returned
	var isVoidReturn = returnType == "void"

	// convert args to jnigi format
	var jniArgs []interface{}
	for _, arg := range args {
		switch v := arg.(type) {
		case string:
			jniArg, err := JavaString(jniEnv, v)
			if err != nil {
				return nil, err
			}
			jniArgs = append(jniArgs, jniArg)
		case int:
			jniArgs = append(jniArgs, int32(v))
		case int64:
			jniArgs = append(jniArgs, v)
		case bool:
			jniArgs = append(jniArgs, v)
		default:
			jniArgs = append(jniArgs, arg)
		}
	}

	var aRt interface{}
	if isVoidReturn {
		aRt = nil
	} else {
		aRt = returnType
	}
	result := jo.objPointer.CallMethod(jniEnv, methodName, aRt, jniArgs...)
	return result, nil
}

func (jo *JavaObject) JObject() uintptr {
	return uintptr(jo.objPointer.JObject())
}

func normalizeClassName(className string) string {
	normalized := ""
	for _, c := range className {
		if c == '.' {
			normalized += "/"
		} else {
			normalized += string(c)
		}
	}
	return normalized
}

func JavaString(env *jnigi.Env, value string) (*jnigi.ObjectRef, error) {
	bytes := []byte(value)
	charset := env.GetUTF8String()
	env.PrecalculateSignature("([BLjava/lang/String;)V")
	return env.NewObject("java/lang/String", bytes, charset)
}
