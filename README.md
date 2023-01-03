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

![insomnia-core-download](https://user-images.githubusercontent.com/17494745/210361035-afac00fb-5fe0-4089-ba09-7557ceb81378.png)


Make sure you download and install **Insomnia**.


#### Example Server

For this tutorial, we will be using an example rest server. Clone this repository,

```
git clone https://github.com/dwyl/learn-insomnia
```

We have included two example servers, 
functionally identical, but written in different styles/languages.

If you're familiar with `Elixir` and/or the MVC pattern, 
use the server in `elixir_server`, 
or if your familiar with go OR have little/no programming knowledge, 
use the server in `go_server`.


##### Elixir Server

If you're familiar with Elixir, there is a Phoenix server included in
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

> If you want to double check the source code, it's in `main.go`. <br />
If you're curious about the `Go` programming language,
see: https://github.com/dwyl/learn-go <br />


### Step 1: Setup a new Workspace

When you first open Insomnia, you will see an empty workspace:


![image](https://user-images.githubusercontent.com/17494745/210362615-110606ed-f242-4518-96df-ea1383831906.png)


Create a new project by clicking the `+` icon button:

Name it "Todo List" or something similar:

![insomnia-name-workspace-todo-list](https://user-images.githubusercontent.com/17494745/210364612-2703719f-620f-4086-abfc-99d002c08bb1.png)


A project in Insomnia is a collection of related routes and configuration.
Within one project, there are several other ways of organising your routes:

+ **Environments**: Useful for global configuration changes: E.g. changing
  URLs from development to production or changing authentication state.

+ **Folders**: Pretty self explanatory, each folder can group routes together
  and can hold there own configuration.

Inside each project, you can have **design documents** and **collections of requests**.

### Step 2: Create our first route

Insomnia is an HTTP API client. 
To consume APIs, we must first create a route or URL.

Let's do that now.
Inside the project we just created,
click on the `New Collection` button.

![insomnia-create-new-request](https://user-images.githubusercontent.com/17494745/210366358-c062ebba-63cd-4e9c-a3c2-96da1b98082d.png)


Call it "Basic Requests" or something similar.
We are ready to create a new request.

![image](https://user-images.githubusercontent.com/17494745/210366597-01da5118-72a9-45a1-ba2a-5cd47fef0757.png)

Click on `New HTTP Request`.
This will create a new request called "New Request".
If you double click it, you can rename it to "List Todos".


![insomnia-new-request-name](https://user-images.githubusercontent.com/17494745/210366865-3c67ca56-533d-48e8-8bea-488dbfac5cb0.png)


Once you've done that, a URL bar should appear at the top of the window,
let's type in the URL of the local server you ran earlier:

```
http://localhost:4000/
```


Click "send", and you should get a response!

![insomnia-response-3-columns](https://user-images.githubusercontent.com/17494745/210367354-a8268c5f-49f8-4d86-a59a-18c17337a301.png)


As we can see, Insomnia is made of three key parts:

+ **Left Pane** This contains all your routes and current Project and
  environments. You'll also see a `cookies` button, Insomnia keeps track
  of cookies as well!

+ **Middle Pane** This contains request information. **Take some time** to 
  investigate each of the tabs. We'll work on adding the `Body` and `auth`
  sections later, but it's good to have an idea of how it all works.

+ **Right Pane** This contains response information. Again, poke around!
  By default it shows the response body.


### Step Two: Sending data

Let's create a new route to send data to the server - POST Request!

Create a new route, called "Create Todo". 
Click on the **+** icon and click on `HTTP Request`.

![create-new-request](https://user-images.githubusercontent.com/17494745/210367777-7a1c16d0-feab-49b7-a598-bd8754021907.png)

This will prompt you to create a new request.
Name it "Create Todo" by double clicking on the newly created request.

![change-to-post](https://user-images.githubusercontent.com/17494745/210368120-736f89ab-eb7a-4004-bc6b-8c65d7333059.png)

Click on the `GET` dropdown and change it to `POST`.

Once you've created the `POST` request, 
you will need to enter some data,
in order to execute the request.

The URL is the same: `http://localhost:4000/` <br />
You will need to send some data, e.g:

```json
{ "item": "hello world"}
```

Click on the `Body` dropdown, choose `JSON` 
and paste the text.

![dropdown](https://user-images.githubusercontent.com/17494745/210368341-db14d212-b3c1-453d-973b-2e2e199627c2.png)

If you click `Send`, 
you should expect to see a **`200`** response from the server
confirming that your new todo `item` was created. 

![post-success](https://user-images.githubusercontent.com/17494745/210376141-0b41798c-146d-4a00-83d5-0dfcb9309cc8.png)


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


![insomnia-manage-environments](https://user-images.githubusercontent.com/17494745/210369234-5e0b0eac-589c-4db2-bdb2-f27cf1e7e22c.png)


Then, press the "+" button next to sub environments. 
Create a public environment. 

![insomnia-create-environment](https://user-images.githubusercontent.com/17494745/210369447-620d508f-f418-4f12-a5b3-937d64a1069c.png)

Double click the environment name ("New Environment") to edit it:
Change it to something like "Development":

![insomnia-environment-developmet](https://user-images.githubusercontent.com/17494745/210369634-e4db7e80-a1f5-45da-9e4e-479f55b42b46.png)


Lets add a host environment variable that we can use in our URL:

```json
{"host": "http://localhost:4000"}
```

![insomnia-host-environment-variable](https://user-images.githubusercontent.com/17494745/210370043-d2075889-a96f-4773-8b33-caca3a90c0b5.png)


Once you've clicked "Close" to finish creating the "Development" environment
with the `"host"` variable, you will be taken back to the home view.

To _use_ the development environment,
you will need to select the "No Environment" drop down
and click on "**Use Development**":

![insomnia-use-development](https://user-images.githubusercontent.com/17494745/210370234-01c34993-8fce-4ebb-a2f2-74e01c1513ff.png)


Let's go back to the requests and use our new template in the URL. Make sure
your environment is not set to "Development" in the top left of the screen.

For both requests, change the URL bar to `{{ host }}`:

![insomnia-host](https://user-images.githubusercontent.com/17494745/210376271-93f6302f-d73e-474f-90d3-6c7c2b317edf.png)


#### Creating a POST request

Our server accepts JSON data, so let's create a JSON request in Insomnia.

Go to the "Body" dropdown and select JSON. 
Our `todo` list schema requires
us to specify an `item` field in our JSON, 
so let's define our request now.

Type this out in the `Body`:

```json
{
	"item": "Review PR"
}
```

Click "Send", 
the server should respond with the complete body, 
with the addition of an `id` assigned to the item:

![insomnia-post-request-review-pr](https://user-images.githubusercontent.com/17494745/210376394-fad572ea-52a4-46dc-82bd-5b2e8b397061.png)


### Step 3: Chaining Responses

Once APIs get complicated and large enough, 
most common tasks involve multiple requests 
and passing arguments and results 
from one request to the next. 
This can get tedious when requests have to be made multiple times,
or involve lots of different parameters.

Luckily, Insomnia has our back here too. 
We can use results from old requests in our new ones. 
Let's do this now.

Create a new get request called "View Item":


In the URL bar, set the url to `{{ host }}` like we did before. 
Add a slash afterwards (`{{ host }}/todo/`) 
so the URL will be expanded to something like: 
`http://localhost:4000/todo/`.

![insomnia-new-request-view-item](https://user-images.githubusercontent.com/17494745/210376564-4f3bf8fd-4d53-41d0-8854-70aa1f22d25f.png)


Press <kbd>Shift + Space</kbd> to bring up the autocomplete menu. 
Scroll down to find `response => body attribute` and select it. 
This is a **template tag**
and these are effectively *functions* that we can use to transform data
from different parts of our app.

![response-autocomplete](https://user-images.githubusercontent.com/17494745/210376665-f297ae5c-7dec-4704-b974-13511b60c7b6.png)


Insomnia shows this tag has an error:

![insomnia-response-error](https://user-images.githubusercontent.com/17494745/210375779-f8d99977-e630-4c59-b66e-b70d39557894.png)


This is because we haven't pointed it at any data yet.
so lets do this now

Click on the tag to edit it.

First, we need to select what request we want to pull data from. 
In this case, we want to use the `ID` of the item we just created,
so select `POST Add Todo`.

Then, we need to select the specific bit of data we want from that response.
Insomnia supports JSONPath - A JSON query language - so lets use that
to pick out the ID:

In `Filter`, write:

```
$.id
```

*`$` is our base response, and where selecting the `id` value*

Your Tag should look like this:

![edit-tag](https://user-images.githubusercontent.com/17494745/210376804-85600a80-6003-4c10-be98-5f9a60e9b574.png)


Press **done**, then 'Send'. 
You should get a response containing the 
Todo you just created!

![view-item](https://user-images.githubusercontent.com/17494745/210378469-bba07a59-b60e-4119-8570-859f57fa1d72.png)


Obviously this is a simple example, 
but very quickly you can put together very powerful request chains 
that transform lots of data throughout your app. 
You can also combine these tags with the environment variables
we specified earlier to make re-useable tags.

### Step 4: Authentication

Most APIs require some form of authentication to access. Insomnia has 
support for many different types of authentication. We won't go into
exploring all of these but we'll walk-through a quick example of using a bearer
token to authenticate.

Open up the environment editor again by going 
`Development -> Manage Environments`, and add another variable called "Token".
Set this to "hunter2":

![add-token-env-var](https://user-images.githubusercontent.com/17494745/210378616-0f1cca19-2f75-40d6-b8c2-37ac4601f137.png)


Create a new GET request called "View Admin Page" and set the URL to 
`{{ host }}/admin`:

Go to the auth tab and select "Bearer token". This sets the `Authorization`
header in the request. Go ahead and set that to our template, `{{ token }}`.

![admin-request](https://user-images.githubusercontent.com/17494745/210378844-0947d525-f9c5-4b7d-9367-4d1070febb70.png)

Ignore the prefix for now.

Press **send** and you should get a "Successfully Authenticated!" Response!

![auth-success](https://user-images.githubusercontent.com/17494745/210378906-1f0fb974-fe18-4e66-b21b-4a93d09e6bf2.png)

----

The beauty of the template system is that you could change your environment in 
future to "Production", declare your production values and then use your 
production servers without having to change any of the individual requests!

