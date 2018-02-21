package eventos

import (
  "encoding/json"
  "net/http"
  "os"
  "io/ioutil"
  "log"
  "fmt"
  "bytes"
)

var FACEBOOK_ACCESS_TOKEN = os.Getenv("FACEBOOK_ACCESS_TOKEN")
var FACEBOOK_API_URL = "http://graph.facebook.com"

type Evento struct {
  Description string `json:description`
  Name string `json:name`
  StartTime string `json:start_time`
  EndTime string `json:end_time`
  Id string `json:id`
  Url string
}

type Data struct {
  Eventos []Evento `json:eventos`
}

type Erro struct {
  Message string `json:message`
}

type EventoErro struct {
  Error Erro `json:error`
}

func ObtemEventos() (interface{},interface{}){
  body, _, statusCode := buscaEventosNoFacebook()
  log.Println(fmt.Sprintf("Trata resposta do Facebook"))
  return trataRespostaFacebook(statusCode, body)
}

func buscaEventosNoFacebook() ([]byte, interface{}, int){
  url := FACEBOOK_API_URL + "/search?q=ciberativismo&type=event&access_token=" + FACEBOOK_ACCESS_TOKEN
  resp, err := http.Get(url)
  log.Println(fmt.Sprintf("Status Code: %v - Url: %v", resp.StatusCode, url))
  if err != nil {
	   log.Println(fmt.Sprintf("Houve um erro: %v", err))
     return nil, err, resp.StatusCode
  }
  defer resp.Body.Close()
  body, erro := ioutil.ReadAll(resp.Body)
  return body, erro, resp.StatusCode
}

func trataRespostaFacebook(statusCode int, body []byte) (interface{}, interface{}) {
  if statusCode == 200 {
    log.Println(fmt.Sprintf("trataSucesso"))
    return trataSucesso(body)
  } else {
    var jsonResponse EventoErro
    bodyReader := bytes.NewReader(body)
    err := json.NewDecoder(bodyReader).Decode(&jsonResponse)
    log.Println(fmt.Sprintf("trataErros"))
    return trataErros(err, jsonResponse.Error.Message, statusCode)
  }
}

func trataSucesso(body []byte) (interface{}, interface{}){
  eventos := &Data{}
  bodyReader := bytes.NewReader(body)
  json.NewDecoder(bodyReader).Decode(&eventos)
  log.Println(fmt.Sprintf("Eventos: %v", eventos.Eventos))
  return eventos.Eventos, nil
}

func converteMapParaEvento(eventoMap map[string]string) (evento Evento) {
  bytes, erro := json.Marshal(eventoMap)
  erro = json.Unmarshal(bytes, &evento)
  if erro != nil {
    fmt.Println(erro)
    evento = Evento{}
  }
  return
}

func trataErros(erro interface{}, mensagemErro interface{}, statusCode int) (resultado interface{}, mensagem string){
  if erro != nil {
      mensagem = fmt.Sprintf("Erro: %v\n", erro)
  }

  if statusCode != 200 {
    mensagem = fmt.Sprintf("Status Code %v - %v", statusCode, mensagemErro)
  }

  return
}
