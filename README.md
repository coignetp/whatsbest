# Whatsbest
## About
An application to rank choices with friends.  
Play now at <https://whatsbest.herokuapp.com/>

## How to use it
Clone the repository with
```
git clone git@github.com:coignetp/whatsbest.git
```
To start it locally, you can use the `docker-compose.yml` at the root of the project.
```
sudo docker-compose up -d
```
It will start both the frontend (in VueJS) on *8080* and the backend (in golang) on *8081* with its embedded database (in SQLite). You can now use it locally.

## Technical notes
### Production environment
The `Dockerfile` at the root of the project is used for production on Heroku. When pushed to heroku, the `Dockerfile` will automatically be built according to `heroku.yml` and the application will be available.
However, the production stack is different from the development stack. Indeed, the free subscriptino on Heroku allows to expose only one `Dockerfile` to be exposed, so when *Whatsbest* is built the backend serves the frontend with `fs := http.FileServer(http.Dir("./web/dist"))` and `http.Handle("/", fs)`. This is why there is a `-dev` option when starting the golang backend.

### Backend
If you use the `-dev` flag when starting the golang backend, it **delete** the last `tournament.db`.  
When you do not use this flag, the backend will consider that the frontend is already built in `backend/web/dist`.

## Sample
In order to have a sample database, there is a sample db file`backend/tournament.db.sample`. You can rename it `tournament.db` to have a tournament when the application starts. Do not forget **not** to use `-dev` flag if you already have a `tournament.db` file. The sample tournament has the id *857454736*, so you can access it at `/mash/331bb890`.