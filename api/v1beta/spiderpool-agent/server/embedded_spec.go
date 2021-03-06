// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2022 Authors of spidernet-io
// SPDX-License-Identifier: Apache-2.0

package server

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "swagger": "2.0",
  "info": {
    "description": "Spiderpool Agent",
    "title": "Spiderpool-Agent API",
    "version": "v1beta"
  },
  "basePath": "/v1beta",
  "paths": {
    "/ipam/healthy": {
      "get": {
        "description": "Check spiderpool daemonset health to make sure whether it's ready\nfor CNI plugin usage\n",
        "tags": [
          "connectivity"
        ],
        "summary": "Get health of spiderpool daemon",
        "responses": {
          "200": {
            "description": "Success"
          },
          "500": {
            "description": "Failed"
          }
        }
      }
    },
    "/ipam/ip": {
      "post": {
        "description": "Send a request to daemonset to ask for an ip assignment\n",
        "tags": [
          "daemonset"
        ],
        "summary": "Get ip from spiderpool daemon",
        "responses": {
          "200": {
            "description": "Success"
          },
          "500": {
            "description": "Allocation failure"
          }
        }
      },
      "delete": {
        "description": "Send a request to daemonset to ask for an ip deleting\n",
        "tags": [
          "daemonset"
        ],
        "summary": "Delete ip from spiderpool daemon",
        "responses": {
          "200": {
            "description": "Success"
          },
          "500": {
            "description": "Failed"
          }
        }
      }
    },
    "/ipam/ips": {
      "post": {
        "description": "Assign multiple ip for a pod, case for spiderflat compent\n",
        "tags": [
          "daemonset"
        ],
        "summary": "Assign multiple ip as a batch",
        "responses": {
          "200": {
            "description": "Success"
          },
          "500": {
            "description": "Allocation failure"
          }
        }
      },
      "delete": {
        "description": "Delete multiple ip for a pod, case for spiderflat compent\n",
        "tags": [
          "daemonset"
        ],
        "summary": "Delete multiple ip as a batch",
        "responses": {
          "200": {
            "description": "Success"
          },
          "500": {
            "description": "Failed"
          }
        }
      }
    },
    "/runtime/liveness": {
      "get": {
        "description": "Check pod liveness probe",
        "tags": [
          "runtime"
        ],
        "summary": "Liveness probe",
        "responses": {
          "200": {
            "description": "Success"
          },
          "500": {
            "description": "Failed"
          }
        }
      }
    },
    "/runtime/readiness": {
      "get": {
        "description": "Check pod readiness probe",
        "tags": [
          "runtime"
        ],
        "summary": "Readiness probe",
        "responses": {
          "200": {
            "description": "Success"
          },
          "500": {
            "description": "Failed"
          }
        }
      }
    },
    "/runtime/startup": {
      "get": {
        "description": "Check pod startup probe",
        "tags": [
          "runtime"
        ],
        "summary": "Startup probe",
        "responses": {
          "200": {
            "description": "Success"
          },
          "500": {
            "description": "Failed"
          }
        }
      }
    },
    "/workloadendpoint": {
      "get": {
        "description": "Get workloadendpoint details for spiderflat use\n",
        "tags": [
          "daemonset"
        ],
        "summary": "Get workloadendpoint status",
        "responses": {
          "200": {
            "description": "Success"
          },
          "500": {
            "description": "Get workloadendpoint failure"
          }
        }
      }
    }
  },
  "x-schemes": [
    "unix"
  ]
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "swagger": "2.0",
  "info": {
    "description": "Spiderpool Agent",
    "title": "Spiderpool-Agent API",
    "version": "v1beta"
  },
  "basePath": "/v1beta",
  "paths": {
    "/ipam/healthy": {
      "get": {
        "description": "Check spiderpool daemonset health to make sure whether it's ready\nfor CNI plugin usage\n",
        "tags": [
          "connectivity"
        ],
        "summary": "Get health of spiderpool daemon",
        "responses": {
          "200": {
            "description": "Success"
          },
          "500": {
            "description": "Failed"
          }
        }
      }
    },
    "/ipam/ip": {
      "post": {
        "description": "Send a request to daemonset to ask for an ip assignment\n",
        "tags": [
          "daemonset"
        ],
        "summary": "Get ip from spiderpool daemon",
        "responses": {
          "200": {
            "description": "Success"
          },
          "500": {
            "description": "Allocation failure"
          }
        }
      },
      "delete": {
        "description": "Send a request to daemonset to ask for an ip deleting\n",
        "tags": [
          "daemonset"
        ],
        "summary": "Delete ip from spiderpool daemon",
        "responses": {
          "200": {
            "description": "Success"
          },
          "500": {
            "description": "Failed"
          }
        }
      }
    },
    "/ipam/ips": {
      "post": {
        "description": "Assign multiple ip for a pod, case for spiderflat compent\n",
        "tags": [
          "daemonset"
        ],
        "summary": "Assign multiple ip as a batch",
        "responses": {
          "200": {
            "description": "Success"
          },
          "500": {
            "description": "Allocation failure"
          }
        }
      },
      "delete": {
        "description": "Delete multiple ip for a pod, case for spiderflat compent\n",
        "tags": [
          "daemonset"
        ],
        "summary": "Delete multiple ip as a batch",
        "responses": {
          "200": {
            "description": "Success"
          },
          "500": {
            "description": "Failed"
          }
        }
      }
    },
    "/runtime/liveness": {
      "get": {
        "description": "Check pod liveness probe",
        "tags": [
          "runtime"
        ],
        "summary": "Liveness probe",
        "responses": {
          "200": {
            "description": "Success"
          },
          "500": {
            "description": "Failed"
          }
        }
      }
    },
    "/runtime/readiness": {
      "get": {
        "description": "Check pod readiness probe",
        "tags": [
          "runtime"
        ],
        "summary": "Readiness probe",
        "responses": {
          "200": {
            "description": "Success"
          },
          "500": {
            "description": "Failed"
          }
        }
      }
    },
    "/runtime/startup": {
      "get": {
        "description": "Check pod startup probe",
        "tags": [
          "runtime"
        ],
        "summary": "Startup probe",
        "responses": {
          "200": {
            "description": "Success"
          },
          "500": {
            "description": "Failed"
          }
        }
      }
    },
    "/workloadendpoint": {
      "get": {
        "description": "Get workloadendpoint details for spiderflat use\n",
        "tags": [
          "daemonset"
        ],
        "summary": "Get workloadendpoint status",
        "responses": {
          "200": {
            "description": "Success"
          },
          "500": {
            "description": "Get workloadendpoint failure"
          }
        }
      }
    }
  },
  "x-schemes": [
    "unix"
  ]
}`))
}
