package gonet

import (
	"os"
)

func WriteAllText(path string, content string) (err Error) {
	
	//os.O_WRONLY | os.O_CREATE
	file, err := os.OpenFile(path, os.O_WRONLY |os.O_TRUNC | os.O_CREATE, 0666)
	if err!=nil{		
		return err
	}
	defer file.Close() 
	 
 
    //使用缓存方式写入
    writer := bufio.NewWriter(file)
 
    count, w_err := writer.WriteString(content)
 
	//将缓存中数据刷新到文本中
    writer.Flush()

	return w_err


}



func ReadAllText(path string) (string, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
