# Example Scrapli code
## High level objectives
1. Demonstrate the ability to consume input from an external source.
2. Iterate over individual commands from the external source or use it to point to another remote source.  In this example, the pre-maintenance steps were in the input file but the configuration file was a separate entity.  Either way would have worked, but I wanted to demonstrate multiple ways of doing the same thing.
## Secondary discovery
1. Learned how to pass arguements with the golang flag library.  I expected it to dynamically prompt me for input.  This library does not function as expected, it acts like a traditional command line tool and expects the arguements to be passed in at run time.
`go run main.go -username your_username -password your_password`
