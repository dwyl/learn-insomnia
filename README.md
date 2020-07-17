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
git clone https://github.com/dwyl/learn-insomnia
```

In `learn-insomnia` you should find a folder called `server`, and within that
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

### Step Two: Sending data

Lets create a new route to send data to the server - POST Request!

Create a new route, called "Create TODO". Same as last time, but select
"POST" instead of "GET" in the creation modal.

We could type the URL in again,
which is fine for two or three routes but once you get to 20+ it starts
getting tedious. Then what happens if the URL changes or we want to access 
our production API? Do we need to go through and change every single URL?

#### Templating

Insomnia has great templating capability, so lets try it out to template in our
URL!

Click the "No environment" dropdown, then "Manage Environments".

Then, press the "+" button next to sub environments. Create a public 
environment. Double click the environemnt name ("New Environment") to edit it,
and change it to something like "Development".

Lets add a host environment variable that we can use in our URL:

![Host Template](https://i.imgur.com/d7BPxi3.png)

Lets go back to the requests and use our new template in the URL. Make sure
your environment is not set to "Development" in the top left of the screen.

For both requests, change the URL bar to `{{ host }}`

#### Creating a POST request

Our server accepts JSON data, so lets create JSON request in Insomnia.

Go to the "Body" dropdown and select JSON. Our TODO list schema requires
us to specify an `item` field in our JSON, so lets define our request now.

Type this out in the `Body`:

```json
{
	"item": "Review PR"
}
```

Click "Send",  the server should respond with the complete body, with the 
addition of an `id` assigned to the item.

### Step 3: Chaining Responses

Once APIs get complicated and large enough, most common tasks involve
multiple requests and passing arguments and results from one request to the
next. This can get tedious when requests have to be made multiple times,
or involve lots of different parameters.

Luckily, Insomnia has our back here too. We can use results from old
requests in our new ones. Lets do this now.

Create a new get request called "View Item", then;

+ In the URL bar, set the url to `{{ host }}` like we did before. Add a 
  slash afterwards (`{{ host }}/`) so the URL will be expanded to something
  like `http://localhost:8080/`

+ Press CTRL-C to bring up the autocomplete menu. Scroll down to find
  `response => body attribute` and select it. This is a **template tag**
  and these are effectively *functions* that we can use to transform data
  from different parts of our app.

+ Insomnia shows this tag has an error, this is becuase we haven't
  pointed it at any data yet - so lets do this now.

Click on the tag to edit it.

First, we need to select what request we want to pull data from. 
In this case, we want to use the `ID` of the item we just created,
so select `POST Add Todo`.

Then, we need to select the specific bit of data we want from that response.
Insomnia supports JSONPath - A JSON query language - so lets use that
to pick out the ID.

In `Filter`, write:

```
$.id
```

*`$` is our base response, and where selecting the `id` value*

Your Tag should look like this:

![Edit Tag](https://i.imgur.com/ZRAv36j.png)

Press done, then 'Send'. You should get a response containing the 
Todo you just created!.

Ovbiously this is a simple example, but very quickly you can put
together very powerful request chains that transform lots of data throughout
your app. You can also combine these tags with the environment variables
we specified earlier to make re-useable tags.

### Step 4: Authentication

Most APIs require some form of authentication to access. Insomnia has 
support for many different types of authentication. We won't go into
exploring all of these but we'll walkthrough a quick example of using a bearer
token to authenticate.

Open up the environemnt editor again by going 
`Development -> Manage Environments`, and add another variable called "Token".
Set this to "hunter2". 

Create a new GET request called "View Admin Page" and set the URL to 
`{{ host }}/admin`. 

Go to the auth tab and select "Bearer token". This sets the `Authorization`
header in the request. Go ahead and set that to our template, `{{ token }}`.

Ignore the prefix for now.

Press send and you should get a "Successfully Authenticated!" Response!.

----

The beauty of the template system is that you could change your environment in 
future to "Production", declare your production values and then use your 
production servers without having to change any of the induvidual requests!.

