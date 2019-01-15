For this assignment you also have to answer a couple of questions.
There is no correct answer and none is mandatory, if you don't know just skip it.

- **What do you think of the initial project structure ?**
I have only be learning Go for 2 weeks and I am unfamiliar with the common design patterns and best practices in architecting services.
But generally there is a fair separation of layers in the project with interfaces between them. 
I am sure I will have more of an opinion about this as my knowledge and experience of Go increases.


- **What you will improve from your solution ?**
This is difficult to answer because I am writing this retrospectively. I have made some changes but mostly in adding a couple of interface methods that were not in the original specification.
One thing I noticed straight away was that while there was plan and tests in place to have a fight between 2 knight there was no http endpoint planned to trigger/ I have taken the liberty of adding an endpoint POST /fight?fighter1?Xfighter2=Y


- **For you, what are the boundaries of a service inside a micro-service architecture ?**
For me the boundary of a micro service is determined mostly by domain area and functional area. It can also have a lot to do with transactional requirements.
In a retail app it makes sense that 'user management' and 'orders' really are separate systems and could be better deployed as micro services. This gives greater resilience to change and promotes scalability. This means the 'busier' components of an app can be replicated as necessary without needing to provision more instances for services that are barely used.
Communication between the services and their respective datastores will be eventually consistent so caution must be taken when designing systems this way as a whole new type of complexity is introduced in the form of eventual consistency.
If an aspect of an app is reusable and has no dependencies on the rest of the app it probably should be a micro service. (ie. email service, accounts system)
Functional areas in an app such as Payments, Collections, Orders, Registration all instantly feel like they could be micro services because each of those areas can probably act and evolve independently as long as other services know how to send messages to it and vice versa.
The concept of the bounded-context in Domain Driven Design maps very well to the idea of micro services.


- **For you, what are the most relevant usage for SQL, NoSQL, key-value and document store ?**
Transactional databases still underpin many systems throughout the world and are are great fit for systems that demand transactional consistency. With modern object relational mappers we can avoid much of the complexity too. The downfall with relational databases is that they don't scale very well.

NoSql or document databases such as mongoDB are a great fit when a service stores unstructured, non-relational data. You can store a huge blob of JSON which conveniently works very well when working with node.js and means a single DB query could fetch all of the relevant data in one go rather than making a huge fragile SQL query with lots and lots of joins.
NoSql databases are sometimes regarded as lesser because they are schema-less and they don't support ACID transactions, however while no schema is enforced at the DB layer, we generally enforce it on the client with libraries that abstract away some of the complexity such as mongoose in nodejs. Also mongoDB does actually support ACID transactions now.

NoSQl databases such as mongoDB scale very well and easily. It makes them a perfect fit because we use message queues to enforce eventual consistency so the need for ACID transactions lessens greatly and the scalability in needed in cloud architecture. This makes NoSQL databases a perfect fit for the modern and rapidly changing micro services landscape.

Key-value stores such as redis are very versatile and can be used for many things from managing configuration between services to real time messaging datastore. While it is usually possible to configure to persist to disk they are generally in-memory databases and are optimised as such. 

Probably the best use for key value store is as a cache as its fast and perfectly suited for caching as it only holds key-value pairs.

