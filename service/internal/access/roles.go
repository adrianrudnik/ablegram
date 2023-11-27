package access

// AdminRole is given to users that authenticate with either the OTP or the configured master password.
const AdminRole = "admin"

// GuestRole is given to anyone saying hello via POST /api/auth.
const GuestRole = "guest"
