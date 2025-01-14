# jdf - JSON Detect and Format
As the name suggests, this tool detects and formats JSON data that is piped into into it

Example:
```bash
devspace dev | jdf
```
This will take the logs from devspace and format any JSON data within them in a nice way!<br>
If you want to pretty print the JSON data stored in a file, you can do so by running the following command:
```bash
cat file.json | jdf
```

# Motivation
I was tired of dealing with the messy devspace logs, so I built this tool based on a colleague's suggestion. It might have some bugs, but it gets the job done effectively! What more can you expect from a tool created in just a couple of hours? XD

# Tests
Todo //🚧
