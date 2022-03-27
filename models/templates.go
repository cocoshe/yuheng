package models

import "github.com/google/uuid"

type DisplayRequest struct {
	Idx1 string `json:"idx1" example:"1"`
	Idx2 string `json:"idx2" example:"3"`
}

type UserPostDto struct {
	Id   string `json:"id" example:"17280000089799"`
	Post string `json:"post,omitempty" example:"文字内容"`
	Pic  string `json:"pic" example:"图片内容"`
}

type CompanyAppealDto struct {
	Id   string `json:"id" example:"17280000089799"`
	Post string `json:"post" example:"文字申诉内容"`
	Pic  string `json:"pic,omitempty" example:"图片内容"`
}

type AccusStatusDto struct {
	Id     uuid.UUID `json:"id" example:"9cade01e-4c66-4c57-910a-2efe5f7e3dff"`
	Status int       `json:"status" example:"1"`
}

type AplStatusDto struct {
	Id     uuid.UUID `json:"id" example:"967d7ae4-1c39-40f8-befd-222d768d70f7"`
	Status int       `json:"status" example:"1"`
}

type SuccessResponse struct {
	Code string `json:"code" example:"200"`
	Msg  string `json:"msg" example:"success"`
}

type AccusationList struct {
	Code string       `json:"code" example:"200"`
	Msg  string       `json:"msg" example:"success"`
	Data []Accusation `json:"data"`
}

type AppealList struct {
	Code string   `json:"code" example:"200"`
	Msg  string   `json:"msg" example:"success"`
	Data []Appeal `json:"data"`
}

type FailureResponse struct {
	Code string `json:"code" example:"401"`
	Msg  string `json:"msg" example:"token invalid"`
}

type LoginRequest struct {
	Id  string `json:"id" gorm:"primary_key" example:"321321"` // 账号
	Pwd string `json:"pwd" example:"321"`                      // 密码
}

type LoginResponse struct {
	Code  string `json:"code" example:"200" example:"200"`
	Msg   string `json:"msg" example:"success" example:"success"`
	Token string `json:"token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjE3MjgwMDAwMDg5OSIsImlhdCI6MTU4NjQwMjQ5NywiZXhwIjoxNTg2NDAyNzk3fQ."`
}

type DisplayResponse struct {
	Code string    `json:"code" example:"200"`
	Msg  string    `json:"msg" example:"success"`
	Data []Company `json:"data"`
}

type RankResponse struct {
	Code string `json:"code" example:"200"`
	Msg  string `json:"msg" example:"success"`
	Data []User `json:"data"`
}

type JSONid struct {
	Id string `json:"id" example:"17280000089799"`
}

type CpnListResponse struct {
	Code string    `json:"code" example:"200"`
	Msg  string    `json:"msg" example:"success"`
	Data []Company `json:"data"`
}

// todo: 完成另一个时间段的采取, 跑两个维度的数据
type RunModelRequest struct {
	PortId                     string  `json:"port_id" example:"64000000000600100000"`        // 污染口ID
	PolutionId                 string  `json:"polution_id" example:"w00000"`                  // 污染物ID
	CompanyId                  string  `json:"company_id" example:"17280000089583"`           // 公司ID
	DataS1                     string  `json:"date_s_1" example:"1/1/2020"`                   // 日/月/年,当前开始时间
	DataE1                     string  `json:"date_e_1" example:"1/2/2020"`                   // 日/月/年,当前结束时间
	DataS2                     string  `json:"date_s_2" example:"1/3/2020"`                   // 日/月/年,对照数据开始时间
	DataE2                     string  `json:"date_e_2" example:"1/4/2020"`                   // 日/月/年,对照数据结束时间
	ThresholdConcentrationLoss float64 `json:"threshold_concentration_loss" example:"200.10"` // 浓度损失阈值
	ThresholdAmountLoss        float64 `json:"threshold_amount_loss" example:"2000.10"`       // 排放量损失阈值
}

type RunModelResponse struct {
	Code string           `json:"code" example:"200"`
	Msg  string           `json:"msg" example:"success"`
	Data RunModelRespData `json:"data"`
}

