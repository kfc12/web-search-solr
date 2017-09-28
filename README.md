# web-search-solr

This is a web application that provides a search box to search for documents in the DI wiki team pages.  The wiki pages have been indexed into solr.  When a match is found from text in the search box, a url is returned along with the highlighted text that the document is matching on. 


For use in Docker, only the Dockerfile is needed.  The docker-compose file is used to create the necessary yaml files, using the 'kompose convert' command in the same directory as the docker-compose file.  The convert will generate the deployment yaml and the service yaml. 

If using with kubernetes, only the yaml files are necessary. 
