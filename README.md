🧪 Test Login with Postman
🔹 Endpoint:
POST http://localhost:8080/login

🔹 Headers:
Content-Type: application/json

🔹 JSON Body:
json
Copy
Edit
{
  "phone": "123456789",
  "password": "secret123"
}
Make sure this phone/password is the same as what you registered with.

🔹 Response:
If success:

json
Copy
Edit
{
  "message": "Login successful",
  "player": {
    "id": "abcde12345",
    "name": "Alice",
    "phone": "123456789",
    "ip_address": "127.0.0.1",
    "origin_url": "http://example.com",
    "wallet": 0
  }
}
If wrong:

json
Copy
Edit
{
  "error": "Invalid password"
}



3. API Endpoints
🔸 Register Player
POST /register

Headers:

Content-Type: application/json

Origin: http://example.com

Body:

json
Copy
Edit
{
  "name": "Alice",
  "phone": "123456789",
  "password": "secret123"
}
🔸 Login Player
POST /login

Body:

json
Copy
Edit
{
  "phone": "123456789",
  "password": "secret123"
}