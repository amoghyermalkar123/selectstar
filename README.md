## SelectStar
SelectStar is a minimalistic personal query management app
It helps you -
  * Add a query
  * Delete a query
  * View All Queries
  * View particular query
  * Search Queries by name (exact/ appropriate keyword match)
  * Copy Query

SelectStar is not a redash replacement. 
It is meant to be your personal tool to manage your plethora of queries (draft/ otherwise) that you want to dump in one place and retrieve anytime you want very quickly and easily.

## Installation -
- Run - `git clone https://github.com/amoghyermalkar123/selectstar`
- Build image - `sudo docker build -t ss .`
- Run image - `sudo docker run -p 8091:8091 --name ss -d ss`
- Stop image - `sudo docker stop ss`
- Restart image - `sudo docker start ss`