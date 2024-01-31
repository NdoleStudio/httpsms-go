package stubs

// MessageThreadIndexResponse response from the /v1/message-threads endpoint
func MessageThreadIndexResponse() []byte {
	return []byte(`
{
  "data": [
    {
      "color": "indigo",
      "contact": "+18005550100",
      "created_at": "2022-06-05T14:26:09.527976+03:00",
      "id": "32343a19-da5e-4b1b-a767-3298a73703ca",
      "is_archived": false,
      "last_message_content": "This is a sample message content",
      "last_message_id": "32343a19-da5e-4b1b-a767-3298a73703ca",
      "order_timestamp": "2022-06-05T14:26:09.527976+03:00",
      "owner": "+18005550199",
      "status": "PENDING",
      "updated_at": "2022-06-05T14:26:09.527976+03:00",
      "user_id": "WB7DRDWrJZRGbYrv2CKGkqbzvqdC"
    }
  ],
  "message": "item created successfully",
  "status": "success"
}
`)
}
