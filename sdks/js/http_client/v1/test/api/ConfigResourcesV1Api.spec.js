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
    // AMD.
    define(['expect.js', '../../src/index'], factory);
  } else if (typeof module === 'object' && module.exports) {
    // CommonJS-like environments that support module.exports, like Node.
    factory(require('expect.js'), require('../../src/index'));
  } else {
    // Browser globals (root is window)
    factory(root.expect, root.PolyaxonSdk);
  }
}(this, function(expect, PolyaxonSdk) {
  'use strict';

  var instance;

  beforeEach(function() {
    instance = new PolyaxonSdk.ConfigResourcesV1Api();
  });

  describe('(package)', function() {
    describe('ConfigResourcesV1Api', function() {
      describe('createConfigResource', function() {
        it('should call createConfigResource successfully', function(done) {
          // TODO: uncomment, update parameter values for createConfigResource call and complete the assertions
          /*
          var owner = "owner_example";
          var body = new PolyaxonSdk.V1ConfigResource();
          body.uuid = "";
          body.name = "";
          body.description = "";
          body.tags = [""];
          body.created_at = new Date();
          body.updated_at = new Date();
          body.frozen = false;
          body.disabled = false;
          body.deleted = false;
          body.kind = new PolyaxonSdk.V1ConfigResourceKind();
          body.schema = new PolyaxonSdk.V1ConfigResourceSchema();
          body.schema.ref = "";
          body.schema.mount_path = "";
          body.schema.items = [""];

          instance.createConfigResource(owner, body, function(error, data, response) {
            if (error) {
              done(error);
              return;
            }
            // TODO: update response assertions
            expect(data).to.be.a(PolyaxonSdk.V1ConfigResource);
            expect(data.uuid).to.be.a('string');
            expect(data.uuid).to.be("");
            expect(data.name).to.be.a('string');
            expect(data.name).to.be("");
            expect(data.description).to.be.a('string');
            expect(data.description).to.be("");
            {
              let dataCtr = data.tags;
              expect(dataCtr).to.be.an(Array);
              expect(dataCtr).to.not.be.empty();
              for (let p in dataCtr) {
                let data = dataCtr[p];
                expect(data).to.be.a('string');
                expect(data).to.be("");
              }
            }
            expect(data.created_at).to.be.a(Date);
            expect(data.created_at).to.be(new Date());
            expect(data.updated_at).to.be.a(Date);
            expect(data.updated_at).to.be(new Date());
            expect(data.frozen).to.be.a('boolean');
            expect(data.frozen).to.be(false);
            expect(data.disabled).to.be.a('boolean');
            expect(data.disabled).to.be(false);
            expect(data.deleted).to.be.a('boolean');
            expect(data.deleted).to.be(false);
            expect(data.kind).to.be.a(PolyaxonSdk.V1ConfigResourceKind);
                expect(data.schema).to.be.a(PolyaxonSdk.V1ConfigResourceSchema);
                  expect(data.schema.ref).to.be.a('string');
              expect(data.schema.ref).to.be("");
              expect(data.schema.mount_path).to.be.a('string');
              expect(data.schema.mount_path).to.be("");
              {
                let dataCtr = data.schema.items;
                expect(dataCtr).to.be.an(Array);
                expect(dataCtr).to.not.be.empty();
                for (let p in dataCtr) {
                  let data = dataCtr[p];
                  expect(data).to.be.a('string');
                  expect(data).to.be("");
                }
              }

            done();
          });
          */
          // TODO: uncomment and complete method invocation above, then delete this line and the next:
          done();
        });
      });
      describe('deleteConfigResource', function() {
        it('should call deleteConfigResource successfully', function(done) {
          // TODO: uncomment, update parameter values for deleteConfigResource call
          /*
          var owner = "owner_example";
          var uuid = "uuid_example";

          instance.deleteConfigResource(owner, uuid, function(error, data, response) {
            if (error) {
              done(error);
              return;
            }

            done();
          });
          */
          // TODO: uncomment and complete method invocation above, then delete this line and the next:
          done();
        });
      });
      describe('getConfigResource', function() {
        it('should call getConfigResource successfully', function(done) {
          // TODO: uncomment, update parameter values for getConfigResource call and complete the assertions
          /*
          var owner = "owner_example";
          var uuid = "uuid_example";

          instance.getConfigResource(owner, uuid, function(error, data, response) {
            if (error) {
              done(error);
              return;
            }
            // TODO: update response assertions
            expect(data).to.be.a(PolyaxonSdk.V1ConfigResource);
            expect(data.uuid).to.be.a('string');
            expect(data.uuid).to.be("");
            expect(data.name).to.be.a('string');
            expect(data.name).to.be("");
            expect(data.description).to.be.a('string');
            expect(data.description).to.be("");
            {
              let dataCtr = data.tags;
              expect(dataCtr).to.be.an(Array);
              expect(dataCtr).to.not.be.empty();
              for (let p in dataCtr) {
                let data = dataCtr[p];
                expect(data).to.be.a('string');
                expect(data).to.be("");
              }
            }
            expect(data.created_at).to.be.a(Date);
            expect(data.created_at).to.be(new Date());
            expect(data.updated_at).to.be.a(Date);
            expect(data.updated_at).to.be(new Date());
            expect(data.frozen).to.be.a('boolean');
            expect(data.frozen).to.be(false);
            expect(data.disabled).to.be.a('boolean');
            expect(data.disabled).to.be(false);
            expect(data.deleted).to.be.a('boolean');
            expect(data.deleted).to.be(false);
            expect(data.kind).to.be.a(PolyaxonSdk.V1ConfigResourceKind);
                expect(data.schema).to.be.a(PolyaxonSdk.V1ConfigResourceSchema);
                  expect(data.schema.ref).to.be.a('string');
              expect(data.schema.ref).to.be("");
              expect(data.schema.mount_path).to.be.a('string');
              expect(data.schema.mount_path).to.be("");
              {
                let dataCtr = data.schema.items;
                expect(dataCtr).to.be.an(Array);
                expect(dataCtr).to.not.be.empty();
                for (let p in dataCtr) {
                  let data = dataCtr[p];
                  expect(data).to.be.a('string');
                  expect(data).to.be("");
                }
              }

            done();
          });
          */
          // TODO: uncomment and complete method invocation above, then delete this line and the next:
          done();
        });
      });
      describe('listConfigResourceNames', function() {
        it('should call listConfigResourceNames successfully', function(done) {
          // TODO: uncomment, update parameter values for listConfigResourceNames call and complete the assertions
          /*
          var owner = "owner_example";
          var opts = {};
          opts.offset = 56;
          opts.limit = 56;
          opts.sort = "sort_example";
          opts.query = "query_example";

          instance.listConfigResourceNames(owner, opts, function(error, data, response) {
            if (error) {
              done(error);
              return;
            }
            // TODO: update response assertions
            expect(data).to.be.a(PolyaxonSdk.V1ListConfigResourcesResponse);
            expect(data.count).to.be.a('number');
            expect(data.count).to.be(0);
            {
              let dataCtr = data.results;
              expect(dataCtr).to.be.an(Array);
              expect(dataCtr).to.not.be.empty();
              for (let p in dataCtr) {
                let data = dataCtr[p];
                expect(data).to.be.a(PolyaxonSdk.V1ConfigResource);
                expect(data.uuid).to.be.a('string');
                expect(data.uuid).to.be("");
                expect(data.name).to.be.a('string');
                expect(data.name).to.be("");
                expect(data.description).to.be.a('string');
                expect(data.description).to.be("");
                {
                  let dataCtr = data.tags;
                  expect(dataCtr).to.be.an(Array);
                  expect(dataCtr).to.not.be.empty();
                  for (let p in dataCtr) {
                    let data = dataCtr[p];
                    expect(data).to.be.a('string');
                    expect(data).to.be("");
                  }
                }
                expect(data.created_at).to.be.a(Date);
                expect(data.created_at).to.be(new Date());
                expect(data.updated_at).to.be.a(Date);
                expect(data.updated_at).to.be(new Date());
                expect(data.frozen).to.be.a('boolean');
                expect(data.frozen).to.be(false);
                expect(data.disabled).to.be.a('boolean');
                expect(data.disabled).to.be(false);
                expect(data.deleted).to.be.a('boolean');
                expect(data.deleted).to.be(false);
                expect(data.kind).to.be.a(PolyaxonSdk.V1ConfigResourceKind);
                    expect(data.schema).to.be.a(PolyaxonSdk.V1ConfigResourceSchema);
                      expect(data.schema.ref).to.be.a('string');
                  expect(data.schema.ref).to.be("");
                  expect(data.schema.mount_path).to.be.a('string');
                  expect(data.schema.mount_path).to.be("");
                  {
                    let dataCtr = data.schema.items;
                    expect(dataCtr).to.be.an(Array);
                    expect(dataCtr).to.not.be.empty();
                    for (let p in dataCtr) {
                      let data = dataCtr[p];
                      expect(data).to.be.a('string');
                      expect(data).to.be("");
                    }
                  }
              }
            }
            expect(data.previous).to.be.a('string');
            expect(data.previous).to.be("");
            expect(data.next).to.be.a('string');
            expect(data.next).to.be("");

            done();
          });
          */
          // TODO: uncomment and complete method invocation above, then delete this line and the next:
          done();
        });
      });
      describe('listConfigResources', function() {
        it('should call listConfigResources successfully', function(done) {
          // TODO: uncomment, update parameter values for listConfigResources call and complete the assertions
          /*
          var owner = "owner_example";
          var opts = {};
          opts.offset = 56;
          opts.limit = 56;
          opts.sort = "sort_example";
          opts.query = "query_example";

          instance.listConfigResources(owner, opts, function(error, data, response) {
            if (error) {
              done(error);
              return;
            }
            // TODO: update response assertions
            expect(data).to.be.a(PolyaxonSdk.V1ListConfigResourcesResponse);
            expect(data.count).to.be.a('number');
            expect(data.count).to.be(0);
            {
              let dataCtr = data.results;
              expect(dataCtr).to.be.an(Array);
              expect(dataCtr).to.not.be.empty();
              for (let p in dataCtr) {
                let data = dataCtr[p];
                expect(data).to.be.a(PolyaxonSdk.V1ConfigResource);
                expect(data.uuid).to.be.a('string');
                expect(data.uuid).to.be("");
                expect(data.name).to.be.a('string');
                expect(data.name).to.be("");
                expect(data.description).to.be.a('string');
                expect(data.description).to.be("");
                {
                  let dataCtr = data.tags;
                  expect(dataCtr).to.be.an(Array);
                  expect(dataCtr).to.not.be.empty();
                  for (let p in dataCtr) {
                    let data = dataCtr[p];
                    expect(data).to.be.a('string');
                    expect(data).to.be("");
                  }
                }
                expect(data.created_at).to.be.a(Date);
                expect(data.created_at).to.be(new Date());
                expect(data.updated_at).to.be.a(Date);
                expect(data.updated_at).to.be(new Date());
                expect(data.frozen).to.be.a('boolean');
                expect(data.frozen).to.be(false);
                expect(data.disabled).to.be.a('boolean');
                expect(data.disabled).to.be(false);
                expect(data.deleted).to.be.a('boolean');
                expect(data.deleted).to.be(false);
                expect(data.kind).to.be.a(PolyaxonSdk.V1ConfigResourceKind);
                    expect(data.schema).to.be.a(PolyaxonSdk.V1ConfigResourceSchema);
                      expect(data.schema.ref).to.be.a('string');
                  expect(data.schema.ref).to.be("");
                  expect(data.schema.mount_path).to.be.a('string');
                  expect(data.schema.mount_path).to.be("");
                  {
                    let dataCtr = data.schema.items;
                    expect(dataCtr).to.be.an(Array);
                    expect(dataCtr).to.not.be.empty();
                    for (let p in dataCtr) {
                      let data = dataCtr[p];
                      expect(data).to.be.a('string');
                      expect(data).to.be("");
                    }
                  }
              }
            }
            expect(data.previous).to.be.a('string');
            expect(data.previous).to.be("");
            expect(data.next).to.be.a('string');
            expect(data.next).to.be("");

            done();
          });
          */
          // TODO: uncomment and complete method invocation above, then delete this line and the next:
          done();
        });
      });
      describe('patchConfigResource', function() {
        it('should call patchConfigResource successfully', function(done) {
          // TODO: uncomment, update parameter values for patchConfigResource call and complete the assertions
          /*
          var owner = "owner_example";
          var config_resource_uuid = "config_resource_uuid_example";
          var body = new PolyaxonSdk.V1ConfigResource();
          body.uuid = "";
          body.name = "";
          body.description = "";
          body.tags = [""];
          body.created_at = new Date();
          body.updated_at = new Date();
          body.frozen = false;
          body.disabled = false;
          body.deleted = false;
          body.kind = new PolyaxonSdk.V1ConfigResourceKind();
          body.schema = new PolyaxonSdk.V1ConfigResourceSchema();
          body.schema.ref = "";
          body.schema.mount_path = "";
          body.schema.items = [""];

          instance.patchConfigResource(owner, config_resource_uuid, body, function(error, data, response) {
            if (error) {
              done(error);
              return;
            }
            // TODO: update response assertions
            expect(data).to.be.a(PolyaxonSdk.V1ConfigResource);
            expect(data.uuid).to.be.a('string');
            expect(data.uuid).to.be("");
            expect(data.name).to.be.a('string');
            expect(data.name).to.be("");
            expect(data.description).to.be.a('string');
            expect(data.description).to.be("");
            {
              let dataCtr = data.tags;
              expect(dataCtr).to.be.an(Array);
              expect(dataCtr).to.not.be.empty();
              for (let p in dataCtr) {
                let data = dataCtr[p];
                expect(data).to.be.a('string');
                expect(data).to.be("");
              }
            }
            expect(data.created_at).to.be.a(Date);
            expect(data.created_at).to.be(new Date());
            expect(data.updated_at).to.be.a(Date);
            expect(data.updated_at).to.be(new Date());
            expect(data.frozen).to.be.a('boolean');
            expect(data.frozen).to.be(false);
            expect(data.disabled).to.be.a('boolean');
            expect(data.disabled).to.be(false);
            expect(data.deleted).to.be.a('boolean');
            expect(data.deleted).to.be(false);
            expect(data.kind).to.be.a(PolyaxonSdk.V1ConfigResourceKind);
                expect(data.schema).to.be.a(PolyaxonSdk.V1ConfigResourceSchema);
                  expect(data.schema.ref).to.be.a('string');
              expect(data.schema.ref).to.be("");
              expect(data.schema.mount_path).to.be.a('string');
              expect(data.schema.mount_path).to.be("");
              {
                let dataCtr = data.schema.items;
                expect(dataCtr).to.be.an(Array);
                expect(dataCtr).to.not.be.empty();
                for (let p in dataCtr) {
                  let data = dataCtr[p];
                  expect(data).to.be.a('string');
                  expect(data).to.be("");
                }
              }

            done();
          });
          */
          // TODO: uncomment and complete method invocation above, then delete this line and the next:
          done();
        });
      });
      describe('updateConfigResource', function() {
        it('should call updateConfigResource successfully', function(done) {
          // TODO: uncomment, update parameter values for updateConfigResource call and complete the assertions
          /*
          var owner = "owner_example";
          var config_resource_uuid = "config_resource_uuid_example";
          var body = new PolyaxonSdk.V1ConfigResource();
          body.uuid = "";
          body.name = "";
          body.description = "";
          body.tags = [""];
          body.created_at = new Date();
          body.updated_at = new Date();
          body.frozen = false;
          body.disabled = false;
          body.deleted = false;
          body.kind = new PolyaxonSdk.V1ConfigResourceKind();
          body.schema = new PolyaxonSdk.V1ConfigResourceSchema();
          body.schema.ref = "";
          body.schema.mount_path = "";
          body.schema.items = [""];

          instance.updateConfigResource(owner, config_resource_uuid, body, function(error, data, response) {
            if (error) {
              done(error);
              return;
            }
            // TODO: update response assertions
            expect(data).to.be.a(PolyaxonSdk.V1ConfigResource);
            expect(data.uuid).to.be.a('string');
            expect(data.uuid).to.be("");
            expect(data.name).to.be.a('string');
            expect(data.name).to.be("");
            expect(data.description).to.be.a('string');
            expect(data.description).to.be("");
            {
              let dataCtr = data.tags;
              expect(dataCtr).to.be.an(Array);
              expect(dataCtr).to.not.be.empty();
              for (let p in dataCtr) {
                let data = dataCtr[p];
                expect(data).to.be.a('string');
                expect(data).to.be("");
              }
            }
            expect(data.created_at).to.be.a(Date);
            expect(data.created_at).to.be(new Date());
            expect(data.updated_at).to.be.a(Date);
            expect(data.updated_at).to.be(new Date());
            expect(data.frozen).to.be.a('boolean');
            expect(data.frozen).to.be(false);
            expect(data.disabled).to.be.a('boolean');
            expect(data.disabled).to.be(false);
            expect(data.deleted).to.be.a('boolean');
            expect(data.deleted).to.be(false);
            expect(data.kind).to.be.a(PolyaxonSdk.V1ConfigResourceKind);
                expect(data.schema).to.be.a(PolyaxonSdk.V1ConfigResourceSchema);
                  expect(data.schema.ref).to.be.a('string');
              expect(data.schema.ref).to.be("");
              expect(data.schema.mount_path).to.be.a('string');
              expect(data.schema.mount_path).to.be("");
              {
                let dataCtr = data.schema.items;
                expect(dataCtr).to.be.an(Array);
                expect(dataCtr).to.not.be.empty();
                for (let p in dataCtr) {
                  let data = dataCtr[p];
                  expect(data).to.be.a('string');
                  expect(data).to.be("");
                }
              }

            done();
          });
          */
          // TODO: uncomment and complete method invocation above, then delete this line and the next:
          done();
        });
      });
    });
  });

}));
