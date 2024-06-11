package httpsms

// EventTypeMessagePhoneReceived is emitted when a new message is received by a mobile phone
const EventTypeMessagePhoneReceived = "message.phone.received"

// EventTypeMessagePhoneSent is emitted when the phone sends a message
const EventTypeMessagePhoneSent = "message.phone.sent"

// EventTypeMessagePhoneDelivered is emitted when the phone delivers a message
const EventTypeMessagePhoneDelivered = "message.phone.delivered"

// EventTypeMessageSendFailed is emitted when the phone could not send
const EventTypeMessageSendFailed = "message.send.failed"

// EventTypeMessageSendExpired is emitted when the phone a message expires
const EventTypeMessageSendExpired = "message.send.expired"

// EventTypePhoneHeartbeatOffline is emitted when the phone is missing a heartbeat
const EventTypePhoneHeartbeatOffline = "phone.heartbeat.offline"

// EventTypePhoneHeartbeatOnline is emitted when the phone is missing a heartbeat
const EventTypePhoneHeartbeatOnline = "phone.heartbeat.online"

// EventTypeMessageCallMissed is emitted when a new message is sent
const EventTypeMessageCallMissed = "message.call.missed"
