package main

import (
  "github.com/stretchr/testify/assert"
  "testing"
  "gopkg.in/h2non/gock.v1"
  "eventos"
  "log"
  "io/ioutil"
)

func TestObtemDiferenteDe200(t *testing.T) {
  log.SetOutput(ioutil.Discard)
  defer gock.Off()

  errorMessage := []byte(`{"error": {"message": "You must use https:// when passing an access token type:OAuthException code:1 fbtrace_id:H6iEbFKqZ4T"}}`)
  gock.New("http://graph.facebook.com").
    Get("/search").
    Reply(400).
    JSON(errorMessage)

  resultado, erro := eventos.ObtemEventos()
  if resultado != nil {
		t.Fatalf("Houve um resultado válido: %v", resultado)
	}
  assert.Equal(t, "Status Code 400 - You must use https:// when passing an access token type:OAuthException code:1 fbtrace_id:H6iEbFKqZ4T", erro)
}

func TestObtem200(t *testing.T) {
  defer gock.Off()

  dataEvento := []byte(`{"data": [{"description": "algo", "name": "algum nome"}]}`)
  gock.New("http://graph.facebook.com").
    Get("/search").
    Reply(200).
    JSON(dataEvento)

  resultado, erro := eventos.ObtemEventos()

  if erro != nil {
		t.Fatalf("Houve um erro no teste: %v", erro)
	}

  var expectedEventos []eventos.Evento
  expectedEvento := eventos.Evento{Description: "algo", Name: "algum nome"}
  expectedEventos = append(expectedEventos, expectedEvento)

  assert.Equal(t, expectedEventos, resultado)
}
