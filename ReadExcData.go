package ReadExcData

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
)

func ReadExcDataByColumn(filename string, sheet string, Column int) (data []string) {
	/*把excel文档中我们需要的列的数据全部读出来，包括第一行（标题行），返回一个String数组
	输入的文件名称需要包含完整路径，输入的表格名称也必须正确否则报错，列数从1开始。
	*/
	if filename == "" || sheet == "" || Column < 0 {
		fmt.Print("You need to give enough and right Param!")
		return
	} else {
		xlsx, err := excelize.OpenFile(filename)
		if err != nil {
			fmt.Println("路径出错？你得给一个正确路径")
			fmt.Println(err)
			return
		}
		allvalues := xlsx.GetRows(sheet)
		for _, row := range allvalues {
			data = append(data, row[Column-1])
		}
		return data
	}
}
