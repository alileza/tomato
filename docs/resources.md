# Resources

Resource are the objects that are going to be used for steps in the scenario. They are listed under the resources key in the pipeline configuration.

## HTTP Client

http client that can be used for sending http requests and comparing the responses

Initialize resource in `config.yml`:
```yaml
- name: # name of the resource
  type: http/client
  ready_check: true 
  params:
    # base url for the http client that will automatically be prepended to any route in the feature.
    base_url: # string
    # timeout for the request round-trip.
    timeout: # duration
    
```

### Actions

#### **Send**
send an http request without a request body
```gherkin
Given $resource send request to $target

```

#### **Send Body**
send an http request with a request body
```gherkin
Given $resource send request to $target with body $body
Given $resource send request to $target with payload $body

```

#### **Response Code**
check http response code
```gherkin
Given $resource response code should be $code

```

#### **Response Header**
check http response headers
```gherkin
Given $resource response header $header_name should be $header_value

```

#### **Response Body**
check response body
```gherkin
Given $resource response body should be $body

```



## HTTP Server

http server that mocks API responses

Initialize resource in `config.yml`:
```yaml
- name: # name of the resource
  type: http/server
  ready_check: true 
  params:
    # http server port to expose
    port: # number
    
```

### Actions

#### **Response**
set a response code and body for any request that comes to the http/server target
```gherkin
Given set $resource response code to $code and response body $body

```

#### **Response Path**
set a response code and body for a given path for the http/server
```gherkin
Given set $resource with path $path response code to $code and response body $body

```



## Database SQL

database driver that interacts with a sql database

Initialize resource in `config.yml`:
```yaml
- name: # name of the resource
  type: database/sql
  ready_check: true 
  params:
    # sql driver (postgres or mysql)
    driver: # string
    # sql database source name (`postgres://user:pass@host:port/dbname?sslmode=disable`)
    datasource: # string
    
```

### Actions

#### **Set**
truncates the target table and sets row results to the passed values
```gherkin
Given set $resource table $table list of content $content

```

#### **Check**
compares table content after an action
```gherkin
Given $resource table $table should look like $content

```



## Queue

messaging queue that that publishes and serves messages

Initialize resource in `config.yml`:
```yaml
- name: # name of the resource
  type: queue
  ready_check: true 
  params:
    # queue driver (rabbitmq)
    driver: # string
    # queue source dsn (`amqp://user:pass@host:port/`)
    datasource: # string
    
```

### Actions

#### **Publish**
publish a message to message queue
```gherkin
Given publish message to $resource target $target with payload $payload

```

#### **Listen**
listen for messages on a given queue. Declaration should be before the publish action

```gherkin
Given listen message from $resource target $target

```

#### **Count**
count messages for a given target. Declaration should be before the publish action
```gherkin
Given message from $resource target $target count should be $count

```

#### **Compare**
compare message payload. Declaration should be before the publish action
```gherkin
Given message from $resource target $target should look like $payload

```



## Shell

to communicate with shell command

Initialize resource in `config.yml`:
```yaml
- name: # name of the resource
  type: shell
  ready_check: true 
```

### Actions

#### **Execute**
execute shell command
```gherkin
Given $resource execute $command

```

#### **Stdout Contains**
check stdout for executed command contains a given value
```gherkin
Given $resource stdout should contains $substring

```

#### **Stdout Not Contains**
check stdout for executed command not contains a given value
```gherkin
Given $resource stdout should not contains $substring

```

#### **Stderr Contains**
check stdout for executed command contains a given value
```gherkin
Given $resource stderr should contains $substring

```

#### **Stderr Not Contains**
check stderr for executed command not contains a given value
```gherkin
Given $resource stderr should not contains $substring

```



