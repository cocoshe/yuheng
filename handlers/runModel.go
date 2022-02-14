package handlers

//
//import (
//	"github.com/gin-gonic/gin"
//)
//
//func RunHandler(c *gin.Context) {
//	// 通过exec.Command函数执行命令或者shell
//	// 第一个参数是命令路径，当然如果PATH路径可以搜索到命令，可以不用输入完整的路径
//	// 第二到第N个参数是命令的参数
//	// 下面语句等价于执行命令: ls -l
//	//cmd := exec.Command("conda", "activate", "pytorch")
//	//cmd := exec.Command("cd", "D:\\PythonProject\\Time_series_anomaly_detection")
//	//var out bytes.Buffer
//	//cmd := exec.Command("sh", "cmd.sh")
//	//cmd := exec.Command("sh", "cd D:\\PythonProject\\Time_series_anomaly_detection")
//	//cmd.Run()
//	//_ = os.Chdir("D:\\PythonProject\\Time_series_anomaly_detection") // 改变当前工作目录
//
//	//cmd := exec.Command("python", "train.py")
//	////cmd := exec.Command("python", "hello-world.py")
//	//cmd.Dir = "D:\\PythonProject\\Time_series_anomaly_detection"
//	//log.Println("cmd.Dir", cmd.Dir)
//	//
//	//workDir, _ := os.Getwd()
//	//workDir2, _ := os.Getwd()
//	//log.Println("当前工作目录:", workDir, "当前工作目录2:", workDir2)
//	//cmd.Run()
//	//cmd.Path = "/d/PythonProject/Time_series_anomaly_detection"
//	//cmd.Path = "/d/PythonProject/Time_series_anomaly_detection/train.py"
//	//cmd := exec.Command("python", "D:\\PythonProject\\Time_series_anomaly_detection\\train.py")
//	//cmd := exec.Command("python", "../PythonProject/Time_series_anomaly_detection/train.py")
//	//var stderr bytes.Buffer
//	////cmd.Stdout = &out
//	//cmd.Stderr = &stderr
//	//output, err := cmd.Output()
//	//if err != nil {
//	//	log.Println("---------------------------------")
//	//	log.Println(err)
//	//}
//	//fmt.Println(string(output))
//	//
//	//os.Chdir(workDir) // 改变当前工作目录
//
//	//cmd = exec.Command("python", "D:\\PythonProject\\Time_series_anomaly_detection\\train.py")
//	// 执行命令，并返回结果
//	//output, err = cmd.Output()
//	//if err != nil {
//	//	log.Println("---------------------------------")
//	//	log.Println(err)
//	//}
//	//fmt.Println(string(output))
//
//}
