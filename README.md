# README

## About

This is a basic setup for sveltekit and wails. 

It has some issues such as not finding $app. I don't know how to fix these so feel free to commit improvements.

## Maintenance
I **do not** plan to actively work on this template.

IF you have fixed something, submit a PR and I will merge it.

## Package Management

By default, this setup uses `PNPM`. If you rather use just NPM or someother JS package manager:
1. Update wails.json to use `npm install`, `npm run build`, and `npm run dev` instead.
2. Delete `pnpm-lock.yaml`
3. run `wails dev` to force the package manager again

## Live Development

To run in live development mode, run `wails dev` in the project directory. This will run a Vite development
server that will provide very fast hot reload of your frontend changes. If you want to develop in a browser
and have access to your Go methods, there is also a dev server that runs on http://localhost:5173 . Connect
to this in your browser, and you can call your Go code from devtools.

## Building

To build a redistributable, production mode package, use `wails build`.
