# go-graphql
This is a learning project. It is a GraphQL API-only web service, with no database or concurrent processes.

### Routing
* pkg/net/http: https://golang.org/pkg/net/http/#ServeMux 
https://github.com/julienschmidt/httprouter
* Goroutines and web requests: https://blog.golang.org/context
  * https://medium.com/@cep21/how-to-correctly-use-context-context-in-go-1-7-8f2c0fafdf39

### Testing
* Standard testing package https://golang.org/pkg/testing/
* Why there aren’t any assertions in pkg/testing:
  * https://golang.org/doc/faq#assertions
  * https://smartystreets.com/blog/2015/02/go-testing-part-1-vanillla
* If you absolutely want to use assertions and mocks: https://github.com/stretchr/testify

### Debugging
* https://golang.org/doc/gdb
* https://github.com/derekparker/delve

### Logging
* Standard log package https://golang.org/pkg/log/
* https://github.com/sirupsen/logrus

### GraphQL 
* https://github.com/graphql-go/graphql
* More features: https://github.com/vektah/gqlgen
  * Guide https://gqlgen.com/getting-started/

### Other resources
* https://github.com/avelino/awesome-go
* https://github.com/astaxie/build-web-application-with-golang/blob/master/en/preface.md

### Stretch goal
* Add persistence to a database
  * https://golang.org/pkg/database/sql/#DB
  * ORM for SQL databases https://github.com/jinzhu/gorm (https://github.com/jinzhu/gorm/tree/master/dialects)
  * Drivers for Redis 
    * https://github.com/astaxie/build-web-application-with-golang/blob/master/en/05.6.md#redis
    * Most starred and recently updated https://github.com/go-redis/redis
* Add features that would benefit from concurrency
  * https://golang.org/doc/effective_go.html#concurrency
  * Basic patterns https://blog.golang.org/pipelines
  * Best practises https://www.youtube.com/watch?v=f6kdp27TYZs
  * Advanced concurrency patterns https://www.youtube.com/watch?v=QDDwwePbDtw
