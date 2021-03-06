// Copyright 2019 Polyaxon, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

/*
 * Polyaxon SDKs and REST API specification.
 * Polyaxon SDKs and REST API specification.
 *
 * OpenAPI spec version: 1.0.0
 * Contact: contact@polyaxon.com
 *
 * NOTE: This class is auto generated by the swagger code generator program.
 * https://github.com/swagger-api/swagger-codegen.git
 *
 * Swagger Codegen version: 2.4.10
 *
 * Do not edit the class manually.
 *
 */

(function(root, factory) {
  if (typeof define === 'function' && define.amd) {
    // AMD. Register as an anonymous module.
    define(['ApiClient', 'model/V1Agent', 'model/V1ListAgentsResponse'], factory);
  } else if (typeof module === 'object' && module.exports) {
    // CommonJS-like environments that support module.exports, like Node.
    module.exports = factory(require('../ApiClient'), require('../model/V1Agent'), require('../model/V1ListAgentsResponse'));
  } else {
    // Browser globals (root is window)
    if (!root.PolyaxonSdk) {
      root.PolyaxonSdk = {};
    }
    root.PolyaxonSdk.AgentsV1Api = factory(root.PolyaxonSdk.ApiClient, root.PolyaxonSdk.V1Agent, root.PolyaxonSdk.V1ListAgentsResponse);
  }
}(this, function(ApiClient, V1Agent, V1ListAgentsResponse) {
  'use strict';

  /**
   * AgentsV1 service.
   * @module api/AgentsV1Api
   * @version 1.0.0
   */

  /**
   * Constructs a new AgentsV1Api. 
   * @alias module:api/AgentsV1Api
   * @class
   * @param {module:ApiClient} [apiClient] Optional API client implementation to use,
   * default to {@link module:ApiClient#instance} if unspecified.
   */
  var exports = function(apiClient) {
    this.apiClient = apiClient || ApiClient.instance;


    /**
     * Callback function to receive the result of the createAgent operation.
     * @callback module:api/AgentsV1Api~createAgentCallback
     * @param {String} error Error message, if any.
     * @param {module:model/V1Agent} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * List runs
     * @param {String} owner Owner of the namespace
     * @param {module:model/V1Agent} body Agent body
     * @param {module:api/AgentsV1Api~createAgentCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/V1Agent}
     */
    this.createAgent = function(owner, body, callback) {
      var postBody = body;

      // verify the required parameter 'owner' is set
      if (owner === undefined || owner === null) {
        throw new Error("Missing the required parameter 'owner' when calling createAgent");
      }

      // verify the required parameter 'body' is set
      if (body === undefined || body === null) {
        throw new Error("Missing the required parameter 'body' when calling createAgent");
      }


      var pathParams = {
        'owner': owner
      };
      var queryParams = {
      };
      var collectionQueryParams = {
      };
      var headerParams = {
      };
      var formParams = {
      };

      var authNames = ['ApiKey'];
      var contentTypes = ['application/json'];
      var accepts = ['application/json'];
      var returnType = V1Agent;

      return this.apiClient.callApi(
        '/api/v1/orgs/{owner}/agents', 'POST',
        pathParams, queryParams, collectionQueryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, callback
      );
    }

    /**
     * Callback function to receive the result of the deleteAgent operation.
     * @callback module:api/AgentsV1Api~deleteAgentCallback
     * @param {String} error Error message, if any.
     * @param data This operation does not return a value.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Patch run
     * @param {String} owner Owner of the namespace
     * @param {String} uuid Unique integer identifier of the entity
     * @param {module:api/AgentsV1Api~deleteAgentCallback} callback The callback function, accepting three arguments: error, data, response
     */
    this.deleteAgent = function(owner, uuid, callback) {
      var postBody = null;

      // verify the required parameter 'owner' is set
      if (owner === undefined || owner === null) {
        throw new Error("Missing the required parameter 'owner' when calling deleteAgent");
      }

      // verify the required parameter 'uuid' is set
      if (uuid === undefined || uuid === null) {
        throw new Error("Missing the required parameter 'uuid' when calling deleteAgent");
      }


      var pathParams = {
        'owner': owner,
        'uuid': uuid
      };
      var queryParams = {
      };
      var collectionQueryParams = {
      };
      var headerParams = {
      };
      var formParams = {
      };

      var authNames = ['ApiKey'];
      var contentTypes = ['application/json'];
      var accepts = ['application/json'];
      var returnType = null;

      return this.apiClient.callApi(
        '/api/v1/orgs/{owner}/agents/{uuid}', 'DELETE',
        pathParams, queryParams, collectionQueryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, callback
      );
    }

    /**
     * Callback function to receive the result of the getAgent operation.
     * @callback module:api/AgentsV1Api~getAgentCallback
     * @param {String} error Error message, if any.
     * @param {module:model/V1Agent} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Create new run
     * @param {String} owner Owner of the namespace
     * @param {String} uuid Unique integer identifier of the entity
     * @param {module:api/AgentsV1Api~getAgentCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/V1Agent}
     */
    this.getAgent = function(owner, uuid, callback) {
      var postBody = null;

      // verify the required parameter 'owner' is set
      if (owner === undefined || owner === null) {
        throw new Error("Missing the required parameter 'owner' when calling getAgent");
      }

      // verify the required parameter 'uuid' is set
      if (uuid === undefined || uuid === null) {
        throw new Error("Missing the required parameter 'uuid' when calling getAgent");
      }


      var pathParams = {
        'owner': owner,
        'uuid': uuid
      };
      var queryParams = {
      };
      var collectionQueryParams = {
      };
      var headerParams = {
      };
      var formParams = {
      };

      var authNames = ['ApiKey'];
      var contentTypes = ['application/json'];
      var accepts = ['application/json'];
      var returnType = V1Agent;

      return this.apiClient.callApi(
        '/api/v1/orgs/{owner}/agents/{uuid}', 'GET',
        pathParams, queryParams, collectionQueryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, callback
      );
    }

    /**
     * Callback function to receive the result of the listAgentNames operation.
     * @callback module:api/AgentsV1Api~listAgentNamesCallback
     * @param {String} error Error message, if any.
     * @param {module:model/V1ListAgentsResponse} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * List bookmarked runs for user
     * @param {String} owner Owner of the namespace
     * @param {Object} opts Optional parameters
     * @param {Number} opts.offset Pagination offset.
     * @param {Number} opts.limit Limit size.
     * @param {String} opts.sort Sort to order the search.
     * @param {String} opts.query Query filter the search search.
     * @param {module:api/AgentsV1Api~listAgentNamesCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/V1ListAgentsResponse}
     */
    this.listAgentNames = function(owner, opts, callback) {
      opts = opts || {};
      var postBody = null;

      // verify the required parameter 'owner' is set
      if (owner === undefined || owner === null) {
        throw new Error("Missing the required parameter 'owner' when calling listAgentNames");
      }


      var pathParams = {
        'owner': owner
      };
      var queryParams = {
        'offset': opts['offset'],
        'limit': opts['limit'],
        'sort': opts['sort'],
        'query': opts['query'],
      };
      var collectionQueryParams = {
      };
      var headerParams = {
      };
      var formParams = {
      };

      var authNames = ['ApiKey'];
      var contentTypes = ['application/json'];
      var accepts = ['application/json'];
      var returnType = V1ListAgentsResponse;

      return this.apiClient.callApi(
        '/api/v1/orgs/{owner}/agents/names', 'GET',
        pathParams, queryParams, collectionQueryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, callback
      );
    }

    /**
     * Callback function to receive the result of the listAgents operation.
     * @callback module:api/AgentsV1Api~listAgentsCallback
     * @param {String} error Error message, if any.
     * @param {module:model/V1ListAgentsResponse} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * List archived runs for user
     * @param {String} owner Owner of the namespace
     * @param {Object} opts Optional parameters
     * @param {Number} opts.offset Pagination offset.
     * @param {Number} opts.limit Limit size.
     * @param {String} opts.sort Sort to order the search.
     * @param {String} opts.query Query filter the search search.
     * @param {module:api/AgentsV1Api~listAgentsCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/V1ListAgentsResponse}
     */
    this.listAgents = function(owner, opts, callback) {
      opts = opts || {};
      var postBody = null;

      // verify the required parameter 'owner' is set
      if (owner === undefined || owner === null) {
        throw new Error("Missing the required parameter 'owner' when calling listAgents");
      }


      var pathParams = {
        'owner': owner
      };
      var queryParams = {
        'offset': opts['offset'],
        'limit': opts['limit'],
        'sort': opts['sort'],
        'query': opts['query'],
      };
      var collectionQueryParams = {
      };
      var headerParams = {
      };
      var formParams = {
      };

      var authNames = ['ApiKey'];
      var contentTypes = ['application/json'];
      var accepts = ['application/json'];
      var returnType = V1ListAgentsResponse;

      return this.apiClient.callApi(
        '/api/v1/orgs/{owner}/agents', 'GET',
        pathParams, queryParams, collectionQueryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, callback
      );
    }

    /**
     * Callback function to receive the result of the patchAgent operation.
     * @callback module:api/AgentsV1Api~patchAgentCallback
     * @param {String} error Error message, if any.
     * @param {module:model/V1Agent} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Update run
     * @param {String} owner Owner of the namespace
     * @param {String} agent_uuid UUID
     * @param {module:model/V1Agent} body Agent body
     * @param {module:api/AgentsV1Api~patchAgentCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/V1Agent}
     */
    this.patchAgent = function(owner, agent_uuid, body, callback) {
      var postBody = body;

      // verify the required parameter 'owner' is set
      if (owner === undefined || owner === null) {
        throw new Error("Missing the required parameter 'owner' when calling patchAgent");
      }

      // verify the required parameter 'agent_uuid' is set
      if (agent_uuid === undefined || agent_uuid === null) {
        throw new Error("Missing the required parameter 'agent_uuid' when calling patchAgent");
      }

      // verify the required parameter 'body' is set
      if (body === undefined || body === null) {
        throw new Error("Missing the required parameter 'body' when calling patchAgent");
      }


      var pathParams = {
        'owner': owner,
        'agent.uuid': agent_uuid
      };
      var queryParams = {
      };
      var collectionQueryParams = {
      };
      var headerParams = {
      };
      var formParams = {
      };

      var authNames = ['ApiKey'];
      var contentTypes = ['application/json'];
      var accepts = ['application/json'];
      var returnType = V1Agent;

      return this.apiClient.callApi(
        '/api/v1/orgs/{owner}/agents/{agent.uuid}', 'PATCH',
        pathParams, queryParams, collectionQueryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, callback
      );
    }

    /**
     * Callback function to receive the result of the updateAgent operation.
     * @callback module:api/AgentsV1Api~updateAgentCallback
     * @param {String} error Error message, if any.
     * @param {module:model/V1Agent} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Get run
     * @param {String} owner Owner of the namespace
     * @param {String} agent_uuid UUID
     * @param {module:model/V1Agent} body Agent body
     * @param {module:api/AgentsV1Api~updateAgentCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/V1Agent}
     */
    this.updateAgent = function(owner, agent_uuid, body, callback) {
      var postBody = body;

      // verify the required parameter 'owner' is set
      if (owner === undefined || owner === null) {
        throw new Error("Missing the required parameter 'owner' when calling updateAgent");
      }

      // verify the required parameter 'agent_uuid' is set
      if (agent_uuid === undefined || agent_uuid === null) {
        throw new Error("Missing the required parameter 'agent_uuid' when calling updateAgent");
      }

      // verify the required parameter 'body' is set
      if (body === undefined || body === null) {
        throw new Error("Missing the required parameter 'body' when calling updateAgent");
      }


      var pathParams = {
        'owner': owner,
        'agent.uuid': agent_uuid
      };
      var queryParams = {
      };
      var collectionQueryParams = {
      };
      var headerParams = {
      };
      var formParams = {
      };

      var authNames = ['ApiKey'];
      var contentTypes = ['application/json'];
      var accepts = ['application/json'];
      var returnType = V1Agent;

      return this.apiClient.callApi(
        '/api/v1/orgs/{owner}/agents/{agent.uuid}', 'PUT',
        pathParams, queryParams, collectionQueryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, callback
      );
    }
  };

  return exports;
}));
