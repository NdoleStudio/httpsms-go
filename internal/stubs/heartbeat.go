package stubs

// HeartbeatIndexResponse response from the /v1/heartbeats endpoint
func HeartbeatIndexResponse() []byte {
	return []byte(`
{
    "data": [
        {
            "id": "9d484671-cac2-41de-9171-d9d2c1835d7b",
            "owner": "+18005550199",
            "user_id": "hT5V2CmN5bbG81glMLmosxPV9Np2",
            "charging": true,
            "timestamp": "2024-01-21T13:07:56.203538Z"
        }
    ],
    "message": "fetched 1 heartbeat",
    "status": "success"
}
`)
}
