Openshift V3 PostgreSQL Example Pod
=================

This is a simple example of deploying PostgreSQL 9.3.5 as a Kubernetes
Pod within Openshift V3.

There are 3 templates within this project.

standalone-pod-template.json defines the basic PostgreSQL Pod.

standalone-service-template.json defines the service to the PostgreSQL Pod.

standalone-consumer-template.json defines a Pod that uses the Service to access the PostgreSQL Pod.

images
------
There are 2 docker images in this repo.  

crunchy-node is the Docker image that is created to run PostgreSQL.

crunchy-admin is the Docker image that is the consumer.  It is a simple REST API server
written in golang that connects to the PostgreSQL database (via the Service).  You
can test it at http://localhost:12000/test.  

Each image has a Makefile you can run to build the Docker image.


example
-------
The consumer source code is found in the example directory.  You can build it using 
the included Makefile.

setpath.sh is used to set your GOPATH to build the golang consumer binary called adminapi.

other
-----

apply is a set of commands that deploy the examples to Openshift

teardown is a set of commands that will delete the deployed Openshift/Kubernetes objects.

written by Jeff McCormick (jeff.mccormick@crunchydatasolutions.com)