type RunModelRespData struct {
	Code     string        `json:"code" example:"200"`
	Meta     string        `json:"meta" example:"single"` // single为单个污染物模式, all为总览模式
	Msg      string        `json:"msg" example:"success"`
	Data     DataModel     `json:"data"`     // 单个污染物视图下的相关数据
	Overview OverviewModel `json:"overview"` // 总览视图下的相关数据
}

type DataModel struct {
	AverageAmount                                     float64   `json:"average_amount" example:"2127.598713285988"`                                                    // 平均排污量
	AverageConcentration                              float64   `json:"average_concentration" example:"52.787238479217876"`                                            // 平均浓度
	Data1CompareData2WithAverageAmountLoss            float64   `json:"data1_compare_data2_with_average_amount_loss" example:"51.1888359508704"`                       // 较[上周]排污量增长量
	Data1CompareData2WithAverageAmountLossRate        float64   `json:"data1_compare_data2_with_average_amount_loss_rate" example:"-0.8009344011185189"`               // 较[上周]排污量增长率
	Data1CompareData2WithAverageConcentrationLoss     float64   `json:"data1_compare_data2_with_average_concentration_loss" example:"-13.094681522291124"`             // 较[上周]浓度增长量
	Data1CompareData2WithAverageConcentrationLossRate float64   `json:"data1_compare_data2_with_average_concentration_loss_rate" example:"-0.8279401442320188"`        // 较[上周]浓度增长率
	Date1AmountTestList                               []float64 `json:"date1_amount_test_list" example:"1853.283905720711,1853.283905720711,314.8520582050088"`        // 模型重构排放量序列
	Date1AmountTruthList                              []float64 `json:"date1_amount_truth_list" example:"870.0000334143641,1757.9999680727722,1917.9999588906767"`     // 原始采集排放量数据
	Date1AmountWarningTags                            []string  `json:"date1_amount_warning_tags" example:"2020-01-07', '2020-01-08'"`                                 // 排放量异常日期列表
	Date1AmountWarningTagsCount                       int       `json:"date1_amount_warning_tags_count" example:"1"`                                                   // 排放量异常日期数量
	Date1CompareAmountLossWithThreshold               float64   `json:"date1_compare_amount_loss_with_threshold" example:"-2375.351830534637"`                         // 平均排放量异常指数较阈值
	Date1CompareConcentrationLossWithThreshold        float64   `json:"date1_compare_concentration_loss_with_threshold" example:"-199.19129530233997"`                 // 平均浓度异常指数较阈值
	Date1ConcentrationTestList                        []float64 `json:"date1_concentration_test_list" example:"90.42889326810837,90.42889326810837,71.88690975308418"` // 模型重构浓度序列
	Date1ConcentrationTruthList                       []float64 `json:"date1_concentration_truth_list" example:"66.25000024214387,73.17999936640263"`                  // 原始采集浓度数据
	Date1ConcentrationWarningTags                     []string  `json:"date1_concentration_warning_tags" example:"'2020-01-07', '2020-01-08'"`                         // 浓度异常日期列表
	Date1ConcentrationWarningTagsCount                int       `json:"date1_concentration_warning_tags_count" example:"1"`                                            // 浓度异常日期数量
	Date2AmountTruthList                              []float64 `json:"date2_amount_truth_list" example:"870.0000334143641,1757.9999680727722,1917.9999588906767"`     // [上周]原始采集排放量数据
	Date2AverageAmountTruth                           float64   `json:"date2_average_amount_truth" example:"2127.598713285988"`                                        // [上周]平均排污量
	Date2AverageConcentrationTruth                    float64   `json:"date2_average_concentration_truth" example:"52.787238479217876"`                                // [上周]平均浓度
	Date2ConcentrationTruthList                       []float64 `json:"date2_concentration_truth_list" example:"66.25000024214387,73.17999936640263"`                  // [上周]原始采集浓度数据
	PolutionId                                        string    `json:"polution_id" example:"w00000"`                                                                  // 污染物ID
	PortAverageAmount                                 float64   `json:"port_average_amount" example:"2127.598713285988"`                                               // 排放口平均排污量
	PortTotalConcentration                            float64   `json:"port_total_concentration" example:"61295.512913"`                                               // 排放口总浓度
	RelativeCpn                                       []string  `json:"relative_cpn" example:"123,456,789"`                                                            // 排放口关联公司
}

