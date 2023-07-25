package stubs

// MessagesSendResponse response from the /v1/messages/send endpoint
func MessagesSendResponse() []byte {
	return []byte(`
		{
			"data": {
				"can_be_polled": false,
				"contact": "+18005550100",
				"content": "This is a sample text message",
				"created_at": "2022-06-05T14:26:02.302718+03:00",
				"delivered_at": "2022-06-05T14:26:09.527976+03:00",
				"expired_at": "2022-06-05T14:26:09.527976+03:00",
				"failed_at": "2022-06-05T14:26:09.527976+03:00",
				"failure_reason": null,
				"id": "32343a19-da5e-4b1b-a767-3298a73703cb",
				"last_attempted_at": "2022-06-05T14:26:09.527976+03:00",
				"max_send_attempts": 1,
				"order_timestamp": "2022-06-05T14:26:09.527976+03:00",
				"owner": "+18005550199",
				"received_at": "2022-06-05T14:26:09.527976+03:00",
				"request_id": "153554b5-ae44-44a0-8f4f-7bbac5657ad4",
				"request_received_at": "2022-06-05T14:26:01.520828+03:00",
				"scheduled_at": "2022-06-05T14:26:09.527976+03:00",
				"send_attempt_count": 0,
				"send_time": 133414,
				"sent_at": "2022-06-05T14:26:09.527976+03:00",
				"sim": "SIM1",
				"status": "pending",
				"type": "mobile-terminated",
				"updated_at": "2022-06-05T14:26:10.303278+03:00",
				"user_id": "WB7DRDWrJZRGbYrv2CKGkqbzvqdC"
			},
			"message": "item created successfully",
			"status": "success"
	}
`)
}

// MessagesSendErrorResponse internal error response
func MessagesSendErrorResponse() []byte {
	return []byte(`
		{
			"message": "We ran into an internal error while handling the request.",
			"status": "error"
		}
`)
}
