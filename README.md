# wag
sWAGger - Web API Generator

## Usage
Note that WAG requires Go 1.7.
### Generating Code
Create a swagger.yml file with your [service definition](http://editor.swagger.io/#/). Note that WAG supports a [subset](https://github.com/Clever/wag#swagger-spec) of the Swagger spec.
Copy the latest `swagger.mk` from the [dev-handbook](https://github.com/Clever/dev-handbook/blob/master/make/swagger.mk).
Set up a `generate` target in your `Makefile` that will generate server and client code:

```
include swagger.mk

SWAGGER_CONFIG := swagger.yml
SWAGGER_CLIENT_NPM_PACKAGE_NAME := @clever/<servicename>
SWAGGER_CLIENT_NPM_PACKAGE_VERSION := 0.1.0
SWAGGER_CLIENT_NPM_PACKAGE_MODULE_NAME := <servicename>

validate: swagger-validate-deps
	$(call swagger-validate,$(SWAGGER_CONFIG))

generate: validate swagger-generate-go-deps swagger-generate-javascript-client-deps
	$(call swagger-generate-go,$(SWAGGER_CONFIG),$(PKG),$(PKG)/gen-go)
	$(call swagger-generate-javascript-client,$(SWAGGER_CONFIG),$(SWAGGER_CLIENT_NPM_PACKAGE_NAME),$(SWAGGER_CLIENT_NPM_PACKAGE_VERSION),$(SWAGGER_CLIENT_NPM_PACKAGE_MODULE_NAME))
```

Then generate your code:
```
make generate
```

This generates four directories. You should not have to modify any of the generated code:
- gen-go/models: contains all the definitions in your Swagger file as well as API input / output definitions
- gen-go/server: contains the router, middleware, and handler logic
- gen-go/client: contains the Go client library
- gen-js: contains the javascript client library

### Using the Go Server
To use the generated code you need to do two things:
- Implement the controller interface defined in `gen-go/server/interface.go`
- Pass the controller into the Server constructor. For example:
```
  s := server.New(myController, ":8000")
  // Serve should not return
  log.Fatal(s.Serve())
```

All interface methods on the controller take in a `context.Context` object.
This object comes with additional features for you to use in implementing your server logic:

* **Logging**. The [kayvee middleware logger](https://godoc.org/gopkg.in/Clever/kayvee-go.v4/middleware) is automatically added to the context object.
  It can be pulled out of the context object and used via the kayvee `FromContext` method:

```go
import "gopkg.in/Clever/kayvee-go.v4/logger"
...
logger.FromContext(ctx).Info(...)
```

  You should use this logger for all logging within your controller implementation.

* **Tracing**: `wag` instruments the context object with tracing-related metadata.
  This is done via [opentracing](http://opentracing.io/).
  In order for it to work, you are required to do two things:

  * Configure your `main()` function to report tracing data to [LightStep](http://lightstep.com/).
     We are testing Lightstep as a way to view tracing-related data:
```go
package main

import (
	lightstep "github.com/lightstep/lightstep-tracer-go"
	opentracing "github.com/opentracing/opentracing-go"
)

func main() {
	tags := make(map[string]interface{})
	tags[lightstep.ComponentNameKey] = "<name of the repo>"
	lightstepTracer := lightstep.NewTracer(lightstep.Options{
	    AccessToken: os.Getenv("LIGHTSTEP_ACCESS_TOKEN"),
	    Tags:        tags,
	})
	defer lightstep.FlushLightStepTracer(lightstepTracer)
	opentracing.InitGlobalTracer(lightstepTracer)
	...
}
```

  * Pass the context object to any subsequent network requests you make.
    Many client libraries accept the context object, e.g.:
    * **net/http**: If you're making HTTP requests, use the [golang.org/x/net/context/ctxhttp](https://godoc.org/golang.org/x/net/context/ctxhttp) package.
    * **wag**: All `wag`-generated clients take in a context object as the first argument.
      If your handler consumes a `wag`-generated client, then pass the context object to these client methods.

### Using the Go Client
Initialize the client with `New`
```
c := client.New("https://url_of_your_service:port")
```

Make an API call
```
books, err := c.GetBooks(ctx, GetBookByIDInput{Authors: []string{"Twain"}})
if err != nil {
  // Do something with the error
}
```

If you're using the client from another WAG-ified service you should pass in the `ctx` object you get in your server handler. Otherwise you can use `context.Background()`

### Using the Javascript Client from a Node + Express server

The Javascript client is generated by a fork of [swagger-codegen](https://github.com/clever/swagger-codegen) that adds support for tracing and retries.
It has currently only been tested in Node + Express servers.

The first step is initializing the object named `ApiClient`, and setting the URL of where its located:

```javascript
import * as yourservice from 'yourservice-js';

const apiClient = new yourservice.ApiClient();
apiClient.basePath = "https://url_of_your_service:port";
```

In a request handler, you can use the `apiClient` object as an argument to construct and call specific APIs.
These specific API objects will be named after the first `tag` you give the operation in your swagger spec.

For example, if you have an operation tagged `Infra` with an `operationId` of `healthCheck`, you would call it like this:

```javascript
app.get('/some_route', (req, res) => {
  const infraAPI = new yourservice.InfraApi(apiClient, {req}); // you must pass the request context through to the client call
  infraAPI.healthCheck().then(function(data) {
    res.send('service called successfully!');
  }, function(error) {
    console.error(error);
  });
});
```

Additionally, for the above to work you will need to set up server middleware to place tracing-related metadata on the `req` object.
We currently are testing out LightStep, a service that collects tracing data and displays it in a nice UI:

```javascript
import * as express from 'express';
import * as Tracer from 'opentracing';
import * as LightStep from 'lightstep-tracer';
import * as lodash from 'lodash';
Tracer.initGlobalTracer(LightStep.tracer({
  access_token   : process.env.LIGHTSTEP_ACCESS_TOKEN,
  component_name : 'repo-name',
}));

const app = express();

// Middleware to look for a span from inbound requests
app.use((req, res, next) => {
  const wireCtx = Tracer.extract(Tracer.FORMAT_HTTP_HEADERS, req.headers);
  req.span = Tracer.startSpan(req.url, {childOf: wireCtx});
  req.span.logEvent("request_received");

  // include trace ID in headers so that we can debug slow requests we see in
  // the browser by looking up the trace ID found in response headers
  const responseHeaders = {};
  Tracer.inject(req.span, Tracer.FORMAT_TEXT_MAP, responseHeaders);
  lodash.forOwn(responseHeaders, (value, key) => res.setHeader(key, responseHeaders[key]));

  // finalize the span when the response is completed
  res.on("finish", () => {
    req.span.logEvent("request_finished");
    req.span.finish();
  });

  next();
});
```

### Custom String Validation
We've added custom string validation for mongo-ids to avoid repeating: "^[0-9a-f]{24}$"` throughout the swagger.yml. To use it you have must:

- Change you swagger.yml file to have the `mongo-id` format. For example:
```
authorID:
        type: string
        format: mongo-id
```

- Import `github.com/Clever/wag/swagger` and call `swagger.InitCustomFormats()` in your server code.

Note that custom string validation only applies to input parameters and does not have any impact on objects defined in '#/definitions'.

Right now we do not allow user-defined custom strings, but this is something we may add if there's sufficient demand.

## Tests
```
make test
```

## Future Features

These are the features we're planning on building in the near future:
- Better logging, monitoring, and metrics
  - Logger from the Context
  - Expose pprof port by default
  - Automatically log process metrics (https://github.com/Clever/go-process-metrics)
- More server middleware
  - Panic recovery / handling (don't crash the server every time you get a panic)
- More client resiliency best practices
  - Smarter retry logic
  - Timeouts


## Swagger Spec

Currently, WAG doesn't implement the entire Swagger Spec. A couple things to keep in mind:
- All schemas should reference type definitions in /definitions. Any schemas defined in /paths will cause an error.
- All parameters and definitions should be defined in their specific path object instead of globally.
- Scheme, produces, and consumers can only be defined in the top-level swagger object, not individual operations. On the top level object the scheme must be 'http', produces must be 'application/json' and consumes must be 'application/json'

Below is a more comprehensive list of the features that aren't currently supported.

### Unsupported Features
Mime Types

Multi-File Swagger Definitions

Schema:
- host
- tags
- scheme (must be http)
- consumes
- produces
- securityDefinitions
- security

Consumes:
- produces (must be application/json)
- consumes (must be application/json)
- schemes
- security

Form parameter type
Parameter:
- file parameter type
- collectionFormat
- global definitions
- possibly the json schema requirements? (uniqueItems, multipleOf, etc...)

Schema object (all these have to be defined in /definitions and are generated by go-swagger)

Discriminators

XML Modeling

Security Objects

Response:
  - Headers
