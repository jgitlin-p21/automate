package api

func init() {
	Swagger.Add("iam_v2_policy", `{
  "swagger": "2.0",
  "info": {
    "title": "components/automate-gateway/api/iam/v2/policy.proto",
    "version": "version not set"
  },
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "paths": {
    "/iam/v2/introspect_projects": {
      "get": {
        "operationId": "IntrospectAllProjects",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.iam.v2.ListProjectsResp"
            }
          }
        },
        "tags": [
          "hidden"
        ]
      }
    },
    "/iam/v2/policies": {
      "get": {
        "summary": "List all policies",
        "description": "List all policies.",
        "operationId": "ListPolicies",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.iam.v2.ListPoliciesResp"
            }
          }
        },
        "tags": [
          "policies"
        ]
      },
      "post": {
        "summary": "Create a new policy",
        "description": "Creates a new IAM policy used to control permissions in Automate.\nA policy is composed of one or more statements that grant permissions to a set of members.\nEach statement contains a role as well as a list of projects.\n\nThe role defines a set of actions that the statement is scoped to.\nThe project list defines the set of resources that the statement is scoped to.\nPass ` + "`" + `\"projects\": [\"*\"]` + "`" + ` to scope a statement to every project.\n\nA policy's top-level projects list defines which project(s) the policy belongs to (for filtering policies by their projects),\nwhereas the statement level projects list defines which project(s) the statement applies to.\n\nThis example creates a new policy not associated with any project (because the top-level ` + "`" + `projects` + "`" + ` property is empty) that grants the ` + "`" + `viewer` + "`" + ` role\non a few projects for all LDAP teams and a custom role ` + "`" + `qa` + "`" + ` on a specific project:\n\n` + "`" + `` + "`" + `` + "`" + `\n{\n\"name\": \"My Viewer Policy\",\n\"id\": \"viewer-policy\",\n\"members\": [\"team:ldap:*\"],\n\"statements\" : [\n{\n\"role\": \"viewer\",\n\"projects\": [\"project1\", \"project2\"]\n},\n{\n\"role\": \"qa\",\n\"projects\": [\"acceptanceProject\"]\n}\n],\n\"projects\": []\n}\n` + "`" + `` + "`" + `` + "`" + `",
        "operationId": "CreatePolicy",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.iam.v2.CreatePolicyResp"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/chef.automate.api.iam.v2.CreatePolicyReq"
            }
          }
        ],
        "tags": [
          "policies"
        ]
      }
    },
    "/iam/v2/policies/{id}": {
      "get": {
        "summary": "Get a policy",
        "description": "Get a policy.",
        "operationId": "GetPolicy",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.iam.v2.GetPolicyResp"
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
          "policies"
        ]
      },
      "delete": {
        "summary": "Delete a policy",
        "description": "Delete a policy.",
        "operationId": "DeletePolicy",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.iam.v2.DeletePolicyResp"
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
          "policies"
        ]
      },
      "put": {
        "summary": "Update an existing policy",
        "description": "Very similar to create except the ID cannot be changed.",
        "operationId": "UpdatePolicy",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.iam.v2.UpdatePolicyResp"
            }
          }
        },
        "parameters": [
          {
            "name": "id",
            "description": "Unique, user-specified ID. Cannot be changed.",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/chef.automate.api.iam.v2.UpdatePolicyReq"
            }
          }
        ],
        "tags": [
          "policies"
        ]
      }
    },
    "/iam/v2/policies/{id}/members": {
      "get": {
        "summary": "List policy members",
        "description": "List all members of a specific policy.",
        "operationId": "ListPolicyMembers",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.iam.v2.ListPolicyMembersResp"
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
          "policies"
        ]
      },
      "put": {
        "summary": "Replace policy members",
        "description": "Replace the entire member list for a specific policy with a new list.",
        "operationId": "ReplacePolicyMembers",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.iam.v2.ReplacePolicyMembersResp"
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
              "$ref": "#/definitions/chef.automate.api.iam.v2.ReplacePolicyMembersReq"
            }
          }
        ],
        "tags": [
          "policies"
        ]
      }
    },
    "/iam/v2/policies/{id}/members:add": {
      "post": {
        "summary": "Add policy members",
        "description": "Add specific members to the member list for a specific policy.",
        "operationId": "AddPolicyMembers",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.iam.v2.AddPolicyMembersResp"
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
              "$ref": "#/definitions/chef.automate.api.iam.v2.AddPolicyMembersReq"
            }
          }
        ],
        "tags": [
          "policies"
        ]
      }
    },
    "/iam/v2/policies/{id}/members:remove": {
      "post": {
        "summary": "Remove policy members",
        "description": "Remove specific members from the member list for a specific policy. Silently ignores\nmembers that are not already part of the member list.",
        "operationId": "RemovePolicyMembers",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.iam.v2.RemovePolicyMembersResp"
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
              "$ref": "#/definitions/chef.automate.api.iam.v2.RemovePolicyMembersReq"
            }
          }
        ],
        "tags": [
          "policies"
        ]
      }
    },
    "/iam/v2/policy_version": {
      "get": {
        "summary": "Get IAM version",
        "description": "Returns the major and minor version of IAM that your automate installation is running.",
        "operationId": "GetPolicyVersion",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.iam.v2.GetPolicyVersionResp"
            }
          }
        },
        "tags": [
          "policies"
        ]
      }
    },
    "/iam/v2/projects": {
      "get": {
        "summary": "List all projects",
        "description": "List all projects.",
        "operationId": "ListProjects",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.iam.v2.ListProjectsResp"
            }
          }
        },
        "tags": [
          "projects"
        ]
      },
      "post": {
        "summary": "Create a project",
        "description": "Creates a new project to be used in the policies that control permissions in Automate.\n\nA project defines the scope of resources in a policy statement. Resources can be in more than one project.",
        "operationId": "CreateProject",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.iam.v2.CreateProjectResp"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/chef.automate.api.iam.v2.CreateProjectReq"
            }
          }
        ],
        "tags": [
          "projects"
        ]
      }
    },
    "/iam/v2/projects/{id}": {
      "get": {
        "summary": "Get a project",
        "description": "Get a project.",
        "operationId": "GetProject",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.iam.v2.GetProjectResp"
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
          "projects"
        ]
      },
      "delete": {
        "summary": "Delete a project",
        "description": "Deletes the project from any resources tagged with it.\n\nAlso deletes this project from any project list in all statements.\nIf the resulting project list for a given statement is empty, it is deleted.\nIf the resulting policy has no statements, it is also deleted.",
        "operationId": "DeleteProject",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.iam.v2.DeleteProjectResp"
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
          "projects"
        ]
      },
      "put": {
        "summary": "Update a project",
        "description": "Update an existing project. Very similar to create except the ID cannot be changed.",
        "operationId": "UpdateProject",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.iam.v2.UpdateProjectResp"
            }
          }
        },
        "parameters": [
          {
            "name": "id",
            "description": "Unique, user-specified ID. Cannot be changed.",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/chef.automate.api.iam.v2.UpdateProjectReq"
            }
          }
        ],
        "tags": [
          "projects"
        ]
      }
    },
    "/iam/v2/roles": {
      "get": {
        "summary": "List all roles",
        "description": "List all roles.",
        "operationId": "ListRoles",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.iam.v2.ListRolesResp"
            }
          }
        },
        "tags": [
          "roles"
        ]
      },
      "post": {
        "summary": "Create a new role",
        "description": "Creates a new role to be used in the policies that control permissions in Automate.\n\nA role defines the scope of actions in a policy statement.",
        "operationId": "CreateRole",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.iam.v2.CreateRoleResp"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/chef.automate.api.iam.v2.CreateRoleReq"
            }
          }
        ],
        "tags": [
          "roles"
        ]
      }
    },
    "/iam/v2/roles/{id}": {
      "get": {
        "summary": "Get a role",
        "description": "Get a role.",
        "operationId": "GetRole",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.iam.v2.GetRoleResp"
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
          "roles"
        ]
      },
      "delete": {
        "summary": "Delete a role",
        "description": "Delete a specified role and removes it from any statements that may have been using it.\nIf such a statement has no other associated actions, the statement is deleted as well.\nSimilarly, if that statement removal results in a policy with no other statements,\nthat policy is removed as well.",
        "operationId": "DeleteRole",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.iam.v2.DeleteRoleResp"
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
          "roles"
        ]
      },
      "put": {
        "summary": "Update a role",
        "description": "Update an existing role. Very similar to create except the ID cannot be changed.",
        "operationId": "UpdateRole",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.iam.v2.UpdateRoleResp"
            }
          }
        },
        "parameters": [
          {
            "name": "id",
            "description": "Unique, user-specified ID. Cannot be changed.",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/chef.automate.api.iam.v2.UpdateRoleReq"
            }
          }
        ],
        "tags": [
          "roles"
        ]
      }
    },
    "/iam/v2beta/introspect_projects": {
      "get": {
        "operationId": "IntrospectAllProjects2",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.iam.v2.ListProjectsResp"
            }
          }
        },
        "tags": [
          "hidden"
        ]
      }
    },
    "/iam/v2beta/policies": {
      "get": {
        "summary": "List all policies",
        "description": "List all policies.",
        "operationId": "ListPolicies2",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.iam.v2.ListPoliciesResp"
            }
          }
        },
        "tags": [
          "policies"
        ]
      },
      "post": {
        "summary": "Create a new policy",
        "description": "Creates a new IAM policy used to control permissions in Automate.\nA policy is composed of one or more statements that grant permissions to a set of members.\nEach statement contains a role as well as a list of projects.\n\nThe role defines a set of actions that the statement is scoped to.\nThe project list defines the set of resources that the statement is scoped to.\nPass ` + "`" + `\"projects\": [\"*\"]` + "`" + ` to scope a statement to every project.\n\nA policy's top-level projects list defines which project(s) the policy belongs to (for filtering policies by their projects),\nwhereas the statement level projects list defines which project(s) the statement applies to.\n\nThis example creates a new policy not associated with any project (because the top-level ` + "`" + `projects` + "`" + ` property is empty) that grants the ` + "`" + `viewer` + "`" + ` role\non a few projects for all LDAP teams and a custom role ` + "`" + `qa` + "`" + ` on a specific project:\n\n` + "`" + `` + "`" + `` + "`" + `\n{\n\"name\": \"My Viewer Policy\",\n\"id\": \"viewer-policy\",\n\"members\": [\"team:ldap:*\"],\n\"statements\" : [\n{\n\"role\": \"viewer\",\n\"projects\": [\"project1\", \"project2\"]\n},\n{\n\"role\": \"qa\",\n\"projects\": [\"acceptanceProject\"]\n}\n],\n\"projects\": []\n}\n` + "`" + `` + "`" + `` + "`" + `",
        "operationId": "CreatePolicy2",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.iam.v2.CreatePolicyResp"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/chef.automate.api.iam.v2.CreatePolicyReq"
            }
          }
        ],
        "tags": [
          "policies"
        ]
      }
    },
    "/iam/v2beta/policies/{id}": {
      "get": {
        "summary": "Get a policy",
        "description": "Get a policy.",
        "operationId": "GetPolicy2",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.iam.v2.GetPolicyResp"
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
          "policies"
        ]
      },
      "delete": {
        "summary": "Delete a policy",
        "description": "Delete a policy.",
        "operationId": "DeletePolicy2",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.iam.v2.DeletePolicyResp"
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
          "policies"
        ]
      },
      "put": {
        "summary": "Update an existing policy",
        "description": "Very similar to create except the ID cannot be changed.",
        "operationId": "UpdatePolicy2",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.iam.v2.UpdatePolicyResp"
            }
          }
        },
        "parameters": [
          {
            "name": "id",
            "description": "Unique, user-specified ID. Cannot be changed.",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/chef.automate.api.iam.v2.UpdatePolicyReq"
            }
          }
        ],
        "tags": [
          "policies"
        ]
      }
    },
    "/iam/v2beta/policies/{id}/members": {
      "get": {
        "summary": "List policy members",
        "description": "List all members of a specific policy.",
        "operationId": "ListPolicyMembers2",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.iam.v2.ListPolicyMembersResp"
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
          "policies"
        ]
      },
      "put": {
        "summary": "Replace policy members",
        "description": "Replace the entire member list for a specific policy with a new list.",
        "operationId": "ReplacePolicyMembers2",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.iam.v2.ReplacePolicyMembersResp"
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
              "$ref": "#/definitions/chef.automate.api.iam.v2.ReplacePolicyMembersReq"
            }
          }
        ],
        "tags": [
          "policies"
        ]
      }
    },
    "/iam/v2beta/policies/{id}/members:add": {
      "post": {
        "summary": "Add policy members",
        "description": "Add specific members to the member list for a specific policy.",
        "operationId": "AddPolicyMembers2",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.iam.v2.AddPolicyMembersResp"
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
              "$ref": "#/definitions/chef.automate.api.iam.v2.AddPolicyMembersReq"
            }
          }
        ],
        "tags": [
          "policies"
        ]
      }
    },
    "/iam/v2beta/policies/{id}/members:remove": {
      "post": {
        "summary": "Remove policy members",
        "description": "Remove specific members from the member list for a specific policy. Silently ignores\nmembers that are not already part of the member list.",
        "operationId": "RemovePolicyMembers2",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.iam.v2.RemovePolicyMembersResp"
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
              "$ref": "#/definitions/chef.automate.api.iam.v2.RemovePolicyMembersReq"
            }
          }
        ],
        "tags": [
          "policies"
        ]
      }
    },
    "/iam/v2beta/policy_version": {
      "get": {
        "summary": "Get IAM version",
        "description": "Returns the major and minor version of IAM that your automate installation is running.",
        "operationId": "GetPolicyVersion2",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.iam.v2.GetPolicyVersionResp"
            }
          }
        },
        "tags": [
          "policies"
        ]
      }
    },
    "/iam/v2beta/projects": {
      "get": {
        "summary": "List all projects",
        "description": "List all projects.",
        "operationId": "ListProjects2",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.iam.v2.ListProjectsResp"
            }
          }
        },
        "tags": [
          "projects"
        ]
      },
      "post": {
        "summary": "Create a project",
        "description": "Creates a new project to be used in the policies that control permissions in Automate.\n\nA project defines the scope of resources in a policy statement. Resources can be in more than one project.",
        "operationId": "CreateProject2",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.iam.v2.CreateProjectResp"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/chef.automate.api.iam.v2.CreateProjectReq"
            }
          }
        ],
        "tags": [
          "projects"
        ]
      }
    },
    "/iam/v2beta/projects/{id}": {
      "get": {
        "summary": "Get a project",
        "description": "Get a project.",
        "operationId": "GetProject2",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.iam.v2.GetProjectResp"
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
          "projects"
        ]
      },
      "delete": {
        "summary": "Delete a project",
        "description": "Deletes the project from any resources tagged with it.\n\nAlso deletes this project from any project list in all statements.\nIf the resulting project list for a given statement is empty, it is deleted.\nIf the resulting policy has no statements, it is also deleted.",
        "operationId": "DeleteProject2",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.iam.v2.DeleteProjectResp"
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
          "projects"
        ]
      },
      "put": {
        "summary": "Update a project",
        "description": "Update an existing project. Very similar to create except the ID cannot be changed.",
        "operationId": "UpdateProject2",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.iam.v2.UpdateProjectResp"
            }
          }
        },
        "parameters": [
          {
            "name": "id",
            "description": "Unique, user-specified ID. Cannot be changed.",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/chef.automate.api.iam.v2.UpdateProjectReq"
            }
          }
        ],
        "tags": [
          "projects"
        ]
      }
    },
    "/iam/v2beta/roles": {
      "get": {
        "summary": "List all roles",
        "description": "List all roles.",
        "operationId": "ListRoles2",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.iam.v2.ListRolesResp"
            }
          }
        },
        "tags": [
          "roles"
        ]
      },
      "post": {
        "summary": "Create a new role",
        "description": "Creates a new role to be used in the policies that control permissions in Automate.\n\nA role defines the scope of actions in a policy statement.",
        "operationId": "CreateRole2",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.iam.v2.CreateRoleResp"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/chef.automate.api.iam.v2.CreateRoleReq"
            }
          }
        ],
        "tags": [
          "roles"
        ]
      }
    },
    "/iam/v2beta/roles/{id}": {
      "get": {
        "summary": "Get a role",
        "description": "Get a role.",
        "operationId": "GetRole2",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.iam.v2.GetRoleResp"
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
          "roles"
        ]
      },
      "delete": {
        "summary": "Delete a role",
        "description": "Delete a specified role and removes it from any statements that may have been using it.\nIf such a statement has no other associated actions, the statement is deleted as well.\nSimilarly, if that statement removal results in a policy with no other statements,\nthat policy is removed as well.",
        "operationId": "DeleteRole2",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.iam.v2.DeleteRoleResp"
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
          "roles"
        ]
      },
      "put": {
        "summary": "Update a role",
        "description": "Update an existing role. Very similar to create except the ID cannot be changed.",
        "operationId": "UpdateRole2",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.iam.v2.UpdateRoleResp"
            }
          }
        },
        "parameters": [
          {
            "name": "id",
            "description": "Unique, user-specified ID. Cannot be changed.",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/chef.automate.api.iam.v2.UpdateRoleReq"
            }
          }
        ],
        "tags": [
          "roles"
        ]
      }
    }
  },
  "definitions": {
    "chef.automate.api.iam.v2.AddPolicyMembersReq": {
      "type": "object",
      "properties": {
        "id": {
          "type": "string"
        },
        "members": {
          "type": "array",
          "items": {
            "type": "string"
          },
          "description": "List of members to add to the policy."
        }
      }
    },
    "chef.automate.api.iam.v2.AddPolicyMembersResp": {
      "type": "object",
      "properties": {
        "members": {
          "type": "array",
          "items": {
            "type": "string"
          },
          "description": "Resulting list of policy members."
        }
      }
    },
    "chef.automate.api.iam.v2.CreatePolicyReq": {
      "type": "object",
      "properties": {
        "id": {
          "type": "string",
          "description": "Unique, user-specified ID. Cannot be changed."
        },
        "name": {
          "type": "string",
          "description": "Name for the new policy."
        },
        "members": {
          "type": "array",
          "items": {
            "type": "string"
          },
          "description": "Members affected by this policy."
        },
        "statements": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/chef.automate.api.iam.v2.Statement"
          },
          "description": "Statements for the new policy. Must contain one or more."
        },
        "projects": {
          "type": "array",
          "items": {
            "type": "string"
          },
          "description": "The list of projects this policy belongs to."
        }
      },
      "description": "Does not contain type as the enduser can only create 'custom' policies."
    },
    "chef.automate.api.iam.v2.CreatePolicyResp": {
      "type": "object",
      "properties": {
        "policy": {
          "$ref": "#/definitions/chef.automate.api.iam.v2.Policy"
        }
      }
    },
    "chef.automate.api.iam.v2.CreateProjectReq": {
      "type": "object",
      "properties": {
        "id": {
          "type": "string",
          "description": "Unique, user-specified ID. Cannot be changed."
        },
        "name": {
          "type": "string",
          "description": "Name for the new project."
        }
      }
    },
    "chef.automate.api.iam.v2.CreateProjectResp": {
      "type": "object",
      "properties": {
        "project": {
          "$ref": "#/definitions/chef.automate.api.iam.v2.Project"
        }
      }
    },
    "chef.automate.api.iam.v2.CreateRoleReq": {
      "type": "object",
      "properties": {
        "id": {
          "type": "string",
          "description": "Unique, user-specified ID. Cannot be changed."
        },
        "name": {
          "type": "string",
          "description": "Name for the new role."
        },
        "actions": {
          "type": "array",
          "items": {
            "type": "string"
          },
          "description": "List of actions that this role scopes to."
        },
        "projects": {
          "type": "array",
          "items": {
            "type": "string"
          },
          "description": "The list of projects this role belongs to."
        }
      },
      "description": "Does not contain type as the enduser can only create 'custom' roles."
    },
    "chef.automate.api.iam.v2.CreateRoleResp": {
      "type": "object",
      "properties": {
        "role": {
          "$ref": "#/definitions/chef.automate.api.iam.v2.Role"
        }
      }
    },
    "chef.automate.api.iam.v2.DeletePolicyResp": {
      "type": "object"
    },
    "chef.automate.api.iam.v2.DeleteProjectResp": {
      "type": "object"
    },
    "chef.automate.api.iam.v2.DeleteRoleResp": {
      "type": "object"
    },
    "chef.automate.api.iam.v2.Flag": {
      "type": "string",
      "enum": [
        "VERSION_2_0",
        "VERSION_2_1"
      ],
      "default": "VERSION_2_0"
    },
    "chef.automate.api.iam.v2.GetPolicyResp": {
      "type": "object",
      "properties": {
        "policy": {
          "$ref": "#/definitions/chef.automate.api.iam.v2.Policy"
        }
      }
    },
    "chef.automate.api.iam.v2.GetPolicyVersionResp": {
      "type": "object",
      "properties": {
        "version": {
          "$ref": "#/definitions/chef.automate.api.iam.v2.Version"
        }
      }
    },
    "chef.automate.api.iam.v2.GetProjectResp": {
      "type": "object",
      "properties": {
        "project": {
          "$ref": "#/definitions/chef.automate.api.iam.v2.Project"
        }
      }
    },
    "chef.automate.api.iam.v2.GetRoleResp": {
      "type": "object",
      "properties": {
        "role": {
          "$ref": "#/definitions/chef.automate.api.iam.v2.Role"
        }
      }
    },
    "chef.automate.api.iam.v2.ListPoliciesResp": {
      "type": "object",
      "properties": {
        "policies": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/chef.automate.api.iam.v2.Policy"
          }
        }
      }
    },
    "chef.automate.api.iam.v2.ListPolicyMembersResp": {
      "type": "object",
      "properties": {
        "members": {
          "type": "array",
          "items": {
            "type": "string"
          },
          "description": "List of policy members."
        }
      }
    },
    "chef.automate.api.iam.v2.ListProjectsResp": {
      "type": "object",
      "properties": {
        "projects": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/chef.automate.api.iam.v2.Project"
          }
        }
      }
    },
    "chef.automate.api.iam.v2.ListRolesResp": {
      "type": "object",
      "properties": {
        "roles": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/chef.automate.api.iam.v2.Role"
          }
        }
      }
    },
    "chef.automate.api.iam.v2.Policy": {
      "type": "object",
      "properties": {
        "name": {
          "type": "string",
          "description": "Name for the policy."
        },
        "id": {
          "type": "string",
          "description": "Unique, user-specified ID. Cannot be changed."
        },
        "type": {
          "$ref": "#/definitions/chef.automate.api.iam.v2.Type",
          "description": "Whether this policy is user created or chef managed.\nOne of ` + "`" + `CUSTOM` + "`" + ` or ` + "`" + `CHEF_MANAGED` + "`" + `, respectively."
        },
        "members": {
          "type": "array",
          "items": {
            "type": "string"
          },
          "description": "Members affected by this policy."
        },
        "statements": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/chef.automate.api.iam.v2.Statement"
          },
          "description": "Statements for the policy. Must contain one or more."
        },
        "projects": {
          "type": "array",
          "items": {
            "type": "string"
          },
          "description": "The list of projects this policy belongs to."
        }
      }
    },
    "chef.automate.api.iam.v2.Project": {
      "type": "object",
      "properties": {
        "name": {
          "type": "string",
          "description": "Name for the project."
        },
        "id": {
          "type": "string",
          "description": "Unique, user-specified ID. Cannot be changed."
        },
        "type": {
          "$ref": "#/definitions/chef.automate.api.iam.v2.Type",
          "description": "Whether this policy is user created or chef managed.\nOne of ` + "`" + `CUSTOM` + "`" + ` or ` + "`" + `CHEF_MANAGED` + "`" + `, respectively."
        },
        "status": {
          "$ref": "#/definitions/chef.automate.api.iam.v2.ProjectRulesStatus",
          "description": "The current status of the rules for this project."
        }
      }
    },
    "chef.automate.api.iam.v2.ProjectRulesStatus": {
      "type": "string",
      "enum": [
        "PROJECT_RULES_STATUS_UNSET",
        "RULES_APPLIED",
        "EDITS_PENDING",
        "NO_RULES"
      ],
      "default": "PROJECT_RULES_STATUS_UNSET"
    },
    "chef.automate.api.iam.v2.RemovePolicyMembersReq": {
      "type": "object",
      "properties": {
        "id": {
          "type": "string"
        },
        "members": {
          "type": "array",
          "items": {
            "type": "string"
          },
          "description": "List of members to remove from the policy."
        }
      }
    },
    "chef.automate.api.iam.v2.RemovePolicyMembersResp": {
      "type": "object",
      "properties": {
        "members": {
          "type": "array",
          "items": {
            "type": "string"
          },
          "description": "Resulting list of policy members."
        }
      }
    },
    "chef.automate.api.iam.v2.ReplacePolicyMembersReq": {
      "type": "object",
      "properties": {
        "id": {
          "type": "string"
        },
        "members": {
          "type": "array",
          "items": {
            "type": "string"
          },
          "description": "List of members that replaces previous policy member list."
        }
      }
    },
    "chef.automate.api.iam.v2.ReplacePolicyMembersResp": {
      "type": "object",
      "properties": {
        "members": {
          "type": "array",
          "items": {
            "type": "string"
          },
          "description": "Resulting list of policy members."
        }
      }
    },
    "chef.automate.api.iam.v2.ResetToV1Resp": {
      "type": "object"
    },
    "chef.automate.api.iam.v2.Role": {
      "type": "object",
      "properties": {
        "name": {
          "type": "string",
          "description": "Name for the role."
        },
        "id": {
          "type": "string",
          "description": "Unique, user-specified ID. Cannot be changed."
        },
        "type": {
          "$ref": "#/definitions/chef.automate.api.iam.v2.Type",
          "description": "Whether this policy is user created or chef managed.\nOne of ` + "`" + `CUSTOM` + "`" + ` or ` + "`" + `CHEF_MANAGED` + "`" + `, respectively."
        },
        "actions": {
          "type": "array",
          "items": {
            "type": "string"
          },
          "description": "List of actions that this role scopes to."
        },
        "projects": {
          "type": "array",
          "items": {
            "type": "string"
          },
          "description": "The list of projects this role belongs to."
        }
      }
    },
    "chef.automate.api.iam.v2.Statement": {
      "type": "object",
      "properties": {
        "effect": {
          "$ref": "#/definitions/chef.automate.api.iam.v2.Statement.Effect",
          "description": "Whether the statement allows or denies.\nOne of ` + "`" + `ALLOW` + "`" + ` or ` + "`" + `DENY` + "`" + `, respectively."
        },
        "actions": {
          "type": "array",
          "items": {
            "type": "string"
          },
          "description": "Actions defined inline. Best practices recommend that you use custom roles where practical."
        },
        "role": {
          "type": "string",
          "description": "The role defines a set of actions that the statement is scoped to."
        },
        "resources": {
          "type": "array",
          "items": {
            "type": "string"
          },
          "description": "DEPRECATED: Resources defined inline. Use projects instead."
        },
        "projects": {
          "type": "array",
          "items": {
            "type": "string"
          },
          "description": "The project list defines the set of resources that the statement is scoped to."
        }
      }
    },
    "chef.automate.api.iam.v2.Statement.Effect": {
      "type": "string",
      "enum": [
        "ALLOW",
        "DENY"
      ],
      "default": "ALLOW"
    },
    "chef.automate.api.iam.v2.Type": {
      "type": "string",
      "enum": [
        "CHEF_MANAGED",
        "CUSTOM"
      ],
      "default": "CHEF_MANAGED"
    },
    "chef.automate.api.iam.v2.UpdatePolicyReq": {
      "type": "object",
      "properties": {
        "id": {
          "type": "string",
          "description": "Unique, user-specified ID. Cannot be changed."
        },
        "members": {
          "type": "array",
          "items": {
            "type": "string"
          },
          "description": "Members affected by this policy."
        },
        "statements": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/chef.automate.api.iam.v2.Statement"
          },
          "description": "Statements for the policy. Must contain one or more."
        },
        "name": {
          "type": "string",
          "description": "Name for this policy."
        },
        "projects": {
          "type": "array",
          "items": {
            "type": "string"
          },
          "description": "The list of projects this policy belongs to."
        }
      },
      "description": "Does not contain type as the enduser can only create 'custom' policies.",
      "required": [
        "id"
      ]
    },
    "chef.automate.api.iam.v2.UpdatePolicyResp": {
      "type": "object",
      "properties": {
        "policy": {
          "$ref": "#/definitions/chef.automate.api.iam.v2.Policy"
        }
      }
    },
    "chef.automate.api.iam.v2.UpdateProjectReq": {
      "type": "object",
      "properties": {
        "id": {
          "type": "string",
          "description": "Unique, user-specified ID. Cannot be changed."
        },
        "name": {
          "type": "string",
          "description": "Name for the project."
        }
      }
    },
    "chef.automate.api.iam.v2.UpdateProjectResp": {
      "type": "object",
      "properties": {
        "project": {
          "$ref": "#/definitions/chef.automate.api.iam.v2.Project"
        }
      }
    },
    "chef.automate.api.iam.v2.UpdateRoleReq": {
      "type": "object",
      "properties": {
        "id": {
          "type": "string",
          "description": "Unique, user-specified ID. Cannot be changed."
        },
        "name": {
          "type": "string",
          "description": "Name for the role."
        },
        "actions": {
          "type": "array",
          "items": {
            "type": "string"
          },
          "description": "List of actions that this role scopes to."
        },
        "projects": {
          "type": "array",
          "items": {
            "type": "string"
          },
          "description": "The list of projects this role belongs to."
        }
      }
    },
    "chef.automate.api.iam.v2.UpdateRoleResp": {
      "type": "object",
      "properties": {
        "role": {
          "$ref": "#/definitions/chef.automate.api.iam.v2.Role"
        }
      }
    },
    "chef.automate.api.iam.v2.UpgradeToV2Resp": {
      "type": "object",
      "properties": {
        "reports": {
          "type": "array",
          "items": {
            "type": "string"
          }
        }
      }
    },
    "chef.automate.api.iam.v2.Version": {
      "type": "object",
      "properties": {
        "major": {
          "$ref": "#/definitions/chef.automate.api.iam.v2.Version.VersionNumber"
        },
        "minor": {
          "$ref": "#/definitions/chef.automate.api.iam.v2.Version.VersionNumber"
        }
      }
    },
    "chef.automate.api.iam.v2.Version.VersionNumber": {
      "type": "string",
      "enum": [
        "V0",
        "V1",
        "V2"
      ],
      "default": "V0"
    }
  }
}
`)
}
