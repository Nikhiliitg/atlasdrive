What you should add (high ROI)

Instead of fake microservices, add one real distributed aspect:

Option A: Async processing
Upload triggers events
Worker processes chunks
Retries + idempotency

Option B: Failure simulation
Kill worker mid-upload
Resume logic
Consistency guarantees

Option C: Split ONE service
Keep monolith
Extract only file-service
Document the migration
This shows maturity.