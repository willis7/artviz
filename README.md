# Artviz

This simple command line tool connects to your artifactory instance using the REST API and generates a graphviz representation using the DOT language.

Usage:

    artviz --url <artifactory base url>

e.g

    artviz --url http://localhost:8081/artifactory

## Development

Spin up a local instance of Artifactory using

    make compose-up

You will need to apply for an evaluation token. Enter that now.

Access the application on:
    
    localhost:8081

login with `un/password`

    admin/password


Get list of repos:

    $ curl http://localhost/artifactory/api/repositories


## Example usage

Pass the output to Graphviz:

    artviz --url http://localhost:8081/artifactory | dot -Tpng > artifactory.png  