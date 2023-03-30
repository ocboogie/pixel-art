# Pixel Art
A work-in-process social media platform with a feed, follower system, custom avatar creator, and content favoriting

Feed             |  User page
:-------------------------:|:-------------------------:
![Feed](https://user-images.githubusercontent.com/8641243/228986548-953aaf88-c9f9-4d8e-a8e8-269c293dfa1f.png) |  ![User page](https://user-images.githubusercontent.com/8641243/228986585-074ac49d-9e4f-4f31-8e36-40cdd0629fa9.png)

## Technical

### Server

#### Migrations

```bash
$ export DATABASE=postgres://postgres:password@localhost:5432/postgres?sslmode=disable
$ migrate -path migrations -database "$DATABASE" command # (See help)
```
