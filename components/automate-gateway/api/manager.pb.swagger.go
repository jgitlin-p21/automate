package api

func init() {
	Swagger.Add("manager", `{
  "swagger": "2.0",
  "info": {
    "title": "api/external/nodes/manager/manager.proto",
    "version": "version not set"
  },
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "paths": {
    "/nodemanagers": {
      "post": {
        "operationId": "Create",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.nodes.manager.v1.Ids"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/chef.automate.api.nodes.manager.v1.NodeManager"
            }
          }
        ],
        "tags": [
          "NodeManagerService"
        ]
      }
    },
    "/nodemanagers/id/{id}": {
      "get": {
        "operationId": "Read",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.nodes.manager.v1.NodeManager"
            }
          }
        },
        "parameters": [
          {
            "name": "id",
            "in": "path",
            "required": true,
            "type": "string"
          }
        ],
        "tags": [
          "NodeManagerService"
        ]
      },
      "delete": {
        "operationId": "Delete",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "properties": {}
            }
          }
        },
        "parameters": [
          {
            "name": "id",
            "in": "path",
            "required": true,
            "type": "string"
          }
        ],
        "tags": [
          "NodeManagerService"
        ]
      },
      "put": {
        "operationId": "Update",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "properties": {}
            }
          }
        },
        "parameters": [
          {
            "name": "id",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/chef.automate.api.nodes.manager.v1.NodeManager"
            }
          }
        ],
        "tags": [
          "NodeManagerService"
        ]
      }
    },
    "/nodemanagers/id/{id}/with-node-state/stopped": {
      "delete": {
        "operationId": "DeleteWithNodeStateStopped",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "properties": {}
            }
          }
        },
        "parameters": [
          {
            "name": "id",
            "in": "path",
            "required": true,
            "type": "string"
          }
        ],
        "tags": [
          "NodeManagerService"
        ]
      }
    },
    "/nodemanagers/id/{id}/with-node-state/terminated": {
      "delete": {
        "operationId": "DeleteWithNodeStateTerminated",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "properties": {}
            }
          }
        },
        "parameters": [
          {
            "name": "id",
            "in": "path",
            "required": true,
            "type": "string"
          }
        ],
        "tags": [
          "NodeManagerService"
        ]
      }
    },
    "/nodemanagers/id/{id}/with-nodes": {
      "delete": {
        "operationId": "DeleteWithNodes",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.nodes.manager.v1.Ids"
            }
          }
        },
        "parameters": [
          {
            "name": "id",
            "in": "path",
            "required": true,
            "type": "string"
          }
        ],
        "tags": [
          "NodeManagerService"
        ]
      }
    },
    "/nodemanagers/id/{node_manager_id}/search-fields": {
      "post": {
        "operationId": "SearchNodeFields",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.nodes.manager.v1.Fields"
            }
          }
        },
        "parameters": [
          {
            "name": "node_manager_id",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/chef.automate.api.nodes.manager.v1.FieldQuery"
            }
          }
        ],
        "tags": [
          "NodeManagerService"
        ]
      }
    },
    "/nodemanagers/id/{node_manager_id}/search-nodes": {
      "post": {
        "operationId": "SearchNodes",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.nodes.manager.v1.Nodes"
            }
          }
        },
        "parameters": [
          {
            "name": "node_manager_id",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/chef.automate.api.nodes.manager.v1.NodeQuery"
            }
          }
        ],
        "tags": [
          "NodeManagerService"
        ]
      }
    },
    "/nodemanagers/rerun/id/{id}": {
      "post": {
        "operationId": "Connect",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.nodes.manager.v1.ConnectResponse"
            }
          }
        },
        "parameters": [
          {
            "name": "id",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/chef.automate.api.nodes.manager.v1.Id"
            }
          }
        ],
        "tags": [
          "NodeManagerService"
        ]
      }
    },
    "/nodemanagers/search": {
      "post": {
        "operationId": "List",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.nodes.manager.v1.NodeManagers"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/chef.automate.api.nodes.manager.v1.Query"
            }
          }
        ],
        "tags": [
          "NodeManagerService"
        ]
      }
    }
  },
  "definitions": {
    "chef.automate.api.common.query.Filter": {
      "type": "object",
      "properties": {
        "key": {
          "type": "string",
          "description": "Field to filter on."
        },
        "exclude": {
          "type": "boolean",
          "format": "boolean",
          "description": "Include matches for this filter.(boolean)\n` + "`" + `true` + "`" + ` (default) *includes* all nodes that match this filter. \n` + "`" + `false` + "`" + ` *excludes* all nodes that match this filter."
        },
        "values": {
          "type": "array",
          "items": {
            "type": "string"
          },
          "description": "Field values to filter on."
        }
      }
    },
    "chef.automate.api.common.query.Kv": {
      "type": "object",
      "properties": {
        "key": {
          "type": "string",
          "description": "Tag key."
        },
        "value": {
          "type": "string",
          "description": "Tag value."
        }
      }
    },
    "chef.automate.api.nodes.manager.v1.ConnectResponse": {
      "type": "object"
    },
    "chef.automate.api.nodes.manager.v1.CredentialsByTags": {
      "type": "object",
      "properties": {
        "tag_key": {
          "type": "string"
        },
        "tag_value": {
          "type": "string"
        },
        "credential_ids": {
          "type": "array",
          "items": {
            "type": "string"
          }
        }
      }
    },
    "chef.automate.api.nodes.manager.v1.FieldQuery": {
      "type": "object",
      "properties": {
        "query": {
          "$ref": "#/definitions/chef.automate.api.nodes.manager.v1.Query"
        },
        "field": {
          "type": "string"
        },
        "node_manager_id": {
          "type": "string"
        }
      }
    },
    "chef.automate.api.nodes.manager.v1.Fields": {
      "type": "object",
      "properties": {
        "fields": {
          "type": "array",
          "items": {
            "type": "string"
          }
        }
      }
    },
    "chef.automate.api.nodes.manager.v1.Id": {
      "type": "object",
      "properties": {
        "id": {
          "type": "string"
        }
      }
    },
    "chef.automate.api.nodes.manager.v1.Ids": {
      "type": "object",
      "properties": {
        "ids": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/chef.automate.api.nodes.manager.v1.Id"
          }
        }
      }
    },
    "chef.automate.api.nodes.manager.v1.NodeManager": {
      "type": "object",
      "properties": {
        "id": {
          "type": "string"
        },
        "name": {
          "type": "string"
        },
        "type": {
          "type": "string"
        },
        "credential_id": {
          "type": "string"
        },
        "instance_credentials": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/chef.automate.api.nodes.manager.v1.CredentialsByTags"
          }
        },
        "status": {
          "type": "string"
        },
        "account_id": {
          "type": "string"
        },
        "date_added": {
          "type": "string",
          "format": "date-time"
        },
        "credential_data": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/chef.automate.api.common.query.Kv"
          }
        }
      }
    },
    "chef.automate.api.nodes.manager.v1.NodeManagers": {
      "type": "object",
      "properties": {
        "managers": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/chef.automate.api.nodes.manager.v1.NodeManager"
          }
        },
        "total": {
          "type": "integer",
          "format": "int32"
        }
      }
    },
    "chef.automate.api.nodes.manager.v1.NodeQuery": {
      "type": "object",
      "properties": {
        "query": {
          "$ref": "#/definitions/chef.automate.api.nodes.manager.v1.Query"
        },
        "node_manager_id": {
          "type": "string"
        }
      }
    },
    "chef.automate.api.nodes.manager.v1.Nodes": {
      "type": "object",
      "properties": {
        "nodes": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "total": {
          "type": "integer",
          "format": "int32"
        }
      }
    },
    "chef.automate.api.nodes.manager.v1.Query": {
      "type": "object",
      "properties": {
        "filter_map": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/chef.automate.api.common.query.Filter"
          }
        },
        "order": {
          "$ref": "#/definitions/chef.automate.api.nodes.manager.v1.Query.OrderType"
        },
        "sort": {
          "type": "string"
        },
        "page": {
          "type": "integer",
          "format": "int32"
        },
        "per_page": {
          "type": "integer",
          "format": "int32"
        }
      }
    },
    "chef.automate.api.nodes.manager.v1.Query.OrderType": {
      "type": "string",
      "enum": [
        "ASC",
        "DESC"
      ],
      "default": "ASC"
    }
  }
}
`)
}
