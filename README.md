# Email (Magic Link) authentication app

This simple app generates a JWT token, appends it to a user-provided URL, and sends it to the email provided by the user.

The app also provides the ability to validate any issued tokens. 

Note that the app is stateless and does not store any information related to issued tokens.


Intended use:

- user lands on a protected frontend route
- frontend checks if cookies for the JWT token
- if found, a request is sent to this app to validate the token
- if successful, access is granted
- if unsuccessful, the app generates a new token and sends it to the user's email


The app only has one endpoint `/token` and only accepts GET requests

When the token needs to be validated, provide it in the `validate` query string (`/token?validate=<token>`).

When new links need to be generated, provide a user's email and protected URL in the query string like so `/token?email=user@example.com&protected-url=https://example.com/protected`

If the request is executed successfully, the server returns 200 status code and the following JSON:

```go
{
    Ok bool
    Message string
}
```

Otherwise it returns status code 400 and a message in plain text or JSON format.

The app relies on the environment variables that are read from the `.env` file that you must create based on `template.env`. It was tested with Gmail but should work with other providers. For Gmail (and potentially other providers) the password will not work and you must generate an app password.

Note: JWT token is valid for 2 weeks, so you may want to change that in `pkg/jwt`.