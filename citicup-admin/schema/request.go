package schema

//申请
/**

  "id": 2315,
  "name": "中国农业银行",
  "no": "K156",
  "code": "NJU2017",
  "email": "sgnb@qq.com",
  "applicant": "https://imgs/img1.png",
  "license": "https://imgs/img2.png"
*/
type Request struct {
	ID   uint   `json:"id"`
	Name string `json:"name"` //名称

}
