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
        "summary": "Create a nodemanager",
        "description": "Creates a nodemanager given a name, credential id or credential data, and type.\n\nExample:\n` + "`" + `` + "`" + `` + "`" + `\n{\n\"name\": \"my aws api integration with session token\",\n\"type\": \"aws-api\",\n\"instance_credentials\": [],\n\"credential_data\": [\n{\"key\": \"AWS_ACCESS_KEY_ID\", \"value\": \"value\" },\n{\"key\": \"AWS_SECRET_ACCESS_KEY\", \"value\": \"value\" },\n{\"key\": \"AWS_SESSION_TOKEN\", \"value\": \"value\" }\n]\n}\n` + "`" + `` + "`" + `` + "`" + `",
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
        "summary": "Read a nodemanager",
        "description": "Read the details of a nodemanager given an id.",
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
            "description": "UUID for the nodemanager.",
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
        "summary": "Delete a nodemanager",
<<<<<<< HEAD:components/automate-gateway/api/manager.pb.swagger.go
        "description": "Delete a nodemanager given an id. Note this only deletes the manager itself. Any nodes\nassociated with the manager will be re-assigned to the Automate node manager.",
=======
        "description": "Delete a nodemanager given an id. Note this only deletes the manager itself. Any nodes\nassociated with the manager will be abandoned.",
>>>>>>> 1649d33856d4707df5f7edc80962d6485dc17949:components/automate-gateway/api/nodes_manager_manager.pb.swagger.go
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
            "description": "UUID for the nodemanager.",
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
        "summary": "Update a nodemanager",
        "description": "PUT operation to update the details for a nodemanager, such as the name or associated credential id or data.\nPlease note that this is a PUT operation, so all nodemanager details included in the create function\nshould be included in the PUT message to update.",
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
            "description": "UUID for the nodemanager.",
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
        "summary": "Delete a nodemanager and set nodes to have a state of stopped",
<<<<<<< HEAD:components/automate-gateway/api/manager.pb.swagger.go
        "description": "Delete a nodemanager and update all associated nodes to have a state of ` + "`" + `stopped` + "`" + `.",
=======
        "description": "Delete a nodemanager and update all associated nodes to have a state of stopped.",
>>>>>>> 1649d33856d4707df5f7edc80962d6485dc17949:components/automate-gateway/api/nodes_manager_manager.pb.swagger.go
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
            "description": "UUID for the nodemanager.",
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
        "summary": "Delete a nodemanager and set nodes to have a state of terminated",
        "description": "Delete a nodemanager and update all associated nodes to have a state of terminated.",
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
            "description": "UUID for the nodemanager.",
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
        "summary": "Delete a nodemanager and all of its nodes",
        "description": "Delete a nodemanager and all associated nodes given a nodemanager id.",
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
            "description": "UUID for the nodemanager.",
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
        "summary": "Search node fields",
<<<<<<< HEAD:components/automate-gateway/api/manager.pb.swagger.go
        "description": "Searches the available values for a given field across all nodes associated with the nodemanager id.\nPossible fields: regions, tags, name, subscription_id\n\nExample:\n` + "`" + `` + "`" + `` + "`" + `\n{\n\"query\": {\n\"filter_map\":[]\n},\n\"field\": \"name\"\n}\n` + "`" + `` + "`" + `` + "`" + `",
=======
        "description": "Searches the available values for a given field and nodemanager id.\nPossible fields: regions, tags, name, subscription_id\n\nExample:\n` + "`" + `` + "`" + `` + "`" + `\n{\n\"query\": {\n\"filter_map\":[]\n},\n\"field\": \"name\"\n}\n` + "`" + `` + "`" + `` + "`" + `",
>>>>>>> 1649d33856d4707df5f7edc80962d6485dc17949:components/automate-gateway/api/nodes_manager_manager.pb.swagger.go
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
            "description": "Nodemanager id for which the search is being made.",
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
        "summary": "Search nodes",
        "description": "Searches the available nodes for a given nodemanager id.\n\nExample:\n` + "`" + `` + "`" + `` + "`" + `\n{\n\"query\": {\n\"filter_map\":[]\n}\t\n}\n` + "`" + `` + "`" + `` + "`" + `",
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
            "description": "Nodemanager id for which the search is being made.",
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
        "summary": "Connect",
<<<<<<< HEAD:components/automate-gateway/api/manager.pb.swagger.go
        "description": "Attempts to reach the API for the given nodemanager id to validate the \ncredentials associated with the nodemanager.",
=======
        "description": "Attempts to reach the API for the given nodemanager id to check validity of the \ncredentials associated with the nodemanager.",
>>>>>>> 1649d33856d4707df5f7edc80962d6485dc17949:components/automate-gateway/api/nodes_manager_manager.pb.swagger.go
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
            "description": "UUID for the nodemanager.",
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
        "summary": "List of nodemanagers",
<<<<<<< HEAD:components/automate-gateway/api/manager.pb.swagger.go
        "description": "Returns a list of nodemanagers matching the query.\nSupports filtering, sorting, and pagination.\n\nValid filtering fields: manager_type\n\nValid sorting fields: name, type, status, status_message, date_added\n\nExample:\n` + "`" + `` + "`" + `` + "`" + `\n{\n\"filter_map\": [\n{\"key\": \"manager_type\", \"values\":[\"aws-ec2\"]}\n], \n\"sort\": \"date_added\"\n}\n` + "`" + `` + "`" + `` + "`" + `",
=======
        "description": "Returns a list of nodemanagers matching the query.\nSupports filtering, sorting, and pagination.\nValid filtering fields: 'manager_type'\nValid sorting fields: name, type, status, status_message, date_added\n\nExample:\n` + "`" + `` + "`" + `` + "`" + `\n{\n\"filter_map\": [\n{\"key\": \"manager_type\", \"values\":[\"aws-ec2\"]}\n], \n\"sort\": \"date_added\"\n}\n` + "`" + `` + "`" + `` + "`" + `",
>>>>>>> 1649d33856d4707df5f7edc80962d6485dc17949:components/automate-gateway/api/nodes_manager_manager.pb.swagger.go
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
    "chef.automate.api.nodes.manager.v1.ConnectResponse": {
      "type": "object"
    },
    "chef.automate.api.nodes.manager.v1.CredentialsByTags": {
      "type": "object",
      "properties": {
        "tag_key": {
          "type": "string",
          "description": "Tag key to match on."
        },
        "tag_value": {
          "type": "string",
          "description": "Tag value to match on."
        },
        "credential_ids": {
          "type": "array",
          "items": {
            "type": "string"
          },
          "description": "List of credential ids to associate with the key/value pair."
        }
      }
    },
    "chef.automate.api.nodes.manager.v1.FieldQuery": {
      "type": "object",
      "properties": {
        "query": {
          "$ref": "#/definitions/chef.automate.api.nodes.manager.v1.Query",
          "description": "Query details (filters) to be applied to the results."
        },
        "field": {
          "type": "string",
          "description": "Field to search on."
        },
        "node_manager_id": {
          "type": "string",
          "description": "Nodemanager id for which the search is being made."
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
          },
          "description": "List of available fields matching the requested criteria."
        }
      }
    },
    "chef.automate.api.nodes.manager.v1.Id": {
      "type": "object",
      "properties": {
        "id": {
          "type": "string",
          "description": "UUID for the nodemanager."
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
          },
          "description": "List of nodemanager UUIDs."
        }
      }
    },
    "chef.automate.api.nodes.manager.v1.NodeManager": {
      "type": "object",
      "properties": {
        "id": {
          "type": "string",
          "description": "UUID for the nodemanager."
        },
        "name": {
          "type": "string",
          "description": "User defined name for the nodemanager."
        },
        "type": {
          "type": "string",
          "description": "Type of nodemanager (aws-ec2, azure-vm, aws-api, azure-api, gcp)."
        },
        "credential_id": {
          "type": "string",
          "description": "UUID of credential containing the information to connect to aws, azure, or gcp."
        },
        "instance_credentials": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/chef.automate.api.nodes.manager.v1.CredentialsByTags"
          },
          "description": "List of tag and credential uuid associations to make. These are ssh, winrm, and sudo creds used to access instances."
        },
        "status": {
          "type": "string",
          "description": "Status of the nodemanager (reachable, unreachable)."
        },
        "account_id": {
          "type": "string",
          "description": "Account id associated with the nodemanager."
        },
        "date_added": {
          "type": "string",
          "format": "date-time",
          "description": "Date the nodemanager was created."
        },
        "credential_data": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/chef.automate.domain.compliance.api.common.Kv"
          },
          "description": "Credential data for the nodemanager. This field is used when a credential\nhas not yet been created, to be able to include credential data (such as AWS_ACCESS_KEY) inline."
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
          },
          "description": "List of nodemanagers."
        },
        "total": {
          "type": "integer",
          "format": "int32",
          "description": "Total count of nodemanagers."
        }
      }
    },
    "chef.automate.api.nodes.manager.v1.NodeQuery": {
      "type": "object",
      "properties": {
        "query": {
          "$ref": "#/definitions/chef.automate.api.nodes.manager.v1.Query",
          "description": "Query details (filters) to be applied to the results."
        },
        "node_manager_id": {
          "type": "string",
          "description": "Nodemanager id for which the search is being made."
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
          },
          "description": "List of node names matching the request."
        },
        "total": {
          "type": "integer",
          "format": "int32",
          "description": "Total count of node names matching the request."
        }
      }
    },
    "chef.automate.api.nodes.manager.v1.Query": {
      "type": "object",
      "properties": {
        "filter_map": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/chef.automate.domain.compliance.api.common.Filter"
          },
          "description": "Filters to be applied to the query."
        },
        "order": {
          "$ref": "#/definitions/chef.automate.api.nodes.manager.v1.Query.OrderType"
        },
        "sort": {
          "type": "string",
          "description": "Field to sort on."
        },
        "page": {
          "type": "integer",
          "format": "int32",
          "description": "Page number of results to return."
        },
        "per_page": {
          "type": "integer",
          "format": "int32",
          "description": "Count of results that should be returned for each page."
        }
      }
    },
    "chef.automate.api.nodes.manager.v1.Query.OrderType": {
      "type": "string",
      "enum": [
        "ASC",
        "DESC"
      ],
      "default": "ASC",
      "description": "Sort the results in ascending or descending order."
    },
    "chef.automate.domain.compliance.api.common.Filter": {
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
    "chef.automate.domain.compliance.api.common.Kv": {
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
    }
  }
}
`)
}
