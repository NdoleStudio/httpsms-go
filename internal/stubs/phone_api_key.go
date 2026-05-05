package stubs

// PhoneAPIKeyStoreResponse response from the POST /v1/phone-api-keys/ endpoint
func PhoneAPIKeyStoreResponse() []byte {
	return []byte(`
{
    "data": {
        "id": "a1b2c3d4-e5f6-7890-abcd-ef1234567890",
        "name": "test-key",
        "user_id": "hT5V2CmN5bbG81glMLmosxPV9Np2",
        "user_email": "test@example.com",
        "phone_numbers": ["+18005550199"],
        "phone_ids": ["9d484671-cac2-41de-9171-d9d2c1835d7b"],
        "api_key": "pk_test_1234567890abcdef",
        "created_at": "2024-01-21T13:07:56.203538Z",
        "updated_at": "2024-01-21T13:07:56.203538Z"
    },
    "message": "phone API key created successfully",
    "status": "success"
}
`)
}
