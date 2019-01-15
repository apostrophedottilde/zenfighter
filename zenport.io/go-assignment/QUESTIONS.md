# Questions

For this assignment you also have to answer a couple of questions.
There is no correct answer and none is mandatory, if you don't know just skip it.

 - **What do you think of the initial project structure ?**



 - **What you will improve from your solution ?**




 - **For you, what are the boundaries of a service inside a micro-service architecture ?**



 - **For you, what are the most relevant usage for SQL, NoSQL, key-value and document store ?**
Transactional databases still underpin many systems throughout the world and are are great fit for systems that demand transactional consistency. With modern object relational mappers we can avoid much of the complexity too. The downall with relational databases is that they dont scale very well.

NoSql or document databases such as mongoDB are a great fit when a service stores unstructured, non-relational data. You can store a huge blob of JSON which conveniently works very well when working with node.js and means a single db query could fetch all of the relevent data in one go rather than making a huge fragile SQL query with lots and lots of joins.
NoSql databases are sometimes regarded as lesser because they are schemaless and they dont support ACID transactions, however while no schema is enforced at the db layer, we generally enforce it on the client with libraries that abstract away some of the complexity such as mongoose in nodejs. Also mongoDB does actually support ACID transactions now.

NoSQl databases such as mongoDB scale very well and easily. It makes them a perfect fit because we use message queues to enforce eventual consistency so the need for ACID transactions lessens greatly and the scalability in needed in cloud architecture. This makes NoSQL databases a perfect fit for the modern and rapidly changing microservices landscape.

Key-value stores such as redis are very versatile and can be used for many things from managing configuration between services to real time messaging datastore. While it is usually possible to configure to persist to disk they are generally in-memory databases and are optimised as such. 

Probably the best use for key value store is as a cache as its fast and perfectly suited for caching as it only holds key-value pairs.
