package apiwha

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

type ApiWha struct {
	Url string
	Key string
}

type Result struct {
	Result interface{}
	Error  error
}

type CreditResponse struct {
	Credit float64 `json:"credit"`
}

type MessageResponse struct {
	ID           string `json:"id"`
	Number       string `json:"number"`
	From         string `json:"from"`
	To           string `json:"to"`
	Type         string `json:"type"`
	Text         string `json:"text"`
	CreationDate string `json:"creation_date"`
	ProcessDate  string `json:"process_date"`
}

type SendMessageResponse struct {
	Success     bool   `json:"success"`
	Description string `json:"description"`
	ResultCode  int    `json:"result_code"`
}

type ApiWhaMessagesParams struct {
	Type             string
	Number           string
	MarkAsPulled     string
	GetNotPulledOnly string
}

type ApiWhaSendMessageParams struct {
	Number string
	Text   string
}

func Init(url, key string) *ApiWha {
	return &ApiWha{url, key}
}

func (r *ApiWha) Credit() Result {
	model := &CreditResponse{}
	err := r.call(http.MethodGet, r.Url+`/get_credit.php?apikey=`+r.Key, model)
	if err != nil {
		return Result{Error: err}
	}
	return Result{Result: model}
}

func (r *ApiWha) Messages(params *ApiWhaMessagesParams) Result {
	model := &[]MessageResponse{}
	err := r.call(http.MethodGet,
		r.Url+`/get_messages.php?apikey=`+r.Key+
			`&type=`+params.Type+
			`&number=`+params.Number+
			`&markaspulled=`+params.MarkAsPulled+
			`&getnotpulledonly=`+params.GetNotPulledOnly,
		model)
	if err != nil {
		return Result{Error: err}
	}
	return Result{Result: model}
}

func (r *ApiWha) SendMessage(params *ApiWhaSendMessageParams) Result {
	model := &SendMessageResponse{}
	err := r.call(http.MethodGet,
		r.Url+`/send_message.php?apikey=`+r.Key+
			`&number=`+params.Number+
			`&text=`+urlEncoded(params.Text),
		model)
	if err != nil {
		return Result{Error: err}
	}
	return Result{Result: model}
}

func (aw *ApiWha) call(method, endpoint string, model interface{}) error {
	req, err := http.NewRequest(method, endpoint, nil)
	req.Header.Add("key", aw.Key)
	req.Header.Add("Content-Type", "application/json")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, &model)
	return err
}

func urlEncoded(str string) string {
	return url.QueryEscape(str)
}
