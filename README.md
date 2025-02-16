# gosplash

![logo](assets/gosplash.webp)

A golang application used to share files and secrets both privately and publicly.

We can login and send the file to another registered user and it will pop up to his homepage.

Public sharing is also supported via secure URLs that are deleted once the time expires or the limit imposed is exceeded.

For sharing the file we need to add a special token facilitated by the user or both users need to be part of the same organization.


## Authors

- [@yelmakioui](https://www.github.com/YasenMakioui)

## Acknowledgements

- [onetimesecret](https://github.com/onetimesecret/onetimesecret)
- [transfersh](https://github.com/dutchcoders/transfer.sh)

## Stack

As for the stack, only golang and postgresql are used. Since we want this app to use the typing system that golang offers and postgres for its huge community.

This project aims to limit as much as possible third party libraries.

All the backend is written using the net/http golang package, since it offers all the functionality we need.

The philosophy in this project is, if we can do it reliably with the standard library then we do it with the standard library.

## ¿Why?

The reason behind this app is the need for securely sharing data in a fast and easy way to shareholders or collegues.

The typical example I use is when we need to share a database dump for operation reasons.

I've seen things such as sharing it with Google Drive or worse, via email.

The best approach is a system hosted by the organization or by a trusted organization that allows us to share with a simple upload to the recipients we want. 

## Database Model

The databsae model aims to be as simple as posible to focus on the main purpose of the app.

![datamodel](assets/gosplash-database.png)