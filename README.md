# gosub

## What is this about?
This project is intended to create a mailing service for blog subscriptions.
It follows the simple routine:
1) request subscription
2) generate new entry
3) send verification link
4) when verified, keep in database & delete all unverified emails after 24 hours
5) MAYBE FOR THE FUTURE: send emails on rss feed changes

## Uses:
- https://github.com/mongodb/mongo-go-driver
- [mongodb](https://www.mongodb.com/)
- https://github.com/joho/godotenv
