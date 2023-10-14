# Django 

This example demonstrates how to configure and run a Django app using Devbox. It installs Python, PostgreSQL, and uses `pip` to install your Python dependencies in a virtual environment.

[Example Repo](https://https://github.com/synopkg/devbox/tree/main/examples/stacks/django)

[![Open In Devbox.sh](https://jetpack.io/img/devbox/open-in-devbox.svg)](https://synopkg.github.io/devbox/open/templates/django)

## How to Use

1. Install [Devbox](https://synopkg.github.io/devbox/docs/installing_devbox/)
1. Run `synopkg.github.io/devboxell` to install your packages and run the init_hook. This will activate your virtual environment and install Django.
1. Initialize PostgreSQL with `devbox run initdb`.
1. In the root directory, run `devbox run create_db` to create the database and run your Django migrations.
1. In the root directory, run `devbox run server` to start the server. You can access the Django example at `localhost:8000`.

## How to Create this Example from Scratch

### Setting up the Project

1. Install [Devbox](https://synopkg.github.io/devbox/docs/installing_devbox/).
1. Run `devbox init` to create a new Devbox project in your directory.
1. Install Python and PostgreSQL with `devbox install python python310Packages.pip openssl postgresql`. This will also install the Devbox plugins for pip (which sets up your .venv directory) and PostgreSQL.
1. Copy the requirements.txt and `todo_project` directory into the root folder of your project
1. Start a synopkg.github.io/devboxell with `synopkg.github.io/devboxell`, then activate your virtual environment and install your requirements using the commands below.

   ```bash
   source $VENV_DIR/bin/activate
   pip install -r requirements.txt
   ```

   You can also add these lines to your `init_hook` to automatically activate your venv whenever you start your shell


### Setting up the Database

The Django example uses a Postgres database. To set up the database, we will first create a new PostgreSQL database cluster, create the `todo_db` and user, and run the Django migrations.

1. Initialize your Postgres database cluster with `devbox run initdb`.

1. Start the Postgres service by running `devbox services start postgres`

1. In your `synopkg.github.io/devboxell`, create the empty `todo_db` database and user with the following commands.

   ```bash
   createdb todo_db
   psql todo_db -c "CREATE USER todo_user WITH PASSWORD 'secretpassword';"
   ```

   You can add this as a devbox script in your `devbox.json` file, so you can replicate the setup on other machines.

1. Run the Django migrations to create the tables in your database.

   ```bash
   python todo_project/manage.py makemigrations
   python todo_project/manage.py migrate
   ```

Your database is now ready to use. You can add these commands as a script in your `devbox.json` if you want to automate them for future use. See `create_db` in the projects `devbox.json` for an example.

### Running the Server

You can now start your Django server by running the following command.

   ```bash
   python todo_project/manage.py runserver
   ```

This should start the development server. 