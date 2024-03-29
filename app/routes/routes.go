// GENERATED CODE - DO NOT EDIT
package routes

import "github.com/revel/revel"


type tApp struct {}
var App tApp


func (_ tApp) Before(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("App.Before", args).Url
}


type tStatic struct {}
var Static tStatic


func (_ tStatic) Serve(
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.Serve", args).Url
}

func (_ tStatic) ServeModule(
		moduleName string,
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "moduleName", moduleName)
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.ServeModule", args).Url
}


type tTestRunner struct {}
var TestRunner tTestRunner


func (_ tTestRunner) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("TestRunner.Index", args).Url
}

func (_ tTestRunner) Run(
		suite string,
		test string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "suite", suite)
	revel.Unbind(args, "test", test)
	return revel.MainRouter.Reverse("TestRunner.Run", args).Url
}

func (_ tTestRunner) List(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("TestRunner.List", args).Url
}


type tMain struct {}
var Main tMain


func (_ tMain) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Main.Index", args).Url
}

func (_ tMain) Upload(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Main.Upload", args).Url
}

func (_ tMain) HandleUpload(
		avatar []byte,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "avatar", avatar)
	return revel.MainRouter.Reverse("Main.HandleUpload", args).Url
}