type OverviewModel struct {
	CompareTotalAmountWarningCount            int     `json:"compare_total_amount_warning_count" example:"1"`                               // 较[上周] 总览排放量异常次数增加
	CompareTotalAmountWarningCountRate        float64 `json:"compare_total_amount_warning_count_rate" example:"-0.8279401442320188"`        // 较[上周] 总览排放量异常次数增加率, 如果为-1,则表明[上周]排放量没有异常
	CompareTotalConcentrationWarningCount     int     `json:"compare_total_concentration_warning_count" example:"1"`                        // 较[上周] 总览浓度异常次数增加
	CompareTotalConcentrationWarningCountRate float64 `json:"compare_total_concentration_warning_count_rate" example:"-0.8279401442320188"` // 较[上周] 总览浓度异常次数增加率, 如果为-1,则表明[上周]浓度没有异常
	CompareTotalWarningCount                  int     `json:"compare_total_warning_count" example:"1"`                                      // 较[上周] 总览排放量和浓度异常次数增加
	CompareTotalWarningCountRate              float64 `json:"compare_total_warning_count_rate" example:"-0.8279401442320188"`               // 较[上周] 总览排放量和浓度异常次数增加率
	Date1TotalAmountWarningCount              int     `json:"date1_total_amount_warning_count" example:"1"`                                 // [本周] 总览排放量异常次数
	Date1TotalConcentrationWarningCount       int     `json:"date1_total_concentration_warning_count" example:"1"`                          // [本周] 总览浓度异常次数
	Date1TotalWarningCount                    int     `json:"date1_total_warning_count" example:"1"`                                        // [本周] 总览浓度与排放量异常总次数
	Date2TotalAmountWarningCount              int     `json:"date2_total_amount_warning_count" example:"1"`                                 // [上周] 总览排放量异常次数
	Date2TotalConcentrationWarningCount       int     `json:"date2_total_concentration_warning_count" example:"1"`                          // [上周] 总览浓度异常次数
	Date2TotalWarningCount                    int     `json:"date2_total_warning_count" example:"1"`                                        // [上周] 总览浓度与排放量异常总次数
}

type DateList struct {
	Date1 []string `json:"date1" example:"2019-01-01,2019-01-02,2019-01-03,2019-01-04,2019-01-05,2019-01-06,2019-01-07"` // 当前检测时段的日期列表
	Date2 []string `json:"date2" example:"2019-01-01,2019-01-02,2019-01-03,2019-01-04,2019-01-05,2019-01-06,2019-01-07"` // [上周]检测时段的日期列表
}

type InfoReq struct {
	CompanyId string `json:"company_id" example:"17280000089583"` // 公司ID
}

type InfoResp struct {
	PortIds     []string `json:"port_ids" example:"64000000000600100000,64000000000600100000,64000000000600100000"` // 排污口ID
	PolutionIds []string `json:"polution_ids" example:"w00000,w00001,w00010"`                                       // 污染物ID
}

type InfoRespDto struct {
	Data InfoResp `json:"data"`
}

type SelfCheckReq struct {
	Data string `json:"data"`
}

type SelfCheckResp struct {
	RebuildData     []float64 `json:"rebuild_data"`      // 重构后的序列
	AbnormalDateIdx []int     `json:"abnormal_date_idx"` // 异常的日期索引
	OriginAvg       float64   `json:"origin_avg"`        // 原始序列平均值
	RebuildAvg      float64   `json:"rebuild_avg"`       // 重构后序列在平均值

}

type ChangeThresholdReq struct {
	Feature   string  ` json:"feature" gorm:"column:feature"`
	Threshold float64 `json:"threshold" gorm:"column:threshold"`
}

type GetFeaturesResp struct {
	Features []string `json:"features"`
	SuccessResponse
}
