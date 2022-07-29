 ```bash
 ✘ ⚡ root@backendcloud  ~/t3  go install  github.com/swaggo/swag/cmd/swag 
 ⚡ root@backendcloud  ~/t3  swag init -g main.go                                
2022/07/29 16:49:49 Generate swagger docs....
2022/07/29 16:49:49 Generate general API Info, search dir:./
2022/07/29 16:49:49 create docs.go at  docs/docs.go
2022/07/29 16:49:49 create swagger.json at  docs/swagger.json
2022/07/29 16:49:49 create swagger.yaml at  docs/swagger.yaml
 ⚡ root@backendcloud  ~/t3  ls
docs  go.mod  go.sum  main.go
 ⚡ root@backendcloud  ~/t3  tree
.
├── docs
│   ├── docs.go
│   ├── swagger.json
│   └── swagger.yaml
├── go.mod
├── go.sum
└── main.go

1 directory, 6 files
 ⚡ root@backendcloud  ~/t3  cd docs 
 ⚡ root@backendcloud  ~/t3/docs  ls
docs.go  swagger.json  swagger.yaml
 ⚡ root@backendcloud  ~/t3/docs  more *
::::::::::::::
docs.go
::::::::::::::
// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "[http://swagger.io/terms/](http://swagger.io/terms/)` + "`" + `",
        "contact": {
            "name": "这里写联系人信息` + "`" + `",
            "url": "[http://www.swagger.io/support](http://www.swagger.io/support)` + "`" + `",
            "email": "support@swagger.io` + "`" + `"
        },
        "license": {
            "name": "Apache 2.0` + "`" + `",
            "url": "[http://www.apache.org/licenses/LICENSE-2.0.html](http://www.apache.org/licenses/LICENSE-2.0.html)` + "`" + `"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {}
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
        Version:          "1.0`",
        Host:             "这里写接口服务的host`",
        BasePath:         "这里写base path`",
        Schemes:          []string{},
        Title:            "这里写标题`",
        Description:      "这里写描述信息`",
        InfoInstanceName: "swagger",
        SwaggerTemplate:  docTemplate,
}

func init() {
        swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
::::::::::::::
swagger.json
::::::::::::::
{
    "swagger": "2.0",
    "info": {
        "description": "这里写描述信息`",
        "title": "这里写标题`",
        "termsOfService": "[http://swagger.io/terms/](http://swagger.io/terms/)`",
        "contact": {
            "name": "这里写联系人信息`",
            "url": "[http://www.swagger.io/support](http://www.swagger.io/support)`",
            "email": "support@swagger.io`"
        },
        "license": {
            "name": "Apache 2.0`",
            "url": "[http://www.apache.org/licenses/LICENSE-2.0.html](http://www.apache.org/licenses/LICENSE-2.0.html)`"
        },
        "version": "1.0`"
    },
    "host": "这里写接口服务的host`",
    "basePath": "这里写base path`",
    "paths": {}
}
::::::::::::::
swagger.yaml
::::::::::::::
basePath: 这里写base path`
host: 这里写接口服务的host`
info:
  contact:
    email: support@swagger.io`
    name: 这里写联系人信息`
    url: '[http://www.swagger.io/support](http://www.swagger.io/support)`'
  description: 这里写描述信息`
  license:
    name: Apache 2.0`
    url: '[http://www.apache.org/licenses/LICENSE-2.0.html](http://www.apache.org/licenses/LICENSE-2.0.html)`'
  termsOfService: '[http://swagger.io/terms/](http://swagger.io/terms/)`'
  title: 这里写标题`
  version: 1.0`
paths: {}
swagger: "2.0"
```