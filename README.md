<div align="center">

# Learn Insomnia
Learn how to quickly develop and test APIs using Insomnia

</div>

## Why?

Insomnia is an easy to use REST API client useful for designing and
debugging API's. 

There are many reasons to use Insomnia:

+ It has a graphical user interface that allows for quick
testing of different routes and configs.
+ A powerful templating system that allows for complex setups with different 
  auth flows and easy chaining of requests.
+ Ability to share workspaces with colleagues or teammates.

It's much easier to use than a web-browser or `cURL`, and can dramatically
increase your development velocity if you find yourself often working with
APIs - both developing and consuming.

## What?

In this guide we will use Insomnia to Authenticate to an API and make some
simple requests that will show you the benefit of using Insomnia over
`cURL` or a web-browser. Towards the end we will (optionally) look at setting
up the Insomnia workspace for the Dwyl `smart-home-security-system`

We will be using `insomnia-core` which is an API *consuming* tool. You still
have to design and build the API yourself, but you can use Insomnia to test
and consume your new API. API design itself it out of scope - if you want to
learn how to do this: https://github.com/dwyl/learn-api-design

## Who

This guide is designed for people who have little to no experience in consuming
APIs or want to know how to do so using Insomnia.

If you already have an excellent understanding of developing REST APIs,
this guide probaly isn't for you.

## How?

### Step 0: Download and install Insomnia and our Example Server

The easiest way to do this is to install it from the official website:

https://insomnia.rest/download/

Make sure you download and install **Insomnia Core**

##### Example Server

For this tutorial, we will be using an example rest server. Clone this repository,

```
git clone https://github.com/dwyl/learn-nerves
```

In `learn-nerves` you should find a folder called `server`, and within that
an executable called `server`, on mac run this using

```
./server
```
This will run a simple Todo-list REST server on port 8080.

If you want to double check the source code, its in `main.go`.

### Step 1: Setup a new Workspace

![Step 1](https://i.imgur.com/O6svSVZ.png)

Go to the top left dropdown to create a new workspace.

Name it "Todo-List" or something similar.

A workspace in Insomnia is a collection of related routes and configuration.
Most people use one workspace per project/website. Within one Workspace
you there are several other ways of organising your routes:

+ **Environments**: Useful for global configuration changes: E.g. changing
  URLs from development to production or changing authentication state.

+ **Folders**: Pretty self explanatory, each folder can group routes together
  and can hold there own configuration.

### Step 2: Create our first route

Insomnia is a HTTP API client, so to consume APIs, we must first create a route,
or a URL.

Lets do that now:

![Step 2](https://i.imgur.com/Cf75cjj.png)

Call it "List todos" or something similar and leave it as a GET request.

Once  you've done that a URL bar should appear at the top of the Window,
lets type in the URL of the local server you ran earlier:

```
http://localhost:8080/
```

Click "send", and you should get a response!

![step 2.1](https://i.imgur.com/5pmT98Q.png)

As we can see, Insomnia is made of three key parts:

+ **Left Pane** This contains all your routes and current Workspace and
  environements. You'll also see a `cookies` button, Insomnia keeps track
  of cookies as well!

+ **Middle Pane** This contains request infomation. **Take some time** to 
  ivestigate each of the tabs. We'll work on adding the `Body` and `auth`
  sections later, but it's good to have an idea of how it all works.

+ **Right Pane** This contains response infomation. Again, poke around!
  By default it shows the response body.

## Step Two: Sending data