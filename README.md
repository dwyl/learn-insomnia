<div align="center">

# Learn Insomnia

Learn how to quickly develop and test APIs using Insomnia.

</div>

## Why?

Insomnia is an easy to use REST API client 
useful for designing and debugging API's. 

There are many reasons to use Insomnia, 
these are our top 3:

+ It has a graphical user interface 
that allows for quick testing of different routes and configs.
+ A powerful templating system that allows for complex setups 
  with different auth flows and easy chaining of requests.
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


> In case you're wondering about the _name_ of the product,
**`insomnia`** is meant to be a play on "REST" ... as in REST API. 

## Who

This guide is designed for people who have little to no experience in consuming
APIs or want to know how to do so using Insomnia.

If you already have an excellent understanding of developing REST APIs,
this guide probaly isn't for you.

## How?

### Step 0: Download and install Insomnia and our Example Server

The easiest way to do this is to install it from the official website:

https://insomnia.rest/download

![insomnia-core-download](https://user-images.githubusercontent.com/194400/88074844-deeb6a80-cb6f-11ea-9daf-ea694d4900a3.png)

Make sure you download and install **Insomnia Core**


#### Example Server

For this tutorial, we will be using an example rest server. Clone this repository,

```
git clone https://github.com/dwyl/learn-insomnia
```

We have included two example Servers, 
functionally identical, but written in different styles/languages.

If your familiar with `Elixir` and/or the MVC pattern, use the server in
`elixir_server`, or if your familiar with go OR have little/no programming 
knowledge, use the server in `go_server`.


##### Elixir Server

If your familiar with Elixir, there is a Phoenix server included in
`elixir_server`.

Presuming you have Elixir setup and installed, run this by:

```
mix deps.get
mix phx.server
```

##### Go Server

In `learn-insomnia` you should find a folder called `server`, and within that
an executable called `go_server`, on mac run this using

```
./server
```
This will run a simple Todo-list REST server on port 4000.

> If you want to double check the source code, its in `main.go`. <br />
If you're curious about teh `Go` programming language,
see: https://github.com/dwyl/learn-go <br />


### Step 1: Setup a new Workspace

When you first open Insomnia core you will see an empty workspace:

![image](https://user-images.githubusercontent.com/194400/88075291-7d77cb80-cb70-11ea-84a0-ef558e6cc11f.png)


Create a new Workspace by clicking the dropdown icon and **Create Workspace**:

![step-1-insomnia-create-workspace](https://user-images.githubusercontent.com/194400/88075760-258d9480-cb71-11ea-94cd-d2dbe0ecc53f.png)


Name it "Todo-List" or something similar:


![insomnia-name-workspace-todo-list](https://user-images.githubusercontent.com/194400/88076197-afd5f880-cb71-11ea-9c21-138ea59b828b.png)


A workspace in Insomnia is a collection of related routes and configuration.
Most people use one workspace per project/website. Within one Workspace
you there are several other ways of organising your routes:

+ **Environments**: Useful for global configuration changes: E.g. changing
  URLs from development to production or changing authentication state.

+ **Folders**: Pretty self explanatory, each folder can group routes together
  and can hold there own configuration.

### Step 2: Create our first route

Insomnia is an HTTP API client. 
To consume APIs, we must first create a route or URL.

Lets do that now.
Click on the **+** dropdown and create a **New Request**:



![insomnia-create-new-request](https://user-images.githubusercontent.com/194400/88076587-296de680-cb72-11ea-9e55-e56fe545c888.png)



Call it "List todos" or something similar 
and leave it as a GET request:

![insomnia-new-request-name](https://user-images.githubusercontent.com/194400/88076734-54f0d100-cb72-11ea-9c25-f3a2a6eb42b3.png)


Once  you've done that a URL bar should appear at the top of the Window,
lets type in the URL of the local server you ran earlier:

```
http://localhost:4000/
```

![insomnia-url-send](https://user-images.githubusercontent.com/194400/88076886-87023300-cb72-11ea-8c62-63eb97264914.png)

Click "send", and you should get a response!

![insomnia-response-3-columns](https://user-images.githubusercontent.com/194400/88077363-1b6c9580-cb73-11ea-86eb-ab01c8fb959e.png)


As we can see, Insomnia is made of three key parts:

+ **Left Pane** This contains all your routes and current Workspace and
  environments. You'll also see a `cookies` button, Insomnia keeps track
  of cookies as well!

+ **Middle Pane** This contains request information. **Take some time** to 
  investigate each of the tabs. We'll work on adding the `Body` and `auth`
  sections later, but it's good to have an idea of how it all works.

+ **Right Pane** This contains response information. Again, poke around!
  By default it shows the response body.


### Step Two: Sending data

Lets create a new route to send data to the server - POST Request!

Create a new route, called "Create TODO". Same as last time, but select
"POST" instead of "GET" in the creation modal.

![insomnia-create-todo](https://user-images.githubusercontent.com/194400/88079300-6b4c5c00-cb75-11ea-9b9e-1ff96f748bcd.png)

Once you've created the `POST` request, 
you will need to enter some data,
in order to execute the request.

The URL is the same: `http://localhost:4000/` <br />
You will need to send some data, e.g:

```json
{ "item": "hello world"}
```

![insomnia-post-request](https://user-images.githubusercontent.com/194400/88094089-e409e300-cb8a-11ea-8986-ada1ee52f880.png)

You should expect to see a **`200`** response from the server
confirming that your new todo `item` was created. 


Manually inputting the URL for each new request
is fine for two or three routes 
but once you get to 20+ it starts getting tedious. 
Then what happens if the URL changes or we want to access 
our production API? Do we need to go through and change every single URL?


#### Templating

Insomnia has great templating capability, 
so lets try it out to template in our
URL!

Click the "No environment" dropdown, then "Manage Environments":

![insomnia-manage-environments](https://user-images.githubusercontent.com/194400/88096864-3ea53e00-cb8f-11ea-9773-6062532bde3e.png)


Then, press the "+" button next to sub environments. 
Create a public environment. 

![insomnia-create-environment](https://user-images.githubusercontent.com/194400/88096984-6bf1ec00-cb8f-11ea-9061-73366a9e093a.png)

Double click the environment name ("New Environment") to edit it:

![insomnia-double-click-to-edit-environment-name](https://user-images.githubusercontent.com/194400/88097605-4fa27f00-cb90-11ea-9843-fd6d14e42b87.png)

Change it to something like "Development":

![insomnia-environment-developmet](https://user-images.githubusercontent.com/194400/88097702-78c30f80-cb90-11ea-8f56-8a21b6152d4e.png)

Lets add a host environment variable that we can use in our URL:

![insomnia-host-environment-variable](https://user-images.githubusercontent.com/194400/88097869-c475b900-cb90-11ea-9efb-0d30b74f24c6.png)

Once you've clicked "Done" to finish creating the "Development" environment
with the `"host"` variable, you will be taken back to the home view.

To _use_ the development environment,
you will need to select the "No Environment" drop down
and click on "**Use Development**":

![insomnia-use-development](https://user-images.githubusercontent.com/194400/88098016-00108300-cb91-11ea-93e0-5a047a88896d.png)


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
  slash afterwards (`{{ host }}/todo/`) so the URL will be expanded to something
  like `http://localhost:4000/todo/`

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

