package controllers

import (
	"fcm_sender/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"reflect"
	"strings"
)

// минимальная инфа
func About(c *gin.Context) {
	fields := getFields(reflect.TypeOf(models.Sender{}))

	var routes []models.Route
	routes = append(routes, models.Route{
		Method: "GET",
		URL:    "/heartbeat",
		Fields: nil,
	})
	routes = append(routes, models.Route{
		Method: "POST",
		URL:    "/send",
		Fields: fields,
	})

	c.JSON(http.StatusOK, routes)
}

// для данных по полям getFields(reflect.TypeOf(models.Sender{}))
func getFields(t reflect.Type) []models.Field {
	value := reflect.New(t).Interface()
	v := reflect.ValueOf(value)
	i := reflect.Indirect(v)
	s := i.Type()
	var fields []models.Field
	for i := 0; i < s.NumField(); i++ {
		t := s.Field(i)
		fieldName := t.Name
		fieldReq := false

		if jsonTag := t.Tag.Get("json"); jsonTag != "" && jsonTag != "-" {
			var commaIdx int
			if commaIdx = strings.Index(jsonTag, ","); commaIdx < 0 {
				commaIdx = len(jsonTag)
			}
			fieldName = jsonTag[:commaIdx]
		}
		if jsonReq := t.Tag.Get("binding"); jsonReq != "" && jsonReq != "-" {
			var commaIdx int
			if commaIdx = strings.Index(jsonReq, ","); commaIdx < 0 {
				commaIdx = len(jsonReq)
			}
			fieldReq = "required" == jsonReq[:commaIdx]
		}

		fields = append(fields, models.Field{
			Name:     fieldName,
			Type:     t.Type.String(),
			Required: fieldReq,
		})
	}
	return fields
}
