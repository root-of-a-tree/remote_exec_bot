# remote_exec_bot

A discord bot for triggering scripts to run remotely on servers. In progress.

## Commands
`!exec <server> <script>`

run script `<script>` over an SSH session with server `<server>`

`!list scripts`

list out the names of all scripts the bot would accept

`!list servers`

list out the names of all servers the bot can open SSH sessions with

`!describe <script> `

print the description for the script `<script>`, from the json config file where it was defined

## Configuration Files

### `servers.json`
list of json objects with the following fields:
- name: string (used for !exec command)
- ipAddr: string (used in SSH command by bot)
- user: string (user on the server that we SSH in as)
- keyPath: string (path to the key file for SSH authentication)

### `scripts.json`
list of json objects with the following fields:
- name: string
- description: string
- scriptPath: string (path to script inside the bot’s docker container)
- servers: list<string> (list of servers that support this script)

## Milestones
1. Golang hello world web app
2. Golang hello world web app inside a docker container
3. hello world discord app
4. ...then we worry about executing scripts over SSH
