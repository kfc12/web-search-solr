# web-search-solr

This is a web application that provides a search box to search for documents in the DI wiki team pages.  The wiki pages have been indexed into solr.  When a match is found from text in the search box, a url is returned along with the highlighted text that the document is matching on. 


For use in Docker, only the Dockerfile is needed.  The docker-compose file is used to create the necessary yaml files, using the 'kompose convert' command in the same directory as the docker-compose file.  The convert will generate the deployment yaml and the service yaml. 

If using with kubernetes, only the yaml files are necessary.

The following commands are given if any of the files need editing.


# Docker Compose

  $ docker build -t <nameyourimage> .
	

# Docker Compose

  $ compose up
	

# Convert the docker-compose to yaml files


  $ kompose convert

# To run the docker-compose in kubernetes without converting to yaml files

  $ kompose convert
