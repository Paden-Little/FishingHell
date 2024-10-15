# NOTES
- If you are adding a library to a go project be sure to re-add go.sum to the 6th line (COPY command) of the docker compose and run "go mod tidy"
- Service DB user and password
- 

**Fishing Hell – Capstone Project Proposal** 

**Note to reader:**

This document is in no way comprehensive; it is meant to be a high-level overview of the content within the application system. If the reader has any questions or desires any further information about the application system the author highly recommends reading through [Fishing Hell – Comprehensive Design Document](https://docs.google.com/document/d/1jgxrtGsgM-mm7dXl3XJ4pR8-fbgXrPShh4YcmxngGvo/edit?usp=sharing). It contains in-depth coverage including, complete user stories, system design, game mechanics, mockups, etcetera. Many of this document’s footnotes will reference or directly link to the design document’s further explanation of the given topic.

**Topic:** 

A fishing game discord bot. Users can catch, sell, trade, buy, and appraise hundreds of fish\! The game is text-based and unique per discord server. Will also feature a custom web-based dashboard where users can see their items, fish, cash, etc easier.

The goal of the project is to work with complex data, and A LOT of it. I want to build a system that is efficient and can handle a large number of requests and complex data interaction.​

The project will also be strictly and fully tested. I fully believe in the power of test driven development and believe this is a great project to show the power of it. Using unit tests within each class inside the services. While doing integration tests from the outside before service (checking endpoint and message quality). Data tests to ensure SQL script reports are being generated properly. And user tests which will be sequences of tests following a sequences of actions of interactable actions within the system; A La users stories.

**Core Requirements:**

**Gameplay:**

* Fishing  
  * Users send messages to initiate fishing and can catch fish that way  
* Appraising  
  * Users can pay a fee to verify the quality of their fish and get a rough value of the fish  
* Selling  
  * Users can sell their fish at the fish market. The fish market has a complex sudo stock market setup where, depending on how much is sold previously, the rarity, quality, freshness of the fish etc etc, the per oz of fish will be variable hour to hour.  
* Trading  
  * Users can trade fish with other users to help complete collections  
* Equipments & Items  
  * Users can purchase equipment and items, different baits, hooks, lines, poles, nets, etc. Which will need to be swapped out based on where you are etc  
* Web dashboard  
  * The web dashboard is a visual representation of the given persons items, fish, equipment, money. As well as data representation of the variability of the fish market.

	**Technical:**

* Data persistence  
  * Data will be persisted and cached to ease load on the system  
* Deployed web app  
  * The web dashboard will be a deployed web app that is user specific. Users can request a link to their specific dashboard via the bot, which will then DM them the link. Then when you click the link it captures the users discord auth information. Compare that to the data you are attempting to fetch, and if it isn’t their dashboard who clicked on the link they cannot access the page.  
* Microservices  
  * All individual services are microservices and contain a postgresql database instance inside it. This is so that as the system grows I can multiply the images without issues

**Stretch Requirements:**

* Metrics and Reporting  
  * Report on data flow and strain on application system  
* Integration Testing  
  * A private discord server with multiple instances of integration-bots interacting with the application system in various ways  
* Caching

**Tech Stack:**

Below is a name of each service/component of the application followed by the technologies planned for building

* Web dashboard: Go, HTMX, Bootstrap  
* Discord Server API Flexible Message Queue: C++, D++ (Discord API Library)  
* Fisher Service: Python/Mojo, PostgreSQL and Redis  
* Selling Service: Python/Mojo, PostgreSQL and Redis  
* Trading Service: Python/Mojo, PostgreSQL and Redis  
* Equipments & Items (shop): Python/Mojo, PostgreSQL and Redis  
* Integration Testing: Python, Discord.py  
* Cloud Infrastructure: AWS (EC2, S3)

**Overview:**

The system will involve a front of stack Flexible Message Queue written from scratch. Where based on the message it will either require or not require a response based on the traffic; e.g. a write to the database does not require a specific response where a read does 

**Schedule:**

Note that the Schedule is broken down by week, then day. Assuming each day has a dedicated 4 hours of work. Each day will show what component I will work on followed by a more detailed description of what will be worked on.

**Week 1:**

1. Database Design  
2. Database \- Init SQL script  
3. Database interface \- design  
4. Database interface \- implementation  
5. Database interface \- testing

**Week 2:**

1. Flexible Message Queue \- Layer 1: direct discord API access  
2. Flexible Message Queue \- Layer 1: direct discord API access  
3. Flexible Message Queue \- Layer 1.5: transform calls to the bot into digestible objects  
4. Flexible Message Queue \- Layer 2: request \- determine reason for call, and where to direct traffic  
5. Flexible Message Queue \- Layer 2: request \- determine whether message requires return

**Week 3:**

1. Flexible Message Queue \- Layer 3: response \- respond for non returns  
2. Flexible Message Queue \- Layer 3: response \- respond from queue response  
3. Flexible Message Queue \- Testing  
4. Flexible Message Queue \- Testing  
5. Flexible Message Queue \- Testing

**Week 4:**

1. Flexible Message Queue interface \- Python/Mojo: accept calls from the Flexible message queue, and interprets them  
2. Flexible Message Queue interface \- Testing  
3. Selling Service \- implementation  
4. Selling Service \- implementation  
5. Selling Service \- testing

**Week 5:**

1. Fisher Service \- implementation  
2. Fisher Service \- implementation  
3. Fisher Service \- implementation  
4. Fisher Service \- testing  
5. Fisher Service \- testing

**Week 6:**

1. Trading Service \- implementation  
2. Trading Service \- implementation  
3. Trading Service \- testing  
4. Equipment Service \- implementation  
5. Equipement Service- implementation

**Week 7:**

1. Equipment Service \- testing  
2. Web Dashboard \- implementation  
3. Web Dashboard \- implementation  
4. Web Dashboard \- implementation  
5. Web Dashboard \- testing

**Week 8:**

1. Integration Testing \- implementation  
2. Integration Testing \- implementation  
3. Integration Testing \- implementation  
4. Integration Testing \- implementation  
5. Integration Testing \- implementation

**Week 9 \- 10:**

I am leaving week 9 through 10 blank currently, will reassess whether stretch goals will be implemented during this time or if I need more time to clean up or fully implement any prior component

Thank you\!

	
