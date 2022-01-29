### About
A recursive, mirroring web crawler that retrieves child links.


### Prerequisites
Before running the application, ensure you have installed and configured
the following tools

- Go v1.7+. Read more at https://go.dev/doc/install


### Usage
1. In a terminal navigate to the root directory of the project then run the following commands in a shell:

   Build the application
   ```
     $ go build
   ```

   Run the application
   ```
     $ ./webcrawler --dir=path/to/destination [URL]
   ```

   or 

    ```
     $ ./webcrawler -d path/to/destination [URL]
   ```

   Where
-  --dir option represent the absolute path to the destination directory
    where the page at the URL will be downloaded.

-  [URL] link for the website. for example:

    > ./webcrawler --dir=/Users/TonyM/ProjectPlato/downloads https://start.url/abc