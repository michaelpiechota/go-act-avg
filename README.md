go-act-avt 
=======
A small Go package that maintains an average time for an action.

Install
=======
Install the package:

> $ go get github.com/michaelpiechota/go-act-avg
	
Preparing
=======
Import the package:

	import "github.com/michaelpiechota/go-act-avg"
	
Usage
=======
(Download/update go 1.13 https://golang.org/dl/)

Included in the main directory is a sample main.go.
To run the sample main function:

>> make prep
>> go run main.go

Note: "make prep" downloads dependencies and runs tests.

Implementation Decisions and Future Considerations
=======
###Service Object Pattern
I've had great success using this type of (Ruby-ish) service pattern with API implementations in the past. 
I believe this pattern helps establish clean architecture and is a great example of the benefits of the
Single Responsibility Principle. In this case, it allows extension of package to be fairly easy by giving 
the ability to add items such as database table configurations and returning HTTP client configurations.

###Zap Logger
Zap (https://godoc.org/go.uber.org/zap) is my favorite Go logging packages and I've had great results using it in production.
Structured logging makes it easier to achieve greater insight and allows for custom logging based on package needs.
Structured logging provides many benefits when used with services such as such AWS Cloudwatch/Insights
and it makes to exporting data into Splunk or Kibana trivial due to named-fields (e.g. Error, Info, Fatal).
 
###Map As A Data Store and Mutex 
For the scope of this project, a map seemed to be the right choice for what appeared to be a k/v input to the AddAction()
func. With this comes the risk of concurrent map writes, which is solved by the use of a mutex to synchronise access to the 
map (https://golang.org/pkg/sync/#Mutex). Perhaps a better implementation would have been to use `syncmap.Map` (https://godoc.org/golang.org/x/sync/syncmap)
but further research and testing would be needed to evaluate proper implementation and performance gain.
The next critical step of this project would be to decide on a persistent data store such as DynamoDB, 
Cassandra, or Redis depending on use cases. 

###Separating Models
I find separating models within a package to a centralized location to be a handy way to increase code cleanliness and repetitiveness.
For example the Input struct that is used in the AddAction func could also be used if a SubtractAction func was added in the future.

###Makefile
I find Makefiles to be very useful. In this case, I created a "prep" action that bundles downloading dependencies and running tests. 
I've used Makefiles heavily in CI/CD deployment pipelines and have seen great results by consolidating steps and de-cluttering a UI.

Future TODO's
=======
#### Add proper data storage
#### Add tests for concurrency

