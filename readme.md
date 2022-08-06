# Senior Backend Technical Assessment

We need you to create a simple article web service to withstand sudden bursts of
read by implementing cache.
The Article table will have at least these columns:
• id INT
• author TEXT
• title TEXT
• body TEXT
• created TIMESTAMP

Please create these HTTP endpoints:

1. [ POST ] /articles to post a new article
2. [ GET ] /articles to get a list of articles.  
   Sort of the articles by newest
   first, with optional query parameters:  
   a) query: to search keyword in article title and body  
   b) author: filter by author name

Requirements:

- Please write the solution in Golang
- Request and response format: JSON
- Cache per Article
- You can use any popular database, search engine, etc(use docker-compose
  for easier deployment)
- Please provide unit testings for the HTTP endpoints
- Big points if you implement CQRS
