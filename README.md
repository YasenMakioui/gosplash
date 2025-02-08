# gosplash

![logo](assets/gosplash.webp)

A golang application used to share secrets and files securely using a web UI, HTTP request or even CLI.

These secrets are one time secrets and will be deleted once readed.

This project is heavily inspired by onetimesecret and transfersh.

Case scenario: 

A user needs a database dump to perform a task. We need to send him the database dump securely so we upload the file to gosplash.

Uploading the file returns the URL of the file we want to share and the secret URL.

We can pass both on the same channel, email or chat but for even more security we can send each URL on a different channel.


## Authors

- [@yelmakioui](https://www.github.com/YasenMakioui)

## Acknowledgements

- [onetimesecret](https://github.com/onetimesecret/onetimesecret)
- [transfersh](https://github.com/dutchcoders/transfer.sh)

## Database Model

The databsae model aims to be as simple as posible to focus on the main purpose of the app.

![datamodel](assets/gosplash-database.png)

