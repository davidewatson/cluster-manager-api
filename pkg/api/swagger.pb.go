package cluster_manager_api

const (
APISwaggerJSON = `{
  "swagger": "2.0",
  "info": {
    "title": "api.proto",
    "version": "version not set"
  },
  "schemes": [
    "http",
    "https"
  ],
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "paths": {
    "/v1/helloworld": {
      "post": {
        "operationId": "HelloWorld",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/cluster_manager_apiHelloWorldReply"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/cluster_manager_apiHelloWorldMsg"
            }
          }
        ],
        "tags": [
          "Cluster"
        ]
      }
    }
  },
  "definitions": {
    "cluster_manager_apiHelloWorldMsg": {
      "type": "object",
      "properties": {
        "name": {
          "type": "string"
        }
      },
      "title": "The Hello World request"
    },
    "cluster_manager_apiHelloWorldReply": {
      "type": "object",
      "properties": {
        "message": {
          "type": "string"
        }
      },
      "title": "The response to Hello World"
    }
  }
}
`
)
