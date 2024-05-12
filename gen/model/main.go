package main

import (
	"chaindraw-fair-ticket-backend/core"
	"chaindraw-fair-ticket-backend/global"
	"chaindraw-fair-ticket-backend/initialize"
	"gorm.io/gen"
	"gorm.io/gorm"
	"strings"
)

// 根据库表重新生成数据库model
func main() {
	core.Viper()
	global.DB = initialize.Gorm()

	g := gen.NewGenerator(gen.Config{
		OutPath: "model",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
	})
	g.UseDB(global.DB)

	var dataMap = map[string]func(gorm.ColumnType) (dataType string){
		// int mapping
		"int": func(columnType gorm.ColumnType) (dataType string) {
			if n, ok := columnType.Nullable(); ok && n {
				return "*int64"
			}
			return "int64"
		},

		// bool mapping
		"tinyint": func(columnType gorm.ColumnType) (dataType string) {
			ct, _ := columnType.ColumnType()
			if strings.HasPrefix(ct, "tinyint(1)") {
				return "*bool"
			}
			return "byte"
		},
	}
	g.WithDataTypeMap(dataMap)

	// 定义json tag命名格式
	g.GenerateAllTable(gen.FieldJSONTagWithNS(func(columnName string) (tagContent string) {
		arr := strings.Split(columnName, "_")
		tagContent = arr[0]
		for i := 1; i < len(arr); i++ {
			word := upFirst(arr[i])
			tagContent += word
		}
		return tagContent
	}))

	// Generate the code
	g.Execute()

}

func upFirst(str string) (word string) {
	byteArr := []byte(str)
	byteArr[0] -= 32
	word = string(byteArr)
	return word
}
