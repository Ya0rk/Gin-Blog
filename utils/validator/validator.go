package validator

import (
	"Gin-Blog/utils/errmsg"
	"fmt"
	"github.com/go-playground/locales/zh_Hans_CN"
	unTrans "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/translations/zh"
	"reflect"
)

/*
后端验证

	进行结构体验证时，将错误信息转换为中文
*/
func Validate(data interface{}) (string, int) {
	validate := validator.New()
	// 转化为中文
	uni := unTrans.New(zh_Hans_CN.New())
	trans, _ := uni.GetTranslator("zh_Hans_CN")

	err := zh.RegisterDefaultTranslations(validate, trans)
	if err != nil {
		fmt.Println("err: ", err)
	}

	// 将结构体字段名映射为label标签中的中文，方便阅读
	// 利用了结构体的反射机制
	validate.RegisterTagNameFunc(func(field reflect.StructField) string {
		label := field.Tag.Get("label")
		return label
	})

	err = validate.Struct(data)
	if err != nil {
		for _, v := range err.(validator.ValidationErrors) {
			return v.Translate(trans), errmsg.ERROR
		}
	}
	return "", errmsg.SUCCESS
}
