package util

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("util", func() {

	Describe("ReverseDomain", func() {
		It("should reverse the string", func() {
			Expect("io.java", ReverseDomain("java.io"))
			Expect("com.example", ReverseDomain("example.com"))
			Expect("no/dots/leave/alone", ReverseDomain("no/dots/leave/alone"))
		})
		It("should leave the string alone", func() {
			Expect("no/dots/leave/alone", ReverseDomain("no/dots/leave/alone"))
			Expect("random string/No|dots;at,all",
				ReverseDomain("random string/No|dots;at,all"))
		})
	})

	Describe("ToCamel", func() {
		It("should convert to Camel", func() {
			Expect("appTest", ToCamel("app_test"))
		})
		It("should convert to Camel when start with _", func() {
			Expect("AppTest", ToCamel("_app_test"))
			Expect("JavaOp", ToCamel("java-op"))
		})
		It("should convert to Camel when has numbers", func() {
			Expect("AppTestK8s", ToCamel("_app_test_k8s"))
		})
		It("should handle special words", func() {
			Expect("egressIPs", ToCamel("egressIPs"))
			Expect("myURL", ToCamel("myURL"))
			Expect("myURL", ToCamel("my_url"))
		})
	})

	Describe("ToClassname", func() {
		It("should capitalize the first letter", func() {
			Expect("AppTest", ToClassname("app_test"))
			Expect("AppTest", ToClassname("_app_test"))
			Expect("JavaOp", ToClassname("java-op"))
			Expect("AppTestK8s", ToClassname("_app_test_k8s"))
			Expect("EgressIPs", ToClassname("egressIPs"))
			Expect("MyURL", ToClassname("myURL"))
			Expect("MyURL", ToClassname("my_url"))
		})
	})
})
