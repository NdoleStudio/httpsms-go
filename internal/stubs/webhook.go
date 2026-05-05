package stubs

// WebhookStoreResponse response from the POST /v1/webhooks endpoint
func WebhookStoreResponse() []byte {
	return []byte(`
{
    "data": {
        "id": "b2c3d4e5-f6a7-8901-bcde-f12345678901",
        "user_id": "hT5V2CmN5bbG81glMLmosxPV9Np2",
        "url": "https://example.com/webhook",
        "signing_key": "whsec_test_signing_key",
        "phone_numbers": ["+18005550199"],
        "events": ["message.phone.received", "message.phone.sent"],
        "created_at": "2024-01-21T13:07:56.203538Z",
        "updated_at": "2024-01-21T13:07:56.203538Z"
    },
    "message": "webhook created successfully",
    "status": "success"
}
`)
}
