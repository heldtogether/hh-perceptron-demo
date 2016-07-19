# hh-perceptron-demo

Demo of a linear classifier that learns boolean operations.

## Installation

To install, you will need to install Go, see [installation instructions](https://golang.org/doc/install).

In your GOPATH, run

	go get github.com/heldtogether/hh-perceptron-demo

to download the app. Then run

	cd $GOPATH/src/github.com/heldtogether/hh-perceptron-demo
	go get -t ./...

to download all dependencies.

Finally, install the app with

	go install github.com/heldtogether/hh-perceptron-demo


## Usage

Run the app using

	hh-perceptron-demo [-generator="{and|or|nand}"] [-rounds=100]

Available flags:

`-generator`: Type of training data generator. Available generators are "and", "or" and "nand". The default is "and".

`-rounds`: Number of rounds of training to perform. The default is 100.
