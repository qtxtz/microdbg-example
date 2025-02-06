package tiger

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"fmt"

	android "github.com/wnxd/microdbg-android"
	gava "github.com/wnxd/microdbg-android/java"
	"github.com/wnxd/microdbg-android/wrapper"
	java "github.com/wnxd/microdbg-java"
)

type Handler struct {
}

func (h Handler) DefineClass(ctx wrapper.FakeContext, name string) gava.FakeClass {
	// fmt.Println("DefineClass", name)
	var cls gava.FakeClass
	switch name {
	case "com.aliyun.TigerTally.common.utils.SecurityUtil":
		cls = ctx.DefineClass(name)
		cls.DefineMethod("getCtx", "()Landroid/content/Context;", gava.Modifier_PUBLIC).BindCall(func(obj java.IObject, args ...any) any {
			return ctx.GetClass("android.content.Context").NewInstance()
		})
	case "com.aliyun.TigerTally.s.A":
		cls = ctx.DefineClass(name)
		cls.DefineMethod("ct", "()Landroid/content/Context;", gava.Modifier_PUBLIC).BindCall(func(obj java.IObject, args ...any) any {
			return ctx.GetClass("android.content.Context").NewInstance()
		})
	case "com.aliyun.TigerTally.s.A$AA":
		cls = ctx.DefineClass(name)
		cls.DefineMethod(gava.ConstructorMethodName, "()V", gava.Modifier_PUBLIC).BindCall(func(obj java.IObject, args ...any) any {
			return cls.NewInstance()
		})
		cls.DefineMethod("en", "(Ljava/lang/String;)Ljava/lang/String;", gava.Modifier_PUBLIC).BindCall(func(obj java.IObject, args ...any) any {
			data := args[0].(java.IString).String()
			sum := md5.Sum([]byte(data))
			return gava.FakeString(hex.EncodeToString(sum[:]))
		})
	case "com.aliyun.TigerTally.s.A$BB":
		cls = ctx.DefineClass(name)
		cls.DefineMethod(gava.ConstructorMethodName, "()V", gava.Modifier_PUBLIC).BindCall(func(obj java.IObject, args ...any) any {
			return cls.NewInstance()
		})
		cls.DefineMethod("en", "(Ljava/lang/String;)Ljava/lang/String;", gava.Modifier_PUBLIC).BindCall(func(obj java.IObject, args ...any) any {
			data := args[0].(java.IString).String()
			sum := sha1.Sum([]byte(data))
			return gava.FakeString(hex.EncodeToString(sum[:]))
		})
	}
	return cls
}

func (h Handler) DefineMethod(ctx wrapper.FakeContext, clazz gava.FakeClass, name, sig string) gava.FakeMethod {
	// fmt.Println("DefineMethod", clazz.GetName(), name, sig)
	return nil
}

func (h Handler) DefineStaticMethod(ctx wrapper.FakeContext, clazz gava.FakeClass, name, sig string) gava.FakeMethod {
	// fmt.Println("DefineStaticMethod", clazz.GetName(), name, sig)
	return nil
}

func (h Handler) DefineField(ctx wrapper.FakeContext, clazz gava.FakeClass, name, sig string) gava.FakeField {
	// fmt.Println("DefineField", clazz.GetName(), name, sig)
	return nil
}

func (h Handler) DefineStaticField(ctx wrapper.FakeContext, clazz gava.FakeClass, name, sig string) gava.FakeField {
	// fmt.Println("DefineStaticField", clazz.GetName(), name, sig)
	return nil
}

func (h Handler) CallMethod(ctx android.JNIContext, obj gava.FakeObject, name, sig string, args ...any) any {
	fmt.Println("CallMethod", obj.ToString(), name, sig)
	return nil
}

func (h Handler) CallStaticMethod(ctx android.JNIContext, clazz gava.FakeClass, name, sig string, args ...any) any {
	fmt.Println("CallStaticMethod", clazz.GetName(), name, sig)
	return nil
}

func (h Handler) GetField(ctx android.JNIContext, obj gava.FakeObject, name string) any {
	fmt.Println("GetField", obj.ToString(), name)
	return nil
}

func (h Handler) SetField(ctx android.JNIContext, obj gava.FakeObject, name string, value any) {
	fmt.Println("GetField", obj.ToString(), name)
}

func (h Handler) GetStaticField(ctx android.JNIContext, clazz gava.FakeClass, name string) any {
	fmt.Println("GetStaticField", clazz.GetName(), name)
	return nil
}

func (h Handler) SetStaticField(ctx android.JNIContext, clazz gava.FakeClass, name string, value any) {
	fmt.Println("GetStaticField", clazz.GetName(), name)
}

func formatSignature(clazz gava.FakeClass, name, sig string) string {
	return clazz.GetName().String() + name + sig
}
