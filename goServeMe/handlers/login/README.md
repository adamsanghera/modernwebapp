# Login µService

## Features

Note, that this service is designed after the Command Query Responsibilty Segregation [CQRS] pattern.  This means that there are two categories of messages parsed by this service.  `Commands`, which change the 'internal' state of the service, and `Queries`, which provide surgical glimpses of that same state.

So, what `commands` would you expect a login service to perform?

1. Create multiple, distinct "categories" of authorization. --> `Registration`
2. Issue tokens, verifiable badges of memers --> `Capability-provision`
3. Verify a token's membership --> `Validation`

What `queries` would you expect a login service to answer?

1. 

## Methodology

Let's illustrate methodology with a simple example.

### Simplest approach

1. User `U` registers for access to category `C`, with a password `p`, which the service hashes with a randomly-generated salt.  The salt and hash are stored.
2. User `U` requests a token representing membership in `C`, with password `p`.
3. This password `p` is validated, and a randomly-generated token `t` is issued to `U`.
4. `U` can use this token to validate her other requests, until the token expires.
5. At any time before expiry, `U` can ask for the old token to be exchanged for a new one.
6. If a request is sent with an expired token, a `challenge` message will be sent back.

### Flaws with this approach

1. If `U` leaks their token, they can be impersonated by Eve.
2. If `U` leaks their token, Eve can make exchange requests, effectively DoS-ing `U`.
3. If Eve can look into the Service, she knows who is logged into what at any given time.
4. If Eve can look into the Service, she can intercept passwords.
5. If Eve can look into the Service, she can intercept tokens, leading to (1) and (2).

### Solutions to these flaws

1. (1) (2) (5) can be neutered by decreasing the value of a given token.  Generate a new token after every single request, and exchange them automatically by piggy-backing on the response.  This has the big downside of effectively making concurrency impossible.
2. (4) can be prevented by moving hashing to the client side.  With this scheme, only hashed passwords and salts are sent over the wire.  This doesn't have any huge downsides, aside from requiring a small amount of computation once every session.
3. (3) can be prevented by random-salting and hashing the username on the client side.  The minimal-obfuscation database row would look something like [usernameSalt, hashedName, passSalt, hashedPass]. Ah wait, except the user would have no way of retrieving the username salt.  So no client-side username salt.  Well, we could salt the received hash and hash it again I guess.  That's a bit extra but why not.

With these changes, a typical handshake would look like this:

1.  User wants to register, so she makes a username X and private password P.  She makes one random salt pSalt, and then salts P.  With a hash function, she generates H(X) and H(P+pSalt).
2.  User sends [H(X), H(P+pSalt), pSalt] to the server.
2.  The server stores this information, replies with success message
3.  The user indicates that they would like to login as user X.  Under the covers, X is hashed to be H(X).  H(X) is sent to the server.
4.  H(X) is looked up, and pSalt is sent back.
5.  User types in password, and client sends back H(X) plus H(pass + pSalt).
6.  Server looks up H(X), and compares H(pass + pSalt).
7.  User is right! So the server generates a token T, associates it with H(X) and sends it back.
8.  User wants to do something that needs validation, sends H(X) and T to server.
9.  Server validates T for H(X), performs the action, generates a new token TT, and piggy-backs it on the reply.

Now, the only sucky part about this is that it requires client applications to hash things locally with a particular algorithm and parameters.

### But wait! There's more...

After login, Eve can only pick up on [H(X) and T].  These are good to impersonate X for at least one shot.  If Eve can spoof their MAC and IP Address to be X's, this impersonation will be impossible to detect.  Eve still doesn't know X's password or even username, but she's in.  X will be left with an old token.  X's next request will fail validation.  Given this scenario's possibility, it's best to forcefully log X out when an invalid token is received.  In this scenario, the worst thing that Eve can do is DoS X for as long as she can intercept X's tokens and ping the server faster than Eve can.  

To transfer the H(X) and tokens securely, we'd need public key crypto.  The server would have to encrypt the message with X's public key.  That would require the user to use HTTPS or TLS or SSH.  Not that big of an overhead, but still an extra complication.  So, we have...

1. User makes a username X and password P.  She makes random pSalt.  Sends [ H(X) , H(P+pSalt) , pSalt ] to the server over encrypted channel.
2. Server stores [ H(X) , H(P+pSalt) , pSalt ].
3. User wants to log in as user X.  Requests pSalt with [ H(X) ] over encrypted channel.
4. Server responds with [ pSalt ] over encrypted channel.
5. User creates H(P+pSalt), sends [ H(P+pSalt) ] over encrypted channel.
6. Server sees that the hash matches.  Server creates token T, and stores [ H(X) , T ].  Server sends [ H(X) , T ] over encrypted channel.
7. User sends request, with footer [ H(X) , T ] over encrypted channel.

So now we're sending hashed data over HTTPS.  Most modern browsers support HTTPS, so nbd.  However, requiring developers to mix and match their data to follow our package's hashing scheme is a deal-breaker of complexity.  We'd have to ship a JS package that does the hashing for them.  Maybe this https://github.com/tonyg/js-scrypt or https://www.npmjs.com/package/bcryptjs.

# Roadmap

## Done
1. Get a basic registration/provisioning/validation system set up.  No fancy hashing.
2. Hash at the server level.  No plaintext passwords are persisted.

## Up and coming
3. Hash at the server level.  No plaintext passwords or usernames are persisted.
4. Hash at the browser level.  No plaintext passwords or usernames are transmitted.

### Interesting feature

Maintain a verifiable history of token validations, blockchain style.

Who validated what, and when?