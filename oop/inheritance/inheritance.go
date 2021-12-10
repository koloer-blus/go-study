package inheritance

import "fmt"

type BaseResponse struct {
	code    int
	message string
	data    interface{}
}

func (bRes *BaseResponse) updateCode(code int) {
	bRes.code = code
}

type GETResponse struct {
	BaseResponse
	status string
}

func GetResponse() {
	res := GETResponse{
		status: "success",
		BaseResponse: BaseResponse{
			code:    200,
			message: "no res",
			data: []string{
				"Jack",
				"Mark",
			},
		},
	}

	fmt.Println(res)

	res.updateCode(400)
	fmt.Println(res)
}
