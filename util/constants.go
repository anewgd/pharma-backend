package util

type userIDString string

var UserID userIDString = "user_id"

type roleString string

const Role roleString = "ROLE"

const AuthorizationHeaderKey = "authorization"
const AuthorizationTypeBearer = "bearer"
const AuthorizationPayloadKey = "authorization_payload"
