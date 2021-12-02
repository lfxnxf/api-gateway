// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import (
	"bytes"
	"encoding/json"
	"strings"
	"text/template"

	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "[http://swagger.io/terms/](http://swagger.io/terms/)` + "`" + `",
        "contact": {
            "name": "xuefeng` + "`" + `",
            "url": "[http://www.swagger.io/support](http://www.swagger.io/support)` + "`" + `",
            "email": "xuefeng6329@126.com` + "`" + `"
        },
        "license": {
            "url": "[http://www.apache.org/licenses/LICENSE-2.0.html](http://www.apache.org/licenses/LICENSE-2.0.html)` + "`" + `"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/v1/vehicle/add": {
            "post": {
                "description": "新增车辆信息",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "新增车辆信息",
                "parameters": [
                    {
                        "description": "新增车辆信息参数",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.AddVehicleReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/utils.WrapResp"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/utils.WrapResp"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.AddVehicleReq": {
            "type": "object",
            "properties": {
                "driver_id": {
                    "description": "司机id",
                    "type": "integer"
                },
                "license_plate": {
                    "description": "车牌号",
                    "type": "string"
                },
                "vehicle_info_id": {
                    "description": "车辆类型Id",
                    "type": "integer"
                }
            }
        },
        "utils.WrapResp": {
            "type": "object",
            "properties": {
                "data": {},
                "dm_error": {
                    "type": "integer"
                },
                "error_msg": {
                    "type": "string"
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "1.0`",
	Host:        "47.241.77.253:10000`",
	BasePath:    "这里写/api/v1/`",
	Schemes:     []string{},
	Title:       "校车通`",
	Description: "校车通接口文档`",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
		"escape": func(v interface{}) string {
			// escape tabs
			str := strings.Replace(v.(string), "\t", "\\t", -1)
			// replace " with \", and if that results in \\", replace that with \\\"
			str = strings.Replace(str, "\"", "\\\"", -1)
			return strings.Replace(str, "\\\\\"", "\\\\\\\"", -1)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register("swagger", &s{})
}