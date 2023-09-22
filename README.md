# KOLLABORATÖR

Empower Your Cyber Security Arsenal with KOLLABORATÖR!

KOLLABORATÖR isn't just a tool; it's your open-source companion for lightning-fast, secure payload hosting. Spin up a server, deploy a Docker container, create custom endpoints, and when your mission's done, delete it all – effortlessly.

Key Features:

* **User-Controlled Setup:** Create your server, install Docker, and take the reins.
* **Payload Hosting:** Define custom endpoints for secure payload delivery.
* **Efficiency Maximized:** Perfect for dynamic cybersecurity projects.
* **Open Source:** Join the community, enhance security, and shape the future.

Experience the future of agile, secure payload hosting with KOLLABORATÖR. Get started now and elevate your cybersecurity game with open-source efficiency!

## Screenshots
![screenshot-1](.screenshots/screen1.png)
![screenshot-2](.screenshots/screen2.png)

## Develop
Go project:
```
$ go build -o kollab *.go
$ ./kollab
```

Frontend project:
```
$ cd frontend
$ npm install
$ npm run dev -- --open
```

To generate static files. Omit prod if not for production environment.
```
npm run build -- --mode prod
```

## Docker
Download the appropriate image from https://github.com/dubs3c/kollab/pkgs/container/kollab.

Create a new container based on the image.
```
sudo docker run -d -p 127.0.0.1:80:80 d7086d146964
```

Visit http://127.0.0.1/mgmt/, login with `kollab:thekollab`.

## TODO
- [ ] When you delete a path, delete corresponding logs
- [ ] Randomize basic auth password on every startup, print new password in console for user
- [ ] DNS pingback functionality
- [ ] Improve frontend - store a subset of event logs in the browser
- [ ] Upload files to be served

## Contributing
Any feedback or ideas are welcome! Want to improve something? Create a pull request!

1. Fork it!
2. Create your feature branch: `git checkout -b my-new-feature`
3. Commit your changes: `git commit -am 'Add some feature'`
4. Push to the branch: `git push origin my-new-feature`
5. Submit a pull request :D
