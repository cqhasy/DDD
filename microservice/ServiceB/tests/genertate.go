package tests

import (
	"fmt"
	"reflect"
	"strings"
	"text/template"
)

const testTemplate = `
package {{.PackageName}}_test

import (
	"testing"
	"{{.PackagePath}}"
	"github.com/stretchr/testify/assert"
)

func Test{{.MethodName}}(t *testing.T) {
	calculator := &{{.StructName}}{}

	// 预设的输入数据
	input := []interface{}{ {{.Input}} }

	// 调用方法
	result := calculator.{{.MethodName}}({{.MethodArgs}})

	// 预期的结果
	expected := {{.Expected}}

	// 使用 assert 断言结果
	assert.Equal(t, expected, result)
}
`

func generateTestCode(structType interface{}) string {
	typ := reflect.TypeOf(structType)
	if typ.Kind() != reflect.Ptr {
		panic("Input must be a pointer to a struct")
	}

	var tests []string
	fmt.Println("structType:", typ)
	fmt.Println("struct_num:", typ.NumMethod())
	// 遍历结构体的所有方法
	for i := 0; i < typ.NumMethod(); i++ {
		method := typ.Method(i)

		// 提取方法的输入参数和返回值
		var methodArgs []string
		var inputValues []string
		var expectedValue string
		for j := 0; j < method.Type.NumIn(); j++ {
			methodArgs = append(methodArgs, fmt.Sprintf("input%d", j+1))
			inputValues = append(inputValues, fmt.Sprintf("input%d", j+1))
		}

		// 使用 strings.Join 将 methodArgs 转换为一个逗号分隔的字符串
		methodArgsStr := strings.Join(methodArgs, ", ")

		// 使用 strings.Join 将 inputValues 转换为一个单独的字符串
		inputValuesStr := strings.Join(inputValues, ", ")

		// 模拟生成一个简单的预期值
		expectedValue = fmt.Sprintf("%s{}", typ.Elem().Name())

		// 渲染测试函数
		testCode := renderTestTemplate(method.Name, inputValuesStr, expectedValue, methodArgsStr, typ.Elem().Name())
		tests = append(tests, testCode)
	}

	return strings.Join(tests, "\n\n")
}

func renderTestTemplate(methodName, inputValues, expectedValue, methodArgs, structName string) string {
	data := map[string]interface{}{
		"PackageName": "main",
		"PackagePath": "github.com/user/project", // Replace with your actual project path
		"MethodName":  methodName,
		"Input":       inputValues,
		"Expected":    expectedValue,
		"MethodArgs":  methodArgs,
		"StructName":  structName,
	}

	tmpl, err := template.New("test").Parse(testTemplate)
	if err != nil {
		panic(err)
	}

	var buf strings.Builder
	if err := tmpl.Execute(&buf, data); err != nil {
		panic(err)
	}

	return buf.String()
}
