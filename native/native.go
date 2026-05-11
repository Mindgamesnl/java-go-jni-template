package main

/*
#include <jni.h>
*/
import "C"
import "time"

//export Java_dev_pixelib_gojni_embedded_EmbeddedApplication_ping
func Java_dev_pixelib_gojni_embedded_EmbeddedApplication_ping(env *C.JNIEnv, cls C.jclass) C.jboolean {
	return C.JNI_TRUE
}

//export Java_dev_pixelib_gojni_embedded_EmbeddedApplication_createExampleObject
func Java_dev_pixelib_gojni_embedded_EmbeddedApplication_createExampleObject(env *C.JNIEnv, cls C.jclass) C.jobject {

	var javaObject = NewJavaObjectNoConstructor(env, "dev.pixelib.gojni.example.HelloWorldContainer")

	if _, err := javaObject.CallMethod("setMessage", "void", "Hello from Go!"); err != nil {
		panic(err)
	}
	if _, err := javaObject.CallMethod("setCount", "void", 42); err != nil {
		panic(err)
	}

	// current unix timestamp in as int64 in milliseconds
	var currentGoTimestamp = time.Now().UnixMilli()
	if _, err := javaObject.CallMethod("setTimestamp", "void", currentGoTimestamp); err != nil {
		panic(err)
	}

	return C.jobject(javaObject.JObject())
}

func main() {}
