# goTor

[TorBot](https://github.com/DedSecInside/TorBoT) is an OSINT tool that allows you to scrape Tor sites and pull relevant information. It's currently ran using CLI which can be tedious for some users and can be offputting. To counter this issue, I'm designing a graphical representation of TorBot that behaves similarly but will be much easier to use. I'm rewriting TorBot in Golang instead of Python so I hope to see performance gains as well.

The only method that works currently is the `-l` argument from TorBot which lists all the associated links of a website. Before the repo will be near production ready, the rest of the arguments must be implemented as well.

## Getting Started

### Development
  - `npm run start` to build code and launch test page locally
  - `go run server.go` to start server that provides API
