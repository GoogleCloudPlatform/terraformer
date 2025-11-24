package myrasec

import (
	"fmt"
	"reflect"
	"strconv"
	"sync"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	mgo "github.com/Myra-Security-GmbH/myrasec-go/v2"
)

// TagSettingGenerator
type TagSettingGenerator struct {
	MyrasecService
}

// createTagSettingResources
func (g *TagSettingGenerator) createTagSettingResources(api *mgo.API, tag mgo.Tag, wg *sync.WaitGroup) error {
	defer wg.Done()

	response, err := api.ListTagSettingsMap(tag.ID)
	if err != nil {
		return err
	}

	r := terraformutils.NewResource(
		strconv.Itoa(tag.ID),
		fmt.Sprintf("%s_%d", tag.Name, tag.ID),
		"myrasec_tag_settings",
		"myrasec",
		map[string]string{
			"tag_id": strconv.Itoa(tag.ID),
		},
		[]string{},
		map[string]any{},
	)

	attributes := structToMap(mgo.Settings{})
	data := *(response.(*map[string]any))
	settings := data["settings"].(map[string]any)
	for k := range attributes {
		if _, ok := settings[k]; !ok {
			r.IgnoreKeys = append(r.IgnoreKeys, k)
		}
	}
	g.Resources = append(g.Resources, r)
	return nil
}

// InitResources
func (g *TagSettingGenerator) InitResources() error {
	wg := sync.WaitGroup{}

	api, err := g.initializeAPI()
	if err != nil {
		return err
	}

	funcs := []func(*mgo.API, mgo.Tag, *sync.WaitGroup) error{
		g.createTagSettingResources,
	}
	err = createResourcesPerTag(api, funcs, &wg, "CONFIG")
	if err != nil {
		return err
	}

	wg.Wait()

	return nil
}

func index(s, substr string) int {
	for i := range s {
		if s[i:i+len(substr)] == substr {
			return i
		}
	}
	return -1
}

func structToMap(s any) map[string]any {
	result := make(map[string]any)
	val := reflect.ValueOf(s)
	typ := reflect.TypeOf(s)

	// Make sure it's a struct
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
		typ = typ.Elem()
	}
	if val.Kind() != reflect.Struct {
		return nil
	}

	for i := 0; i < val.NumField(); i++ {
		field := typ.Field(i)
		value := val.Field(i)

		// get JSON tag
		jsonTag := field.Tag.Get("json")

		// if no json tag or "-", skip or fallback to field name
		if jsonTag == "" || jsonTag == "-" {
			jsonTag = field.Name
		} else {
			if commaIdx := reflect.ValueOf(jsonTag).String(); commaIdx != "" {
				if idx := index(jsonTag, ","); idx != -1 {
					jsonTag = jsonTag[:idx]
				}
			}
		}
		result[jsonTag] = value.Interface()
	}
	return result
}
