package main

import (
	"encoding/json"
	"log"
	"reflect"
	"strconv"
	"strings"
)

type (
	FieldInfo struct {
		Type     string     `json:"type"`
		Tags     TagsInfo   `json:"tags,omitempty"`
		Embedded FieldsInfo `json:"embedded,omitempty"`
	}

	FieldsInfo map[string]FieldInfo

	TagsInfo map[string][]string
)

func (f FieldsInfo) String() string {
	bz, err := json.MarshalIndent(f, "", "   ")
	if err != nil {
		log.Print(err)
	}

	return string(bz)
}

func GetStructTags(obj any) FieldsInfo {
	fieldsInfo := make(FieldsInfo)

	var objType reflect.Type
	if t, ok := obj.(reflect.Type); ok {
		objType = t
	} else {
		objType = reflect.ValueOf(obj).Type()
	}

	if objType.Kind() == reflect.Ptr {
		objType = objType.Elem()
	}

	if objType.Kind() != reflect.Struct {
		return nil
	}

	for idx := range objType.NumField() {
		field := objType.Field(idx)
		fieldsInfo[field.Name] = FieldInfo{
			Type:     field.Type.String(),
			Tags:     parseTagString(string(field.Tag)),
			Embedded: GetStructTags(field.Type),
		}
	}

	return fieldsInfo
}

func parseTagString(tagRaw string) TagsInfo {
	tagsInfo := make(TagsInfo)

	for _, tag := range strings.Split(tagRaw, " ") {
		if tag = strings.TrimSpace(tag); tag == "" {
			continue
		}

		tagParts := strings.Split(tag, ":")
		if len(tagParts) != 2 { //nolint:mnd // example
			continue
		}

		tagName := strings.TrimSpace(tagParts[0])
		if _, ok := tagsInfo[tagName]; ok {
			continue
		}

		var tagValues []string
		tagValuesRaw, _ := strconv.Unquote(tagParts[1])

		for _, value := range strings.Split(tagValuesRaw, ",") {
			if trimValue := strings.TrimSpace(value); trimValue != "" {
				tagValues = append(tagValues, trimValue)
			}
		}

		tagsInfo[tagName] = tagValues
	}

	return tagsInfo
}

type (
	TestStruct struct {
		ID        string `json:"id"        example:"68b69bd2-8db6-4b7f-b7f0-7c78739046c6" format:"uuid"`
		Name      string `json:"name"      example:"Bob"`
		Group     Group  `json:"group"`
		CreatedAt int64  `json:"createdAt" example:"1622647813"                           format:"unix"`
	}

	Group struct {
		ID             uint64   `json:"id"`
		PermsOverrides []string `json:"overrides" example:"USERS_RW,COMPANY_RWC"`
	}
)

func main() {
	var s *TestStruct
	log.Println(GetStructTags(s))
}
