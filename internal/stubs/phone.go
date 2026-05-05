package stubs

// PhoneUpsertResponse response from the PUT /v1/phones endpoint
func PhoneUpsertResponse() []byte {
	return []byte(`
{
    "data": {
        "id": "9d484671-cac2-41de-9171-d9d2c1835d7b",
        "user_id": "hT5V2CmN5bbG81glMLmosxPV9Np2",
        "phone_number": "+18005550199",
        "fcm_token": "test-fcm-token",
        "messages_per_minute": 5,
        "max_send_attempts": 3,
        "message_expiration_seconds": 3600,
        "sim": "SIM1",
        "created_at": "2024-01-21T13:07:56.203538Z",
        "updated_at": "2024-01-21T13:07:56.203538Z"
    },
    "message": "phone upserted successfully",
    "status": "success"
}
`)
}

// PhoneUpsertFCMTokenResponse response from the PUT /v1/phones/fcm-token endpoint
func PhoneUpsertFCMTokenResponse() []byte {
	return []byte(`
{
    "data": {
        "id": "9d484671-cac2-41de-9171-d9d2c1835d7b",
        "user_id": "hT5V2CmN5bbG81glMLmosxPV9Np2",
        "phone_number": "+18005550199",
        "fcm_token": "new-fcm-token",
        "messages_per_minute": 5,
        "max_send_attempts": 3,
        "message_expiration_seconds": 3600,
        "sim": "SIM1",
        "created_at": "2024-01-21T13:07:56.203538Z",
        "updated_at": "2024-01-21T13:07:56.203538Z"
    },
    "message": "phone FCM token updated successfully",
    "status": "success"
}
`)
}
