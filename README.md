# Welcome!

This is a pet project of mine, where I design and implement a Modern Web App<sup>TM</sup>.

This project accompanies a series of blog posts, which you can follow <a href=""> here </a>.

# How Do I Run this?

>Note that this process *will* change as the project evolves!!  This means that a build process that works today MIGHT NOT EXIST tomorrow!!  You've been warned!

## Requirements

1. Install Docker.  
    * This will save you from having to install Nginx.
2. Install Golang.  
    *   This requirement will eventually disappear, and will be replaced by a build server, as we only used Go to build a binary for the server.  Then you'll only need Docker.

On macs with homebrew, the above is a simple `brew install docker golang`

## Running the darn thing

1. Execute the build script. This will build the Go binary, and our two docker images
2. Execute the run script. Just a wrapper for `docker-compose up`.

# Interacting with the site
Open up your favorite web browser, and navigate to localhost:80, which is exposed in docker-compose.  You'll then see the web page in all its glory.

# Design Decisions

1.  For the go microservices, the only information that should be logged is information that is <b>functionally opaque</b> to the nearest external observer.  Functional Opacity is a relation between observer and information: <b>FO(O, I)</b>.
    1. I interpret <b>O</b> to be the Operating System on which the logging service is run.
    2. This relation covers information that is not immediately/trivially recordable.  Information that is functinoally opque can still be approximately derived by a careful observer.  For example, the OS can guess at the state of a service based on the power consumption of the CPU, or the number of threads currently assigned to its process.  In spite of this reality, the state of the service is still considered to be functionally opaque.
    3.  Therefore, a service should log [list is WIP]
        1. Its state
    4.  Therefore, a service should <em>not</em> log [this list is WIP]
        1. Timestamps – an observer can easily generate timestamps.
        2. Hostname – the OS knows its own host name
    
